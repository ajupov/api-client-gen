package parser

type SwaggerCallbackOrReference struct {
	Ref *string `json:"$ref,omitempty"`
}
