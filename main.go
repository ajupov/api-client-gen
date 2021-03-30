package main

import (
	"encoding/json"
	"flag"
	"fmt"

	parser "github.com/ajupov/api-client-gen/parser"
)

func main() {
	var (
		inputFile       = flag.String("inputFile", "", "Path to swagger.json file")
		outputDirectory = flag.String("outputDirectory", "", "Path to output files directory")
		language        = flag.String("language", "", "Programming language for which clients will be generated")
	)

	flag.Parse()

	fmt.Println("Input file: " + *inputFile)
	fmt.Println("Output directory: " + *outputDirectory)
	fmt.Println("Language: " + *language)

	content := ReadFile(*inputFile)
	swagger := parser.Parse(*content)

	CreateDirectory(*outputDirectory)

	res, err := json.MarshalIndent(swagger, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}

	WriteFile(*outputDirectory+"/"+"file.json", string(res))
}
