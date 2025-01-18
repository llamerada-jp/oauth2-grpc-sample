package main

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type serverSecret struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
}

func main() {
	certFile := pflag.String("cert-file", "localhost.crt", "TLS certificate file")
	keyFile := pflag.String("key-file", "localhost.key", "TLS key file")
	secretFile := pflag.String("secret", "secret.json", "Secret file")
	pflag.Parse()

	secret, err := readSecret(*secretFile)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// setup session
	sessionStore := newSessionStore(ctx, 10*time.Minute)

	// grpc server
	listener, err := net.Listen("tcp", ":8444")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(authenticateUnaryInterceptor),
		grpc.StreamInterceptor(authenticateStreamInterceptor),
	)

	registerGRPCAuthenticate(ctx, grpcServer, sessionStore, secret.ClientID, secret.ClientSecret)
	registerGRPCCommands(grpcServer)

	reflection.Register(grpcServer)

	go func() {
		grpcServer.Serve(listener)
	}()

	// http & grpc server
	// use grpcweb instead of proxy to simplify startup
	wrapGrpc := grpcweb.WrapServer(grpcServer)
	httpHandler := newHTTPHandler(ctx, sessionStore, secret.ClientID, secret.ClientSecret)
	httpServer := &http.Server{
		Addr: ":8443",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if wrapGrpc.IsGrpcWebRequest(r) {
				wrapGrpc.ServeHTTP(w, r)
			} else {
				httpHandler.ServeHTTP(w, r)
			}
		}),
	}

	go func() {
		err = httpServer.ListenAndServeTLS(*certFile, *keyFile)
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	grpcServer.GracefulStop()
	httpServer.Shutdown(ctx)
}

func readSecret(file string) (*serverSecret, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var secret serverSecret
	err = json.Unmarshal(b, &secret)
	if err != nil {
		return nil, err
	}

	return &secret, nil
}
