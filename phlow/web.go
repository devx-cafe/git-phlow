package phlow

import (
	"github.com/praqma/git-phlow/executor"
	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/plugins"
	"fmt"
	"os"
)

func Web(userIssue int) {

	var issueListURL = "https://github.com/%s/%s/issues"
	var issueInfoURL = "https://github.com/%s/%s/issues/%d"

	//Get organisation and repository for repository
	orgAndRepo, err := githandler.Remote()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Get your current branch
	branchInfo, err := githandler.Branch()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	branchIssue := plugins.IssueFromBranchName(branchInfo.Current)

	if userIssue != -1 {
		issueInfoURL = fmt.Sprintf(issueInfoURL, orgAndRepo.Organisation, orgAndRepo.Repository, userIssue)
		_, err := executor.ExecuteCommand("open", issueInfoURL)
		if err != nil {
			fmt.Println(err)
		}
		os.Exit(0)
	}

	if branchIssue != -1 {
		issueInfoURL = fmt.Sprintf(issueInfoURL, orgAndRepo.Organisation, orgAndRepo.Repository, branchIssue)
		executor.ExecuteCommand("open", issueInfoURL)
		os.Exit(0)
	}

	issueListURL = fmt.Sprintf(issueListURL, orgAndRepo.Organisation, orgAndRepo.Repository)
	executor.ExecuteCommand("open", issueListURL)
}
