package gitwrapper

import (
	. "github.com/smartystreets/goconvey/convey"

	"github.com/praqma/git-phlow/testfixture"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckout(t *testing.T) {

	Convey("Checkout Test", t, func() {

		testfixture.SetupTestRepo()

		Convey("Checkout other branch should be possible", func() {
			assert.Fail(t, "not yet implemented")
		})

		Convey("Checkout current branch should not result in error", func() {
			assert.Fail(t, "not yet implemented")
		})

		Convey("Checkout nonexisting branch should fail", func() {
			assert.Fail(t, "not yet implemented")
		})

		testfixture.TearDownTestRepo()

	})
}
