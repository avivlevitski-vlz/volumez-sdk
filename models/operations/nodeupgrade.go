// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type NodeUpgradeRequest struct {
	// Name of node to upgrade
	Node string `pathParam:"style=simple,explode=false,name=node"`
	// Connector Version
	NodeVersion *components.NodeVersion `request:"mediaType=application/json"`
}

func (o *NodeUpgradeRequest) GetNode() string {
	if o == nil {
		return ""
	}
	return o.Node
}

func (o *NodeUpgradeRequest) GetNodeVersion() *components.NodeVersion {
	if o == nil {
		return nil
	}
	return o.NodeVersion
}

type NodeUpgradeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Node upgrade started successfully
	RegularResponse *components.RegularResponse
	Headers         map[string][]string
}

func (o *NodeUpgradeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *NodeUpgradeResponse) GetRegularResponse() *components.RegularResponse {
	if o == nil {
		return nil
	}
	return o.RegularResponse
}

func (o *NodeUpgradeResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}