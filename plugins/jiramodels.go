package plugins

import (
	"bytes"

	"github.com/praqma/git-phlow/ui"
)

//JiraIssue ...
type JiraIssue struct {
	Key    string `json:"key"`
	Fields struct {
		Summary string `json:"summary"`
		Status  struct {
			Name string `json:"name"`
		} `json:"status"`
	} `json:"fields"`
}

//Transitions ...
type Transitions struct {
	Transitions []Transition `json:"transitions"`
}

//QueryResult ...
type QueryResult struct {
	Issues []JiraIssue `json:"issues"`
}

//Transition ...
type Transition struct {
	ID string `json:"id"`
	To struct {
		StatusCategory struct {
			Name string `json:"name"`
		} `json:"statusCategory"`
	} `json:"to"`
}

//AssignBody ...
type AssignBody struct {
	Name string `json:"name"`
}

//TransitionBody ...
type TransitionBody struct {
	Transition struct {
		ID string `json:"id"`
	} `json:"transition"`
}

//ToString ...
//Formats issue
func (issue JiraIssue) ToString() string {
	var buffer bytes.Buffer
	buffer.WriteString(ui.Format.Bold(issue.Key + ": "))
	buffer.WriteString(issue.Fields.Summary)
	return buffer.String()
}
