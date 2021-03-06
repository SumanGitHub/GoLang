// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/api/serviceconfig/auth.proto
// DO NOT EDIT!

package google_api // import "google.golang.org/genproto/googleapis/api/serviceconfig"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// `Authentication` defines the authentication configuration for an API.
//
// Example for an API targeted for external use:
//
//     name: calendar.googleapis.com
//     authentication:
//       rules:
//       - selector: "*"
//         oauth:
//           canonical_scopes: https://www.googleapis.com/auth/calendar
//
//       - selector: google.calendar.Delegate
//         oauth:
//           canonical_scopes: https://www.googleapis.com/auth/calendar.read
type Authentication struct {
	// Individual rules for authentication.
	Rules []*AuthenticationRule `protobuf:"bytes,3,rep,name=rules" json:"rules,omitempty"`
	// Defines a set of authentication providers that a service supports.
	Providers []*AuthProvider `protobuf:"bytes,4,rep,name=providers" json:"providers,omitempty"`
}

func (m *Authentication) Reset()                    { *m = Authentication{} }
func (m *Authentication) String() string            { return proto.CompactTextString(m) }
func (*Authentication) ProtoMessage()               {}
func (*Authentication) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

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
//
type AuthenticationRule struct {
	// Selects the methods to which this rule applies.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector" json:"selector,omitempty"`
	// The requirements for OAuth credentials.
	Oauth *OAuthRequirements `protobuf:"bytes,2,opt,name=oauth" json:"oauth,omitempty"`
	// Whether to allow requests without a credential.  If quota is enabled, an
	// API key is required for such request to pass the quota check.
	//
	AllowWithoutCredential bool `protobuf:"varint,5,opt,name=allow_without_credential,json=allowWithoutCredential" json:"allow_without_credential,omitempty"`
	// Requirements for additional authentication providers.
	Requirements []*AuthRequirement `protobuf:"bytes,7,rep,name=requirements" json:"requirements,omitempty"`
}

func (m *AuthenticationRule) Reset()                    { *m = AuthenticationRule{} }
func (m *AuthenticationRule) String() string            { return proto.CompactTextString(m) }
func (*AuthenticationRule) ProtoMessage()               {}
func (*AuthenticationRule) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *AuthenticationRule) GetOauth() *OAuthRequirements {
	if m != nil {
		return m.Oauth
	}
	return nil
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
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Identifies the principal that issued the JWT. See
	// https://tools.ietf.org/html/draft-ietf-oauth-json-web-token-32#section-4.1.1
	// Usually a URL or an email address.
	//
	// Example: https://securetoken.google.com
	// Example: 1234567-compute@developer.gserviceaccount.com
	Issuer string `protobuf:"bytes,2,opt,name=issuer" json:"issuer,omitempty"`
	// URL of the provider's public key set to validate signature of the JWT. See
	// [OpenID Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	// Optional if the key set document:
	//  - can be retrieved from
	//    [OpenID Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html
	//    of the issuer.
	//  - can be inferred from the email domain of the issuer (e.g. a Google service account).
	//
	// Example: https://www.googleapis.com/oauth2/v1/certs
	JwksUri string `protobuf:"bytes,3,opt,name=jwks_uri,json=jwksUri" json:"jwks_uri,omitempty"`
}

func (m *AuthProvider) Reset()                    { *m = AuthProvider{} }
func (m *AuthProvider) String() string            { return proto.CompactTextString(m) }
func (*AuthProvider) ProtoMessage()               {}
func (*AuthProvider) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

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
//
type OAuthRequirements struct {
	// The list of publicly documented OAuth scopes that are allowed access. An
	// OAuth token containing any of these scopes will be accepted.
	//
	// Example:
	//
	//      canonical_scopes: https://www.googleapis.com/auth/calendar,
	//                        https://www.googleapis.com/auth/calendar.read
	CanonicalScopes string `protobuf:"bytes,1,opt,name=canonical_scopes,json=canonicalScopes" json:"canonical_scopes,omitempty"`
}

func (m *OAuthRequirements) Reset()                    { *m = OAuthRequirements{} }
func (m *OAuthRequirements) String() string            { return proto.CompactTextString(m) }
func (*OAuthRequirements) ProtoMessage()               {}
func (*OAuthRequirements) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

// User-defined authentication requirements, including support for
// [JSON Web Token (JWT)](https://tools.ietf.org/html/draft-ietf-oauth-json-web-token-32).
type AuthRequirement struct {
	// [id][google.api.AuthProvider.id] from authentication provider.
	//
	// Example:
	//
	//     provider_id: bookstore_auth
	ProviderId string `protobuf:"bytes,1,opt,name=provider_id,json=providerId" json:"provider_id,omitempty"`
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
	Audiences string `protobuf:"bytes,2,opt,name=audiences" json:"audiences,omitempty"`
}

func (m *AuthRequirement) Reset()                    { *m = AuthRequirement{} }
func (m *AuthRequirement) String() string            { return proto.CompactTextString(m) }
func (*AuthRequirement) ProtoMessage()               {}
func (*AuthRequirement) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func init() {
	proto.RegisterType((*Authentication)(nil), "google.api.Authentication")
	proto.RegisterType((*AuthenticationRule)(nil), "google.api.AuthenticationRule")
	proto.RegisterType((*AuthProvider)(nil), "google.api.AuthProvider")
	proto.RegisterType((*OAuthRequirements)(nil), "google.api.OAuthRequirements")
	proto.RegisterType((*AuthRequirement)(nil), "google.api.AuthRequirement")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/api/serviceconfig/auth.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x52, 0xc1, 0x72, 0xd3, 0x30,
	0x10, 0x1d, 0xb7, 0xa4, 0x8d, 0x37, 0x9d, 0x14, 0x74, 0xe8, 0x88, 0x42, 0xa1, 0xe3, 0x53, 0xb9,
	0xd8, 0x33, 0x2d, 0xc3, 0x70, 0x82, 0x21, 0x9c, 0x7a, 0x22, 0x88, 0x61, 0x38, 0x7a, 0x84, 0x2c,
	0x5c, 0x81, 0xaa, 0x0d, 0x92, 0xdc, 0xdc, 0xf8, 0x56, 0x3e, 0x05, 0x59, 0x76, 0x1d, 0x37, 0xb9,
	0x71, 0xc9, 0x64, 0xf7, 0xbd, 0x7d, 0x4f, 0x6f, 0xd7, 0xb0, 0xa8, 0x11, 0x6b, 0x2d, 0xf3, 0x1a,
	0x35, 0x37, 0x75, 0x8e, 0xb6, 0x2e, 0x6a, 0x69, 0x56, 0x16, 0x3d, 0x16, 0x1d, 0xc4, 0x57, 0xca,
	0x15, 0xe1, 0xa7, 0x70, 0xd2, 0xde, 0x29, 0x21, 0x05, 0x9a, 0x1f, 0xaa, 0x2e, 0x78, 0xe3, 0x6f,
	0xf2, 0xc8, 0x23, 0xd0, 0x6b, 0x04, 0xd2, 0xe9, 0xf5, 0x7f, 0xeb, 0x19, 0x83, 0x9e, 0x7b, 0x85,
	0xc6, 0x75, 0xb2, 0xd9, 0x1f, 0x98, 0x7f, 0x08, 0x26, 0xd2, 0x78, 0x25, 0x22, 0x40, 0x5e, 0xc3,
	0xc4, 0x36, 0x5a, 0x3a, 0xba, 0x7f, 0xbe, 0x7f, 0x31, 0xbb, 0x7c, 0x91, 0x6f, 0x8c, 0xf3, 0x87,
	0x54, 0x16, 0x68, 0xac, 0x23, 0x93, 0x37, 0x90, 0x06, 0xc1, 0x3b, 0x55, 0x49, 0xeb, 0xe8, 0xa3,
	0x38, 0x49, 0xb7, 0x27, 0x97, 0x3d, 0x81, 0x6d, 0xa8, 0xd9, 0xdf, 0x04, 0xc8, 0xae, 0x2a, 0x39,
	0x85, 0xa9, 0x93, 0x5a, 0x0a, 0x8f, 0x96, 0x26, 0xe7, 0xc9, 0x45, 0xca, 0x86, 0x9a, 0x5c, 0xc1,
	0x04, 0xdb, 0xc5, 0xd0, 0xbd, 0x00, 0xcc, 0x2e, 0xcf, 0xc6, 0x36, 0x9f, 0x5a, 0x2d, 0x26, 0x7f,
	0x37, 0xca, 0xca, 0xdb, 0xa0, 0xe9, 0x58, 0xc7, 0x25, 0x6f, 0x81, 0x72, 0xad, 0x71, 0x5d, 0xae,
	0x95, 0xbf, 0xc1, 0xc6, 0x97, 0xc2, 0xca, 0xaa, 0x35, 0xe5, 0x9a, 0x4e, 0x82, 0xce, 0x94, 0x9d,
	0x44, 0xfc, 0x5b, 0x07, 0x7f, 0x1c, 0x50, 0xf2, 0x1e, 0x8e, 0xec, 0x48, 0x90, 0x1e, 0xc6, 0x70,
	0xcf, 0xb6, 0xc3, 0x8d, 0x4c, 0xd9, 0x83, 0x81, 0xec, 0x33, 0x1c, 0x8d, 0xd3, 0x93, 0x39, 0xec,
	0xa9, 0xaa, 0x4f, 0x15, 0xfe, 0x91, 0x13, 0x38, 0x50, 0xce, 0x35, 0xd2, 0xc6, 0x40, 0x29, 0xeb,
	0x2b, 0xf2, 0x14, 0xa6, 0x3f, 0xd7, 0xbf, 0x5c, 0xd9, 0x58, 0x15, 0x6e, 0xd1, 0x22, 0x87, 0x6d,
	0xfd, 0xd5, 0xaa, 0xec, 0x1d, 0x3c, 0xd9, 0x49, 0x4a, 0x5e, 0xc1, 0x63, 0xc1, 0x0d, 0x9a, 0xb0,
	0x47, 0x5d, 0x3a, 0x81, 0xab, 0x70, 0xc3, 0xce, 0xe5, 0x78, 0xe8, 0x7f, 0x89, 0xed, 0x6c, 0x09,
	0xc7, 0x5b, 0xe3, 0xe4, 0x25, 0xcc, 0xee, 0xaf, 0x52, 0x0e, 0xcf, 0x83, 0xfb, 0xd6, 0x75, 0x45,
	0x9e, 0x43, 0xca, 0x9b, 0x4a, 0x49, 0x23, 0x82, 0x6e, 0xf7, 0xd2, 0x4d, 0x63, 0x71, 0x06, 0x73,
	0x81, 0xb7, 0xa3, 0xa5, 0x2c, 0xd2, 0x3e, 0xb4, 0xc7, 0x65, 0xf2, 0xfd, 0x20, 0x7e, 0x6d, 0x57,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xf2, 0x70, 0xa7, 0x0a, 0x03, 0x00, 0x00,
}
