package parser

type SwaggerResponseOrSwaggerReference struct {
	Description *string                                     `json:"description,omitempty"`
	Headers     *map[string]SwaggerHeaderOrSwaggerReference `json:"headers,omitempty"`
	Content     *map[string]SwaggerMediaType                `json:"content,omitempty"`
	Links       *map[string]SwaggerLinkOrSwaggerReference   `json:"links,omitempty"`
	Ref         *string                                     `json:"$ref,omitempty"`
}
