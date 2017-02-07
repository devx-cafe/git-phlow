/*
Test package for githubwrapper Branch
*/
package githubwrapper

import (
	"testing"

	"fmt"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDefaultBranch(t *testing.T) {
	Convey("Test function DefaultBranch", t, func() {

		//Convey("Test function 'DefaultBranch' should contain default branch", func() {
		//b := []byte(`"default_branch": "master"`)
		//repo := UnmarshalRepoInfo(b)
		//fmt.Println(repo.DefaultBranch)

		//So(json, ShouldEqual, json)

		//})
	})
}

func TestGetRepoInfo(t *testing.T) {
	Convey("Test function GetRepoInfo", t, func() {

		reponame := "Praqma/git-phlow"

		var res Repo
		res = GetRepoInfo(reponame)

		Convey("Function should return non-empty result for valid public repo", func() {
			fmt.Println("")
			fmt.Printf("%+v\n", res)

			So(res, ShouldNotBeEmpty)
		})
	})
}
