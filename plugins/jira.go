package plugins

import (
	"net/http"
	"github.com/go-errors/errors"
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
