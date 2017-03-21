package plugins

import (
	"bytes"
	"github.com/fatih/color"
	"regexp"
	"strconv"
	"strings"
)

//PhlowLabels
//Map of labels in the phlow
var PhlowLabels map[string]*PhlowLabel

//PhlowLabel
//label struct for storing GitHub labels
type PhlowLabel struct {
	Title string
	Group int
	ID    int
}

func init() {
	PhlowLabels = make(map[string]*PhlowLabel)
	PhlowLabels["Action - awaiting feed-back"] = &PhlowLabel{"Action - awaiting feed-back", 1, 1}
	PhlowLabels["Action - needs grooming"] = &PhlowLabel{"Action - needs grooming", 1, 2}
	PhlowLabels["Prio 1 - must have"] = &PhlowLabel{"Prio 1 - must have", 2, 3}
	PhlowLabels["Prio 2 - should have"] = &PhlowLabel{"Prio 2 - should have", 2, 4}
	PhlowLabels["Prio 3 - could have"] = &PhlowLabel{"Prio 3 - could have", 2, 5}
	PhlowLabels["Prio 4 - won't have"] = &PhlowLabel{"Prio 4 - won't have", 2, 6}
	PhlowLabels["Size 0 - briefing"] = &PhlowLabel{"Size 0 - briefing", 3, 7}
	PhlowLabels["Size 1 - small"] = &PhlowLabel{"Size 1 - small", 3, 8}
	PhlowLabels["Size 2 - medium"] = &PhlowLabel{"Size 2 - medium", 3, 9}
	PhlowLabels["Size 3 - large"] = &PhlowLabel{"Size 3 - large", 3, 10}
	PhlowLabels["Size 4 - too big"] = &PhlowLabel{"Size 4 - too big", 3, 11}
	PhlowLabels["Status - duplicate"] = &PhlowLabel{"Status - duplicate", 4, 12}
	PhlowLabels["Status - workable"] = &PhlowLabel{"Status - workable", 4, 13}
	PhlowLabels["Status - in progress"] = &PhlowLabel{"Status - in progress", 4, 14}
	PhlowLabels["Status - up next"] = &PhlowLabel{"Status - up next", 4, 15}
}

//Colorizer
//Determines color based on label group
func Colorizer(label string) string {
	if l := PhlowLabels[label]; l != nil {
		switch l.Group {
		case 1:
			return color.New(color.FgHiGreen).Add(color.Bold).Add(color.BgGreen).SprintFunc()(label)
		case 2:
			return color.New(color.FgWhite).Add(color.Bold).Add(color.BgHiRed).SprintFunc()(label)
		case 3:
			return color.New(color.FgHiBlack).Add(color.Bold).Add(color.BgHiBlue).SprintFunc()(label)
		case 4:
			return color.New(color.FgBlack).Add(color.Bold).Add(color.BgHiWhite).SprintFunc()(label)
		default:
			return color.New(color.FgBlack).Add(color.Bold).Add(color.BgWhite).SprintFunc()(label)
		}
	}
	return color.New(color.FgBlack).Add(color.Bold).Add(color.BgWhite).SprintFunc()(label)
}

func MilestoneColor(milestone string) string {
	return color.New(color.FgGreen).SprintFunc()(milestone)
}

//BranchNameFromIssue
//Converts issues to branch names nby removing illegal characters and inserting hyphens
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


//IssueFromBranchName
//Extracts the issue number from the branch name
func IssueFromBranchName(branch string) int {
	iss, err := strconv.Atoi(strings.Split(branch, "-")[0])
	if err != nil {
		return -1
	}
	return iss
}


//efficientConcatString
//Concatenate strings in an effective way
func efficientConcatString(args ...string) string {
	buffer := bytes.Buffer{}
	for _, str := range args {
		buffer.WriteString(str)
	}
	return buffer.String()
}
