package parser

type Swagger struct {
	Openapi    *string            `json:"openapi,omitempty"`
	Info       *SwaggerInfo       `json:"info,omitempty"`
	Paths      *SwaggerPaths      `json:"paths,omitempty"`
	Components *SwaggerComponents `json:"components,omitempty"`
}
