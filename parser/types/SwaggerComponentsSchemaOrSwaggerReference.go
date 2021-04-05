package parser

type SwaggerComponentsSchemaOrSwaggerReference struct {
	Type                 *string                                     `json:"type,omitempty"`
	AllOf                *SwaggerSchemaOrSwaggerReference            `json:"allOf,omitempty"`
	OneOf                *SwaggerSchemaOrSwaggerReference            `json:"oneOf,omitempty"`
	AnyOf                *SwaggerSchemaOrSwaggerReference            `json:"anyOf,omitempty"`
	Not                  *SwaggerSchemaOrSwaggerReference            `json:"not,omitempty"`
	Items                *SwaggerSchemaOrSwaggerReference            `json:"items,omitempty"`
	Properties           *map[string]SwaggerSchemaOrSwaggerReference `json:"properties,omitempty"`
	AdditionalProperties *bool                                       `json:"additionalProperties,omitempty"`
	Description          *string                                     `json:"description,omitempty"`
	Format               *string                                     `json:"format,omitempty"`
	Default              *string                                     `json:"default,omitempty"`
	Content              *map[string]SwaggerMediaType                `json:"content,omitempty"`
	Required             *[]string                                   `json:"required,omitempty"`
	Enum                 *[]int                                      `json:"enum,omitempty"`
	XEnumNames           *[]string                                   `json:"x-enumNames,omitempty"`
	Nullable             *bool                                       `json:"nullable,omitempty"`
	Ref                  *string                                     `json:"$ref,omitempty"`
}
