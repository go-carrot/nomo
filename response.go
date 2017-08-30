package nomo 

// Response Object: https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#responseObject
type Response struct {
	Description string            `json:"description"`
	Schema      Schema            `json:"schema,omitempty"`
	Headers     map[string]Header `json:"headers,omitempty"`
}
