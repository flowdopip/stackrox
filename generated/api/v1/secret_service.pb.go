// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v4.25.3
// source: api/v1/secret_service.proto

package v1

import (
	storage "github.com/stackrox/rox/generated/storage"
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

// A list of secrets (free of scoped information)
// Next Tag: 2
type SecretList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Secrets       []*storage.Secret      `protobuf:"bytes,1,rep,name=secrets,proto3" json:"secrets,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SecretList) Reset() {
	*x = SecretList{}
	mi := &file_api_v1_secret_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SecretList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretList) ProtoMessage() {}

func (x *SecretList) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_secret_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretList.ProtoReflect.Descriptor instead.
func (*SecretList) Descriptor() ([]byte, []int) {
	return file_api_v1_secret_service_proto_rawDescGZIP(), []int{0}
}

func (x *SecretList) GetSecrets() []*storage.Secret {
	if x != nil {
		return x.Secrets
	}
	return nil
}

// A list of secrets with their relationships.
// Next Tag: 2
type ListSecretsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Secrets       []*storage.ListSecret  `protobuf:"bytes,1,rep,name=secrets,proto3" json:"secrets,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListSecretsResponse) Reset() {
	*x = ListSecretsResponse{}
	mi := &file_api_v1_secret_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSecretsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSecretsResponse) ProtoMessage() {}

func (x *ListSecretsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_secret_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSecretsResponse.ProtoReflect.Descriptor instead.
func (*ListSecretsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_secret_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListSecretsResponse) GetSecrets() []*storage.ListSecret {
	if x != nil {
		return x.Secrets
	}
	return nil
}

type CountSecretsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Count         int32                  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CountSecretsResponse) Reset() {
	*x = CountSecretsResponse{}
	mi := &file_api_v1_secret_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CountSecretsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountSecretsResponse) ProtoMessage() {}

func (x *CountSecretsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_secret_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountSecretsResponse.ProtoReflect.Descriptor instead.
func (*CountSecretsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_secret_service_proto_rawDescGZIP(), []int{2}
}

func (x *CountSecretsResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_api_v1_secret_service_proto protoreflect.FileDescriptor

var file_api_v1_secret_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76,
	0x31, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x14, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x0a, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x22, 0x44, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x07, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x22, 0x2c, 0x0a, 0x14, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x32, 0xf6, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x0f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x50, 0x0a, 0x0c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x12, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x18,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x49, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x12, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a,
	0x17, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d,
	0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x42, 0x27, 0x0a,
	0x18, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x58, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_secret_service_proto_rawDescOnce sync.Once
	file_api_v1_secret_service_proto_rawDescData = file_api_v1_secret_service_proto_rawDesc
)

func file_api_v1_secret_service_proto_rawDescGZIP() []byte {
	file_api_v1_secret_service_proto_rawDescOnce.Do(func() {
		file_api_v1_secret_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_secret_service_proto_rawDescData)
	})
	return file_api_v1_secret_service_proto_rawDescData
}

var file_api_v1_secret_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_v1_secret_service_proto_goTypes = []any{
	(*SecretList)(nil),           // 0: v1.SecretList
	(*ListSecretsResponse)(nil),  // 1: v1.ListSecretsResponse
	(*CountSecretsResponse)(nil), // 2: v1.CountSecretsResponse
	(*storage.Secret)(nil),       // 3: storage.Secret
	(*storage.ListSecret)(nil),   // 4: storage.ListSecret
	(*ResourceByID)(nil),         // 5: v1.ResourceByID
	(*RawQuery)(nil),             // 6: v1.RawQuery
}
var file_api_v1_secret_service_proto_depIdxs = []int32{
	3, // 0: v1.SecretList.secrets:type_name -> storage.Secret
	4, // 1: v1.ListSecretsResponse.secrets:type_name -> storage.ListSecret
	5, // 2: v1.SecretService.GetSecret:input_type -> v1.ResourceByID
	6, // 3: v1.SecretService.CountSecrets:input_type -> v1.RawQuery
	6, // 4: v1.SecretService.ListSecrets:input_type -> v1.RawQuery
	3, // 5: v1.SecretService.GetSecret:output_type -> storage.Secret
	2, // 6: v1.SecretService.CountSecrets:output_type -> v1.CountSecretsResponse
	1, // 7: v1.SecretService.ListSecrets:output_type -> v1.ListSecretsResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1_secret_service_proto_init() }
func file_api_v1_secret_service_proto_init() {
	if File_api_v1_secret_service_proto != nil {
		return
	}
	file_api_v1_common_proto_init()
	file_api_v1_search_service_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_secret_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_secret_service_proto_goTypes,
		DependencyIndexes: file_api_v1_secret_service_proto_depIdxs,
		MessageInfos:      file_api_v1_secret_service_proto_msgTypes,
	}.Build()
	File_api_v1_secret_service_proto = out.File
	file_api_v1_secret_service_proto_rawDesc = nil
	file_api_v1_secret_service_proto_goTypes = nil
	file_api_v1_secret_service_proto_depIdxs = nil
}
