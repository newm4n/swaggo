// http://spec.openapis.org/oas/v3.0.3
package v303

// OpenAPI is the root document object of the OpenAPI document.
// http://spec.openapis.org/oas/v3.0.3#openapi-object
type OpenAPI struct {
	OpenAPI      string                 `json:"openapi"`
	Info         *Info                  `json:"info"`
	Servers      *Server                `json:"servers"`
	Paths        map[string]*PathItem   `json:"paths"`
	Components   *Components            `json:"components"`
	Security     map[string][]string    `json:"security"`
	Tags         *Tag                   `json:"tags"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs"`
	Schemes      []string               `json:"schemes"`
}

// Info provides metadata about the API. The metadata MAY be used by the clients if needed, and MAY be presented in editing or documentation generation tools for convenience.
// http://spec.openapis.org/oas/v3.0.3#info-object
type Info struct {
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	TermsOfService string   `json:"termsOfService"`
	Contact        *Contact `json:"contact"`
	License        *License `json:"license"`
	Version        string   `json:"version"`
}

// Tag Adds metadata to a single tag that is used by the Operation Object. It is not mandatory to have a Tag Object per tag defined in the Operation Object instances.
// http://spec.openapis.org/oas/v3.0.3#tag-object
type Tag struct {
	Name         string                 `json:"name"`
	Description  string                 `json:"description"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs"`
}

// Contact information for the exposed API.
// http://spec.openapis.org/oas/v3.0.3#contact-object
type Contact struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Email string `json:"email"`
}

// License information for the exposed API.
// http://spec.openapis.org/oas/v3.0.3#license-object
type License struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// Server An object representing a Server.
// http://spec.openapis.org/oas/v3.0.3#server-object
type Server struct {
	Url         string                     `json:"url"`
	Description string                     `json:"description"`
	Variables   map[string]*ServerVariable `json:"variables"`
}

// ServerVariable An object representing a Server Variable for server URL template substitution.
// http://spec.openapis.org/oas/v3.0.3#server-variable-object
type ServerVariable struct {
	Enum        []string `json:"enum"`
	Default     string   `json:"default"`
	Description string   `json:"description"`
}

// Reference A simple object to allow referencing other components in the specification, internally and externally.
// http://spec.openapis.org/oas/v3.0.3#reference-object
type Reference struct {
	Ref string `json:"$ref"`
}

// PathItem Describes the operations available on a single path. A Path Item MAY be empty, due to ACL constraints. The path itself is still exposed to the documentation viewer but they will not know which operations and parameters are available.
// http://spec.openapis.org/oas/v3.0.3#path-item-object
type PathItem struct {
	Reference
	Summary     string       `json:"summary"`
	Description string       `json:"description"`
	Get         *Operation   `json:"get"`
	Put         *Operation   `json:"put"`
	Post        *Operation   `json:"post"`
	Delete      *Operation   `json:"delete"`
	Options     *Operation   `json:"options"`
	Head        *Operation   `json:"head"`
	Patch       *Operation   `json:"patch"`
	Trace       *Operation   `json:"trace"`
	Servers     []*Server    `json:"server"`
	Parameters  []*Parameter `json:"parameters"`
}

// Operation Describes a single API operation on a path.
// http://spec.openapis.org/oas/v3.0.3#operation-object
type Operation struct {
	Tags         []string                        `json:"tags"`
	Summary      string                          `json:"summary"`
	Description  string                          `json:"description"`
	ExternalDocs *ExternalDocumentation          `json:"externalDocs"`
	OperationID  string                          `json:"operationId"`
	Parameters   []*Parameter                    `json:"parameters"`
	RequestBody  *RequestBody                    `json:"requestBody"`
	Responses    map[string]*Response            `json:"responses"`
	Callbacks    map[string]map[string]*PathItem `json:"callbacks"`
	Deprecated   bool                            `json:"deprecated"`
	Security     map[string][]string             `json:"security"`
	Servers      []*Server                       `json:"servers"`
}

// ExternalDocumentation Allows referencing an external resource for extended documentation.
// http://spec.openapis.org/oas/v3.0.3#external-documentation-object
type ExternalDocumentation struct {
	Description string `json:"description"`
	URL         string `json:"url"`
}

// Describes a single request body.
// http://spec.openapis.org/oas/v3.0.3#request-body-object
type RequestBody struct {
	Reference
	Description string                `json:"description"`
	Content     map[string]*MediaType `json:"content"`
	Required    bool                  `json:"required"`
}

// MediaType Each Media Type Object provides schema and examples for the media type identified by its key.
// http://spec.openapis.org/oas/v3.0.3#media-type-object
type MediaType struct {
	Schema   *Schema              `json:"schema"`
	Example  *Example             `json:"example"`
	Examples map[string]*Example  `json:"examples"`
	Encoding map[string]*Encoding `json:"encoding"`
}

// Encoding A single encoding definition applied to a single schema property.
// http://spec.openapis.org/oas/v3.0.3#encoding-object
type Encoding struct {
	ContentType   string             `json:"content_type"`
	Headers       map[string]*Header `json:"headers"`
	Style         string             `json:"style"`
	Explode       bool               `json:"explode"`
	AllowReserved bool               `json:"allowReserved"`
}

// Parameter Describes a single operation parameter.
// http://spec.openapis.org/oas/v3.0.3#parameter-object
type Parameter struct {
	Reference
	Name            string                `json:"name"`
	In              string                `json:"in"`
	Description     string                `json:"description"`
	Required        bool                  `json:"required"`
	Deprecated      bool                  `json:"deprecated"`
	AllowEmptyValue bool                  `json:"allowEmptyValue"`
	Style           string                `json:"style"`
	Explode         bool                  `json:"explode"`
	AllowReserved   bool                  `json:"allowReserved"`
	Schema          *Schema               `json:"schema"`
	Example         string                `json:"example"`
	Examples        map[string]*Example   `json:"examples"`
	Content         map[string]*MediaType `json:"content"`
}

// Example is simply an example
// http://spec.openapis.org/oas/v3.0.3#example-object
type Example struct {
	Reference
	Summary       string      `json:"summary"`
	Description   string      `json:"description"`
	Value         interface{} `json:"value"`
	ExternalValue string      `json:"externalValue"`
}

// Schema Object allows the definition of input and output data types. These types can be objects, but also primitives and arrays.
// http://spec.openapis.org/oas/v3.0.3#schema-object
type Schema struct {
	Reference
	Title                string                 `json:"title"`
	MultipleOf           int                    `json:"multipleOf"`
	Maximum              int                    `json:"maximum"`
	ExclusiveMaximum     bool                   `json:"exclusiveMaximum"`
	Minimum              int                    `json:"minimum"`
	ExclusiveMinimum     bool                   `json:"exclusiveMinimum"`
	MaxLength            int                    `json:"maxLength"`
	MinLength            int                    `json:"minLength"`
	Pattern              string                 `json:"pattern"`
	MaxItems             int                    `json:"maxItems"`
	MinItems             int                    `json:"minItems"`
	UniqueItems          bool                   `json:"uniqueItems"`
	MaxProperties        int                    `json:"maxProperties"`
	MinProperties        int                    `json:"minProperties"`
	Required             []string               `json:"required"`
	Enum                 []interface{}          `json:"enum"`
	Type                 string                 `json:"type"`
	AllOf                []*Schema              `json:"allOf"`
	OneOf                []*Schema              `json:"oneOf"`
	AnyOf                []*Schema              `json:"anyOf"`
	Not                  *Schema                `json:"not"`
	Items                []*Schema              `json:"items"`
	Properties           map[string]*Schema     `json:"properties"`
	AdditionalProperties map[string]*Schema     `json:"additionalProperties"`
	Description          string                 `json:"description"`
	Format               string                 `json:"format"` //date-time,email, hostname,ipv4, ipv6,uri,uriref
	Default              string                 `json:"default"`
	Nullable             bool                   `json:"nullable"`
	Discriminator        *Discriminator         `json:"discriminator"`
	ReadOnly             bool                   `json:"readOnly"`
	WriteOnly            bool                   `json:"writeOnly"`
	Xml                  *XML                   `json:"xml"`
	ExternalDocs         *ExternalDocumentation `json:"externalDocs"`
	Example              *Example               `json:"example"`
	Deprecated           bool                   `json:"deprecated"`
}

// XML A metadata object that allows for more fine-tuned XML model definitions.
// http://spec.openapis.org/oas/v3.0.3#xml-object
type XML struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Prefix    string `json:"prefix"`
	Attribute bool   `json:"attribute"`
	Wrapped   bool   `json:"wrapped"`
}

// Discriminator When request bodies or response payloads may be one of a number of different schemas, a discriminator object can be used to aid in serialization, deserialization, and validation.
// http://spec.openapis.org/oas/v3.0.3#discriminator-object
type Discriminator struct {
	PropertyName string
	Mapping      map[string]string
}

// Componenets Holds a set of reusable objects for different aspects of the OAS
// http://spec.openapis.org/oas/v3.0.3#components-object
type Components struct {
	Schema          map[string]*Schema              `json:"schema"`
	Responses       map[string]*Response            `json:"responses"`
	Parameters      map[string]*Parameter           `json:"parameters"`
	Examples        map[string]*Example             `json:"examples"`
	RequestBodies   map[string]*RequestBody         `json:"requestBodies"`
	Headers         map[string]*Header              `json:"headers"`
	SecuritySchemes map[string]*SecurityScheme      `json:"securitySchemes"`
	Links           map[string]*Link                `json:"links"`
	Callbacks       map[string]map[string]*PathItem `json:"callbacks"`
}

// SecurityScheme Defines a security scheme that can be used by the operations
// http://spec.openapis.org/oas/v3.0.3#security-scheme-object
type SecurityScheme struct {
	Reference
	Type             string      `json:"type"`
	Description      string      `json:"description"`
	Name             string      `json:"name"`
	In               string      `json:"in"`
	Scheme           string      `json:"scheme"`
	BearerFormat     string      `json:"bearerFormat"`
	Flows            *OAuthFlows `json:"flows"`
	OpenIDConnectURL string      `json:"openIdConnectUrl"`
}

// OAuthFlows Allows configuration of the supported OAuth Flows.
// http://spec.openapis.org/oas/v3.0.3#oauth-flows-object
type OAuthFlows struct {
	Implicit          *OAuthFlow `json:"implicit"`
	Password          *OAuthFlow `json:"password"`
	ClientCredentials *OAuthFlow `json:"clientCredentials"`
	AuthorizationCode *OAuthFlow `json:"authorizationCode"`
}

// OAuthFlow Configuration details for a supported OAuth Flow
// http://spec.openapis.org/oas/v3.0.3#oauth-flow-object
type OAuthFlow struct {
	AuthorizationURL string            `json:"authorizationUrl"`
	TokenURL         string            `json:"tokenUrl"`
	RefreshURL       string            `json:"refreshUrl"`
	Scopes           map[string]string `json:"scopes"`
}

// Response A container for the expected responses of an operation
// http://spec.openapis.org/oas/v3.0.3#responses-object
type Response struct {
	Reference
	Description string                `json:"description"`
	Headers     map[string]*Header    `json:"headers"`
	Content     map[string]*MediaType `json:"content"`
	Links       map[string]*Link      `json:"links"`
}

// Link represents a possible design-time link for a response
// http://spec.openapis.org/oas/v3.0.3#link-object
type Link struct {
	Reference
	OperationRef string
	OperationID  string
	Parameters   map[string]interface{}
	RequestBody  interface{}
	Description  string
	Server       *Server
}

// Header follows the structure of the Parameter Object
// http://spec.openapis.org/oas/v3.0.3#header-object
type Header struct {
	Reference
	Description     string                `json:"description"`
	Required        bool                  `json:"required"`
	Deprecated      bool                  `json:"deprecated"`
	AllowEmptyValue bool                  `json:"allowEmptyValue"`
	Style           string                `json:"style"`
	Explode         bool                  `json:"explode"`
	AllowReserved   bool                  `json:"allowReserved"`
	Schema          *Schema               `json:"schema"`
	Example         string                `json:"example"`
	Examples        map[string]*Example   `json:"examples"`
	Content         map[string]*MediaType `json:"content"`
}
