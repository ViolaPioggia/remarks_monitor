// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: MapReduce.proto

package pb

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

type GetMapWorkReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MNum  string   `protobuf:"bytes,1,opt,name=MNum,proto3" json:"MNum,omitempty"`
	Paths []string `protobuf:"bytes,2,rep,name=paths,proto3" json:"paths,omitempty"`
}

func (x *GetMapWorkReq) Reset() {
	*x = GetMapWorkReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMapWorkReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMapWorkReq) ProtoMessage() {}

func (x *GetMapWorkReq) ProtoReflect() protoreflect.Message {
	mi := &file_map_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMapWorkReq.ProtoReflect.Descriptor instead.
func (*GetMapWorkReq) Descriptor() ([]byte, []int) {
	return file_map_proto_rawDescGZIP(), []int{0}
}

func (x *GetMapWorkReq) GetMNum() string {
	if x != nil {
		return x.MNum
	}
	return ""
}

func (x *GetMapWorkReq) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

type GetMapWorkResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Director string `protobuf:"bytes,1,opt,name=director,proto3" json:"director,omitempty"`
}

func (x *GetMapWorkResp) Reset() {
	*x = GetMapWorkResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMapWorkResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMapWorkResp) ProtoMessage() {}

func (x *GetMapWorkResp) ProtoReflect() protoreflect.Message {
	mi := &file_map_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMapWorkResp.ProtoReflect.Descriptor instead.
func (*GetMapWorkResp) Descriptor() ([]byte, []int) {
	return file_map_proto_rawDescGZIP(), []int{1}
}

func (x *GetMapWorkResp) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

var File_map_proto protoreflect.FileDescriptor

var file_map_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22,
	0x39, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x70, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71,
	0x12, 0x12, 0x0a, 0x04, 0x4d, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4d, 0x4e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x22, 0x2c, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x70, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x3b, 0x0a, 0x07, 0x4d, 0x61, 0x70, 0x57,
	0x6f, 0x72, 0x6b, 0x12, 0x30, 0x0a, 0x07, 0x4d, 0x61, 0x70, 0x57, 0x6f, 0x72, 0x6b, 0x12, 0x11,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x70, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65,
	0x71, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x70, 0x57, 0x6f, 0x72,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_map_proto_rawDescOnce sync.Once
	file_map_proto_rawDescData = file_map_proto_rawDesc
)

func file_map_proto_rawDescGZIP() []byte {
	file_map_proto_rawDescOnce.Do(func() {
		file_map_proto_rawDescData = protoimpl.X.CompressGZIP(file_map_proto_rawDescData)
	})
	return file_map_proto_rawDescData
}

var file_map_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_map_proto_goTypes = []interface{}{
	(*GetMapWorkReq)(nil),  // 0: pb.GetMapWorkReq
	(*GetMapWorkResp)(nil), // 1: pb.GetMapWorkResp
}
var file_map_proto_depIdxs = []int32{
	0, // 0: pb.MapWork.MapWork:input_type -> pb.GetMapWorkReq
	1, // 1: pb.MapWork.MapWork:output_type -> pb.GetMapWorkResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_map_proto_init() }
func file_map_proto_init() {
	if File_map_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_map_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMapWorkReq); i {
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
		file_map_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMapWorkResp); i {
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
			RawDescriptor: file_map_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_map_proto_goTypes,
		DependencyIndexes: file_map_proto_depIdxs,
		MessageInfos:      file_map_proto_msgTypes,
	}.Build()
	File_map_proto = out.File
	file_map_proto_rawDesc = nil
	file_map_proto_goTypes = nil
	file_map_proto_depIdxs = nil
}
