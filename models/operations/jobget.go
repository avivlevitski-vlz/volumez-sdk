// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
)

type JobGetRequest struct {
	Job int64 `pathParam:"style=simple,explode=false,name=job"`
}

func (o *JobGetRequest) GetJob() int64 {
	if o == nil {
		return 0
	}
	return o.Job
}

type JobGetResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Properties of a job
	Job     *components.Job
	Headers map[string][]string
}

func (o *JobGetResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *JobGetResponse) GetJob() *components.Job {
	if o == nil {
		return nil
	}
	return o.Job
}

func (o *JobGetResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
