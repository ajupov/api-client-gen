package parser

type SwaggerPathResponses map[SwaggerPathResponsesStatusCode]struct {
	Description string                      `json:"description"`
	Content     SwaggerPathResponsesContent `json:"content"`
}
