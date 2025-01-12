// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type ChangePasswordLoggedInResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// New password changed successfully
	RegularResponse *components.RegularResponse
	Headers         map[string][]string
}

func (o *ChangePasswordLoggedInResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ChangePasswordLoggedInResponse) GetRegularResponse() *components.RegularResponse {
	if o == nil {
		return nil
	}
	return o.RegularResponse
}

func (o *ChangePasswordLoggedInResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
