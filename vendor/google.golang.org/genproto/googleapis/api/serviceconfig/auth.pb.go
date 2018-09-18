// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/auth.proto

package serviceconfig // import "google.golang.org/genproto/googleapis/api/serviceconfig"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// `Authentication` defines the authentication configuration for an API.
//
// Example for an API targeted for external use:
//
//     name: calendar.googleapis.com
//     authentication:
//       providers:
//       - id: google_calendar_auth
//         jwks_uri: https://www.googleapis.com/oauth2/v1/certs
//         issuer: https://securetoken.google.com
//       rules:
//       - selector: "*"
//         requirements:
//           provider_id: google_calendar_auth
type Authentication struct {
	// A list of authentication rules that apply to individual API methods.
	//
	// **NOTE:** All service configuration rules follow "last one wins" order.
	Rules []*AuthenticationRule `protobuf:"bytes,3,rep,name=rules,proto3" json:"rules,omitempty"`
	// Defines a set of authentication providers that a service supports.
	Providers            []*AuthProvider `protobuf:"bytes,4,rep,name=providers,proto3" json:"providers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Authentication) Reset()         { *m = Authentication{} }
func (m *Authentication) String() string { return proto.CompactTextString(m) }
func (*Authentication) ProtoMessage()    {}
func (*Authentication) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_a85b9ac2217fee89, []int{0}
}
func (m *Authentication) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Authentication.Unmarshal(m, b)
}
func (m *Authentication) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Authentication.Marshal(b, m, deterministic)
}
func (dst *Authentication) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Authentication.Merge(dst, src)
}
func (m *Authentication) XXX_Size() int {
	return xxx_messageInfo_Authentication.Size(m)
}
func (m *Authentication) XXX_DiscardUnknown() {
	xxx_messageInfo_Authentication.DiscardUnknown(m)
}

var xxx_messageInfo_Authentication proto.InternalMessageInfo

func (m *Authentication) GetRules() []*AuthenticationRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *Authentication) GetProviders() []*AuthProvider {
	if m != nil {
		return m.Providers
	}
	return nil
}

// Authentication rules for the service.
//
// By default, if a method has any authentication requirements, every request
// must include a valid credential matching one of the requirements.
// It's an error to include more than one kind of credential in a single
// request.
//
// If a method doesn't have any auth requirements, request credentials will be
// ignored.
type AuthenticationRule struct {
	// Selects the methods to which this rule applies.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// The requirements for OAuth credentials.
	Oauth *OAuthRequirements `protobuf:"bytes,2,opt,name=oauth,proto3" json:"oauth,omitempty"`
	// If true, the service accepts API keys without any other credential.
	AllowWithoutCredential bool `protobuf:"varint,5,opt,name=allow_without_credential,json=allowWithoutCredential,proto3" json:"allow_without_credential,omitempty"`
	// Requirements for additional authentication providers.
	Requirements         []*AuthRequirement `protobuf:"bytes,7,rep,name=requirements,proto3" json:"requirements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AuthenticationRule) Reset()         { *m = AuthenticationRule{} }
func (m *AuthenticationRule) String() string { return proto.CompactTextString(m) }
func (*AuthenticationRule) ProtoMessage()    {}
func (*AuthenticationRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_a85b9ac2217fee89, []int{1}
}
func (m *AuthenticationRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationRule.Unmarshal(m, b)
}
func (m *AuthenticationRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationRule.Marshal(b, m, deterministic)
}
func (dst *AuthenticationRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationRule.Merge(dst, src)
}
func (m *AuthenticationRule) XXX_Size() int {
	return xxx_messageInfo_AuthenticationRule.Size(m)
}
func (m *AuthenticationRule) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationRule.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationRule proto.InternalMessageInfo

func (m *AuthenticationRule) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

func (m *AuthenticationRule) GetOauth() *OAuthRequirements {
	if m != nil {
		return m.Oauth
	}
	return nil
}

func (m *AuthenticationRule) GetAllowWithoutCredential() bool {
	if m != nil {
		return m.AllowWithoutCredential
	}
	return false
}

func (m *AuthenticationRule) GetRequirements() []*AuthRequirement {
	if m != nil {
		return m.Requirements
	}
	return nil
}

// Configuration for an anthentication provider, including support for
// [JSON Web Token (JWT)](https://tools.ietf.org/html/draft-ietf-oauth-json-web-token-32).
type AuthProvider struct {
	// The unique identifier of the auth provider. It will be referred to by
	// `AuthRequirement.provider_id`.
	//
	// Example: "bookstore_auth".
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Identifies the principal that issued the JWT. See
	// https://tools.ietf.org/html/draft-ietf-oauth-json-web-token-32#section-4.1.1
	// Usually a URL or an email address.
	//
	// Example: https://securetoken.google.com
	// Example: 1234567-compute@developer.gserviceaccount.com
	Issuer string `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// URL of the provider's public key set to validate signature of the JWT. See
	// [OpenID Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	// Optional if the key set document:
	//  - can be retrieved from
	//    [OpenID Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html
	//    of the issuer.
	//  - can be inferred from the email domain of the issuer (e.g. a Google service account).
	//
	// Example: https://www.googleapis.com/oauth2/v1/certs
	JwksUri string `protobuf:"bytes,3,opt,name=jwks_uri,json=jwksUri,proto3" json:"jwks_uri,omitempty"`
	// The list of JWT
	// [audiences](https://tools.ietf.org/html/draft-ietf-oauth-json-web-token-32#section-4.1.3).
	// that are allowed to access. A JWT containing any of these audiences will
	// be accepted. When this setting is absent, only JWTs with audience
	// "https://[Service_name][google.api.Service.name]/[API_name][google.protobuf.Api.name]"
	// will be accepted. For example, if no audiences are in the setting,
	// LibraryService API will only accept JWTs with the following audience
	// "https://library-example.googleapis.com/google.example.library.v1.LibraryService".
	//
	// Example:
	//
	//     audiences: bookstore_android.apps.googleusercontent.com,
	//                bookstore_web.apps.googleusercontent.com
	Audiences string `protobuf:"bytes,4,opt,name=audiences,proto3" json:"audiences,omitempty"`
	// Redirect URL if JWT token is required but no present or is expired.
	// Implement authorizationUrl of securityDefinitions in OpenAPI spec.
	AuthorizationUrl     string   `protobuf:"bytes,5,opt,name=authorization_url,json=authorizationUrl,proto3" json:"authorization_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthProvider) Reset()         { *m = AuthProvider{} }
func (m *AuthProvider) String() string { return proto.CompactTextString(m) }
func (*AuthProvider) ProtoMessage()    {}
func (*AuthProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_a85b9ac2217fee89, []int{2}
}
func (m *AuthProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthProvider.Unmarshal(m, b)
}
func (m *AuthProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthProvider.Marshal(b, m, deterministic)
}
func (dst *AuthProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthProvider.Merge(dst, src)
}
func (m *AuthProvider) XXX_Size() int {
	return xxx_messageInfo_AuthProvider.Size(m)
}
func (m *AuthProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthProvider.DiscardUnknown(m)
}

var xxx_messageInfo_AuthProvider proto.InternalMessageInfo

func (m *AuthProvider) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AuthProvider) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *AuthProvider) GetJwksUri() string {
	if m != nil {
		return m.JwksUri
	}
	return ""
}

func (m *AuthProvider) GetAudiences() string {
	if m != nil {
		return m.Audiences
	}
	return ""
}

func (m *AuthProvider) GetAuthorizationUrl() string {
	if m != nil {
		return m.AuthorizationUrl
	}
	return ""
}

// OAuth scopes are a way to define data and permissions on data. For example,
// there are scopes defined for "Read-only access to Google Calendar" and
// "Access to Cloud Platform". Users can consent to a scope for an application,
// giving it permission to access that data on their behalf.
//
// OAuth scope specifications should be fairly coarse grained; a user will need
// to see and understand the text description of what your scope means.
//
// In most cases: use one or at most two OAuth scopes for an entire family of
// products. If your product has multiple APIs, you should probably be sharing
// the OAuth scope across all of those APIs.
//
// When you need finer grained OAuth consent screens: talk with your product
// management about how developers will use them in practice.
//
// Please note that even though each of the canonical scopes is enough for a
// request to be accepted and passed to the backend, a request can still fail
// due to the backend requiring additional scopes or permissions.
type OAuthRequirements struct {
	// The list of publicly documented OAuth scopes that are allowed access. An
	// OAuth token containing any of these scopes will be accepted.
	//
	// Example:
	//
	//      canonical_scopes: https://www.googleapis.com/auth/calendar,
	//                        https://www.googleapis.com/auth/calendar.read
	CanonicalScopes      string   `protobuf:"bytes,1,opt,name=canonical_scopes,json=canonicalScopes,proto3" json:"canonical_scopes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OAuthRequirements) Reset()         { *m = OAuthRequirements{} }
func (m *OAuthRequirements) String() string { return proto.CompactTextString(m) }
func (*OAuthRequirements) ProtoMessage()    {}
func (*OAuthRequirements) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_a85b9ac2217fee89, []int{3}
}
func (m *OAuthRequirements) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthRequirements.Unmarshal(m, b)
}
func (m *OAuthRequirements) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthRequirements.Marshal(b, m, deterministic)
}
func (dst *OAuthRequirements) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthRequirements.Merge(dst, src)
}
func (m *OAuthRequirements) XXX_Size() int {
	return xxx_messageInfo_OAuthRequirements.Size(m)
}
func (m *OAuthRequirements) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthRequirements.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthRequirements proto.InternalMessageInfo

func (m *OAuthRequirements) GetCanonicalScopes() string {
	if m != nil {
		return m.CanonicalScopes
	}
	return ""
}

// User-defined authentication requirements, including support for
// [JSON Web Token (JWT)](https://tools.ietf.org/html/draft-ietf-oauth-json-web-token-32).
type AuthRequirement struct {
	// [id][google.api.AuthProvider.id] from authentication provider.
	//
	// Example:
	//
	//     provider_id: bookstore_auth
	ProviderId string `protobuf:"bytes,1,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
	// NOTE: This will be deprecated soon, once AuthProvider.audiences is
	// implemented and accepted in all the runtime components.
	//
	// The list of JWT
	// [audiences](https://tools.ietf.org/html/draft-ietf-oauth-json-web-token-32#section-4.1.3).
	// that are allowed to access. A JWT containing any of these audiences will
	// be accepted. When this setting is absent, only JWTs with audience
	// "https://[Service_name][google.api.Service.name]/[API_name][google.protobuf.Api.name]"
	// will be accepted. For example, if no audiences are in the setting,
	// LibraryService API will only accept JWTs with the following audience
	// "https://library-example.googleapis.com/google.example.library.v1.LibraryService".
	//
	// Example:
	//
	//     audiences: bookstore_android.apps.googleusercontent.com,
	//                bookstore_web.apps.googleusercontent.com
	Audiences            string   `protobuf:"bytes,2,opt,name=audiences,proto3" json:"audiences,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequirement) Reset()         { *m = AuthRequirement{} }
func (m *AuthRequirement) String() string { return proto.CompactTextString(m) }
func (*AuthRequirement) ProtoMessage()    {}
func (*AuthRequirement) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_a85b9ac2217fee89, []int{4}
}
func (m *AuthRequirement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequirement.Unmarshal(m, b)
}
func (m *AuthRequirement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequirement.Marshal(b, m, deterministic)
}
func (dst *AuthRequirement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequirement.Merge(dst, src)
}
func (m *AuthRequirement) XXX_Size() int {
	return xxx_messageInfo_AuthRequirement.Size(m)
}
func (m *AuthRequirement) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequirement.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequirement proto.InternalMessageInfo

func (m *AuthRequirement) GetProviderId() string {
	if m != nil {
		return m.ProviderId
	}
	return ""
}

func (m *AuthRequirement) GetAudiences() string {
	if m != nil {
		return m.Audiences
	}
	return ""
}

func init() {
	proto.RegisterType((*Authentication)(nil), "google.api.Authentication")
	proto.RegisterType((*AuthenticationRule)(nil), "google.api.AuthenticationRule")
	proto.RegisterType((*AuthProvider)(nil), "google.api.AuthProvider")
	proto.RegisterType((*OAuthRequirements)(nil), "google.api.OAuthRequirements")
	proto.RegisterType((*AuthRequirement)(nil), "google.api.AuthRequirement")
}

func init() { proto.RegisterFile("google/api/auth.proto", fileDescriptor_auth_a85b9ac2217fee89) }

var fileDescriptor_auth_a85b9ac2217fee89 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x45, 0x72, 0x9c, 0x58, 0xe3, 0xe0, 0x24, 0x0b, 0x0d, 0xea, 0xb7, 0xd1, 0xc9, 0xa5, 0x20,
	0x43, 0x52, 0x4a, 0xa1, 0xd0, 0x92, 0x94, 0x52, 0x72, 0xaa, 0xd9, 0x12, 0x0a, 0xbd, 0x88, 0xed,
	0x6a, 0x2b, 0x6f, 0xb3, 0xd1, 0xa8, 0xfb, 0x11, 0x43, 0x0f, 0xfd, 0x25, 0x3d, 0xf5, 0x97, 0xf5,
	0xa7, 0x04, 0xad, 0x14, 0x5b, 0x72, 0x8e, 0x33, 0xef, 0xbd, 0x99, 0x79, 0x6f, 0x17, 0x1e, 0x14,
	0x88, 0x85, 0x12, 0x73, 0x56, 0xc9, 0x39, 0x73, 0x76, 0x99, 0x56, 0x1a, 0x2d, 0x12, 0x68, 0xda,
	0x29, 0xab, 0x64, 0xf2, 0x07, 0x26, 0x67, 0xce, 0x2e, 0x45, 0x69, 0x25, 0x67, 0x56, 0x62, 0x49,
	0x5e, 0xc1, 0x50, 0x3b, 0x25, 0x4c, 0x3c, 0x98, 0x0e, 0x66, 0xe3, 0x93, 0x67, 0xe9, 0x86, 0x9d,
	0xf6, 0xa9, 0xd4, 0x29, 0x41, 0x1b, 0x32, 0x79, 0x0d, 0x51, 0xa5, 0xf1, 0x46, 0xe6, 0x42, 0x9b,
	0x78, 0xc7, 0x2b, 0xe3, 0x6d, 0xe5, 0xa2, 0x25, 0xd0, 0x0d, 0x35, 0xf9, 0x1f, 0x00, 0xb9, 0x3f,
	0x95, 0x3c, 0x82, 0x91, 0x11, 0x4a, 0x70, 0x8b, 0x3a, 0x0e, 0xa6, 0xc1, 0x2c, 0xa2, 0xeb, 0x9a,
	0x9c, 0xc2, 0x10, 0x6b, 0x37, 0x71, 0x38, 0x0d, 0x66, 0xe3, 0x93, 0xa7, 0xdd, 0x35, 0x9f, 0xeb,
	0x59, 0x54, 0xfc, 0x72, 0x52, 0x8b, 0x6b, 0x51, 0x5a, 0x43, 0x1b, 0x2e, 0x79, 0x03, 0x31, 0x53,
	0x0a, 0x57, 0xd9, 0x4a, 0xda, 0x25, 0x3a, 0x9b, 0x71, 0x2d, 0xf2, 0x7a, 0x29, 0x53, 0xf1, 0x70,
	0x1a, 0xcc, 0x46, 0xf4, 0xd8, 0xe3, 0x5f, 0x1b, 0xf8, 0xc3, 0x1a, 0x25, 0xef, 0x61, 0x5f, 0x77,
	0x06, 0xc6, 0x7b, 0xde, 0xdc, 0xe3, 0x6d, 0x73, 0x9d, 0xa5, 0xb4, 0x27, 0x48, 0xfe, 0x06, 0xb0,
	0xdf, 0xb5, 0x4f, 0x26, 0x10, 0xca, 0xbc, 0xb5, 0x15, 0xca, 0x9c, 0x1c, 0xc3, 0xae, 0x34, 0xc6,
	0x09, 0xed, 0x1d, 0x45, 0xb4, 0xad, 0xc8, 0x43, 0x18, 0xfd, 0x5c, 0x5d, 0x99, 0xcc, 0x69, 0x19,
	0x0f, 0x3c, 0xb2, 0x57, 0xd7, 0x97, 0x5a, 0x92, 0x27, 0x10, 0x31, 0x97, 0x4b, 0x51, 0x72, 0x51,
	0xc7, 0x5d, 0x63, 0x9b, 0x06, 0x79, 0x09, 0x47, 0xb5, 0x69, 0xd4, 0xf2, 0xb7, 0x8f, 0x34, 0x73,
	0xba, 0x71, 0x19, 0xd1, 0xc3, 0x1e, 0x70, 0xa9, 0x55, 0xf2, 0x0e, 0x8e, 0xee, 0xa5, 0x46, 0x5e,
	0xc0, 0x21, 0x67, 0x25, 0x96, 0x92, 0x33, 0x95, 0x19, 0x8e, 0x95, 0x30, 0xed, 0xc1, 0x07, 0xeb,
	0xfe, 0x17, 0xdf, 0x4e, 0x16, 0x70, 0xb0, 0x25, 0x27, 0xcf, 0x61, 0x7c, 0xf7, 0xc2, 0xd9, 0xda,
	0x29, 0xdc, 0xb5, 0x2e, 0xf2, 0xfe, 0xf9, 0xe1, 0xd6, 0xf9, 0xe7, 0x57, 0x30, 0xe1, 0x78, 0xdd,
	0x09, 0xf8, 0x3c, 0x6a, 0xf3, 0xb3, 0xb8, 0x08, 0xbe, 0x7d, 0x6c, 0x81, 0x02, 0x15, 0x2b, 0x8b,
	0x14, 0x75, 0x31, 0x2f, 0x44, 0xe9, 0xbf, 0xf6, 0xbc, 0x81, 0x58, 0x25, 0x8d, 0xff, 0xf4, 0x46,
	0xe8, 0x1b, 0xc9, 0x05, 0xc7, 0xf2, 0x87, 0x2c, 0xde, 0xf6, 0xaa, 0x7f, 0xe1, 0xce, 0xa7, 0xb3,
	0xc5, 0xc5, 0xf7, 0x5d, 0x2f, 0x3c, 0xbd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x54, 0x91, 0x94, 0x96,
	0x2c, 0x03, 0x00, 0x00,
}