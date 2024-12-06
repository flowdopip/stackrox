// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.3
// source: api/v1/rbac_service.proto

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

// A list of k8s roles (free of scoped information)
// Next Tag: 2
type ListRolesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roles []*storage.K8SRole `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *ListRolesResponse) Reset() {
	*x = ListRolesResponse{}
	mi := &file_api_v1_rbac_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRolesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRolesResponse) ProtoMessage() {}

func (x *ListRolesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRolesResponse.ProtoReflect.Descriptor instead.
func (*ListRolesResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListRolesResponse) GetRoles() []*storage.K8SRole {
	if x != nil {
		return x.Roles
	}
	return nil
}

type GetRoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role *storage.K8SRole `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *GetRoleResponse) Reset() {
	*x = GetRoleResponse{}
	mi := &file_api_v1_rbac_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleResponse) ProtoMessage() {}

func (x *GetRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleResponse.ProtoReflect.Descriptor instead.
func (*GetRoleResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetRoleResponse) GetRole() *storage.K8SRole {
	if x != nil {
		return x.Role
	}
	return nil
}

// A list of k8s role bindings (free of scoped information)
// Next Tag: 2
type ListRoleBindingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bindings []*storage.K8SRoleBinding `protobuf:"bytes,1,rep,name=bindings,proto3" json:"bindings,omitempty"`
}

func (x *ListRoleBindingsResponse) Reset() {
	*x = ListRoleBindingsResponse{}
	mi := &file_api_v1_rbac_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRoleBindingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoleBindingsResponse) ProtoMessage() {}

func (x *ListRoleBindingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoleBindingsResponse.ProtoReflect.Descriptor instead.
func (*ListRoleBindingsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListRoleBindingsResponse) GetBindings() []*storage.K8SRoleBinding {
	if x != nil {
		return x.Bindings
	}
	return nil
}

type GetRoleBindingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Binding *storage.K8SRoleBinding `protobuf:"bytes,1,opt,name=binding,proto3" json:"binding,omitempty"`
}

func (x *GetRoleBindingResponse) Reset() {
	*x = GetRoleBindingResponse{}
	mi := &file_api_v1_rbac_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRoleBindingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleBindingResponse) ProtoMessage() {}

func (x *GetRoleBindingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleBindingResponse.ProtoReflect.Descriptor instead.
func (*GetRoleBindingResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetRoleBindingResponse) GetBinding() *storage.K8SRoleBinding {
	if x != nil {
		return x.Binding
	}
	return nil
}

// A list of k8s subjects (users and groups only, for service accounts, try the service account service)
// Next Tag: 2
type ListSubjectsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubjectAndRoles []*SubjectAndRoles `protobuf:"bytes,1,rep,name=subject_and_roles,json=subjectAndRoles,proto3" json:"subject_and_roles,omitempty"`
}

func (x *ListSubjectsResponse) Reset() {
	*x = ListSubjectsResponse{}
	mi := &file_api_v1_rbac_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSubjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubjectsResponse) ProtoMessage() {}

func (x *ListSubjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubjectsResponse.ProtoReflect.Descriptor instead.
func (*ListSubjectsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListSubjectsResponse) GetSubjectAndRoles() []*SubjectAndRoles {
	if x != nil {
		return x.SubjectAndRoles
	}
	return nil
}

type SubjectAndRoles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject *storage.Subject   `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Roles   []*storage.K8SRole `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *SubjectAndRoles) Reset() {
	*x = SubjectAndRoles{}
	mi := &file_api_v1_rbac_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubjectAndRoles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubjectAndRoles) ProtoMessage() {}

func (x *SubjectAndRoles) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubjectAndRoles.ProtoReflect.Descriptor instead.
func (*SubjectAndRoles) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_service_proto_rawDescGZIP(), []int{5}
}

func (x *SubjectAndRoles) GetSubject() *storage.Subject {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *SubjectAndRoles) GetRoles() []*storage.K8SRole {
	if x != nil {
		return x.Roles
	}
	return nil
}

type GetSubjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject      *storage.Subject   `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	ClusterRoles []*storage.K8SRole `protobuf:"bytes,2,rep,name=cluster_roles,json=clusterRoles,proto3" json:"cluster_roles,omitempty"`
	ScopedRoles  []*ScopedRoles     `protobuf:"bytes,3,rep,name=scoped_roles,json=scopedRoles,proto3" json:"scoped_roles,omitempty"`
}

func (x *GetSubjectResponse) Reset() {
	*x = GetSubjectResponse{}
	mi := &file_api_v1_rbac_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSubjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubjectResponse) ProtoMessage() {}

func (x *GetSubjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubjectResponse.ProtoReflect.Descriptor instead.
func (*GetSubjectResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetSubjectResponse) GetSubject() *storage.Subject {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *GetSubjectResponse) GetClusterRoles() []*storage.K8SRole {
	if x != nil {
		return x.ClusterRoles
	}
	return nil
}

func (x *GetSubjectResponse) GetScopedRoles() []*ScopedRoles {
	if x != nil {
		return x.ScopedRoles
	}
	return nil
}

type ScopedRoles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string             `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Roles     []*storage.K8SRole `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *ScopedRoles) Reset() {
	*x = ScopedRoles{}
	mi := &file_api_v1_rbac_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScopedRoles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScopedRoles) ProtoMessage() {}

func (x *ScopedRoles) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rbac_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScopedRoles.ProtoReflect.Descriptor instead.
func (*ScopedRoles) Descriptor() ([]byte, []int) {
	return file_api_v1_rbac_service_proto_rawDescGZIP(), []int{7}
}

func (x *ScopedRoles) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ScopedRoles) GetRoles() []*storage.K8SRole {
	if x != nil {
		return x.Roles
	}
	return nil
}

var File_api_v1_rbac_service_proto protoreflect.FileDescriptor

var file_api_v1_rbac_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a,
	0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x4b, 0x38, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73,
	0x22, 0x37, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4b, 0x38, 0x73, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x4f, 0x0a, 0x18, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x4b, 0x38, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x52, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x4b, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x4b, 0x38, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x07,
	0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x57, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3f, 0x0a, 0x11, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x6e, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52,
	0x0f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x6e, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x73,
	0x22, 0x65, 0x0a, 0x0f, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x6e, 0x64, 0x52, 0x6f,
	0x6c, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x26, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4b, 0x38, 0x73, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0xab, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x35, 0x0a, 0x0d, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4b, 0x38, 0x73, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65,
	0x73, 0x12, 0x32, 0x0a, 0x0c, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x0b, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x64,
	0x52, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x53, 0x0a, 0x0b, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x52,
	0x6f, 0x6c, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4b, 0x38, 0x73, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x32, 0x8b, 0x04, 0x0a, 0x0b, 0x52,
	0x62, 0x61, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x48, 0x0a, 0x09, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x77, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x1a, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2f, 0x72, 0x6f,
	0x6c, 0x65, 0x73, 0x12, 0x5e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x62, 0x61, 0x63, 0x2f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x59, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x42,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x77,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x62, 0x61, 0x63, 0x2f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x55,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x10, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x16,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x51, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x77, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x1a, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2f,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x42, 0x27, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x58, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_rbac_service_proto_rawDescOnce sync.Once
	file_api_v1_rbac_service_proto_rawDescData = file_api_v1_rbac_service_proto_rawDesc
)

func file_api_v1_rbac_service_proto_rawDescGZIP() []byte {
	file_api_v1_rbac_service_proto_rawDescOnce.Do(func() {
		file_api_v1_rbac_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_rbac_service_proto_rawDescData)
	})
	return file_api_v1_rbac_service_proto_rawDescData
}

var file_api_v1_rbac_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_v1_rbac_service_proto_goTypes = []any{
	(*ListRolesResponse)(nil),        // 0: v1.ListRolesResponse
	(*GetRoleResponse)(nil),          // 1: v1.GetRoleResponse
	(*ListRoleBindingsResponse)(nil), // 2: v1.ListRoleBindingsResponse
	(*GetRoleBindingResponse)(nil),   // 3: v1.GetRoleBindingResponse
	(*ListSubjectsResponse)(nil),     // 4: v1.ListSubjectsResponse
	(*SubjectAndRoles)(nil),          // 5: v1.SubjectAndRoles
	(*GetSubjectResponse)(nil),       // 6: v1.GetSubjectResponse
	(*ScopedRoles)(nil),              // 7: v1.ScopedRoles
	(*storage.K8SRole)(nil),          // 8: storage.K8sRole
	(*storage.K8SRoleBinding)(nil),   // 9: storage.K8sRoleBinding
	(*storage.Subject)(nil),          // 10: storage.Subject
	(*ResourceByID)(nil),             // 11: v1.ResourceByID
	(*RawQuery)(nil),                 // 12: v1.RawQuery
}
var file_api_v1_rbac_service_proto_depIdxs = []int32{
	8,  // 0: v1.ListRolesResponse.roles:type_name -> storage.K8sRole
	8,  // 1: v1.GetRoleResponse.role:type_name -> storage.K8sRole
	9,  // 2: v1.ListRoleBindingsResponse.bindings:type_name -> storage.K8sRoleBinding
	9,  // 3: v1.GetRoleBindingResponse.binding:type_name -> storage.K8sRoleBinding
	5,  // 4: v1.ListSubjectsResponse.subject_and_roles:type_name -> v1.SubjectAndRoles
	10, // 5: v1.SubjectAndRoles.subject:type_name -> storage.Subject
	8,  // 6: v1.SubjectAndRoles.roles:type_name -> storage.K8sRole
	10, // 7: v1.GetSubjectResponse.subject:type_name -> storage.Subject
	8,  // 8: v1.GetSubjectResponse.cluster_roles:type_name -> storage.K8sRole
	7,  // 9: v1.GetSubjectResponse.scoped_roles:type_name -> v1.ScopedRoles
	8,  // 10: v1.ScopedRoles.roles:type_name -> storage.K8sRole
	11, // 11: v1.RbacService.GetRole:input_type -> v1.ResourceByID
	12, // 12: v1.RbacService.ListRoles:input_type -> v1.RawQuery
	11, // 13: v1.RbacService.GetRoleBinding:input_type -> v1.ResourceByID
	12, // 14: v1.RbacService.ListRoleBindings:input_type -> v1.RawQuery
	11, // 15: v1.RbacService.GetSubject:input_type -> v1.ResourceByID
	12, // 16: v1.RbacService.ListSubjects:input_type -> v1.RawQuery
	1,  // 17: v1.RbacService.GetRole:output_type -> v1.GetRoleResponse
	0,  // 18: v1.RbacService.ListRoles:output_type -> v1.ListRolesResponse
	3,  // 19: v1.RbacService.GetRoleBinding:output_type -> v1.GetRoleBindingResponse
	2,  // 20: v1.RbacService.ListRoleBindings:output_type -> v1.ListRoleBindingsResponse
	6,  // 21: v1.RbacService.GetSubject:output_type -> v1.GetSubjectResponse
	4,  // 22: v1.RbacService.ListSubjects:output_type -> v1.ListSubjectsResponse
	17, // [17:23] is the sub-list for method output_type
	11, // [11:17] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_api_v1_rbac_service_proto_init() }
func file_api_v1_rbac_service_proto_init() {
	if File_api_v1_rbac_service_proto != nil {
		return
	}
	file_api_v1_common_proto_init()
	file_api_v1_search_service_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_rbac_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_rbac_service_proto_goTypes,
		DependencyIndexes: file_api_v1_rbac_service_proto_depIdxs,
		MessageInfos:      file_api_v1_rbac_service_proto_msgTypes,
	}.Build()
	File_api_v1_rbac_service_proto = out.File
	file_api_v1_rbac_service_proto_rawDesc = nil
	file_api_v1_rbac_service_proto_goTypes = nil
	file_api_v1_rbac_service_proto_depIdxs = nil
}
