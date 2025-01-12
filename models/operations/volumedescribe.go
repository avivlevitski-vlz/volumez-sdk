// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type VolumeDescribeRequest struct {
	Volume string `pathParam:"style=simple,explode=false,name=volume"`
}

func (o *VolumeDescribeRequest) GetVolume() string {
	if o == nil {
		return ""
	}
	return o.Volume
}

type VolumeDescribeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// the Volume information
	VolumeGroup *components.VolumeGroup
	Headers     map[string][]string
}

func (o *VolumeDescribeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *VolumeDescribeResponse) GetVolumeGroup() *components.VolumeGroup {
	if o == nil {
		return nil
	}
	return o.VolumeGroup
}

func (o *VolumeDescribeResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
