// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type RefreshToken struct {
	Token *string `json:"token,omitempty"`
}

func (o *RefreshToken) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}
