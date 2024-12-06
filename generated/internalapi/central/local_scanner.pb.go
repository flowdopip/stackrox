// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.3
// source: internalapi/central/local_scanner.proto

package central

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

type LocalScannerCertsIssueError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *LocalScannerCertsIssueError) Reset() {
	*x = LocalScannerCertsIssueError{}
	mi := &file_internalapi_central_local_scanner_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LocalScannerCertsIssueError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalScannerCertsIssueError) ProtoMessage() {}

func (x *LocalScannerCertsIssueError) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_local_scanner_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalScannerCertsIssueError.ProtoReflect.Descriptor instead.
func (*LocalScannerCertsIssueError) Descriptor() ([]byte, []int) {
	return file_internalapi_central_local_scanner_proto_rawDescGZIP(), []int{0}
}

func (x *LocalScannerCertsIssueError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type IssueLocalScannerCertsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *IssueLocalScannerCertsRequest) Reset() {
	*x = IssueLocalScannerCertsRequest{}
	mi := &file_internalapi_central_local_scanner_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IssueLocalScannerCertsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueLocalScannerCertsRequest) ProtoMessage() {}

func (x *IssueLocalScannerCertsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_local_scanner_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueLocalScannerCertsRequest.ProtoReflect.Descriptor instead.
func (*IssueLocalScannerCertsRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_central_local_scanner_proto_rawDescGZIP(), []int{1}
}

func (x *IssueLocalScannerCertsRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type IssueLocalScannerCertsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Types that are assignable to Response:
	//
	//	*IssueLocalScannerCertsResponse_Certificates
	//	*IssueLocalScannerCertsResponse_Error
	Response isIssueLocalScannerCertsResponse_Response `protobuf_oneof:"response"`
}

func (x *IssueLocalScannerCertsResponse) Reset() {
	*x = IssueLocalScannerCertsResponse{}
	mi := &file_internalapi_central_local_scanner_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IssueLocalScannerCertsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueLocalScannerCertsResponse) ProtoMessage() {}

func (x *IssueLocalScannerCertsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_central_local_scanner_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueLocalScannerCertsResponse.ProtoReflect.Descriptor instead.
func (*IssueLocalScannerCertsResponse) Descriptor() ([]byte, []int) {
	return file_internalapi_central_local_scanner_proto_rawDescGZIP(), []int{2}
}

func (x *IssueLocalScannerCertsResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (m *IssueLocalScannerCertsResponse) GetResponse() isIssueLocalScannerCertsResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *IssueLocalScannerCertsResponse) GetCertificates() *storage.TypedServiceCertificateSet {
	if x, ok := x.GetResponse().(*IssueLocalScannerCertsResponse_Certificates); ok {
		return x.Certificates
	}
	return nil
}

func (x *IssueLocalScannerCertsResponse) GetError() *LocalScannerCertsIssueError {
	if x, ok := x.GetResponse().(*IssueLocalScannerCertsResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isIssueLocalScannerCertsResponse_Response interface {
	isIssueLocalScannerCertsResponse_Response()
}

type IssueLocalScannerCertsResponse_Certificates struct {
	Certificates *storage.TypedServiceCertificateSet `protobuf:"bytes,2,opt,name=certificates,proto3,oneof"`
}

type IssueLocalScannerCertsResponse_Error struct {
	Error *LocalScannerCertsIssueError `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*IssueLocalScannerCertsResponse_Certificates) isIssueLocalScannerCertsResponse_Response() {}

func (*IssueLocalScannerCertsResponse_Error) isIssueLocalScannerCertsResponse_Response() {}

var File_internalapi_central_local_scanner_proto protoreflect.FileDescriptor

var file_internalapi_central_local_scanner_proto_rawDesc = []byte{
	0x0a, 0x27, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x65,
	0x6e, 0x74, 0x72, 0x61, 0x6c, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x73, 0x63, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x65, 0x6e, 0x74, 0x72,
	0x61, 0x6c, 0x1a, 0x1e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x37, 0x0a, 0x1b, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x63, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x73, 0x49, 0x73, 0x73, 0x75, 0x65, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3e, 0x0a, 0x1d, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x43, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0xd4, 0x01, 0x0a, 0x1e,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x43, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x49, 0x0a,
	0x0c, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61,
	0x6c, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x43, 0x65,
	0x72, 0x74, 0x73, 0x49, 0x73, 0x73, 0x75, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x3b, 0x63, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internalapi_central_local_scanner_proto_rawDescOnce sync.Once
	file_internalapi_central_local_scanner_proto_rawDescData = file_internalapi_central_local_scanner_proto_rawDesc
)

func file_internalapi_central_local_scanner_proto_rawDescGZIP() []byte {
	file_internalapi_central_local_scanner_proto_rawDescOnce.Do(func() {
		file_internalapi_central_local_scanner_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalapi_central_local_scanner_proto_rawDescData)
	})
	return file_internalapi_central_local_scanner_proto_rawDescData
}

var file_internalapi_central_local_scanner_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internalapi_central_local_scanner_proto_goTypes = []any{
	(*LocalScannerCertsIssueError)(nil),        // 0: central.LocalScannerCertsIssueError
	(*IssueLocalScannerCertsRequest)(nil),      // 1: central.IssueLocalScannerCertsRequest
	(*IssueLocalScannerCertsResponse)(nil),     // 2: central.IssueLocalScannerCertsResponse
	(*storage.TypedServiceCertificateSet)(nil), // 3: storage.TypedServiceCertificateSet
}
var file_internalapi_central_local_scanner_proto_depIdxs = []int32{
	3, // 0: central.IssueLocalScannerCertsResponse.certificates:type_name -> storage.TypedServiceCertificateSet
	0, // 1: central.IssueLocalScannerCertsResponse.error:type_name -> central.LocalScannerCertsIssueError
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internalapi_central_local_scanner_proto_init() }
func file_internalapi_central_local_scanner_proto_init() {
	if File_internalapi_central_local_scanner_proto != nil {
		return
	}
	file_internalapi_central_local_scanner_proto_msgTypes[2].OneofWrappers = []any{
		(*IssueLocalScannerCertsResponse_Certificates)(nil),
		(*IssueLocalScannerCertsResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internalapi_central_local_scanner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internalapi_central_local_scanner_proto_goTypes,
		DependencyIndexes: file_internalapi_central_local_scanner_proto_depIdxs,
		MessageInfos:      file_internalapi_central_local_scanner_proto_msgTypes,
	}.Build()
	File_internalapi_central_local_scanner_proto = out.File
	file_internalapi_central_local_scanner_proto_rawDesc = nil
	file_internalapi_central_local_scanner_proto_goTypes = nil
	file_internalapi_central_local_scanner_proto_depIdxs = nil
}
