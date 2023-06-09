// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: input.proto

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

// req 、resp
type InputReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Domain   string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Content  string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Time     string `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *InputReq) Reset() {
	*x = InputReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_input_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputReq) ProtoMessage() {}

func (x *InputReq) ProtoReflect() protoreflect.Message {
	mi := &file_input_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputReq.ProtoReflect.Descriptor instead.
func (*InputReq) Descriptor() ([]byte, []int) {
	return file_input_proto_rawDescGZIP(), []int{0}
}

func (x *InputReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *InputReq) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *InputReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *InputReq) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

type InputResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InputResp) Reset() {
	*x = InputResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_input_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputResp) ProtoMessage() {}

func (x *InputResp) ProtoReflect() protoreflect.Message {
	mi := &file_input_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputResp.ProtoReflect.Descriptor instead.
func (*InputResp) Descriptor() ([]byte, []int) {
	return file_input_proto_rawDescGZIP(), []int{1}
}

type SearchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *SearchReq) Reset() {
	*x = SearchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_input_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchReq) ProtoMessage() {}

func (x *SearchReq) ProtoReflect() protoreflect.Message {
	mi := &file_input_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchReq.ProtoReflect.Descriptor instead.
func (*SearchReq) Descriptor() ([]byte, []int) {
	return file_input_proto_rawDescGZIP(), []int{2}
}

func (x *SearchReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type SearchResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*InputReq `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SearchResp) Reset() {
	*x = SearchResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_input_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResp) ProtoMessage() {}

func (x *SearchResp) ProtoReflect() protoreflect.Message {
	mi := &file_input_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResp.ProtoReflect.Descriptor instead.
func (*SearchResp) Descriptor() ([]byte, []int) {
	return file_input_proto_rawDescGZIP(), []int{3}
}

func (x *SearchResp) GetData() []*InputReq {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_input_proto protoreflect.FileDescriptor

var file_input_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x6c, 0x0a, 0x08, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22,
	0x0b, 0x0a, 0x09, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x27, 0x0a, 0x09,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x56, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x24,
	0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x0d,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_input_proto_rawDescOnce sync.Once
	file_input_proto_rawDescData = file_input_proto_rawDesc
)

func file_input_proto_rawDescGZIP() []byte {
	file_input_proto_rawDescOnce.Do(func() {
		file_input_proto_rawDescData = protoimpl.X.CompressGZIP(file_input_proto_rawDescData)
	})
	return file_input_proto_rawDescData
}

var file_input_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_input_proto_goTypes = []interface{}{
	(*InputReq)(nil),   // 0: pb.InputReq
	(*InputResp)(nil),  // 1: pb.InputResp
	(*SearchReq)(nil),  // 2: pb.SearchReq
	(*SearchResp)(nil), // 3: pb.SearchResp
}
var file_input_proto_depIdxs = []int32{
	0, // 0: pb.SearchResp.data:type_name -> pb.InputReq
	0, // 1: pb.input.input:input_type -> pb.InputReq
	2, // 2: pb.input.search:input_type -> pb.SearchReq
	1, // 3: pb.input.input:output_type -> pb.InputResp
	3, // 4: pb.input.search:output_type -> pb.SearchResp
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_input_proto_init() }
func file_input_proto_init() {
	if File_input_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_input_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputReq); i {
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
		file_input_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputResp); i {
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
		file_input_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchReq); i {
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
		file_input_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResp); i {
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
			RawDescriptor: file_input_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_input_proto_goTypes,
		DependencyIndexes: file_input_proto_depIdxs,
		MessageInfos:      file_input_proto_msgTypes,
	}.Build()
	File_input_proto = out.File
	file_input_proto_rawDesc = nil
	file_input_proto_goTypes = nil
	file_input_proto_depIdxs = nil
}
