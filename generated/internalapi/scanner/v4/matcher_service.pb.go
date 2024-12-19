// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v4.25.3
// source: internalapi/scanner/v4/matcher_service.proto

package v4

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type GetVulnerabilitiesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	HashId        string                 `protobuf:"bytes,1,opt,name=hash_id,json=hashId,proto3" json:"hash_id,omitempty"`
	Contents      *Contents              `protobuf:"bytes,2,opt,name=contents,proto3" json:"contents,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetVulnerabilitiesRequest) Reset() {
	*x = GetVulnerabilitiesRequest{}
	mi := &file_internalapi_scanner_v4_matcher_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVulnerabilitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVulnerabilitiesRequest) ProtoMessage() {}

func (x *GetVulnerabilitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_scanner_v4_matcher_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVulnerabilitiesRequest.ProtoReflect.Descriptor instead.
func (*GetVulnerabilitiesRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_scanner_v4_matcher_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetVulnerabilitiesRequest) GetHashId() string {
	if x != nil {
		return x.HashId
	}
	return ""
}

func (x *GetVulnerabilitiesRequest) GetContents() *Contents {
	if x != nil {
		return x.Contents
	}
	return nil
}

type Metadata struct {
	state                   protoimpl.MessageState `protogen:"open.v1"`
	LastVulnerabilityUpdate *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=LastVulnerabilityUpdate,proto3" json:"LastVulnerabilityUpdate,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	mi := &file_internalapi_scanner_v4_matcher_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_scanner_v4_matcher_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_internalapi_scanner_v4_matcher_service_proto_rawDescGZIP(), []int{1}
}

func (x *Metadata) GetLastVulnerabilityUpdate() *timestamppb.Timestamp {
	if x != nil {
		return x.LastVulnerabilityUpdate
	}
	return nil
}

var File_internalapi_scanner_v4_matcher_service_proto protoreflect.FileDescriptor

var file_internalapi_scanner_v4_matcher_service_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x34, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x34,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x2f, 0x76, 0x34, 0x2f, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x66, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x68, 0x61, 0x73, 0x68, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x2e, 0x76, 0x34, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x08,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x60, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x54, 0x0a, 0x17, 0x4c, 0x61, 0x73, 0x74, 0x56, 0x75, 0x6c, 0x6e,
	0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x17, 0x4c, 0x61, 0x73, 0x74, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x32, 0xa4, 0x01, 0x0a, 0x07, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x5c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x56, 0x75, 0x6c,
	0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x25, 0x2e, 0x73,
	0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x75, 0x6c,
	0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x34,
	0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x73, 0x63,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x1d, 0x5a, 0x1b, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x34, 0x3b, 0x76, 0x34,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internalapi_scanner_v4_matcher_service_proto_rawDescOnce sync.Once
	file_internalapi_scanner_v4_matcher_service_proto_rawDescData = file_internalapi_scanner_v4_matcher_service_proto_rawDesc
)

func file_internalapi_scanner_v4_matcher_service_proto_rawDescGZIP() []byte {
	file_internalapi_scanner_v4_matcher_service_proto_rawDescOnce.Do(func() {
		file_internalapi_scanner_v4_matcher_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalapi_scanner_v4_matcher_service_proto_rawDescData)
	})
	return file_internalapi_scanner_v4_matcher_service_proto_rawDescData
}

var file_internalapi_scanner_v4_matcher_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internalapi_scanner_v4_matcher_service_proto_goTypes = []any{
	(*GetVulnerabilitiesRequest)(nil), // 0: scanner.v4.GetVulnerabilitiesRequest
	(*Metadata)(nil),                  // 1: scanner.v4.Metadata
	(*Contents)(nil),                  // 2: scanner.v4.Contents
	(*timestamppb.Timestamp)(nil),     // 3: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),             // 4: google.protobuf.Empty
	(*VulnerabilityReport)(nil),       // 5: scanner.v4.VulnerabilityReport
}
var file_internalapi_scanner_v4_matcher_service_proto_depIdxs = []int32{
	2, // 0: scanner.v4.GetVulnerabilitiesRequest.contents:type_name -> scanner.v4.Contents
	3, // 1: scanner.v4.Metadata.LastVulnerabilityUpdate:type_name -> google.protobuf.Timestamp
	0, // 2: scanner.v4.Matcher.GetVulnerabilities:input_type -> scanner.v4.GetVulnerabilitiesRequest
	4, // 3: scanner.v4.Matcher.GetMetadata:input_type -> google.protobuf.Empty
	5, // 4: scanner.v4.Matcher.GetVulnerabilities:output_type -> scanner.v4.VulnerabilityReport
	1, // 5: scanner.v4.Matcher.GetMetadata:output_type -> scanner.v4.Metadata
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internalapi_scanner_v4_matcher_service_proto_init() }
func file_internalapi_scanner_v4_matcher_service_proto_init() {
	if File_internalapi_scanner_v4_matcher_service_proto != nil {
		return
	}
	file_internalapi_scanner_v4_common_proto_init()
	file_internalapi_scanner_v4_vulnerability_report_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internalapi_scanner_v4_matcher_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internalapi_scanner_v4_matcher_service_proto_goTypes,
		DependencyIndexes: file_internalapi_scanner_v4_matcher_service_proto_depIdxs,
		MessageInfos:      file_internalapi_scanner_v4_matcher_service_proto_msgTypes,
	}.Build()
	File_internalapi_scanner_v4_matcher_service_proto = out.File
	file_internalapi_scanner_v4_matcher_service_proto_rawDesc = nil
	file_internalapi_scanner_v4_matcher_service_proto_goTypes = nil
	file_internalapi_scanner_v4_matcher_service_proto_depIdxs = nil
}
