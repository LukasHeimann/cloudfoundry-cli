package commands_test

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/i18n"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/util/testhelpers/configuration"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"

	"testing"
)

func TestCommands(t *testing.T) {
	config := configuration.NewRepositoryWithDefaults()
	i18n.T = i18n.Init(config)

	_ = commands.API{}

	RegisterFailHandler(Fail)
	RunSpecs(t, "Commands Suite")
}

var _ = BeforeEach(func() {
	log.SetLevel(log.PanicLevel)
})

type passingRequirement struct {
	Name string
}

func (r passingRequirement) Execute() error {
	return nil
}
