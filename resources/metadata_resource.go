package resources

import "github.com/LukasHeimann/cloudfoundrycli/v8/types"

type Metadata struct {
	Labels map[string]types.NullString `json:"labels,omitempty"`
}

type ResourceMetadata struct {
	Metadata *Metadata `json:"metadata,omitempty"`
}
