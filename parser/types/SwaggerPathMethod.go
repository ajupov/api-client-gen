package parser

type SwaggerPathsMethod string

const (
	Get    SwaggerPathsMethod = "get"
	Post   SwaggerPathsMethod = "post"
	Put    SwaggerPathsMethod = "put"
	Patch  SwaggerPathsMethod = "patch"
	Delete SwaggerPathsMethod = "delete"
)
