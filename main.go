package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"regexp"
	"sort"
	"strings"

	filesystem "github.com/ajupov/api-client-gen/filesystem"
	parser "github.com/ajupov/api-client-gen/parser"
	types "github.com/ajupov/api-client-gen/parser/types"
)

const applicationJsonContentType = "application/json"

type Api struct {
	apiClients []ApiClient
	apiModels  []ApiModel
}

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
	Response            *ApiClientMethodResponse
}

type ApiClientMethodParameter struct {
	Name               string
	Type               string
	IsArrayOfType      bool
	IsDictionaryOfType bool
	Nullable           bool
}

type ApiClientMethodResponse struct {
	Type               string
	IsArrayOfType      bool
	IsDictionaryOfType bool
	Nullable           bool
}

type ApiModel struct {
	Name       string
	Imports    []string
	IsEnum     bool
	Properties []ApiModelProperty
	EnumItems  []ApiModelEnumItem
}

type ApiModelProperty struct {
	Name               string
	Type               string
	IsArrayOfType      bool
	IsDictionaryOfType bool
	Nullable           bool
}

type ApiModelEnumItem struct {
	Name  string
	Value int
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

	api := Convert(swagger, regex)

	apiClientsSerialized, error := json.MarshalIndent(api.apiModels, "", "  ")
	if error != nil {
		fmt.Println(error.Error())
	}

	apiClientsSerializedOutputPath := *outputDirectory + "/" + "apiClientsSerialized.json"
	filesystem.WriteToFile(apiClientsSerializedOutputPath, &apiClientsSerialized)

	// serialized := parser.Serialize(swagger)
	// outputPath := *outputDirectory + "/" + "swagger.json"
	// filesystem.WriteToFile(outputPath, serialized)
}

func Convert(swagger *types.Swagger, regex *string) *Api {
	matched, error := regexp.MatchString("3(.\\d+)*", *swagger.Openapi)
	if error != nil {
		fmt.Println(error.Error())
	}

	if !matched {
		return nil
	}

	if swagger.Paths == nil {
		return nil
	}

	apiClients := make([]ApiClient, 0)

	for path, pathItem := range *swagger.Paths {
		ConvertPath(&apiClients, regex, path, &pathItem)
	}

	apiModels := make([]ApiModel, 0)

	if swagger.Components != nil && *swagger.Components.Schemas != nil {
		for schemaName, schema := range *swagger.Components.Schemas {
			ConvertSchema(&apiModels, schemaName, &schema)
		}
	}

	sort.Slice(apiClients, func(i, j int) bool {
		return apiClients[i].Name < apiClients[j].Name
	})

	for _, item := range apiClients {
		sort.Slice(item.Methods, func(i, j int) bool {
			return item.Methods[i].Name < item.Methods[j].Name
		})
	}

	return &Api{
		apiClients: apiClients,
		apiModels:  apiModels,
	}
}

func ConvertPath(apiClients *[]ApiClient, regex *string, path string, pathItem *types.SwaggerPathItem) {
	if pathItem.Get != nil {
		ConvertHttpMethod(apiClients, regex, path, "GET", pathItem.Get)
	}

	if pathItem.Post != nil {
		ConvertHttpMethod(apiClients, regex, path, "POST", pathItem.Post)
	}

	if pathItem.Put != nil {
		ConvertHttpMethod(apiClients, regex, path, "PUT", pathItem.Put)
	}

	if pathItem.Patch != nil {
		ConvertHttpMethod(apiClients, regex, path, "PATCH", pathItem.Patch)
	}

	if pathItem.Delete != nil {
		ConvertHttpMethod(apiClients, regex, path, "DELETE", pathItem.Delete)
	}
}

func ConvertHttpMethod(apiClients *[]ApiClient, regex *string, path string, method string, operation *types.SwaggerOperation) {
	if *operation.Tags == nil || len(*operation.Tags) == 0 {
		return
	}

	apiClientName := (*operation.Tags)[0]
	apiClient := GetOrAddApiClient(apiClients, apiClientName)
	apiClientMethod := AddApiClientMethod(apiClient, regex, path, method)

	ConvertParameters(apiClientMethod, &apiClient.Imports, operation.Parameters)
	ConvertResponse(apiClientMethod, &apiClient.Imports, operation.Responses)
	ConvertRequestBody(apiClientMethod, &apiClient.Imports, operation.RequestBody)

	apiClient.Methods = append(apiClient.Methods, *apiClientMethod)
}

func GetApiClientMethodName(regex *string, path string, method string) string {
	passedActonRegex := strings.Replace(*regex, "{action}", "\\w+", 1)
	compiledRegex := regexp.MustCompile(passedActonRegex)

	matched := compiledRegex.FindStringSubmatch(path)

	apiClientMethodName := ""
	if len(matched) == 2 {
		apiClientMethodName = matched[1]
	} else {
		apiClientMethodName = strings.Title(strings.ToLower(method))
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
			Nullable: parameter.Schema.Required != nil && !*parameter.Schema.Required,
		}

		apiClientMethod.Parameters = append(apiClientMethod.Parameters, apiClientParameter)
	}
}

func ConvertResponse(apiClientMethod *ApiClientMethod, apiClientImports *[]string, responses *map[string]types.SwaggerResponseOrSwaggerReference) {
	if responses == nil {
		return
	}

	okResponse, isExistsOkResponse := (*responses)["200"]
	if !isExistsOkResponse || okResponse.Content == nil {
		return
	}

	applicationJson, isExistsContent := (*okResponse.Content)[applicationJsonContentType]
	if !isExistsContent {
		return
	}

	apiClientMethod.Response = &ApiClientMethodResponse{
		IsArrayOfType:      false,
		IsDictionaryOfType: false,
	}

	if applicationJson.Schema.Ref != nil {
		_type, _isImportType := GetTypeWithImport(*applicationJson.Schema)
		if _isImportType && !StringArrayContains(apiClientImports, _type) {
			*apiClientImports = append(*apiClientImports, _type)
		}

		apiClientMethod.Response.Type = _type
		apiClientMethod.Response.Nullable = applicationJson.Schema.Nullable != nil && *applicationJson.Schema.Nullable
	} else if *applicationJson.Schema.Type == "object" && applicationJson.Schema.AdditionalProperties != nil {
		apiClientMethod.Response.IsDictionaryOfType = true

		_type, _isImportType := GetTypeWithImport(*applicationJson.Schema.AdditionalProperties)
		if _isImportType && !StringArrayContains(apiClientImports, _type) {
			*apiClientImports = append(*apiClientImports, _type)
		}

		apiClientMethod.Response.Type = _type
		apiClientMethod.Response.Nullable = applicationJson.Schema.AdditionalProperties.Nullable != nil && *applicationJson.Schema.AdditionalProperties.Nullable
	} else if *applicationJson.Schema.Type == "array" && applicationJson.Schema.Items != nil {
		apiClientMethod.Response.IsArrayOfType = true

		_type, _isImportType := GetTypeWithImport(*applicationJson.Schema.Items)
		if _isImportType && !StringArrayContains(apiClientImports, _type) {
			*apiClientImports = append(*apiClientImports, _type)
		}

		apiClientMethod.Response.Type = _type
		apiClientMethod.Response.Nullable = applicationJson.Schema.Items.Nullable != nil && *applicationJson.Schema.Items.Nullable
	} else if applicationJson.Schema.Type != nil {
		apiClientMethod.Response.Type = *applicationJson.Schema.Type
		apiClientMethod.Response.Nullable = applicationJson.Schema.Nullable != nil && *applicationJson.Schema.Nullable
	}
}

func ConvertRequestBody(apiClientMethod *ApiClientMethod, apiClientImports *[]string, requestBody *types.SwaggerRequestBodyOrReference) {
	if requestBody == nil || requestBody.Content == nil {
		return
	}

	applicationJson, isExistsContent := (*requestBody.Content)[applicationJsonContentType]
	if !isExistsContent {
		return
	}

	apiClientParameter := ApiClientMethodParameter{
		Name:     "request",
		Nullable: applicationJson.Schema.Required != nil && !*applicationJson.Schema.Required,
	}

	if applicationJson.Schema.Ref != nil {
		_type, _isImportType := GetTypeWithImport(*applicationJson.Schema)
		if _isImportType && !StringArrayContains(apiClientImports, _type) {
			*apiClientImports = append(*apiClientImports, _type)
		}

		apiClientParameter.Type = _type
		apiClientParameter.Nullable = applicationJson.Schema.Nullable != nil && *applicationJson.Schema.Nullable
	} else if *applicationJson.Schema.Type == "object" && applicationJson.Schema.AdditionalProperties != nil {
		apiClientParameter.IsDictionaryOfType = true

		_type, _isImportType := GetTypeWithImport(*applicationJson.Schema.AdditionalProperties)
		if _isImportType && !StringArrayContains(apiClientImports, _type) {
			*apiClientImports = append(*apiClientImports, _type)
		}

		apiClientParameter.Type = _type
		apiClientParameter.Nullable = applicationJson.Schema.AdditionalProperties.Nullable != nil && *applicationJson.Schema.AdditionalProperties.Nullable
	} else if *applicationJson.Schema.Type == "array" && applicationJson.Schema.Items != nil {
		apiClientParameter.IsArrayOfType = true

		_type, _isImportType := GetTypeWithImport(*applicationJson.Schema.Items)
		if _isImportType && !StringArrayContains(apiClientImports, _type) {
			*apiClientImports = append(*apiClientImports, _type)
		}

		apiClientParameter.Type = _type
		apiClientParameter.Nullable = applicationJson.Schema.Items.Nullable != nil && *applicationJson.Schema.Items.Nullable
	} else if applicationJson.Schema.Type != nil {
		apiClientParameter.Type = *applicationJson.Schema.Type
		apiClientParameter.Nullable = applicationJson.Schema.Nullable != nil && *applicationJson.Schema.Nullable
	}

	apiClientMethod.Parameters = append(apiClientMethod.Parameters, apiClientParameter)
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

	for i := 0; i < len(*apiClients); i++ {
		if (*apiClients)[i].Name == apiClientName {
			return &(*apiClients)[i]
		}
	}

	return nil
}

func AddApiClientMethod(apiClient *ApiClient, regex *string, path string, method string) *ApiClientMethod {
	requestContentType := ""
	if method != "GET" {
		requestContentType = applicationJsonContentType
	}

	apiClientMethod := &ApiClientMethod{
		Name:                GetApiClientMethodName(regex, path, method),
		Url:                 path,
		Method:              method,
		RequestContentType:  requestContentType,
		ResponseContentType: applicationJsonContentType,
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

////////////////////////

func ConvertSchema(apiModels *[]ApiModel, schemaName string, schema *types.SwaggerComponentsSchemaOrSwaggerReference) {
	apiModel := ApiModel{
		Name: schemaName,
	}

	if schema.Enum != nil {
		apiModel.IsEnum = true
		apiModel.EnumItems = make([]ApiModelEnumItem, len(*schema.Enum))

		for i := 0; i < len(*schema.Enum); i++ {
			apiModelEnumItem := ApiModelEnumItem{
				Name:  (*schema.XEnumNames)[i],
				Value: (*schema.Enum)[i],
			}

			apiModel.EnumItems[i] = apiModelEnumItem
		}
	} else if schema.Properties != nil {
		for propertyName, property := range *schema.Properties {
			apiModelProperty := ApiModelProperty{
				Name: propertyName,
			}

			if property.Ref != nil {
				_type, _isImportType := GetTypeWithImport(property)
				if _isImportType && !StringArrayContains(&apiModel.Imports, _type) {
					apiModel.Imports = append(apiModel.Imports, _type)
				}

				apiModelProperty.Type = _type
				apiModelProperty.Nullable = property.Nullable != nil && *property.Nullable
			} else if *property.Type == "object" && property.AdditionalProperties != nil {
				apiModelProperty.IsDictionaryOfType = true

				_type, _isImportType := GetTypeWithImport(*property.AdditionalProperties)
				if _isImportType && !StringArrayContains(&apiModel.Imports, _type) {
					apiModel.Imports = append(apiModel.Imports, _type)
				}

				apiModelProperty.Type = _type
				apiModelProperty.Nullable = property.AdditionalProperties.Nullable != nil && *property.AdditionalProperties.Nullable
			} else if *property.Type == "array" && property.Items != nil {
				apiModelProperty.IsArrayOfType = true

				_type, _isImportType := GetTypeWithImport(*property.Items)
				if _isImportType && !StringArrayContains(&apiModel.Imports, _type) {
					apiModel.Imports = append(apiModel.Imports, _type)
				}

				apiModelProperty.Type = _type
				apiModelProperty.Nullable = property.Items.Nullable != nil && *property.Items.Nullable
			} else if property.Type != nil {
				apiModelProperty.Type = *property.Type
				apiModelProperty.Nullable = property.Nullable != nil && *property.Nullable
			}

			apiModel.Properties = append(apiModel.Properties, apiModelProperty)

		}
	}

	*apiModels = append(*apiModels, apiModel)
}
