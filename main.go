package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		inputFile    = flag.String("inputFile", "", "Path to swagger.json file")
		outputFolder = flag.String("outputFolder", "", "Path to output files folder")
		language     = flag.String("language", "", "Programming language for which clients will be generated")
	)

	flag.Parse()

	fmt.Println("Input file: " + *inputFile)
	fmt.Println("Output folder: " + *outputFolder)
	fmt.Println("Language: " + *language)
}
