package phlow

import (
	"regexp"
	"strconv"
	"strings"
	"bytes"
)

//SanitizeIssueToBranchName ...
//Removes illigal characters from your issues
//Git branches can only contain certian characters
func SanitizeIssueToBranchName(issue int, name string) string {
	var result string

	result = strings.ToLower(name)
	result = strings.Replace(result, " ", "-", -1)

	if strings.HasPrefix(result, ".") {
		result = result[1:]
	}

	if strings.HasSuffix(result, "/") {
		result = result[:len(result)-1]

	}

	result = strings.Replace(result, ".", "-", -1)
	result = strings.Replace(result, "\"", "", -1)
	r := regexp.MustCompile("[\040\177 ~^:?*[]+")
	result = r.ReplaceAllString(result, "")
	result = strconv.Itoa(issue) + "-" + result
	return result
}

func efficientConcatString(args ...string) string {
	buffer := bytes.Buffer{}
	for _, str := range args {
		buffer.WriteString(str)
	}
	return buffer.String()
}
