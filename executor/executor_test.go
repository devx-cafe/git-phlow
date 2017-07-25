package executor_test

import (
	. "github.com/praqma/git-phlow/executor"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"bytes"
	"os/exec"
)

var _ = Describe("Executor", func() {

	Describe("The ExecuteCommand function", func() {

		Context("called with valid command ls", func() {
			It("should not return an error", func() {
				_, err := ExecuteCommand("echo", "hello")

				Ω(err).Should(BeNil())
			})
		})

		Context("called with invalid command", func() {
			It("should fail", func() {
				output, err := ExecuteCommand("lsk", "-lah")
				Ω(output).Should(BeEmpty())
				Ω(err).ShouldNot(BeNil())
			})

		})

	})

	Describe("The ExecPipeCommand function", func() {
		Context("should run", func() {
			It("with 3 commands", func() {
				var buf bytes.Buffer
				err := ExecPipeCommand(&buf,
					exec.Command("echo", "cyan read yellow"),
					exec.Command("grep", "c"),
					exec.Command("sort", "-r"))

				Ω(err).Should(BeNil())
				Ω(buf.String()).ShouldNot(BeEmpty())
			})
		})

		Context("should run", func() {
			It("with 2 commands", func() {
				var buf bytes.Buffer
				err := ExecPipeCommand(&buf,
					exec.Command("echo", "cyan"),
					exec.Command("grep", "c"))

				Ω(err).Should(BeNil())
				Ω(buf.String()).ShouldNot(BeEmpty())
			})
		})

		Context("should run", func() {
			It("with 1 command", func() {
				var buf bytes.Buffer
				err := ExecPipeCommand(&buf, exec.Command("echo", "hello"))

				Ω(err).Should(BeNil())
				Ω(buf.String()).ShouldNot(BeEmpty())
			})
		})

		Context("should fail", func() {
			It("in first function", func() {
				var buf bytes.Buffer

				err := ExecPipeCommand(&buf,
					exec.Command("argh", "blash"),
					exec.Command("grep", "stuff"))

				Ω(err).ShouldNot(BeNil())
			})
		})

		Context("should fail", func() {
			It("in second function", func() {
				var buf bytes.Buffer

				err := ExecPipeCommand(&buf,
					exec.Command("ls", "-lah"),
					exec.Command("jklasd", "stuff"))

				Ω(err).ShouldNot(BeNil())
			})
		})
	})

})
