package parser

type SwaggerPaths map[SwaggerPathsMethod]struct {
	Path SwaggerPath `json:"path"`
}
