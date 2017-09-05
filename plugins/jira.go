package plugins

import (
	"net/http"
	"github.com/go-errors/errors"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"bytes"
)

//AuthenticateJIRA ...
func AuthenticateJIRA(url, user, pass string) (bool, error) {
	verifyURL := "/rest/api/latest/mypermissions"
	req, _ := http.NewRequest("GET", url+verifyURL, nil)
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(user, pass)
	client := http.DefaultClient

	res, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return false, errors.New("Could not verify credentials")
	}
	return true, nil
}

//AuthorizeJIRA ...
//JIRA basic auth do not require authorization, so we just validate the user have giving right credentials with AuthenticateJIRA
func AuthorizeJIRA(URL, user, pass string) (token string, err error) {

	authorized, err := AuthenticateJIRA(URL, user, pass)
	if err != nil {
		return "", err
	}

	if authorized == false {
		return "", errors.New("not authorized")
	}

	return pass, nil
}

//GetJiraIssue ...
func GetJiraIssue(URL, key, user, pass string) (*JiraIssue, error) {

	issueURL := "/rest/api/latest/issue/"
	req, _ := http.NewRequest("GET", URL+issueURL+key, nil)
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(user, pass)
	client := http.DefaultClient

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body);
	if err != nil {
		return nil, err
	}

	re := JiraIssue{}
	if err = json.Unmarshal(body, &re); err != nil {
		return nil, err
	}

	return &re, err
}

//GetTranstions ...
//Retrieve transitions for a specific issue
func GetTransitions(URL, key, user, pass string) (*Transitions, error) {
	issueURL := "/rest/api/latest/issue/%s/transitions"

	req, _ := http.NewRequest("GET", URL+fmt.Sprintf(issueURL, key), nil)
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(user, pass)
	client := http.DefaultClient

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body);
	if err != nil {
		return nil, err
	}

	re := Transitions{}
	if err = json.Unmarshal(body, &re); err != nil {
		return nil, err
	}

	return &re, err
}

//DoTransition ...
func DoTransition(URL, key, user, pass, transitionID string) (error) {
	issueURL := "/rest/api/latest/issue/%s/transitions"
	transition := TransitionBody{}

	transition.Transition.ID = transitionID
	data, err := json.Marshal(&transition)
	if err != nil {
		return err
	}

	req, _ := http.NewRequest("POST", URL+fmt.Sprintf(issueURL, key), bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(user, pass)
	client := http.DefaultClient

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 204 {
		return errors.New("Could not transition issue")
	}
	return nil
}

//AssignUser ...
func AssignUser(URL, key, user, pass string) (error) {
	issueURL := "/rest/api/latest/issue/%s/assignee"

	data, err := json.Marshal(AssignBody{Name: user})
	if err != nil {
		return err
	}
	req, _ := http.NewRequest("PUT", URL+fmt.Sprintf(issueURL, key), bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(user, pass)
	client := http.DefaultClient

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 204 {
		return errors.New("Could not assign user")
	}
	return nil
}
