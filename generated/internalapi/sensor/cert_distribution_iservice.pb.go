// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.3
// source: internalapi/sensor/cert_distribution_iservice.proto

package sensor

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

type FetchCertificateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceType         storage.ServiceType `protobuf:"varint,1,opt,name=service_type,json=serviceType,proto3,enum=storage.ServiceType" json:"service_type,omitempty"`
	ServiceAccountToken string              `protobuf:"bytes,2,opt,name=service_account_token,json=serviceAccountToken,proto3" json:"service_account_token,omitempty"`
}

func (x *FetchCertificateRequest) Reset() {
	*x = FetchCertificateRequest{}
	mi := &file_internalapi_sensor_cert_distribution_iservice_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchCertificateRequest) ProtoMessage() {}

func (x *FetchCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_sensor_cert_distribution_iservice_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchCertificateRequest.ProtoReflect.Descriptor instead.
func (*FetchCertificateRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_sensor_cert_distribution_iservice_proto_rawDescGZIP(), []int{0}
}

func (x *FetchCertificateRequest) GetServiceType() storage.ServiceType {
	if x != nil {
		return x.ServiceType
	}
	return storage.ServiceType(0)
}

func (x *FetchCertificateRequest) GetServiceAccountToken() string {
	if x != nil {
		return x.ServiceAccountToken
	}
	return ""
}

type FetchCertificateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PemCert string `protobuf:"bytes,1,opt,name=pem_cert,json=pemCert,proto3" json:"pem_cert,omitempty"`
	PemKey  string `protobuf:"bytes,2,opt,name=pem_key,json=pemKey,proto3" json:"pem_key,omitempty"`
}

func (x *FetchCertificateResponse) Reset() {
	*x = FetchCertificateResponse{}
	mi := &file_internalapi_sensor_cert_distribution_iservice_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchCertificateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchCertificateResponse) ProtoMessage() {}

func (x *FetchCertificateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_sensor_cert_distribution_iservice_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchCertificateResponse.ProtoReflect.Descriptor instead.
func (*FetchCertificateResponse) Descriptor() ([]byte, []int) {
	return file_internalapi_sensor_cert_distribution_iservice_proto_rawDescGZIP(), []int{1}
}

func (x *FetchCertificateResponse) GetPemCert() string {
	if x != nil {
		return x.PemCert
	}
	return ""
}

func (x *FetchCertificateResponse) GetPemKey() string {
	if x != nil {
		return x.PemKey
	}
	return ""
}

var File_internalapi_sensor_cert_distribution_iservice_proto protoreflect.FileDescriptor

var file_internalapi_sensor_cert_distribution_iservice_proto_rawDesc = []byte{
	0x0a, 0x33, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x1a, 0x1e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01,
	0x0a, 0x17, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0c, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4e, 0x0a, 0x18, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x6d, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x65, 0x6d, 0x43, 0x65, 0x72, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x70, 0x65, 0x6d, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x65, 0x6d, 0x4b, 0x65, 0x79, 0x32, 0x70, 0x0a, 0x17, 0x43, 0x65, 0x72, 0x74, 0x44, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x55, 0x0a, 0x10, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1d, 0x5a, 0x1b, 0x2e, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x3b, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internalapi_sensor_cert_distribution_iservice_proto_rawDescOnce sync.Once
	file_internalapi_sensor_cert_distribution_iservice_proto_rawDescData = file_internalapi_sensor_cert_distribution_iservice_proto_rawDesc
)

func file_internalapi_sensor_cert_distribution_iservice_proto_rawDescGZIP() []byte {
	file_internalapi_sensor_cert_distribution_iservice_proto_rawDescOnce.Do(func() {
		file_internalapi_sensor_cert_distribution_iservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalapi_sensor_cert_distribution_iservice_proto_rawDescData)
	})
	return file_internalapi_sensor_cert_distribution_iservice_proto_rawDescData
}

var file_internalapi_sensor_cert_distribution_iservice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internalapi_sensor_cert_distribution_iservice_proto_goTypes = []any{
	(*FetchCertificateRequest)(nil),  // 0: sensor.FetchCertificateRequest
	(*FetchCertificateResponse)(nil), // 1: sensor.FetchCertificateResponse
	(storage.ServiceType)(0),         // 2: storage.ServiceType
}
var file_internalapi_sensor_cert_distribution_iservice_proto_depIdxs = []int32{
	2, // 0: sensor.FetchCertificateRequest.service_type:type_name -> storage.ServiceType
	0, // 1: sensor.CertDistributionService.FetchCertificate:input_type -> sensor.FetchCertificateRequest
	1, // 2: sensor.CertDistributionService.FetchCertificate:output_type -> sensor.FetchCertificateResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internalapi_sensor_cert_distribution_iservice_proto_init() }
func file_internalapi_sensor_cert_distribution_iservice_proto_init() {
	if File_internalapi_sensor_cert_distribution_iservice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internalapi_sensor_cert_distribution_iservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internalapi_sensor_cert_distribution_iservice_proto_goTypes,
		DependencyIndexes: file_internalapi_sensor_cert_distribution_iservice_proto_depIdxs,
		MessageInfos:      file_internalapi_sensor_cert_distribution_iservice_proto_msgTypes,
	}.Build()
	File_internalapi_sensor_cert_distribution_iservice_proto = out.File
	file_internalapi_sensor_cert_distribution_iservice_proto_rawDesc = nil
	file_internalapi_sensor_cert_distribution_iservice_proto_goTypes = nil
	file_internalapi_sensor_cert_distribution_iservice_proto_depIdxs = nil
}
