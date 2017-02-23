package plugins

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"
)

var (
	//LabelActionAwating ...
	LabelActionAwating = "Action - awaiting feed-back"
	//LabelActionGroomig ...
	LabelActionGroomig = "Action - needs grooming"
	//LabelPrioOne ...
	LabelPrioOne = "Prio 1 - must have"
	//LabelPrioTwo ...
	LabelPrioTwo = "Prio 2 - should have"
	//LabelPrioThree ...
	LabelPrioThree = "Prio 3 - could have"
	//LabelPrioFour ...
	LabelPrioFour = "Prio 4 - won't have"
	//LabelSizeSmall ...
	LabelSizeSmall = "Size 1 - small"
	//LabelSizeMedium ...
	LabelSizeMedium = "Size 2 - medium"
	//LabelSizeLarge ...
	LabelSizeLarge = "Size 3 - large"
	//LabelSizeTooBig ...
	LabelSizeTooBig = "Size 4 - too big"
	//LabelStatusDublicate ...
	LabelStatusDublicate = "Status - duplicate"
	//LabelStatusWorkable ...
	LabelStatusWorkable = "Status - workable"
	//LabelStatusInProgress ...
	LabelStatusInProgress = "Status - in progress"
	//LabelStatusUpNext ...
	LabelStatusUpNext = "Status - up next"
)

//BranchNameFromIssue ...
func BranchNameFromIssue(issue int, name string) string {
	var result string
	
	removeNonAlphaNumeric := regexp.MustCompile("([^\\w-])+")
	result = removeNonAlphaNumeric.ReplaceAllString(name, "-")

	removeDuplicateChars := regexp.MustCompile("([-_]{2,})")
	result = removeDuplicateChars.ReplaceAllString(result, "")

	result = strings.Trim(result, "-")
	result = strings.ToLower(result)
	return strconv.Itoa(issue) + "-" + result
}

func efficientConcatString(args ...string) string {
	buffer := bytes.Buffer{}
	for _, str := range args {
		buffer.WriteString(str)
	}
	return buffer.String()
}
