package plugins

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/praqma/git-phlow/githandler"
)

var (
	authBody = `{"scopes": ["public_repo"],"note": "git phlow"}`
	AuthURL  = "https://api.github.com/authorizations"
	RepoUrl  = "https://api.github.com/repos/"
)

//Auth ...
type Auth struct {
	Token string `json:"token"`
}

//Repo ...
type Repo struct {
	DefaultBranch string `json:"default_branch"`
}

//Issues ...
type Issues struct {
	Title  string `json:"title"`
	Number int    `json:"number"`
}

//GetOpenIssues ...
func GetOpenIssues(url string) ([]Issues, error) {
	info, err := githandler.Remote()
	res, _ := http.Get(url + info.Organisation + "/" + info.Repository + "/issues")

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request did not respond 200 OK: %s", res.Status)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	re := []Issues{}
	err = json.Unmarshal(body, &re)

	if err != nil {
		return nil, err
	}
	return re, nil

}

//Authorize ...
func Authorize(user, pass, url string) (string, error) {
	var auth Auth
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(authBody)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(user, pass)

	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("githup responses with %s", resp.Status)
	}
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	output, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(output, &auth)
	if err != nil {
		return "", err
	}
	return auth.Token, nil
}

//GetDefaultBranch ...
func GetDefaultBranch(url string) (string, error) {

	info, err := githandler.Remote()
	res, _ := http.Get(url + info.Organisation + "/" + info.Repository)
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

//SetAssignee ...
//Sets you as an assignee
func SetAssignee() error {
	return nil
}

//SetLabel ...
func SetLabel(label string) error {
	return nil //label - 'Status - in progress'
}
