package parser

type SwaggerHeaderOrSwaggerReference struct {
	Description *string                          `json:"description,omitempty"`
	Schema      *SwaggerSchemaOrSwaggerReference `json:"schema,omitempty"`
	Ref         *string                          `json:"$ref,omitempty"`
}
