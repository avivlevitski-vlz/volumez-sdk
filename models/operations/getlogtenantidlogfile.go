// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type GetLogTenantidLogfileRequest struct {
	Logfile  string `pathParam:"style=simple,explode=false,name=logfile"`
	Tenantid string `pathParam:"style=simple,explode=false,name=tenantid"`
}

func (o *GetLogTenantidLogfileRequest) GetLogfile() string {
	if o == nil {
		return ""
	}
	return o.Logfile
}

func (o *GetLogTenantidLogfileRequest) GetTenantid() string {
	if o == nil {
		return ""
	}
	return o.Tenantid
}

type GetLogTenantidLogfileResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Headers  map[string][]string
}

func (o *GetLogTenantidLogfileResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetLogTenantidLogfileResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
