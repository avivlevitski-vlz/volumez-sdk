// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
)

type SignUpResponse struct {
	Status  int64  `json:"status"`
	Message string `json:"message"`
}

var _ error = &SignUpResponse{}

func (e *SignUpResponse) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
