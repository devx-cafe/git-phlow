package plugins

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/go-errors/errors"
)

//Authorization ...
//interface for authorization with external service
type Authorization func(URL, user, pass string) (token string, err error)

//Authentication ...
//interface for authentication with external service
type Authentication func(URL, user, pass string) (authenticated bool, err error)

//DefaultBranch ...
//interface for getting the default branch of the external service
type DefaultBranch func(URL, org, repo, token string) (defaultBranch string, err error)

//Stringer ...
//interface for github formats
type Stringer interface {
	ToString() string
}

//PhlowLabels ...
//Map of labels in the phlow
var PhlowLabels map[string]*PhlowLabel

//PhlowLabel ...
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

//GroupID ...
func GroupID(name string) int {
	if PhlowLabels[name] != nil {
		return PhlowLabels[name].Group
	}
	return 4
}

//BranchNameFromIssue ...
//Converts issues to branch names nby removing illegal characters and inserting hyphens
func BranchNameFromIssue(issue string, name string) string {
	var result string

	removeNonAlphaNumeric := regexp.MustCompile("([^\\w-])+")
	result = removeNonAlphaNumeric.ReplaceAllString(name, "-")

	removeDuplicateChars := regexp.MustCompile("([-_]{2,})")
	result = removeDuplicateChars.ReplaceAllString(result, "")

	result = strings.Trim(result, "-")
	result = strings.ToLower(result)
	return issue + "-" + result
}

//IssueExtractor ...
type IssueExtractor func(branch string) (string, error)

//IssueFromBranchName ...
//Extracts github issue from the branch name. Will error if no issue ID is identified
func IssueFromBranchName(branch string) (string, error) {
	arr := strings.Split(branch, "-")
	_, err := strconv.Atoi(arr[0])
	if err != nil {
		return "", errors.New("Could not extract github issue ID from branch")
	}
	return arr[0], nil
}

//KeyFromBranchName ...
//Extracts a Jira key from a branch name. Will error if no key is identified
func KeyFromBranchName(branch string) (string, error) {
	parts := strings.Split(branch, "-")
	//Jira issues must have 3 parts PRJ-123-NAME
	if len(parts) < 3 {
		return "", errors.New("Could not get Jira Key from branch name")
	}

	if _, err := strconv.Atoi(parts[0]); err != nil && len(parts) > 1 {
		return parts[0] + "-" + parts[1], nil
	}
	return "", errors.New("could not get Jira key from branch name")

}
