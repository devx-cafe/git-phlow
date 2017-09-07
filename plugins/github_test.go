package plugins_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http/httptest"
	"net/http"
	"github.com/praqma/git-phlow/plugins"
)

var _ = Describe("Github", func() {

	Describe("Testing SetAssigneeGitHub ", func() {
		It("SetAssignee should not return error", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				Ω(r.Method).Should(Equal("POST"))
				Ω(r.URL.RawQuery).Should(Equal("access_token=abc"))
				Ω(r.URL.EscapedPath()).Should(Equal("/repos/org/repo/issues/10/assignees"))

				w.WriteHeader(http.StatusCreated)

			}))
			defer ts.Close()

			err := plugins.SetAssigneeGitHub(ts.URL, "org", "repo", "abc", "10", "john")
			Ω(err).Should(BeNil())
		})
	})

	Describe("Testing SetLabelGitHub", func() {

		It("SetLabel should return array of labels", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				Ω(r.Method).Should(Equal("POST"))
				Ω(r.URL.RawQuery).Should(Equal("access_token=abc"))
				Ω(r.URL.EscapedPath()).Should(Equal("/repos/org/repo/issues/10/labels"))

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(labelResponse))
			}))

			defer ts.Close()
			labels, err := plugins.SetLabelGitHub(ts.URL, "org", "repo", "abc", "Status - in progress", "10")
			Ω(labels).Should(HaveLen(4))
			Ω(err).Should(BeNil())
		})

	})

	Describe("Testing GetIssueGithub", func() {

		It("Should return a single issue", func() {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				Ω(r.Method).Should(Equal("GET"))
				Ω(r.URL.RawQuery).Should(Equal("access_token=abc"))
				Ω(r.URL.EscapedPath()).Should(Equal("/repos/org/repo/issues/1347"))

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(gitHubIssueResponse))
			}))

			defer ts.Close()
			issue, err := plugins.GetIssueGitHub(ts.URL, "org", "repo", "1347", "abc")
			Ω(issue.Number).Should(Equal(1347))
			Ω(err).Should(BeNil())
		})

	})

})
