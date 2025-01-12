// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type GetSubscriptionRequest struct {
	Token string `queryParam:"style=form,explode=true,name=token"`
}

func (o *GetSubscriptionRequest) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type GetSubscriptionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// confirmation of subscription registered
	RegularResponse *components.RegularResponse
	Headers         map[string][]string
}

func (o *GetSubscriptionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetSubscriptionResponse) GetRegularResponse() *components.RegularResponse {
	if o == nil {
		return nil
	}
	return o.RegularResponse
}

func (o *GetSubscriptionResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
