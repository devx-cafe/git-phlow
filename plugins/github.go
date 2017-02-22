package plugins

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/praqma/git-phlow/githandler"
	"strconv"
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

type Label struct {
	Id    int     `json:"id"`
	Url   string  `json:"url"`
	Name  string  `json:"name"`
	Color string  `json:"color"`
}

type Assignee struct {
	Assignees []string `json:"assignees"`
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

//SetLabel ...
func SetLabel(label, url, token string, number int) ([]Label, error) {
	client := &http.Client{}
	info, err := githandler.Remote()
	var body = `[ "` + label + `" ]`
	var uri = url + info.Organisation + "/" + info.Repository + "/issues/" + strconv.Itoa(number) + "/labels" //"/repos/:owner/:repo/issues/:number/labels"

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("githup responses with %s", resp.Status)
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	re := []Label{}
	output, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(output, &re)
	if err != nil {
		return nil, err
	}
	return re, nil

}

//SetAssignee ...
//Sets you as an assignee
func SetAssignee(assignee, url, token string, number int) error {
	client := &http.Client{}
	info, _ := githandler.Remote()

	var apiURL = fmt.Sprintf(url+"%s/%s/issues/%s/assignees", info.Organisation, info.Repository, strconv.Itoa(number))
	jsonBytes, _ := json.Marshal(Assignee{Assignees: []string{assignee}})

	req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+token)

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("githup responses with %s", resp.Status)
	}

	return nil
}
