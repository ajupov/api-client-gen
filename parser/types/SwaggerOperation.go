package parser

type SwaggerOperation struct {
	Tags         *[]string                                     `json:"tags,omitempty"`
	Summary      *string                                       `json:"summary,omitempty"`
	Description  *string                                       `json:"description,omitempty"`
	ExternalDocs *SwaggerExternalDocumentation                 `json:"externalDocs,omitempty"`
	OperationId  *string                                       `json:"operationId,omitempty"`
	Parameters   *[]SwaggerParameterOrReference                `json:"parameters,omitempty"`
	RequestBody  *SwaggerRequestBodyOrReference                `json:"requestBody,omitempty"`
	Responses    *map[string]SwaggerResponseOrSwaggerReference `json:"responses,omitempty"`
	Callbacks    *map[string]SwaggerPathItem                   `json:"callbacks,omitempty"`
	Deprecated   *bool                                         `json:"deprecated,omitempty"`
	Security     *SwaggerSecurityRequirement                   `json:"security,omitempty"`
	Servers      *[]SwaggerServer                              `json:"servers,omitempty"`
}
