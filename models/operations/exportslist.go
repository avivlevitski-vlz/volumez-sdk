// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type ExportsListRequest struct {
	Startfrom *string `queryParam:"style=form,explode=true,name=startfrom"`
	Count     *int64  `queryParam:"style=form,explode=true,name=count"`
}

func (o *ExportsListRequest) GetStartfrom() *string {
	if o == nil {
		return nil
	}
	return o.Startfrom
}

func (o *ExportsListRequest) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

type ExportsListResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List of exports
	Exports []components.Export
	Headers map[string][]string
}

func (o *ExportsListResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ExportsListResponse) GetExports() []components.Export {
	if o == nil {
		return nil
	}
	return o.Exports
}

func (o *ExportsListResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
