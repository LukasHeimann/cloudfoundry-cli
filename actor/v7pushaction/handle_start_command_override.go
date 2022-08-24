package v7pushaction

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/command/translatableerror"
	"github.com/LukasHeimann/cloudfoundrycli/v8/util/manifestparser"
)

func HandleStartCommandOverride(manifest manifestparser.Manifest, overrides FlagOverrides) (manifestparser.Manifest, error) {
	if overrides.StartCommand.IsSet {
		if manifest.ContainsMultipleApps() {
			return manifest, translatableerror.CommandLineArgsWithMultipleAppsError{}
		}

		webProcess := manifest.GetFirstAppWebProcess()
		if webProcess != nil {
			webProcess.SetStartCommand(overrides.StartCommand.Value)
		} else {
			app := manifest.GetFirstApp()
			app.SetStartCommand(overrides.StartCommand.Value)
		}
	}

	return manifest, nil
}
