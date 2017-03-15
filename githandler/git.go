package githandler

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/praqma/git-phlow/executor"
)

//ConfigBranchRemote ...
func ConfigBranchRemote(branch string) string {
	configArg := fmt.Sprintf("branch.%s.remote", branch)
	output, _ := executor.ExecuteCommand("git", "config", configArg)
	return strings.Replace(output, "\n", "", -1)
}

//ConfigGet ...
func ConfigGet(key, group string) string {
	pair := fmt.Sprintf("%s.%s", group, key)
	output, _ := executor.ExecuteCommand("git", "config", "--global", "--get", pair)
	return strings.Replace(output, "\n", "", -1)
}

//ConfigSet ...
func ConfigSet(key, value, group string) error {
	pair := fmt.Sprintf("%s.%s", group, key)
	_, err := executor.ExecuteCommand("git", "config", "--global", pair, value)
	return err
}

//CheckOut ...
func CheckOut(branch string) error {
	_, err := executor.ExecuteCommand("git", "checkout", branch)
	return err
}

//CheckoutNewBranchFromRemote ...
func CheckoutNewBranchFromRemote(branch, defaultBranch string) error {
	remote := ConfigBranchRemote(defaultBranch)
	_, err := executor.ExecuteCommand("git", "checkout", "-b", branch, remote+"/"+defaultBranch)
	return err
}

//Status ...
func Status() error {
	_, err := executor.ExecuteCommand("git", "status")
	return err
}

//Add ...
func Add() error {
	_, err := executor.ExecuteCommand("git", "add", "--all")
	return err
}

//Commit ...
func Commit(message string) (string, error) {
	return executor.ExecuteCommand("git", "commit", "-m", message)
}

//Fetch ...
func Fetch() error {
	_, err := executor.ExecuteCommand("git", "fetch", "--all")
	return err
}

//FetchPrune ...
func FetchPrune() error {
	_, err := executor.ExecuteCommand("git", "fetch", "--prune")
	return err
}

//Pull ...
func Pull() (string, error) {
	return executor.ExecuteCommand("git", "pull", "--rebase")
}

//Push ...
func Push() (string, error) {
	return executor.ExecuteCommand("git", "push")
}

//PushRename ...
func PushRename(branch, defaultBranch string) (string, error) {
	remote := ConfigBranchRemote(defaultBranch)
	str := fmt.Sprintf("%s:ready/%s", branch, branch)

	return executor.ExecuteCommand("git", "push", remote, str)
}

//Merge ...
func Merge(branch string) error {
	_, err := executor.ExecuteCommand("git", "merge", branch)
	return err
}

//RemoteInfo ...
type RemoteInfo struct {
	Organisation string
	Repository   string
}

//Remote ...
func Remote(defaultBranch string) (*RemoteInfo, error) {
	re := regexp.MustCompile(`.+:(\S+)\/(\S+)\.git`)
	var res string
	var err error

	if res, err = executor.ExecuteCommand("git", "config", fmt.Sprintf("branch.%s.remote", defaultBranch)); err != nil {
		return nil, err
	}
	res = strings.Trim(res, "\n")
	if res, err = executor.ExecuteCommand("git", "config", "--get", fmt.Sprintf("remote.%s.url", res)); err != nil {
		return nil, err
	}
	res = strings.Trim(res, "\n")
	match := re.FindStringSubmatch(res)
	return &RemoteInfo{match[1], match[2]}, nil
}
