package v7pushaction

import "github.com/LukasHeimann/cloudfoundrycli/v8/util/manifestparser"

type HandleFlagOverrideFunc func(manifest manifestparser.Manifest, overrides FlagOverrides) (manifestparser.Manifest, error)
