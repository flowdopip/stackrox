// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v4.25.3
// source: api/v2/pagination.proto

package v2

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

type Aggregation int32

const (
	Aggregation_UNSET Aggregation = 0
	Aggregation_COUNT Aggregation = 1
	Aggregation_MIN   Aggregation = 2
	Aggregation_MAX   Aggregation = 3
)

// Enum value maps for Aggregation.
var (
	Aggregation_name = map[int32]string{
		0: "UNSET",
		1: "COUNT",
		2: "MIN",
		3: "MAX",
	}
	Aggregation_value = map[string]int32{
		"UNSET": 0,
		"COUNT": 1,
		"MIN":   2,
		"MAX":   3,
	}
)

func (x Aggregation) Enum() *Aggregation {
	p := new(Aggregation)
	*p = x
	return p
}

func (x Aggregation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Aggregation) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v2_pagination_proto_enumTypes[0].Descriptor()
}

func (Aggregation) Type() protoreflect.EnumType {
	return &file_api_v2_pagination_proto_enumTypes[0]
}

func (x Aggregation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Aggregation.Descriptor instead.
func (Aggregation) EnumDescriptor() ([]byte, []int) {
	return file_api_v2_pagination_proto_rawDescGZIP(), []int{0}
}

type AggregateBy struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AggrFunc      Aggregation            `protobuf:"varint,1,opt,name=aggrFunc,proto3,enum=v2.Aggregation" json:"aggrFunc,omitempty"`
	Distinct      bool                   `protobuf:"varint,2,opt,name=distinct,proto3" json:"distinct,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AggregateBy) Reset() {
	*x = AggregateBy{}
	mi := &file_api_v2_pagination_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AggregateBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregateBy) ProtoMessage() {}

func (x *AggregateBy) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_pagination_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregateBy.ProtoReflect.Descriptor instead.
func (*AggregateBy) Descriptor() ([]byte, []int) {
	return file_api_v2_pagination_proto_rawDescGZIP(), []int{0}
}

func (x *AggregateBy) GetAggrFunc() Aggregation {
	if x != nil {
		return x.AggrFunc
	}
	return Aggregation_UNSET
}

func (x *AggregateBy) GetDistinct() bool {
	if x != nil {
		return x.Distinct
	}
	return false
}

type SortOption struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	Field    string                 `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Reversed bool                   `protobuf:"varint,2,opt,name=reversed,proto3" json:"reversed,omitempty"`
	// This field is under development. It is not supported on any REST APIs.
	AggregateBy   *AggregateBy `protobuf:"bytes,3,opt,name=aggregate_by,json=aggregateBy,proto3" json:"aggregate_by,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SortOption) Reset() {
	*x = SortOption{}
	mi := &file_api_v2_pagination_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SortOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortOption) ProtoMessage() {}

func (x *SortOption) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_pagination_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortOption.ProtoReflect.Descriptor instead.
func (*SortOption) Descriptor() ([]byte, []int) {
	return file_api_v2_pagination_proto_rawDescGZIP(), []int{1}
}

func (x *SortOption) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *SortOption) GetReversed() bool {
	if x != nil {
		return x.Reversed
	}
	return false
}

func (x *SortOption) GetAggregateBy() *AggregateBy {
	if x != nil {
		return x.AggregateBy
	}
	return nil
}

type Pagination struct {
	state      protoimpl.MessageState `protogen:"open.v1"`
	Limit      int32                  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset     int32                  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	SortOption *SortOption            `protobuf:"bytes,3,opt,name=sort_option,json=sortOption,proto3" json:"sort_option,omitempty"`
	// This field is under development. It is not supported on any REST APIs.
	SortOptions   []*SortOption `protobuf:"bytes,4,rep,name=sort_options,json=sortOptions,proto3" json:"sort_options,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	mi := &file_api_v2_pagination_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_pagination_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_api_v2_pagination_proto_rawDescGZIP(), []int{2}
}

func (x *Pagination) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Pagination) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Pagination) GetSortOption() *SortOption {
	if x != nil {
		return x.SortOption
	}
	return nil
}

func (x *Pagination) GetSortOptions() []*SortOption {
	if x != nil {
		return x.SortOptions
	}
	return nil
}

var File_api_v2_pagination_proto protoreflect.FileDescriptor

var file_api_v2_pagination_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x32, 0x22, 0x56, 0x0a,
	0x0b, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x2b, 0x0a, 0x08,
	0x61, 0x67, 0x67, 0x72, 0x46, 0x75, 0x6e, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f,
	0x2e, 0x76, 0x32, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x61, 0x67, 0x67, 0x72, 0x46, 0x75, 0x6e, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x63, 0x74, 0x22, 0x72, 0x0a, 0x0a, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76,
	0x65, 0x72, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x76,
	0x65, 0x72, 0x73, 0x65, 0x64, 0x12, 0x32, 0x0a, 0x0c, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x32,
	0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x42, 0x79, 0x52, 0x0b, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x42, 0x79, 0x22, 0x9e, 0x01, 0x0a, 0x0a, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x2f, 0x0a, 0x0b, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x32,
	0x2e, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x73, 0x6f, 0x72,
	0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x0c, 0x73, 0x6f, 0x72, 0x74, 0x5f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x76, 0x32, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x73,
	0x6f, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2a, 0x35, 0x0a, 0x0b, 0x41, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x4e, 0x53,
	0x45, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x01, 0x12,
	0x07, 0x0a, 0x03, 0x4d, 0x49, 0x4e, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x41, 0x58, 0x10,
	0x03, 0x42, 0x27, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x5a, 0x0b, 0x2e,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x3b, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_v2_pagination_proto_rawDescOnce sync.Once
	file_api_v2_pagination_proto_rawDescData = file_api_v2_pagination_proto_rawDesc
)

func file_api_v2_pagination_proto_rawDescGZIP() []byte {
	file_api_v2_pagination_proto_rawDescOnce.Do(func() {
		file_api_v2_pagination_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v2_pagination_proto_rawDescData)
	})
	return file_api_v2_pagination_proto_rawDescData
}

var file_api_v2_pagination_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_v2_pagination_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_v2_pagination_proto_goTypes = []any{
	(Aggregation)(0),    // 0: v2.Aggregation
	(*AggregateBy)(nil), // 1: v2.AggregateBy
	(*SortOption)(nil),  // 2: v2.SortOption
	(*Pagination)(nil),  // 3: v2.Pagination
}
var file_api_v2_pagination_proto_depIdxs = []int32{
	0, // 0: v2.AggregateBy.aggrFunc:type_name -> v2.Aggregation
	1, // 1: v2.SortOption.aggregate_by:type_name -> v2.AggregateBy
	2, // 2: v2.Pagination.sort_option:type_name -> v2.SortOption
	2, // 3: v2.Pagination.sort_options:type_name -> v2.SortOption
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_v2_pagination_proto_init() }
func file_api_v2_pagination_proto_init() {
	if File_api_v2_pagination_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v2_pagination_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v2_pagination_proto_goTypes,
		DependencyIndexes: file_api_v2_pagination_proto_depIdxs,
		EnumInfos:         file_api_v2_pagination_proto_enumTypes,
		MessageInfos:      file_api_v2_pagination_proto_msgTypes,
	}.Build()
	File_api_v2_pagination_proto = out.File
	file_api_v2_pagination_proto_rawDesc = nil
	file_api_v2_pagination_proto_goTypes = nil
	file_api_v2_pagination_proto_depIdxs = nil
}
