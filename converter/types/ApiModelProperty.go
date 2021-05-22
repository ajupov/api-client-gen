package converter

type ApiModelProperty struct {
	Name               string
	Type               string
	IsArrayOfType      bool
	IsDictionaryOfType bool
	Nullable           bool
}
