package githandler

import (
	"regexp"
	"strings"

	"github.com/praqma/git-phlow/executor"
	"bytes"
	"os/exec"
	"fmt"
)

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

//FormatPatch ...
//dry runs patch to see if we can auto merge
func FormatPatch(buf *bytes.Buffer, remoteBranch string) (err error) {
	err = executor.ExecPipeCommand(buf,
		exec.Command("git", "format-patch", remoteBranch, "--stdout"),
		exec.Command("git", "apply", "check"),
	)
	return
}

//Status ...
//git status
func Status() error {
	_, err := executor.ExecuteCommand("git", "status")
	return err
}

//StatusPorcelain ...
//generates behind and ahead status
func StatusPorcelain() (string, error) {
	out, err := executor.ExecuteCommand("git", "status", "short", "--branch", "--porcelain")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out), nil
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
//Must have either origin or upstream
//THIS NEEDS TO BE REVISITED
func Remote() (*RemoteInfo, error) {
	var res string
	var err error
	if res, err = executor.ExecuteCommand("git", "ls-remote", "--get-url", "origin"); err != nil {
		return nil, err
	}
	res = strings.Trim(res, "\n")
	return remoteURLExtractor(res), nil
}

func remoteURLExtractor(url string) *RemoteInfo {
	re := regexp.MustCompile(`.+:(\S+)\/(\S+)\.git`)

	//Extracts repo and org from ssh url format
	if strings.HasPrefix(url, "git@") {
		match := re.FindStringSubmatch(url)
		return &RemoteInfo{match[1], match[2]}
	}
	//Extracts repo and org from http url format
	if strings.HasPrefix(url, "http") {
		splitURL := strings.Split(strings.TrimSuffix(url, ".git"), "/")
		org := splitURL[len(splitURL)-2]
		repo := splitURL[len(splitURL)-1]
		return &RemoteInfo{org, repo}
	}

	//Clone from local repo
	return &RemoteInfo{}
}

//ConfigBranchRemote ...
func ConfigBranchRemote(branch string) string {
	configArg := fmt.Sprintf("branch.%s.remote", branch)
	output, _ := executor.ExecuteCommand("git", "config", configArg)
	return strings.Replace(output, "\n", "", -1)
}
