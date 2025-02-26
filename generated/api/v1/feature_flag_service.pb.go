// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v4.25.3
// source: api/v1/feature_flag_service.proto

package v1

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

type FeatureFlag struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	EnvVar        string                 `protobuf:"bytes,2,opt,name=env_var,json=envVar,proto3" json:"env_var,omitempty"`
	Enabled       bool                   `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FeatureFlag) Reset() {
	*x = FeatureFlag{}
	mi := &file_api_v1_feature_flag_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeatureFlag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureFlag) ProtoMessage() {}

func (x *FeatureFlag) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_feature_flag_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureFlag.ProtoReflect.Descriptor instead.
func (*FeatureFlag) Descriptor() ([]byte, []int) {
	return file_api_v1_feature_flag_service_proto_rawDescGZIP(), []int{0}
}

func (x *FeatureFlag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FeatureFlag) GetEnvVar() string {
	if x != nil {
		return x.EnvVar
	}
	return ""
}

func (x *FeatureFlag) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type GetFeatureFlagsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FeatureFlags  []*FeatureFlag         `protobuf:"bytes,1,rep,name=feature_flags,json=featureFlags,proto3" json:"feature_flags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetFeatureFlagsResponse) Reset() {
	*x = GetFeatureFlagsResponse{}
	mi := &file_api_v1_feature_flag_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFeatureFlagsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeatureFlagsResponse) ProtoMessage() {}

func (x *GetFeatureFlagsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_feature_flag_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeatureFlagsResponse.ProtoReflect.Descriptor instead.
func (*GetFeatureFlagsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_feature_flag_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetFeatureFlagsResponse) GetFeatureFlags() []*FeatureFlag {
	if x != nil {
		return x.FeatureFlags
	}
	return nil
}

var File_api_v1_feature_flag_service_proto protoreflect.FileDescriptor

var file_api_v1_feature_flag_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x5f, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x0b, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x65, 0x6e, 0x76, 0x5f, 0x76, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65,
	0x6e, 0x76, 0x56, 0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4a,
	0x04, 0x08, 0x04, 0x10, 0x05, 0x22, 0x4f, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x0d, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x66, 0x6c, 0x61, 0x67,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x0c, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x32, 0x69, 0x0a, 0x12, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x46, 0x6c, 0x61, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12,
	0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12,
	0x10, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x66, 0x6c, 0x61, 0x67,
	0x73, 0x42, 0x27, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x0b, 0x2e,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x58, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_feature_flag_service_proto_rawDescOnce sync.Once
	file_api_v1_feature_flag_service_proto_rawDescData = file_api_v1_feature_flag_service_proto_rawDesc
)

func file_api_v1_feature_flag_service_proto_rawDescGZIP() []byte {
	file_api_v1_feature_flag_service_proto_rawDescOnce.Do(func() {
		file_api_v1_feature_flag_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_feature_flag_service_proto_rawDescData)
	})
	return file_api_v1_feature_flag_service_proto_rawDescData
}

var file_api_v1_feature_flag_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_v1_feature_flag_service_proto_goTypes = []any{
	(*FeatureFlag)(nil),             // 0: v1.FeatureFlag
	(*GetFeatureFlagsResponse)(nil), // 1: v1.GetFeatureFlagsResponse
	(*Empty)(nil),                   // 2: v1.Empty
}
var file_api_v1_feature_flag_service_proto_depIdxs = []int32{
	0, // 0: v1.GetFeatureFlagsResponse.feature_flags:type_name -> v1.FeatureFlag
	2, // 1: v1.FeatureFlagService.GetFeatureFlags:input_type -> v1.Empty
	1, // 2: v1.FeatureFlagService.GetFeatureFlags:output_type -> v1.GetFeatureFlagsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_v1_feature_flag_service_proto_init() }
func file_api_v1_feature_flag_service_proto_init() {
	if File_api_v1_feature_flag_service_proto != nil {
		return
	}
	file_api_v1_empty_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_feature_flag_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_feature_flag_service_proto_goTypes,
		DependencyIndexes: file_api_v1_feature_flag_service_proto_depIdxs,
		MessageInfos:      file_api_v1_feature_flag_service_proto_msgTypes,
	}.Build()
	File_api_v1_feature_flag_service_proto = out.File
	file_api_v1_feature_flag_service_proto_rawDesc = nil
	file_api_v1_feature_flag_service_proto_goTypes = nil
	file_api_v1_feature_flag_service_proto_depIdxs = nil
}
