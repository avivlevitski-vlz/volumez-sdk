// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type ConnectivityDeleteRequest struct {
	Connectivity string `pathParam:"style=simple,explode=false,name=connectivity"`
}

func (o *ConnectivityDeleteRequest) GetConnectivity() string {
	if o == nil {
		return ""
	}
	return o.Connectivity
}

type ConnectivityDeleteResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A connectivity was deleted successfully
	RegularResponse *components.RegularResponse
	Headers         map[string][]string
}

func (o *ConnectivityDeleteResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ConnectivityDeleteResponse) GetRegularResponse() *components.RegularResponse {
	if o == nil {
		return nil
	}
	return o.RegularResponse
}

func (o *ConnectivityDeleteResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
