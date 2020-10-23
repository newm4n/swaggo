package v200

type Swagger struct {
	Swagger             string                     `json:"swagger"`
	Info                *Info                      `json:"info"`
	Host                string                     `json:"host"`
	BasePath            string                     `json:"basePath"`
	Schemes             []string                   `json:"schemes"`
	Consumes            []string                   `json:"consumes"`
	Produces            []string                   `json:"produces"`
	Paths               map[string]*PathItem       `json:"paths"`
	Definitions         map[string]*Schema         `json:"definitions"`
	Parameters          map[string]*Parameter      `json:"parameters"`
	Responses           map[string]*Response       `json:"responses"`
	SecurityDefinitions map[string]*SecurityScheme `json:"securityDefinitions"`
	Security            map[string][]string        `json:"security"`
	Tags                []*Tag                     `json:"tags"`
	ExternalDocs        *ExternalDocumentation     `json:"externalDocs"`
}

type Tag struct {
	Name         string                 `json:"name"`
	Description  string                 `json:"description"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs"`
}

type SecurityScheme struct {
	Reference
	Type             string            `json:"type"`
	Description      string            `json:"description"`
	Name             string            `json:"name"`
	In               string            `json:"in"`
	Flow             string            `json:"flow"`
	AuthorizationURL string            `json:"authorizationUrl"`
	TokenURL         string            `json:"tokenUrl"`
	Scopes           map[string]string `json:"scopes"`
}

type Response struct {
	Reference
	Description string                 `json:"description"`
	Schema      *Schema                `json:"schema"`
	Headers     map[string]*Header     `json:"headers"`
	Example     map[string]interface{} `json:"example"`
}

type Header struct {
	Reference
	Description      string        `json:"description"`
	Type             string        `json:"type"`
	Format           string        `json:"format"`
	Items            *Items        `json:"items"`
	CollectionFormat string        `json:"collectionFormat"`
	Default          interface{}   `json:"default"`
	Maximum          int           `json:"maximum"`
	ExclusiveMaximum bool          `json:"exclusiveMaximum"`
	Minimum          int           `json:"minimum"`
	ExclusiveMinimum bool          `json:"exclusiveMinimum"`
	MaxLength        int           `json:"maxLength"`
	MinLength        int           `json:"minLength"`
	Pattern          string        `json:"pattern"`
	MaxItems         int           `json:"maxItems"`
	MinItems         int           `json:"minItems"`
	UniqueItems      bool          `json:"uniqueItems"`
	Enum             []interface{} `json:"enum"`
	MultipleOf       int           `json:"multipleOf"`
}

type Parameter struct {
	Reference
	Name             string        `json:"name"`
	In               string        `json:"in"`
	Description      string        `json:"description"`
	Required         bool          `json:"required"`
	Schema           *Schema       `json:"schema"`
	Type             string        `json:"type"`
	Format           string        `json:"format"`
	AllowEmptyValue  bool          `json:"allowEmptyValue"`
	Items            *Items        `json:"items"`
	CollectionFormat string        `json:"collectionFormat"`
	Default          interface{}   `json:"default"`
	Maximum          int           `json:"maximum"`
	ExclusiveMaximum bool          `json:"exclusiveMaximum"`
	Minimum          int           `json:"minimum"`
	ExclusiveMinimum bool          `json:"exclusiveMinimum"`
	MaxLength        int           `json:"maxLength"`
	MinLength        int           `json:"minLength"`
	Pattern          string        `json:"pattern"`
	MaxItems         int           `json:"maxItems"`
	MinItems         int           `json:"minItems"`
	UniqueItems      bool          `json:"uniqueItems"`
	Enum             []interface{} `json:"enum"`
	MultipleOf       int           `json:"multipleOf"`
}

type Items struct {
	Type             string        `json:"type"`
	Format           string        `json:"format"`
	AllowEmptyValue  bool          `json:"allowEmptyValue"`
	Items            *Items        `json:"items"`
	CollectionFormat string        `json:"collection_format"`
	Default          interface{}   `json:"default"`
	Maximum          int           `json:"maximum"`
	ExclusiveMaximum bool          `json:"exclusiveMaximum"`
	Minimum          int           `json:"minimum"`
	ExclusiveMinimum bool          `json:"exclusiveMinimum"`
	MaxLength        int           `json:"maxLength"`
	MinLength        int           `json:"minLength"`
	Pattern          string        `json:"pattern"`
	MaxItems         int           `json:"maxItems"`
	MinItems         int           `json:"minItems"`
	UniqueItems      bool          `json:"uniqueItems"`
	Enum             []interface{} `json:"enum"`
	MultipleOf       int           `json:"multipleOf"`
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
	Parameters  []*Parameter `json:"parameters"`
}

type Reference struct {
	Ref string `json:"$ref"`
}

type Operation struct {
	Tags         []string               `json:"tags"`
	Summary      string                 `json:"summary"`
	Description  string                 `json:"description"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs"`
	OperationID  string                 `json:"operationId"`
	Consumes     []string               `json:"consumes"`
	Produces     []string               `json:"produces"`
	Parameters   []*Parameter           `json:"parameters"`
	Responses    map[string]*Response   `json:"responses"`
	Schemes      []string               `json:"schemes"`
	Deprecated   bool                   `json:"deprecated"`
	Security     map[string][]string    `json:"security"`
}

type Schema struct {
	Reference
	Format               string                 `json:"format"` //date-time,email, hostname,ipv4, ipv6,uri,uriref
	Title                string                 `json:"title"`
	Description          string                 `json:"description"`
	Default              string                 `json:"default"`
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
	Items                []*Schema              `json:"items"`
	AllOf                []*Schema              `json:"allOf"`
	Properties           map[string]*Schema     `json:"properties"`
	AdditionalProperties map[string]*Schema     `json:"additionalProperties"`
	Discriminator        string                 `json:"discriminator"`
	ReadOnly             bool                   `json:"readOnly"`
	Xml                  *XML                   `json:"xml"`
	ExternalDocs         *ExternalDocumentation `json:"externalDocs"`
	Example              map[string]interface{} `json:"example"`
}

type XML struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Prefix    string `json:"prefix"`
	Attribute bool   `json:"attribute"`
	Wrapped   bool   `json:"wrapped"`
}

type ExternalDocumentation struct {
	Description string `json:"description"`
	URL         string `json:"url"`
}
