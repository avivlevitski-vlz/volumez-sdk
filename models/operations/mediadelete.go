// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/internal/utils"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type MediaDeleteRequest struct {
	Media string `pathParam:"style=simple,explode=false,name=media"`
	Force *bool  `default:"false" queryParam:"style=form,explode=true,name=force"`
}

func (m MediaDeleteRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MediaDeleteRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MediaDeleteRequest) GetMedia() string {
	if o == nil {
		return ""
	}
	return o.Media
}

func (o *MediaDeleteRequest) GetForce() *bool {
	if o == nil {
		return nil
	}
	return o.Force
}

type MediaDeleteResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A node delete job was created successfully
	RegularResponse *components.RegularResponse
	Headers         map[string][]string
}

func (o *MediaDeleteResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *MediaDeleteResponse) GetRegularResponse() *components.RegularResponse {
	if o == nil {
		return nil
	}
	return o.RegularResponse
}

func (o *MediaDeleteResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}