package plugins

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"strconv"

	"github.com/praqma/git-phlow/githandler"
	"errors"
	"github.com/praqma/git-phlow/options"
)

var (
	authBody = `{"scopes": ["public_repo"],"note": "git phlow"}`
	//AuthURL ...
	AuthURL = "https://api.github.com/authorizations"
	//RepoURL ...
	RepoURL = "https://api.github.com/repos/"
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

//Label ...
type Label struct {
	ID    int    `json:"id"`
	URL   string `json:"url"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

//Assignee ...
type Assignee struct {
	Assignees []string `json:"assignees"`
}

//GetOpenIssues ...
func GetOpenIssues(url string) ([]Issues, error) {

	info, err := githandler.Remote("master")
	if err != nil {
		return nil, err
	}

	if options.GlobalFlagVerbose {
		fmt.Println(info)
		fmt.Println(url + info.Organisation + "/" + info.Repository + "/issues")
	}

	res, _ := http.Get(url + info.Organisation + "/" + info.Repository + "/issues")

	if res.StatusCode != http.StatusOK {
		switch res.StatusCode {
		case http.StatusUnprocessableEntity:
			break
			return nil, errors.New("Token with git-phlow signature already exists")
		case http.StatusNotFound:
			break
			return nil, fmt.Errorf("responded with %s - Url is wrong", res.Status)
		default:
			return nil, fmt.Errorf("request did not respond 200 OK: %s", res.Status)
		}
	}

	if res.StatusCode != http.StatusUnprocessableEntity {

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

	if options.GlobalFlagVerbose {
		fmt.Println(url)
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(authBody)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(user, pass)

	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("githup responsed with %s", resp.Status)
	}
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	output, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(output, &auth)
	if err != nil {
		return "", err
	}
	return auth.Token, nil
}

//GetDefaultBranch ...
func GetDefaultBranch(url string) (string, error) {
	info, err := githandler.Remote("master")
	if err != nil {
		return "", err
	}
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
func SetLabel(label, url, token string, number int, info *githandler.RemoteInfo) ([]Label, error) {
	client := &http.Client{}
	var body = `[ "` + label + `" ]`
	var uri = url + info.Organisation + "/" + info.Repository + "/issues/" + strconv.Itoa(number) + "/labels" //"/repos/:owner/:repo/issues/:number/labels"

	req, _ := http.NewRequest("POST", uri, bytes.NewBuffer([]byte(body)))
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
	output, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(output, &re)
	if err != nil {
		return nil, err
	}
	return re, nil

}

//SetAssignee ...
//Sets you as an assignee
func SetAssignee(assignee, url, token string, number int, info *githandler.RemoteInfo) error {
	client := &http.Client{}

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
