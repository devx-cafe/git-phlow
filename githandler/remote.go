package githandler

import (
	"strings"
	"regexp"
)

//RemoteInfo ...
type RemoteInfo struct {
	Organisation string
	Repository   string
}

//OrgAndRepo ...
//Expects input from LSRemote
//Extracts the Organization and repository name from a remote URL returned from git ls-remote
//this is only meant to work for github because they have a multi-tenant system where organization and repository
//is used to query their API
func OrgAndRepo(url string) *RemoteInfo {

	//Extracts repo and org from ssh url format
	re := regexp.MustCompile(`.+:(\S+)\/(\S+)\.git`)

	//if there is no protocol prefix we the two capture groups will fetch
	//organization and repo
	if strings.HasPrefix(url, "git@") {
		match := re.FindStringSubmatch(url)
		return &RemoteInfo{match[1], match[2]}
	}
	//Is the protocol prefix is http we split the urls backwards and remove .git
	if strings.HasPrefix(url, "http") {
		splitURL := strings.Split(strings.TrimSuffix(url, ".git"), "/")
		org := splitURL[len(splitURL)-2]
		repo := splitURL[len(splitURL)-1]
		return &RemoteInfo{org, repo}
	}

	//If any protocol is given (apart from file:// that will fail) the capture groups
	//will fetch the organization and repo
	//There is a redundancy because this regex will also work in the first case,
	//but this has been added as a fix and will work even though it is not completely optimized
	if strings.Contains(url, "://") {
		protoExp := regexp.MustCompile(`\w*:\/\/[\w.@]+(?:\:\d+)?\/(\w+)\/(\w+)\.git`)
		match := protoExp.FindStringSubmatch(url)
		return &RemoteInfo{match[1], match[2]}
	}

	//If for any case we miss an url specification we will return an empty RemoteInfo
	return &RemoteInfo{}
}
