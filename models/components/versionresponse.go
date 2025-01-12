// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type VersionResponse struct {
	Version       *string `json:"version,omitempty"`
	ComponentName *string `json:"componentName,omitempty"`
}

func (o *VersionResponse) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}

func (o *VersionResponse) GetComponentName() *string {
	if o == nil {
		return nil
	}
	return o.ComponentName
}