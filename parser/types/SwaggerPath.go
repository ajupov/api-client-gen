package parser

type SwaggerPath map[SwaggerPathsMethod]struct {
	Tags        *[]string                `json:"tags,omitempty"`
	Parameters  *[]SwaggerPathParameters `json:"parameters,omitempty"`
	RequestBody *SwaggerPathRequestBody  `json:"requestBody,omitempty"`
	Responses   *SwaggerPathResponses    `json:"responses,omitempty"`
}
