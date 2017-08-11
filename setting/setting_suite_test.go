package setting_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSetting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Setting Suite")
}
