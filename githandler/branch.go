package githandler

import (
	"strconv"
	"strings"

	"github.com/praqma/git-phlow/executor"
)

//BranchInfo ...
type BranchInfo struct {
	Current string
	List    []string
}

//AsList ...
func AsList(branchOutput string) *BranchInfo {
	var info BranchInfo

	for _, branch := range strings.Split(branchOutput, "\n") {
		if branch != "" {
			if strings.HasPrefix(branch, "*") {
				branch = strings.TrimPrefix(branch, "*")
				branch = strings.TrimSpace(branch)
				info.Current = branch
			}
			branch = strings.TrimSpace(branch)
			info.List = append(info.List, branch)
		}
	}
	return &info
}

//Delivered ...
func Delivered(info *BranchInfo, remote string) (localBranches []string, remoteBranches []string) {

	for _, branch := range info.List {
		if strings.HasPrefix(branch, "delivered/") {
			localBranches = append(localBranches, branch)
		}
		if strings.HasPrefix(branch, "remotes/"+remote+"/delivered") {
			branch = strings.TrimPrefix(branch, "remotes/"+remote+"/")
			remoteBranches = append(remoteBranches, branch)
		}
	}
	return
}

//Ready ...
func Ready(info *BranchInfo, remote string, prefix string) (remoteBranches []string) {
	for _, branch := range info.List {
		if strings.HasPrefix(branch, "remotes/"+remote+"/"+prefix) {
			branch = strings.TrimPrefix(branch, "remotes/")
			remoteBranches = append(remoteBranches, branch)
		}
	}
	return
}

//DEPRECETED SECTION - USE GIT
//BranchRename ...
func BranchRename(name string) error {
	_, err := executor.RunCommand("git", "branch", "-m", name, "delivered/"+name)
	return err
}

//BranchTime ...
func BranchTime(name string) (int, error) {
	output, err := executor.RunCommand("git", "log", "-n 1", name, "--format=format:%ct")
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(output)
}
