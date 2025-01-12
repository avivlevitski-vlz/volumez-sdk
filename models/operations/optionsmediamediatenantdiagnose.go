// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type OptionsMediaMediaTenantDiagnoseRequest struct {
	Media string `pathParam:"style=simple,explode=false,name=media"`
	// Tenant ID
	Tenant string `pathParam:"style=simple,explode=false,name=tenant"`
}

func (o *OptionsMediaMediaTenantDiagnoseRequest) GetMedia() string {
	if o == nil {
		return ""
	}
	return o.Media
}

func (o *OptionsMediaMediaTenantDiagnoseRequest) GetTenant() string {
	if o == nil {
		return ""
	}
	return o.Tenant
}

type OptionsMediaMediaTenantDiagnoseResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// 500 response
	ErrorResponse *components.ErrorResponse
	Headers       map[string][]string
}

func (o *OptionsMediaMediaTenantDiagnoseResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *OptionsMediaMediaTenantDiagnoseResponse) GetErrorResponse() *components.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *OptionsMediaMediaTenantDiagnoseResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
