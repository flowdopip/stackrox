// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v4.25.3
// source: storage/auth_provider.proto

package storage

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

// Next Tag: 15.
type AuthProvider struct {
	state      protoimpl.MessageState `protogen:"open.v1"`
	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk"`     // @gotags: sql:"pk"
	Name       string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" sql:"unique" search:"AuthProvider Name,store,hidden"` // @gotags: sql:"unique" search:"AuthProvider Name,store,hidden"
	Type       string                 `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	UiEndpoint string                 `protobuf:"bytes,4,opt,name=ui_endpoint,json=uiEndpoint,proto3" json:"ui_endpoint,omitempty"`
	Enabled    bool                   `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Config holds auth provider specific configuration. Each configuration options
	// are different based on the given auth provider type.
	// OIDC:
	//   - "issuer": the OIDC issuer according to https://openid.net/specs/openid-connect-core-1_0.html#IssuerIdentifier.
	//   - "client_id": the client ID according to https://www.rfc-editor.org/rfc/rfc6749.html#section-2.2.
	//   - "client_secret": the client secret according to https://www.rfc-editor.org/rfc/rfc6749.html#section-2.3.1.
	//   - "do_not_use_client_secret": set to "true" if you want to create a configuration with only
	//     a client ID and no client secret.
	//   - "mode": the OIDC callback mode, choosing from "fragment", "post", or "query".
	//   - "disable_offline_access_scope": set to "true" if no offline tokens shall be issued.
	//   - "extra_scopes": a space-delimited string of additional scopes to request in addition to "openid profile email"
	//     according to https://www.rfc-editor.org/rfc/rfc6749.html#section-3.3.
	//
	// OpenShift Auth: supports no extra configuration options.
	//
	// User PKI:
	// - "keys": the trusted certificates PEM encoded.
	//
	// SAML:
	// - "sp_issuer": the service provider issuer according to https://datatracker.ietf.org/doc/html/rfc7522#section-3.
	// - "idp_metadata_url": the metadata URL according to https://docs.oasis-open.org/security/saml/v2.0/saml-metadata-2.0-os.pdf.
	// - "idp_issuer": the IdP issuer.
	// - "idp_cert_pem": the cert PEM encoded for the IdP endpoint.
	// - "idp_sso_url": the IdP SSO URL.
	// - "idp_nameid_format": the IdP name ID format.
	//
	// IAP:
	// - "audience": the audience to use.
	Config map[string]string `protobuf:"bytes,6,rep,name=config,proto3" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value" scrub:"map-values"` // @gotags: scrub:"map-values"
	// The login URL will be provided by the backend, and may not be specified in a request.
	LoginUrl string `protobuf:"bytes,7,opt,name=login_url,json=loginUrl,proto3" json:"login_url,omitempty"`
	// Deprecated: Marked as deprecated in storage/auth_provider.proto.
	Validated bool `protobuf:"varint,8,opt,name=validated,proto3" json:"validated,omitempty"`
	// UI endpoints which to allow in addition to `ui_endpoint`. I.e., if a login request
	// is coming from any of these, the auth request will use these for the callback URL,
	// not ui_endpoint.
	ExtraUiEndpoints   []string                          `protobuf:"bytes,9,rep,name=extra_ui_endpoints,json=extraUiEndpoints,proto3" json:"extra_ui_endpoints,omitempty"`
	Active             bool                              `protobuf:"varint,10,opt,name=active,proto3" json:"active,omitempty"`
	RequiredAttributes []*AuthProvider_RequiredAttribute `protobuf:"bytes,11,rep,name=required_attributes,json=requiredAttributes,proto3" json:"required_attributes,omitempty"`
	Traits             *Traits                           `protobuf:"bytes,12,opt,name=traits,proto3" json:"traits,omitempty"`
	// Specifies claims from IdP token that will be copied to Rox token attributes.
	//
	// Each key in this map contains a path in IdP token we want to map. Path is separated by "." symbol.
	// For example, if IdP token payload looks like:
	//
	// {
	//
	//	"a": {
	//
	//	    "b" : "c",
	//
	//	    "d": true,
	//
	//	    "e": [ "val1", "val2", "val3" ],
	//
	//	    "f": [ true, false, false ],
	//
	//	    "g": 123.0,
	//
	//	    "h": [ 1, 2, 3]
	//
	//	}
	//
	// }
	//
	// then "a.b" would be a valid key and "a.z" is not.
	//
	// We support the following types of claims:
	// * string(path "a.b")
	// * bool(path "a.d")
	// * string array(path "a.e")
	// * bool array (path "a.f.")
	//
	// We do NOT support the following types of claims:
	// * complex claims(path "a")
	// * float/integer claims(path "a.g")
	// * float/integer array claims(path "a.h")
	//
	// Each value in this map contains a Rox token attribute name we want to add claim to.
	// If, for example, value is "groups", claim would be found in "external_user.Attributes.groups" in token.
	//
	// Note: we only support this feature for OIDC auth provider.
	ClaimMappings map[string]string `protobuf:"bytes,13,rep,name=claim_mappings,json=claimMappings,proto3" json:"claim_mappings,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Last updated indicates the last time the auth provider has been updated.
	//
	// In case there have been tokens issued by an auth provider _before_ this timestamp, they will be considered
	// invalid. Subsequently, all clients will have to re-issue their tokens (either by refreshing or by an additional
	// login attempt).
	LastUpdated   *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthProvider) Reset() {
	*x = AuthProvider{}
	mi := &file_storage_auth_provider_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthProvider) ProtoMessage() {}

func (x *AuthProvider) ProtoReflect() protoreflect.Message {
	mi := &file_storage_auth_provider_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthProvider.ProtoReflect.Descriptor instead.
func (*AuthProvider) Descriptor() ([]byte, []int) {
	return file_storage_auth_provider_proto_rawDescGZIP(), []int{0}
}

func (x *AuthProvider) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AuthProvider) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuthProvider) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AuthProvider) GetUiEndpoint() string {
	if x != nil {
		return x.UiEndpoint
	}
	return ""
}

func (x *AuthProvider) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *AuthProvider) GetConfig() map[string]string {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *AuthProvider) GetLoginUrl() string {
	if x != nil {
		return x.LoginUrl
	}
	return ""
}

// Deprecated: Marked as deprecated in storage/auth_provider.proto.
func (x *AuthProvider) GetValidated() bool {
	if x != nil {
		return x.Validated
	}
	return false
}

func (x *AuthProvider) GetExtraUiEndpoints() []string {
	if x != nil {
		return x.ExtraUiEndpoints
	}
	return nil
}

func (x *AuthProvider) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *AuthProvider) GetRequiredAttributes() []*AuthProvider_RequiredAttribute {
	if x != nil {
		return x.RequiredAttributes
	}
	return nil
}

func (x *AuthProvider) GetTraits() *Traits {
	if x != nil {
		return x.Traits
	}
	return nil
}

func (x *AuthProvider) GetClaimMappings() map[string]string {
	if x != nil {
		return x.ClaimMappings
	}
	return nil
}

func (x *AuthProvider) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

// RequiredAttribute allows to specify a set of attributes which ALL are required to be returned
// by the auth provider.
// If any attribute is missing within the external claims of the token issued by Central, the
// authentication request to this IdP is considered failed.
type AuthProvider_RequiredAttribute struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	AttributeKey   string                 `protobuf:"bytes,1,opt,name=attribute_key,json=attributeKey,proto3" json:"attribute_key,omitempty"`
	AttributeValue string                 `protobuf:"bytes,2,opt,name=attribute_value,json=attributeValue,proto3" json:"attribute_value,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *AuthProvider_RequiredAttribute) Reset() {
	*x = AuthProvider_RequiredAttribute{}
	mi := &file_storage_auth_provider_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthProvider_RequiredAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthProvider_RequiredAttribute) ProtoMessage() {}

func (x *AuthProvider_RequiredAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_storage_auth_provider_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthProvider_RequiredAttribute.ProtoReflect.Descriptor instead.
func (*AuthProvider_RequiredAttribute) Descriptor() ([]byte, []int) {
	return file_storage_auth_provider_proto_rawDescGZIP(), []int{0, 1}
}

func (x *AuthProvider_RequiredAttribute) GetAttributeKey() string {
	if x != nil {
		return x.AttributeKey
	}
	return ""
}

func (x *AuthProvider_RequiredAttribute) GetAttributeValue() string {
	if x != nil {
		return x.AttributeValue
	}
	return ""
}

var File_storage_auth_provider_proto protoreflect.FileDescriptor

var file_storage_auth_provider_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2f, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x06,
	0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x69, 0x5f, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x69, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x12, 0x39, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x09, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x42, 0x02, 0x18, 0x01,
	0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x5f, 0x75, 0x69, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x65, 0x78, 0x74, 0x72, 0x61, 0x55, 0x69,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x12, 0x58, 0x0a, 0x13, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x06, 0x74,
	0x72, 0x61, 0x69, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73, 0x52, 0x06, 0x74, 0x72,
	0x61, 0x69, 0x74, 0x73, 0x12, 0x4f, 0x0a, 0x0e, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x6d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x1a, 0x39, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x61, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x1a, 0x40, 0x0a, 0x12, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x4d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x2e, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x3b, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_auth_provider_proto_rawDescOnce sync.Once
	file_storage_auth_provider_proto_rawDescData = file_storage_auth_provider_proto_rawDesc
)

func file_storage_auth_provider_proto_rawDescGZIP() []byte {
	file_storage_auth_provider_proto_rawDescOnce.Do(func() {
		file_storage_auth_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_auth_provider_proto_rawDescData)
	})
	return file_storage_auth_provider_proto_rawDescData
}

var file_storage_auth_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_storage_auth_provider_proto_goTypes = []any{
	(*AuthProvider)(nil),                   // 0: storage.AuthProvider
	nil,                                    // 1: storage.AuthProvider.ConfigEntry
	(*AuthProvider_RequiredAttribute)(nil), // 2: storage.AuthProvider.RequiredAttribute
	nil,                                    // 3: storage.AuthProvider.ClaimMappingsEntry
	(*Traits)(nil),                         // 4: storage.Traits
	(*timestamppb.Timestamp)(nil),          // 5: google.protobuf.Timestamp
}
var file_storage_auth_provider_proto_depIdxs = []int32{
	1, // 0: storage.AuthProvider.config:type_name -> storage.AuthProvider.ConfigEntry
	2, // 1: storage.AuthProvider.required_attributes:type_name -> storage.AuthProvider.RequiredAttribute
	4, // 2: storage.AuthProvider.traits:type_name -> storage.Traits
	3, // 3: storage.AuthProvider.claim_mappings:type_name -> storage.AuthProvider.ClaimMappingsEntry
	5, // 4: storage.AuthProvider.last_updated:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_storage_auth_provider_proto_init() }
func file_storage_auth_provider_proto_init() {
	if File_storage_auth_provider_proto != nil {
		return
	}
	file_storage_traits_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_auth_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_auth_provider_proto_goTypes,
		DependencyIndexes: file_storage_auth_provider_proto_depIdxs,
		MessageInfos:      file_storage_auth_provider_proto_msgTypes,
	}.Build()
	File_storage_auth_provider_proto = out.File
	file_storage_auth_provider_proto_rawDesc = nil
	file_storage_auth_provider_proto_goTypes = nil
	file_storage_auth_provider_proto_depIdxs = nil
}
