// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.3
// source: api/v2/compliance_results_stats_service.proto

package v2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// ComplianceScanStatsShim models statistics of checks for a given scan configuration
type ComplianceScanStatsShim struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScanName     string                        `protobuf:"bytes,1,opt,name=scan_name,json=scanName,proto3" json:"scan_name,omitempty"`
	CheckStats   []*ComplianceCheckStatusCount `protobuf:"bytes,2,rep,name=check_stats,json=checkStats,proto3" json:"check_stats,omitempty"`
	LastScan     *timestamppb.Timestamp        `protobuf:"bytes,3,opt,name=last_scan,json=lastScan,proto3" json:"last_scan,omitempty"`
	ScanConfigId string                        `protobuf:"bytes,4,opt,name=scan_config_id,json=scanConfigId,proto3" json:"scan_config_id,omitempty"`
}

func (x *ComplianceScanStatsShim) Reset() {
	*x = ComplianceScanStatsShim{}
	mi := &file_api_v2_compliance_results_stats_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ComplianceScanStatsShim) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplianceScanStatsShim) ProtoMessage() {}

func (x *ComplianceScanStatsShim) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_compliance_results_stats_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplianceScanStatsShim.ProtoReflect.Descriptor instead.
func (*ComplianceScanStatsShim) Descriptor() ([]byte, []int) {
	return file_api_v2_compliance_results_stats_service_proto_rawDescGZIP(), []int{0}
}

func (x *ComplianceScanStatsShim) GetScanName() string {
	if x != nil {
		return x.ScanName
	}
	return ""
}

func (x *ComplianceScanStatsShim) GetCheckStats() []*ComplianceCheckStatusCount {
	if x != nil {
		return x.CheckStats
	}
	return nil
}

func (x *ComplianceScanStatsShim) GetLastScan() *timestamppb.Timestamp {
	if x != nil {
		return x.LastScan
	}
	return nil
}

func (x *ComplianceScanStatsShim) GetScanConfigId() string {
	if x != nil {
		return x.ScanConfigId
	}
	return ""
}

// ComplianceProfileScanStats provides scan stats overview based on profile
type ComplianceProfileScanStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckStats  []*ComplianceCheckStatusCount `protobuf:"bytes,1,rep,name=check_stats,json=checkStats,proto3" json:"check_stats,omitempty"`
	ProfileName string                        `protobuf:"bytes,2,opt,name=profile_name,json=profileName,proto3" json:"profile_name,omitempty"`
	Title       string                        `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Version     string                        `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Benchmarks  []*ComplianceBenchmark        `protobuf:"bytes,5,rep,name=benchmarks,proto3" json:"benchmarks,omitempty"`
}

func (x *ComplianceProfileScanStats) Reset() {
	*x = ComplianceProfileScanStats{}
	mi := &file_api_v2_compliance_results_stats_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ComplianceProfileScanStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplianceProfileScanStats) ProtoMessage() {}

func (x *ComplianceProfileScanStats) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_compliance_results_stats_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplianceProfileScanStats.ProtoReflect.Descriptor instead.
func (*ComplianceProfileScanStats) Descriptor() ([]byte, []int) {
	return file_api_v2_compliance_results_stats_service_proto_rawDescGZIP(), []int{1}
}

func (x *ComplianceProfileScanStats) GetCheckStats() []*ComplianceCheckStatusCount {
	if x != nil {
		return x.CheckStats
	}
	return nil
}

func (x *ComplianceProfileScanStats) GetProfileName() string {
	if x != nil {
		return x.ProfileName
	}
	return ""
}

func (x *ComplianceProfileScanStats) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ComplianceProfileScanStats) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ComplianceProfileScanStats) GetBenchmarks() []*ComplianceBenchmark {
	if x != nil {
		return x.Benchmarks
	}
	return nil
}

// ComplianceClusterScanStats provides scan stats overview based on cluster
type ComplianceClusterScanStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScanStats *ComplianceScanStatsShim `protobuf:"bytes,1,opt,name=scan_stats,json=scanStats,proto3" json:"scan_stats,omitempty"`
	Cluster   *ComplianceScanCluster   `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *ComplianceClusterScanStats) Reset() {
	*x = ComplianceClusterScanStats{}
	mi := &file_api_v2_compliance_results_stats_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ComplianceClusterScanStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplianceClusterScanStats) ProtoMessage() {}

func (x *ComplianceClusterScanStats) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_compliance_results_stats_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplianceClusterScanStats.ProtoReflect.Descriptor instead.
func (*ComplianceClusterScanStats) Descriptor() ([]byte, []int) {
	return file_api_v2_compliance_results_stats_service_proto_rawDescGZIP(), []int{2}
}

func (x *ComplianceClusterScanStats) GetScanStats() *ComplianceScanStatsShim {
	if x != nil {
		return x.ScanStats
	}
	return nil
}

func (x *ComplianceClusterScanStats) GetCluster() *ComplianceScanCluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type ComplianceScanClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId string    `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	Query     *RawQuery `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *ComplianceScanClusterRequest) Reset() {
	*x = ComplianceScanClusterRequest{}
	mi := &file_api_v2_compliance_results_stats_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ComplianceScanClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplianceScanClusterRequest) ProtoMessage() {}

func (x *ComplianceScanClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_compliance_results_stats_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplianceScanClusterRequest.ProtoReflect.Descriptor instead.
func (*ComplianceScanClusterRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_compliance_results_stats_service_proto_rawDescGZIP(), []int{3}
}

func (x *ComplianceScanClusterRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *ComplianceScanClusterRequest) GetQuery() *RawQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

// ListComplianceProfileScanStatsResponse provides stats for the profiles within the scans
type ListComplianceProfileScanStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScanStats  []*ComplianceProfileScanStats `protobuf:"bytes,1,rep,name=scan_stats,json=scanStats,proto3" json:"scan_stats,omitempty"`
	TotalCount int32                         `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *ListComplianceProfileScanStatsResponse) Reset() {
	*x = ListComplianceProfileScanStatsResponse{}
	mi := &file_api_v2_compliance_results_stats_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListComplianceProfileScanStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListComplianceProfileScanStatsResponse) ProtoMessage() {}

func (x *ListComplianceProfileScanStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_compliance_results_stats_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListComplianceProfileScanStatsResponse.ProtoReflect.Descriptor instead.
func (*ListComplianceProfileScanStatsResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_compliance_results_stats_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListComplianceProfileScanStatsResponse) GetScanStats() []*ComplianceProfileScanStats {
	if x != nil {
		return x.ScanStats
	}
	return nil
}

func (x *ListComplianceProfileScanStatsResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

// ListComplianceClusterProfileStatsResponse provides stats for the profiles within the scans
type ListComplianceClusterProfileStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScanStats   []*ComplianceProfileScanStats `protobuf:"bytes,1,rep,name=scan_stats,json=scanStats,proto3" json:"scan_stats,omitempty"`
	ClusterId   string                        `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	ClusterName string                        `protobuf:"bytes,3,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	TotalCount  int32                         `protobuf:"varint,4,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *ListComplianceClusterProfileStatsResponse) Reset() {
	*x = ListComplianceClusterProfileStatsResponse{}
	mi := &file_api_v2_compliance_results_stats_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListComplianceClusterProfileStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListComplianceClusterProfileStatsResponse) ProtoMessage() {}

func (x *ListComplianceClusterProfileStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_compliance_results_stats_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListComplianceClusterProfileStatsResponse.ProtoReflect.Descriptor instead.
func (*ListComplianceClusterProfileStatsResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_compliance_results_stats_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListComplianceClusterProfileStatsResponse) GetScanStats() []*ComplianceProfileScanStats {
	if x != nil {
		return x.ScanStats
	}
	return nil
}

func (x *ListComplianceClusterProfileStatsResponse) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *ListComplianceClusterProfileStatsResponse) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *ListComplianceClusterProfileStatsResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

// ListComplianceClusterScanStatsResponse provides stats for the clusters within the scans
type ListComplianceClusterScanStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScanStats  []*ComplianceClusterScanStats `protobuf:"bytes,1,rep,name=scan_stats,json=scanStats,proto3" json:"scan_stats,omitempty"`
	TotalCount int32                         `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *ListComplianceClusterScanStatsResponse) Reset() {
	*x = ListComplianceClusterScanStatsResponse{}
	mi := &file_api_v2_compliance_results_stats_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListComplianceClusterScanStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListComplianceClusterScanStatsResponse) ProtoMessage() {}

func (x *ListComplianceClusterScanStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_compliance_results_stats_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListComplianceClusterScanStatsResponse.ProtoReflect.Descriptor instead.
func (*ListComplianceClusterScanStatsResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_compliance_results_stats_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListComplianceClusterScanStatsResponse) GetScanStats() []*ComplianceClusterScanStats {
	if x != nil {
		return x.ScanStats
	}
	return nil
}

func (x *ListComplianceClusterScanStatsResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_api_v2_compliance_results_stats_service_proto protoreflect.FileDescriptor

var file_api_v2_compliance_results_stats_service_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x76, 0x32, 0x1a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x01,
	0x0a, 0x17, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x63, 0x61, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x68, 0x69, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63, 0x61,
	0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x63,
	0x61, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x76, 0x32,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0a, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x73, 0x63, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x63, 0x61, 0x6e,
	0x12, 0x24, 0x0a, 0x0e, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x63, 0x61, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x22, 0xe9, 0x01, 0x0a, 0x1a, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x69, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x63, 0x61, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x3f, 0x0a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x76, 0x32, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0a, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x0a, 0x62, 0x65, 0x6e,
	0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x65, 0x6e,
	0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x0a, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72,
	0x6b, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x1a, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x63, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x69, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x68,
	0x69, 0x6d, 0x52, 0x09, 0x73, 0x63, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x33, 0x0a,
	0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x63,
	0x61, 0x6e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x22, 0x61, 0x0a, 0x1c, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65,
	0x53, 0x63, 0x61, 0x6e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x22, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x61, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x88, 0x01, 0x0a, 0x26, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53,
	0x63, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3d, 0x0a, 0x0a, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69,
	0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x09, 0x73, 0x63, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0xcd, 0x01, 0x0a, 0x29, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61,
	0x6e, 0x63, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d,
	0x0a, 0x0a, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x09, 0x73, 0x63, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x88, 0x01, 0x0a, 0x26, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61,
	0x6e, 0x63, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x63, 0x61, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x73,
	0x63, 0x61, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x63, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x09, 0x73, 0x63, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xbc, 0x09, 0x0a, 0x1d,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa7, 0x01,
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x23, 0x2e, 0x76, 0x32,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2a, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69,
	0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x33, 0x12, 0x31, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69,
	0x61, 0x6e, 0x63, 0x65, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x82, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x0c, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x61, 0x77, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x1a, 0x2a, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53,
	0x63, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x12, 0x22, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0xb6, 0x01, 0x0a,
	0x21, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x20, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e,
	0x63, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x40, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3a, 0x12, 0x38, 0x2f, 0x76, 0x32,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x73, 0x63, 0x61, 0x6e,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0xb4, 0x01, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x21, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x76, 0x32,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x4d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x47, 0x12, 0x45, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x69, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2f,
	0x7b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0xb0, 0x01, 0x0a,
	0x1d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x63, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x20,
	0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x63,
	0x61, 0x6e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2a, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69,
	0x61, 0x6e, 0x63, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x63, 0x61, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x41, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x3b, 0x12, 0x39, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69,
	0x61, 0x6e, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12,
	0x92, 0x01, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63,
	0x65, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x12, 0x0c, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x61, 0x77, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x1a, 0x2d, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x76, 0x65,
	0x72, 0x61, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x12, 0x29, 0x2f, 0x76, 0x32, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x2f, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x2f, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x12, 0xb3, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x23, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x4f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3c, 0x12, 0x3a,
	0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x73,
	0x63, 0x61, 0x6e, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x7d, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x42, 0x27, 0x0a, 0x18, 0x69, 0x6f,
	0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32,
	0x3b, 0x76, 0x32, 0x58, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v2_compliance_results_stats_service_proto_rawDescOnce sync.Once
	file_api_v2_compliance_results_stats_service_proto_rawDescData = file_api_v2_compliance_results_stats_service_proto_rawDesc
)

func file_api_v2_compliance_results_stats_service_proto_rawDescGZIP() []byte {
	file_api_v2_compliance_results_stats_service_proto_rawDescOnce.Do(func() {
		file_api_v2_compliance_results_stats_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v2_compliance_results_stats_service_proto_rawDescData)
	})
	return file_api_v2_compliance_results_stats_service_proto_rawDescData
}

var file_api_v2_compliance_results_stats_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_v2_compliance_results_stats_service_proto_goTypes = []any{
	(*ComplianceScanStatsShim)(nil),                   // 0: v2.ComplianceScanStatsShim
	(*ComplianceProfileScanStats)(nil),                // 1: v2.ComplianceProfileScanStats
	(*ComplianceClusterScanStats)(nil),                // 2: v2.ComplianceClusterScanStats
	(*ComplianceScanClusterRequest)(nil),              // 3: v2.ComplianceScanClusterRequest
	(*ListComplianceProfileScanStatsResponse)(nil),    // 4: v2.ListComplianceProfileScanStatsResponse
	(*ListComplianceClusterProfileStatsResponse)(nil), // 5: v2.ListComplianceClusterProfileStatsResponse
	(*ListComplianceClusterScanStatsResponse)(nil),    // 6: v2.ListComplianceClusterScanStatsResponse
	(*ComplianceCheckStatusCount)(nil),                // 7: v2.ComplianceCheckStatusCount
	(*timestamppb.Timestamp)(nil),                     // 8: google.protobuf.Timestamp
	(*ComplianceBenchmark)(nil),                       // 9: v2.ComplianceBenchmark
	(*ComplianceScanCluster)(nil),                     // 10: v2.ComplianceScanCluster
	(*RawQuery)(nil),                                  // 11: v2.RawQuery
	(*ComplianceProfileResultsRequest)(nil),           // 12: v2.ComplianceProfileResultsRequest
	(*ComplianceProfileCheckRequest)(nil),             // 13: v2.ComplianceProfileCheckRequest
	(*ListComplianceProfileResults)(nil),              // 14: v2.ListComplianceProfileResults
	(*ListComplianceClusterOverallStatsResponse)(nil), // 15: v2.ListComplianceClusterOverallStatsResponse
}
var file_api_v2_compliance_results_stats_service_proto_depIdxs = []int32{
	7,  // 0: v2.ComplianceScanStatsShim.check_stats:type_name -> v2.ComplianceCheckStatusCount
	8,  // 1: v2.ComplianceScanStatsShim.last_scan:type_name -> google.protobuf.Timestamp
	7,  // 2: v2.ComplianceProfileScanStats.check_stats:type_name -> v2.ComplianceCheckStatusCount
	9,  // 3: v2.ComplianceProfileScanStats.benchmarks:type_name -> v2.ComplianceBenchmark
	0,  // 4: v2.ComplianceClusterScanStats.scan_stats:type_name -> v2.ComplianceScanStatsShim
	10, // 5: v2.ComplianceClusterScanStats.cluster:type_name -> v2.ComplianceScanCluster
	11, // 6: v2.ComplianceScanClusterRequest.query:type_name -> v2.RawQuery
	1,  // 7: v2.ListComplianceProfileScanStatsResponse.scan_stats:type_name -> v2.ComplianceProfileScanStats
	1,  // 8: v2.ListComplianceClusterProfileStatsResponse.scan_stats:type_name -> v2.ComplianceProfileScanStats
	2,  // 9: v2.ListComplianceClusterScanStatsResponse.scan_stats:type_name -> v2.ComplianceClusterScanStats
	12, // 10: v2.ComplianceResultsStatsService.GetComplianceProfileStats:input_type -> v2.ComplianceProfileResultsRequest
	11, // 11: v2.ComplianceResultsStatsService.GetComplianceProfilesStats:input_type -> v2.RawQuery
	3,  // 12: v2.ComplianceResultsStatsService.GetComplianceProfilesClusterStats:input_type -> v2.ComplianceScanClusterRequest
	13, // 13: v2.ComplianceResultsStatsService.GetComplianceProfileCheckStats:input_type -> v2.ComplianceProfileCheckRequest
	3,  // 14: v2.ComplianceResultsStatsService.GetComplianceClusterScanStats:input_type -> v2.ComplianceScanClusterRequest
	11, // 15: v2.ComplianceResultsStatsService.GetComplianceOverallClusterStats:input_type -> v2.RawQuery
	12, // 16: v2.ComplianceResultsStatsService.GetComplianceClusterStats:input_type -> v2.ComplianceProfileResultsRequest
	4,  // 17: v2.ComplianceResultsStatsService.GetComplianceProfileStats:output_type -> v2.ListComplianceProfileScanStatsResponse
	4,  // 18: v2.ComplianceResultsStatsService.GetComplianceProfilesStats:output_type -> v2.ListComplianceProfileScanStatsResponse
	5,  // 19: v2.ComplianceResultsStatsService.GetComplianceProfilesClusterStats:output_type -> v2.ListComplianceClusterProfileStatsResponse
	14, // 20: v2.ComplianceResultsStatsService.GetComplianceProfileCheckStats:output_type -> v2.ListComplianceProfileResults
	6,  // 21: v2.ComplianceResultsStatsService.GetComplianceClusterScanStats:output_type -> v2.ListComplianceClusterScanStatsResponse
	15, // 22: v2.ComplianceResultsStatsService.GetComplianceOverallClusterStats:output_type -> v2.ListComplianceClusterOverallStatsResponse
	15, // 23: v2.ComplianceResultsStatsService.GetComplianceClusterStats:output_type -> v2.ListComplianceClusterOverallStatsResponse
	17, // [17:24] is the sub-list for method output_type
	10, // [10:17] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_api_v2_compliance_results_stats_service_proto_init() }
func file_api_v2_compliance_results_stats_service_proto_init() {
	if File_api_v2_compliance_results_stats_service_proto != nil {
		return
	}
	file_api_v2_compliance_common_proto_init()
	file_api_v2_search_query_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v2_compliance_results_stats_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v2_compliance_results_stats_service_proto_goTypes,
		DependencyIndexes: file_api_v2_compliance_results_stats_service_proto_depIdxs,
		MessageInfos:      file_api_v2_compliance_results_stats_service_proto_msgTypes,
	}.Build()
	File_api_v2_compliance_results_stats_service_proto = out.File
	file_api_v2_compliance_results_stats_service_proto_rawDesc = nil
	file_api_v2_compliance_results_stats_service_proto_goTypes = nil
	file_api_v2_compliance_results_stats_service_proto_depIdxs = nil
}
