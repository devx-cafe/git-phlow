package phlowimpl

import (
	"strings"
	"regexp"
)

//ConvertToBranchName
func ConvertToBranchName(name string) string {
	var result string

	result = strings.ToLower(name)
	result = strings.Replace(result, " ", "-", -1)

	if strings.HasPrefix(result, ".") {
		result = result[1:]
	}

	if strings.HasSuffix(result, "/"){
		result = result[:len(result)-1]

	}

	result = strings.Replace(result, ".", "-", -1)
	result = strings.Replace(result, "\"", "", -1)
	r := regexp.MustCompile("[\040\177 ~^:?*[]+")
	result = r.ReplaceAllString(result, "")

	return result
}
