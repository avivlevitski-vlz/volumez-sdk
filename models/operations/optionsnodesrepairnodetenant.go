// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type OptionsNodesRepairNodeTenantRequest struct {
	// Name of node
	Node string `pathParam:"style=simple,explode=false,name=node"`
	// Tenant ID
	Tenant string `pathParam:"style=simple,explode=false,name=tenant"`
}

func (o *OptionsNodesRepairNodeTenantRequest) GetNode() string {
	if o == nil {
		return ""
	}
	return o.Node
}

func (o *OptionsNodesRepairNodeTenantRequest) GetTenant() string {
	if o == nil {
		return ""
	}
	return o.Tenant
}

type OptionsNodesRepairNodeTenantResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// 500 response
	ErrorResponse *components.ErrorResponse
	Headers       map[string][]string
}

func (o *OptionsNodesRepairNodeTenantResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *OptionsNodesRepairNodeTenantResponse) GetErrorResponse() *components.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *OptionsNodesRepairNodeTenantResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}