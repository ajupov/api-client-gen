package converter

type ApiClientMethodRequestBody struct {
	Type               string
	IsArrayOfType      bool
	IsDictionaryOfType bool
	Nullable           bool
}
