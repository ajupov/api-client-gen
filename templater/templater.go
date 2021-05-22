package generator

import (
	"encoding/json"
	"fmt"

	converter "github.com/ajupov/api-client-gen/converter/types"
	templater "github.com/ajupov/api-client-gen/templater/types"
)

func Template(language string, api *converter.Api) *[]templater.Directory {
	return &[]templater.Directory{
		{
			Name:  "clients",
			Files: *templateApiClients(language, &api.ApiClients),
		},
		{
			Name:  "models",
			Files: *templateApiModels(language, &api.ApiModels),
		},
	}
}

func templateApiClients(language string, clients *[]converter.ApiClient) *[]templater.File {
	files := make([]templater.File, len(*clients))

	for i, apiClient := range *clients {
		content, error := json.MarshalIndent(apiClient, "", "  ")
		if error != nil {
			fmt.Print(error.Error())
		}

		files[i] = templater.File{
			Name:    apiClient.Name + ".json",
			Content: &content,
		}
	}

	return &files
}

func templateApiModels(language string, models *[]converter.ApiModel) *[]templater.File {
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
