package phlow

import (
	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/testfixture"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestClean(t *testing.T) {

	Convey("Runnign tests on 'Clean' function", t, func() {

		testfixture.CreateTestRepository(t, false)

		Convey("Testing output of local clean function", func() {
			options.GlobalFlagLocal = true
			Clean("origin")

		})

		testfixture.RemoveTestRepository(t)
	})

}

func TestCleanRemote(t *testing.T) {

	Convey("Runnign tests on 'Clean' function", t, func() {

		testfixture.CreateTestRepository(t, false)

		Convey("Testing output of clean function", func() {
			options.GlobalFlagLocal = false
			Clean("origin")

		})

		testfixture.RemoveTestRepository(t)
	})

}
