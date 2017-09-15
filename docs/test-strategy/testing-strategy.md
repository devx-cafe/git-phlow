# Testing strategy for Git Phlow
Looking at the structure of Git Phlow, it is relevant to define what is to be considered:
- unit tests (this one is less necessary)
- integration tests
- component tests 
- API tests

As well as whether or not to treat the 'issues' commands as GUI. 

For a normal project, it is worthwhile to take into consideration the testing pyramid: 

![test pyramid](./testpyramid.png "Testing pyramid")

For Praqma projects, the "manual" testing is only ever to explore new possibilities for more automated testing, however the rest of the pyramid stands. 

For the sake of this document, integration and component have been switched around, as Git Phlow has defined its integration modules as 'components'. 

Going by the bounded contexts of the Git Phlow command line tool, is worthwhile as this defines what top level modules actually exist. 

## Bounded contexts (API level)
Since there is not yet an architectural overview for git phlow, the propposed bounded contexts are *workon*, *upnext*, *cleanup*, *deliver*, *wraup* and *issues*. 

These points are exellent options for API tests, and would be the only place a system test would make any sense as well. This can be done on an integration test level, keeping it in golang, or from shell/powershell to mimic user interaction. 

Depending on what sort of issues has been seen throughout Git Phlow's life cycle, it is better to keep them in the golang test suite if possible to avoid having to tinker with external dependencies (as powershell/bash requires more maintenance presumably).

Maximum 1-2 tests per bounded context, as anything else is likely too expensive to maintain. 

## Layers (Component level)
Looking at it from a layered approach, these commands go through git, jira and github. For git there is another layer, which is the operating system Git Phlow is running on, while for both github and jira there is authentication. 

The tests here are likely to exercise entire component blocks, wrapping functionality of github, jira and git phlow internals respectively. This will include Windows, Linux and OS X as each of these are perceived as a component black box, and should have black box tests (component tests) covering their happy path, as well as catch OS specific problems. 

Components each are likely to have a varying count of tests, but should have at least one test for each bounded context that calls them. Still only happy path testing should be necessary. 

Pay specific attention to the git component, as tests here might be redundant if all git functions do, is call OS components then it is not worth to put tests here. 

## Low-level building blocks (Integration level)
These are some of what we already test, but should include setting up a repository in git or doing OS-centric calls directly to processes. 
The issue here is that traditionally integration tests either exercise database or network calls, and Git Phlow has a different need. 

Maintaining a repository which can be spun up as code in similar fashion should be a task, similarly to how a testing database spun up as code is the norm. This is the recommended approach we give in our code maturity, and we should heed that here. Here it is probably important to make it external, as otherwise it loses its value by being tied to OS. Note that this solution should be the very last thing we do, as it is extremely high maintenance and quite complex to maintain external test dependencies. (Same as for a database). 

As for network calls, Git Phlow instead has OS calls. Having tests that communicate and do every git command for any of the OS is necessary, and should cover most of the problem. There are also several network calls, which can be modelled as usual. 

For Linux, containerizing and using different linux distros makes it extremely easy to reuse and scale integration tests, as it's possible to rerun the same integration test for different platforms. 
Similarly for OS X. With Windows this is not entirely the same, and different windows slaves may need to be created for the used build system, however when Windows Containers are mature the same can be done here. 

None of the above tests should ever be necessary at the higher levels of testing. Furthermore, testing things that are not happy path may be worthwhile here, depending on where bugs showed up in the past, and future. 

An exellent example of this can be found in [gh_test.go](https://github.com/Praqma/git-phlow/blob/master/plugins/gh_test.go) : 
```
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
			GitHub.Auth.URL = ts.URL + "/authorizations"
			token, err := GitHub.Auth.Auth("simon", "password")
			t.Log(err)
			So(token, ShouldEqual, "abcdefgh12345678")
			So(err, ShouldBeNil)
		})
	})
}
```

As well as [branch_test.go](https://github.com/Praqma/git-phlow/blob/master/githandler/branch_test.go)

```
func TestBranch(t *testing.T) {
	Convey("Running tests on 'Branch' function", t, func() {

		testfixture.CreateTestRepository(t, false)

		Convey("branch should return List of branches", func() {
			info, err := Branch()
			So(len(info.List), ShouldEqual, 11)
			So(err, ShouldBeNil)
		})

		Convey("branch should return Current branch", func() {
			info, err := Branch()
			So(info.Current, ShouldEqual, "master")
			So(err, ShouldBeNil)
		})

		testfixture.RemoveTestRepository(t)

	})
}
```
These are marked as integration tests since neither of them execute on a singular method, but need to authenticate against github (network call) and create a repository to execute on (system call) respectively. 


## Methods (Unit test level)
Every method should have a unit test, as it's a straight forward project to run TDD on. However as with most prototypes this is probably not the case, and thus it is necessary to determine which parts of the system are still exploratory. 
Writing a lot of unit tests for existing methods, which are still being restructured is a waste of productivity. 

Good examples of unit tests would be this method in [plugin_test.go](https://github.com/Praqma/git-phlow/blob/master/plugins/plugin_test.go)

```
Describe("Branch name from issue", func() {
		testsToRun := [7]testCase{
			{issue: 12, branchName: "work on iss", expected: "12-work-on-iss", casedesc: "Test replaces whitespaces with dash '-'"},
			{issue: 45, branchName: "Case SENsitivity", expected: "45-case-sensitivity", casedesc: "Test converts charecters to lowercase"},
			{issue: 15, branchName: ".branch name", expected: "15-branch-name", casedesc: "Test removes . prefix"},
			{issue: 220, branchName: "^^..:~:name", expected: "220-name", casedesc: "removes ASCII control characters"},
			{issue: 2735, branchName: "name/", expected: "2735-name", casedesc: "test removes end / "},
			{issue: 234567, branchName: ".NAME.is\"dotted", expected: "234567-name-is-dotted", casedesc: "test removes backslash"},
			{issue: 672, branchName: "add big data blog /", expected: "672-add-big-data-blog", casedesc: "remove forward slash"},
		}

		Context("Names should follow format rules", func() {

			for _, currentTest := range testsToRun {

				It(currentTest.casedesc, func() {
					actualName := BranchNameFromIssue(strconv.Itoa(currentTest.issue), currentTest.branchName)

					Ω(actualName).Should(Equal(currentTest.expected))
				})
			}
		})
	})
```

This is a unit test, as no calls are made other than the single method in question (which compares the name to format rules)


## Handling "legacy" code
Legacy code for Git Phlow is from here defined as : 

A method, code-block, component or API call not covered by an automated test. 

Allowing legacy code in the CI or CD is perfectly fine, as long as pipeline breaks in the case where someone contributes and lowers code test coverage. Also keep in mind that tests have diminishing returns after 80% coverage (read: not worth doing in most cases). 






