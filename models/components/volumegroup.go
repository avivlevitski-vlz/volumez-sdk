// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type VolumeGroup struct {
	Volumegroupname          *string      `json:"Volumegroupname,omitempty"`
	Encryption               *bool        `json:"encryption,omitempty"`
	Integrity                *bool        `json:"Integrity,omitempty"`
	Allocation               *string      `json:"allocation,omitempty"`
	Compression              *bool        `json:"compression,omitempty"`
	Deduplication            *bool        `json:"deduplication,omitempty"`
	Writecache               *bool        `json:"writecache,omitempty"`
	Redundancy               *int64       `json:"redundancy,omitempty"`
	Size                     *int64       `json:"size,omitempty"`
	Targetsecret             *string      `json:"targetsecret,omitempty"`
	Allocatedsize            *int64       `json:"allocatedsize,omitempty"`
	Resiliency               *string      `json:"resiliency,omitempty"`
	Raidcolumns              *int64       `json:"raidcolumns,omitempty"`
	Mediasize                *int64       `json:"mediasize,omitempty"`
	Mediabandwidthwrite      *int64       `json:"mediabandwidthwrite,omitempty"`
	Mediabandwidthread       *int64       `json:"mediabandwidthread,omitempty"`
	Mediaiopswrite           *int64       `json:"mediaiopswrite,omitempty"`
	Mediaiopsread            *int64       `json:"mediaiopsread,omitempty"`
	Media                    []string     `json:"media,omitempty"`
	Cachesize                *int64       `json:"cachesize,omitempty"`
	Cacheresiliency          *string      `json:"cacheresiliency,omitempty"`
	Cacheredundancy          *int64       `json:"cacheredundancy,omitempty"`
	Cacheraidcolumns         *int64       `json:"cacheraidcolumns,omitempty"`
	Cachemediasize           *int64       `json:"cachemediasize,omitempty"`
	Cachemediabandwidthwrite *int64       `json:"cachemediabandwidthwrite,omitempty"`
	Cachemediabandwidthread  *int64       `json:"cachemediabandwidthread,omitempty"`
	Cachemediaiopswrite      *int64       `json:"cachemediaiopswrite,omitempty"`
	Cachemediaiopsread       *int64       `json:"cachemediaiopsread,omitempty"`
	Cachemedia               []string     `json:"cachemedia,omitempty"`
	Attachments              []Attachment `json:"attachments,omitempty"`
}

func (o *VolumeGroup) GetVolumegroupname() *string {
	if o == nil {
		return nil
	}
	return o.Volumegroupname
}

func (o *VolumeGroup) GetEncryption() *bool {
	if o == nil {
		return nil
	}
	return o.Encryption
}

func (o *VolumeGroup) GetIntegrity() *bool {
	if o == nil {
		return nil
	}
	return o.Integrity
}

func (o *VolumeGroup) GetAllocation() *string {
	if o == nil {
		return nil
	}
	return o.Allocation
}

func (o *VolumeGroup) GetCompression() *bool {
	if o == nil {
		return nil
	}
	return o.Compression
}

func (o *VolumeGroup) GetDeduplication() *bool {
	if o == nil {
		return nil
	}
	return o.Deduplication
}

func (o *VolumeGroup) GetWritecache() *bool {
	if o == nil {
		return nil
	}
	return o.Writecache
}

func (o *VolumeGroup) GetRedundancy() *int64 {
	if o == nil {
		return nil
	}
	return o.Redundancy
}

func (o *VolumeGroup) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *VolumeGroup) GetTargetsecret() *string {
	if o == nil {
		return nil
	}
	return o.Targetsecret
}

func (o *VolumeGroup) GetAllocatedsize() *int64 {
	if o == nil {
		return nil
	}
	return o.Allocatedsize
}

func (o *VolumeGroup) GetResiliency() *string {
	if o == nil {
		return nil
	}
	return o.Resiliency
}

func (o *VolumeGroup) GetRaidcolumns() *int64 {
	if o == nil {
		return nil
	}
	return o.Raidcolumns
}

func (o *VolumeGroup) GetMediasize() *int64 {
	if o == nil {
		return nil
	}
	return o.Mediasize
}

func (o *VolumeGroup) GetMediabandwidthwrite() *int64 {
	if o == nil {
		return nil
	}
	return o.Mediabandwidthwrite
}

func (o *VolumeGroup) GetMediabandwidthread() *int64 {
	if o == nil {
		return nil
	}
	return o.Mediabandwidthread
}

func (o *VolumeGroup) GetMediaiopswrite() *int64 {
	if o == nil {
		return nil
	}
	return o.Mediaiopswrite
}

func (o *VolumeGroup) GetMediaiopsread() *int64 {
	if o == nil {
		return nil
	}
	return o.Mediaiopsread
}

func (o *VolumeGroup) GetMedia() []string {
	if o == nil {
		return nil
	}
	return o.Media
}

func (o *VolumeGroup) GetCachesize() *int64 {
	if o == nil {
		return nil
	}
	return o.Cachesize
}

func (o *VolumeGroup) GetCacheresiliency() *string {
	if o == nil {
		return nil
	}
	return o.Cacheresiliency
}

func (o *VolumeGroup) GetCacheredundancy() *int64 {
	if o == nil {
		return nil
	}
	return o.Cacheredundancy
}

func (o *VolumeGroup) GetCacheraidcolumns() *int64 {
	if o == nil {
		return nil
	}
	return o.Cacheraidcolumns
}

func (o *VolumeGroup) GetCachemediasize() *int64 {
	if o == nil {
		return nil
	}
	return o.Cachemediasize
}

func (o *VolumeGroup) GetCachemediabandwidthwrite() *int64 {
	if o == nil {
		return nil
	}
	return o.Cachemediabandwidthwrite
}

func (o *VolumeGroup) GetCachemediabandwidthread() *int64 {
	if o == nil {
		return nil
	}
	return o.Cachemediabandwidthread
}

func (o *VolumeGroup) GetCachemediaiopswrite() *int64 {
	if o == nil {
		return nil
	}
	return o.Cachemediaiopswrite
}

func (o *VolumeGroup) GetCachemediaiopsread() *int64 {
	if o == nil {
		return nil
	}
	return o.Cachemediaiopsread
}

func (o *VolumeGroup) GetCachemedia() []string {
	if o == nil {
		return nil
	}
	return o.Cachemedia
}

func (o *VolumeGroup) GetAttachments() []Attachment {
	if o == nil {
		return nil
	}
	return o.Attachments
}
