package plugins

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/options"
	"os"
)

//GitHub ...
var GitHub GitHubRequest

//AuthURL ...
var authURL = "https://api.github.com/authorizations"

//RepoURL ...
var repoURL = "https://api.github.com/repos/%s/%s"

//IssueUrl
var issueURL = "https://api.github.com/repos/%s/%s/issues"

//labelURL
var labelURL = "https://api.github.com/repos/%s/%s/issues/%d/labels"

var assigneeURL = "https://api.github.com/repos/%s/%s/issues/%d/assignees"

//GitHubRequest ...
//request object for github
type GitHubRequest struct {
	Assignee AssigneeRequest
	Issue    IssueRequest
	Label    LabelRequest
	Branch   BranchRequest
	Auth     AuthRequest
}

type requestData struct {
	URL    string
	repo   string
	org    string
	client *http.Client
}

//IssueRequest ...
//request for GitHub issues
type IssueRequest struct {
	requestData
}

//AssigneeRequest ...
//request for assignees on GitHub
type AssigneeRequest struct {
	requestData
	token string
}

//BranchRequest ...
//request for default branch on GitHub repository
type BranchRequest struct {
	requestData
}

//LabelRequest ...
//request for getting and setting labels on GitHub issues
type LabelRequest struct {
	requestData
	token string
}

//AuthRequest ...
//request for authentication on GitHub
type AuthRequest struct {
	requestData
}

func init() {
	info, _ := githandler.Remote()
	org := info.Organisation
	repo := info.Repository
	token := githandler.ConfigGet("token", "phlow")

	GitHub.Issue = IssueRequest{requestData{issueURL, repo, org, http.DefaultClient}}
	GitHub.Branch = BranchRequest{requestData{repoURL, repo, org, http.DefaultClient}}
	GitHub.Label = LabelRequest{requestData{labelURL, repo, org, http.DefaultClient}, token}
	GitHub.Assignee = AssigneeRequest{requestData{assigneeURL, repo, org, http.DefaultClient}, token}
	GitHub.Auth = AuthRequest{requestData{authURL, repo, org, http.DefaultClient}}
}

//Get ...
//Get All Issues from GitHub
func (i *IssueRequest) Get() ([]Issues, error) {
	var resp *http.Response
	var err error
	var body []byte

	i.URL = fmt.Sprintf(i.URL, i.org, i.repo)

	token := githandler.ConfigGet("token", "phlow")
	request, err := http.NewRequest("GET", i.URL, nil)
	if err != nil {
		return nil, err
	}
	q := request.URL.Query()
	q.Add("access_token", token)
	request.URL.RawQuery = q.Encode()

	if options.GlobalFlagVerbose {
		fmt.Println("github uri: " + request.URL.String())
	}

	if resp, err = i.client.Do(request); err != nil {
		return nil, err
	}
	if err = requestStatus(resp); err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if body, err = ioutil.ReadAll(resp.Body); err == nil {
		re := []Issues{}
		if err = json.Unmarshal(body, &re); err != nil {
			return nil, err
		}
		return re, nil
	}
	return nil, err
}

//Permissions ...
//data struct for permissions
type Permissions struct {
	Scopes []string `json:"scopes"`
	Note   string   `json:"note"`
}

//createPermissions ...
func createPermissions() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}

	note := "git phlow " + hostname
	if options.GlobalFlagVerbose {
		fmt.Println("github plugin: " + note)
	}

	perm := Permissions{
		Scopes: []string{"public_repo", "repo", "repo_deployment", "repo:status"},
		Note:   note,
	}
	b2b, err := json.Marshal(&perm)
	if err != nil {
		return "", err
	}
	return string(b2b), nil
}

//Auth ...
//Auth request to github
func (a *AuthRequest) Auth(user, pass string) (string, error) {
	var auth Auth

	perm, err := createPermissions()
	if err != nil {
		return "", err
	}

	if options.GlobalFlagVerbose {
		fmt.Println(a.URL)
	}
	req, err := http.NewRequest("POST", a.URL, bytes.NewBuffer([]byte(perm)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(user, pass)

	resp, err := a.client.Do(req)
	if err != nil {
		return "", err
	}
	if err = requestStatus(resp); err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if output, err := ioutil.ReadAll(resp.Body); err == nil {
		if err = json.Unmarshal(output, &auth); err != nil {
			return "", err
		}
		return auth.Token, nil
	}

	return "", errors.New("Could not parse json request")
}

//Set ...
//Set assignee on a GitHub Issue
func (a *AssigneeRequest) Set(assignee string, issue int) error {
	var resp *http.Response
	var err error
	a.URL = fmt.Sprintf(a.URL, a.org, a.repo, issue)

	jsonBytes, _ := json.Marshal(Assignee{Assignees: []string{assignee}})
	req, err := http.NewRequest("POST", a.URL, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+a.token)

	if resp, err = a.client.Do(req); err != nil {
		return err
	}
	if err = requestStatus(resp); err != nil {
		return err
	}
	return nil
}

//Set ...
//Set label on a GitHub issue
func (l *LabelRequest) Set(label string, issue int) ([]Label, error) {
	client := &http.Client{}
	var body = `[ "` + label + `" ]`
	var err error
	var resp *http.Response
	var req *http.Request

	l.URL = fmt.Sprintf(l.URL, l.org, l.repo, issue)
	if options.GlobalFlagVerbose {
		fmt.Println("github uri: " + l.URL)
	}

	if req, err = http.NewRequest("POST", l.URL, bytes.NewBuffer([]byte(body))); err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+l.token)

	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err = requestStatus(resp); err != nil {
		return nil, err
	}

	var output []byte
	if output, err = ioutil.ReadAll(resp.Body); err == nil {
		re := []Label{}
		if err = json.Unmarshal(output, &re); err == nil {
			return re, nil
		}
	}
	return nil, err
}

//Default ...
//Get default branch of a GitHub issue
func (b *BranchRequest) Default() (string, error) {
	var resp *http.Response
	var err error
	var body []byte

	b.URL = fmt.Sprintf(b.URL, b.org, b.repo)
	if options.GlobalFlagVerbose {
		fmt.Println("github uri: " + b.URL)
	}
	if resp, err = http.Get(b.URL); err != nil {
		return "", err
	}
	if err = requestStatus(resp); err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if body, err = ioutil.ReadAll(resp.Body); err == nil {
		re := Repo{}
		if err = json.Unmarshal(body, &re); err != nil {
			return "", err
		}
		return re.DefaultBranch, nil
	}
	return "", err

}

//requestStatus ...
//Determines the status of the request
func requestStatus(res *http.Response) error {
	switch res.StatusCode {
	case http.StatusCreated:
		return nil
	case http.StatusOK:
		return nil
	case 422: //StatusUnprocessableEntity
		//For POST requests
		return errors.New("git-phlow token already exists")
	case http.StatusNotFound:
		//For GET and POST
		return fmt.Errorf("responded with %s - malformed url", res.Status)
	default:
		//Default behaviour if status is not OK
		return fmt.Errorf("request did not respond with 200 OK, but %s", res.Status)
	}
}
