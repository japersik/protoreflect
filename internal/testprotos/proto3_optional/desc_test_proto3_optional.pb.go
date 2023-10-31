// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.22.0
// source: proto3_optional/desc_test_proto3_optional.proto

package testprotos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MessageWithOptionalFields struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foo *string `protobuf:"bytes,1,opt,name=foo,proto3,oneof" json:"foo,omitempty"`
	Bar *int64  `protobuf:"varint,2,opt,name=bar,proto3,oneof" json:"bar,omitempty"`
}

func (x *MessageWithOptionalFields) Reset() {
	*x = MessageWithOptionalFields{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto3_optional_desc_test_proto3_optional_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithOptionalFields) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithOptionalFields) ProtoMessage() {}

func (x *MessageWithOptionalFields) ProtoReflect() protoreflect.Message {
	mi := &file_proto3_optional_desc_test_proto3_optional_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithOptionalFields.ProtoReflect.Descriptor instead.
func (*MessageWithOptionalFields) Descriptor() ([]byte, []int) {
	return file_proto3_optional_desc_test_proto3_optional_proto_rawDescGZIP(), []int{0}
}

func (x *MessageWithOptionalFields) GetFoo() string {
	if x != nil && x.Foo != nil {
		return *x.Foo
	}
	return ""
}

func (x *MessageWithOptionalFields) GetBar() int64 {
	if x != nil && x.Bar != nil {
		return *x.Bar
	}
	return 0
}

var file_proto3_optional_desc_test_proto3_optional_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         44444,
		Name:          "some_custom_options",
		Tag:           "bytes,44444,opt,name=some_custom_options",
		Filename:      "proto3_optional/desc_test_proto3_optional.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional string some_custom_options = 44444;
	E_SomeCustomOptions = &file_proto3_optional_desc_test_proto3_optional_proto_extTypes[0]
)

var File_proto3_optional_desc_test_proto3_optional_proto protoreflect.FileDescriptor

var file_proto3_optional_desc_test_proto3_optional_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x19, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69,
	0x74, 0x68, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x12, 0x15, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x03, 0x66, 0x6f, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x62, 0x61, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x03, 0x62, 0x61, 0x72, 0x88, 0x01, 0x01, 0x42, 0x06,
	0x0a, 0x04, 0x5f, 0x66, 0x6f, 0x6f, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x62, 0x61, 0x72, 0x3a, 0x54,
	0x0a, 0x13, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9c, 0xdb, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x73, 0x6f, 0x6d, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x88, 0x01, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6a, 0x68, 0x75, 0x6d, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x65,
	0x66, 0x6c, 0x65, 0x63, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto3_optional_desc_test_proto3_optional_proto_rawDescOnce sync.Once
	file_proto3_optional_desc_test_proto3_optional_proto_rawDescData = file_proto3_optional_desc_test_proto3_optional_proto_rawDesc
)

func file_proto3_optional_desc_test_proto3_optional_proto_rawDescGZIP() []byte {
	file_proto3_optional_desc_test_proto3_optional_proto_rawDescOnce.Do(func() {
		file_proto3_optional_desc_test_proto3_optional_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto3_optional_desc_test_proto3_optional_proto_rawDescData)
	})
	return file_proto3_optional_desc_test_proto3_optional_proto_rawDescData
}

var file_proto3_optional_desc_test_proto3_optional_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto3_optional_desc_test_proto3_optional_proto_goTypes = []interface{}{
	(*MessageWithOptionalFields)(nil),   // 0: MessageWithOptionalFields
	(*descriptorpb.MessageOptions)(nil), // 1: google.protobuf.MessageOptions
}
var file_proto3_optional_desc_test_proto3_optional_proto_depIdxs = []int32{
	1, // 0: some_custom_options:extendee -> google.protobuf.MessageOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto3_optional_desc_test_proto3_optional_proto_init() }
func file_proto3_optional_desc_test_proto3_optional_proto_init() {
	if File_proto3_optional_desc_test_proto3_optional_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto3_optional_desc_test_proto3_optional_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithOptionalFields); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_proto3_optional_desc_test_proto3_optional_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto3_optional_desc_test_proto3_optional_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_proto3_optional_desc_test_proto3_optional_proto_goTypes,
		DependencyIndexes: file_proto3_optional_desc_test_proto3_optional_proto_depIdxs,
		MessageInfos:      file_proto3_optional_desc_test_proto3_optional_proto_msgTypes,
		ExtensionInfos:    file_proto3_optional_desc_test_proto3_optional_proto_extTypes,
	}.Build()
	File_proto3_optional_desc_test_proto3_optional_proto = out.File
	file_proto3_optional_desc_test_proto3_optional_proto_rawDesc = nil
	file_proto3_optional_desc_test_proto3_optional_proto_goTypes = nil
	file_proto3_optional_desc_test_proto3_optional_proto_depIdxs = nil
}
