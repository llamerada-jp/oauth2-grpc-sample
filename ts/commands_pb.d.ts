import * as jspb from 'google-protobuf'



export class UnaryRequest extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): UnaryRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UnaryRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UnaryRequest): UnaryRequest.AsObject;
  static serializeBinaryToWriter(message: UnaryRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UnaryRequest;
  static deserializeBinaryFromReader(message: UnaryRequest, reader: jspb.BinaryReader): UnaryRequest;
}

export namespace UnaryRequest {
  export type AsObject = {
    message: string,
  }
}

export class UnaryResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): UnaryResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UnaryResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UnaryResponse): UnaryResponse.AsObject;
  static serializeBinaryToWriter(message: UnaryResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UnaryResponse;
  static deserializeBinaryFromReader(message: UnaryResponse, reader: jspb.BinaryReader): UnaryResponse;
}

export namespace UnaryResponse {
  export type AsObject = {
    message: string,
  }
}

export class ServerStreamRequest extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): ServerStreamRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ServerStreamRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ServerStreamRequest): ServerStreamRequest.AsObject;
  static serializeBinaryToWriter(message: ServerStreamRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ServerStreamRequest;
  static deserializeBinaryFromReader(message: ServerStreamRequest, reader: jspb.BinaryReader): ServerStreamRequest;
}

export namespace ServerStreamRequest {
  export type AsObject = {
    message: string,
  }
}

export class ServerStreamResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): ServerStreamResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ServerStreamResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ServerStreamResponse): ServerStreamResponse.AsObject;
  static serializeBinaryToWriter(message: ServerStreamResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ServerStreamResponse;
  static deserializeBinaryFromReader(message: ServerStreamResponse, reader: jspb.BinaryReader): ServerStreamResponse;
}

export namespace ServerStreamResponse {
  export type AsObject = {
    message: string,
  }
}

