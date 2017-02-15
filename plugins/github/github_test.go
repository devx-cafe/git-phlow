/*
Test package for githubwrapper Branch
*/
package github

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"net/http/httptest"
	"net/http"
)

func TestGetRepoNameAndOrg(t *testing.T) {
	Convey("Test Function GetRepoNameAndOrg", t, func() {

		Convey("Organization/name and Repo Name should be extracted", func() {
			var fetchUrl = "origin	git@github.com:Praqma/git-phlow.git (fetch)"

			name, org := getRepoNameAndOrg(fetchUrl)

			So(org, ShouldEqual, "Praqma")
			So(name, ShouldEqual, "git-phlow")
		})
	})
}

func TestMakeRequest(t *testing.T) {

	Convey("Test GetMainBranch function", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected Request 'POST', got '%s'", r.Method)
			}

			if r.URL.EscapedPath() != "/repos" {
				t.Errorf("Expected request to 'repos', got '%s'", r.URL.EscapedPath())
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{\"default_branch\":\"master\"}"))

		}))

		defer ts.Close()

		name, err := GetMainBranch(ts.URL + "/repos")
		So(name, ShouldEqual, "master")
		So(err, ShouldBeNil)
	})
}
