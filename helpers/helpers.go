package helpers

import (
	"regexp"
	"strings"
)

func GetActionName(regex string, path string, method string) (string, bool) {
	passedActonRegex := strings.Replace(regex, "{action}", `\w+`, 1)
	compiledRegex := regexp.MustCompile(passedActonRegex)

	matched := compiledRegex.FindStringSubmatch(path)

	if len(matched) == 2 {
		return matched[1], true
	} else {
		return CapitalizeString(method), false
	}
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

func CapitalizeString(value string) string {
	return strings.Title(strings.ToLower(value))
}
