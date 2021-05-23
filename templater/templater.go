package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
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
			Name:  config.ApiClientDirectory,
			Files: *templateApiClients(languageDirectoryPath+"/"+config.ApiClientTemplate, &api.ApiClients),
		},
		// {
		// 	Name:  config.ApiModelDirectory,
		// 	Files: *templateApiModels(languageDirectoryPath+"/"+config.ApiModelTemplate, &api.ApiModels),
		// },
	}
}

func templateApiClients(templatePath string, clients *[]converter.ApiClient) *[]templater.File {
	apiClientTemplater, error := template.New("ApiClientTemlater").ParseFiles(templatePath)
	if error != nil {
		fmt.Println("Cannot parse content: " + error.Error())
		os.Exit(1)
	}

	files := make([]templater.File, len(*clients))

	for i, apiClient := range *clients {
		var buffer bytes.Buffer

		apiClientTemplater.Execute(&buffer, apiClient)

		bytes := buffer.Bytes()

		files[i] = templater.File{
			Name:    apiClient.Name + ".json",
			Content: &bytes,
		}
	}

	return &files
}

func templateApiModels(templatePath string, models *[]converter.ApiModel) *[]templater.File {
	files := make([]templater.File, len(*models))

	for i, apiModel := range *models {
		content, error := json.MarshalIndent(apiModel, "", "  ")
		if error != nil {
			fmt.Print(error.Error())
		}

		files[i] = templater.File{
			Name:    apiModel.Name + ".json",
			Content: &content,
		}
	}

	return &files
}
