// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type OptionsVolumesVolumeSnapshotsSnapshotAttachmentsNodeRequest struct {
	Volume   string `pathParam:"style=simple,explode=false,name=volume"`
	Snapshot string `pathParam:"style=simple,explode=false,name=snapshot"`
	Node     string `pathParam:"style=simple,explode=false,name=node"`
}

func (o *OptionsVolumesVolumeSnapshotsSnapshotAttachmentsNodeRequest) GetVolume() string {
	if o == nil {
		return ""
	}
	return o.Volume
}

func (o *OptionsVolumesVolumeSnapshotsSnapshotAttachmentsNodeRequest) GetSnapshot() string {
	if o == nil {
		return ""
	}
	return o.Snapshot
}

func (o *OptionsVolumesVolumeSnapshotsSnapshotAttachmentsNodeRequest) GetNode() string {
	if o == nil {
		return ""
	}
	return o.Node
}

type OptionsVolumesVolumeSnapshotsSnapshotAttachmentsNodeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// 500 response
	ErrorResponse *components.ErrorResponse
	Headers       map[string][]string
}

func (o *OptionsVolumesVolumeSnapshotsSnapshotAttachmentsNodeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *OptionsVolumesVolumeSnapshotsSnapshotAttachmentsNodeResponse) GetErrorResponse() *components.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *OptionsVolumesVolumeSnapshotsSnapshotAttachmentsNodeResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
