package converter

type ApiClientMethod struct {
	Name                   string
	IsNameMatchedFromRegex bool
	Url                    string
	Method                 string
	RequestContentType     string
	ResponseContentType    string
	AllParameters          []ApiClientMethodParameterOrBody
	QueryParameters        []ApiClientMethodParameterOrBody
	RequestBody            *ApiClientMethodParameterOrBody
	Response               ApiClientMethodResponse
}
