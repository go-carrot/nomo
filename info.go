package spec

// Info Object: https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#infoObject
type Info struct {
	Title          string  `json:"title"`
	Description    string  `json:"description,omitempty"`
	TermsOfService string  `json:"termsOfService,omitempty"`
	Contact        Contact `json:"contact,omitempty"`
	License        License `json:"license,omitempty"`
	Version        string  `json:"version"`
}
