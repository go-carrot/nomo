package nomo 

// Path Item Object: https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#pathItemObject
type Path struct {
	Get    Operation `json:"get,omitempty"`
	Put    Operation `json:"put,omitempty"`
	Post   Operation `json:"post,omitempty"`
	Delete Operation `json:"delete,omitempty"`
}
