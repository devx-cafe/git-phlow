package githubwrapper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Repo contains repo information
type Repo struct {
	DefaultBranch string `json:"default_branch"`
}

//GetRepoInfo gets repo info from name
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
