// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v4.25.3
// source: storage/node_component.proto

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

type NodeComponent struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" search:"Component ID,store,hidden" sql:"pk,id"`                                  // This field is composite id over name, version, and operating system. // @gotags: search:"Component ID,store,hidden" sql:"pk,id"
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" search:"Component,store"`                              // @gotags: search:"Component,store"
	Version   string                 `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty" search:"Component Version,store"`                        // @gotags: search:"Component Version,store"
	Priority  int64                  `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty" search:"Component Risk Priority,hidden"`                     // @gotags: search:"Component Risk Priority,hidden"
	RiskScore float32                `protobuf:"fixed32,7,opt,name=risk_score,json=riskScore,proto3" json:"risk_score,omitempty" search:"Component Risk Score,hidden"` // @gotags: search:"Component Risk Score,hidden"
	// Types that are valid to be assigned to SetTopCvss:
	//
	//	*NodeComponent_TopCvss
	SetTopCvss      isNodeComponent_SetTopCvss `protobuf_oneof:"set_top_cvss"`
	OperatingSystem string                     `protobuf:"bytes,9,opt,name=operating_system,json=operatingSystem,proto3" json:"operating_system,omitempty" search:"Operating System"` // @gotags: search:"Operating System"
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *NodeComponent) Reset() {
	*x = NodeComponent{}
	mi := &file_storage_node_component_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeComponent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeComponent) ProtoMessage() {}

func (x *NodeComponent) ProtoReflect() protoreflect.Message {
	mi := &file_storage_node_component_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeComponent.ProtoReflect.Descriptor instead.
func (*NodeComponent) Descriptor() ([]byte, []int) {
	return file_storage_node_component_proto_rawDescGZIP(), []int{0}
}

func (x *NodeComponent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NodeComponent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NodeComponent) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *NodeComponent) GetPriority() int64 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *NodeComponent) GetRiskScore() float32 {
	if x != nil {
		return x.RiskScore
	}
	return 0
}

func (x *NodeComponent) GetSetTopCvss() isNodeComponent_SetTopCvss {
	if x != nil {
		return x.SetTopCvss
	}
	return nil
}

func (x *NodeComponent) GetTopCvss() float32 {
	if x != nil {
		if x, ok := x.SetTopCvss.(*NodeComponent_TopCvss); ok {
			return x.TopCvss
		}
	}
	return 0
}

func (x *NodeComponent) GetOperatingSystem() string {
	if x != nil {
		return x.OperatingSystem
	}
	return ""
}

type isNodeComponent_SetTopCvss interface {
	isNodeComponent_SetTopCvss()
}

type NodeComponent_TopCvss struct {
	TopCvss float32 `protobuf:"fixed32,8,opt,name=top_cvss,json=topCvss,proto3,oneof" search:"Component Top CVSS,store"` // @gotags: search:"Component Top CVSS,store"
}

func (*NodeComponent_TopCvss) isNodeComponent_SetTopCvss() {}

var File_storage_node_component_proto protoreflect.FileDescriptor

var file_storage_node_component_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0xe0, 0x01, 0x0a, 0x0d, 0x4e, 0x6f, 0x64, 0x65,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x69, 0x73, 0x6b, 0x5f, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x72, 0x69, 0x73, 0x6b, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x1b, 0x0a, 0x08, 0x74, 0x6f, 0x70, 0x5f, 0x63, 0x76, 0x73, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x07, 0x74, 0x6f, 0x70, 0x43, 0x76, 0x73, 0x73, 0x12,
	0x29, 0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x42, 0x0e, 0x0a, 0x0c, 0x73, 0x65,
	0x74, 0x5f, 0x74, 0x6f, 0x70, 0x5f, 0x63, 0x76, 0x73, 0x73, 0x42, 0x2e, 0x0a, 0x19, 0x69, 0x6f,
	0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_storage_node_component_proto_rawDescOnce sync.Once
	file_storage_node_component_proto_rawDescData = file_storage_node_component_proto_rawDesc
)

func file_storage_node_component_proto_rawDescGZIP() []byte {
	file_storage_node_component_proto_rawDescOnce.Do(func() {
		file_storage_node_component_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_node_component_proto_rawDescData)
	})
	return file_storage_node_component_proto_rawDescData
}

var file_storage_node_component_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_storage_node_component_proto_goTypes = []any{
	(*NodeComponent)(nil), // 0: storage.NodeComponent
}
var file_storage_node_component_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_storage_node_component_proto_init() }
func file_storage_node_component_proto_init() {
	if File_storage_node_component_proto != nil {
		return
	}
	file_storage_node_component_proto_msgTypes[0].OneofWrappers = []any{
		(*NodeComponent_TopCvss)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_node_component_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_node_component_proto_goTypes,
		DependencyIndexes: file_storage_node_component_proto_depIdxs,
		MessageInfos:      file_storage_node_component_proto_msgTypes,
	}.Build()
	File_storage_node_component_proto = out.File
	file_storage_node_component_proto_rawDesc = nil
	file_storage_node_component_proto_goTypes = nil
	file_storage_node_component_proto_depIdxs = nil
}
