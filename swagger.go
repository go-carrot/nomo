package spec

type SwaggerVersion string
type Scheme string

const (
	SwaggerVersion20 SwaggerVersion = "2.0"

	SchemeHTTP  Scheme = "http"
	SchemeHTTPS        = "https"
	SchemeWS           = "ws"
	SchemeWSS          = "wss"
)

// Swagger Object: https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#swagger-object
type Swagger struct {
	Swagger     SwaggerVersion    `json:"swagger"`
	Info        Info              `json:"info"`
	Host        string            `json:"host,omitempty"`
	BasePath    string            `json:"basePath,omitempty"`
	Schemes     []Scheme          `json:"schemes,omitempty"`
	Consumes    []string          `json:"consumes,omitempty"`
	Produces    []string          `json:"produces,omitempty"`
	Paths       map[string]Path   `json:"paths"`
	Definitions map[string]Schema `json:"definitions,omitempty"`
	// TODO parameters
	// TODO responses
	// TODO securityDefinitions
	// TODO security
	// TODO tags
	ExternalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
}
