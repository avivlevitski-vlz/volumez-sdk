// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type TenantCreateAdminUserRequest struct {
	Email         string  `json:"email"`
	Password      string  `json:"password"`
	Name          string  `json:"name"`
	Cloudprovider *string `json:"cloudprovider,omitempty"`
	Markettoken   *string `json:"markettoken,omitempty"`
}

func (o *TenantCreateAdminUserRequest) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *TenantCreateAdminUserRequest) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *TenantCreateAdminUserRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *TenantCreateAdminUserRequest) GetCloudprovider() *string {
	if o == nil {
		return nil
	}
	return o.Cloudprovider
}

func (o *TenantCreateAdminUserRequest) GetMarkettoken() *string {
	if o == nil {
		return nil
	}
	return o.Markettoken
}
