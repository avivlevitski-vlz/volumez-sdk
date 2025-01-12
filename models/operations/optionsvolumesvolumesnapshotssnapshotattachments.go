// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type OptionsVolumesVolumeSnapshotsSnapshotAttachmentsRequest struct {
	Volume   string `pathParam:"style=simple,explode=false,name=volume"`
	Snapshot string `pathParam:"style=simple,explode=false,name=snapshot"`
}

func (o *OptionsVolumesVolumeSnapshotsSnapshotAttachmentsRequest) GetVolume() string {
	if o == nil {
		return ""
	}
	return o.Volume
}

func (o *OptionsVolumesVolumeSnapshotsSnapshotAttachmentsRequest) GetSnapshot() string {
	if o == nil {
		return ""
	}
	return o.Snapshot
}

type OptionsVolumesVolumeSnapshotsSnapshotAttachmentsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// 500 response
	ErrorResponse *components.ErrorResponse
	Headers       map[string][]string
}

func (o *OptionsVolumesVolumeSnapshotsSnapshotAttachmentsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *OptionsVolumesVolumeSnapshotsSnapshotAttachmentsResponse) GetErrorResponse() *components.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *OptionsVolumesVolumeSnapshotsSnapshotAttachmentsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
