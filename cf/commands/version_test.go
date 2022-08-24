package commands_test

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commandregistry"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/flags"

	testterm "github.com/LukasHeimann/cloudfoundrycli/v8/cf/util/testhelpers/terminal"

	"github.com/LukasHeimann/cloudfoundrycli/v8/cf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("version Command", func() {
	var (
		ui  *testterm.FakeUI
		cmd commandregistry.Command
	)

	BeforeEach(func() {
		ui = &testterm.FakeUI{}

		deps := commandregistry.Dependency{
			UI: ui,
		}

		cmd = &commands.Version{}
		cmd.SetDependency(deps, false)
	})

	Describe("Execute", func() {
		var flagContext flags.FlagContext

		BeforeEach(func() {
			cf.Name = "my-special-cf"
		})

		It("prints the version", func() {
			cmd.Execute(flagContext)

			Expect(ui.Outputs()).To(Equal([]string{
				"my-special-cf version 0.0.0-unknown-version",
			}))
		})
	})
})
