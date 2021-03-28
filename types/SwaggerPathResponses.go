package main

type SwaggerPathResponses map[SwaggerPathResponsesStatusCode]struct {
	description string
	content     SwaggerPathResponsesContent
}
