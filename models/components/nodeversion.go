// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type NodeVersion struct {
	Version *string `json:"version,omitempty"`
}

func (o *NodeVersion) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}
