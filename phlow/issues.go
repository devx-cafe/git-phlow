package phlow

import (
	"fmt"
	"os"
	"github.com/praqma/git-phlow/plugins"
	"os/exec"
	"strings"
	"github.com/praqma/git-phlow/executor"
	"bytes"
	"github.com/praqma/git-phlow/ui"
	"runtime"
)

func Issues() {
	ui.PhlowSpinner.Start("")
	issues, err := plugins.GitHub.GetIssues()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	ui.PhlowSpinner.Stop()

	//Collection Issues to a string
	var buffer bytes.Buffer
	for _, issue := range issues {
		buffer.WriteString(issue.ToString())
	}

	pager := GetPager()

	if pager == "" {
		fmt.Println(buffer.String())
	} else {
		IssuesInPager("less", buffer.String())
	}
}

//GetPager ...
//return the pager if set
func GetPager() string {
	pager := os.Getenv("PAGER")
	if pager != "" {
		return pager
	}

	if runtime.GOOS == "windows" {
		return "more"
	}
	return ""
}

func IssuesInPager(pager, text string) error {
	cmd := exec.Command(pager)
	cmd.Stdin = strings.NewReader(text)
	cmd.Stdout = os.Stdout
	err := executor.ExecuteCommander(cmd)
	if err != nil {
		return err
	}
	return nil
}

//IssueList ...
//List open issues from GitHub
func IssueList() {

	//ui.PhlowSpinner.Start("")
	//
	//ui.PhlowSpinner.Stop()
	//
	////Nested function for finding user issues
	//var userIssue = func(issue []plugins.AssigneeIssue) bool {
	//	user := githandler.ConfigGet("user", "phlow")
	//	for _, u := range issue {
	//		if u.Login == user {
	//			return true
	//		}
	//	}
	//	return false
	//}
	//
	//fmt.Println(ui.Format.MileStone("# Issue"))
	//
	//for _, issue := range issues {
	//	assignees := issue.Assignees
	//	//If mine is true we print on issues assigned to a user
	//	if options.GlobalFlagMine {
	//		if userIssue(assignees) {
	//			printIssue(issue)
	//		}
	//	} else {
	//		printIssue(issue)
	//	}
	//}

}
