// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: load.proto

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

type LoadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start int32 `protobuf:"varint,1,opt,name=Start,proto3" json:"Start,omitempty"`
}

func (x *LoadRequest) Reset() {
	*x = LoadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_load_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadRequest) ProtoMessage() {}

func (x *LoadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_load_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadRequest.ProtoReflect.Descriptor instead.
func (*LoadRequest) Descriptor() ([]byte, []int) {
	return file_load_proto_rawDescGZIP(), []int{0}
}

func (x *LoadRequest) GetStart() int32 {
	if x != nil {
		return x.Start
	}
	return 0
}

type LoadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Result  string `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Data    int32  `protobuf:"varint,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *LoadResponse) Reset() {
	*x = LoadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_load_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadResponse) ProtoMessage() {}

func (x *LoadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_load_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadResponse.ProtoReflect.Descriptor instead.
func (*LoadResponse) Descriptor() ([]byte, []int) {
	return file_load_proto_rawDescGZIP(), []int{1}
}

func (x *LoadResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *LoadResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *LoadResponse) GetData() int32 {
	if x != nil {
		return x.Data
	}
	return 0
}

var File_load_proto protoreflect.FileDescriptor

var file_load_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x74, 0x61, 0x6c, 0x65, 0x74, 0x6c, 0x22, 0x23, 0x0a, 0x0b, 0x4c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x22, 0x54,
	0x0a, 0x0c, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x32, 0x43, 0x0a, 0x04, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x3b, 0x0a, 0x04,
	0x57, 0x6f, 0x72, 0x6b, 0x12, 0x18, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x61, 0x6c, 0x65,
	0x74, 0x6c, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x61, 0x6c, 0x65, 0x74, 0x6c, 0x2e, 0x4c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_load_proto_rawDescOnce sync.Once
	file_load_proto_rawDescData = file_load_proto_rawDesc
)

func file_load_proto_rawDescGZIP() []byte {
	file_load_proto_rawDescOnce.Do(func() {
		file_load_proto_rawDescData = protoimpl.X.CompressGZIP(file_load_proto_rawDescData)
	})
	return file_load_proto_rawDescData
}

var file_load_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_load_proto_goTypes = []interface{}{
	(*LoadRequest)(nil),  // 0: tempotaletl.LoadRequest
	(*LoadResponse)(nil), // 1: tempotaletl.LoadResponse
}
var file_load_proto_depIdxs = []int32{
	0, // 0: tempotaletl.Load.Work:input_type -> tempotaletl.LoadRequest
	1, // 1: tempotaletl.Load.Work:output_type -> tempotaletl.LoadResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_load_proto_init() }
func file_load_proto_init() {
	if File_load_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_load_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadRequest); i {
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
		file_load_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadResponse); i {
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
			RawDescriptor: file_load_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_load_proto_goTypes,
		DependencyIndexes: file_load_proto_depIdxs,
		MessageInfos:      file_load_proto_msgTypes,
	}.Build()
	File_load_proto = out.File
	file_load_proto_rawDesc = nil
	file_load_proto_goTypes = nil
	file_load_proto_depIdxs = nil
}
