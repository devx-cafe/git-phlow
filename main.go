package main

import (
	"github.com/praqma/git-phlow/plugins"
	"fmt"
)

func main() {
	//cmd.Execute()
	/*issue, err := plugins.GetJiraIssue("http://jira.teamsinspace.com:8080", "TIS-41", "admin", "Charlie!")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(issue.Key, issue.Fields.Summary, issue.Fields.Status.Name)

	tr, err := plugins.GetTranstions("http://jira.teamsinspace.com:8080", "TIS-41", "admin", "Charlie!")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tr.Transitions)

	err = plugins.DoTransition("http://jira.teamsinspace.com:8080", "TIS-41", "admin", "Charlie!",tr.Transitions[0].ID)
	if err != nil {
		fmt.Println(err)
	}*/

	err := plugins.AssignUser("http://jira.teamsinspace.com:8080", "TIS-41", "admin", "Charlie!")
	if err != nil {
		fmt.Println(err)
	}


}
