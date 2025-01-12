// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type PolicyDeleteRequest struct {
	Policy string `pathParam:"style=simple,explode=false,name=policy"`
}

func (o *PolicyDeleteRequest) GetPolicy() string {
	if o == nil {
		return ""
	}
	return o.Policy
}

type PolicyDeleteResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A policy was deleted successfully
	RegularResponse *components.RegularResponse
	Headers         map[string][]string
}

func (o *PolicyDeleteResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PolicyDeleteResponse) GetRegularResponse() *components.RegularResponse {
	if o == nil {
		return nil
	}
	return o.RegularResponse
}

func (o *PolicyDeleteResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
