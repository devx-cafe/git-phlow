package phlow_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/praqma/git-phlow/phlow"
	. "github.com/onsi/gomega"
	"bytes"
)

var _ = Describe("Auth", func() {

	Describe("reading input from user", func() {
		Context("in osx and windows", func() {

			It("should read line breaks on mac", func() {
				ACTUAL := phlow.ReadInput("", bytes.NewReader([]byte("username\n")))
				EXPECTED := "username"

				Ω(ACTUAL).Should(Equal(EXPECTED))
			})

			It("should read line breaks on windows", func() {
				ACTUAL := phlow.ReadInput("", bytes.NewReader([]byte("username\n\r")))
				EXPECTED := "username"

				Ω(ACTUAL).Should(Equal(EXPECTED))
			})
		})
	})
})
