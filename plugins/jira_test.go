package plugins_test

import (
	. "github.com/praqma/git-phlow/plugins"

	. "github.com/onsi/ginkgo"
	"net/http/httptest"
	"net/http"
	. "github.com/onsi/gomega"
)

var _ = Describe("Jira", func() {

	Describe("Testing GetJiraIssue", func() {

		Context("call to a valid url", func() {

			It("should return a jira issue object", func() {

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					Ω(r.Method).Should(Equal("GET"))
					Ω(r.Header.Get("Content-Type")).Should(Equal("application/json"))
					Ω(r.URL.EscapedPath()).Should(Equal("/rest/api/latest/issue/TIS-41"))

					w.WriteHeader(http.StatusOK)
					w.Write([]byte(jiraIssueResponse))

				}))
				defer ts.Close()

				issue, err := GetJiraIssue(ts.URL, "TIS-41", "user", "pass")
				Ω(issue.Fields.Status.Name).Should(Equal("open"))
				Ω(err).Should(BeNil())
			})
		})

		Context("call to a invalid url", func() {

			It("send no response and should error", func() {
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					Ω(r.Method).Should(Equal("GET"))
					Ω(r.Header.Get("Content-Type")).Should(Equal("application/json"))

				}))
				defer ts.Close()

				_, err := GetJiraIssue(ts.URL, "TIS-41", "user", "pass")
				Ω(err).ShouldNot(BeNil())
			})
		})

	})

	Describe("Testing GetTransitions", func() {

		Context("with a valid url", func() {

			It("Should return transitions object", func() {

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					Ω(r.Method).Should(Equal("GET"))
					Ω(r.Header.Get("Content-Type")).Should(Equal("application/json"))
					Ω(r.URL.EscapedPath()).Should(Equal("/rest/api/latest/issue/TIS-41/transitions"))

					w.WriteHeader(http.StatusOK)
					w.Write([]byte(jiraTransitionResponse))

				}))
				defer ts.Close()

				transitions, err := GetTransitions(ts.URL, "TIS-41", "user", "pass")
				Ω(transitions.Transitions[0].To.StatusCategory.Name).Should(Equal("In Progress"))
				Ω(err).Should(BeNil())
			})

		})

	})

	Describe("Testing DoTransition", func() {

		Context("Transition with a valid transition ID", func() {

			//Consult API documentation for status codes - https://docs.atlassian.com/jira/REST/cloud/#api/2/issue-doTransition
			It("Should return status code 204 ", func() {
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					Ω(r.Method).Should(Equal("POST"))
					Ω(r.Header.Get("Content-Type")).Should(Equal("application/json"))
					Ω(r.URL.EscapedPath()).Should(Equal("/rest/api/latest/issue/TIS-41/transitions"))

					w.WriteHeader(204)
					w.Write([]byte(jiraTransitionResponse))

				}))
				defer ts.Close()

				err := DoTransition(ts.URL, "TIS-41", "user", "pass", "11")
				Ω(err).Should(BeNil())
			})

		})

		Context("Transition with an invalid transition ID", func() {

			//Consult API documentation for status codes - https://docs.atlassian.com/jira/REST/cloud/#api/2/issue-doTransition
			It("Should return status code 400 ", func() {
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					Ω(r.Method).Should(Equal("POST"))
					Ω(r.Header.Get("Content-Type")).Should(Equal("application/json"))
					Ω(r.URL.EscapedPath()).Should(Equal("/rest/api/latest/issue/TIS-41/transitions"))

					w.WriteHeader(400)
					w.Write([]byte("error"))

				}))
				defer ts.Close()

				err := DoTransition(ts.URL, "TIS-41", "user", "pass", "111")
				Ω(err).ShouldNot(BeNil())
			})

		})

	})

	Describe("Testing assign user", func() {

		Context("Assign with a valid user", func() {
			It("Should return status code 204 ", func() {
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					Ω(r.Method).Should(Equal("PUT"))
					Ω(r.Header.Get("Content-Type")).Should(Equal("application/json"))
					Ω(r.URL.EscapedPath()).Should(Equal("/rest/api/latest/issue/TIS-41/assignee"))

					w.WriteHeader(204)
					w.Write([]byte(jiraTransitionResponse))

				}))
				defer ts.Close()

				err := AssignUser(ts.URL, "TIS-41", "admin", "pass")
				Ω(err).Should(BeNil())
			})
		})
	})
})
