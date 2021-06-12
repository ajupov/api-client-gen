package generator

import (
	"fmt"
	"os"
	"regexp"
	"sort"

	converter "github.com/ajupov/api-client-gen/converter/types"
	helpers "github.com/ajupov/api-client-gen/helpers"
	parser "github.com/ajupov/api-client-gen/parser/types"
)

const applicationJsonContentType = "application/json"

func Convert(swagger *parser.Swagger, regex string) *converter.Api {
	matched, error := regexp.MatchString(`3(.\d+)*`, *swagger.Openapi)
	if error != nil {
		fmt.Println("Cannot match OpenAPI version: " + error.Error())
		os.Exit(1)
	}

	if !matched {
		fmt.Println("Cannot convert OpenAPI with version: " + *swagger.Openapi + ". Version should be `3.x.x`")
		os.Exit(1)
	}

	if swagger.Paths == nil {
		return nil
	}

	apiClients := make([]converter.ApiClient, 0)

	for path, pathItem := range *swagger.Paths {
		convertPath(&apiClients, regex, path, &pathItem)
	}

	apiModels := make([]converter.ApiModel, 0)

	if swagger.Components != nil && *swagger.Components.Schemas != nil {
		for schemaName, schema := range *swagger.Components.Schemas {
			convertSchema(&apiModels, schemaName, &schema)
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

	api := converter.Api{
		ApiClients: apiClients,
		ApiModels:  apiModels,
	}

	return &api
}

func convertPath(apiClients *[]converter.ApiClient, regex string, path string, pathItem *parser.SwaggerPathItem) {
	if pathItem.Get != nil {
		convertHttpMethod(apiClients, regex, path, "GET", pathItem.Get)
	}

	if pathItem.Post != nil {
		convertHttpMethod(apiClients, regex, path, "POST", pathItem.Post)
	}

	if pathItem.Put != nil {
		convertHttpMethod(apiClients, regex, path, "PUT", pathItem.Put)
	}

	if pathItem.Patch != nil {
		convertHttpMethod(apiClients, regex, path, "PATCH", pathItem.Patch)
	}

	if pathItem.Delete != nil {
		convertHttpMethod(apiClients, regex, path, "DELETE", pathItem.Delete)
	}
}

func convertHttpMethod(apiClients *[]converter.ApiClient, regex string, path string, method string, operation *parser.SwaggerOperation) {
	if *operation.Tags == nil || len(*operation.Tags) == 0 {
		return
	}

	apiClientName := (*operation.Tags)[0]
	apiClient := getOrAddApiClient(apiClients, apiClientName)
	apiClientMethod := addApiClientMethod(regex, path, method)

	convertParameters(apiClientMethod, &apiClient.Imports, operation.Parameters)
	convertResponse(apiClientMethod, &apiClient.Imports, operation.Responses)
	convertRequestBody(apiClientMethod, &apiClient.Imports, operation.RequestBody)

	apiClient.Methods = append(apiClient.Methods, *apiClientMethod)
}

func convertParameters(apiClientMethod *converter.ApiClientMethod, apiClientImports *[]string, parameters *[]parser.SwaggerParameterOrReference) {
	if parameters == nil {
		return
	}

	for _, parameter := range *parameters {
		_type, _isImportType := getTypeWithImport(*parameter.Schema)
		if _isImportType && !helpers.StringArrayContains(apiClientImports, _type) {
			*apiClientImports = append(*apiClientImports, _type)
		}

		apiClientParameter := converter.ApiClientMethodParameterOrBody{
			Name:      *parameter.Name,
			Type:      _type,
			IsInPath:  *parameter.In == "path",
			IsInQuery: *parameter.In == "query",
			Nullable:  parameter.Schema.Required != nil && !*parameter.Schema.Required,
		}

		apiClientMethod.AllParameters = append(apiClientMethod.AllParameters, apiClientParameter)

		if *parameter.In == "query" {
			apiClientMethod.QueryParameters = append(apiClientMethod.QueryParameters, apiClientParameter)
		}

		if !apiClientMethod.IsNameMatchedFromRegex {
			apiClientMethod.Name = apiClientMethod.Name + helpers.CapitalizeString(apiClientParameter.Name)
		}
	}
}

func convertResponse(apiClientMethod *converter.ApiClientMethod, apiClientImports *[]string, responses *map[string]parser.SwaggerResponseOrSwaggerReference) {
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

	apiClientMethod.Response = converter.ApiClientMethodResponse{
		IsArrayOfType:      false,
		IsDictionaryOfType: false,
	}

	if applicationJson.Schema.Ref != nil {
		_type, _isImportType := getTypeWithImport(*applicationJson.Schema)
		if _isImportType && !helpers.StringArrayContains(apiClientImports, _type) {
			*apiClientImports = append(*apiClientImports, _type)
		}

		apiClientMethod.Response.Type = _type
		apiClientMethod.Response.Nullable = applicationJson.Schema.Nullable != nil && *applicationJson.Schema.Nullable
	} else if *applicationJson.Schema.Type == "object" && applicationJson.Schema.AdditionalProperties != nil {
		apiClientMethod.Response.IsDictionaryOfType = true

		_type, _isImportType := getTypeWithImport(*applicationJson.Schema.AdditionalProperties)
		if _isImportType && !helpers.StringArrayContains(apiClientImports, _type) {
			*apiClientImports = append(*apiClientImports, _type)
		}

		apiClientMethod.Response.Type = _type
		apiClientMethod.Response.Nullable = applicationJson.Schema.AdditionalProperties.Nullable != nil && *applicationJson.Schema.AdditionalProperties.Nullable
	} else if *applicationJson.Schema.Type == "array" && applicationJson.Schema.Items != nil {
		apiClientMethod.Response.IsArrayOfType = true

		_type, _isImportType := getTypeWithImport(*applicationJson.Schema.Items)
		if _isImportType && !helpers.StringArrayContains(apiClientImports, _type) {
			*apiClientImports = append(*apiClientImports, _type)
		}

		apiClientMethod.Response.Type = _type
		apiClientMethod.Response.Nullable = applicationJson.Schema.Items.Nullable != nil && *applicationJson.Schema.Items.Nullable
	} else if applicationJson.Schema.Type != nil {
		apiClientMethod.Response.Type = *applicationJson.Schema.Type
		apiClientMethod.Response.Nullable = applicationJson.Schema.Nullable != nil && *applicationJson.Schema.Nullable
	}
}

func convertRequestBody(apiClientMethod *converter.ApiClientMethod, apiClientImports *[]string, requestBody *parser.SwaggerRequestBodyOrReference) {
	if requestBody == nil || requestBody.Content == nil {
		return
	}

	applicationJson, isExistsContent := (*requestBody.Content)[applicationJsonContentType]
	if !isExistsContent {
		return
	}

	apiClientMethodParameterOrBody := converter.ApiClientMethodParameterOrBody{
		Name:     "body",
		Nullable: applicationJson.Schema.Required != nil && !*applicationJson.Schema.Required,
	}

	if applicationJson.Schema.Ref != nil {
		_type, _isImportType := getTypeWithImport(*applicationJson.Schema)
		if _isImportType && !helpers.StringArrayContains(apiClientImports, _type) {
			*apiClientImports = append(*apiClientImports, _type)
		}

		apiClientMethodParameterOrBody.Type = _type
		apiClientMethodParameterOrBody.Nullable = applicationJson.Schema.Nullable != nil && *applicationJson.Schema.Nullable
	} else if *applicationJson.Schema.Type == "object" && applicationJson.Schema.AdditionalProperties != nil {
		apiClientMethodParameterOrBody.IsDictionaryOfType = true

		_type, _isImportType := getTypeWithImport(*applicationJson.Schema.AdditionalProperties)
		if _isImportType && !helpers.StringArrayContains(apiClientImports, _type) {
			*apiClientImports = append(*apiClientImports, _type)
		}

		apiClientMethodParameterOrBody.Type = _type
		apiClientMethodParameterOrBody.Nullable = applicationJson.Schema.AdditionalProperties.Nullable != nil && *applicationJson.Schema.AdditionalProperties.Nullable
	} else if *applicationJson.Schema.Type == "array" && applicationJson.Schema.Items != nil {
		apiClientMethodParameterOrBody.IsArrayOfType = true

		_type, _isImportType := getTypeWithImport(*applicationJson.Schema.Items)
		if _isImportType && !helpers.StringArrayContains(apiClientImports, _type) {
			*apiClientImports = append(*apiClientImports, _type)
		}

		apiClientMethodParameterOrBody.Type = _type
		apiClientMethodParameterOrBody.Nullable = applicationJson.Schema.Items.Nullable != nil && *applicationJson.Schema.Items.Nullable
	} else if applicationJson.Schema.Type != nil {
		apiClientMethodParameterOrBody.Type = *applicationJson.Schema.Type
		apiClientMethodParameterOrBody.Nullable = applicationJson.Schema.Nullable != nil && *applicationJson.Schema.Nullable
	}

	apiClientMethod.RequestBody = &apiClientMethodParameterOrBody
	apiClientMethod.AllParameters = append(apiClientMethod.AllParameters, apiClientMethodParameterOrBody)

	if !apiClientMethod.IsNameMatchedFromRegex {
		apiClientMethod.Name = apiClientMethod.Name + helpers.CapitalizeString(apiClientMethodParameterOrBody.Name)
	}
}

func convertSchema(apiModels *[]converter.ApiModel, schemaName string, schema *parser.SwaggerComponentsSchemaOrSwaggerReference) {
	apiModel := converter.ApiModel{
		Name: schemaName,
	}

	if schema.Enum != nil {
		apiModel.IsEnum = true
		apiModel.EnumItems = make([]converter.ApiModelEnumItem, len(*schema.Enum))

		for i := 0; i < len(*schema.Enum); i++ {
			apiModelEnumItem := converter.ApiModelEnumItem{
				Name:  (*schema.XEnumNames)[i],
				Value: (*schema.Enum)[i],
			}

			apiModel.EnumItems[i] = apiModelEnumItem
		}
	} else if schema.Properties != nil {
		for propertyName, property := range *schema.Properties {
			apiModelProperty := converter.ApiModelProperty{
				Name: propertyName,
			}

			if property.Ref != nil {
				_type, _isImportType := getTypeWithImport(property)
				if _isImportType && !helpers.StringArrayContains(&apiModel.Imports, _type) {
					apiModel.Imports = append(apiModel.Imports, _type)
				}

				apiModelProperty.Type = _type
				apiModelProperty.Nullable = property.Nullable != nil && *property.Nullable
			} else if *property.Type == "object" && property.AdditionalProperties != nil {
				apiModelProperty.IsDictionaryOfType = true

				_type, _isImportType := getTypeWithImport(*property.AdditionalProperties)
				if _isImportType && !helpers.StringArrayContains(&apiModel.Imports, _type) {
					apiModel.Imports = append(apiModel.Imports, _type)
				}

				apiModelProperty.Type = _type
				apiModelProperty.Nullable = property.AdditionalProperties.Nullable != nil && *property.AdditionalProperties.Nullable
			} else if *property.Type == "array" && property.Items != nil {
				apiModelProperty.IsArrayOfType = true

				_type, _isImportType := getTypeWithImport(*property.Items)
				if _isImportType && !helpers.StringArrayContains(&apiModel.Imports, _type) {
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

func getOrAddApiClient(apiClients *[]converter.ApiClient, apiClientName string) *converter.ApiClient {
	for i := 0; i < len(*apiClients); i++ {
		if (*apiClients)[i].Name == apiClientName {
			return &(*apiClients)[i]
		}
	}

	apiClient := &converter.ApiClient{
		Name:    apiClientName,
		Imports: make([]string, 0),
		Methods: make([]converter.ApiClientMethod, 0),
	}

	*apiClients = append(*apiClients, *apiClient)

	for i := 0; i < len(*apiClients); i++ {
		if (*apiClients)[i].Name == apiClientName {
			return &(*apiClients)[i]
		}
	}

	return nil
}

func addApiClientMethod(regex string, path string, method string) *converter.ApiClientMethod {
	requestContentType := ""
	if method != "GET" {
		requestContentType = applicationJsonContentType
	}

	name, isMatched := helpers.GetActionName(regex, path, method)
	apiClientMethod := &converter.ApiClientMethod{
		Name:                   name,
		IsNameMatchedFromRegex: isMatched,
		Url:                    path,
		Method:                 method,
		RequestContentType:     requestContentType,
		ResponseContentType:    applicationJsonContentType,
	}

	return apiClientMethod
}

func getTypeWithImport(schema parser.SwaggerSchemaOrSwaggerReference) (_type string, _isImportType bool) {
	_isImportType = false

	if schema.Ref != nil {
		_isImportType = true
		_type = helpers.GetPathLastPart(*schema.Ref)
	} else if schema.Type != nil {
		_type = *schema.Type
	}

	return
}
