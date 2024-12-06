// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.3
// source: test/test.proto

package test

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TestClone_CloneEnum int32

const (
	TestClone_UNSET TestClone_CloneEnum = 0
	TestClone_Val1  TestClone_CloneEnum = 1
	TestClone_Val2  TestClone_CloneEnum = 2
)

// Enum value maps for TestClone_CloneEnum.
var (
	TestClone_CloneEnum_name = map[int32]string{
		0: "UNSET",
		1: "Val1",
		2: "Val2",
	}
	TestClone_CloneEnum_value = map[string]int32{
		"UNSET": 0,
		"Val1":  1,
		"Val2":  2,
	}
)

func (x TestClone_CloneEnum) Enum() *TestClone_CloneEnum {
	p := new(TestClone_CloneEnum)
	*p = x
	return p
}

func (x TestClone_CloneEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestClone_CloneEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_test_test_proto_enumTypes[0].Descriptor()
}

func (TestClone_CloneEnum) Type() protoreflect.EnumType {
	return &file_test_test_proto_enumTypes[0]
}

func (x TestClone_CloneEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestClone_CloneEnum.Descriptor instead.
func (TestClone_CloneEnum) EnumDescriptor() ([]byte, []int) {
	return file_test_test_proto_rawDescGZIP(), []int{1, 0}
}

type TestCloneSubMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Int32   int32  `protobuf:"varint,1,opt,name=int32,proto3" json:"int32,omitempty"`
	String_ string `protobuf:"bytes,2,opt,name=string,proto3" json:"string,omitempty"`
}

func (x *TestCloneSubMessage) Reset() {
	*x = TestCloneSubMessage{}
	mi := &file_test_test_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestCloneSubMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestCloneSubMessage) ProtoMessage() {}

func (x *TestCloneSubMessage) ProtoReflect() protoreflect.Message {
	mi := &file_test_test_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestCloneSubMessage.ProtoReflect.Descriptor instead.
func (*TestCloneSubMessage) Descriptor() ([]byte, []int) {
	return file_test_test_proto_rawDescGZIP(), []int{0}
}

func (x *TestCloneSubMessage) GetInt32() int32 {
	if x != nil {
		return x.Int32
	}
	return 0
}

func (x *TestCloneSubMessage) GetString_() string {
	if x != nil {
		return x.String_
	}
	return ""
}

type TestClone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntSlice    []int32                         `protobuf:"varint,1,rep,packed,name=int_slice,json=intSlice,proto3" json:"int_slice,omitempty"`
	StringSlice []string                        `protobuf:"bytes,2,rep,name=string_slice,json=stringSlice,proto3" json:"string_slice,omitempty"`
	SubMessages []*TestCloneSubMessage          `protobuf:"bytes,3,rep,name=sub_messages,json=subMessages,proto3" json:"sub_messages,omitempty"`
	MessageMap  map[string]*TestCloneSubMessage `protobuf:"bytes,4,rep,name=message_map,json=messageMap,proto3" json:"message_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	StringMap   map[string]string               `protobuf:"bytes,5,rep,name=string_map,json=stringMap,proto3" json:"string_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	EnumSlice   []TestClone_CloneEnum           `protobuf:"varint,6,rep,packed,name=enum_slice,json=enumSlice,proto3,enum=test.TestClone_CloneEnum" json:"enum_slice,omitempty"`
	Ts          *timestamppb.Timestamp          `protobuf:"bytes,7,opt,name=ts,proto3" json:"ts,omitempty"`
	// Types that are assignable to Primitive:
	//
	//	*TestClone_Int32
	//	*TestClone_String_
	//	*TestClone_Msg
	Primitive  isTestClone_Primitive `protobuf_oneof:"primitive"`
	Any        *anypb.Any            `protobuf:"bytes,11,opt,name=any,proto3" json:"any,omitempty"`
	BytesMap   map[string][]byte     `protobuf:"bytes,12,rep,name=bytes_map,json=bytesMap,proto3" json:"bytes_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	BytesSlice [][]byte              `protobuf:"bytes,13,rep,name=bytes_slice,json=bytesSlice,proto3" json:"bytes_slice,omitempty"`
	Bytes      []byte                `protobuf:"bytes,14,opt,name=bytes,proto3" json:"bytes,omitempty"`
	SubMessage *TestCloneSubMessage  `protobuf:"bytes,15,opt,name=sub_message,json=subMessage,proto3" json:"sub_message,omitempty"`
}

func (x *TestClone) Reset() {
	*x = TestClone{}
	mi := &file_test_test_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestClone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestClone) ProtoMessage() {}

func (x *TestClone) ProtoReflect() protoreflect.Message {
	mi := &file_test_test_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestClone.ProtoReflect.Descriptor instead.
func (*TestClone) Descriptor() ([]byte, []int) {
	return file_test_test_proto_rawDescGZIP(), []int{1}
}

func (x *TestClone) GetIntSlice() []int32 {
	if x != nil {
		return x.IntSlice
	}
	return nil
}

func (x *TestClone) GetStringSlice() []string {
	if x != nil {
		return x.StringSlice
	}
	return nil
}

func (x *TestClone) GetSubMessages() []*TestCloneSubMessage {
	if x != nil {
		return x.SubMessages
	}
	return nil
}

func (x *TestClone) GetMessageMap() map[string]*TestCloneSubMessage {
	if x != nil {
		return x.MessageMap
	}
	return nil
}

func (x *TestClone) GetStringMap() map[string]string {
	if x != nil {
		return x.StringMap
	}
	return nil
}

func (x *TestClone) GetEnumSlice() []TestClone_CloneEnum {
	if x != nil {
		return x.EnumSlice
	}
	return nil
}

func (x *TestClone) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (m *TestClone) GetPrimitive() isTestClone_Primitive {
	if m != nil {
		return m.Primitive
	}
	return nil
}

func (x *TestClone) GetInt32() int32 {
	if x, ok := x.GetPrimitive().(*TestClone_Int32); ok {
		return x.Int32
	}
	return 0
}

func (x *TestClone) GetString_() string {
	if x, ok := x.GetPrimitive().(*TestClone_String_); ok {
		return x.String_
	}
	return ""
}

func (x *TestClone) GetMsg() *TestCloneSubMessage {
	if x, ok := x.GetPrimitive().(*TestClone_Msg); ok {
		return x.Msg
	}
	return nil
}

func (x *TestClone) GetAny() *anypb.Any {
	if x != nil {
		return x.Any
	}
	return nil
}

func (x *TestClone) GetBytesMap() map[string][]byte {
	if x != nil {
		return x.BytesMap
	}
	return nil
}

func (x *TestClone) GetBytesSlice() [][]byte {
	if x != nil {
		return x.BytesSlice
	}
	return nil
}

func (x *TestClone) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *TestClone) GetSubMessage() *TestCloneSubMessage {
	if x != nil {
		return x.SubMessage
	}
	return nil
}

type isTestClone_Primitive interface {
	isTestClone_Primitive()
}

type TestClone_Int32 struct {
	Int32 int32 `protobuf:"varint,8,opt,name=int32,proto3,oneof"`
}

type TestClone_String_ struct {
	String_ string `protobuf:"bytes,9,opt,name=string,proto3,oneof"`
}

type TestClone_Msg struct {
	Msg *TestCloneSubMessage `protobuf:"bytes,10,opt,name=msg,proto3,oneof"`
}

func (*TestClone_Int32) isTestClone_Primitive() {}

func (*TestClone_String_) isTestClone_Primitive() {}

func (*TestClone_Msg) isTestClone_Primitive() {}

var File_test_test_proto protoreflect.FileDescriptor

var file_test_test_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x13, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6c, 0x6f, 0x6e, 0x65,
	0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0xb6, 0x07, 0x0a, 0x09, 0x54, 0x65, 0x73,
	0x74, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x5f, 0x73, 0x6c,
	0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x53, 0x6c,
	0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x6c,
	0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x53, 0x75, 0x62,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x6d, 0x61, 0x70, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4d, 0x61, 0x70, 0x12, 0x3d, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x6d, 0x61, 0x70, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x4d, 0x61, 0x70, 0x12, 0x38, 0x0a, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x6c,
	0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x6e, 0x65,
	0x45, 0x6e, 0x75, 0x6d, 0x52, 0x09, 0x65, 0x6e, 0x75, 0x6d, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12,
	0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x05, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x2d, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x53, 0x75, 0x62, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x26, 0x0a, 0x03,
	0x61, 0x6e, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x03, 0x61, 0x6e, 0x79, 0x12, 0x3a, 0x0a, 0x09, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x6d, 0x61,
	0x70, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x62, 0x79, 0x74, 0x65, 0x73, 0x4d, 0x61, 0x70,
	0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x18,
	0x0d, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x53, 0x6c, 0x69, 0x63,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x53, 0x75, 0x62,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x58, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x43, 0x6c, 0x6f, 0x6e, 0x65, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3c, 0x0a,
	0x0e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2a, 0x0a, 0x09, 0x43, 0x6c, 0x6f, 0x6e,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x56, 0x61, 0x6c, 0x31, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x56, 0x61,
	0x6c, 0x32, 0x10, 0x02, 0x42, 0x0b, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76,
	0x65, 0x42, 0x25, 0x0a, 0x16, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x5a, 0x0b, 0x2e, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x3b, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_test_proto_rawDescOnce sync.Once
	file_test_test_proto_rawDescData = file_test_test_proto_rawDesc
)

func file_test_test_proto_rawDescGZIP() []byte {
	file_test_test_proto_rawDescOnce.Do(func() {
		file_test_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_test_proto_rawDescData)
	})
	return file_test_test_proto_rawDescData
}

var file_test_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_test_test_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_test_test_proto_goTypes = []any{
	(TestClone_CloneEnum)(0),      // 0: test.TestClone.CloneEnum
	(*TestCloneSubMessage)(nil),   // 1: test.TestCloneSubMessage
	(*TestClone)(nil),             // 2: test.TestClone
	nil,                           // 3: test.TestClone.MessageMapEntry
	nil,                           // 4: test.TestClone.StringMapEntry
	nil,                           // 5: test.TestClone.BytesMapEntry
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*anypb.Any)(nil),             // 7: google.protobuf.Any
}
var file_test_test_proto_depIdxs = []int32{
	1,  // 0: test.TestClone.sub_messages:type_name -> test.TestCloneSubMessage
	3,  // 1: test.TestClone.message_map:type_name -> test.TestClone.MessageMapEntry
	4,  // 2: test.TestClone.string_map:type_name -> test.TestClone.StringMapEntry
	0,  // 3: test.TestClone.enum_slice:type_name -> test.TestClone.CloneEnum
	6,  // 4: test.TestClone.ts:type_name -> google.protobuf.Timestamp
	1,  // 5: test.TestClone.msg:type_name -> test.TestCloneSubMessage
	7,  // 6: test.TestClone.any:type_name -> google.protobuf.Any
	5,  // 7: test.TestClone.bytes_map:type_name -> test.TestClone.BytesMapEntry
	1,  // 8: test.TestClone.sub_message:type_name -> test.TestCloneSubMessage
	1,  // 9: test.TestClone.MessageMapEntry.value:type_name -> test.TestCloneSubMessage
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_test_test_proto_init() }
func file_test_test_proto_init() {
	if File_test_test_proto != nil {
		return
	}
	file_test_test_proto_msgTypes[1].OneofWrappers = []any{
		(*TestClone_Int32)(nil),
		(*TestClone_String_)(nil),
		(*TestClone_Msg)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_test_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_test_proto_goTypes,
		DependencyIndexes: file_test_test_proto_depIdxs,
		EnumInfos:         file_test_test_proto_enumTypes,
		MessageInfos:      file_test_test_proto_msgTypes,
	}.Build()
	File_test_test_proto = out.File
	file_test_test_proto_rawDesc = nil
	file_test_test_proto_goTypes = nil
	file_test_test_proto_depIdxs = nil
}
