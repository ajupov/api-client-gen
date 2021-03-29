package parser

type SwaggerPathParameters struct {
	Name        string                      `json:"name"`
	In          string                      `json:"in"`
	Required    bool                        `json:"required"`
	Description string                      `json:"description"`
	Schema      SwaggerPathParametersSchema `json:"schema"`
	XEnumNames  []string                    `json:"x-enumNames"`
}
