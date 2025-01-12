// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type OptionsVolumesSnapshotSnapshotGroupNameRequest struct {
	SnapshotGroupName string `pathParam:"style=simple,explode=false,name=snapshot_group_name"`
}

func (o *OptionsVolumesSnapshotSnapshotGroupNameRequest) GetSnapshotGroupName() string {
	if o == nil {
		return ""
	}
	return o.SnapshotGroupName
}

type OptionsVolumesSnapshotSnapshotGroupNameResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// 500 response
	ErrorResponse *components.ErrorResponse
	Headers       map[string][]string
}

func (o *OptionsVolumesSnapshotSnapshotGroupNameResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *OptionsVolumesSnapshotSnapshotGroupNameResponse) GetErrorResponse() *components.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *OptionsVolumesSnapshotSnapshotGroupNameResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}