// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v3.12.4
// source: proto/commands.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UnaryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnaryRequest) Reset() {
	*x = UnaryRequest{}
	mi := &file_proto_commands_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnaryRequest) ProtoMessage() {}

func (x *UnaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_commands_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnaryRequest.ProtoReflect.Descriptor instead.
func (*UnaryRequest) Descriptor() ([]byte, []int) {
	return file_proto_commands_proto_rawDescGZIP(), []int{0}
}

func (x *UnaryRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UnaryResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnaryResponse) Reset() {
	*x = UnaryResponse{}
	mi := &file_proto_commands_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnaryResponse) ProtoMessage() {}

func (x *UnaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_commands_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnaryResponse.ProtoReflect.Descriptor instead.
func (*UnaryResponse) Descriptor() ([]byte, []int) {
	return file_proto_commands_proto_rawDescGZIP(), []int{1}
}

func (x *UnaryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ServerStreamRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ServerStreamRequest) Reset() {
	*x = ServerStreamRequest{}
	mi := &file_proto_commands_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServerStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStreamRequest) ProtoMessage() {}

func (x *ServerStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_commands_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStreamRequest.ProtoReflect.Descriptor instead.
func (*ServerStreamRequest) Descriptor() ([]byte, []int) {
	return file_proto_commands_proto_rawDescGZIP(), []int{2}
}

func (x *ServerStreamRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ServerStreamResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ServerStreamResponse) Reset() {
	*x = ServerStreamResponse{}
	mi := &file_proto_commands_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServerStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStreamResponse) ProtoMessage() {}

func (x *ServerStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_commands_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStreamResponse.ProtoReflect.Descriptor instead.
func (*ServerStreamResponse) Descriptor() ([]byte, []int) {
	return file_proto_commands_proto_rawDescGZIP(), []int{3}
}

func (x *ServerStreamResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_commands_proto protoreflect.FileDescriptor

var file_proto_commands_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a,
	0x0c, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x55, 0x6e, 0x61, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x2f, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x30, 0x0a, 0x14, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x8f, 0x01, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x50, 0x43, 0x12, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x6e, 0x61, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0f, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x50, 0x43, 0x12, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x6c, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x64, 0x61, 0x2d,
	0x6a, 0x70, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_commands_proto_rawDescOnce sync.Once
	file_proto_commands_proto_rawDescData = file_proto_commands_proto_rawDesc
)

func file_proto_commands_proto_rawDescGZIP() []byte {
	file_proto_commands_proto_rawDescOnce.Do(func() {
		file_proto_commands_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_commands_proto_rawDescData)
	})
	return file_proto_commands_proto_rawDescData
}

var file_proto_commands_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_commands_proto_goTypes = []any{
	(*UnaryRequest)(nil),         // 0: proto.UnaryRequest
	(*UnaryResponse)(nil),        // 1: proto.UnaryResponse
	(*ServerStreamRequest)(nil),  // 2: proto.ServerStreamRequest
	(*ServerStreamResponse)(nil), // 3: proto.ServerStreamResponse
}
var file_proto_commands_proto_depIdxs = []int32{
	0, // 0: proto.Commands.UnaryRPC:input_type -> proto.UnaryRequest
	2, // 1: proto.Commands.ServerStreamRPC:input_type -> proto.ServerStreamRequest
	1, // 2: proto.Commands.UnaryRPC:output_type -> proto.UnaryResponse
	3, // 3: proto.Commands.ServerStreamRPC:output_type -> proto.ServerStreamResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_commands_proto_init() }
func file_proto_commands_proto_init() {
	if File_proto_commands_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_commands_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_commands_proto_goTypes,
		DependencyIndexes: file_proto_commands_proto_depIdxs,
		MessageInfos:      file_proto_commands_proto_msgTypes,
	}.Build()
	File_proto_commands_proto = out.File
	file_proto_commands_proto_rawDesc = nil
	file_proto_commands_proto_goTypes = nil
	file_proto_commands_proto_depIdxs = nil
}
