// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type VolumesListRequest struct {
	Capacity *bool `queryParam:"style=form,explode=true,name=capacity"`
}

func (o *VolumesListRequest) GetCapacity() *bool {
	if o == nil {
		return nil
	}
	return o.Capacity
}

type VolumesListResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List of volumes tests
	Volumes []components.Volume
	Headers map[string][]string
}

func (o *VolumesListResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *VolumesListResponse) GetVolumes() []components.Volume {
	if o == nil {
		return nil
	}
	return o.Volumes
}

func (o *VolumesListResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
