// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type SnapshotRollbackRequest struct {
	Volume   string `pathParam:"style=simple,explode=false,name=volume"`
	Snapshot string `pathParam:"style=simple,explode=false,name=snapshot"`
}

func (o *SnapshotRollbackRequest) GetVolume() string {
	if o == nil {
		return ""
	}
	return o.Volume
}

func (o *SnapshotRollbackRequest) GetSnapshot() string {
	if o == nil {
		return ""
	}
	return o.Snapshot
}

type SnapshotRollbackResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Rollback successful
	RegularResponse *components.RegularResponse
	Headers         map[string][]string
}

func (o *SnapshotRollbackResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SnapshotRollbackResponse) GetRegularResponse() *components.RegularResponse {
	if o == nil {
		return nil
	}
	return o.RegularResponse
}

func (o *SnapshotRollbackResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
