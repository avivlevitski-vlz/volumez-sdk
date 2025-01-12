// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type MediaAssignRequest struct {
	Media       string                  `pathParam:"style=simple,explode=false,name=media"`
	MediaAssign *components.MediaAssign `request:"mediaType=application/json"`
}

func (o *MediaAssignRequest) GetMedia() string {
	if o == nil {
		return ""
	}
	return o.Media
}

func (o *MediaAssignRequest) GetMediaAssign() *components.MediaAssign {
	if o == nil {
		return nil
	}
	return o.MediaAssign
}

type MediaAssignResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Operation completed successfully
	RegularResponse *components.RegularResponse
	Headers         map[string][]string
}

func (o *MediaAssignResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *MediaAssignResponse) GetRegularResponse() *components.RegularResponse {
	if o == nil {
		return nil
	}
	return o.RegularResponse
}

func (o *MediaAssignResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
