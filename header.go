package spec

// Header Object: https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#headerObject
type Header struct {
	Description string `json:"description,omitempty"`
	Type        string `json:"type"`
	// TODO, other props
}
