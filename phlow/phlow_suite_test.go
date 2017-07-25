package phlow_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPhlow(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Phlow Suite")
}
