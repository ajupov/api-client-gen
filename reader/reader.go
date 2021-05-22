package reader

import (
	utils "github.com/ajupov/api-client-gen/utils"
)

func Read(inputFile string) *[]byte {
	return utils.ReadFromFile(inputFile)
}
