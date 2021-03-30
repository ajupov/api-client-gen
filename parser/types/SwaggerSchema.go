package parser

type SwaggerSchema struct {
	Type                 *SwaggerSchemaType                 `json:"type,omitempty"`
	Format               *string                            `json:"format,omitempty"`
	Ref                  *string                            `json:"$ref,omitempty"`
	Items                *SwaggerSchemaItems                `json:"items,omitempty"`
	AdditionalProperties *SwaggerSchemaAdditionalProperties `json:"additionalProperties,omitempty"`
}
