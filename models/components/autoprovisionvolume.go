// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type AutoProvisionVolumeOsType string

const (
	AutoProvisionVolumeOsTypeLinux  AutoProvisionVolumeOsType = "Linux"
	AutoProvisionVolumeOsTypeRhel   AutoProvisionVolumeOsType = "Rhel"
	AutoProvisionVolumeOsTypeUbuntu AutoProvisionVolumeOsType = "Ubuntu"
)

func (e AutoProvisionVolumeOsType) ToPointer() *AutoProvisionVolumeOsType {
	return &e
}
func (e *AutoProvisionVolumeOsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Linux":
		fallthrough
	case "Rhel":
		fallthrough
	case "Ubuntu":
		*e = AutoProvisionVolumeOsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AutoProvisionVolumeOsType: %v", v)
	}
}

type AutoProvisionVolume struct {
	Volume            VolumeInput               `json:"volume"`
	InfraPlan         *AutoProvisionInfraPlan   `json:"infraPlan,omitempty"`
	CloudProvider     string                    `json:"cloudProvider"`
	AccountID         string                    `json:"accountId"`
	SSHKeyName        *string                   `json:"sshKeyName,omitempty"`
	ImageID           *string                   `json:"imageId,omitempty"`
	Region            string                    `json:"region"`
	AvailabilityZones []string                  `json:"availabilityZones"`
	Subnets           []string                  `json:"subnets"`
	OsType            AutoProvisionVolumeOsType `json:"osType"`
}

func (o *AutoProvisionVolume) GetVolume() VolumeInput {
	if o == nil {
		return VolumeInput{}
	}
	return o.Volume
}

func (o *AutoProvisionVolume) GetInfraPlan() *AutoProvisionInfraPlan {
	if o == nil {
		return nil
	}
	return o.InfraPlan
}

func (o *AutoProvisionVolume) GetCloudProvider() string {
	if o == nil {
		return ""
	}
	return o.CloudProvider
}

func (o *AutoProvisionVolume) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *AutoProvisionVolume) GetSSHKeyName() *string {
	if o == nil {
		return nil
	}
	return o.SSHKeyName
}

func (o *AutoProvisionVolume) GetImageID() *string {
	if o == nil {
		return nil
	}
	return o.ImageID
}

func (o *AutoProvisionVolume) GetRegion() string {
	if o == nil {
		return ""
	}
	return o.Region
}

func (o *AutoProvisionVolume) GetAvailabilityZones() []string {
	if o == nil {
		return []string{}
	}
	return o.AvailabilityZones
}

func (o *AutoProvisionVolume) GetSubnets() []string {
	if o == nil {
		return []string{}
	}
	return o.Subnets
}

func (o *AutoProvisionVolume) GetOsType() AutoProvisionVolumeOsType {
	if o == nil {
		return AutoProvisionVolumeOsType("")
	}
	return o.OsType
}
