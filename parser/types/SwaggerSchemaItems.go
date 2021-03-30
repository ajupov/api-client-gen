package parser

type SwaggerSchemaItems struct {
	Type   *SwaggerSchemaType `json:"type,omitempty"`
	Format *string            `json:"format,omitempty"`
	Ref    *string            `json:"$ref,omitempty"`
}
