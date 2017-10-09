package phlow

import (
	"fmt"
	"strings"

	"github.com/praqma/git-phlow/plugins"

	"os"

	"errors"
	"strconv"

	"github.com/praqma/git-phlow/executor"
	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/setting"
	"github.com/praqma/git-phlow/options"
)

//WrapUpCaller ...
func WrapUpCaller() {
	conf := setting.NewProjectStg(options.GlobalFlagTarget)

	if conf.Service == "jira" {
		WrapUp("#close ", options.GlobalFlagForceMessage, plugins.KeyFromBranchName)
		return
	}

	if conf.Service == "github" {
		WrapUp("close #", options.GlobalFlagForceMessage, plugins.IssueFromBranchName)
		return
	}

}

//GenerateMessage ...
//generates a message from a branch with a given issue extractor
func GenerateMessage(branch string, smartCommitPrefix string, extractor plugins.IssueExtractor, forceMessage string) (string, error) {
	iss, err := extractor(branch)
	if err != nil {
		return "", err
	}

	//remove the issue from the branch name
	branch = strings.TrimPrefix(branch, iss)

	if forceMessage != "" {
		return smartCommitPrefix + iss + " " + forceMessage, nil
	}

	//Replace - with white space
	msg := strings.Replace(branch, "-", " ", -1)

	//prepend smart commit prefix
	return smartCommitPrefix + iss + msg, nil
}

//WrapUp ...
func WrapUp(smartCommitPrefix, force string, extractor plugins.IssueExtractor) {
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

	msg, err := GenerateMessage(info.Current, smartCommitPrefix, extractor, force)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = git.Commit("-m", msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(msg)
}

func GetJIRAIssue(branch string) (string, error) {

	parts := strings.Split(branch, "-")
	if _, err := strconv.Atoi(parts[0]); err != nil && len(parts) > 1 {
		return parts[0] + "-" + parts[1], nil
	}
	return "", errors.New("not a jira branch")
}
