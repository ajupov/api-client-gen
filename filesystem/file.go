package filesystem

import (
	"fmt"
	"io"
	"os"
)

func ReadFromFile(path string) *[]byte {
	file, error := os.Open(path)
	if error != nil {
		fmt.Println("Cannot open file: " + error.Error())

		os.Exit(1)
	}

	defer file.Close()

	content := ""
	buffer := make([]byte, 64)

	for {
		count, error := file.Read(buffer)
		if error != nil {
			if error == io.EOF {
				break
			}

			fmt.Println("Cannot open file: " + error.Error())

			os.Exit(1)
		}

		content += string(buffer[:count])
	}

	result := []byte(content)

	return &result
}

func WriteToFile(path string, content *[]byte) {
	file, error := os.Create(path)
	if error != nil {
		fmt.Println("Cannot write to file: " + error.Error())

		os.Exit(1)
	}

	defer file.Close()

	file.Write(*content)
}
