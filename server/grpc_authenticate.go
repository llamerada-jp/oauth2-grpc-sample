package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/go-github/github"
	"github.com/llamerada-jp/oauth2-grpc-sample/proto"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var _ proto.AuthenticateServer = (*authenticateImpl)(nil)
var authenticate *authenticateImpl

const (
	metadataSessionID = "session-id"
	sessionKeyUser    = "user"
)

type authenticateRecord struct {
	timestamp    time.Time
	authResponse *oauth2.DeviceAuthResponse
}

type authenticateImpl struct {
	proto.UnimplementedAuthenticateServer
	ctx          context.Context
	mtx          sync.Mutex
	oauth2Github *oauth2.Config
	records      map[string]*authenticateRecord
	sessionStore *sessionStore
}

func registerGRPCAuthenticate(ctx context.Context, server *grpc.Server, sessionStore *sessionStore, clientID, clientSecret string) {
	authenticate = &authenticateImpl{
		ctx: ctx,
		oauth2Github: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Endpoint:     endpoints.GitHub,
			RedirectURL:  "urn:ietf:params:oauth:grant-type:device_code",
			Scopes:       []string{"user"},
		},
		records:      make(map[string]*authenticateRecord),
		sessionStore: sessionStore,
	}

	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				authenticate.cleanupAuthResponse()
			}
		}
	}()

	proto.RegisterAuthenticateServer(server, authenticate)
}

func authenticateUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if info.Server == authenticate {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("unauthorized(metadata not found)")
	}

	sessionID := md.Get(metadataSessionID)
	if len(sessionID) != 1 {
		return nil, fmt.Errorf("unauthorized(session-id not found)")
	}

	ctxWithSession, s := authenticate.sessionStore.newContextWithSession(ctx, sessionID[0])
	if s == nil {
		return nil, fmt.Errorf("unauthorized(session not found)")
	}

	return handler(ctxWithSession, req)
}

type ssWrapper struct {
	grpc.ServerStream
	ctxWithSession context.Context
}

func (s *ssWrapper) Context() context.Context {
	return s.ctxWithSession
}

func authenticateStreamInterceptor(server interface{}, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if server == authenticate {
		return handler(server, ss)
	}

	ctx := ss.Context()
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("unauthorized(metadata not found)")
	}

	sessionID := md.Get(metadataSessionID)
	if len(sessionID) != 1 {
		return fmt.Errorf("unauthorized(session-id not found)")
	}

	ctxWithSession, s := authenticate.sessionStore.newContextWithSession(ctx, sessionID[0])
	if s == nil {
		return fmt.Errorf("unauthorized(session not found)")
	}

	ssWithSession := &ssWrapper{
		ss,
		ctxWithSession,
	}

	return handler(server, ssWithSession)
}

// Signin implements proto.SampleServer.
func (a *authenticateImpl) Signin(ctx context.Context, request *proto.SigninRequest) (*proto.SigninResponse, error) {
	authResponse, err := a.oauth2Github.DeviceAuth(ctx)
	if err != nil {
		log.Printf("error on DeviceAuth: %v", err)
		return nil, err
	}

	key := a.storeAuthResponse(authResponse)

	return &proto.SigninResponse{
		UserCode:        authResponse.UserCode,
		VerificationUri: authResponse.VerificationURI,
		SigninId:        key,
	}, nil
}

func (a *authenticateImpl) GetSessionInfo(ctx context.Context, request *proto.GetSessionInfoRequest) (*proto.GetSessionInfoResponse, error) {
	record := a.popAuthenticateRecord(request.SigninId)
	if record == nil {
		return nil, fmt.Errorf("signin process expired or invalid id")
	}

	token, err := a.oauth2Github.DeviceAccessToken(ctx, record.authResponse)
	if err != nil {
		log.Printf("error on DeviceAccessToken: %v", err)
		return nil, err
	}

	client := github.NewClient(a.oauth2Github.Client(ctx, token))
	user, _, err := client.Users.Get(ctx, "")
	if err != nil {
		log.Printf("error on Users.Get: %v", err)
		return nil, err
	}

	session, sessionID := a.sessionStore.create()
	session.set(sessionKeyUser, *user.Login)

	return &proto.GetSessionInfoResponse{
		SessionId: sessionID,
	}, nil
}

func (a *authenticateImpl) storeAuthResponse(authResponse *oauth2.DeviceAuthResponse) string {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	// generate random string for key
	key := randomString(32)
	for {
		if _, ok := a.records[key]; !ok {
			break
		}
		key = randomString(32)
	}
	a.records[key] = &authenticateRecord{
		timestamp:    time.Now(),
		authResponse: authResponse,
	}

	return key
}

func (a *authenticateImpl) popAuthenticateRecord(key string) *authenticateRecord {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	record, ok := a.records[key]
	if !ok {
		return nil
	}
	delete(a.records, key)

	return record
}

func (a *authenticateImpl) cleanupAuthResponse() {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	for key, store := range a.records {
		if time.Since(store.timestamp) > 5*time.Minute {
			delete(a.records, key)
		}
	}
}
