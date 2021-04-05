package parser

type SwaggerExampleOrSwaggerReference struct {
	Summary       *string `json:"summary,omitempty"`
	Description   *string `json:"description,omitempty"`
	Value         *string `json:"value,omitempty"`
	ExternalValue *string `json:"externalValue,omitempty"`
	Ref           *string `json:"$ref,omitempty"`
}
