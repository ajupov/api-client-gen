package parser

type SwaggerMediaType struct {
	Schema   *SwaggerSchemaOrSwaggerReference             `json:"schema,omitempty"`
	Example  *string                                      `json:"example,omitempty"`
	Examples *map[string]SwaggerExampleOrSwaggerReference `json:"examples,omitempty"`
	Encoding *map[string]SwaggerEncoding                  `json:"encoding,omitempty"`
}
