// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type RequestChangePasswordRequest struct {
	Email *string `json:"email,omitempty"`
}

func (o *RequestChangePasswordRequest) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}
