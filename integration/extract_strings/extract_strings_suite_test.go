package extract_strings_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/krishicks/i18n4go/integration/test_helpers"
	"github.com/onsi/gomega/gexec"

	"testing"
)

func TestExtractStrings(t *testing.T) {
	BeforeSuite(test_helpers.BuildExecutable)

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	RegisterFailHandler(Fail)
	RunSpecs(t, "ExtractStrings Suite")
}
