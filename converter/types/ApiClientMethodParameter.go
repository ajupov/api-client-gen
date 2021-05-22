package converter

type ApiClientMethodParameter struct {
	Name               string
	Type               string
	IsArrayOfType      bool
	IsDictionaryOfType bool
	Nullable           bool
}
