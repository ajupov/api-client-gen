package parser

type SwaggerSchemaOrSwaggerReference struct {
	Type                 *string                                     `json:"type,omitempty"`
	AllOf                *SwaggerSchemaOrSwaggerReference            `json:"allOf,omitempty"`
	OneOf                *SwaggerSchemaOrSwaggerReference            `json:"oneOf,omitempty"`
	AnyOf                *SwaggerSchemaOrSwaggerReference            `json:"anyOf,omitempty"`
	Not                  *SwaggerSchemaOrSwaggerReference            `json:"not,omitempty"`
	Items                *SwaggerSchemaOrSwaggerReference            `json:"items,omitempty"`
	Properties           *map[string]SwaggerSchemaOrSwaggerReference `json:"properties,omitempty"`
	AdditionalProperties *SwaggerSchemaOrSwaggerReference            `json:"additionalProperties,omitempty"`
	Description          *string                                     `json:"description,omitempty"`
	Format               *string                                     `json:"format,omitempty"`
	Default              *string                                     `json:"default,omitempty"`
	Nullable             *bool                                       `json:"nullable,omitempty"`
	Discriminator        *SwaggerSchemaDiscriminator                 `json:"discriminator,omitempty"`
	ReadOnly             *bool                                       `json:"readOnly,omitempty"`
	WriteOnly            *bool                                       `json:"writeOnly,omitempty"`
	Xml                  *SwaggerXml                                 `json:"xml,omitempty"`
	ExternalDocs         *SwaggerExternalDocumentation               `json:"externalDocs,omitempty"`
	Example              *string                                     `json:"example,omitempty"`
	Deprecated           *bool                                       `json:"deprecated,omitempty"`
	Content              *map[string]SwaggerMediaType                `json:"content,omitempty"`
	Required             *[]string                                   `json:"required,omitempty"`
	Enum                 *[]int                                      `json:"enum,omitempty"`
	XEnumNames           *[]string                                   `json:"x-enumNames,omitempty"`
	Ref                  *string                                     `json:"$ref,omitempty"`
}
