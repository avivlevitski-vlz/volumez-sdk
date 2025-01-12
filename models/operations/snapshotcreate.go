// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type SnapshotCreateRequest struct {
	Volume string `pathParam:"style=simple,explode=false,name=volume"`
	// A Snapshot object
	Snapshot components.SnapshotInput `request:"mediaType=application/json"`
}

func (o *SnapshotCreateRequest) GetVolume() string {
	if o == nil {
		return ""
	}
	return o.Volume
}

func (o *SnapshotCreateRequest) GetSnapshot() components.SnapshotInput {
	if o == nil {
		return components.SnapshotInput{}
	}
	return o.Snapshot
}

type SnapshotCreateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// New snapshot was created successfully
	RegularResponse *components.RegularResponse
	Headers         map[string][]string
}

func (o *SnapshotCreateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SnapshotCreateResponse) GetRegularResponse() *components.RegularResponse {
	if o == nil {
		return nil
	}
	return o.RegularResponse
}

func (o *SnapshotCreateResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
