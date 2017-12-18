package phlow

import (
	"fmt"

	"github.com/praqma/git-phlow/executor"
	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/plugins"
	"github.com/praqma/git-phlow/setting"
)

//IssueCaller ...
//prints issue with given target
func IssueCaller() {
	conf := setting.NewProjectStg(options.GlobalFlagTarget)

	if conf.Service == "github" {
		PrintIssues(conf, FetchGH)
	}

	if conf.Service == "jira" {
		PrintIssues(conf, FetchJ)
	}
}

//Fetch ...
//Type for getting issues
type Fetch func(*setting.ProjectSetting) ([]plugins.Stringer, error)

//FetchGH ...
//Fetch for github
func FetchGH(conf *setting.ProjectSetting) ([]plugins.Stringer, error) {
	git := githandler.Git{Run: executor.RunGit}
	remote, err := git.LSRemote("--get-url", conf.Remote)
	if err != nil {
		return nil, err
	}

	token, err := git.Config("--get", "phlow.token")
	if err != nil {
		return nil, err
	}

	oar := githandler.OrgAndRepo(remote)

	list, err := plugins.GetIssuesGitHub(conf.IssueApi, oar.Organisation, oar.Repository, token)
	if err != nil {
		return nil, err
	}
	return list, nil
}

//FetchJ ...
//fetch for Jira
func FetchJ(conf *setting.ProjectSetting) ([]plugins.Stringer, error) {
	git := githandler.Git{Run: executor.RunGit}
	user, err := git.Config("--get", "phlow.jirauser")
	if err != nil {
		return nil, err
	}

	token, err := git.Config("--get", "phlow.jiratoken")
	if err != nil {
		return nil, err
	}
	list, err := plugins.QueryIssues(conf.IssueApi, user, token)
	if err != nil {
		return nil, err
	}
	return list, nil
}

//PrintIssues ...
func PrintIssues(conf *setting.ProjectSetting, fetch Fetch) {
	list, err := fetch(conf)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, iss := range list {
		fmt.Println(iss.ToString())
	}

}
