// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type NetworkType string

const (
	NetworkTypeManagement NetworkType = "management"
	NetworkTypeStorage    NetworkType = "storage"
)

func (e NetworkType) ToPointer() *NetworkType {
	return &e
}
func (e *NetworkType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "management":
		fallthrough
	case "storage":
		*e = NetworkType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NetworkType: %v", v)
	}
}

type Network struct {
	Name    string      `json:"name"`
	Zone    *string     `json:"zone,omitempty"`
	Type    NetworkType `json:"type"`
	Ipstart string      `json:"ipstart"`
	Ipend   string      `json:"ipend"`
}

func (o *Network) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Network) GetZone() *string {
	if o == nil {
		return nil
	}
	return o.Zone
}

func (o *Network) GetType() NetworkType {
	if o == nil {
		return NetworkType("")
	}
	return o.Type
}

func (o *Network) GetIpstart() string {
	if o == nil {
		return ""
	}
	return o.Ipstart
}

func (o *Network) GetIpend() string {
	if o == nil {
		return ""
	}
	return o.Ipend
}
