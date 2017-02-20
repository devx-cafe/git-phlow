package plugins

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUnmarshal(t *testing.T) {
	Convey("Test unmarshal method", t, func() {

		Convey("Unmarshal should return token", func() {
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

var authResponse = `{
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
