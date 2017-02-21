package plugins

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"strings"
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
				w.WriteHeader(http.StatusOK)
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
	Convey("Runnign tests on 'GetDefaultBranch' request", t, func() {
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
	Convey("Running tests on 'GetOpenIssues' request", t, func() {
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

			for _, v := range repo {
				t.Log(v)
			}

			So(err, ShouldBeNil)
		})
	})
}

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
      "creator": {
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
  },
  {
    "url": "https://api.github.com/repos/Praqma/git-phlow/issues/45",
    "repository_url": "https://api.github.com/repos/Praqma/git-phlow",
    "labels_url": "https://api.github.com/repos/Praqma/git-phlow/issues/45/labels{/name}",
    "comments_url": "https://api.github.com/repos/Praqma/git-phlow/issues/45/comments",
    "events_url": "https://api.github.com/repos/Praqma/git-phlow/issues/45/events",
    "html_url": "https://github.com/Praqma/git-phlow/issues/45",
    "id": 207844777,
    "number": 45,
    "title": "Support setting default branch in git config",
    "user": {
      "login": "JKrag",
      "id": 1712950,
      "avatar_url": "https://avatars.githubusercontent.com/u/1712950?v=3",
      "gravatar_id": "",
      "url": "https://api.github.com/users/JKrag",
      "html_url": "https://github.com/JKrag",
      "followers_url": "https://api.github.com/users/JKrag/followers",
      "following_url": "https://api.github.com/users/JKrag/following{/other_user}",
      "gists_url": "https://api.github.com/users/JKrag/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/JKrag/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/JKrag/subscriptions",
      "organizations_url": "https://api.github.com/users/JKrag/orgs",
      "repos_url": "https://api.github.com/users/JKrag/repos",
      "events_url": "https://api.github.com/users/JKrag/events{/privacy}",
      "received_events_url": "https://api.github.com/users/JKrag/received_events",
      "type": "User",
      "site_admin": false
    },
    "labels": [

    ],
    "state": "open",
    "locked": false,
    "assignee": null,
    "assignees": [

    ],
    "milestone": null,
    "comments": 0,
    "created_at": "2017-02-15T15:56:15Z",
    "updated_at": "2017-02-15T15:56:15Z",
    "closed_at": null,
    "body": "It should be possible to set the name of the integration branch in the normal git config hierarchy.\r\nOnly if the value is not set, it should default to asking GitHub for its \"default branch\".\r\n\r\nThis feature is useful for situations where you are working on e.g. a maintenance branch. \r\nIt will also be useful for future support of other Git servers that do not have a \"default-branch\" API endpoint.\r\n\r\nThis issue probably depends on or relates to #15 "
  },
  {
    "url": "https://api.github.com/repos/Praqma/git-phlow/issues/42",
    "repository_url": "https://api.github.com/repos/Praqma/git-phlow",
    "labels_url": "https://api.github.com/repos/Praqma/git-phlow/issues/42/labels{/name}",
    "comments_url": "https://api.github.com/repos/Praqma/git-phlow/issues/42/comments",
    "events_url": "https://api.github.com/repos/Praqma/git-phlow/issues/42/events",
    "html_url": "https://github.com/Praqma/git-phlow/issues/42",
    "id": 207296792,
    "number": 42,
    "title": "Deliver - Add gitwrapper push",
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
        "id": 524293923,
        "url": "https://api.github.com/repos/Praqma/git-phlow/labels/Size%202%20-%20medium",
        "name": "Size 2 - medium",
        "color": "208fe5",
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
      "creator": {
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
      "open_issues": 10,
      "closed_issues": 8,
      "state": "open",
      "created_at": "2017-02-08T09:59:49Z",
      "updated_at": "2017-02-17T12:01:12Z",
      "due_on": "2017-02-20T08:00:00Z",
      "closed_at": null
    },
    "comments": 0,
    "created_at": "2017-02-13T18:12:49Z",
    "updated_at": "2017-02-13T18:12:49Z",
    "closed_at": null,
    "body": "Add push gitwrapper "
  },
  {
    "url": "https://api.github.com/repos/Praqma/git-phlow/issues/41",
    "repository_url": "https://api.github.com/repos/Praqma/git-phlow",
    "labels_url": "https://api.github.com/repos/Praqma/git-phlow/issues/41/labels{/name}",
    "comments_url": "https://api.github.com/repos/Praqma/git-phlow/issues/41/comments",
    "events_url": "https://api.github.com/repos/Praqma/git-phlow/issues/41/events",
    "html_url": "https://github.com/Praqma/git-phlow/issues/41",
    "id": 207294158,
    "number": 41,
    "title": "Colored printers",
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
        "id": 524293863,
        "url": "https://api.github.com/repos/Praqma/git-phlow/labels/Prio%203%20-%20could%20have",
        "name": "Prio 3 - could have",
        "color": "e8850f",
        "default": false
      },
      {
        "id": 524293915,
        "url": "https://api.github.com/repos/Praqma/git-phlow/labels/Size%201%20-%20small",
        "name": "Size 1 - small",
        "color": "20b4e5",
        "default": false
      },
      {
        "id": 524293956,
        "url": "https://api.github.com/repos/Praqma/git-phlow/labels/Status%20-%20in%20progress",
        "name": "Status - in progress",
        "color": "ededed",
        "default": false
      },
      {
        "id": 524293957,
        "url": "https://api.github.com/repos/Praqma/git-phlow/labels/Status%20-%20up%20next",
        "name": "Status - up next",
        "color": "eeeeee",
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
      "creator": {
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
      "open_issues": 10,
      "closed_issues": 8,
      "state": "open",
      "created_at": "2017-02-08T09:59:49Z",
      "updated_at": "2017-02-17T12:01:12Z",
      "due_on": "2017-02-20T08:00:00Z",
      "closed_at": null
    },
    "comments": 0,
    "created_at": "2017-02-13T18:01:47Z",
    "updated_at": "2017-02-13T18:15:43Z",
    "closed_at": null,
    "body": "Print colored printers in terminal"
  },
  {
    "url": "https://api.github.com/repos/Praqma/git-phlow/issues/40",
    "repository_url": "https://api.github.com/repos/Praqma/git-phlow",
    "labels_url": "https://api.github.com/repos/Praqma/git-phlow/issues/40/labels{/name}",
    "comments_url": "https://api.github.com/repos/Praqma/git-phlow/issues/40/comments",
    "events_url": "https://api.github.com/repos/Praqma/git-phlow/issues/40/events",
    "html_url": "https://github.com/Praqma/git-phlow/issues/40",
    "id": 206199667,
    "number": 40,
    "title": "User story - create and work on issue",
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
        "id": 524293863,
        "url": "https://api.github.com/repos/Praqma/git-phlow/labels/Prio%203%20-%20could%20have",
        "name": "Prio 3 - could have",
        "color": "e8850f",
        "default": false
      }
    ],
    "state": "open",
    "locked": false,
    "assignee": null,
    "assignees": [

    ],
    "milestone": null,
    "comments": 0,
    "created_at": "2017-02-08T13:35:49Z",
    "updated_at": "2017-02-08T13:36:07Z",
    "closed_at": null,
    "body": "As a:\r\ndeveloper\r\n\r\nI dream of:\r\nbeing able to create an issue and start working on it, in a single command\r\n\r\nSo that:\r\nI can do quick fixes without having to separately create an issue\r\n\r\nsuggested by @JKrag "
  }
]`
