// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type SnapshotGetRequest struct {
	Volume   string `pathParam:"style=simple,explode=false,name=volume"`
	Snapshot string `pathParam:"style=simple,explode=false,name=snapshot"`
}

func (o *SnapshotGetRequest) GetVolume() string {
	if o == nil {
		return ""
	}
	return o.Volume
}

func (o *SnapshotGetRequest) GetSnapshot() string {
	if o == nil {
		return ""
	}
	return o.Snapshot
}

type SnapshotGetResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Properties of a snapshot
	Snapshot *components.Snapshot
	Headers  map[string][]string
}

func (o *SnapshotGetResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SnapshotGetResponse) GetSnapshot() *components.Snapshot {
	if o == nil {
		return nil
	}
	return o.Snapshot
}

func (o *SnapshotGetResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
