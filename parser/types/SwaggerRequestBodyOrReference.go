package parser

type SwaggerRequestBodyOrReference struct {
	Description *string                      `json:"description,omitempty"`
	Content     *map[string]SwaggerMediaType `json:"content,omitempty"`
	Required    *bool                        `json:"required,omitempty"`
	Ref         *string                      `json:"$ref,omitempty"`
}
