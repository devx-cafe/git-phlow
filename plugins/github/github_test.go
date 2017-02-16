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
