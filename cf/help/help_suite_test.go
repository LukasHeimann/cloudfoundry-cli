package help_test

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commandsloader"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHelp(t *testing.T) {
	RegisterFailHandler(Fail)

	commandsloader.Load()

	RunSpecs(t, "Help Suite")
}
