package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"sort"
	"strings"

	filesystem "github.com/ajupov/api-client-gen/filesystem"
	parser "github.com/ajupov/api-client-gen/parser"
	types "github.com/ajupov/api-client-gen/parser/types"
)

type ApiClient struct {
	Name    string
	Imports []string
	Methods []ApiClientMethod
}

type ApiClientMethod struct {
	Name                string
	Url                 string
	Method              string
	RequestContentType  string
	ResponseContentType string
	Parameters          []ApiClientMethodParameter
	Response            ApiClientMethodResponse
}

type ApiClientMethodParameter struct {
	Name     string
	Type     string
	Required bool
}

type ApiClientMethodResponse struct {
	Type               string
	IsArrayOfType      bool
	IsDictionaryOfType bool
}

func main() {
	var (
		inputFile       = flag.String("inputFile", "", "Path to swagger.json file")
		outputDirectory = flag.String("outputDirectory", "", "Path to output files directory")
		regex           = flag.String("regex", "", "Regex")
		language        = flag.String("language", "", "Programming language for which clients will be generated")
	)

	flag.Parse()

	fmt.Println("Input file: " + *inputFile)
	fmt.Println("Output directory: " + *outputDirectory)
	fmt.Println("Regex: " + *regex)
	fmt.Println("Language: " + *language)

	filesystem.CreateDirectory(*outputDirectory)
	content := filesystem.ReadFromFile(*inputFile)
	swagger := parser.Parse(content)

	apiClients := Convert(swagger)

	apiClientsSerialized, error := json.MarshalIndent(*apiClients, "", "  ")
	if error != nil {
		fmt.Println(error.Error())
	}

	apiClientsSerializedOutputPath := *outputDirectory + "/" + "apiClientsSerialized.json"
	filesystem.WriteToFile(apiClientsSerializedOutputPath, &apiClientsSerialized)

	// serialized := parser.Serialize(swagger)
	// outputPath := *outputDirectory + "/" + "swagger.json"
	// filesystem.WriteToFile(outputPath, serialized)
}

func Convert(swagger *types.Swagger) *[]ApiClient {
	if swagger.Paths == nil {
		return nil
	}

	apiClients := make([]ApiClient, 0)

	for path, pathItem := range *swagger.Paths {
		ConvertPath(&apiClients, path, &pathItem)
	}

	sort.Slice(apiClients, func(i, j int) bool {
		return apiClients[i].Name < apiClients[j].Name
	})

	return &apiClients
}

func ConvertPath(apiClients *[]ApiClient, path string, pathItem *types.SwaggerPathItem) {
	if pathItem.Get != nil {
		ConvertHttpMethod(apiClients, path, "GET", pathItem.Get)
	}

	if pathItem.Post != nil {
		ConvertHttpMethod(apiClients, path, "POST", pathItem.Post)
	}

	if pathItem.Put != nil {
		ConvertHttpMethod(apiClients, path, "PUT", pathItem.Put)
	}

	if pathItem.Patch != nil {
		ConvertHttpMethod(apiClients, path, "PATCH", pathItem.Patch)
	}

	if pathItem.Delete != nil {
		ConvertHttpMethod(apiClients, path, "DELETE", pathItem.Delete)
	}
}

func ConvertHttpMethod(apiClients *[]ApiClient, path string, method string, operation *types.SwaggerOperation) {
	if *operation.Tags == nil || len(*operation.Tags) == 0 {
		return
	}

	apiClient := GetOrAddApiClient(apiClients, (*operation.Tags)[0])
	apiClientMethod := AddApiClientMethod(apiClient, path, method)

	ConvertParameters(apiClientMethod, &apiClient.Imports, operation.Parameters)

	ConvertResponse(apiClientMethod, &apiClient.Imports, operation)

	apiClient.Methods = append(apiClient.Methods, *apiClientMethod)

	sort.Slice(apiClient.Methods, func(i, j int) bool {
		return apiClient.Methods[i].Name < apiClient.Methods[j].Name
	})
}

func GetApiClientMethodName(path string) string {
	apiClientMethodName := GetPathLastPart(path)
	if len(apiClientMethodName) == 0 {
		apiClientMethodName = "Default"
	}

	return apiClientMethodName
}

func ConvertParameters(apiClientMethod *ApiClientMethod, apiClientImports *[]string, parameters *[]types.SwaggerParameterOrReference) {
	if parameters == nil {
		return
	}

	for _, parameter := range *parameters {
		_type, _isImportType := GetTypeWithImport(*parameter.Schema)
		if _isImportType && !StringArrayContains(apiClientImports, _type) {
			*apiClientImports = append(*apiClientImports, _type)
		}

		apiClientParameter := ApiClientMethodParameter{
			Name:     *parameter.Name,
			Type:     _type,
			Required: parameter.Schema.Required == nil || *parameter.Schema.Required,
		}

		apiClientMethod.Parameters = append(apiClientMethod.Parameters, apiClientParameter)
	}
}

func ConvertResponse(apiClientMethod *ApiClientMethod, apiClientImports *[]string, operation *types.SwaggerOperation) {
	if operation.Responses == nil {
		return
	}

	okResponse, isExistsOkResponse := (*operation.Responses)["200"]
	if !isExistsOkResponse || okResponse.Content == nil {
		return
	}

	applicationJson, isExistsContent := (*okResponse.Content)["application/json"]
	if !isExistsContent {
		return
	}

	apiClientMethod.Response = ApiClientMethodResponse{
		IsArrayOfType:      false,
		IsDictionaryOfType: false,
	}

	if applicationJson.Schema.Ref != nil {
		_type, _isImportType := GetTypeWithImport(*applicationJson.Schema)
		if _isImportType && !StringArrayContains(apiClientImports, _type) {
			*apiClientImports = append(*apiClientImports, _type)
		}

		apiClientMethod.Response.Type = _type
	} else if *applicationJson.Schema.Type == "object" && applicationJson.Schema.AdditionalProperties != nil {
		apiClientMethod.Response.IsDictionaryOfType = true

		_type, _isImportType := GetTypeWithImport(*applicationJson.Schema.AdditionalProperties)
		if _isImportType && !StringArrayContains(apiClientImports, _type) {
			*apiClientImports = append(*apiClientImports, _type)
		}

		apiClientMethod.Response.Type = _type
	} else if *applicationJson.Schema.Type == "array" && applicationJson.Schema.Items != nil {
		apiClientMethod.Response.IsArrayOfType = true

		_type, _isImportType := GetTypeWithImport(*applicationJson.Schema.Items)
		if _isImportType && !StringArrayContains(apiClientImports, _type) {
			*apiClientImports = append(*apiClientImports, _type)
		}

		apiClientMethod.Response.Type = _type
	} else if applicationJson.Schema.Type != nil {
		apiClientMethod.Response.Type = *applicationJson.Schema.Type
	}
}

func GetOrAddApiClient(apiClients *[]ApiClient, apiClientName string) *ApiClient {
	for i := 0; i < len(*apiClients); i++ {
		if (*apiClients)[i].Name == apiClientName {
			return &(*apiClients)[i]
		}
	}

	apiClient := &ApiClient{
		Name:    apiClientName,
		Imports: make([]string, 0),
		Methods: make([]ApiClientMethod, 0),
	}

	*apiClients = append(*apiClients, *apiClient)

	return apiClient
}

func AddApiClientMethod(apiClient *ApiClient, path string, method string) *ApiClientMethod {
	apiClientMethod := &ApiClientMethod{
		Name:                GetApiClientMethodName(path),
		Url:                 path,
		Method:              method,
		RequestContentType:  "application/json",
		ResponseContentType: "application/json",
	}

	return apiClientMethod
}

func GetTypeWithImport(schema types.SwaggerSchemaOrSwaggerReference) (_type string, _isImportType bool) {
	_isImportType = false

	if schema.Ref != nil {
		_isImportType = true
		_type = GetPathLastPart(*schema.Ref)
	} else if schema.Type != nil {
		_type = *schema.Type
	}

	return
}

func GetPathLastPart(value string) string {
	parts := strings.Split(value, "/")

	return parts[len(parts)-1]
}

func StringArrayContains(array *[]string, element string) bool {
	for _, foundImport := range *array {
		if foundImport == element {
			return true
		}
	}

	return false
}
