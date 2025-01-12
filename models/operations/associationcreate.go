// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type AssociationCreateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Association has been created successfully
	SuccessJobResponse *components.SuccessJobResponse
	Headers            map[string][]string
}

func (o *AssociationCreateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AssociationCreateResponse) GetSuccessJobResponse() *components.SuccessJobResponse {
	if o == nil {
		return nil
	}
	return o.SuccessJobResponse
}

func (o *AssociationCreateResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}