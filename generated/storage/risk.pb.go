// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.3
// source: storage/risk.proto

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

// Next tag: 9
type RiskSubjectType int32

const (
	RiskSubjectType_UNKNOWN         RiskSubjectType = 0
	RiskSubjectType_DEPLOYMENT      RiskSubjectType = 1
	RiskSubjectType_NAMESPACE       RiskSubjectType = 2
	RiskSubjectType_CLUSTER         RiskSubjectType = 3
	RiskSubjectType_NODE            RiskSubjectType = 7
	RiskSubjectType_NODE_COMPONENT  RiskSubjectType = 8
	RiskSubjectType_IMAGE           RiskSubjectType = 4
	RiskSubjectType_IMAGE_COMPONENT RiskSubjectType = 6
	RiskSubjectType_SERVICEACCOUNT  RiskSubjectType = 5
)

// Enum value maps for RiskSubjectType.
var (
	RiskSubjectType_name = map[int32]string{
		0: "UNKNOWN",
		1: "DEPLOYMENT",
		2: "NAMESPACE",
		3: "CLUSTER",
		7: "NODE",
		8: "NODE_COMPONENT",
		4: "IMAGE",
		6: "IMAGE_COMPONENT",
		5: "SERVICEACCOUNT",
	}
	RiskSubjectType_value = map[string]int32{
		"UNKNOWN":         0,
		"DEPLOYMENT":      1,
		"NAMESPACE":       2,
		"CLUSTER":         3,
		"NODE":            7,
		"NODE_COMPONENT":  8,
		"IMAGE":           4,
		"IMAGE_COMPONENT": 6,
		"SERVICEACCOUNT":  5,
	}
)

func (x RiskSubjectType) Enum() *RiskSubjectType {
	p := new(RiskSubjectType)
	*p = x
	return p
}

func (x RiskSubjectType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RiskSubjectType) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_risk_proto_enumTypes[0].Descriptor()
}

func (RiskSubjectType) Type() protoreflect.EnumType {
	return &file_storage_risk_proto_enumTypes[0]
}

func (x RiskSubjectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RiskSubjectType.Descriptor instead.
func (RiskSubjectType) EnumDescriptor() ([]byte, []int) {
	return file_storage_risk_proto_rawDescGZIP(), []int{0}
}

type Risk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk"` // @gotags: sql:"pk"
	Subject *RiskSubject   `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Score   float32        `protobuf:"fixed32,3,opt,name=score,proto3" json:"score,omitempty" search:"Risk Score,hidden"` // @gotags: search:"Risk Score,hidden"
	Results []*Risk_Result `protobuf:"bytes,4,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *Risk) Reset() {
	*x = Risk{}
	mi := &file_storage_risk_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Risk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Risk) ProtoMessage() {}

func (x *Risk) ProtoReflect() protoreflect.Message {
	mi := &file_storage_risk_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Risk.ProtoReflect.Descriptor instead.
func (*Risk) Descriptor() ([]byte, []int) {
	return file_storage_risk_proto_rawDescGZIP(), []int{0}
}

func (x *Risk) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Risk) GetSubject() *RiskSubject {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *Risk) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Risk) GetResults() []*Risk_Result {
	if x != nil {
		return x.Results
	}
	return nil
}

type RiskSubject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Namespace string          `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty" search:"Namespace,store"`                     // @gotags: search:"Namespace,store"
	ClusterId string          `protobuf:"bytes,3,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty" search:"Cluster ID,store,hidden" sql:"type(uuid)"`    // @gotags: search:"Cluster ID,store,hidden" sql:"type(uuid)"
	Type      RiskSubjectType `protobuf:"varint,4,opt,name=type,proto3,enum=storage.RiskSubjectType" json:"type,omitempty" search:"Risk Subject Type,hidden"` // @gotags: search:"Risk Subject Type,hidden"
}

func (x *RiskSubject) Reset() {
	*x = RiskSubject{}
	mi := &file_storage_risk_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RiskSubject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RiskSubject) ProtoMessage() {}

func (x *RiskSubject) ProtoReflect() protoreflect.Message {
	mi := &file_storage_risk_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RiskSubject.ProtoReflect.Descriptor instead.
func (*RiskSubject) Descriptor() ([]byte, []int) {
	return file_storage_risk_proto_rawDescGZIP(), []int{1}
}

func (x *RiskSubject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RiskSubject) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *RiskSubject) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *RiskSubject) GetType() RiskSubjectType {
	if x != nil {
		return x.Type
	}
	return RiskSubjectType_UNKNOWN
}

type Risk_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Factors []*Risk_Result_Factor `protobuf:"bytes,2,rep,name=factors,proto3" json:"factors,omitempty"`
	Score   float32               `protobuf:"fixed32,3,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *Risk_Result) Reset() {
	*x = Risk_Result{}
	mi := &file_storage_risk_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Risk_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Risk_Result) ProtoMessage() {}

func (x *Risk_Result) ProtoReflect() protoreflect.Message {
	mi := &file_storage_risk_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Risk_Result.ProtoReflect.Descriptor instead.
func (*Risk_Result) Descriptor() ([]byte, []int) {
	return file_storage_risk_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Risk_Result) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Risk_Result) GetFactors() []*Risk_Result_Factor {
	if x != nil {
		return x.Factors
	}
	return nil
}

func (x *Risk_Result) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type Risk_Result_Factor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Url     string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Risk_Result_Factor) Reset() {
	*x = Risk_Result_Factor{}
	mi := &file_storage_risk_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Risk_Result_Factor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Risk_Result_Factor) ProtoMessage() {}

func (x *Risk_Result_Factor) ProtoReflect() protoreflect.Message {
	mi := &file_storage_risk_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Risk_Result_Factor.ProtoReflect.Descriptor instead.
func (*Risk_Result_Factor) Descriptor() ([]byte, []int) {
	return file_storage_risk_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Risk_Result_Factor) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Risk_Result_Factor) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_storage_risk_proto protoreflect.FileDescriptor

var file_storage_risk_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x69, 0x73, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0xae, 0x02,
	0x0a, 0x04, 0x52, 0x69, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x52, 0x69, 0x73, 0x6b, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x2e, 0x0a, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x69, 0x73, 0x6b, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x9f, 0x01, 0x0a,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x66,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x69, 0x73, 0x6b, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x2e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x66, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x34, 0x0a, 0x06, 0x46, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x88,
	0x01, 0x0a, 0x0b, 0x52, 0x69, 0x73, 0x6b, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x52, 0x69, 0x73, 0x6b, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x9c, 0x01, 0x0a, 0x0f, 0x52, 0x69,
	0x73, 0x6b, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x45,
	0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x41,
	0x4d, 0x45, 0x53, 0x50, 0x41, 0x43, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4c, 0x55,
	0x53, 0x54, 0x45, 0x52, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x44, 0x45, 0x10, 0x07,
	0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4f, 0x4e, 0x45,
	0x4e, 0x54, 0x10, 0x08, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x04, 0x12,
	0x13, 0x0a, 0x0f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4f, 0x4e, 0x45,
	0x4e, 0x54, 0x10, 0x06, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x41,
	0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x05, 0x42, 0x2e, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_risk_proto_rawDescOnce sync.Once
	file_storage_risk_proto_rawDescData = file_storage_risk_proto_rawDesc
)

func file_storage_risk_proto_rawDescGZIP() []byte {
	file_storage_risk_proto_rawDescOnce.Do(func() {
		file_storage_risk_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_risk_proto_rawDescData)
	})
	return file_storage_risk_proto_rawDescData
}

var file_storage_risk_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_storage_risk_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_storage_risk_proto_goTypes = []any{
	(RiskSubjectType)(0),       // 0: storage.RiskSubjectType
	(*Risk)(nil),               // 1: storage.Risk
	(*RiskSubject)(nil),        // 2: storage.RiskSubject
	(*Risk_Result)(nil),        // 3: storage.Risk.Result
	(*Risk_Result_Factor)(nil), // 4: storage.Risk.Result.Factor
}
var file_storage_risk_proto_depIdxs = []int32{
	2, // 0: storage.Risk.subject:type_name -> storage.RiskSubject
	3, // 1: storage.Risk.results:type_name -> storage.Risk.Result
	0, // 2: storage.RiskSubject.type:type_name -> storage.RiskSubjectType
	4, // 3: storage.Risk.Result.factors:type_name -> storage.Risk.Result.Factor
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_storage_risk_proto_init() }
func file_storage_risk_proto_init() {
	if File_storage_risk_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_risk_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_risk_proto_goTypes,
		DependencyIndexes: file_storage_risk_proto_depIdxs,
		EnumInfos:         file_storage_risk_proto_enumTypes,
		MessageInfos:      file_storage_risk_proto_msgTypes,
	}.Build()
	File_storage_risk_proto = out.File
	file_storage_risk_proto_rawDesc = nil
	file_storage_risk_proto_goTypes = nil
	file_storage_risk_proto_depIdxs = nil
}
