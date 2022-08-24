package resources

import "github.com/LukasHeimann/cloudfoundrycli/v8/types"

type Sidecar struct {
	GUID    string               `json:"guid"`
	Name    string               `json:"name"`
	Command types.FilteredString `json:"command"`
}
