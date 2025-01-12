// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type OptionsVolumesVolumeRecoverRequest struct {
	Volume string `pathParam:"style=simple,explode=false,name=volume"`
}

func (o *OptionsVolumesVolumeRecoverRequest) GetVolume() string {
	if o == nil {
		return ""
	}
	return o.Volume
}

type OptionsVolumesVolumeRecoverResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// 500 response
	ErrorResponse *components.ErrorResponse
	Headers       map[string][]string
}

func (o *OptionsVolumesVolumeRecoverResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *OptionsVolumesVolumeRecoverResponse) GetErrorResponse() *components.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *OptionsVolumesVolumeRecoverResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
