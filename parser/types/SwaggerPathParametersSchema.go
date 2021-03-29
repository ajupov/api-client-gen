package parser

type SwaggerPathParametersSchema struct {
	Type   SwaggerPathParametersSchemaType `json:"type"`
	Format string                          `json:"format"`
	Ref    string                          `json:"$ref"`
}
