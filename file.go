package main

import (
	"fmt"
	"io"
	"os"
)

func openFile(path string) *os.File {
	file, error := os.Open(path)

	if error != nil {
		fmt.Println("Cannot open file: " + error.Error())

		os.Exit(1)
	}

	return file
}

func readContent(file *os.File) *string {
	content := ""
	buffer := make([]byte, 64)

	for {
		count, error := file.Read(buffer)
		if error == io.EOF {
			break
		}

		content += string(buffer[:count])
	}

	return &content
}

func ReadFile(path string) *string {
	file := openFile(path)

	defer file.Close()

	return readContent(file)
}

func CreateDirectory(path string) {
	error := os.MkdirAll(path, os.ModePerm)
	if error != nil {
		fmt.Println("Cannot create directory: " + error.Error())

		os.Exit(1)
	}
}

func WriteFile(path, content string) {
	file, error := os.Create(path)
	if error != nil {
		fmt.Println("Cannot write file: " + error.Error())

		os.Exit(1)
	}

	defer file.Close()

	file.WriteString(content)
}
