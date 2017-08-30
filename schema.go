package spec

type Type string

const (
	TypeString  Type = "string"
	TypeNumber       = "number"
	TypeInteger      = "integer"
	TypeBoolean      = "boolean"
	TypeArray        = "array"
	TypeFile         = "file"
)

// Schema Object: https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#schemaObject
type Schema struct {
	Ref        string             `json:"ref,omitempty"`
	Type       Type               `json:"type,omitempty"`
	Format     string             `json:"format,omitempty"`
	Properties map[string]*Schema `json:"properties,omitempty"`
	Items      *Schema            `json:"items,omitempty"`
	Enum       []string           `json:"enum,omitempty"`
	Example    interface{}        `json:"example,omitempty"`
}
