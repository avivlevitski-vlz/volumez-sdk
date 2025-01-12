# github.com/avivlevitski-vlz/volumez-sdk

Developer-friendly & type-safe Go SDK specifically catered to leverage *github.com/avivlevitski-vlz/volumez-sdk* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=github-com/avivlevitski-vlz/volumez-sdk&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/volumez/saas). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

avivl-test-apigw: Volumez orchestrator API
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [github.com/avivlevitski-vlz/volumez-sdk](#githubcomavivlevitski-vlzvolumez-sdk)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/avivlevitski-vlz/volumez-sdk
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := volumezsdk.New(
		volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
	)

	res, err := s.VolumesList(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.Volumes != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name                  | Type   | Scheme  | Environment Variable               |
| --------------------- | ------ | ------- | ---------------------------------- |
| `StorageIoAuthorizer` | apiKey | API key | `VOLUMEZSDK_STORAGE_IO_AUTHORIZER` |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := volumezsdk.New(
		volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
	)

	res, err := s.VolumesList(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.Volumes != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [VolumezSDK](docs/sdks/volumezsdk/README.md)

* [VolumesList](docs/sdks/volumezsdk/README.md#volumeslist) - Get a list of volumes
* [VolumeCreate](docs/sdks/volumezsdk/README.md#volumecreate) - Create a new volume
* [OptionsVolumes](docs/sdks/volumezsdk/README.md#optionsvolumes)
* [BatchVolumesPlan](docs/sdks/volumezsdk/README.md#batchvolumesplan) - check if volumes can be created
* [OptionsVolumesPlan](docs/sdks/volumezsdk/README.md#optionsvolumesplan)
* [ConsistencyGroupSnapshotCreate](docs/sdks/volumezsdk/README.md#consistencygroupsnapshotcreate) - Create a new snapshot for given consistency group
* [OptionsVolumesSnapshot](docs/sdks/volumezsdk/README.md#optionsvolumessnapshot)
* [ConsistencyGroupGet](docs/sdks/volumezsdk/README.md#consistencygroupget) - List of snapshots group
* [OptionsVolumesSnapshotSnapshotGroupName](docs/sdks/volumezsdk/README.md#optionsvolumessnapshotsnapshotgroupname)
* [VolumeGet](docs/sdks/volumezsdk/README.md#volumeget) - Get the properties of a volume
* [VolumeDelete](docs/sdks/volumezsdk/README.md#volumedelete) - Delete a volume
* [OptionsVolumesVolume](docs/sdks/volumezsdk/README.md#optionsvolumesvolume)
* [VolumeModify](docs/sdks/volumezsdk/README.md#volumemodify) - Modify a volume
* [VolumeDescribe](docs/sdks/volumezsdk/README.md#volumedescribe) - describe existing volume 
* [OptionsVolumesVolumeDescribe](docs/sdks/volumezsdk/README.md#optionsvolumesvolumedescribe)
* [VolumeRecoverInitiate](docs/sdks/volumezsdk/README.md#volumerecoverinitiate) - Initiate recover on volume
* [OptionsVolumesVolumeRecover](docs/sdks/volumezsdk/README.md#optionsvolumesvolumerecover)
* [SnapshotsList](docs/sdks/volumezsdk/README.md#snapshotslist) - Get a list of snapshots
* [SnapshotCreate](docs/sdks/volumezsdk/README.md#snapshotcreate) - Create a new snapshot
* [OptionsVolumesVolumeSnapshots](docs/sdks/volumezsdk/README.md#optionsvolumesvolumesnapshots)
* [SnapshotGet](docs/sdks/volumezsdk/README.md#snapshotget) - Get the properties of a snapshot
* [SnapshotDelete](docs/sdks/volumezsdk/README.md#snapshotdelete) - Delete a snapshot
* [OptionsVolumesVolumeSnapshotsSnapshot](docs/sdks/volumezsdk/README.md#optionsvolumesvolumesnapshotssnapshot)
* [SnapshotModify](docs/sdks/volumezsdk/README.md#snapshotmodify) - Modify a snapshot
* [OptionsVolumesVolumeSnapshotsSnapshotRollback](docs/sdks/volumezsdk/README.md#optionsvolumesvolumesnapshotssnapshotrollback)
* [SnapshotRollback](docs/sdks/volumezsdk/README.md#snapshotrollback) - Roll back to snapshot
* [AssociationsList](docs/sdks/volumezsdk/README.md#associationslist) - Get a list of associations
* [AssociationCreate](docs/sdks/volumezsdk/README.md#associationcreate) - Create association
* [OptionsAssociations](docs/sdks/volumezsdk/README.md#optionsassociations)
* [AssociationDelete](docs/sdks/volumezsdk/README.md#associationdelete) - Delete an association
* [OptionsAssociationsAssociation](docs/sdks/volumezsdk/README.md#optionsassociationsassociation)
* [AssociationModify](docs/sdks/volumezsdk/README.md#associationmodify) - Association modify
* [ExportsList](docs/sdks/volumezsdk/README.md#exportslist) - Get a list of associations
* [ExportCreate](docs/sdks/volumezsdk/README.md#exportcreate) - Create export
* [OptionsExports](docs/sdks/volumezsdk/README.md#optionsexports)
* [ExportDelete](docs/sdks/volumezsdk/README.md#exportdelete) - Delete an export
* [OptionsExportsExport](docs/sdks/volumezsdk/README.md#optionsexportsexport)
* [AutoProvisionVolumes](docs/sdks/volumezsdk/README.md#autoprovisionvolumes) - Create a new auto provisioned volume
* [OptionsAutoprovisionvolumes](docs/sdks/volumezsdk/README.md#optionsautoprovisionvolumes)
* [SnapshotsListAll](docs/sdks/volumezsdk/README.md#snapshotslistall) - Get a list of all snapshots
* [OptionsSnapshots](docs/sdks/volumezsdk/README.md#optionssnapshots)
* [AttachmentsListAll](docs/sdks/volumezsdk/README.md#attachmentslistall) - Get a list of all attachments
* [OptionsAttachments](docs/sdks/volumezsdk/README.md#optionsattachments)
* [AttachmentsListForVolume](docs/sdks/volumezsdk/README.md#attachmentslistforvolume) - Get a list of attachments for a given volume
* [OptionsVolumesVolumeAttachments](docs/sdks/volumezsdk/README.md#optionsvolumesvolumeattachments)
* [AttachmentsList](docs/sdks/volumezsdk/README.md#attachmentslist) - Get a list of attachments for a given volume and snapshot
* [AttachmentCreate](docs/sdks/volumezsdk/README.md#attachmentcreate) - Create a new attachment
* [OptionsVolumesVolumeSnapshotsSnapshotAttachments](docs/sdks/volumezsdk/README.md#optionsvolumesvolumesnapshotssnapshotattachments)
* [AttachmentGet](docs/sdks/volumezsdk/README.md#attachmentget) - Get the properties of an attachment
* [AttachmentDelete](docs/sdks/volumezsdk/README.md#attachmentdelete) - Delete an attachment
* [OptionsVolumesVolumeSnapshotsSnapshotAttachmentsNode](docs/sdks/volumezsdk/README.md#optionsvolumesvolumesnapshotssnapshotattachmentsnode)
* [AttachmentModify](docs/sdks/volumezsdk/README.md#attachmentmodify) - Modify an attachment
* [NodesList](docs/sdks/volumezsdk/README.md#nodeslist) - Get a list of nodes
* [OptionsNodes](docs/sdks/volumezsdk/README.md#optionsnodes)
* [OptionsNodesTagsNode](docs/sdks/volumezsdk/README.md#optionsnodestagsnode)
* [NodeSetTags](docs/sdks/volumezsdk/README.md#nodesettags) - Set the tags of a node
* [NodeGet](docs/sdks/volumezsdk/README.md#nodeget) - Get the properties of a node
* [NodeDelete](docs/sdks/volumezsdk/README.md#nodedelete) - Delete a node
* [OptionsNodesNode](docs/sdks/volumezsdk/README.md#optionsnodesnode)
* [NodeDrain](docs/sdks/volumezsdk/README.md#nodedrain) - Node drain
* [OptionsNodesNodeDrain](docs/sdks/volumezsdk/README.md#optionsnodesnodedrain)
* [NodeDescribe](docs/sdks/volumezsdk/README.md#nodedescribe) - Node describe
* [OptionsNodesNodeDescribe](docs/sdks/volumezsdk/README.md#optionsnodesnodedescribe)
* [NodeCollectLogs](docs/sdks/volumezsdk/README.md#nodecollectlogs) - Node collect logs
* [OptionsNodesLogsNodeTenant](docs/sdks/volumezsdk/README.md#optionsnodeslogsnodetenant)
* [NodeRepair](docs/sdks/volumezsdk/README.md#noderepair) - Execute commands on node for repair
* [OptionsNodesRepairNodeTenant](docs/sdks/volumezsdk/README.md#optionsnodesrepairnodetenant)
* [NodeUpgrade](docs/sdks/volumezsdk/README.md#nodeupgrade) - performing node version upgrade
* [OptionsNodesUpgradeNode](docs/sdks/volumezsdk/README.md#optionsnodesupgradenode)
* [NodeUpgradeForSupport](docs/sdks/volumezsdk/README.md#nodeupgradeforsupport) - performing node version upgrade
* [OptionsNodesUpgradeNodeTenantTenant](docs/sdks/volumezsdk/README.md#optionsnodesupgradenodetenanttenant)
* [PoliciesList](docs/sdks/volumezsdk/README.md#policieslist) - Get a list of policies
* [PolicyCreate](docs/sdks/volumezsdk/README.md#policycreate) - Create a new policy
* [OptionsPolicies](docs/sdks/volumezsdk/README.md#optionspolicies)
* [PolicyGet](docs/sdks/volumezsdk/README.md#policyget) - Get the properties of a policy
* [PolicyDelete](docs/sdks/volumezsdk/README.md#policydelete) - Delete a policy
* [OptionsPoliciesPolicy](docs/sdks/volumezsdk/README.md#optionspoliciespolicy)
* [PolicyModify](docs/sdks/volumezsdk/README.md#policymodify) - Modify a policy
* [PolicyGetVolumes](docs/sdks/volumezsdk/README.md#policygetvolumes) - Get the properties of a policy
* [OptionsPoliciesPolicyVolumes](docs/sdks/volumezsdk/README.md#optionspoliciespolicyvolumes)
* [PolicyPlan](docs/sdks/volumezsdk/README.md#policyplan) - Show policy volume create plan
* [OptionsPoliciesPolicySizeSizeZoneZone](docs/sdks/volumezsdk/README.md#optionspoliciespolicysizesizezonezone)
* [NetworksList](docs/sdks/volumezsdk/README.md#networkslist) - Get a list of networks
* [NetworkCreate](docs/sdks/volumezsdk/README.md#networkcreate) - Create a new network
* [OptionsNetworks](docs/sdks/volumezsdk/README.md#optionsnetworks)
* [NetworkGet](docs/sdks/volumezsdk/README.md#networkget) - Get the properties of a network
* [NetworkDelete](docs/sdks/volumezsdk/README.md#networkdelete) - Delete a network
* [OptionsNetworksNetwork](docs/sdks/volumezsdk/README.md#optionsnetworksnetwork)
* [NetworkModify](docs/sdks/volumezsdk/README.md#networkmodify) - Modify a network
* [JobsList](docs/sdks/volumezsdk/README.md#jobslist) - Get a list of jobs
* [OptionsJobs](docs/sdks/volumezsdk/README.md#optionsjobs)
* [JobGet](docs/sdks/volumezsdk/README.md#jobget) - Get the properties of a job
* [JobDelete](docs/sdks/volumezsdk/README.md#jobdelete) - Delete a job
* [OptionsJobsJob](docs/sdks/volumezsdk/README.md#optionsjobsjob)
* [OptionsJobsJobResumeSuspendState](docs/sdks/volumezsdk/README.md#optionsjobsjobresumesuspendstate)
* [JobResumeSuspend](docs/sdks/volumezsdk/README.md#jobresumesuspend) - Resume or Suspend a job
* [MediaList](docs/sdks/volumezsdk/README.md#medialist) - Get a list of media
* [OptionsMedia](docs/sdks/volumezsdk/README.md#optionsmedia)
* [MediaGet](docs/sdks/volumezsdk/README.md#mediaget) - Get the properties of a media
* [MediaDelete](docs/sdks/volumezsdk/README.md#mediadelete) - Delete a media
* [OptionsMediaMedia](docs/sdks/volumezsdk/README.md#optionsmediamedia)
* [MediaModify](docs/sdks/volumezsdk/README.md#mediamodify) - modify a media properties
* [MediaAssignLegacy](docs/sdks/volumezsdk/README.md#mediaassignlegacy) - Assign media
* [OptionsMediaMediaAssign](docs/sdks/volumezsdk/README.md#optionsmediamediaassign)
* [MediaAssign](docs/sdks/volumezsdk/README.md#mediaassign) - Assign media
* [MediaUnassignLegacy](docs/sdks/volumezsdk/README.md#mediaunassignlegacy) - Unassign media
* [OptionsMediaMediaUnassign](docs/sdks/volumezsdk/README.md#optionsmediamediaunassign)
* [MediaUnassign](docs/sdks/volumezsdk/README.md#mediaunassign) - Unassign media
* [OptionsMediaMediaProfile](docs/sdks/volumezsdk/README.md#optionsmediamediaprofile)
* [MediaProfileModify](docs/sdks/volumezsdk/README.md#mediaprofilemodify) - Modify a media profile
* [MediaDrain](docs/sdks/volumezsdk/README.md#mediadrain) - Media drain
* [OptionsMediaMediaDrain](docs/sdks/volumezsdk/README.md#optionsmediamediadrain)
* [MediaDiagnose](docs/sdks/volumezsdk/README.md#mediadiagnose) - Media diagnose
* [OptionsMediaMediaTenantDiagnose](docs/sdks/volumezsdk/README.md#optionsmediamediatenantdiagnose)
* [VirtualMediaList](docs/sdks/volumezsdk/README.md#virtualmedialist) - Get a list of virtual media
* [VirtualMediaCreate](docs/sdks/volumezsdk/README.md#virtualmediacreate) - Create (virtual) media
* [OptionsVirtualmedia](docs/sdks/volumezsdk/README.md#optionsvirtualmedia)
* [VirtualMediaDelete](docs/sdks/volumezsdk/README.md#virtualmediadelete) - Delete virtual media
* [OptionsVirtualmediaMedia](docs/sdks/volumezsdk/README.md#optionsvirtualmediamedia)
* [ConnectivitiesList](docs/sdks/volumezsdk/README.md#connectivitieslist) - Get a list of connectivities
* [ConnectivityCreate](docs/sdks/volumezsdk/README.md#connectivitycreate) - Create a new connectivity
* [OptionsConnectivities](docs/sdks/volumezsdk/README.md#optionsconnectivities)
* [ConnectivityGet](docs/sdks/volumezsdk/README.md#connectivityget) - Get the properties of a connectivity
* [ConnectivityDelete](docs/sdks/volumezsdk/README.md#connectivitydelete) - Delete a connectivity
* [OptionsConnectivitiesConnectivity](docs/sdks/volumezsdk/README.md#optionsconnectivitiesconnectivity)
* [ConnectivityModify](docs/sdks/volumezsdk/README.md#connectivitymodify) - Modify a connectivity
* [CapacityGroupGet](docs/sdks/volumezsdk/README.md#capacitygroupget) - Get available capacity groups
* [OptionsCapacitygroups](docs/sdks/volumezsdk/README.md#optionscapacitygroups)
* [UserCreate](docs/sdks/volumezsdk/README.md#usercreate) - Create Tenant's first user
* [OptionsSignup](docs/sdks/volumezsdk/README.md#optionssignup)
* [Signout](docs/sdks/volumezsdk/README.md#signout) - Signs out user from all devices
* [OptionsSignout](docs/sdks/volumezsdk/README.md#optionssignout)
* [SignIn](docs/sdks/volumezsdk/README.md#signin) - SIO Sign In
* [OptionsSignin](docs/sdks/volumezsdk/README.md#optionssignin)
* [Tenanthostdelete](docs/sdks/volumezsdk/README.md#tenanthostdelete) - Delete a tenant host
* [OptionsTenantTenantIDTenanthostsTenantHost](docs/sdks/volumezsdk/README.md#optionstenanttenantidtenanthoststenanthost)
* [TenanthostAccessCredentials](docs/sdks/volumezsdk/README.md#tenanthostaccesscredentials) - Given the tenant's host token (Refresh Token), get the ID Token for this machine
* [OptionsTenantTenanthostCredentials](docs/sdks/volumezsdk/README.md#optionstenanttenanthostcredentials)
* [Tenanthostget](docs/sdks/volumezsdk/README.md#tenanthostget) - Get a tenant host
* [OptionsTenantTenanthost](docs/sdks/volumezsdk/README.md#optionstenanttenanthost)
* [TenantIDget](docs/sdks/volumezsdk/README.md#tenantidget) - Get a Tenant ID from Tenant's Token
* [OptionsTenantTenantid](docs/sdks/volumezsdk/README.md#optionstenanttenantid)
* [TenantRefreshToken](docs/sdks/volumezsdk/README.md#tenantrefreshtoken) - Get the tenant's Refresh Token
* [OptionsTenantRefreshtoken](docs/sdks/volumezsdk/README.md#optionstenantrefreshtoken)
* [TenantAPIAccessRefreshToken](docs/sdks/volumezsdk/README.md#tenantapiaccessrefreshtoken) - Given the Tenant's API Access Refresh Token, return new Access Token, Identity Token and a newer Refresh Token
* [OptionsTenantApiaccessCredentialsRefresh](docs/sdks/volumezsdk/README.md#optionstenantapiaccesscredentialsrefresh)
* [TenantToken](docs/sdks/volumezsdk/README.md#tenanttoken) - Get the Tenant's Token
* [OptionsTenantToken](docs/sdks/volumezsdk/README.md#optionstenanttoken)
* [Tenantuserget](docs/sdks/volumezsdk/README.md#tenantuserget) - Get tenant user
* [TenantUserCreate](docs/sdks/volumezsdk/README.md#tenantusercreate) - Create Tenant's additional user
* [OptionsTenantUser](docs/sdks/volumezsdk/README.md#optionstenantuser)
* [UserConfirm](docs/sdks/volumezsdk/README.md#userconfirm) - Confirm user signup
* [OptionsTenantUserConfirmation](docs/sdks/volumezsdk/README.md#optionstenantuserconfirmation)
* [RequestChangePassword](docs/sdks/volumezsdk/README.md#requestchangepassword) - Request Change clients password
* [OptionsTenantUserRequestchangepassword](docs/sdks/volumezsdk/README.md#optionstenantuserrequestchangepassword)
* [ChangePasswordLoggedIn](docs/sdks/volumezsdk/README.md#changepasswordloggedin) - Change clients password when user is logged in
* [OptionsTenantUserChangepasswordloggedin](docs/sdks/volumezsdk/README.md#optionstenantuserchangepasswordloggedin)
* [ChangePassword](docs/sdks/volumezsdk/README.md#changepassword) - Change clients password
* [OptionsTenantUserChangepassword](docs/sdks/volumezsdk/README.md#optionstenantuserchangepassword)
* [VersionGet](docs/sdks/volumezsdk/README.md#versionget) - Get version of sio
* [OptionsVersion](docs/sdks/volumezsdk/README.md#optionsversion)
* [GetSubscription](docs/sdks/volumezsdk/README.md#getsubscription) - get subscription
* [OptionsAzuremarketplaceSubscription](docs/sdks/volumezsdk/README.md#optionsazuremarketplacesubscription)
* [AlertsList](docs/sdks/volumezsdk/README.md#alertslist) - Get a list of alerts
* [OptionsAlerts](docs/sdks/volumezsdk/README.md#optionsalerts)
* [AlertAcknowledge](docs/sdks/volumezsdk/README.md#alertacknowledge) - Acknowledge a pending alert
* [OptionsAlertsAlertAcknowledge](docs/sdks/volumezsdk/README.md#optionsalertsalertacknowledge)
* [GetMachineInfo](docs/sdks/volumezsdk/README.md#getmachineinfo) - get system info
* [OptionsSystemMachineinfo](docs/sdks/volumezsdk/README.md#optionssystemmachineinfo)
* [OptionsCharts](docs/sdks/volumezsdk/README.md#optionscharts)
* [OptionsChartsProxyPlus](docs/sdks/volumezsdk/README.md#optionschartsproxyplus)
* [GetLogTenantidLogfile](docs/sdks/volumezsdk/README.md#getlogtenantidlogfile)
* [OptionsLogTenantidLogfile](docs/sdks/volumezsdk/README.md#optionslogtenantidlogfile)

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/retry"
	"log"
	"models/operations"
	"os"
)

func main() {
	ctx := context.Background()

	s := volumezsdk.New(
		volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
	)

	res, err := s.VolumesList(ctx, nil, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.Volumes != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/retry"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := volumezsdk.New(
		volumezsdk.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
	)

	res, err := s.VolumesList(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.Volumes != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `VolumesList` function may return the following errors:

| Error Type              | Status Code   | Content Type     |
| ----------------------- | ------------- | ---------------- |
| apierrors.ErrorResponse | 400, 404, 500 | application/json |
| apierrors.APIError      | 4XX, 5XX      | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/apierrors"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := volumezsdk.New(
		volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
	)

	res, err := s.VolumesList(ctx, nil)
	if err != nil {

		var e *apierrors.ErrorResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := volumezsdk.New(
		volumezsdk.WithServerURL("https://api.dev.volumez.com"),
		volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
	)

	res, err := s.VolumesList(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.Volumes != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/avivlevitski-vlz/volumez-sdk&utm_campaign=go)
