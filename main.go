package main

import (
	"encoding/json"
	"flag"
	"fmt"
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

	fmt.Println(len(content))

	var swagger Swagger

	json.Unmarshal([]byte(content), &swagger)

	CreateDirectory(*outputDirectory)

	WriteFile(*outputDirectory+"/"+"file", swagger.Info)
}
