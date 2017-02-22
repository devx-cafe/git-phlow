package plugins

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"
)

var (
	LabelActionAwating    = "Action - awaiting feed-back"
	LabelActionGroomig    = "Action - needs grooming"
	LabelPrioOne          = "Prio 1 - must have"
	LabelPrioTwo          = "Prio 2 - should have"
	LabelPrioThree        = "Prio 3 - could have"
	LabelPrioFour         = "Prio 4 - won't have"
	LabelSizeSmall        = "Size 1 - small"
	LabelSizeMedium       = "Size 2 - medium"
	LabelSizeLarge        = "Size 3 - large"
	LabelSizeTooBig       = "Size 4 - too big"
	LabelStatusDublicate  = "Status - duplicate"
	LabelStatusWorkable   = "Status - workable"
	LabelStatusInProgress = "Status - in progress"
	LabelStatusUpNext     = "Status - up next"
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
