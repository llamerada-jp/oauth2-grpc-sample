/**
 * @fileoverview gRPC-Web generated client stub for proto
 * @enhanceable
 * @public
 */
import * as grpcWeb from 'grpc-web';
import * as commands_pb from './commands_pb';
export declare class CommandsClient {
    client_: grpcWeb.AbstractClientBase;
    hostname_: string;
    credentials_: null | {
        [index: string]: string;
    };
    options_: null | {
        [index: string]: any;
    };
    constructor(hostname: string, credentials?: null | {
        [index: string]: string;
    }, options?: null | {
        [index: string]: any;
    });
    methodDescriptorUnaryRPC: any;
    unaryRPC(request: commands_pb.UnaryRequest, metadata?: grpcWeb.Metadata | null): Promise<commands_pb.UnaryResponse>;
    unaryRPC(request: commands_pb.UnaryRequest, metadata: grpcWeb.Metadata | null, callback: (err: grpcWeb.RpcError, response: commands_pb.UnaryResponse) => void): grpcWeb.ClientReadableStream<commands_pb.UnaryResponse>;
    methodDescriptorServerStreamRPC: any;
    serverStreamRPC(request: commands_pb.ServerStreamRequest, metadata?: grpcWeb.Metadata): grpcWeb.ClientReadableStream<commands_pb.ServerStreamResponse>;
}
