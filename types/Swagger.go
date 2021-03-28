package main

type Swagger struct {
	Openapi string       `json:"openapi"`
	Info    SwaggerInfo  `json:"info"`
	Paths   SwaggerPaths `json:"paths"`
}
