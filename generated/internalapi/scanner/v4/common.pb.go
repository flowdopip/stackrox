// This contains protobuf types in pair with ClairCore's types, with
// minimal differences. See https://github.com/quay/claircore for comments
// on the fields.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.3
// source: internalapi/scanner/v4/common.proto

package v4

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

type Contents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Packages      []*Package                   `protobuf:"bytes,1,rep,name=packages,proto3" json:"packages,omitempty"`
	Distributions []*Distribution              `protobuf:"bytes,2,rep,name=distributions,proto3" json:"distributions,omitempty"`
	Repositories  []*Repository                `protobuf:"bytes,3,rep,name=repositories,proto3" json:"repositories,omitempty"`
	Environments  map[string]*Environment_List `protobuf:"bytes,4,rep,name=environments,proto3" json:"environments,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Contents) Reset() {
	*x = Contents{}
	mi := &file_internalapi_scanner_v4_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Contents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contents) ProtoMessage() {}

func (x *Contents) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_scanner_v4_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contents.ProtoReflect.Descriptor instead.
func (*Contents) Descriptor() ([]byte, []int) {
	return file_internalapi_scanner_v4_common_proto_rawDescGZIP(), []int{0}
}

func (x *Contents) GetPackages() []*Package {
	if x != nil {
		return x.Packages
	}
	return nil
}

func (x *Contents) GetDistributions() []*Distribution {
	if x != nil {
		return x.Distributions
	}
	return nil
}

func (x *Contents) GetRepositories() []*Repository {
	if x != nil {
		return x.Repositories
	}
	return nil
}

func (x *Contents) GetEnvironments() map[string]*Environment_List {
	if x != nil {
		return x.Environments
	}
	return nil
}

type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version           string             `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	NormalizedVersion *NormalizedVersion `protobuf:"bytes,4,opt,name=normalized_version,json=normalizedVersion,proto3" json:"normalized_version,omitempty"`
	FixedInVersion    string             `protobuf:"bytes,5,opt,name=fixed_in_version,json=fixedInVersion,proto3" json:"fixed_in_version,omitempty"`
	Kind              string             `protobuf:"bytes,6,opt,name=kind,proto3" json:"kind,omitempty"`
	Source            *Package           `protobuf:"bytes,7,opt,name=source,proto3" json:"source,omitempty"`
	PackageDb         string             `protobuf:"bytes,8,opt,name=package_db,json=packageDb,proto3" json:"package_db,omitempty"`
	RepositoryHint    string             `protobuf:"bytes,9,opt,name=repository_hint,json=repositoryHint,proto3" json:"repository_hint,omitempty"`
	Module            string             `protobuf:"bytes,10,opt,name=module,proto3" json:"module,omitempty"`
	Arch              string             `protobuf:"bytes,11,opt,name=arch,proto3" json:"arch,omitempty"`
	Cpe               string             `protobuf:"bytes,12,opt,name=cpe,proto3" json:"cpe,omitempty"`
}

func (x *Package) Reset() {
	*x = Package{}
	mi := &file_internalapi_scanner_v4_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_scanner_v4_common_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_internalapi_scanner_v4_common_proto_rawDescGZIP(), []int{1}
}

func (x *Package) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Package) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Package) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Package) GetNormalizedVersion() *NormalizedVersion {
	if x != nil {
		return x.NormalizedVersion
	}
	return nil
}

func (x *Package) GetFixedInVersion() string {
	if x != nil {
		return x.FixedInVersion
	}
	return ""
}

func (x *Package) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Package) GetSource() *Package {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *Package) GetPackageDb() string {
	if x != nil {
		return x.PackageDb
	}
	return ""
}

func (x *Package) GetRepositoryHint() string {
	if x != nil {
		return x.RepositoryHint
	}
	return ""
}

func (x *Package) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (x *Package) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

func (x *Package) GetCpe() string {
	if x != nil {
		return x.Cpe
	}
	return ""
}

type NormalizedVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind string  `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	V    []int32 `protobuf:"varint,2,rep,packed,name=v,proto3" json:"v,omitempty"`
}

func (x *NormalizedVersion) Reset() {
	*x = NormalizedVersion{}
	mi := &file_internalapi_scanner_v4_common_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NormalizedVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NormalizedVersion) ProtoMessage() {}

func (x *NormalizedVersion) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_scanner_v4_common_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NormalizedVersion.ProtoReflect.Descriptor instead.
func (*NormalizedVersion) Descriptor() ([]byte, []int) {
	return file_internalapi_scanner_v4_common_proto_rawDescGZIP(), []int{2}
}

func (x *NormalizedVersion) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *NormalizedVersion) GetV() []int32 {
	if x != nil {
		return x.V
	}
	return nil
}

type Distribution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Did             string `protobuf:"bytes,2,opt,name=did,proto3" json:"did,omitempty"`
	Name            string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Version         string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	VersionCodeName string `protobuf:"bytes,5,opt,name=version_code_name,json=versionCodeName,proto3" json:"version_code_name,omitempty"`
	VersionId       string `protobuf:"bytes,6,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	Arch            string `protobuf:"bytes,7,opt,name=arch,proto3" json:"arch,omitempty"`
	Cpe             string `protobuf:"bytes,8,opt,name=cpe,proto3" json:"cpe,omitempty"`
	PrettyName      string `protobuf:"bytes,9,opt,name=pretty_name,json=prettyName,proto3" json:"pretty_name,omitempty"`
}

func (x *Distribution) Reset() {
	*x = Distribution{}
	mi := &file_internalapi_scanner_v4_common_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Distribution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Distribution) ProtoMessage() {}

func (x *Distribution) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_scanner_v4_common_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Distribution.ProtoReflect.Descriptor instead.
func (*Distribution) Descriptor() ([]byte, []int) {
	return file_internalapi_scanner_v4_common_proto_rawDescGZIP(), []int{3}
}

func (x *Distribution) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Distribution) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *Distribution) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Distribution) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Distribution) GetVersionCodeName() string {
	if x != nil {
		return x.VersionCodeName
	}
	return ""
}

func (x *Distribution) GetVersionId() string {
	if x != nil {
		return x.VersionId
	}
	return ""
}

func (x *Distribution) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

func (x *Distribution) GetCpe() string {
	if x != nil {
		return x.Cpe
	}
	return ""
}

func (x *Distribution) GetPrettyName() string {
	if x != nil {
		return x.PrettyName
	}
	return ""
}

type Repository struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Key  string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Uri  string `protobuf:"bytes,4,opt,name=uri,proto3" json:"uri,omitempty"`
	Cpe  string `protobuf:"bytes,5,opt,name=cpe,proto3" json:"cpe,omitempty"`
}

func (x *Repository) Reset() {
	*x = Repository{}
	mi := &file_internalapi_scanner_v4_common_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Repository) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Repository) ProtoMessage() {}

func (x *Repository) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_scanner_v4_common_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Repository.ProtoReflect.Descriptor instead.
func (*Repository) Descriptor() ([]byte, []int) {
	return file_internalapi_scanner_v4_common_proto_rawDescGZIP(), []int{4}
}

func (x *Repository) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Repository) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Repository) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Repository) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Repository) GetCpe() string {
	if x != nil {
		return x.Cpe
	}
	return ""
}

// Environment describes the surrounding environment a package was
// discovered in.
type Environment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageDb      string   `protobuf:"bytes,1,opt,name=package_db,json=packageDb,proto3" json:"package_db,omitempty"`
	IntroducedIn   string   `protobuf:"bytes,2,opt,name=introduced_in,json=introducedIn,proto3" json:"introduced_in,omitempty"`
	DistributionId string   `protobuf:"bytes,3,opt,name=distribution_id,json=distributionId,proto3" json:"distribution_id,omitempty"`
	RepositoryIds  []string `protobuf:"bytes,4,rep,name=repository_ids,json=repositoryIds,proto3" json:"repository_ids,omitempty"`
}

func (x *Environment) Reset() {
	*x = Environment{}
	mi := &file_internalapi_scanner_v4_common_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Environment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Environment) ProtoMessage() {}

func (x *Environment) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_scanner_v4_common_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Environment.ProtoReflect.Descriptor instead.
func (*Environment) Descriptor() ([]byte, []int) {
	return file_internalapi_scanner_v4_common_proto_rawDescGZIP(), []int{5}
}

func (x *Environment) GetPackageDb() string {
	if x != nil {
		return x.PackageDb
	}
	return ""
}

func (x *Environment) GetIntroducedIn() string {
	if x != nil {
		return x.IntroducedIn
	}
	return ""
}

func (x *Environment) GetDistributionId() string {
	if x != nil {
		return x.DistributionId
	}
	return ""
}

func (x *Environment) GetRepositoryIds() []string {
	if x != nil {
		return x.RepositoryIds
	}
	return nil
}

type Environment_List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Environments []*Environment `protobuf:"bytes,1,rep,name=environments,proto3" json:"environments,omitempty"`
}

func (x *Environment_List) Reset() {
	*x = Environment_List{}
	mi := &file_internalapi_scanner_v4_common_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Environment_List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Environment_List) ProtoMessage() {}

func (x *Environment_List) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_scanner_v4_common_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Environment_List.ProtoReflect.Descriptor instead.
func (*Environment_List) Descriptor() ([]byte, []int) {
	return file_internalapi_scanner_v4_common_proto_rawDescGZIP(), []int{5, 0}
}

func (x *Environment_List) GetEnvironments() []*Environment {
	if x != nil {
		return x.Environments
	}
	return nil
}

var File_internalapi_scanner_v4_common_proto protoreflect.FileDescriptor

var file_internalapi_scanner_v4_common_proto_rawDesc = []byte{
	0x0a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x34, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76,
	0x34, 0x22, 0xe2, 0x02, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2f,
	0x0a, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x2e, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x3e, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x2e, 0x76, 0x34, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0d, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x3a, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e,
	0x76, 0x34, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x0c, 0x72,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x4a, 0x0a, 0x0c, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x5d, 0x0a, 0x11, 0x45, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x32,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x86, 0x03, 0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x4c, 0x0a, 0x12, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73,
	0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x2e, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x6e, 0x6f, 0x72,
	0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28,
	0x0a, 0x10, 0x66, 0x69, 0x78, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x69, 0x78, 0x65, 0x64, 0x49,
	0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x2b, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73,
	0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x62, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x44, 0x62, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x68, 0x69, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x48, 0x69, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x63,
	0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x63, 0x68, 0x12, 0x10, 0x0a,
	0x03, 0x63, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x65, 0x22,
	0x35, 0x0a, 0x11, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x01, 0x76, 0x22, 0xf0, 0x01, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x63, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x61, 0x72, 0x63, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x74,
	0x74, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x72, 0x65, 0x74, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x66, 0x0a, 0x0a, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12,
	0x10, 0x0a, 0x03, 0x63, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70,
	0x65, 0x22, 0xe6, 0x01, 0x0a, 0x0b, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x62, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x44, 0x62,
	0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x64, 0x5f, 0x69,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x64, 0x49, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x25,
	0x0a, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x49, 0x64, 0x73, 0x1a, 0x43, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3b, 0x0a,
	0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x34,
	0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x1d, 0x5a, 0x1b, 0x2e, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x2f, 0x76, 0x34, 0x3b, 0x76, 0x34, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_internalapi_scanner_v4_common_proto_rawDescOnce sync.Once
	file_internalapi_scanner_v4_common_proto_rawDescData = file_internalapi_scanner_v4_common_proto_rawDesc
)

func file_internalapi_scanner_v4_common_proto_rawDescGZIP() []byte {
	file_internalapi_scanner_v4_common_proto_rawDescOnce.Do(func() {
		file_internalapi_scanner_v4_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalapi_scanner_v4_common_proto_rawDescData)
	})
	return file_internalapi_scanner_v4_common_proto_rawDescData
}

var file_internalapi_scanner_v4_common_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_internalapi_scanner_v4_common_proto_goTypes = []any{
	(*Contents)(nil),          // 0: scanner.v4.Contents
	(*Package)(nil),           // 1: scanner.v4.Package
	(*NormalizedVersion)(nil), // 2: scanner.v4.NormalizedVersion
	(*Distribution)(nil),      // 3: scanner.v4.Distribution
	(*Repository)(nil),        // 4: scanner.v4.Repository
	(*Environment)(nil),       // 5: scanner.v4.Environment
	nil,                       // 6: scanner.v4.Contents.EnvironmentsEntry
	(*Environment_List)(nil),  // 7: scanner.v4.Environment.List
}
var file_internalapi_scanner_v4_common_proto_depIdxs = []int32{
	1, // 0: scanner.v4.Contents.packages:type_name -> scanner.v4.Package
	3, // 1: scanner.v4.Contents.distributions:type_name -> scanner.v4.Distribution
	4, // 2: scanner.v4.Contents.repositories:type_name -> scanner.v4.Repository
	6, // 3: scanner.v4.Contents.environments:type_name -> scanner.v4.Contents.EnvironmentsEntry
	2, // 4: scanner.v4.Package.normalized_version:type_name -> scanner.v4.NormalizedVersion
	1, // 5: scanner.v4.Package.source:type_name -> scanner.v4.Package
	7, // 6: scanner.v4.Contents.EnvironmentsEntry.value:type_name -> scanner.v4.Environment.List
	5, // 7: scanner.v4.Environment.List.environments:type_name -> scanner.v4.Environment
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_internalapi_scanner_v4_common_proto_init() }
func file_internalapi_scanner_v4_common_proto_init() {
	if File_internalapi_scanner_v4_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internalapi_scanner_v4_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internalapi_scanner_v4_common_proto_goTypes,
		DependencyIndexes: file_internalapi_scanner_v4_common_proto_depIdxs,
		MessageInfos:      file_internalapi_scanner_v4_common_proto_msgTypes,
	}.Build()
	File_internalapi_scanner_v4_common_proto = out.File
	file_internalapi_scanner_v4_common_proto_rawDesc = nil
	file_internalapi_scanner_v4_common_proto_goTypes = nil
	file_internalapi_scanner_v4_common_proto_depIdxs = nil
}
