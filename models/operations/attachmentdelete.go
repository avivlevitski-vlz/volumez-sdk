// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/internal/utils"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type AttachmentDeleteRequest struct {
	Volume   string `pathParam:"style=simple,explode=false,name=volume"`
	Snapshot string `pathParam:"style=simple,explode=false,name=snapshot"`
	Node     string `pathParam:"style=simple,explode=false,name=node"`
	Force    *bool  `default:"false" queryParam:"style=form,explode=true,name=force"`
}

func (a AttachmentDeleteRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AttachmentDeleteRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AttachmentDeleteRequest) GetVolume() string {
	if o == nil {
		return ""
	}
	return o.Volume
}

func (o *AttachmentDeleteRequest) GetSnapshot() string {
	if o == nil {
		return ""
	}
	return o.Snapshot
}

func (o *AttachmentDeleteRequest) GetNode() string {
	if o == nil {
		return ""
	}
	return o.Node
}

func (o *AttachmentDeleteRequest) GetForce() *bool {
	if o == nil {
		return nil
	}
	return o.Force
}

type AttachmentDeleteResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// An attachment was deleted successfully
	RegularResponse *components.RegularResponse
	Headers         map[string][]string
}

func (o *AttachmentDeleteResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AttachmentDeleteResponse) GetRegularResponse() *components.RegularResponse {
	if o == nil {
		return nil
	}
	return o.RegularResponse
}

func (o *AttachmentDeleteResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}