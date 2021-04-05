package parser

type SwaggerOAuthFlows struct {
	Implicit          *SwaggerOAuthFlow `json:"implicit,omitempty"`
	Password          *SwaggerOAuthFlow `json:"password,omitempty"`
	ClientCredentials *SwaggerOAuthFlow `json:"clientCredentials,omitempty"`
	AuthorizationCode *SwaggerOAuthFlow `json:"authorizationCode,omitempty"`
}
