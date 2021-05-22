package main

import (
	"encoding/json"
	"flag"
	"fmt"

	converter "github.com/ajupov/api-client-gen/converter"
	parser "github.com/ajupov/api-client-gen/parser"
	utils "github.com/ajupov/api-client-gen/utils"
)

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

	utils.CreateDirectory(*outputDirectory)
	content := utils.ReadFromFile(*inputFile)
	swagger := parser.Parse(content)

	api := converter.Convert(swagger, regex)

	apiClientsSerialized, error := json.MarshalIndent(api.ApiClients, "", "  ")
	if error != nil {
		fmt.Println(error.Error())
	}

	apiClientsSerializedOutputPath := *outputDirectory + "/" + "apiClientsSerialized.json"
	utils.WriteToFile(apiClientsSerializedOutputPath, &apiClientsSerialized)

	// serialized := parser.Serialize(swagger)
	// outputPath := *outputDirectory + "/" + "swagger.json"
	// filesystem.WriteToFile(outputPath, serialized)
}
