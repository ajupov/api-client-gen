package parser

type SwaggerComponent struct {
	Openapi      *string                       `json:"openapi,omitempty"`
	Info         *SwaggerInfo                  `json:"info,omitempty"`
	Servers      *[]SwaggerServer              `json:"servers,omitempty"`
	Paths        *SwaggerPaths                 `json:"paths,omitempty"`
	Components   *SwaggerComponents            `json:"components,omitempty"`
	Security     *[]SwaggerSecurityRequirement `json:"security,omitempty"`
	Tags         *[]SwaggerTag                 `json:"tags,omitempty"`
	ExternalDocs *SwaggerExternalDocumentation `json:"externalDocs,omitempty"`
}
