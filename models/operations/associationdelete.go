// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type AssociationDeleteRequest struct {
	Association string `pathParam:"style=simple,explode=false,name=association"`
}

func (o *AssociationDeleteRequest) GetAssociation() string {
	if o == nil {
		return ""
	}
	return o.Association
}

type AssociationDeleteResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// An associations was deleted successfully
	RegularResponse *components.RegularResponse
	Headers         map[string][]string
}

func (o *AssociationDeleteResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AssociationDeleteResponse) GetRegularResponse() *components.RegularResponse {
	if o == nil {
		return nil
	}
	return o.RegularResponse
}

func (o *AssociationDeleteResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}