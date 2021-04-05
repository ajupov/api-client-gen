package main

import (
	"flag"
	"fmt"

	filesystem "github.com/ajupov/api-client-gen/filesystem"
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

	filesystem.CreateDirectory(*outputDirectory)
	content := filesystem.ReadFromFile(*inputFile)
	swagger := parser.Parse(content)
	serialized := parser.Serialize(swagger)

	filesystem.WriteToFile(*outputDirectory+"/"+"file1.json", serialized)
}
