package phlow

import (
	"fmt"
	"strings"

	"os"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/executor"
	"strconv"
	"errors"
)

//WrapUp ...
func WrapUp() {
	git := githandler.Git{Run: executor.RunGit}

	//Add all files to index
	_, err := git.Add("--all")
	if err != nil {
		fmt.Println("Project files could not be added: " + err.Error())
		return
	}

	out, err := git.Branch("-a")
	if err != nil {
		fmt.Println(err)
		return
	}

	//Retrieve branch info - current branch
	info := githandler.AsList(out)
	var commitMessage string

	issue, err := GetJIRAIssue(info.Current)
	if err != nil {
		
		if options.GlobalFlagForceMessage != "" {
			commitMessage = "close #" + strings.Split(info.Current, "-")[0] + " " + options.GlobalFlagForceMessage
		} else {
			commitMessage = "close #" + strings.Replace(info.Current, "-", " ", -1)
		}
	} else {
		msg := strings.TrimPrefix(info.Current, issue)

		if options.GlobalFlagForceMessage != "" {
			commitMessage = "close #" + issue + " " + options.GlobalFlagForceMessage
		} else {
			commitMessage = "close #" + issue + strings.Replace(msg, "-", " ", -1)
		}

	}

	_, err = git.Commit("-m", commitMessage)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintln(os.Stdout, commitMessage)
}

func GetJIRAIssue(branch string) (string, error) {

	parts := strings.Split(branch, "-")
	if _, err := strconv.Atoi(parts[0]); err != nil && len(parts) > 1 {
		return parts[0] + "-" + parts[1], nil
	}
	return "", errors.New("not a jira branch")
}
