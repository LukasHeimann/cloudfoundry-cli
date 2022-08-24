package plugin_transition

import (
	"os"

	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/cmd"
	"github.com/LukasHeimann/cloudfoundrycli/v8/command"
	"github.com/LukasHeimann/cloudfoundrycli/v8/util/configv3"
)

func RunPlugin(plugin configv3.Plugin, _ command.UI) error {
	// ugly workaround to maintain v7 api in v7 main
	plugin = configv3.Plugin{}
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
