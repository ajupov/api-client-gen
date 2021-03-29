package parser

type SwaggerPathResponsesContentSchema map[string]struct {
	Type                 SwaggerPathResponsesContentSchemaType                 `json:"type"`
	Items                SwaggerPathResponsesContentSchemaItems                `json:"items"`
	AdditionalProperties SwaggerPathResponsesContentSchemaAdditionalProperties `json:"additionalProperties"`
}
