// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type AlertsListRequest struct {
	Capacity *bool `queryParam:"style=form,explode=true,name=capacity"`
}

func (o *AlertsListRequest) GetCapacity() *bool {
	if o == nil {
		return nil
	}
	return o.Capacity
}

type AlertsListResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List of alerts
	Alerts  []components.Alert
	Headers map[string][]string
}

func (o *AlertsListResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AlertsListResponse) GetAlerts() []components.Alert {
	if o == nil {
		return nil
	}
	return o.Alerts
}

func (o *AlertsListResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
