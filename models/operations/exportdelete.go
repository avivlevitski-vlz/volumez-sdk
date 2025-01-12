// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/internal/utils"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type ExportDeleteRequest struct {
	Export string `pathParam:"style=simple,explode=false,name=export"`
	Force  *bool  `default:"false" queryParam:"style=form,explode=true,name=force"`
}

func (e ExportDeleteRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *ExportDeleteRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ExportDeleteRequest) GetExport() string {
	if o == nil {
		return ""
	}
	return o.Export
}

func (o *ExportDeleteRequest) GetForce() *bool {
	if o == nil {
		return nil
	}
	return o.Force
}

type ExportDeleteResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Export was deleted successfully
	SuccessJobResponse *components.SuccessJobResponse
	Headers            map[string][]string
}

func (o *ExportDeleteResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ExportDeleteResponse) GetSuccessJobResponse() *components.SuccessJobResponse {
	if o == nil {
		return nil
	}
	return o.SuccessJobResponse
}

func (o *ExportDeleteResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
