package parser

type SwaggerPathResponsesContent map[SwaggerPathResponsesContentType]struct {
	Schema SwaggerPathResponsesContentSchema `json:"schema"`
}
