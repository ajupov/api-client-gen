package writer

import (
	templater "github.com/ajupov/api-client-gen/templater/types"
	utils "github.com/ajupov/api-client-gen/utils"
)

func Write(outputDirectory string, directories *[]templater.Directory) {
	utils.CreateDirectory(outputDirectory)

	for _, directory := range *directories {
		fullOutputDirectory := outputDirectory + "/" + directory.Name

		utils.CreateDirectory(fullOutputDirectory)

		for _, file := range directory.Files {
			utils.WriteToFile(fullOutputDirectory+"/"+file.Name, file.Content)
		}
	}
}
