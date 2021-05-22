package converter

type ApiClientMethod struct {
	Name                string
	Url                 string
	Method              string
	RequestContentType  string
	ResponseContentType string
	Parameters          []ApiClientMethodParameter
	Response            *ApiClientMethodResponse
}
