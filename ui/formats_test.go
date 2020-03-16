package ui_test

import (
	. "github.com/code-cafe/git-phlow/ui"

	"github.com/code-cafe/git-phlow/options"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Formats", func() {

	Describe("The no-color Function", func() {

		Context("called with no-color = true", func() {
			It("should not return ANSI", func() {
				options.GlobalFlagNoColor = true
				str := Format.MileStone("sdsdsd")

				Expect(str).ShouldNot(ContainSubstring("\u001b["))
			})
		})
	})
})
