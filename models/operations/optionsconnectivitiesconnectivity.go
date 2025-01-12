// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type OptionsConnectivitiesConnectivityRequest struct {
	Connectivity string `pathParam:"style=simple,explode=false,name=connectivity"`
}

func (o *OptionsConnectivitiesConnectivityRequest) GetConnectivity() string {
	if o == nil {
		return ""
	}
	return o.Connectivity
}

type OptionsConnectivitiesConnectivityResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// 500 response
	ErrorResponse *components.ErrorResponse
	Headers       map[string][]string
}

func (o *OptionsConnectivitiesConnectivityResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *OptionsConnectivitiesConnectivityResponse) GetErrorResponse() *components.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *OptionsConnectivitiesConnectivityResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
