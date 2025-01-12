// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type OptionsSnapshotsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// 500 response
	ErrorResponse *components.ErrorResponse
	Headers       map[string][]string
}

func (o *OptionsSnapshotsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *OptionsSnapshotsResponse) GetErrorResponse() *components.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *OptionsSnapshotsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}