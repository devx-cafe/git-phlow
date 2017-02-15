package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	RequestRepositoriesUri  = "https://api.github.com/repos/"
	RequestAuthorizationUri = "https://api.github.com/authorizations"
)

//Repo contains repo information
type Repo struct {
	DefaultBranch string `json:"default_branch"`
}

func GetMainBranch(RequestRepositoriesUri string) (string, error) {

	res, err := http.Get(RequestRepositoriesUri)

	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request did not respond 200 OK: %s", res.Status)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	re := Repo{}
	err = json.Unmarshal(body, &re)

	if err != nil {
		return "", err
	}

	return re.DefaultBranch, nil
}

//GetRepoInfo ...  gets repo info from name
func GetRepoInfo(repoName string) (out Repo) {

	res, err := http.Get("https://api.github.com/repos/" + repoName)
	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	re := Repo{}
	err = json.Unmarshal(body, &re)
	if err != nil {
		panic(err)
	}

	fmt.Println(re.DefaultBranch)

	return re
}

func getRepoNameAndOrg(remote string) (name, org string) {
	var repoInfo []string

	tmp := strings.Split(remote, ":")[1]
	tmp = strings.Split(tmp, ".")[0]

	repoInfo = strings.Split(tmp, "/")

	org = repoInfo[0]
	name = repoInfo[1]
	return
}
