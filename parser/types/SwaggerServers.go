package parser

type SwaggerServer struct {
	Url         *string                `json:"url,omitempty"`
	Description *string                `json:"description,omitempty"`
	Variables   *SwaggerServerVariable `json:"variables,omitempty"`
}
