package utils

import (
	"fmt"
	"os"
)

func CreateDirectory(path string) {
	error := os.MkdirAll(path, os.ModePerm)
	if error != nil {
		fmt.Println("Cannot create directory: " + error.Error())

		os.Exit(1)
	}
}
