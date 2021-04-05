package parser

type SwaggerParameterOrReference struct {
	Name            *string                                      `json:"name,omitempty"`
	In              *string                                      `json:"in,omitempty"`
	Description     *string                                      `json:"description,omitempty"`
	Required        *bool                                        `json:"required,omitempty"`
	Deprecated      *bool                                        `json:"deprecated,omitempty"`
	AllowEmptyValue *bool                                        `json:"allowEmptyValue,omitempty"`
	Style           *string                                      `json:"style,omitempty"`
	Explode         *bool                                        `json:"explode,omitempty"`
	AllowReserved   *bool                                        `json:"allowReserved,omitempty"`
	Schema          *SwaggerSchemaOrSwaggerReference             `json:"schema,omitempty"`
	Example         *string                                      `json:"example,omitempty"`
	Examples        *map[string]SwaggerExampleOrSwaggerReference `json:"examples,omitempty"`
	Content         *map[string]SwaggerMediaType                 `json:"content,omitempty"`
	XEnumNames      *[]string                                    `json:"x-enumNames,omitempty"`
	Ref             *string                                      `json:"$ref,omitempty"`
}
