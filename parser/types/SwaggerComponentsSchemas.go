package parser

type SwaggerComponentsSchemas map[string]struct {
	Type                 *SwaggerSchemaType                  `json:"type,omitempty"`
	Format               *string                             `json:"format,omitempty"`
	Description          *string                             `json:"description,omitempty"`
	Enum                 *[]int                              `json:"enum,omitempty"`
	XEnumNames           *[]string                           `json:"x-enumNames,omitempty"`
	Properties           *SwaggerComponentsSchemasProperties `json:"properties,omitempty"`
	AdditionalProperties *bool                               `json:"additionalProperties,omitempty"`
	Required             *[]string                           `json:"required,omitempty"`
}
