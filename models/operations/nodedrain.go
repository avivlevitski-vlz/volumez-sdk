// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/internal/utils"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type NodeDrainRequest struct {
	Node    string `pathParam:"style=simple,explode=false,name=node"`
	Cleanup *bool  `default:"false" queryParam:"style=form,explode=true,name=cleanup"`
}

func (n NodeDrainRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(n, "", false)
}

func (n *NodeDrainRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &n, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *NodeDrainRequest) GetNode() string {
	if o == nil {
		return ""
	}
	return o.Node
}

func (o *NodeDrainRequest) GetCleanup() *bool {
	if o == nil {
		return nil
	}
	return o.Cleanup
}

type NodeDrainResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Operation completed successfully
	RegularResponse *components.RegularResponse
	Headers         map[string][]string
}

func (o *NodeDrainResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *NodeDrainResponse) GetRegularResponse() *components.RegularResponse {
	if o == nil {
		return nil
	}
	return o.RegularResponse
}

func (o *NodeDrainResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
