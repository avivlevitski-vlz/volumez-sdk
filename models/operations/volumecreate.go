// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type VolumeCreateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Volume has been created successfully
	RegularResponse *components.RegularResponse
	// 409 response
	ErrorResponse *components.ErrorResponse
	Headers       map[string][]string
}

func (o *VolumeCreateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *VolumeCreateResponse) GetRegularResponse() *components.RegularResponse {
	if o == nil {
		return nil
	}
	return o.RegularResponse
}

func (o *VolumeCreateResponse) GetErrorResponse() *components.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *VolumeCreateResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
