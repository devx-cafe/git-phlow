package plugins

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/code-cafe/git-phlow/options"
)

//pluginWebURL ...
//required paths to api call
type pluginWebURL struct {
	base        string
	authURL     string
	issueURL    string
	labelURL    string
	assigneeURL string
	repo        string
	userRepo    string
}

//pluginWebRequest ...
type pluginWebRequest struct {
	client *http.Client
	res    *http.Response
	body   *[]byte
}

//URLNoEsc ...
//returns full path of url
func (p *pluginWebURL) URLNoEsc(URL string) string {
	return p.base + URL
}

//NewPWRequest ...
//Creates a new pluginWebRequest
func NewPWRequest() *pluginWebRequest {
	return &pluginWebRequest{res: &http.Response{}, body: &[]byte{}, client: http.DefaultClient}
}

//Do ...
//Execute a web request
func (g *pluginWebRequest) Do(r *http.Request) (body []byte, err error) {

	if options.GlobalFlagVerbose {
		fmt.Println("web request URL: " + r.URL.String())
	}

	if g.res, err = g.client.Do(r); err != nil {
		return nil, err
	}

	defer g.res.Body.Close()

	if err = requestStatus(g.res); err != nil {
		return nil, err
	}

	if body, err = ioutil.ReadAll(g.res.Body); err != nil {
		return nil, err
	}
	return body, err
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
		return errors.New("git-phlow token already exists - go to GitHub and delete your git-phlow Personal Access Token")
	case http.StatusNotFound:
		//For GET and POST
		return fmt.Errorf("responded with %s - malformed url", res.Status)
	case http.StatusUnauthorized:
		//For 401 unauthorized
		return fmt.Errorf("Request returned 401 unauthorized. Please run 'git phlow auth' and reauthenticate")
	default:
		//Default behaviour if status is not OK
		return fmt.Errorf("request did not respond with 200 OK, but %s", res.Status)
	}
}
