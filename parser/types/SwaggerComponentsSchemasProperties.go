package parser

type SwaggerComponentsSchemasProperties map[string]struct {
	Type                 *SwaggerSchemaType  `json:"type,omitempty"`
	Format               *string             `json:"format,omitempty"`
	Ref                  *string             `json:"$ref,omitempty"`
	Items                *SwaggerSchemaItems `json:"items,omitempty"`
	Nullable             *bool               `json:"nullable,omitempty"`
	AdditionalProperties *SwaggerSchema      `json:"additionalProperties,omitempty"`
}
