package parser

type SwaggerTag struct {
	Name         *string                       `json:"name,omitempty"`
	Description  *string                       `json:"description,omitempty"`
	ExternalDocs *SwaggerExternalDocumentation `json:"externalDocs,omitempty"`
}
