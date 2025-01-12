// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type TenantAPIAccessRefreshTokenRequest struct {
	Refreshtoken string `header:"style=simple,explode=false,name=refreshtoken"`
}

func (o *TenantAPIAccessRefreshTokenRequest) GetRefreshtoken() string {
	if o == nil {
		return ""
	}
	return o.Refreshtoken
}

type TenantAPIAccessRefreshTokenResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// 200 response
	RefreshTokenResponse *components.RefreshTokenResponse
	Headers              map[string][]string
}

func (o *TenantAPIAccessRefreshTokenResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *TenantAPIAccessRefreshTokenResponse) GetRefreshTokenResponse() *components.RefreshTokenResponse {
	if o == nil {
		return nil
	}
	return o.RefreshTokenResponse
}

func (o *TenantAPIAccessRefreshTokenResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
