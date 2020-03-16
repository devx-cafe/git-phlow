package plugins

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const OAuthHost = "github.com"

var (
	// ClientID for github OAuth app
	ClientID = "1dd42fc0690c4fa72620"
	// ClientSecret for Github OAuth app
	ClientSecret = "a45bac52a29c0f96ed08997d3ff0981cbf92d405"
)

type OAuthFlow struct {
	Hostname         string
	ClientID         string
	ClientSecret     string
	WriteSuccessHTML func(io.Writer)
	VerboseStream    io.Writer
}

//randomString
func randomString(length int) (string, error) {
	b := make([]byte, length/2)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// ForOS produces an exec.Cmd to open the web browser for different OS
func OpenInBrowser(goos, url string) *exec.Cmd {
	var args []string
	switch goos {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
		r := strings.NewReplacer("&", "^&")
		url = r.Replace(url)
	default:
		args = []string{"xdg-open"}
	}

	args = append(args, url)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	return cmd
}

// ObtainAccessToken guides the user through the browser OAuth flow on GitHub
// and returns the OAuth access token upon completion.
func (oa *OAuthFlow) ObtainAccessToken() (accessToken string, err error) {
	state, _ := randomString(20)

	code := ""
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return
	}
	port := listener.Addr().(*net.TCPAddr).Port

	q := url.Values{}
	q.Set("client_id", oa.ClientID)
	q.Set("redirect_uri", fmt.Sprintf("http://localhost:%d/callback", port))

	q.Set("scope", "repo")
	q.Set("state", state)

	startURL := fmt.Sprintf("https://%s/login/oauth/authorize?%s", oa.Hostname, q.Encode())
	oa.logf("open %s\n", startURL)
	if err := OpenInBrowser(runtime.GOOS, startURL).Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error opening web browser: %s\n", err)
		fmt.Fprintf(os.Stderr, "Please open the following URL manually:\n%s\n", startURL)
	}

	http.Serve(listener, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		oa.logf("server handler: %s\n", r.URL.Path)
		if r.URL.Path != "/callback" {
			w.WriteHeader(404)
			return
		}
		defer listener.Close()
		rq := r.URL.Query()
		if state != rq.Get("state") {
			fmt.Fprintf(w, "Error: state mismatch")
			return
		}
		code = rq.Get("code")
		oa.logf("server received code %q\n", code)
		w.Header().Add("content-type", "text/html")
		if oa.WriteSuccessHTML != nil {
			oa.WriteSuccessHTML(w)
		} else {
			fmt.Fprintf(w, "<p>You have successfully authenticated. You may now close this page.</p>")
		}
	}))

	tokenURL := fmt.Sprintf("https://%s/login/oauth/access_token", oa.Hostname)
	oa.logf("POST %s\n", tokenURL)
	resp, err := http.PostForm(tokenURL,
		url.Values{
			"client_id":     {oa.ClientID},
			"client_secret": {oa.ClientSecret},
			"code":          {code},
			"state":         {state},
		})
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("HTTP %d error while obtaining OAuth access token", resp.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	tokenValues, err := url.ParseQuery(string(body))
	if err != nil {
		return
	}
	accessToken = tokenValues.Get("access_token")
	if accessToken == "" {
		err = errors.New("the access token could not be read from HTTP response")
	}
	return
}

func (oa *OAuthFlow) logf(format string, args ...interface{}) {
	if oa.VerboseStream == nil {
		return
	}
	fmt.Fprintf(oa.VerboseStream, format, args...)
}
