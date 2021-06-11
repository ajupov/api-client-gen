package converter

type ApiClientMethod struct {
	Name                string
	Url                 string
	Method              string
	RequestContentType  string
	ResponseContentType string
	AllParameters       []ApiClientMethodParameterOrBody
	QueryParameters     []ApiClientMethodParameterOrBody
	RequestBody         *ApiClientMethodParameterOrBody
	Response            ApiClientMethodResponse
}
