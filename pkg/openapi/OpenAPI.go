package openapi

import (
	"golang.org/x/net/http2"
)

type Document struct {
	OpenAPI      string                 `json:"openapi"`
	Info         *Info                  `json:"info"`
	Servers      *Server                `json:"servers"`
	Paths        map[string]*PathItem   `json:"paths"`
	Components   *Components            `json:"components"`
	Security     *SecurityRequirement   `json:"security"`
	Tags         *Tag                   `json:"tags"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs"`
	Schemes      []string               `json:"schemes"`
}

type Info struct {
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	TermsOfService string   `json:"termsOfService"`
	Contact        *Contact `json:"contact"`
	License        *License `json:"license"`
	Version        string   `json:"version"`
}

type Contact struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Email string `json:"email"`
}

type License struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Server struct {
	Url         string                     `json:"url"`
	Description string                     `json:"description"`
	Variables   map[string]*ServerVariable `json:"variables"`
}

type ServerVariable struct {
	Enum        []string `json:"enum"`
	Default     string   `json:"default"`
	Description string   `json:"description"`
}

type Reference struct {
	Ref string `json:"$ref"`
}

type PathItem struct {
	Ref         string       `json:"$ref"`
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

type Operation struct {
	Tags         []string               `json:"tags"`
	Summary      string                 `json:"summary"`
	Description  string                 `json:"description"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs"`
	OperationID  string                 `json:"operationId"`
	Parameters   []*Parameter           `json:"parameters"`
	RequestBody  *RequestBody           `json:"requestBody"`
	Responses    *Responses             `json:"responses"`
	Callbacks    map[string]*Callback   `json:"callbacks"`
	Deprecated   bool                   `json:"deprecated"`
	Security     *SecurityRequirement   `json:"security"`
	Servers      []*Server              `json:"servers"`
}

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
	Examples        []*Example            `json:"examples"`
	Content         map[string]*MediaType `json:"content"`
}

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
	Discriminator        *Disriminator          `json:"discriminator"`
	ReadOnly             bool                   `json:"readOnly"`
	WriteOnly            bool                   `json:"writeOnly"`
	Xml                  *XmlObject             `json:"xml"`
	ExternalDocs         *ExternalDocumentation `json:"externalDocs"`
	Example              string                 `json:"example"`
	Deprecated           bool                   `json:"deprecated"`
}

type Components struct {
	Schema          map[string]*Schema         `json:"schema"`
	Responses       map[string]*Response       `json:"responses"`
	Parameters      map[string]*Parameter      `json:"parameters"`
	Examples        map[string]*Example        `json:"examples"`
	RequestBodies   map[string]*RequestBody    `json:"requestBodies"`
	Headers         map[string]*Header         `json:"headers"`
	SecuritySchemes map[string]*SecurityScheme `json:"securitySchemes"`
	Links           map[string]*Link           `json:"links"`
	Callbacks       map[string]*Callback       `json:"callbacks"`
}

type Response struct {
	Description string
	Headers     map[string]*Header
	Content     map[string]*MediaType
	Links       map[string]*Link
}

type Callback struct {
}
