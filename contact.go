package nomo 

// Contact Object: https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#contactObject
type Contact struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}
