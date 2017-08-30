package nomo 

type ParameterIn string

const (
	ParameterInQuery    ParameterIn = "query"
	ParameterInHeader               = "header"
	ParameterInPath                 = "path"
	ParameterInFormData             = "formData"
	ParameterInBody                 = "body"
)

// Parameter Object: https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#parameterObject
type Parameter struct {
	Name        string      `json:"name"`
	In          ParameterIn `json:"in"`
	Description string      `json:"description,omitempty"`
	Required    bool        `json:"required,omitempty"`
}
