package main

import (
	"flag"
	"fmt"

	converter "github.com/ajupov/api-client-gen/converter"
	parser "github.com/ajupov/api-client-gen/parser"
	reader "github.com/ajupov/api-client-gen/reader"
	templater "github.com/ajupov/api-client-gen/templater"
	writer "github.com/ajupov/api-client-gen/writer"
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

	readed := reader.Read(*inputFile)
	parsed := parser.Parse(readed)
	converted := converter.Convert(parsed, *regex)
	temlated := templater.Template(*language, converted)
	writer.Write(*outputDirectory, temlated)
}
