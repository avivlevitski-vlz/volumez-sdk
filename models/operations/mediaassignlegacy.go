// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/internal/utils"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type MediaAssignLegacyRequest struct {
	Media         string  `pathParam:"style=simple,explode=false,name=media"`
	CapacityGroup *string `queryParam:"style=form,explode=true,name=capacity_group"`
	Reprofile     *bool   `default:"false" queryParam:"style=form,explode=true,name=reprofile"`
}

func (m MediaAssignLegacyRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MediaAssignLegacyRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MediaAssignLegacyRequest) GetMedia() string {
	if o == nil {
		return ""
	}
	return o.Media
}

func (o *MediaAssignLegacyRequest) GetCapacityGroup() *string {
	if o == nil {
		return nil
	}
	return o.CapacityGroup
}

func (o *MediaAssignLegacyRequest) GetReprofile() *bool {
	if o == nil {
		return nil
	}
	return o.Reprofile
}

type MediaAssignLegacyResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Operation completed successfully
	RegularResponse *components.RegularResponse
	Headers         map[string][]string
}

func (o *MediaAssignLegacyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *MediaAssignLegacyResponse) GetRegularResponse() *components.RegularResponse {
	if o == nil {
		return nil
	}
	return o.RegularResponse
}

func (o *MediaAssignLegacyResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
