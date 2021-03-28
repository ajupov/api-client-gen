package main

type SwaggerPathParameters struct {
	name        string
	in          string
	required    bool
	description string
	schema      SwaggerPathParametersSchema
	xEnumNames  []string `json:"x-enumNames"`
}
