package plugins

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"
)

var (
	labelActionAwating    = "Action - awaiting feed-back"
	labelActionGroomig    = "Action - needs grooming"
	labelPrioOne          = "Prio 1 - must have"
	labelPrioTwo          = "Prio 2 - should have"
	labelPrioThree        = "Prio 3 - could have"
	labelPrioFour         = "Prio 4 - won't have"
	labelSizeSmall        = "Size 1 - small"
	labelSizeMedium       = "Size 2 - medium"
	labelSizeLarge        = "Size 3 - large"
	labelSizeTooBig       = "Size 4 - too big"
	labelStatusDublicate  = "Status - duplicate"
	labelStatusWorkable   = "Status - workable"
	labelStatusInProgress = "Status - in progress"
	labelStatusUpNext     = "Status - up next"
)

//BranchNameFromIssue ...
func BranchNameFromIssue(issue int, name string) string {
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
