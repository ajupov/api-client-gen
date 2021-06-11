package converter

type ApiClientMethodParameterOrBody struct {
	Name               string
	Type               string
	IsInPath           bool
	IsInQuery          bool
	IsArrayOfType      bool
	IsDictionaryOfType bool
	Nullable           bool
}
