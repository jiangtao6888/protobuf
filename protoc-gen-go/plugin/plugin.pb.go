// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.19.0-devel
// 	protoc        (unknown)
// 	go            v1.12.5
// source: github.com/golang/protobuf/protoc-gen-go/plugin/plugin.proto

package plugin_go

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	pluginpb "google.golang.org/protobuf/types/pluginpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(19 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 19)
)

// Symbols defined in public import of google/protobuf/compiler/plugin.proto.

type Version = pluginpb.Version
type CodeGeneratorRequest = pluginpb.CodeGeneratorRequest
type CodeGeneratorResponse = pluginpb.CodeGeneratorResponse
type CodeGeneratorResponse_File = pluginpb.CodeGeneratorResponse_File

var File_github_com_golang_protobuf_protoc_gen_go_plugin_plugin_proto protoreflect.FileDescriptor

var file_github_com_golang_protobuf_protoc_gen_go_plugin_plugin_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67,
	0x6f, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x3b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f,
	0x67, 0x6f, 0x50, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
}

var (
	file_github_com_golang_protobuf_protoc_gen_go_plugin_plugin_proto_rawDescOnce sync.Once
	file_github_com_golang_protobuf_protoc_gen_go_plugin_plugin_proto_rawDescData = file_github_com_golang_protobuf_protoc_gen_go_plugin_plugin_proto_rawDesc
)

func file_github_com_golang_protobuf_protoc_gen_go_plugin_plugin_proto_rawDescGZIP() []byte {
	file_github_com_golang_protobuf_protoc_gen_go_plugin_plugin_proto_rawDescOnce.Do(func() {
		file_github_com_golang_protobuf_protoc_gen_go_plugin_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_golang_protobuf_protoc_gen_go_plugin_plugin_proto_rawDescData)
	})
	return file_github_com_golang_protobuf_protoc_gen_go_plugin_plugin_proto_rawDescData
}

var file_github_com_golang_protobuf_protoc_gen_go_plugin_plugin_proto_goTypes = []interface{}{}
var file_github_com_golang_protobuf_protoc_gen_go_plugin_plugin_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_golang_protobuf_protoc_gen_go_plugin_plugin_proto_init() }
func file_github_com_golang_protobuf_protoc_gen_go_plugin_plugin_proto_init() {
	if File_github_com_golang_protobuf_protoc_gen_go_plugin_plugin_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_golang_protobuf_protoc_gen_go_plugin_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_golang_protobuf_protoc_gen_go_plugin_plugin_proto_goTypes,
		DependencyIndexes: file_github_com_golang_protobuf_protoc_gen_go_plugin_plugin_proto_depIdxs,
	}.Build()
	File_github_com_golang_protobuf_protoc_gen_go_plugin_plugin_proto = out.File
	file_github_com_golang_protobuf_protoc_gen_go_plugin_plugin_proto_rawDesc = nil
	file_github_com_golang_protobuf_protoc_gen_go_plugin_plugin_proto_goTypes = nil
	file_github_com_golang_protobuf_protoc_gen_go_plugin_plugin_proto_depIdxs = nil
}
