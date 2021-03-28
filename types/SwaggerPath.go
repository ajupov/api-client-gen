package main

type SwaggerPath struct {
	tags        []string
	parameters  []SwaggerPathParameters
	requestBody SwaggerPathRequestBody
	responses   SwaggerPathResponses
}
