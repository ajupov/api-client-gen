package parser

type SwaggerComponents struct {
	Schemas         *map[string]SwaggerComponentsSchemaOrSwaggerReference `json:"schemas,omitempty"`
	Responses       *map[string]SwaggerResponseOrSwaggerReference         `json:"responses,omitempty"`
	Parameters      *map[string]SwaggerParameterOrReference               `json:"parameters,omitempty"`
	Examples        *map[string]SwaggerExampleOrSwaggerReference          `json:"examples,omitempty"`
	RequestBodies   *map[string]SwaggerRequestBodyOrReference             `json:"requestBodies,omitempty"`
	Headers         *map[string]SwaggerHeaderOrSwaggerReference           `json:"headers,omitempty"`
	SecuritySchemes *map[string]SwaggerSecuritySchemeOrReference          `json:"securitySchemes,omitempty"`
	Links           *map[string]SwaggerLinkOrSwaggerReference             `json:"links,omitempty"`
	Callbacks       *map[string]SwaggerCallbackOrReference                `json:"callbacks,omitempty"`
}
