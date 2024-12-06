// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.3
// source: storage/api_token.proto

package storage

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type TokenMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk"` // @gotags: sql:"pk"
	Name       string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Roles      []string               `protobuf:"bytes,7,rep,name=roles,proto3" json:"roles,omitempty"`
	IssuedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	Expiration *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=expiration,proto3" json:"expiration,omitempty" search:"Expiration,store"` // @gotags: search:"Expiration,store"
	Revoked    bool                   `protobuf:"varint,6,opt,name=revoked,proto3" json:"revoked,omitempty" search:"Revoked,store"`      // @gotags: search:"Revoked,store"
	// Deprecated: Marked as deprecated in storage/api_token.proto.
	Role string `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *TokenMetadata) Reset() {
	*x = TokenMetadata{}
	mi := &file_storage_api_token_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TokenMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenMetadata) ProtoMessage() {}

func (x *TokenMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_storage_api_token_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenMetadata.ProtoReflect.Descriptor instead.
func (*TokenMetadata) Descriptor() ([]byte, []int) {
	return file_storage_api_token_proto_rawDescGZIP(), []int{0}
}

func (x *TokenMetadata) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TokenMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TokenMetadata) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *TokenMetadata) GetIssuedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.IssuedAt
	}
	return nil
}

func (x *TokenMetadata) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

func (x *TokenMetadata) GetRevoked() bool {
	if x != nil {
		return x.Revoked
	}
	return false
}

// Deprecated: Marked as deprecated in storage/api_token.proto.
func (x *TokenMetadata) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

var File_storage_api_token_proto protoreflect.FileDescriptor

var file_storage_api_token_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x01, 0x0a, 0x0d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12,
	0x37, 0x0a, 0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x12, 0x16,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01,
	0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x2e, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x3b, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_api_token_proto_rawDescOnce sync.Once
	file_storage_api_token_proto_rawDescData = file_storage_api_token_proto_rawDesc
)

func file_storage_api_token_proto_rawDescGZIP() []byte {
	file_storage_api_token_proto_rawDescOnce.Do(func() {
		file_storage_api_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_api_token_proto_rawDescData)
	})
	return file_storage_api_token_proto_rawDescData
}

var file_storage_api_token_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_storage_api_token_proto_goTypes = []any{
	(*TokenMetadata)(nil),         // 0: storage.TokenMetadata
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_storage_api_token_proto_depIdxs = []int32{
	1, // 0: storage.TokenMetadata.issued_at:type_name -> google.protobuf.Timestamp
	1, // 1: storage.TokenMetadata.expiration:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_storage_api_token_proto_init() }
func file_storage_api_token_proto_init() {
	if File_storage_api_token_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_api_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_api_token_proto_goTypes,
		DependencyIndexes: file_storage_api_token_proto_depIdxs,
		MessageInfos:      file_storage_api_token_proto_msgTypes,
	}.Build()
	File_storage_api_token_proto = out.File
	file_storage_api_token_proto_rawDesc = nil
	file_storage_api_token_proto_goTypes = nil
	file_storage_api_token_proto_depIdxs = nil
}
