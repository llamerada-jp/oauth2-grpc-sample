package main

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/llamerada-jp/oauth2-grpc-sample/proto"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func signin(ctx context.Context, conn *grpc.ClientConn) *string {
	auth := proto.NewAuthenticateClient(conn)

	// signin
	signinRes, err := auth.Signin(ctx, &proto.SigninRequest{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("URL:%s\nCode:%s\n", signinRes.VerificationUri, signinRes.UserCode)

	// get session id
	sessionInfo, err := auth.GetSessionInfo(ctx, &proto.GetSessionInfoRequest{
		SigninId: signinRes.SigninId,
	})
	if err != nil {
		panic(err)
	}

	return &sessionInfo.SessionId
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serverURL := pflag.String("server-url", "localhost:8444", "Server URL")
	sessionID := pflag.String("session-id", "", "Session ID")
	pflag.Parse()

	conn, err := grpc.NewClient(*serverURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cmds := proto.NewCommandsClient(conn)

	if len(*sessionID) == 0 {
		// try request before signin
		unaryRes, err := cmds.UnaryRPC(ctx, &proto.UnaryRequest{
			Message: "Hello before signin",
		})
		if err != nil {
			fmt.Printf("UnaryRPC error: %v\n", err)
		} else {
			fmt.Printf("UnaryRPC response: %s\n", unaryRes.Message)
		}

		// signin
		sessionID = signin(ctx, conn)
	}

	// set session id for metadata
	md := metadata.New(map[string]string{
		"session-id": *sessionID,
	})
	ctxWithMD := metadata.NewOutgoingContext(ctx, md)

	// request after signin
	unaryRes, err := cmds.UnaryRPC(ctxWithMD, &proto.UnaryRequest{
		Message: "Hello after signin",
	})
	if err != nil {
		fmt.Printf("UnaryRPC error: %v\n", err)
	} else {
		fmt.Printf("UnaryRPC response: %s\n", unaryRes.Message)
	}

	streamRes, err := cmds.ServerStreamRPC(ctxWithMD, &proto.ServerStreamRequest{
		Message: "Hello after signin",
	})
	if err != nil {
		fmt.Printf("ServerStreamRPC error: %v\n", err)
	} else {
		for {
			res, err := streamRes.Recv()
			if err != nil {
				if !errors.Is(err, io.EOF) {
					fmt.Printf("ServerStreamRPC error: %v\n", err)
				}
				break
			}
			fmt.Printf("ServerStreamRPC response: %s\n", res.Message)
		}
	}
}
