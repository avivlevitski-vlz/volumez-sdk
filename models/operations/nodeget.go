// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type NodeGetRequest struct {
	// Name of node to return
	Node string `pathParam:"style=simple,explode=false,name=node"`
}

func (o *NodeGetRequest) GetNode() string {
	if o == nil {
		return ""
	}
	return o.Node
}

type NodeGetResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Properties of a node
	Node    *components.Node
	Headers map[string][]string
}

func (o *NodeGetResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *NodeGetResponse) GetNode() *components.Node {
	if o == nil {
		return nil
	}
	return o.Node
}

func (o *NodeGetResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
