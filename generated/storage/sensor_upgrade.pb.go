// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v4.25.3
// source: storage/sensor_upgrade.proto

package storage

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

// SensorUpgradeConfig encapsulates configuration relevant to sensor auto-upgrades.
type SensorUpgradeConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Whether to automatically trigger upgrades for out-of-date sensors.
	EnableAutoUpgrade bool `protobuf:"varint,1,opt,name=enable_auto_upgrade,json=enableAutoUpgrade,proto3" json:"enable_auto_upgrade,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *SensorUpgradeConfig) Reset() {
	*x = SensorUpgradeConfig{}
	mi := &file_storage_sensor_upgrade_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SensorUpgradeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SensorUpgradeConfig) ProtoMessage() {}

func (x *SensorUpgradeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_storage_sensor_upgrade_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SensorUpgradeConfig.ProtoReflect.Descriptor instead.
func (*SensorUpgradeConfig) Descriptor() ([]byte, []int) {
	return file_storage_sensor_upgrade_proto_rawDescGZIP(), []int{0}
}

func (x *SensorUpgradeConfig) GetEnableAutoUpgrade() bool {
	if x != nil {
		return x.EnableAutoUpgrade
	}
	return false
}

var File_storage_sensor_upgrade_proto protoreflect.FileDescriptor

var file_storage_sensor_upgrade_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x5f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0x45, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2e,
	0x0a, 0x13, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x75, 0x70,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x42, 0x2e,
	0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5a, 0x11, 0x2e, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_sensor_upgrade_proto_rawDescOnce sync.Once
	file_storage_sensor_upgrade_proto_rawDescData = file_storage_sensor_upgrade_proto_rawDesc
)

func file_storage_sensor_upgrade_proto_rawDescGZIP() []byte {
	file_storage_sensor_upgrade_proto_rawDescOnce.Do(func() {
		file_storage_sensor_upgrade_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_sensor_upgrade_proto_rawDescData)
	})
	return file_storage_sensor_upgrade_proto_rawDescData
}

var file_storage_sensor_upgrade_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_storage_sensor_upgrade_proto_goTypes = []any{
	(*SensorUpgradeConfig)(nil), // 0: storage.SensorUpgradeConfig
}
var file_storage_sensor_upgrade_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_storage_sensor_upgrade_proto_init() }
func file_storage_sensor_upgrade_proto_init() {
	if File_storage_sensor_upgrade_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_sensor_upgrade_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_sensor_upgrade_proto_goTypes,
		DependencyIndexes: file_storage_sensor_upgrade_proto_depIdxs,
		MessageInfos:      file_storage_sensor_upgrade_proto_msgTypes,
	}.Build()
	File_storage_sensor_upgrade_proto = out.File
	file_storage_sensor_upgrade_proto_rawDesc = nil
	file_storage_sensor_upgrade_proto_goTypes = nil
	file_storage_sensor_upgrade_proto_depIdxs = nil
}
