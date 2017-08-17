package phlow

import (
	"fmt"
	"os"

	"github.com/praqma/git-phlow/executor"
	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/plugins"
	"github.com/praqma/git-phlow/setting"
)

//Web ...
//opens the browser on the current or specified issue
func Web(userIssue int) {
	git := githandler.Git{Run: executor.RunGit}
	conf := setting.NewProjectStg("default")

	var issueListURL = "https://github.com/%s/%s/issues"
	var issueInfoURL = "https://github.com/%s/%s/issues/%d"

	//Get organisation and repository for repository
	remote, err := git.LSRemote("--get-url", conf.Remote)
	if err != nil {
		fmt.Println(err)
		return
	}

	branch, err := git.Branch()
	if err != nil {
		fmt.Println(err)
		return
	}
	branchInfo := githandler.AsList(branch)
	orgAndRepo := githandler.OrgAndRepo(remote)

	branchIssue := plugins.IssueFromBranchName(branchInfo.Current)

	if userIssue != -1 {
		issueInfoURL = fmt.Sprintf(issueInfoURL, orgAndRepo.Organisation, orgAndRepo.Repository, userIssue)
		_, err := executor.RunCommand("open", issueInfoURL)
		if err != nil {
			fmt.Println(err)
		}
		os.Exit(0)
	}

	if branchIssue != -1 {
		issueInfoURL = fmt.Sprintf(issueInfoURL, orgAndRepo.Organisation, orgAndRepo.Repository, branchIssue)
		executor.RunCommand("open", issueInfoURL)
		os.Exit(0)
	}

	issueListURL = fmt.Sprintf(issueListURL, orgAndRepo.Organisation, orgAndRepo.Repository)
	executor.RunCommand("open", issueListURL)
}
