package options

import (
	"github.com/fatih/color"
	"strconv"
)

//Bold ...
func Bold(message string) string {
	return color.New(color.Bold).SprintFunc()(message)
}

//SuccessFormat ...
func SuccessFormat(err string) string {
	return color.New(color.FgHiGreen).SprintFunc()(err)
}

//ErrorFormat ...
func ErrorFormat(err string) string {
	return color.New(color.FgHiRed).SprintFunc()(err)
}

//LabelFormat ...
func LabelFormat(message string) string {
	return color.New(color.FgBlack).Add(color.Bold).Add(color.BgHiWhite).SprintFunc()(message)
}

//BranchFormat ...
func BranchFormat(name string) string {
	return color.New(color.FgHiGreen).Add(color.Bold).SprintFunc()(name)
}

//IssueFormat ...
func IssueFormat(issue int) string {
	return Bold("#" + strconv.Itoa(issue))
}

//AssigneeFormat ...
func AssigneeFormat(assignee string) string {
	return color.New(color.FgYellow).Add(color.Bold).SprintFunc()("@" + assignee)
}

//AliasFormat ...
func AliasFormat(alias string) string {
	return color.New(color.FgHiCyan).Add(color.Bold).SprintFunc()(alias)
}
