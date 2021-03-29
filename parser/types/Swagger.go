package parser

// Swagger
type Swagger struct {
	Openapi string       `json:"openapi"`
	Info    SwaggerInfo  `json:"info"`
	Paths   SwaggerPaths `json:"paths"`
}
