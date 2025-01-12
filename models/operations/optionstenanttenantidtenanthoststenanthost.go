// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type OptionsTenantTenantIDTenanthostsTenantHostRequest struct {
	TenantID   string `pathParam:"style=simple,explode=false,name=tenantID"`
	TenantHost string `pathParam:"style=simple,explode=false,name=tenantHost"`
}

func (o *OptionsTenantTenantIDTenanthostsTenantHostRequest) GetTenantID() string {
	if o == nil {
		return ""
	}
	return o.TenantID
}

func (o *OptionsTenantTenantIDTenanthostsTenantHostRequest) GetTenantHost() string {
	if o == nil {
		return ""
	}
	return o.TenantHost
}

type OptionsTenantTenantIDTenanthostsTenantHostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// 500 response
	ErrorResponse *components.ErrorResponse
	Headers       map[string][]string
}

func (o *OptionsTenantTenantIDTenanthostsTenantHostResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *OptionsTenantTenantIDTenanthostsTenantHostResponse) GetErrorResponse() *components.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *OptionsTenantTenantIDTenanthostsTenantHostResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}