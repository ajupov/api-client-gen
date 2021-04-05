package parser

type SwaggerOAuthFlow struct {
	AuthorizationUrl *string            `json:"authorizationUrl,omitempty"`
	TokenUrl         *string            `json:"tokenUrl,omitempty"`
	RefreshUrl       *string            `json:"refreshUrl,omitempty"`
	Scopes           *map[string]string `json:"scopes,omitempty"`
}
