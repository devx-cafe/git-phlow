package plugins

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	authBody = `{"scopes": ["public_repo"],"note": "admin script"}`
	authURL  = "https://api.github.com/authorizations"
	repoUrl  = "https://api.github.com/repos/"
)

//Auth ...
type Auth struct {
	Token string `json:"token"`
}

//Repo ...
type Repo struct {
	DefaultBranch string `json:"default_branch"`
}

//Authorize ...
func Authorize(user, pass, url string) (string, error) {
	var auth Auth
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(authBody)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(user, pass)

	resp, err := client.Do(req)
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
func GetDefaultBranch() (string, error) {

	res, _ := http.Get(repoUrl)

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
