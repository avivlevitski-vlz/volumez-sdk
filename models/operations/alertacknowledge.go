// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type AlertAcknowledgeRequest struct {
	Alert string `pathParam:"style=simple,explode=false,name=alert"`
}

func (o *AlertAcknowledgeRequest) GetAlert() string {
	if o == nil {
		return ""
	}
	return o.Alert
}

type AlertAcknowledgeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Alert was successfully acknowledged
	RegularResponse *components.RegularResponse
	Headers         map[string][]string
}

func (o *AlertAcknowledgeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AlertAcknowledgeResponse) GetRegularResponse() *components.RegularResponse {
	if o == nil {
		return nil
	}
	return o.RegularResponse
}

func (o *AlertAcknowledgeResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}