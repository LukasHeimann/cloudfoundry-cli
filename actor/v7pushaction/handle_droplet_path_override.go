package v7pushaction

import (
	"github.com/LukasHeimann/cloudfoundrycli/v8/command/translatableerror"
	"github.com/LukasHeimann/cloudfoundrycli/v8/util/manifestparser"
)

func HandleDropletPathOverride(manifest manifestparser.Manifest, overrides FlagOverrides) (manifestparser.Manifest, error) {
	if overrides.DropletPath != "" {
		if manifest.ContainsMultipleApps() {
			return manifest, translatableerror.CommandLineArgsWithMultipleAppsError{}
		}

		app := manifest.GetFirstApp()

		if app.Docker != nil {
			return manifest, translatableerror.ArgumentManifestMismatchError{
				Arg:              "--droplet",
				ManifestProperty: "docker",
			}
		}

		if app.Path != "" {
			return manifest, translatableerror.ArgumentManifestMismatchError{
				Arg:              "--droplet",
				ManifestProperty: "path",
			}
		}

		if app.HasBuildpacks() {
			return manifest, translatableerror.ArgumentManifestMismatchError{
				Arg:              "--droplet",
				ManifestProperty: "buildpacks",
			}
		}
	}

	return manifest, nil
}
