package githandler

import (
	. "github.com/praqma/git-phlow/executor"
	"strconv"
	"strings"
)

//BranchInfo ...
type BranchInfo struct {
	Current string
	List    []string
}

//Branch ...
func Branch() (*BranchInfo, error) {
	var err error
	info := BranchInfo{}

	current, cErr := ExecuteCommand("git", "rev-parse", "--abbrev-ref", "HEAD")
	if cErr != nil {
		return nil, err
	}

	output, lErr := ExecuteCommand("git", "branch", "-a")
	if lErr != nil {
		return nil, err
	}

	info.Current = strings.TrimSpace(current)
	for _, branch := range strings.Split(output, "\n") {
		if branch != "" {
			branch = strings.TrimPrefix(branch, "*")
			branch = strings.TrimSpace(branch)
			info.List = append(info.List, branch)
		}
	}
	return &info, err
}

//BranchRename ...
func BranchRename(name string) error {
	_, err := ExecuteCommand("git", "branch", "-m", name, "delivered/"+name)
	return err
}

//BranchDelete ...
func BranchDelete(name, remote string, deleteRemote, force bool) (string, error) {
	if deleteRemote {
		return ExecuteCommand("git", "push", remote, "--delete", name)
	}

	if force {
		return ExecuteCommand("git", "branch", "-D", name)
	}
	return ExecuteCommand("git", "branch", "-d", name)
}

//BranchDelivered ...
func BranchDelivered(remote string) (localBranches []string, remoteBranches []string) {
	info, err := Branch()

	if err != nil {
		return
	}

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

//BranchReady ...
func BranchReady(remote string) (remoteBranches []string) {
	info, err := Branch()
	if err != nil {
		return
	}

	for _, branch := range info.List {
		if strings.HasPrefix(branch, "remotes/"+remote+"/ready") {
			branch = strings.TrimPrefix(branch, "remotes/")
			remoteBranches = append(remoteBranches, branch)
		}
	}
	return
}

//BranchTime ...
func BranchTime(name string) (int, error) {
	output, err := ExecuteCommand("git", "log", "-n 1", name, "--format=format:%ct")
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(output)
}
