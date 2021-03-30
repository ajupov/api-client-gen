package parser

type SwaggerPathRequestBody struct {
	Content  *SwaggerPathRequestBodyContent `json:"content,omitempty"`
	Required *bool                          `json:"required,omitempty"`
}
