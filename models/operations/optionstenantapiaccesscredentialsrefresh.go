// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type OptionsTenantApiaccessCredentialsRefreshRequest struct {
	Refreshtoken string `header:"style=simple,explode=false,name=refreshtoken"`
}

func (o *OptionsTenantApiaccessCredentialsRefreshRequest) GetRefreshtoken() string {
	if o == nil {
		return ""
	}
	return o.Refreshtoken
}

type OptionsTenantApiaccessCredentialsRefreshResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// 500 response
	ErrorResponse *components.ErrorResponse
	Headers       map[string][]string
}

func (o *OptionsTenantApiaccessCredentialsRefreshResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *OptionsTenantApiaccessCredentialsRefreshResponse) GetErrorResponse() *components.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *OptionsTenantApiaccessCredentialsRefreshResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}