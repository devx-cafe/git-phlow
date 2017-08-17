package githandler

import (
	"strings"
	"regexp"
	"fmt"
	"github.com/praqma/git-phlow/executor"
)

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
	if res, err = executor.RunCommand("git", "ls-remote", "--get-url", "origin"); err != nil {
		return nil, err
	}
	res = strings.Trim(res, "\n")
	return remoteURLExtractor(res), nil
}

//OrgAndRepo ...
//Expects input from LSRemote
func OrgAndRepo(url string) *RemoteInfo {
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
	output, _ := executor.RunCommand("git", "config", configArg)
	return strings.Replace(output, "\n", "", -1)
}
