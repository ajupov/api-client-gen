package parser

type SwaggerEncoding struct {
	ContentType   *string                                     `json:"contentType,omitempty"`
	Headers       *map[string]SwaggerHeaderOrSwaggerReference `json:"headers,omitempty"`
	Style         *string                                     `json:"style,omitempty"`
	Explode       *bool                                       `json:"explode,omitempty"`
	AllowReserved *bool                                       `json:"allowReserved,omitempty"`
}
