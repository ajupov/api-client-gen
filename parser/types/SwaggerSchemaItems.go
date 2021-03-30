package parser

type SwaggerSchemaItems struct {
	Ref    *string            `json:"$ref,omitempty"`
	Type   *SwaggerSchemaType `json:"type,omitempty"`
	Format *string            `json:"format,omitempty"`
}
