package githandler

import (
	"fmt"
	"regexp"
	"strings"

	. "github.com/praqma/git-phlow/executor"
)

//ConfigBranchRemote ...
func ConfigBranchRemote(branch string) string {
	configArg := fmt.Sprintf("branch.%s.remote", branch)
	output, _ := ExecuteCommand("git", "config", configArg)
	return strings.Replace(output, "\n", "", -1)
}

//ConfigGet ...
func ConfigGet(key, group string) string {
	pair := fmt.Sprintf("%s.%s", group, key)
	output, _ := ExecuteCommand("git", "config", "--global", "--get", pair)
	return strings.Replace(output, "\n", "", -1)
}

//ConfigSet ...
func ConfigSet(key, value, group string) error {
	pair := fmt.Sprintf("%s.%s", group, key)
	_, err := ExecuteCommand("git", "config", "--global", pair, value)
	return err
}

//CheckOut ...
func CheckOut(branch string) error {
	_, err := ExecuteCommand("git", "checkout", branch)
	return err
}

//CheckoutNewBranchFromRemote ...
func CheckoutNewBranchFromRemote(branch, defaultBranch string) error {
	remote := ConfigBranchRemote(defaultBranch)
	_, err := ExecuteCommand("git", "checkout", "-b", branch, remote+"/"+defaultBranch)
	return err
}

//Status ...
func Status() error {
	_, err := ExecuteCommand("git", "status")
	return err
}

//Add ...
func Add() error {
	_, err := ExecuteCommand("git", "add", "--all")
	return err
}

//Commit ...
func Commit(message string) (string, error) {
	return ExecuteCommand("git", "commit", "-m", message)
}

//Fetch ...
func Fetch() error {
	_, err := ExecuteCommand("git", "fetch", "--all")
	return err
}

//Pull ...
func Pull() (string, error) {
	return ExecuteCommand("git", "pull", "--rebase")
}

//Push ...
func Push() (string, error) {
	return ExecuteCommand("git", "push")
}

//PushRename ...
func PushRename(branch, defaultBranch string) (string, error) {
	remote := ConfigBranchRemote(defaultBranch)
	str := fmt.Sprintf("%s:ready/%s", branch, branch)

	return ExecuteCommand("git", "push", remote, str)
}

//Merge ...
func Merge(branch string) error {
	_, err := ExecuteCommand("git", "merge", branch)
	return err
}

//RemoteInfo ...
type RemoteInfo struct {
	Organisation string
	Repository   string
}

//Remote ...
func Remote() (*RemoteInfo, error) {
	re := regexp.MustCompile(`.+:(\S+)\/(\S+)\.git`)
	output, err := ExecuteCommand("git", "remote", "-v")
	if err != nil {
		return nil, err
	}
	match := re.FindStringSubmatch(output)
	return &RemoteInfo{match[1], match[2]}, nil
}

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

	output, lErr := ExecuteCommand("git", "branch")
	if lErr != nil {
		return nil, err
	}

	info.Current = strings.TrimSpace(current)
	for _, branch := range strings.Split(output, "\n") {
		if branch != "" {
			info.List = append(info.List, strings.TrimSpace(branch))
		}
	}
	return &info, err
}

//BranchRename ...
func BranchRename(name string) error {
	_, err := ExecuteCommand("git", "branch", "-m", name, "delivered/"+name)
	return err
}
