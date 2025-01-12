// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type MediaModifyRequest struct {
	Media string `pathParam:"style=simple,explode=false,name=media"`
	// A Media Modify Object
	MediaModify components.MediaModify `request:"mediaType=application/json"`
}

func (o *MediaModifyRequest) GetMedia() string {
	if o == nil {
		return ""
	}
	return o.Media
}

func (o *MediaModifyRequest) GetMediaModify() components.MediaModify {
	if o == nil {
		return components.MediaModify{}
	}
	return o.MediaModify
}

type MediaModifyResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Properties of a media to patch
	RegularResponse *components.RegularResponse
	Headers         map[string][]string
}

func (o *MediaModifyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *MediaModifyResponse) GetRegularResponse() *components.RegularResponse {
	if o == nil {
		return nil
	}
	return o.RegularResponse
}

func (o *MediaModifyResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
