package main

import (
	"context"
	"fmt"

	"github.com/llamerada-jp/oauth2-grpc-sample/proto"
	"google.golang.org/grpc"
)

var _ proto.CommandsServer = (*commandsImpl)(nil)

type commandsImpl struct {
	proto.UnimplementedCommandsServer
}

func registerGRPCCommands(server *grpc.Server) {
	ci := &commandsImpl{}

	proto.RegisterCommandsServer(server, ci)
}

func (c *commandsImpl) UnaryRPC(ctx context.Context, request *proto.UnaryRequest) (*proto.UnaryResponse, error) {
	session := getSession(ctx)
	user, _ := session.get(sessionKeyUser)
	return &proto.UnaryResponse{
		Message: fmt.Sprintf("[%s] response for %s", user, request.Message),
	}, nil
}

func (c *commandsImpl) ServerStreamRPC(request *proto.ServerStreamRequest, stream grpc.ServerStreamingServer[proto.ServerStreamResponse]) error {
	session := getSession(stream.Context())
	user, _ := session.get(sessionKeyUser)

	for i := 0; i < 10; i++ {
		if err := stream.Send(&proto.ServerStreamResponse{
			Message: fmt.Sprintf("[%s] response for %s (%d)", user, request.Message, i),
		}); err != nil {
			return err
		}
	}
	return nil
}
