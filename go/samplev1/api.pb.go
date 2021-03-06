// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: sample/v1/api.proto

package samplev1

import (
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

// Authentication ...
type Authentication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Required bool `protobuf:"varint,1,opt,name=required,proto3" json:"required,omitempty"`
}

func (x *Authentication) Reset() {
	*x = Authentication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_v1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authentication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authentication) ProtoMessage() {}

func (x *Authentication) ProtoReflect() protoreflect.Message {
	mi := &file_sample_v1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authentication.ProtoReflect.Descriptor instead.
func (*Authentication) Descriptor() ([]byte, []int) {
	return file_sample_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *Authentication) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

// GetFooRequest ...
type GetFooRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foo string `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
}

func (x *GetFooRequest) Reset() {
	*x = GetFooRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_v1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFooRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFooRequest) ProtoMessage() {}

func (x *GetFooRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sample_v1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFooRequest.ProtoReflect.Descriptor instead.
func (*GetFooRequest) Descriptor() ([]byte, []int) {
	return file_sample_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetFooRequest) GetFoo() string {
	if x != nil {
		return x.Foo
	}
	return ""
}

// GetFooResponse ...
type GetFooResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bar string `protobuf:"bytes,1,opt,name=bar,proto3" json:"bar,omitempty"`
}

func (x *GetFooResponse) Reset() {
	*x = GetFooResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_v1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFooResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFooResponse) ProtoMessage() {}

func (x *GetFooResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sample_v1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFooResponse.ProtoReflect.Descriptor instead.
func (*GetFooResponse) Descriptor() ([]byte, []int) {
	return file_sample_v1_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetFooResponse) GetBar() string {
	if x != nil {
		return x.Bar
	}
	return ""
}

// Foo ...
type Foo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hoge string `protobuf:"bytes,1,opt,name=hoge,proto3" json:"hoge,omitempty"`
}

func (x *Foo) Reset() {
	*x = Foo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_v1_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Foo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Foo) ProtoMessage() {}

func (x *Foo) ProtoReflect() protoreflect.Message {
	mi := &file_sample_v1_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Foo.ProtoReflect.Descriptor instead.
func (*Foo) Descriptor() ([]byte, []int) {
	return file_sample_v1_api_proto_rawDescGZIP(), []int{3}
}

func (x *Foo) GetHoge() string {
	if x != nil {
		return x.Hoge
	}
	return ""
}

var file_sample_v1_api_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*Authentication)(nil),
		Field:         50000,
		Name:          "sample.v1.authentication",
		Tag:           "bytes,50000,opt,name=authentication",
		Filename:      "sample/v1/api.proto",
	},
}

// Extension fields to descriptor.MethodOptions.
var (
	// optional sample.v1.Authentication authentication = 50000;
	E_Authentication = &file_sample_v1_api_proto_extTypes[0]
)

var File_sample_v1_api_proto protoreflect.FileDescriptor

var file_sample_v1_api_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x22, 0x21, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x66, 0x6f, 0x6f, 0x22, 0x22, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x62, 0x61, 0x72, 0x22, 0x19, 0x0a, 0x03, 0x46, 0x6f, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x6f, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f,
	0x67, 0x65, 0x32, 0x52, 0x0a, 0x09, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x41, 0x50, 0x49, 0x12,
	0x45, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x12, 0x18, 0x2e, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06,
	0x82, 0xb5, 0x18, 0x02, 0x08, 0x01, 0x3a, 0x63, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x5a, 0x0b, 0x67,
	0x6f, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_sample_v1_api_proto_rawDescOnce sync.Once
	file_sample_v1_api_proto_rawDescData = file_sample_v1_api_proto_rawDesc
)

func file_sample_v1_api_proto_rawDescGZIP() []byte {
	file_sample_v1_api_proto_rawDescOnce.Do(func() {
		file_sample_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_sample_v1_api_proto_rawDescData)
	})
	return file_sample_v1_api_proto_rawDescData
}

var file_sample_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sample_v1_api_proto_goTypes = []interface{}{
	(*Authentication)(nil),           // 0: sample.v1.Authentication
	(*GetFooRequest)(nil),            // 1: sample.v1.GetFooRequest
	(*GetFooResponse)(nil),           // 2: sample.v1.GetFooResponse
	(*Foo)(nil),                      // 3: sample.v1.Foo
	(*descriptor.MethodOptions)(nil), // 4: google.protobuf.MethodOptions
}
var file_sample_v1_api_proto_depIdxs = []int32{
	4, // 0: sample.v1.authentication:extendee -> google.protobuf.MethodOptions
	0, // 1: sample.v1.authentication:type_name -> sample.v1.Authentication
	1, // 2: sample.v1.SampleAPI.GetFoo:input_type -> sample.v1.GetFooRequest
	2, // 3: sample.v1.SampleAPI.GetFoo:output_type -> sample.v1.GetFooResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sample_v1_api_proto_init() }
func file_sample_v1_api_proto_init() {
	if File_sample_v1_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sample_v1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authentication); i {
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
		file_sample_v1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFooRequest); i {
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
		file_sample_v1_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFooResponse); i {
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
		file_sample_v1_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Foo); i {
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
			RawDescriptor: file_sample_v1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 1,
			NumServices:   1,
		},
		GoTypes:           file_sample_v1_api_proto_goTypes,
		DependencyIndexes: file_sample_v1_api_proto_depIdxs,
		MessageInfos:      file_sample_v1_api_proto_msgTypes,
		ExtensionInfos:    file_sample_v1_api_proto_extTypes,
	}.Build()
	File_sample_v1_api_proto = out.File
	file_sample_v1_api_proto_rawDesc = nil
	file_sample_v1_api_proto_goTypes = nil
	file_sample_v1_api_proto_depIdxs = nil
}
