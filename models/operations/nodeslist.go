// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type NodesListResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List of nodes
	Nodes   []components.Node
	Headers map[string][]string
}

func (o *NodesListResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *NodesListResponse) GetNodes() []components.Node {
	if o == nil {
		return nil
	}
	return o.Nodes
}

func (o *NodesListResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
