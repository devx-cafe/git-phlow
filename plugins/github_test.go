package plugins

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"strings"

	"github.com/praqma/git-phlow/githandler"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAuthorize(t *testing.T) {
	Convey("Running tests on 'Authorize' request", t, func() {

		Convey("Authorize should return token", func() {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "POST" {
					t.Errorf("Expected Request 'POST', got '%s'", r.Method)
				}
				if r.URL.EscapedPath() != "/authorizations" {
					t.Errorf("Expected request to 'repos', got '%s'", r.URL.EscapedPath())
				}
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(authResponse))

			}))

			defer ts.Close()
			token, err := Authorize("simon", "password", ts.URL+"/authorizations")
			So(token, ShouldEqual, "abcdefgh12345678")
			So(err, ShouldBeNil)
		})
	})
}

func TestGetDefaultBranch(t *testing.T) {
	SkipConvey("Runnign tests on 'GetDefaultBranch' request", t, func() {
		Convey("GetDefaultBranch should return master", func() {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "GET" {
					t.Errorf("Expected Request 'GET', got '%s'", r.Method)
				}
				if strings.Contains(r.URL.EscapedPath(), "/repos/") {
					t.Errorf("Expected request to 'repo', got '%s'", r.URL.EscapedPath())
				}
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(repoResponse))

			}))

			defer ts.Close()
			defaultBranch, err := GetDefaultBranch(ts.URL + "/repo/")
			So(defaultBranch, ShouldEqual, "master")
			So(err, ShouldBeNil)
		})
	})
}

func TestGetOpenIssues(t *testing.T) {
	SkipConvey("Running tests on 'GetOpenIssues' request", t, func() {
		Convey("GetOpenIssues should return array of issues", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "GET" {
					t.Errorf("Expected Request 'GET', got '%s'", r.Method)
				}
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(issueResponse))

			}))

			defer ts.Close()
			repo, err := GetOpenIssues(ts.URL + "/issues/")

			So(len(repo), ShouldEqual, 1)
			So(err, ShouldBeNil)
		})
	})
}

func TestSetLabel(t *testing.T) {
	Convey("Running tests on 'SetLabel' request", t, func() {

		Convey("SetLabel should return array of labels", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "POST" {
					t.Errorf("Expected Request 'POST', got '%s'", r.Method)
				}

				if r.Header.Get("Authorization") != "token abc" {
					t.Errorf("Authorization error, was '%s'", r.Header.Get("Authorization"))
				}

				if r.URL.EscapedPath() != "/issues/org/phlow-repo/issues/1/labels" {
					t.Errorf("Expected request to '/issues/org/phlow-repo/issues/1/labels', got '%s'", r.URL.EscapedPath())
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(labelResponse))

			}))

			defer ts.Close()
			info := githandler.RemoteInfo{Organisation: "org", Repository: "phlow-repo"}

			labels, err := SetLabel(LabelStatusInProgress, ts.URL+"/issues/", "abc", 1, &info)
			So(len(labels), ShouldEqual, 4)
			So(err, ShouldBeNil)
		})

	})
}

func TestSetAssignee(t *testing.T) {
	Convey("Runnig tests on 'SetAssignee' function", t, func() {
		Convey("SetAssignee should not return error", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "POST" {
					t.Errorf("Expected Request 'POST', got '%s'", r.Method)
				}

				if r.Header.Get("Authorization") != "token abc" {
					t.Errorf("Authorization error, was '%s'", r.Header.Get("Authorization"))
				}

				if r.URL.EscapedPath() != "/issues/org/phlow-repo/issues/1/assignees" {
					t.Errorf("Expected request to '/issues/org/phlow-repo/issues/1/assignees', got '%s'", r.URL.EscapedPath())
				}

				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(labelResponse))

			}))

			defer ts.Close()

			info := githandler.RemoteInfo{Organisation: "org", Repository: "phlow-repo"}
			err := SetAssignee("john markom", ts.URL+"/issues/", "abc", 1, &info)
			So(err, ShouldBeNil)
		})
	})
}

var expected = `{"assignees":["john markom"]}`

var labelResponse = `
[
  {
    "id": 544302811,
    "url": "https://api.github.com/repos/Praqma/phlow-test/labels/Action%20-%20awaiting%20feed-back",
    "name": "Action - awaiting feed-back",
    "color": "6eb82c",
    "default": false
  },
  {
    "id": 545150499,
    "url": "https://api.github.com/repos/Praqma/phlow-test/labels/Label1",
    "name": "Label1",
    "color": "ededed",
    "default": false
  },
  {
    "id": 545150500,
    "url": "https://api.github.com/repos/Praqma/phlow-test/labels/Label2",
    "name": "Label2",
    "color": "ededed",
    "default": false
  },
  {
    "id": 544302897,
    "url": "https://api.github.com/repos/Praqma/phlow-test/labels/Size%202%20-%20medium",
    "name": "Size 2 - medium",
    "color": "208fe5",
    "default": false
  }
]`

var repoResponse = `
{
   "id":79813336,
   "name":"git-phlow",
   "full_name":"Praqma/git-phlow",
   "private":false,
   "html_url":"https://github.com/Praqma/git-phlow",
   "description":"The official repository for the git-phlow extension",
   "language":"Go",
   "has_issues":true,
   "has_downloads":true,
   "has_wiki":true,
   "open_issues_count":28,
   "forks":0,
   "open_issues":28,
   "watchers":0,
   "default_branch":"master",
   "network_count":0,
   "subscribers_count":3
}`

var authResponse = `
{
  "id": 1,
  "url": "https://api.github.com/authorizations/1",
  "scopes": [
    "public_repo"
  ],
  "token": "abcdefgh12345678",
  "token_last_eight": "12345678",
  "hashed_token": "25f94a2a5c7fbaf499c665bc73d67c1c87e496da8985131633ee0a95819db2e8",
  "app": {
    "url": "http://my-github-app.com",
    "name": "my github app",
    "client_id": "abcde12345fghij67890"
  },
  "note": "optional note",
  "note_url": "http://optional/note/url",
  "updated_at": "2011-09-06T20:39:23Z",
  "created_at": "2011-09-06T17:26:27Z",
  "fingerprint": ""
}`

var issueResponse = `[
  {
    "url": "https://api.github.com/repos/Praqma/git-phlow/issues/46",
    "repository_url": "https://api.github.com/repos/Praqma/git-phlow",
    "labels_url": "https://api.github.com/repos/Praqma/git-phlow/issues/46/labels{/name}",
    "comments_url": "https://api.github.com/repos/Praqma/git-phlow/issues/46/comments",
    "events_url": "https://api.github.com/repos/Praqma/git-phlow/issues/46/events",
    "html_url": "https://github.com/Praqma/git-phlow/issues/46",
    "id": 208421587,
    "number": 46,
    "title": "rewrite of repo",
    "user": {
      "login": "groenborg",
      "id": 5576954,
      "avatar_url": "https://avatars.githubusercontent.com/u/5576954?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/groenborg",
      "html_url": "https://github.com/groenborg",
      "followers_url": "https://api.github.com/users/groenborg/followers",
      "following_url": "https://api.github.com/users/groenborg/following{/other_user}",
      "gists_url": "https://api.github.com/users/groenborg/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/groenborg/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/groenborg/subscriptions",
      "organizations_url": "https://api.github.com/users/groenborg/orgs",
      "repos_url": "https://api.github.com/users/groenborg/repos",
      "events_url": "https://api.github.com/users/groenborg/events{/privacy}",
      "received_events_url": "https://api.github.com/users/groenborg/received_events",
      "type": "User",
      "site_admin": false
    },
    "labels": [
      {
        "id": 524293854,
        "url": "https://api.github.com/repos/Praqma/git-phlow/labels/Prio%201%20-%20must%20have",
        "name": "Prio 1 - must have",
        "color": "e83d0f",
        "default": false
      },
      {
        "id": 524293924,
        "url": "https://api.github.com/repos/Praqma/git-phlow/labels/Size%203%20-%20large",
        "name": "Size 3 - large",
        "color": "0052cc",
        "default": false
      },
      {
        "id": 524293956,
        "url": "https://api.github.com/repos/Praqma/git-phlow/labels/Status%20-%20in%20progress",
        "name": "Status - in progress",
        "color": "ededed",
        "default": false
      }
    ],
    "state": "open",
    "locked": false,
    "assignee": {
      "login": "groenborg",
      "id": 5576954,
      "avatar_url": "https://avatars.githubusercontent.com/u/5576954?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/groenborg",
      "html_url": "https://github.com/groenborg",
      "followers_url": "https://api.github.com/users/groenborg/followers",
      "following_url": "https://api.github.com/users/groenborg/following{/other_user}",
      "gists_url": "https://api.github.com/users/groenborg/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/groenborg/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/groenborg/subscriptions",
      "organizations_url": "https://api.github.com/users/groenborg/orgs",
      "repos_url": "https://api.github.com/users/groenborg/repos",
      "events_url": "https://api.github.com/users/groenborg/events{/privacy}",
      "received_events_url": "https://api.github.com/users/groenborg/received_events",
      "type": "User",
      "site_admin": false
    },
    "assignees": [
      {
        "login": "groenborg",
        "gravatar_id": "",
        "url": "https://api.github.com/users/groenborg",
        "html_url": "https://github.com/groenborg",
        "followers_url": "https://api.github.com/users/groenborg/followers",
        "following_url": "https://api.github.com/users/groenborg/following{/other_user}",
        "gists_url": "https://api.github.com/users/groenborg/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/groenborg/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/groenborg/subscriptions",
        "organizations_url": "https://api.github.com/users/groenborg/orgs",
        "repos_url": "https://api.github.com/users/groenborg/repos",
        "events_url": "https://api.github.com/users/groenborg/events{/privacy}",
        "received_events_url": "https://api.github.com/users/groenborg/received_events",
        "type": "User",
        "site_admin": false
      }
    ],
    "milestone": {
      "url": "https://api.github.com/repos/Praqma/git-phlow/milestones/3",
      "html_url": "https://github.com/Praqma/git-phlow/milestone/3",
      "labels_url": "https://api.github.com/repos/Praqma/git-phlow/milestones/3/labels",
      "id": 2309002,
      "number": 3,
      "title": "Implement workon, init wrapup commands ",
      "description": "Milestone for the first three basic commands",
      "open_issues": 10,
      "closed_issues": 8,
      "state": "open",
      "created_at": "2017-02-08T09:59:49Z",
      "updated_at": "2017-02-17T12:01:12Z",
      "due_on": "2017-02-20T08:00:00Z",
      "closed_at": null
    },
    "comments": 0,
    "created_at": "2017-02-17T12:00:10Z",
    "updated_at": "2017-02-20T09:43:00Z",
    "closed_at": null,
    "body": "start over!"
  }
]`
