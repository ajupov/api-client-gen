package parser

type SwaggerPathResponses map[SwaggerPathResponsesStatusCode]struct {
	Description *string                                                          `json:"description,omitempty"`
	Content     *map[SwaggerPathResponsesContentType]SwaggerPathResponsesContent `json:"content,omitempty"`
}
