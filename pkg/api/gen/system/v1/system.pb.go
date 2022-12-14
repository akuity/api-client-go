// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: system/v1/system.proto

package systemv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetVersionRequest) Reset() {
	*x = GetVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1_system_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVersionRequest) ProtoMessage() {}

func (x *GetVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1_system_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVersionRequest.ProtoReflect.Descriptor instead.
func (*GetVersionRequest) Descriptor() ([]byte, []int) {
	return file_system_v1_system_proto_rawDescGZIP(), []int{0}
}

type GetVersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *GetVersionResponse) Reset() {
	*x = GetVersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1_system_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVersionResponse) ProtoMessage() {}

func (x *GetVersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1_system_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVersionResponse.ProtoReflect.Descriptor instead.
func (*GetVersionResponse) Descriptor() ([]byte, []int) {
	return file_system_v1_system_proto_rawDescGZIP(), []int{1}
}

func (x *GetVersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type GetSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSettingsRequest) Reset() {
	*x = GetSettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1_system_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSettingsRequest) ProtoMessage() {}

func (x *GetSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1_system_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSettingsRequest.ProtoReflect.Descriptor instead.
func (*GetSettingsRequest) Descriptor() ([]byte, []int) {
	return file_system_v1_system_proto_rawDescGZIP(), []int{2}
}

type GetSettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Env                    string      `protobuf:"bytes,1,opt,name=env,proto3" json:"env,omitempty"`
	DomainSuffix           string      `protobuf:"bytes,2,opt,name=domain_suffix,json=domainSuffix,proto3" json:"domain_suffix,omitempty"`
	SentryDsn              string      `protobuf:"bytes,3,opt,name=sentry_dsn,json=sentryDsn,proto3" json:"sentry_dsn,omitempty"`
	GoogleAnalytics_4TagId string      `protobuf:"bytes,4,opt,name=google_analytics_4_tag_id,json=googleAnalytics4TagId,proto3" json:"google_analytics_4_tag_id,omitempty"`
	NameConfig             *NameConfig `protobuf:"bytes,5,opt,name=name_config,json=nameConfig,proto3" json:"name_config,omitempty"`
	SelfHosted             bool        `protobuf:"varint,6,opt,name=self_hosted,json=selfHosted,proto3" json:"self_hosted,omitempty"`
	BillingEnabled         bool        `protobuf:"varint,7,opt,name=billing_enabled,json=billingEnabled,proto3" json:"billing_enabled,omitempty"`
}

func (x *GetSettingsResponse) Reset() {
	*x = GetSettingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1_system_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSettingsResponse) ProtoMessage() {}

func (x *GetSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1_system_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSettingsResponse.ProtoReflect.Descriptor instead.
func (*GetSettingsResponse) Descriptor() ([]byte, []int) {
	return file_system_v1_system_proto_rawDescGZIP(), []int{3}
}

func (x *GetSettingsResponse) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *GetSettingsResponse) GetDomainSuffix() string {
	if x != nil {
		return x.DomainSuffix
	}
	return ""
}

func (x *GetSettingsResponse) GetSentryDsn() string {
	if x != nil {
		return x.SentryDsn
	}
	return ""
}

func (x *GetSettingsResponse) GetGoogleAnalytics_4TagId() string {
	if x != nil {
		return x.GoogleAnalytics_4TagId
	}
	return ""
}

func (x *GetSettingsResponse) GetNameConfig() *NameConfig {
	if x != nil {
		return x.NameConfig
	}
	return nil
}

func (x *GetSettingsResponse) GetSelfHosted() bool {
	if x != nil {
		return x.SelfHosted
	}
	return false
}

func (x *GetSettingsResponse) GetBillingEnabled() bool {
	if x != nil {
		return x.BillingEnabled
	}
	return false
}

type ListFeaturesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *ListFeaturesRequest) Reset() {
	*x = ListFeaturesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1_system_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFeaturesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFeaturesRequest) ProtoMessage() {}

func (x *ListFeaturesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1_system_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFeaturesRequest.ProtoReflect.Descriptor instead.
func (*ListFeaturesRequest) Descriptor() ([]byte, []int) {
	return file_system_v1_system_proto_rawDescGZIP(), []int{4}
}

func (x *ListFeaturesRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type ListFeaturesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Features map[string]bool `protobuf:"bytes,1,rep,name=features,proto3" json:"features,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *ListFeaturesResponse) Reset() {
	*x = ListFeaturesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1_system_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFeaturesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFeaturesResponse) ProtoMessage() {}

func (x *ListFeaturesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1_system_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFeaturesResponse.ProtoReflect.Descriptor instead.
func (*ListFeaturesResponse) Descriptor() ([]byte, []int) {
	return file_system_v1_system_proto_rawDescGZIP(), []int{5}
}

func (x *ListFeaturesResponse) GetFeatures() map[string]bool {
	if x != nil {
		return x.Features
	}
	return nil
}

type NameConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinOrganizationNameLength int64 `protobuf:"varint,1,opt,name=min_organization_name_length,json=minOrganizationNameLength,proto3" json:"min_organization_name_length,omitempty"`
	MinInstanceNameLength     int64 `protobuf:"varint,2,opt,name=min_instance_name_length,json=minInstanceNameLength,proto3" json:"min_instance_name_length,omitempty"`
	MinClusterNameLength      int64 `protobuf:"varint,3,opt,name=min_cluster_name_length,json=minClusterNameLength,proto3" json:"min_cluster_name_length,omitempty"`
	MinSubdomainNameLength    int64 `protobuf:"varint,4,opt,name=min_subdomain_name_length,json=minSubdomainNameLength,proto3" json:"min_subdomain_name_length,omitempty"`
}

func (x *NameConfig) Reset() {
	*x = NameConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1_system_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameConfig) ProtoMessage() {}

func (x *NameConfig) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1_system_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameConfig.ProtoReflect.Descriptor instead.
func (*NameConfig) Descriptor() ([]byte, []int) {
	return file_system_v1_system_proto_rawDescGZIP(), []int{6}
}

func (x *NameConfig) GetMinOrganizationNameLength() int64 {
	if x != nil {
		return x.MinOrganizationNameLength
	}
	return 0
}

func (x *NameConfig) GetMinInstanceNameLength() int64 {
	if x != nil {
		return x.MinInstanceNameLength
	}
	return 0
}

func (x *NameConfig) GetMinClusterNameLength() int64 {
	if x != nil {
		return x.MinClusterNameLength
	}
	return 0
}

func (x *NameConfig) GetMinSubdomainNameLength() int64 {
	if x != nil {
		return x.MinSubdomainNameLength
	}
	return 0
}

var File_system_v1_system_proto protoreflect.FileDescriptor

var file_system_v1_system_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2e, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x14, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0xae, 0x02, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65,
	0x6e, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x23, 0x0a,
	0x0d, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x75, 0x66, 0x66,
	0x69, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x64, 0x73, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x44, 0x73,
	0x6e, 0x12, 0x38, 0x0a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x34, 0x5f, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x34, 0x54, 0x61, 0x67, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0b, 0x6e,
	0x61, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0a,
	0x6e, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65,
	0x6c, 0x66, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x73, 0x65, 0x6c, 0x66, 0x48, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x22, 0x2b, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x22, 0xa5, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x08, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x61,
	0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x1a, 0x3b, 0x0a, 0x0d,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf8, 0x01, 0x0a, 0x0a, 0x4e, 0x61,
	0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3f, 0x0a, 0x1c, 0x6d, 0x69, 0x6e, 0x5f,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x19,
	0x6d, 0x69, 0x6e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x37, 0x0a, 0x18, 0x6d, 0x69, 0x6e,
	0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x15, 0x6d, 0x69, 0x6e,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x12, 0x35, 0x0a, 0x17, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x14, 0x6d, 0x69, 0x6e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x39, 0x0a, 0x19, 0x6d, 0x69, 0x6e,
	0x5f, 0x73, 0x75, 0x62, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x6d, 0x69,
	0x6e, 0x53, 0x75, 0x62, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x32, 0x85, 0x03, 0x0a, 0x0d, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x77, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x6b, 0x75, 0x69,
	0x74, 0x79, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x7b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x24,
	0x2e, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x7e, 0x0a, 0x0c,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x25, 0x2e, 0x61,
	0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x42, 0xc5, 0x01, 0x0a,
	0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2e, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x53, 0x58, 0xaa, 0x02, 0x10, 0x41, 0x6b, 0x75,
	0x69, 0x74, 0x79, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10,
	0x41, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x5c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x1c, 0x41, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x5c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x12, 0x41, 0x6b, 0x75, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_system_v1_system_proto_rawDescOnce sync.Once
	file_system_v1_system_proto_rawDescData = file_system_v1_system_proto_rawDesc
)

func file_system_v1_system_proto_rawDescGZIP() []byte {
	file_system_v1_system_proto_rawDescOnce.Do(func() {
		file_system_v1_system_proto_rawDescData = protoimpl.X.CompressGZIP(file_system_v1_system_proto_rawDescData)
	})
	return file_system_v1_system_proto_rawDescData
}

var file_system_v1_system_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_system_v1_system_proto_goTypes = []interface{}{
	(*GetVersionRequest)(nil),    // 0: akuity.system.v1.GetVersionRequest
	(*GetVersionResponse)(nil),   // 1: akuity.system.v1.GetVersionResponse
	(*GetSettingsRequest)(nil),   // 2: akuity.system.v1.GetSettingsRequest
	(*GetSettingsResponse)(nil),  // 3: akuity.system.v1.GetSettingsResponse
	(*ListFeaturesRequest)(nil),  // 4: akuity.system.v1.ListFeaturesRequest
	(*ListFeaturesResponse)(nil), // 5: akuity.system.v1.ListFeaturesResponse
	(*NameConfig)(nil),           // 6: akuity.system.v1.NameConfig
	nil,                          // 7: akuity.system.v1.ListFeaturesResponse.FeaturesEntry
}
var file_system_v1_system_proto_depIdxs = []int32{
	6, // 0: akuity.system.v1.GetSettingsResponse.name_config:type_name -> akuity.system.v1.NameConfig
	7, // 1: akuity.system.v1.ListFeaturesResponse.features:type_name -> akuity.system.v1.ListFeaturesResponse.FeaturesEntry
	0, // 2: akuity.system.v1.SystemService.GetVersion:input_type -> akuity.system.v1.GetVersionRequest
	2, // 3: akuity.system.v1.SystemService.GetSettings:input_type -> akuity.system.v1.GetSettingsRequest
	4, // 4: akuity.system.v1.SystemService.ListFeatures:input_type -> akuity.system.v1.ListFeaturesRequest
	1, // 5: akuity.system.v1.SystemService.GetVersion:output_type -> akuity.system.v1.GetVersionResponse
	3, // 6: akuity.system.v1.SystemService.GetSettings:output_type -> akuity.system.v1.GetSettingsResponse
	5, // 7: akuity.system.v1.SystemService.ListFeatures:output_type -> akuity.system.v1.ListFeaturesResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_system_v1_system_proto_init() }
func file_system_v1_system_proto_init() {
	if File_system_v1_system_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_system_v1_system_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVersionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_system_v1_system_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVersionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_system_v1_system_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSettingsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_system_v1_system_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSettingsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_system_v1_system_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFeaturesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_system_v1_system_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFeaturesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_system_v1_system_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_system_v1_system_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_system_v1_system_proto_goTypes,
		DependencyIndexes: file_system_v1_system_proto_depIdxs,
		MessageInfos:      file_system_v1_system_proto_msgTypes,
	}.Build()
	File_system_v1_system_proto = out.File
	file_system_v1_system_proto_rawDesc = nil
	file_system_v1_system_proto_goTypes = nil
	file_system_v1_system_proto_depIdxs = nil
}
