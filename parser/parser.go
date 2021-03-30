package parser

import (
	"encoding/json"
	"fmt"
	"os"

	parser "github.com/ajupov/api-client-gen/parser/types"
)

func Parse(content *[]byte) *parser.Swagger {
	var swagger parser.Swagger

	error := json.Unmarshal(*content, &swagger)
	if error != nil {
		fmt.Println("Cannot parse content: " + error.Error())

		os.Exit(1)
	}

	return &swagger
}

func Serialize(swagger *parser.Swagger) *[]byte {
	serialized, err := json.MarshalIndent(swagger, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}

	return &serialized
}
