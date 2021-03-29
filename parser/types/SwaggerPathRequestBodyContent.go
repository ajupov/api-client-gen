package parser

type SwaggerPathRequestBodyContent map[SwaggerPathResponsesContentType]struct {
	Schema SwaggerPathResponsesContentSchema `json:"schema"`
}
