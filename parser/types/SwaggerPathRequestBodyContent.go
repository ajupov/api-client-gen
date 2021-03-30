package parser

type SwaggerPathRequestBodyContent map[SwaggerPathResponsesContentType]struct {
	Schema *SwaggerSchema `json:"schema,omitempty"`
}
