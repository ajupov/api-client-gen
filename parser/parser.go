package parser

import (
	"encoding/json"
	"fmt"
	"os"

	parser "github.com/ajupov/api-client-gen/parser/types"
)

func Parse(content string) *parser.Swagger {
	var swagger parser.Swagger

	error := json.Unmarshal([]byte(content), &swagger)
	if error != nil {
		fmt.Println("Cannot parse content: " + error.Error())

		os.Exit(1)
	}

	return &swagger
}
