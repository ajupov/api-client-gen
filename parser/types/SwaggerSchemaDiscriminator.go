package parser

type SwaggerSchemaDiscriminator struct {
	PropertyName *string            `json:"propertyName,omitempty"`
	Mapping      *map[string]string `json:"mapping,omitempty"`
}
