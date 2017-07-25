package plugins

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"os"

	"github.com/praqma/git-phlow/githandler"
	"github.com/praqma/git-phlow/options"
)

var GitHub *GitHubImpl
var urls *pluginWebURL

//GitHubImpl ...
//GitHub api implementation
type GitHubImpl struct {
	*pluginWebURL
	repo  string
	org   string
	token string
}

//init ...
//creates a new GitHub request object with all the gh api urls
func init() {
	urls = &pluginWebURL{
		base:        "https://api.github.com",
		authURL:     "/authorizations",
		assigneeURL: "/repos/%s/%s/issues/%d/assignees",
		issueURL:    "/repos/%s/%s/issues",
		labelURL:    "/repos/%s/%s/issues/%d/labels",
		repo:        "/repos/%s/%s",
		userRepo:    "/user/repos",
	}

	info, _ := githandler.Remote()
	org := info.Organisation
	repo := info.Repository
	token := githandler.ConfigGet("token", "phlow")

	GitHub = &GitHubImpl{
		urls,
		repo,
		org,
		token,
	}
}

//GetIssues ...
func (g *GitHubImpl) GetIssues() (issues []Issues, err error) {
	URL := fmt.Sprintf(g.URLNoEsc(urls.issueURL), g.org, g.repo)

	req, _ := http.NewRequest("GET", URL, nil)
	q := req.URL.Query()
	q.Add("access_token", g.token)
	req.URL.RawQuery = q.Encode()

	body, err := NewPWRequest().Do(req)
	if err != nil {
		return nil, err
	}

	re := []Issues{}
	if err = json.Unmarshal(body, &re); err != nil {
		return nil, err
	}
	return re, nil
}

//SetLabel ...
func (g *GitHubImpl) SetLabel(label string, issue int) (labels []Label, err error) {

	URL := fmt.Sprintf(g.URLNoEsc(urls.labelURL), g.org, g.repo, issue)

	req, _ := http.NewRequest("POST", URL, bytes.NewBuffer([]byte(`[ "`+label+`" ]`)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+g.token)

	body, err := NewPWRequest().Do(req)
	if err != nil {
		return nil, err
	}

	re := []Label{}
	if err = json.Unmarshal(body, &re); err != nil {
		return nil, err
	}
	return re, nil
}

//Default ...
//Get default branch of a GitHub issue
func (g *GitHubImpl) Default() (defUrl string, err error) {

	URL := fmt.Sprintf(g.URLNoEsc(g.pluginWebURL.repo), g.org, g.repo)

	req, _ := http.NewRequest("GET", URL, nil)
	q := req.URL.Query()
	q.Add("access_token", g.token)
	req.URL.RawQuery = q.Encode()

	body, err := NewPWRequest().Do(req)
	if err != nil {
		return "", err
	}

	re := Repo{}
	if err = json.Unmarshal(body, &re); err != nil {
		return "", err
	}
	return re.DefaultBranch, nil
}

//SetAssignee ...
//Set assignee on a GitHub Issue
func (g *GitHubImpl) SetAssignee(assignee string, issue int) (err error) {

	URL := fmt.Sprintf(g.URLNoEsc(urls.assigneeURL), g.org, g.repo, issue)

	jsonBytes, _ := json.Marshal(Assignee{Assignees: []string{assignee}})
	req, _ := http.NewRequest("POST", URL, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+g.token)

	_, err = NewPWRequest().Do(req)
	if err != nil {
		return err
	}
	return nil
}

//GhPermissions ...
//data struct for permissions
type GhPermissions struct {
	Scopes []string `json:"scopes"`
	Note   string   `json:"note"`
}

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

//Auth ...
//Auth request to gh
func (g *GitHubImpl) Auth(user, pass string) (token string, err error) {

	perm, err := createGHPermissions()
	if err != nil {
		return "", err
	}

	req, _ := http.NewRequest("POST", g.URLNoEsc(urls.authURL), bytes.NewBuffer([]byte(perm)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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

//CheckAuth ...
//Checks personal access token validity by requesting private repositories and checking status code
func (g *GitHubImpl) CheckAuth() (bool, error) {

	URL := fmt.Sprintf(g.URLNoEsc(urls.userRepo))

	req, _ := http.NewRequest("GET", URL, nil)
	q := req.URL.Query()
	q.Add("access_token", g.token)
	req.URL.RawQuery = q.Encode()

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return false, err
	}
	if res.StatusCode != http.StatusOK {
		return false, errors.New(strconv.Itoa(res.StatusCode))
	}
	return true, nil
}
