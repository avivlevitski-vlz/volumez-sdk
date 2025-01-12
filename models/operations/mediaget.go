// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type MediaGetRequest struct {
	Media string `pathParam:"style=simple,explode=false,name=media"`
}

func (o *MediaGetRequest) GetMedia() string {
	if o == nil {
		return ""
	}
	return o.Media
}

type MediaGetResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Properties of a media
	Media   *components.Media
	Headers map[string][]string
}

func (o *MediaGetResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *MediaGetResponse) GetMedia() *components.Media {
	if o == nil {
		return nil
	}
	return o.Media
}

func (o *MediaGetResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
