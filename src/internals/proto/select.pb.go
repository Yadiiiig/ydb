// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: internals/proto/select.proto

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

type SelectValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table  string    `protobuf:"bytes,1,opt,name=Table,proto3" json:"Table,omitempty"`
	Fields []string  `protobuf:"bytes,2,rep,name=Fields,proto3" json:"Fields,omitempty"`
	Values []*Values `protobuf:"bytes,3,rep,name=Values,proto3" json:"Values,omitempty"`
}

func (x *SelectValues) Reset() {
	*x = SelectValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internals_proto_select_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectValues) ProtoMessage() {}

func (x *SelectValues) ProtoReflect() protoreflect.Message {
	mi := &file_internals_proto_select_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectValues.ProtoReflect.Descriptor instead.
func (*SelectValues) Descriptor() ([]byte, []int) {
	return file_internals_proto_select_proto_rawDescGZIP(), []int{0}
}

func (x *SelectValues) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *SelectValues) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *SelectValues) GetValues() []*Values {
	if x != nil {
		return x.Values
	}
	return nil
}

type Values struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operator string `protobuf:"bytes,1,opt,name=Operator,proto3" json:"Operator,omitempty"`
	Row      string `protobuf:"bytes,2,opt,name=Row,proto3" json:"Row,omitempty"`
	Value    string `protobuf:"bytes,3,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *Values) Reset() {
	*x = Values{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internals_proto_select_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Values) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Values) ProtoMessage() {}

func (x *Values) ProtoReflect() protoreflect.Message {
	mi := &file_internals_proto_select_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Values.ProtoReflect.Descriptor instead.
func (*Values) Descriptor() ([]byte, []int) {
	return file_internals_proto_select_proto_rawDescGZIP(), []int{1}
}

func (x *Values) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *Values) GetRow() string {
	if x != nil {
		return x.Row
	}
	return ""
}

func (x *Values) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SelectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *SelectResponse) Reset() {
	*x = SelectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internals_proto_select_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectResponse) ProtoMessage() {}

func (x *SelectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internals_proto_select_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectResponse.ProtoReflect.Descriptor instead.
func (*SelectResponse) Descriptor() ([]byte, []int) {
	return file_internals_proto_select_proto_rawDescGZIP(), []int{2}
}

func (x *SelectResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_internals_proto_select_proto protoreflect.FileDescriptor

var file_internals_proto_select_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d,
	0x0a, 0x0c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x1f, 0x0a, 0x06,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x4c, 0x0a,
	0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x52, 0x6f, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x28, 0x0a, 0x0e, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x39, 0x0a, 0x06, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x12,
	0x2f, 0x0a, 0x0b, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x0d,
	0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x0f, 0x2e,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x21, 0x5a, 0x1f, 0x79, 0x61, 0x64, 0x69, 0x69, 0x69, 0x67, 0x2e, 0x64, 0x65, 0x76, 0x2f,
	0x79, 0x64, 0x62, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internals_proto_select_proto_rawDescOnce sync.Once
	file_internals_proto_select_proto_rawDescData = file_internals_proto_select_proto_rawDesc
)

func file_internals_proto_select_proto_rawDescGZIP() []byte {
	file_internals_proto_select_proto_rawDescOnce.Do(func() {
		file_internals_proto_select_proto_rawDescData = protoimpl.X.CompressGZIP(file_internals_proto_select_proto_rawDescData)
	})
	return file_internals_proto_select_proto_rawDescData
}

var file_internals_proto_select_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internals_proto_select_proto_goTypes = []interface{}{
	(*SelectValues)(nil),   // 0: SelectValues
	(*Values)(nil),         // 1: Values
	(*SelectResponse)(nil), // 2: SelectResponse
}
var file_internals_proto_select_proto_depIdxs = []int32{
	1, // 0: SelectValues.Values:type_name -> Values
	0, // 1: Select.SelectQuery:input_type -> SelectValues
	2, // 2: Select.SelectQuery:output_type -> SelectResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internals_proto_select_proto_init() }
func file_internals_proto_select_proto_init() {
	if File_internals_proto_select_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internals_proto_select_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectValues); i {
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
		file_internals_proto_select_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Values); i {
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
		file_internals_proto_select_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectResponse); i {
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
			RawDescriptor: file_internals_proto_select_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internals_proto_select_proto_goTypes,
		DependencyIndexes: file_internals_proto_select_proto_depIdxs,
		MessageInfos:      file_internals_proto_select_proto_msgTypes,
	}.Build()
	File_internals_proto_select_proto = out.File
	file_internals_proto_select_proto_rawDesc = nil
	file_internals_proto_select_proto_goTypes = nil
	file_internals_proto_select_proto_depIdxs = nil
}
