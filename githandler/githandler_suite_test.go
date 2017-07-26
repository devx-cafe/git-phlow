package githandler_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGithandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Githandler Suite")
}
