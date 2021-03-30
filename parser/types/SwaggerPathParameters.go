package parser

type SwaggerPathParameters struct {
	Name        *string        `json:"name,omitempty"`
	In          *string        `json:"in,omitempty"`
	Required    *bool          `json:"required,omitempty"`
	Description *string        `json:"description,omitempty"`
	Schema      *SwaggerSchema `json:"schema,omitempty"`
	XEnumNames  *[]string      `json:"x-enumNames,omitempty"`
}
