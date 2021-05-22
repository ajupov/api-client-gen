package helpers

import (
	"regexp"
	"strings"
)

func GetActionName(regex *string, path string, method string) string {
	passedActonRegex := strings.Replace(*regex, "{action}", "\\w+", 1)
	compiledRegex := regexp.MustCompile(passedActonRegex)

	matched := compiledRegex.FindStringSubmatch(path)

	apiClientMethodName := ""
	if len(matched) == 2 {
		apiClientMethodName = matched[1]
	} else {
		apiClientMethodName = strings.Title(strings.ToLower(method))
	}

	return apiClientMethodName
}

func StringArrayContains(array *[]string, element string) bool {
	for _, foundImport := range *array {
		if foundImport == element {
			return true
		}
	}

	return false
}

func GetPathLastPart(value string) string {
	parts := strings.Split(value, "/")

	return parts[len(parts)-1]
}
