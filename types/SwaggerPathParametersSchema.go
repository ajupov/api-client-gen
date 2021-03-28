package main

type SwaggerPathParametersSchema struct {
	_type  SwaggerPathParametersSchemaType `json:"type"`
	format string
	ref    string `json:"$ref"`
}
