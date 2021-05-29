package templater

type Config struct {
	ApiClientTemplate      string    `json:"ApiClientTemplate"`
	ApiModelTemplate       string    `json:"ApiModelTemplate"`
	ApiClientDirectory     string    `json:"ApiClientDirectory"`
	ApiModelDirectory      string    `json:"ApiModelDirectory"`
	ApiClientFileExtension string    `json:"ApiClientFileExtension"`
	ApiModelFileExtension  string    `json:"ApiModelFileExtension"`
	CopyWithoutTemplating  *[]string `json:"CopyWithoutTemplating,omitempty"`
}
