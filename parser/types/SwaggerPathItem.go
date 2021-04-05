package parser

type SwaggerPathItem struct {
	Ref         *string                        `json:"$ref,omitempty"`
	Summary     *string                        `json:"summary,omitempty"`
	Description *string                        `json:"description,omitempty"`
	Get         *SwaggerOperation              `json:"get,omitempty"`
	Put         *SwaggerOperation              `json:"put,omitempty"`
	Post        *SwaggerOperation              `json:"post,omitempty"`
	Delete      *SwaggerOperation              `json:"delete,omitempty"`
	Options     *SwaggerOperation              `json:"options,omitempty"`
	Head        *SwaggerOperation              `json:"head,omitempty"`
	Patch       *SwaggerOperation              `json:"patch,omitempty"`
	Trace       *SwaggerOperation              `json:"trace,omitempty"`
	Servers     *[]SwaggerServer               `json:"servers,omitempty"`
	Parameters  *[]SwaggerParameterOrReference `json:"parameters,omitempty"`
}
