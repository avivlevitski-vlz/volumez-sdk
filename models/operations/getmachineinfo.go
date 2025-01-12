// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type GetMachineInfoResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// confirmation of subscription registered
	MachineInfo *components.MachineInfo
	Headers     map[string][]string
}

func (o *GetMachineInfoResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetMachineInfoResponse) GetMachineInfo() *components.MachineInfo {
	if o == nil {
		return nil
	}
	return o.MachineInfo
}

func (o *GetMachineInfoResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
