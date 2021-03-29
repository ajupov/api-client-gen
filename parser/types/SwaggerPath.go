package parser

type SwaggerPath struct {
	Tags        []string                `json:"tags"`
	Parameters  []SwaggerPathParameters `json:"parameters"`
	RequestBody SwaggerPathRequestBody  `json:"requestBody"`
	Responses   SwaggerPathResponses    `json:"responses"`
}
