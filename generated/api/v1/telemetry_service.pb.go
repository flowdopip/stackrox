// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.3
// source: api/v1/telemetry_service.proto

package v1

import (
	central "github.com/stackrox/rox/generated/internalapi/central"
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

type ConfigureTelemetryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Marked as deprecated in api/v1/telemetry_service.proto.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *ConfigureTelemetryRequest) Reset() {
	*x = ConfigureTelemetryRequest{}
	mi := &file_api_v1_telemetry_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfigureTelemetryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureTelemetryRequest) ProtoMessage() {}

func (x *ConfigureTelemetryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_telemetry_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureTelemetryRequest.ProtoReflect.Descriptor instead.
func (*ConfigureTelemetryRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_telemetry_service_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Marked as deprecated in api/v1/telemetry_service.proto.
func (x *ConfigureTelemetryRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

var File_api_v1_telemetry_service_proto protoreflect.FileDescriptor

var file_api_v1_telemetry_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x76, 0x31, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x42, 0x02, 0x18, 0x01, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x32,
	0xcc, 0x02, 0x0a, 0x10, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x6b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x22, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x88, 0x02,
	0x01, 0x12, 0x7b, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x54, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a,
	0x01, 0x2a, 0x1a, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x88, 0x02, 0x01, 0x12, 0x4e,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x09, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c,
	0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x27,
	0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x58, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_v1_telemetry_service_proto_rawDescOnce sync.Once
	file_api_v1_telemetry_service_proto_rawDescData = file_api_v1_telemetry_service_proto_rawDesc
)

func file_api_v1_telemetry_service_proto_rawDescGZIP() []byte {
	file_api_v1_telemetry_service_proto_rawDescOnce.Do(func() {
		file_api_v1_telemetry_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_telemetry_service_proto_rawDescData)
	})
	return file_api_v1_telemetry_service_proto_rawDescData
}

var file_api_v1_telemetry_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_v1_telemetry_service_proto_goTypes = []any{
	(*ConfigureTelemetryRequest)(nil),      // 0: v1.ConfigureTelemetryRequest
	(*Empty)(nil),                          // 1: v1.Empty
	(*storage.TelemetryConfiguration)(nil), // 2: storage.TelemetryConfiguration
	(*central.TelemetryConfig)(nil),        // 3: central.TelemetryConfig
}
var file_api_v1_telemetry_service_proto_depIdxs = []int32{
	1, // 0: v1.TelemetryService.GetTelemetryConfiguration:input_type -> v1.Empty
	0, // 1: v1.TelemetryService.ConfigureTelemetry:input_type -> v1.ConfigureTelemetryRequest
	1, // 2: v1.TelemetryService.GetConfig:input_type -> v1.Empty
	2, // 3: v1.TelemetryService.GetTelemetryConfiguration:output_type -> storage.TelemetryConfiguration
	2, // 4: v1.TelemetryService.ConfigureTelemetry:output_type -> storage.TelemetryConfiguration
	3, // 5: v1.TelemetryService.GetConfig:output_type -> central.TelemetryConfig
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_telemetry_service_proto_init() }
func file_api_v1_telemetry_service_proto_init() {
	if File_api_v1_telemetry_service_proto != nil {
		return
	}
	file_api_v1_empty_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_telemetry_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_telemetry_service_proto_goTypes,
		DependencyIndexes: file_api_v1_telemetry_service_proto_depIdxs,
		MessageInfos:      file_api_v1_telemetry_service_proto_msgTypes,
	}.Build()
	File_api_v1_telemetry_service_proto = out.File
	file_api_v1_telemetry_service_proto_rawDesc = nil
	file_api_v1_telemetry_service_proto_goTypes = nil
	file_api_v1_telemetry_service_proto_depIdxs = nil
}
