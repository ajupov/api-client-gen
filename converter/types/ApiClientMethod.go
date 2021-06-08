package converter

type ApiClientMethod struct {
	Name                string
	Url                 string
	Method              string
	RequestContentType  string
	ResponseContentType string
	PathParameters      []ApiClientMethodParameter
	QueryParameters     []ApiClientMethodParameter
	RequestBody         *ApiClientMethodRequestBody
	Response            ApiClientMethodResponse
}
