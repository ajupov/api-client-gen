package parser

type SwaggerServerVariable map[string]struct {
	Enum        *[]string `json:"enum,omitempty"`
	Default     *string   `json:"default,omitempty"`
	Description *string   `json:"description,omitempty"`
}
