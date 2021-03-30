package parser

type SwaggerComponentsSchemas map[string]struct {
	Enum                 *[]int                              `json:"enum,omitempty"`
	Type                 *SwaggerSchemaType                  `json:"type,omitempty"`
	Required             *[]string                           `json:"required,omitempty"`
	Description          *string                             `json:"description,omitempty"`
	Format               *string                             `json:"format,omitempty"`
	XEnumNames           *[]string                           `json:"x-enumNames,omitempty"`
	Properties           *SwaggerComponentsSchemasProperties `json:"properties,omitempty"`
	AdditionalProperties *bool                               `json:"additionalProperties,omitempty"`
}
