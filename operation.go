package nomo 

// Operation Object: https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#operationObject
type Operation struct {
	Tags         []string              `json:"tags,omitempty"`
	Summary      string                `json:"summary,omitempty"`
	Description  string                `json:"description,omitempty"`
	ExternalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
	OperationID  string                `json:"operationId,omitempty"`
	Consumes     []string              `json:"consumes,omitempty"`
	Produces     []string              `json:"produces,omitempty"`
	Parameters   []Parameter           `json:"parameters,omitempty"`
	Responses    map[string]Response   `json:"responses"`
	Schemes      []Scheme              `json:"schemes,omitempty"`
	Deprecated   bool                  `json:"deprecated,omitempty"`
	// TODO security
}
