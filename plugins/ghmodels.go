package plugins

import (
	"bytes"
	"strconv"
	"github.com/praqma/git-phlow/ui"
)

//Auth ...
type Auth struct {
	Token string `json:"token"`
}

//Repo ...
type Repo struct {
	DefaultBranch string `json:"default_branch"`
}

//Issue ...
type Issue struct {
	Assignees []AssigneeIssue `json:"assignees"`
	Title     string          `json:"title"`
	Number    int             `json:"number"`
	Body      string          `json:"body"`
	Labels    []Label         `json:"labels"`
	Milestone Milestone       `json:"milestone"`
}

//Label ...
type Label struct {
	ID    int    `json:"id"`
	URL   string `json:"url"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type Milestone struct {
	Title string `json:"title"`
}

//AssigneeIssue ...
type AssigneeIssue struct {
	Login string `json:"login"`
}

//Assignee ...
type Assignee struct {
	Assignees []string `json:"assignees"`
}

//ToString ...
//Formats issue
func (issue Issue) ToString() string {
	var buffer bytes.Buffer

	buffer.WriteString(ui.Format.Bold(strconv.Itoa(issue.Number) + ": "))
	buffer.WriteString(issue.Title)

	for _, label := range issue.Labels {
		buffer.WriteString(" " + ui.Format.FByG(GroupID(label.Name))(label.Name))
	}

	for _, user := range issue.Assignees {
		buffer.WriteString(" " + ui.Format.Assignee(user.Login))
	}
	buffer.WriteString(" " + ui.Format.MileStone(issue.Milestone.Title))

	return buffer.String()
}
