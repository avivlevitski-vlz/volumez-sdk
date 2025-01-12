// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type CapacityGroupGetResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// capacity groups
	CapacityGroups []components.CapacityGroup
	Headers        map[string][]string
}

func (o *CapacityGroupGetResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CapacityGroupGetResponse) GetCapacityGroups() []components.CapacityGroup {
	if o == nil {
		return nil
	}
	return o.CapacityGroups
}

func (o *CapacityGroupGetResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
