package v7pushaction

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/api/cloudcontroller/ccv3/constant"
	"github.com/LukasHeimann/cloudfoundrycli/v8/command/translatableerror"
	"github.com/LukasHeimann/cloudfoundrycli/v8/util/manifestparser"
)

func HandleHealthCheckEndpointOverride(manifest manifestparser.Manifest, overrides FlagOverrides) (manifestparser.Manifest, error) {
	if overrides.HealthCheckEndpoint != "" {

		if manifest.ContainsMultipleApps() {
			return manifest, translatableerror.CommandLineArgsWithMultipleAppsError{}
		}

		var healthCheckType constant.HealthCheckType

		webProcess := manifest.GetFirstAppWebProcess()
		if webProcess != nil {
			webProcess.HealthCheckEndpoint = overrides.HealthCheckEndpoint
			healthCheckType = webProcess.HealthCheckType
		} else {
			app := manifest.GetFirstApp()
			app.HealthCheckEndpoint = overrides.HealthCheckEndpoint
			healthCheckType = app.HealthCheckType
		}

		if healthCheckType != "" && healthCheckType != constant.HTTP {
			return manifest, translatableerror.ArgumentManifestMismatchError{
				Arg:              "--endpoint",
				ManifestProperty: "health-check-type",
				ManifestValue:    string(healthCheckType),
			}
		}
	}

	return manifest, nil
}
