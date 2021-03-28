package main

type SwaggerPathResponsesContentSchema map[string]struct {
	Type                 SwaggerPathResponsesContentSchemaType `json:"type"`
	items                SwaggerPathResponsesContentSchemaItems
	additionalProperties SwaggerPathResponsesContentSchemaAdditionalProperties
}
