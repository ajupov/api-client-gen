package parser

type SwaggerSecuritySchemeOrReference struct {
	Type             *string            `json:"type,omitempty"`
	Description      *string            `json:"description,omitempty"`
	Name             *string            `json:"name,omitempty"`
	In               *string            `json:"in,omitempty"`
	Scheme           *string            `json:"scheme,omitempty"`
	BearerFormat     *string            `json:"bearerFormat,omitempty"`
	Flows            *SwaggerOAuthFlows `json:"flows,omitempty"`
	OpenIdConnectUrl *string            `json:"openIdConnectUrl,omitempty"`
	Ref              *string            `json:"$ref,omitempty"`
}
