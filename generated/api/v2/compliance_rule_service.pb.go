// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v4.25.3
// source: api/v2/compliance_rule_service.proto

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

type RuleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RuleName      string                 `protobuf:"bytes,1,opt,name=rule_name,json=ruleName,proto3" json:"rule_name,omitempty"`
	Query         *RawQuery              `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RuleRequest) Reset() {
	*x = RuleRequest{}
	mi := &file_api_v2_compliance_rule_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleRequest) ProtoMessage() {}

func (x *RuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_compliance_rule_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleRequest.ProtoReflect.Descriptor instead.
func (*RuleRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_compliance_rule_service_proto_rawDescGZIP(), []int{0}
}

func (x *RuleRequest) GetRuleName() string {
	if x != nil {
		return x.RuleName
	}
	return ""
}

func (x *RuleRequest) GetQuery() *RawQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

var File_api_v2_compliance_rule_service_proto protoreflect.FileDescriptor

var file_api_v2_compliance_rule_service_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x32, 0x1a, 0x1e, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x32, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x0b, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x22, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x76, 0x32, 0x2e, 0x52, 0x61, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x32, 0x82, 0x01, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x75,
	0x6c, 0x65, 0x12, 0x0f, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x12,
	0x27, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2f,
	0x72, 0x75, 0x6c, 0x65, 0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2f, 0x7b, 0x72, 0x75,
	0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x42, 0x27, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x3b, 0x76,
	0x32, 0x58, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v2_compliance_rule_service_proto_rawDescOnce sync.Once
	file_api_v2_compliance_rule_service_proto_rawDescData = file_api_v2_compliance_rule_service_proto_rawDesc
)

func file_api_v2_compliance_rule_service_proto_rawDescGZIP() []byte {
	file_api_v2_compliance_rule_service_proto_rawDescOnce.Do(func() {
		file_api_v2_compliance_rule_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v2_compliance_rule_service_proto_rawDescData)
	})
	return file_api_v2_compliance_rule_service_proto_rawDescData
}

var file_api_v2_compliance_rule_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_v2_compliance_rule_service_proto_goTypes = []any{
	(*RuleRequest)(nil),    // 0: v2.RuleRequest
	(*RawQuery)(nil),       // 1: v2.RawQuery
	(*ComplianceRule)(nil), // 2: v2.ComplianceRule
}
var file_api_v2_compliance_rule_service_proto_depIdxs = []int32{
	1, // 0: v2.RuleRequest.query:type_name -> v2.RawQuery
	0, // 1: v2.ComplianceRuleService.GetComplianceRule:input_type -> v2.RuleRequest
	2, // 2: v2.ComplianceRuleService.GetComplianceRule:output_type -> v2.ComplianceRule
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_v2_compliance_rule_service_proto_init() }
func file_api_v2_compliance_rule_service_proto_init() {
	if File_api_v2_compliance_rule_service_proto != nil {
		return
	}
	file_api_v2_compliance_common_proto_init()
	file_api_v2_search_query_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v2_compliance_rule_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v2_compliance_rule_service_proto_goTypes,
		DependencyIndexes: file_api_v2_compliance_rule_service_proto_depIdxs,
		MessageInfos:      file_api_v2_compliance_rule_service_proto_msgTypes,
	}.Build()
	File_api_v2_compliance_rule_service_proto = out.File
	file_api_v2_compliance_rule_service_proto_rawDesc = nil
	file_api_v2_compliance_rule_service_proto_goTypes = nil
	file_api_v2_compliance_rule_service_proto_depIdxs = nil
}
