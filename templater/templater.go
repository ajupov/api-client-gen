package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"

	converter "github.com/ajupov/api-client-gen/converter/types"
	templater "github.com/ajupov/api-client-gen/templater/types"
	"github.com/ajupov/api-client-gen/utils"
)

func Template(language string, api *converter.Api) *[]templater.Directory {
	var config templater.Config

	languageDirectoryPath := "./templater/languages/" + language

	bytes := utils.ReadFromFile(languageDirectoryPath + "/config.json")
	error := json.Unmarshal(*bytes, &config)
	if error != nil {
		fmt.Println("Cannot parse content: " + error.Error())
		os.Exit(1)
	}

	return &[]templater.Directory{
		{
			Name:  "",
			Files: *copyWithoutTemplating(languageDirectoryPath, config.CopyWithoutTemplating),
		},
		{
			Name:  config.ApiClientDirectory,
			Files: *templateApiClients(config, languageDirectoryPath+"/"+config.ApiClientTemplate, &api.ApiClients),
		},
		{
			Name:  config.ApiModelDirectory,
			Files: *templateApiModels(config, languageDirectoryPath+"/"+config.ApiModelTemplate, &api.ApiModels),
		},
	}
}

func copyWithoutTemplating(rootPath string, paths *[]string) *[]templater.File {
	files := make([]templater.File, len(*paths))

	for index, path := range *paths {
		files[index] = templater.File{
			Name:    path,
			Content: utils.ReadFromFile(rootPath + "/" + path),
		}
	}

	return &files
}

func templateApiClients(config templater.Config, templatePath string, clients *[]converter.ApiClient) *[]templater.File {
	compiledRegex := regexp.MustCompile(`\{(\w+)\}`)

	funcMap := template.FuncMap{
		"ToLower": strings.ToLower,
		"FilterIsInQueryParameters": func(parameters []converter.ApiClientMethodParameterOrBody) []converter.ApiClientMethodParameterOrBody {
			result := make([]converter.ApiClientMethodParameterOrBody, 0)

			for _, parameter := range parameters {
				if parameter.IsInQuery {
					result = append(result, parameter)
				}
			}

			return result
		},
		"GetMappedType": func(oldType string) string {
			return getMappedType(config.TypeMappings, oldType)
		},
		"PathParameterReplace": func(value string) string {
			if config.PathParameterReplacePattern == nil || len(*config.PathParameterReplacePattern) == 0 {
				return value
			}

			replaceValue := strings.Replace(*config.PathParameterReplacePattern, "{parameter}", "{$1}", 1)

			return compiledRegex.ReplaceAllString(value, replaceValue)
		},
	}

	templateFile := utils.ReadFromFile(templatePath)

	apiClientTemplater, error := template.New("ApiClientTemplater").Funcs(funcMap).Parse(string(*templateFile))
	if error != nil {
		fmt.Println(error.Error())
		os.Exit(1)
	}

	files := make([]templater.File, len(*clients))

	for i, apiClient := range *clients {
		var buffer bytes.Buffer

		apiClientTemplater.Execute(&buffer, apiClient)
		bytes := buffer.Bytes()

		files[i] = templater.File{
			Name:    apiClient.Name + "Client." + config.ApiClientFileExtension,
			Content: &bytes,
		}
	}

	return &files
}

func templateApiModels(config templater.Config, templatePath string, models *[]converter.ApiModel) *[]templater.File {
	templateFile := utils.ReadFromFile(templatePath)

	funcMap := template.FuncMap{
		"GetMappedType": func(oldType string) string {
			return getMappedType(config.TypeMappings, oldType)
		},
	}

	apiModelTemplater, error := template.New("ApiModelTemplater").Funcs(funcMap).Parse(string(*templateFile))
	if error != nil {
		fmt.Println(error.Error())
		os.Exit(1)
	}

	files := make([]templater.File, len(*models))

	for i, apiModel := range *models {
		var buffer bytes.Buffer

		apiModelTemplater.Execute(&buffer, apiModel)
		bytes := buffer.Bytes()

		files[i] = templater.File{
			Name:    apiModel.Name + "." + config.ApiModelFileExtension,
			Content: &bytes,
		}
	}

	return &files
}

func getMappedType(typeMappings *map[string]string, oldType string) string {
	if typeMappings == nil || len(*typeMappings) == 0 {
		return oldType
	}

	newType, isExists := (*typeMappings)[oldType]
	if !isExists || newType == "" {
		return oldType
	}

	return newType
}
