package v7pushaction

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/api/cloudcontroller/ccv3/constant"
	"github.com/LukasHeimann/cloudfoundrycli/v8/command/translatableerror"
	"github.com/LukasHeimann/cloudfoundrycli/v8/util/manifestparser"
)

func HandleHealthCheckTypeOverride(manifest manifestparser.Manifest, overrides FlagOverrides) (manifestparser.Manifest, error) {
	if overrides.HealthCheckType != "" {
		if manifest.ContainsMultipleApps() {
			return manifest, translatableerror.CommandLineArgsWithMultipleAppsError{}
		}

		webProcess := manifest.GetFirstAppWebProcess()
		if webProcess != nil {
			webProcess.HealthCheckType = overrides.HealthCheckType
			if webProcess.HealthCheckType != constant.HTTP {
				webProcess.HealthCheckEndpoint = ""
			}
		} else {
			app := manifest.GetFirstApp()
			app.HealthCheckType = overrides.HealthCheckType
			if app.HealthCheckType != constant.HTTP {
				app.HealthCheckEndpoint = ""
			}
		}
	}

	return manifest, nil
}
