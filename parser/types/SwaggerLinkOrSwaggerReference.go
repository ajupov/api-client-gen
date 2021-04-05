package parser

type SwaggerLinkOrSwaggerReference struct {
	OperationRef *string            `json:"operationRef,omitempty"`
	OperationId  *string            `json:"operationId,omitempty"`
	Parameters   *map[string]string `json:"parameters,omitempty"`
	RequestBody  *map[string]string `json:"requestBody,omitempty"`
	Description  *string            `json:"description,omitempty"`
	Server       *SwaggerServer     `json:"server,omitempty"`
	Ref          *string            `json:"$ref,omitempty"`
}
