package converter

type ApiModel struct {
	Name       string
	Imports    []string
	IsEnum     bool
	Properties []ApiModelProperty
	EnumItems  []ApiModelEnumItem
}
