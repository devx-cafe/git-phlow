package phlow

import (
	"fmt"
	"github.com/praqma/git-phlow/executor"
	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/setting"
	"github.com/praqma/git-phlow/options"
	"runtime"
	"github.com/praqma/git-phlow/plugins"
)

//WebCaller ...
//Executes command to web
func WebCaller(args []string) {
	conf := setting.NewProjectStg(options.GlobalFlagTarget)

	var issue = ""
	if len(args) > 0 {
		issue = args[0]
	}

	if conf.Service == "github" {
		OpenGitHub(issue, conf, plugins.IssueFromBranchName, executor.RunCommand)
		return
	}

	if conf.Service == "jira" {
		OpenJira(issue, conf, plugins.KeyFromBranchName, executor.RunCommand)
		return
	}

}

//OpenJira ...
//try to display jira configuration
func OpenJira(issue string, conf *setting.ProjectSetting, extractor plugins.IssueExtractor, runner executor.Runner) {
	git := githandler.Git{Run: executor.RunGit}

	branch, err := git.Branch()
	if err != nil {
		fmt.Println(err)
		return
	}
	branchInfo := githandler.AsList(branch)

	if issue != "" {
		LaunchBrowser(conf.IssueWeb+"/browse/"+issue, runner)
		return
	}

	//No issues to get, just open the browser with the code.. For good measure
	key, err := extractor(branchInfo.Current)
	if err != nil {
		fmt.Println(err)
		LaunchBrowser(conf.IssueWeb, runner)
		return
	}
	LaunchBrowser(conf.IssueWeb+"/browse/"+key, runner)
}

//OpenGitHub ...
//Try to display github issues
func OpenGitHub(issue string, conf *setting.ProjectSetting, extractor plugins.IssueExtractor, runner executor.Runner) {
	git := githandler.Git{Run: executor.RunGit}

	branch, err := git.Branch()
	if err != nil {
		fmt.Println(err)
		return
	}

	branchInfo := githandler.AsList(branch)

	remote, err := git.LSRemote("--get-url", conf.Remote)
	if err != nil {
		fmt.Println(err)
		return
	}

	orgAndRepo := githandler.OrgAndRepo(remote)

	ext := "/" + orgAndRepo.Organisation + "/" + orgAndRepo.Repository

	if issue != "" {
		LaunchBrowser(conf.IssueWeb+ext+"/issues/"+issue, runner)
		return
	}

	//No issues to get, just open the browser with the code.. For good measure
	key, err := extractor(branchInfo.Current)
	if err != nil {
		fmt.Println(err)
		LaunchBrowser(conf.IssueWeb+ext, runner)
		return
	}
	LaunchBrowser(conf.IssueWeb+ext+"/issues/"+key, runner)
}

func LaunchBrowser(link string, run executor.Runner) {

	if runtime.GOOS == "windows" {
		run("explorer", link)
		return
	}
	run("open", link)
}
