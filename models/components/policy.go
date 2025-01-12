// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// Capacityoptimization - Choosing “Capacity” directs Volumez to prefer using capacity-saving methods (such as compression, deduplication, erasure coding and thin provisioning) where relevant, in order to consume the minimum amount of raw media. Using such methods might take some CPU cycles, and might reduce the performance of your volumes (it will still be within the range you specified). Choosing “Balanced” directs Volumez to prefer using some capacity-saving methods where relevant, in order to use less raw media, while consuming a small amount of CPU cycles. “Performance Optimized” directs Volumez to avoid using capacity-saving any methods (such as compression and deduplication) that reduce media consumption. This way applications can get the optimal performance from their media, however more raw media might be consumed to provision Performance-Optimized volumes.
type Capacityoptimization string

const (
	CapacityoptimizationCapacity    Capacityoptimization = "capacity"
	CapacityoptimizationBalanced    Capacityoptimization = "balanced"
	CapacityoptimizationPerformance Capacityoptimization = "performance"
)

func (e Capacityoptimization) ToPointer() *Capacityoptimization {
	return &e
}
func (e *Capacityoptimization) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "capacity":
		fallthrough
	case "balanced":
		fallthrough
	case "performance":
		*e = Capacityoptimization(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Capacityoptimization: %v", v)
	}
}

type Policy struct {
	// A name for the policy. The name can be any non-empty string that does not contain a white space.
	Name string `json:"name"`
	// Enter the maximum write IOPS that a volume is expected to sustain (assuming 8K writes). Write IOPS should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes.
	Iopswrite *int64 `json:"iopswrite,omitempty"`
	// Enter the maximum read IOPS that a volume is expected to sustain (assuming 8K reads). Read IOPS should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes.
	Iopsread *int64 `json:"iopsread,omitempty"`
	// Enter the maximum write bandwidth that a volume is expected to sustain. Write Bandwidth should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes.
	Bandwidthwrite *int64 `json:"bandwidthwrite,omitempty"`
	// Enter the maximum read bandwidth that a volume is expected to sustain. Read Bandwidth should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes.
	Bandwidthread *int64 `json:"bandwidthread,omitempty"`
	// Enter the maximum latency that a volume is expected to sustain. Write latency should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes.
	Latencywrite *int64 `json:"latencywrite,omitempty"`
	// Enter the maximum read IOPS that a volume is expected to sustain. Read latency should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes.
	Latencyread *int64 `json:"latencyread,omitempty"`
	//  If not all the reads are hot (i.e., Percentage of Cold Reads is >0) – Enter the more relaxed constraints for read latencies of cold data.  Valid values: non-negative integer number, that is larger than “Read Latency”.
	Latencyreadcold *int64 `json:"latencyreadcold,omitempty"`
	// Enter the percentage of the volume’s capacity that is expected to be “cold” (i.e. expected to only have infrequent reads). Default is 0%. Values that are greater than 0 give Volumez the option to use more economic media with more relaxed read performance requirements. Valid values: Integers in the range of 0..100.
	Colddata *int64 `json:"colddata,omitempty"`
	// Setting this value to “Yes” directs Volumez to prefer volume configurations where reads are usually happening from disks that are in the same zone as the application. This saves east-west network traffic across zones, however more media per zone will be required to achieve read-IOPs requirements. Set this value to “Yes” if you have network constraints (bandwidth or cost) across your zones; otherwise set it to “No”.
	Localzoneread *bool `json:"localzoneread,omitempty"`
	// Setting this value to “Yes” directs Volumez to over-provision volumes in a way that even after having a failure, the volumes will have the desired performance. Setting this value to “No” directs Volumez to provision volumes according to the desired performance, however in a case of failure – performance may be impacted. The default value is “No”.
	Failureperformance *bool `json:"failureperformance,omitempty"`
	// Choosing “Capacity” directs Volumez to prefer using capacity-saving methods (such as compression, deduplication, erasure coding and thin provisioning) where relevant, in order to consume the minimum amount of raw media. Using such methods might take some CPU cycles, and might reduce the performance of your volumes (it will still be within the range you specified). Choosing “Balanced” directs Volumez to prefer using some capacity-saving methods where relevant, in order to use less raw media, while consuming a small amount of CPU cycles. “Performance Optimized” directs Volumez to avoid using capacity-saving any methods (such as compression and deduplication) that reduce media consumption. This way applications can get the optimal performance from their media, however more raw media might be consumed to provision Performance-Optimized volumes.
	Capacityoptimization Capacityoptimization `json:"capacityoptimization"`
	// Enter how much logical capacity is reserved up-front for the applications to use. If more capacity is needed for the volume, it will be allocated based on availability of media. Capacities that are reserved can be used for the volume itself and for its snapshots. For example – Use 0% for thin-provisioned volumes, 130% for thick-provisioned volumes with estimated 30% of space for snapshots. Valid values are 0%-500%, default is 20%.
	Capacityreservation *int64 `json:"capacityreservation,omitempty"`
	//  Enter how many media failures (e.g. disk, memory card) the system is required to sustain, and still serve volumes of this policy. A value of “0” means any disk failure will result data unavailability and loss. Valid values are 0..3, default value is 2.
	Resiliencymedia *int64 `json:"resiliencymedia,omitempty"`
	// Enter how many Volumez node (e.g. EC2 instance, server) failures the system is required to sustain, and still serve volumes of this policy. This is different than “Media failures” because sometimes multiple media copies may end on a single node. A value of “0” means any node failure will result data unavailability and loss. Valid values are 0..3, default value is 1.
	Resiliencynode *int64 `json:"resiliencynode,omitempty"`
	// Enter how many zones (e.g. AWS availability zones, DataCenters Buildings) failures the system is required to sustain, and still serve volumes of this policy. Note: zones are assumed to be within the same metro distance, and resiliency to zone failures means cross-zone network traffic. Valid values are 0..2, default value is 1.
	Resiliencyzone *int64 `json:"resiliencyzone,omitempty"`
	// Enter how many regions (e.g. AWS regions zones, DataCenters across continents) failures the system is required to sustain, and still serve volumes of this policy. Note: regions are assumed to reside across WAN distance, with some bandwidth limitations. Valid values are 0 and 1, default value is 0.
	Resiliencyregion *int64 `json:"resiliencyregion,omitempty"`
	// Enter how many seconds are allowed for the replica to stay behind the primary storage. 0 means synchronous replication. Valid values are 0..3600, default value is 0. Max value: 3600. (1 hour).
	Replicationrpo *int64 `json:"replicationrpo,omitempty"`
	// Specifies the maximum bandwidth that Volumez is allowed to consume for replication of this volume (MB/s). 0 means no bandwidth limitation.
	Replicationbandwidth *int64 `json:"replicationbandwidth,omitempty"`
	// Enter “YES” to encrypt the data in server where the application is running. Note: No change is needed in the applications themselves, however encryption will consume some CPU cycles on the application server. Default value NO.
	Encryption *bool `json:"encryption,omitempty"`
	// Enter “YES” to direct Volumez to only use media with disk encryption capabilities. Note that specifying “NO” can still use such media, however it is not a must to use it. Default value: NO.
	Sed *bool `json:"sed,omitempty"`
	// Enter “YES” to direct Volumez to activate the “Device Mapper integrity” protection for the volume. This capability provides strong integrity checking. Note: No change is needed in the applications themselves, however Data Integrity will consume non-negligible CPU cycles on the application server. Default value: NO.
	Integrity          *bool   `json:"integrity,omitempty"`
	Snapshotkeep       *int64  `json:"snapshotkeep,omitempty"`
	Snapshotfrequency  *string `json:"snapshotfrequency,omitempty"`
	Snapshotday        *int64  `json:"snapshotday,omitempty"`
	Snapshothour       *int64  `json:"snapshothour,omitempty"`
	Snapshotminute     *int64  `json:"snapshotminute,omitempty"`
	CreatedbyuserName  *string `json:"createdbyuserName,omitempty"`
	Createdbyuseremail *string `json:"createdbyuseremail,omitempty"`
	Createdtime        *string `json:"createdtime,omitempty"`
	Updatebyusername   *string `json:"updatebyusername,omitempty"`
	UpdatebyUseremail  *string `json:"updatebyUseremail,omitempty"`
	Updatetime         *string `json:"updatetime,omitempty"`
}

func (o *Policy) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Policy) GetIopswrite() *int64 {
	if o == nil {
		return nil
	}
	return o.Iopswrite
}

func (o *Policy) GetIopsread() *int64 {
	if o == nil {
		return nil
	}
	return o.Iopsread
}

func (o *Policy) GetBandwidthwrite() *int64 {
	if o == nil {
		return nil
	}
	return o.Bandwidthwrite
}

func (o *Policy) GetBandwidthread() *int64 {
	if o == nil {
		return nil
	}
	return o.Bandwidthread
}

func (o *Policy) GetLatencywrite() *int64 {
	if o == nil {
		return nil
	}
	return o.Latencywrite
}

func (o *Policy) GetLatencyread() *int64 {
	if o == nil {
		return nil
	}
	return o.Latencyread
}

func (o *Policy) GetLatencyreadcold() *int64 {
	if o == nil {
		return nil
	}
	return o.Latencyreadcold
}

func (o *Policy) GetColddata() *int64 {
	if o == nil {
		return nil
	}
	return o.Colddata
}

func (o *Policy) GetLocalzoneread() *bool {
	if o == nil {
		return nil
	}
	return o.Localzoneread
}

func (o *Policy) GetFailureperformance() *bool {
	if o == nil {
		return nil
	}
	return o.Failureperformance
}

func (o *Policy) GetCapacityoptimization() Capacityoptimization {
	if o == nil {
		return Capacityoptimization("")
	}
	return o.Capacityoptimization
}

func (o *Policy) GetCapacityreservation() *int64 {
	if o == nil {
		return nil
	}
	return o.Capacityreservation
}

func (o *Policy) GetResiliencymedia() *int64 {
	if o == nil {
		return nil
	}
	return o.Resiliencymedia
}

func (o *Policy) GetResiliencynode() *int64 {
	if o == nil {
		return nil
	}
	return o.Resiliencynode
}

func (o *Policy) GetResiliencyzone() *int64 {
	if o == nil {
		return nil
	}
	return o.Resiliencyzone
}

func (o *Policy) GetResiliencyregion() *int64 {
	if o == nil {
		return nil
	}
	return o.Resiliencyregion
}

func (o *Policy) GetReplicationrpo() *int64 {
	if o == nil {
		return nil
	}
	return o.Replicationrpo
}

func (o *Policy) GetReplicationbandwidth() *int64 {
	if o == nil {
		return nil
	}
	return o.Replicationbandwidth
}

func (o *Policy) GetEncryption() *bool {
	if o == nil {
		return nil
	}
	return o.Encryption
}

func (o *Policy) GetSed() *bool {
	if o == nil {
		return nil
	}
	return o.Sed
}

func (o *Policy) GetIntegrity() *bool {
	if o == nil {
		return nil
	}
	return o.Integrity
}

func (o *Policy) GetSnapshotkeep() *int64 {
	if o == nil {
		return nil
	}
	return o.Snapshotkeep
}

func (o *Policy) GetSnapshotfrequency() *string {
	if o == nil {
		return nil
	}
	return o.Snapshotfrequency
}

func (o *Policy) GetSnapshotday() *int64 {
	if o == nil {
		return nil
	}
	return o.Snapshotday
}

func (o *Policy) GetSnapshothour() *int64 {
	if o == nil {
		return nil
	}
	return o.Snapshothour
}

func (o *Policy) GetSnapshotminute() *int64 {
	if o == nil {
		return nil
	}
	return o.Snapshotminute
}

func (o *Policy) GetCreatedbyuserName() *string {
	if o == nil {
		return nil
	}
	return o.CreatedbyuserName
}

func (o *Policy) GetCreatedbyuseremail() *string {
	if o == nil {
		return nil
	}
	return o.Createdbyuseremail
}

func (o *Policy) GetCreatedtime() *string {
	if o == nil {
		return nil
	}
	return o.Createdtime
}

func (o *Policy) GetUpdatebyusername() *string {
	if o == nil {
		return nil
	}
	return o.Updatebyusername
}

func (o *Policy) GetUpdatebyUseremail() *string {
	if o == nil {
		return nil
	}
	return o.UpdatebyUseremail
}

func (o *Policy) GetUpdatetime() *string {
	if o == nil {
		return nil
	}
	return o.Updatetime
}

type PolicyInput struct {
	// A name for the policy. The name can be any non-empty string that does not contain a white space.
	Name string `json:"name"`
	// Enter the maximum write IOPS that a volume is expected to sustain (assuming 8K writes). Write IOPS should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes.
	Iopswrite *int64 `json:"iopswrite,omitempty"`
	// Enter the maximum read IOPS that a volume is expected to sustain (assuming 8K reads). Read IOPS should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes.
	Iopsread *int64 `json:"iopsread,omitempty"`
	// Enter the maximum write bandwidth that a volume is expected to sustain. Write Bandwidth should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes.
	Bandwidthwrite *int64 `json:"bandwidthwrite,omitempty"`
	// Enter the maximum read bandwidth that a volume is expected to sustain. Read Bandwidth should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes.
	Bandwidthread *int64 `json:"bandwidthread,omitempty"`
	// Enter the maximum latency that a volume is expected to sustain. Write latency should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes.
	Latencywrite *int64 `json:"latencywrite,omitempty"`
	// Enter the maximum read IOPS that a volume is expected to sustain. Read latency should be a positive integer number. Volumez will guarantee to provide this performance, regardless of the volume size or other volumes.
	Latencyread *int64 `json:"latencyread,omitempty"`
	// Setting this value to “Yes” directs Volumez to prefer volume configurations where reads are usually happening from disks that are in the same zone as the application. This saves east-west network traffic across zones, however more media per zone will be required to achieve read-IOPs requirements. Set this value to “Yes” if you have network constraints (bandwidth or cost) across your zones; otherwise set it to “No”.
	Localzoneread *bool `json:"localzoneread,omitempty"`
	// Setting this value to “Yes” directs Volumez to over-provision volumes in a way that even after having a failure, the volumes will have the desired performance. Setting this value to “No” directs Volumez to provision volumes according to the desired performance, however in a case of failure – performance may be impacted. The default value is “No”.
	Failureperformance *bool `json:"failureperformance,omitempty"`
	// Choosing “Capacity” directs Volumez to prefer using capacity-saving methods (such as compression, deduplication, erasure coding and thin provisioning) where relevant, in order to consume the minimum amount of raw media. Using such methods might take some CPU cycles, and might reduce the performance of your volumes (it will still be within the range you specified). Choosing “Balanced” directs Volumez to prefer using some capacity-saving methods where relevant, in order to use less raw media, while consuming a small amount of CPU cycles. “Performance Optimized” directs Volumez to avoid using capacity-saving any methods (such as compression and deduplication) that reduce media consumption. This way applications can get the optimal performance from their media, however more raw media might be consumed to provision Performance-Optimized volumes.
	Capacityoptimization Capacityoptimization `json:"capacityoptimization"`
	// Enter how much logical capacity is reserved up-front for the applications to use. If more capacity is needed for the volume, it will be allocated based on availability of media. Capacities that are reserved can be used for the volume itself and for its snapshots. For example – Use 0% for thin-provisioned volumes, 130% for thick-provisioned volumes with estimated 30% of space for snapshots. Valid values are 0%-500%, default is 20%.
	Capacityreservation *int64 `json:"capacityreservation,omitempty"`
	//  Enter how many media failures (e.g. disk, memory card) the system is required to sustain, and still serve volumes of this policy. A value of “0” means any disk failure will result data unavailability and loss. Valid values are 0..3, default value is 2.
	Resiliencymedia *int64 `json:"resiliencymedia,omitempty"`
	// Enter how many Volumez node (e.g. EC2 instance, server) failures the system is required to sustain, and still serve volumes of this policy. This is different than “Media failures” because sometimes multiple media copies may end on a single node. A value of “0” means any node failure will result data unavailability and loss. Valid values are 0..3, default value is 1.
	Resiliencynode *int64 `json:"resiliencynode,omitempty"`
	// Enter how many zones (e.g. AWS availability zones, DataCenters Buildings) failures the system is required to sustain, and still serve volumes of this policy. Note: zones are assumed to be within the same metro distance, and resiliency to zone failures means cross-zone network traffic. Valid values are 0..2, default value is 1.
	Resiliencyzone *int64 `json:"resiliencyzone,omitempty"`
	// Enter how many regions (e.g. AWS regions zones, DataCenters across continents) failures the system is required to sustain, and still serve volumes of this policy. Note: regions are assumed to reside across WAN distance, with some bandwidth limitations. Valid values are 0 and 1, default value is 0.
	Resiliencyregion *int64 `json:"resiliencyregion,omitempty"`
	// Enter how many seconds are allowed for the replica to stay behind the primary storage. 0 means synchronous replication. Valid values are 0..3600, default value is 0. Max value: 3600. (1 hour).
	Replicationrpo *int64 `json:"replicationrpo,omitempty"`
	// Specifies the maximum bandwidth that Volumez is allowed to consume for replication of this volume (MB/s). 0 means no bandwidth limitation.
	Replicationbandwidth *int64 `json:"replicationbandwidth,omitempty"`
	// Enter “YES” to encrypt the data in server where the application is running. Note: No change is needed in the applications themselves, however encryption will consume some CPU cycles on the application server. Default value NO.
	Encryption *bool `json:"encryption,omitempty"`
	// Enter “YES” to direct Volumez to only use media with disk encryption capabilities. Note that specifying “NO” can still use such media, however it is not a must to use it. Default value: NO.
	Sed *bool `json:"sed,omitempty"`
	// Enter “YES” to direct Volumez to activate the “Device Mapper integrity” protection for the volume. This capability provides strong integrity checking. Note: No change is needed in the applications themselves, however Data Integrity will consume non-negligible CPU cycles on the application server. Default value: NO.
	Integrity          *bool   `json:"integrity,omitempty"`
	Snapshotkeep       *int64  `json:"snapshotkeep,omitempty"`
	Snapshotfrequency  *string `json:"snapshotfrequency,omitempty"`
	Snapshotday        *int64  `json:"snapshotday,omitempty"`
	Snapshothour       *int64  `json:"snapshothour,omitempty"`
	Snapshotminute     *int64  `json:"snapshotminute,omitempty"`
	CreatedbyuserName  *string `json:"createdbyuserName,omitempty"`
	Createdbyuseremail *string `json:"createdbyuseremail,omitempty"`
	Createdtime        *string `json:"createdtime,omitempty"`
	Updatebyusername   *string `json:"updatebyusername,omitempty"`
	UpdatebyUseremail  *string `json:"updatebyUseremail,omitempty"`
	Updatetime         *string `json:"updatetime,omitempty"`
}

func (o *PolicyInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PolicyInput) GetIopswrite() *int64 {
	if o == nil {
		return nil
	}
	return o.Iopswrite
}

func (o *PolicyInput) GetIopsread() *int64 {
	if o == nil {
		return nil
	}
	return o.Iopsread
}

func (o *PolicyInput) GetBandwidthwrite() *int64 {
	if o == nil {
		return nil
	}
	return o.Bandwidthwrite
}

func (o *PolicyInput) GetBandwidthread() *int64 {
	if o == nil {
		return nil
	}
	return o.Bandwidthread
}

func (o *PolicyInput) GetLatencywrite() *int64 {
	if o == nil {
		return nil
	}
	return o.Latencywrite
}

func (o *PolicyInput) GetLatencyread() *int64 {
	if o == nil {
		return nil
	}
	return o.Latencyread
}

func (o *PolicyInput) GetLocalzoneread() *bool {
	if o == nil {
		return nil
	}
	return o.Localzoneread
}

func (o *PolicyInput) GetFailureperformance() *bool {
	if o == nil {
		return nil
	}
	return o.Failureperformance
}

func (o *PolicyInput) GetCapacityoptimization() Capacityoptimization {
	if o == nil {
		return Capacityoptimization("")
	}
	return o.Capacityoptimization
}

func (o *PolicyInput) GetCapacityreservation() *int64 {
	if o == nil {
		return nil
	}
	return o.Capacityreservation
}

func (o *PolicyInput) GetResiliencymedia() *int64 {
	if o == nil {
		return nil
	}
	return o.Resiliencymedia
}

func (o *PolicyInput) GetResiliencynode() *int64 {
	if o == nil {
		return nil
	}
	return o.Resiliencynode
}

func (o *PolicyInput) GetResiliencyzone() *int64 {
	if o == nil {
		return nil
	}
	return o.Resiliencyzone
}

func (o *PolicyInput) GetResiliencyregion() *int64 {
	if o == nil {
		return nil
	}
	return o.Resiliencyregion
}

func (o *PolicyInput) GetReplicationrpo() *int64 {
	if o == nil {
		return nil
	}
	return o.Replicationrpo
}

func (o *PolicyInput) GetReplicationbandwidth() *int64 {
	if o == nil {
		return nil
	}
	return o.Replicationbandwidth
}

func (o *PolicyInput) GetEncryption() *bool {
	if o == nil {
		return nil
	}
	return o.Encryption
}

func (o *PolicyInput) GetSed() *bool {
	if o == nil {
		return nil
	}
	return o.Sed
}

func (o *PolicyInput) GetIntegrity() *bool {
	if o == nil {
		return nil
	}
	return o.Integrity
}

func (o *PolicyInput) GetSnapshotkeep() *int64 {
	if o == nil {
		return nil
	}
	return o.Snapshotkeep
}

func (o *PolicyInput) GetSnapshotfrequency() *string {
	if o == nil {
		return nil
	}
	return o.Snapshotfrequency
}

func (o *PolicyInput) GetSnapshotday() *int64 {
	if o == nil {
		return nil
	}
	return o.Snapshotday
}

func (o *PolicyInput) GetSnapshothour() *int64 {
	if o == nil {
		return nil
	}
	return o.Snapshothour
}

func (o *PolicyInput) GetSnapshotminute() *int64 {
	if o == nil {
		return nil
	}
	return o.Snapshotminute
}

func (o *PolicyInput) GetCreatedbyuserName() *string {
	if o == nil {
		return nil
	}
	return o.CreatedbyuserName
}

func (o *PolicyInput) GetCreatedbyuseremail() *string {
	if o == nil {
		return nil
	}
	return o.Createdbyuseremail
}

func (o *PolicyInput) GetCreatedtime() *string {
	if o == nil {
		return nil
	}
	return o.Createdtime
}

func (o *PolicyInput) GetUpdatebyusername() *string {
	if o == nil {
		return nil
	}
	return o.Updatebyusername
}

func (o *PolicyInput) GetUpdatebyUseremail() *string {
	if o == nil {
		return nil
	}
	return o.UpdatebyUseremail
}

func (o *PolicyInput) GetUpdatetime() *string {
	if o == nil {
		return nil
	}
	return o.Updatetime
}
