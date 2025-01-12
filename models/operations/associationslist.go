// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type AssociationsListRequest struct {
	Startfrom *string `queryParam:"style=form,explode=true,name=startfrom"`
	Count     *int64  `queryParam:"style=form,explode=true,name=count"`
}

func (o *AssociationsListRequest) GetStartfrom() *string {
	if o == nil {
		return nil
	}
	return o.Startfrom
}

func (o *AssociationsListRequest) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

type AssociationsListResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List of associations
	Associations []components.Association
	Headers      map[string][]string
}

func (o *AssociationsListResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AssociationsListResponse) GetAssociations() []components.Association {
	if o == nil {
		return nil
	}
	return o.Associations
}

func (o *AssociationsListResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}