package plugins

import (
	"net/http"
	"net/http/httptest"
	"testing"

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
			GitHub.base = ts.URL
			token, err := GitHub.Auth("simon", "password")
			t.Log(err)
			So(token, ShouldEqual, "abcdefgh12345678")
			So(err, ShouldBeNil)
		})
	})
}

func TestCheckAuth(t *testing.T) {
	Convey("Running tests on 'CheckAuth' request", t, func() {

		Convey("CheckAuth should return true on status 200", func() {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(checkAuthResponse))
			}))

			defer ts.Close()
			GitHub.base = ts.URL
			isAuthorized, err := GitHub.CheckAuth()
			So(isAuthorized, ShouldBeTrue)
			So(err, ShouldBeNil)

		})
	})
}

func TestCreatePermissions(t *testing.T) {
	Convey("Running tests on 'createGHPermissions' function", t, func() {
		Convey("should return json permissions as string", func() {
			str, err := createGHPermissions()

			t.Log(str)
			So(str, ShouldContainSubstring, "repo")
			So(err, ShouldBeNil)

		})
	})
}

func TestGetDefaultBranch(t *testing.T) {
	Convey("Runnign tests on 'GetDefaultBranch' request", t, func() {
		Convey("GetDefaultBranch should return master", func() {
			var repo = "i-am-a-REPO"
			var org = "this-org"
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "GET" {
					t.Errorf("Expected Request 'GET', got '%s'", r.Method)
				}
				if r.URL.EscapedPath() == "repos/"+org+"/"+repo {
					t.Errorf("Expected got '%s'", r.URL.EscapedPath())
				}

				t.Log(r.URL.String())
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(repoResponse))

			}))
			defer ts.Close()
			GitHub.base = ts.URL
			GitHub.org = org
			GitHub.repo = repo
			def, err := GitHub.Default()
			t.Log(def)
			So(def, ShouldEqual, "master")
			So(err, ShouldBeNil)
		})
	})
}

func TestGetIssues(t *testing.T) {
	Convey("Running tests on 'GetIssues' request", t, func() {
		Convey("GetIssues should return array of issues", func() {
			var org = "org"
			var repo = "some-repo-name"

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "GET" {
					t.Errorf("Expected Request 'GET', got '%s'", r.Method)
				}
				t.Log(r.URL.EscapedPath())
				if r.URL.EscapedPath() != "/repos/"+org+"/"+repo+"/issues" {
					t.Errorf("expected %s but got %s", GitHub.issueURL, r.URL.EscapedPath())
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(issueResponse))

			}))
			defer ts.Close()
			GitHub.org = org
			GitHub.repo = repo
			GitHub.base = ts.URL

			issues, err := GitHub.GetIssues()
			So(issues[0].Assignees[0].Login, ShouldEqual, "groenborg")
			So(issues, ShouldHaveLength, 1)
			So(err, ShouldBeNil)
		})
	})
}

func TestSetLabel(t *testing.T) {
	Convey("Running tests on 'SetLabel' request", t, func() {

		Convey("SetLabel should return array of labels", func() {
			var org = "org"
			var repo = "some-repo-name"

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "POST" {
					t.Errorf("Expected Request 'POST', got '%s'", r.Method)
				}

				if r.Header.Get("Authorization") != "token abc" {
					t.Errorf("Authorization error, was '%s'", r.Header.Get("Authorization"))
				}

				if r.URL.EscapedPath() != "/repos/"+org+"/"+repo+"/issues/1/labels" {
					t.Errorf("Expected request to '/issues/org/phlow-repo/issues/1/labels', got '%s'", r.URL.EscapedPath())
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(labelResponse))
			}))

			defer ts.Close()
			GitHub.base = ts.URL
			GitHub.token = "abc"
			GitHub.org = org
			GitHub.repo = repo
			labels, err := GitHub.SetLabel("Status - in progress", 1)
			t.Log(err)
			So(labels, ShouldHaveLength, 4)
			So(err, ShouldBeNil)
		})

	})
}

func TestSetAssignee(t *testing.T) {
	Convey("Runnig tests on 'SetAssignee' function", t, func() {
		Convey("SetAssignee should not return error", func() {
			var org = "org"
			var repo = "some-repo-name"
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "POST" {
					t.Errorf("Expected Request 'POST', got '%s'", r.Method)
				}

				if r.Header.Get("Authorization") != "token abc" {
					t.Errorf("Authorization error, was '%s'", r.Header.Get("Authorization"))
				}

				if r.URL.EscapedPath() != "/repos/"+org+"/"+repo+"/issues/1/assignees" {
					t.Errorf("Expected request to 'issues/"+org+"/"+repo+"/issues/1/assignees', got '%s'", r.URL.EscapedPath())
				}

				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(labelResponse))

			}))
			defer ts.Close()

			GitHub.base = ts.URL
			GitHub.org = org
			GitHub.repo = repo
			GitHub.token = "abc"
			err := GitHub.SetAssignee("john markom", 1)
			So(err, ShouldBeNil)
		})
	})
}

var expected = `{"assignees":["john markom"]}`

var labelResponse = `
[
  {
    "id": 544302811,
    "url": "https://api.gh.com/repos/Praqma/phlow-test/labels/Action%20-%20awaiting%20feed-back",
    "name": "Action - awaiting feed-back",
    "color": "6eb82c",
    "default": false
  },
  {
    "id": 545150499,
    "url": "https://api.gh.com/repos/Praqma/phlow-test/labels/Label1",
    "name": "Label1",
    "color": "ededed",
    "default": false
  },
  {
    "id": 545150500,
    "url": "https://api.gh.com/repos/Praqma/phlow-test/labels/Label2",
    "name": "Label2",
    "color": "ededed",
    "default": false
  },
  {
    "id": 544302897,
    "url": "https://api.gh.com/repos/Praqma/phlow-test/labels/Size%202%20-%20medium",
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
   "html_url":"https://gh.com/Praqma/git-phlow",
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
  "url": "https://api.gh.com/authorizations/1",
  "scopes": [
    "public_repo"
  ],
  "token": "abcdefgh12345678",
  "token_last_eight": "12345678",
  "hashed_token": "25f94a2a5c7fbaf499c665bc73d67c1c87e496da8985131633ee0a95819db2e8",
  "app": {
    "url": "http://my-gh-app.com",
    "name": "my gh app",
    "client_id": "abcde12345fghij67890"
  },
  "note": "optional note",
  "note_url": "http://optional/note/url",
  "updated_at": "2011-09-06T20:39:23Z",
  "created_at": "2011-09-06T17:26:27Z",
  "fingerprint": ""
}`

var checkAuthResponse = `
{
  "id": 75939914,
  "name": "pj",
  "full_name": "Andrey9kin/pj",
  "owner": {
    "login": "Andrey9kin",
    "id": 1671664,
    "avatar_url": "https://avatars3.githubusercontent.com/u/1671664?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/Andrey9kin",
    "html_url": "https://github.com/Andrey9kin",
    "followers_url": "https://api.github.com/users/Andrey9kin/followers",
    "following_url": "https://api.github.com/users/Andrey9kin/following{/other_user}",
    "gists_url": "https://api.github.com/users/Andrey9kin/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/Andrey9kin/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/Andrey9kin/subscriptions",
    "organizations_url": "https://api.github.com/users/Andrey9kin/orgs",
    "repos_url": "https://api.github.com/users/Andrey9kin/repos",
    "events_url": "https://api.github.com/users/Andrey9kin/events{/privacy}",
    "received_events_url": "https://api.github.com/users/Andrey9kin/received_events",
    "type": "User",
    "site_admin": false
  },
  "private": false,
  "html_url": "https://github.com/Andrey9kin/pj",
  "description": "Jenkins voice message generator (p - polly, j - jenkins)",
  "fork": false,
  "url": "https://api.github.com/repos/Andrey9kin/pj",
  "forks_url": "https://api.github.com/repos/Andrey9kin/pj/forks",
  "keys_url": "https://api.github.com/repos/Andrey9kin/pj/keys{/key_id}",
  "collaborators_url": "https://api.github.com/repos/Andrey9kin/pj/collaborators{/collaborator}",
  "teams_url": "https://api.github.com/repos/Andrey9kin/pj/teams",
  "hooks_url": "https://api.github.com/repos/Andrey9kin/pj/hooks",
  "issue_events_url": "https://api.github.com/repos/Andrey9kin/pj/issues/events{/number}",
  "events_url": "https://api.github.com/repos/Andrey9kin/pj/events",
  "assignees_url": "https://api.github.com/repos/Andrey9kin/pj/assignees{/user}",
  "branches_url": "https://api.github.com/repos/Andrey9kin/pj/branches{/branch}",
  "tags_url": "https://api.github.com/repos/Andrey9kin/pj/tags",
  "blobs_url": "https://api.github.com/repos/Andrey9kin/pj/git/blobs{/sha}",
  "git_tags_url": "https://api.github.com/repos/Andrey9kin/pj/git/tags{/sha}",
  "git_refs_url": "https://api.github.com/repos/Andrey9kin/pj/git/refs{/sha}",
  "trees_url": "https://api.github.com/repos/Andrey9kin/pj/git/trees{/sha}",
  "statuses_url": "https://api.github.com/repos/Andrey9kin/pj/statuses/{sha}",
  "languages_url": "https://api.github.com/repos/Andrey9kin/pj/languages",
  "stargazers_url": "https://api.github.com/repos/Andrey9kin/pj/stargazers",
  "contributors_url": "https://api.github.com/repos/Andrey9kin/pj/contributors",
  "subscribers_url": "https://api.github.com/repos/Andrey9kin/pj/subscribers",
  "subscription_url": "https://api.github.com/repos/Andrey9kin/pj/subscription",
  "commits_url": "https://api.github.com/repos/Andrey9kin/pj/commits{/sha}",
  "git_commits_url": "https://api.github.com/repos/Andrey9kin/pj/git/commits{/sha}",
  "comments_url": "https://api.github.com/repos/Andrey9kin/pj/comments{/number}",
  "issue_comment_url": "https://api.github.com/repos/Andrey9kin/pj/issues/comments{/number}",
  "contents_url": "https://api.github.com/repos/Andrey9kin/pj/contents/{+path}",
  "compare_url": "https://api.github.com/repos/Andrey9kin/pj/compare/{base}...{head}",
  "merges_url": "https://api.github.com/repos/Andrey9kin/pj/merges",
  "archive_url": "https://api.github.com/repos/Andrey9kin/pj/{archive_format}{/ref}",
  "downloads_url": "https://api.github.com/repos/Andrey9kin/pj/downloads",
  "issues_url": "https://api.github.com/repos/Andrey9kin/pj/issues{/number}",
  "pulls_url": "https://api.github.com/repos/Andrey9kin/pj/pulls{/number}",
  "milestones_url": "https://api.github.com/repos/Andrey9kin/pj/milestones{/number}",
  "notifications_url": "https://api.github.com/repos/Andrey9kin/pj/notifications{?since,all,participating}",
  "labels_url": "https://api.github.com/repos/Andrey9kin/pj/labels{/name}",
  "releases_url": "https://api.github.com/repos/Andrey9kin/pj/releases{/id}",
  "deployments_url": "https://api.github.com/repos/Andrey9kin/pj/deployments",
  "created_at": "2016-12-08T13:25:04Z",
  "updated_at": "2016-12-09T10:46:29Z",
  "pushed_at": "2016-12-09T11:24:52Z",
  "git_url": "git://github.com/Andrey9kin/pj.git",
  "ssh_url": "git@github.com:Andrey9kin/pj.git",
  "clone_url": "https://github.com/Andrey9kin/pj.git",
  "svn_url": "https://github.com/Andrey9kin/pj",
  "homepage": "",
  "size": 17,
  "stargazers_count": 0,
  "watchers_count": 0,
  "language": "Python",
  "has_issues": true,
  "has_projects": true,
  "has_downloads": true,
  "has_wiki": true,
  "has_pages": false,
  "forks_count": 0,
  "mirror_url": null,
  "open_issues_count": 0,
  "forks": 0,
  "open_issues": 0,
  "watchers": 0,
  "default_branch": "master",
  "permissions": {
    "admin": false,
    "push": true,
    "pull": true
  }
}
`

var issueResponse = `[
  {
    "url": "https://api.gh.com/repos/Praqma/git-phlow/issues/46",
    "repository_url": "https://api.gh.com/repos/Praqma/git-phlow",
    "labels_url": "https://api.gh.com/repos/Praqma/git-phlow/issues/46/labels{/name}",
    "comments_url": "https://api.gh.com/repos/Praqma/git-phlow/issues/46/comments",
    "events_url": "https://api.gh.com/repos/Praqma/git-phlow/issues/46/events",
    "html_url": "https://gh.com/Praqma/git-phlow/issues/46",
    "id": 208421587,
    "number": 46,
    "title": "rewrite of repo",
    "user": {
      "login": "groenborg",
      "id": 5576954,
      "avatar_url": "https://avatars.githubusercontent.com/u/5576954?v=3",
      "gravatar_id": "",
      "url": "https://api.gh.com/users/groenborg",
      "html_url": "https://gh.com/groenborg",
      "followers_url": "https://api.gh.com/users/groenborg/followers",
      "following_url": "https://api.gh.com/users/groenborg/following{/other_user}",
      "gists_url": "https://api.gh.com/users/groenborg/gists{/gist_id}",
      "starred_url": "https://api.gh.com/users/groenborg/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.gh.com/users/groenborg/subscriptions",
      "organizations_url": "https://api.gh.com/users/groenborg/orgs",
      "repos_url": "https://api.gh.com/users/groenborg/repos",
      "events_url": "https://api.gh.com/users/groenborg/events{/privacy}",
      "received_events_url": "https://api.gh.com/users/groenborg/received_events",
      "type": "User",
      "site_admin": false
    },
    "labels": [
      {
        "id": 524293854,
        "url": "https://api.gh.com/repos/Praqma/git-phlow/labels/Prio%201%20-%20must%20have",
        "name": "Prio 1 - must have",
        "color": "e83d0f",
        "default": false
      },
      {
        "id": 524293924,
        "url": "https://api.gh.com/repos/Praqma/git-phlow/labels/Size%203%20-%20large",
        "name": "Size 3 - large",
        "color": "0052cc",
        "default": false
      },
      {
        "id": 524293956,
        "url": "https://api.gh.com/repos/Praqma/git-phlow/labels/Status%20-%20in%20progress",
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
      "url": "https://api.gh.com/users/groenborg",
      "html_url": "https://gh.com/groenborg",
      "followers_url": "https://api.gh.com/users/groenborg/followers",
      "following_url": "https://api.gh.com/users/groenborg/following{/other_user}",
      "gists_url": "https://api.gh.com/users/groenborg/gists{/gist_id}",
      "starred_url": "https://api.gh.com/users/groenborg/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.gh.com/users/groenborg/subscriptions",
      "organizations_url": "https://api.gh.com/users/groenborg/orgs",
      "repos_url": "https://api.gh.com/users/groenborg/repos",
      "events_url": "https://api.gh.com/users/groenborg/events{/privacy}",
      "received_events_url": "https://api.gh.com/users/groenborg/received_events",
      "type": "User",
      "site_admin": false
    },
    "assignees": [
      {
        "login": "groenborg",
        "gravatar_id": "",
        "url": "https://api.gh.com/users/groenborg",
        "html_url": "https://gh.com/groenborg",
        "followers_url": "https://api.gh.com/users/groenborg/followers",
        "following_url": "https://api.gh.com/users/groenborg/following{/other_user}",
        "gists_url": "https://api.gh.com/users/groenborg/gists{/gist_id}",
        "starred_url": "https://api.gh.com/users/groenborg/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.gh.com/users/groenborg/subscriptions",
        "organizations_url": "https://api.gh.com/users/groenborg/orgs",
        "repos_url": "https://api.gh.com/users/groenborg/repos",
        "events_url": "https://api.gh.com/users/groenborg/events{/privacy}",
        "received_events_url": "https://api.gh.com/users/groenborg/received_events",
        "type": "User",
        "site_admin": false
      }
    ],
    "milestone": {
      "url": "https://api.gh.com/repos/Praqma/git-phlow/milestones/3",
      "html_url": "https://gh.com/Praqma/git-phlow/milestone/3",
      "labels_url": "https://api.gh.com/repos/Praqma/git-phlow/milestones/3/labels",
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
