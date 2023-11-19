// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.8
// source: sso/example.proto

package sso

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

type MyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *MyRequest) Reset() {
	*x = MyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyRequest) ProtoMessage() {}

func (x *MyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyRequest.ProtoReflect.Descriptor instead.
func (*MyRequest) Descriptor() ([]byte, []int) {
	return file_sso_example_proto_rawDescGZIP(), []int{0}
}

func (x *MyRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type MyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *MyResponse) Reset() {
	*x = MyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyResponse) ProtoMessage() {}

func (x *MyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyResponse.ProtoReflect.Descriptor instead.
func (*MyResponse) Descriptor() ([]byte, []int) {
	return file_sso_example_proto_rawDescGZIP(), []int{1}
}

func (x *MyResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_sso_example_proto protoreflect.FileDescriptor

var file_sso_example_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x73, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x1f, 0x0a, 0x09, 0x4d, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x24, 0x0a, 0x0a, 0x4d, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x32, 0x3a, 0x0a, 0x09, 0x4d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a,
	0x08, 0x4d, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x0f, 0x2e, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x4d, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x4d, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33, 0x5a, 0x31,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6d, 0x61, 0x6e, 0x75,
	0x6b, 0x79, 0x61, 0x6e, 0x56, 0x4d, 0x77, 0x2f, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x42, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x73,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sso_example_proto_rawDescOnce sync.Once
	file_sso_example_proto_rawDescData = file_sso_example_proto_rawDesc
)

func file_sso_example_proto_rawDescGZIP() []byte {
	file_sso_example_proto_rawDescOnce.Do(func() {
		file_sso_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_sso_example_proto_rawDescData)
	})
	return file_sso_example_proto_rawDescData
}

var file_sso_example_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sso_example_proto_goTypes = []interface{}{
	(*MyRequest)(nil),  // 0: main.MyRequest
	(*MyResponse)(nil), // 1: main.MyResponse
}
var file_sso_example_proto_depIdxs = []int32{
	0, // 0: main.MyService.MyMethod:input_type -> main.MyRequest
	1, // 1: main.MyService.MyMethod:output_type -> main.MyResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sso_example_proto_init() }
func file_sso_example_proto_init() {
	if File_sso_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sso_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyRequest); i {
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
		file_sso_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sso_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sso_example_proto_goTypes,
		DependencyIndexes: file_sso_example_proto_depIdxs,
		MessageInfos:      file_sso_example_proto_msgTypes,
	}.Build()
	File_sso_example_proto = out.File
	file_sso_example_proto_rawDesc = nil
	file_sso_example_proto_goTypes = nil
	file_sso_example_proto_depIdxs = nil
}
