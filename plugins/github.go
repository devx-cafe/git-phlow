package plugins

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/praqma/git-phlow/options"
)

//AuthorizeGitHub ...
//Retrieve token from github for authorization
func AuthorizeGitHub(githubBaseURL, user, pass string) (token string, err error) {

	perm, err := createGHPermissions()
	if err != nil {
		return "", err
	}

	req, _ := http.NewRequest("POST", githubBaseURL+"/authorizations", bytes.NewBuffer([]byte(perm)))
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.SetBasicAuth(user, pass)

	body, err := NewPWRequest().Do(req)
	if err != nil {
		return "", err
	}

	re := Auth{}
	if err = json.Unmarshal(body, &re); err != nil {
		return "", err
	}
	return re.Token, nil
}

//AuthenticateGitHub ...
//Checks personal access token validity by requesting private repositories and checking status code
func AuthenticateGitHub(githubBaseURL string, user, token string) (bool, error) {

	req, _ := http.NewRequest("GET", githubBaseURL+"/user/repos", nil)
	q := req.URL.Query()
	q.Add("access_token", token)
	req.URL.RawQuery = q.Encode()
	client := http.DefaultClient

	res, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return false, errors.New(strconv.Itoa(res.StatusCode))
	}
	return true, nil
}

//DefaultBranchGitHub ...
//return the default branch of the repository
func DefaultBranchGitHub(URL, org, repo, token string) (defaultBranch string, err error) {

	req, _ := http.NewRequest("GET", URL+fmt.Sprintf("/repos/%s/%s", org, repo), nil)
	q := req.URL.Query()
	q.Add("access_token", token)
	req.URL.RawQuery = q.Encode()

	client := http.DefaultClient

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	re := Repo{}
	err = json.NewDecoder(res.Body).Decode(&re)
	if err != nil {
		return "", err
	}
	return re.DefaultBranch, nil
}

//This should be rewritten
//createGHPermissions ...
func createGHPermissions() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}

	note := "git phlow " + hostname
	if options.GlobalFlagVerbose {
		fmt.Println("gh plugin: " + note)
	}

	perm := GhPermissions{
		Scopes: []string{"public_repo", "repo", "repo_deployment", "repo:status"},
		Note:   note,
	}
	b2b, err := json.Marshal(&perm)
	if err != nil {
		return "", err
	}
	return string(b2b), nil
}

//GetIssuesGitHub ...
//return an issue with from the number of the issue
func GetIssuesGitHub(URL, org, repo, token string) ([]Stringer, error) {

	req, _ := http.NewRequest("GET", URL+fmt.Sprintf("/repos/%s/%s/issues", org, repo), nil)
	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("access_token", token)
	q.Add("per_page", "30")
	req.URL.RawQuery = q.Encode()

	if options.GlobalFlagVerbose {
		fmt.Println(req.URL)
	}

	client := http.DefaultClient

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == 401 {
		return nil, errors.New("Not Authorized \nVerify that you are authorized by running 'git phlow auth' with the same configuration")
	}

	if res.StatusCode == 404 && (org == "" || repo == "") {
		return nil, errors.New("Could not reach GitHub API - Malformed URL \nVerify 'Remote' field is correct in configuration" +
			"\ntry 'git ls-remote --get <Remote from config>' should return: git@github.com:org/repo.git")
	}

	if res.StatusCode != 200 {
		return nil, errors.New("Could not get list of issues ")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	re := []Issue{}
	if err = json.Unmarshal(body, &re); err != nil {
		return nil, err
	}

	iss := make([]Stringer, len(re))
	for i, v := range re {
		iss[i] = v
	}

	return iss, nil
}

//GetIssueGitHub ...
//get a single issue from github
func GetIssueGitHub(URL, org, repo, key, token string) (*Issue, error) {

	req, _ := http.NewRequest("GET", URL+fmt.Sprintf("/repos/%s/%s/issues/", org, repo)+key, nil)
	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("access_token", token)
	req.URL.RawQuery = q.Encode()

	if options.GlobalFlagVerbose {
		fmt.Println(req.URL)
	}

	client := http.DefaultClient

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == 401 {
		return nil, errors.New("Not Authorized \nVerify that you are authorized by running 'git phlow auth' with the same configuration")
	}

	if res.StatusCode == 404 && (org == "" || repo == "") {
		return nil, errors.New("Could not reach GitHub API - Malformed URL \nVerify 'Remote' field is correct in configuration" +
			"\ntry 'git ls-remote --get <Remote from config>' should return: git@github.com:org/repo.git")
	}

	if res.StatusCode != 200 {
		return nil, errors.New("An issue with number: " + key +
			" could not be found on GitHub \ncheck that it exists or that you have access to it on GitHub")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	re := Issue{}
	if err = json.Unmarshal(body, &re); err != nil {
		return nil, err
	}

	return &re, err
}

//SetLabelGitHub ...
//Sets a label on a github repository with
func SetLabelGitHub(URL, org, repo, token, label, issue string) (labels []Label, err error) {

	req, _ := http.NewRequest("POST", URL+fmt.Sprintf("/repos/%s/%s/issues/%s/labels", org, repo, issue), bytes.NewBuffer([]byte(`[ "`+label+`" ]`)))
	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("access_token", token)
	req.URL.RawQuery = q.Encode()

	if options.GlobalFlagVerbose {
		fmt.Println(req.URL)
	}

	client := http.DefaultClient

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	re := []Label{}
	if err = json.Unmarshal(body, &re); err != nil {
		return nil, err
	}
	return re, nil
}

//SetAssigneeGitHub ...
//Assigns a user to a github issue
func SetAssigneeGitHub(URL, org, repo, token, issue, assignee string) (err error) {

	jsonBytes, _ := json.Marshal(Assignee{Assignees: []string{assignee}})

	req, _ := http.NewRequest("POST", URL+fmt.Sprintf("/repos/%s/%s/issues/%s/assignees", org, repo, issue), bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("access_token", token)
	req.URL.RawQuery = q.Encode()

	if options.GlobalFlagVerbose {
		fmt.Println(req.URL)
	}

	client := http.DefaultClient

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
