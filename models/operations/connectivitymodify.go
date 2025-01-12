// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type ConnectivityModifyRequest struct {
	ConnectivityPathParameter string `pathParam:"style=simple,explode=false,name=connectivity"`
	// A Connectivity object
	Connectivity1 components.Connectivity `request:"mediaType=application/json"`
}

func (o *ConnectivityModifyRequest) GetConnectivityPathParameter() string {
	if o == nil {
		return ""
	}
	return o.ConnectivityPathParameter
}

func (o *ConnectivityModifyRequest) GetConnectivity1() components.Connectivity {
	if o == nil {
		return components.Connectivity{}
	}
	return o.Connectivity1
}

type ConnectivityModifyResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A connectivity was updated successfully
	RegularResponse *components.RegularResponse
	Headers         map[string][]string
}

func (o *ConnectivityModifyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ConnectivityModifyResponse) GetRegularResponse() *components.RegularResponse {
	if o == nil {
		return nil
	}
	return o.RegularResponse
}

func (o *ConnectivityModifyResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}