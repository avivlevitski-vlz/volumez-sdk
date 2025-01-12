// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type PolicyGetVolumesRequest struct {
	Policy string `pathParam:"style=simple,explode=false,name=policy"`
}

func (o *PolicyGetVolumesRequest) GetPolicy() string {
	if o == nil {
		return ""
	}
	return o.Policy
}

type PolicyGetVolumesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List of volumes, using given policy
	Volumes []components.Volume
	Headers map[string][]string
}

func (o *PolicyGetVolumesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PolicyGetVolumesResponse) GetVolumes() []components.Volume {
	if o == nil {
		return nil
	}
	return o.Volumes
}

func (o *PolicyGetVolumesResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}