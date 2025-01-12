# VolumezSDK

## Overview

avivl-test-apigw: Volumez orchestrator API

### Available Operations

* [VolumesList](#volumeslist) - Get a list of volumes
* [VolumeCreate](#volumecreate) - Create a new volume
* [OptionsVolumes](#optionsvolumes)
* [BatchVolumesPlan](#batchvolumesplan) - check if volumes can be created
* [OptionsVolumesPlan](#optionsvolumesplan)
* [ConsistencyGroupSnapshotCreate](#consistencygroupsnapshotcreate) - Create a new snapshot for given consistency group
* [OptionsVolumesSnapshot](#optionsvolumessnapshot)
* [ConsistencyGroupGet](#consistencygroupget) - List of snapshots group
* [OptionsVolumesSnapshotSnapshotGroupName](#optionsvolumessnapshotsnapshotgroupname)
* [VolumeGet](#volumeget) - Get the properties of a volume
* [VolumeDelete](#volumedelete) - Delete a volume
* [OptionsVolumesVolume](#optionsvolumesvolume)
* [VolumeModify](#volumemodify) - Modify a volume
* [VolumeDescribe](#volumedescribe) - describe existing volume 
* [OptionsVolumesVolumeDescribe](#optionsvolumesvolumedescribe)
* [VolumeRecoverInitiate](#volumerecoverinitiate) - Initiate recover on volume
* [OptionsVolumesVolumeRecover](#optionsvolumesvolumerecover)
* [SnapshotsList](#snapshotslist) - Get a list of snapshots
* [SnapshotCreate](#snapshotcreate) - Create a new snapshot
* [OptionsVolumesVolumeSnapshots](#optionsvolumesvolumesnapshots)
* [SnapshotGet](#snapshotget) - Get the properties of a snapshot
* [SnapshotDelete](#snapshotdelete) - Delete a snapshot
* [OptionsVolumesVolumeSnapshotsSnapshot](#optionsvolumesvolumesnapshotssnapshot)
* [SnapshotModify](#snapshotmodify) - Modify a snapshot
* [OptionsVolumesVolumeSnapshotsSnapshotRollback](#optionsvolumesvolumesnapshotssnapshotrollback)
* [SnapshotRollback](#snapshotrollback) - Roll back to snapshot
* [AssociationsList](#associationslist) - Get a list of associations
* [AssociationCreate](#associationcreate) - Create association
* [OptionsAssociations](#optionsassociations)
* [AssociationDelete](#associationdelete) - Delete an association
* [OptionsAssociationsAssociation](#optionsassociationsassociation)
* [AssociationModify](#associationmodify) - Association modify
* [ExportsList](#exportslist) - Get a list of associations
* [ExportCreate](#exportcreate) - Create export
* [OptionsExports](#optionsexports)
* [ExportDelete](#exportdelete) - Delete an export
* [OptionsExportsExport](#optionsexportsexport)
* [AutoProvisionVolumes](#autoprovisionvolumes) - Create a new auto provisioned volume
* [OptionsAutoprovisionvolumes](#optionsautoprovisionvolumes)
* [SnapshotsListAll](#snapshotslistall) - Get a list of all snapshots
* [OptionsSnapshots](#optionssnapshots)
* [AttachmentsListAll](#attachmentslistall) - Get a list of all attachments
* [OptionsAttachments](#optionsattachments)
* [AttachmentsListForVolume](#attachmentslistforvolume) - Get a list of attachments for a given volume
* [OptionsVolumesVolumeAttachments](#optionsvolumesvolumeattachments)
* [AttachmentsList](#attachmentslist) - Get a list of attachments for a given volume and snapshot
* [AttachmentCreate](#attachmentcreate) - Create a new attachment
* [OptionsVolumesVolumeSnapshotsSnapshotAttachments](#optionsvolumesvolumesnapshotssnapshotattachments)
* [AttachmentGet](#attachmentget) - Get the properties of an attachment
* [AttachmentDelete](#attachmentdelete) - Delete an attachment
* [OptionsVolumesVolumeSnapshotsSnapshotAttachmentsNode](#optionsvolumesvolumesnapshotssnapshotattachmentsnode)
* [AttachmentModify](#attachmentmodify) - Modify an attachment
* [NodesList](#nodeslist) - Get a list of nodes
* [OptionsNodes](#optionsnodes)
* [OptionsNodesTagsNode](#optionsnodestagsnode)
* [NodeSetTags](#nodesettags) - Set the tags of a node
* [NodeGet](#nodeget) - Get the properties of a node
* [NodeDelete](#nodedelete) - Delete a node
* [OptionsNodesNode](#optionsnodesnode)
* [NodeDrain](#nodedrain) - Node drain
* [OptionsNodesNodeDrain](#optionsnodesnodedrain)
* [NodeDescribe](#nodedescribe) - Node describe
* [OptionsNodesNodeDescribe](#optionsnodesnodedescribe)
* [NodeCollectLogs](#nodecollectlogs) - Node collect logs
* [OptionsNodesLogsNodeTenant](#optionsnodeslogsnodetenant)
* [NodeRepair](#noderepair) - Execute commands on node for repair
* [OptionsNodesRepairNodeTenant](#optionsnodesrepairnodetenant)
* [NodeUpgrade](#nodeupgrade) - performing node version upgrade
* [OptionsNodesUpgradeNode](#optionsnodesupgradenode)
* [NodeUpgradeForSupport](#nodeupgradeforsupport) - performing node version upgrade
* [OptionsNodesUpgradeNodeTenantTenant](#optionsnodesupgradenodetenanttenant)
* [PoliciesList](#policieslist) - Get a list of policies
* [PolicyCreate](#policycreate) - Create a new policy
* [OptionsPolicies](#optionspolicies)
* [PolicyGet](#policyget) - Get the properties of a policy
* [PolicyDelete](#policydelete) - Delete a policy
* [OptionsPoliciesPolicy](#optionspoliciespolicy)
* [PolicyModify](#policymodify) - Modify a policy
* [PolicyGetVolumes](#policygetvolumes) - Get the properties of a policy
* [OptionsPoliciesPolicyVolumes](#optionspoliciespolicyvolumes)
* [PolicyPlan](#policyplan) - Show policy volume create plan
* [OptionsPoliciesPolicySizeSizeZoneZone](#optionspoliciespolicysizesizezonezone)
* [NetworksList](#networkslist) - Get a list of networks
* [NetworkCreate](#networkcreate) - Create a new network
* [OptionsNetworks](#optionsnetworks)
* [NetworkGet](#networkget) - Get the properties of a network
* [NetworkDelete](#networkdelete) - Delete a network
* [OptionsNetworksNetwork](#optionsnetworksnetwork)
* [NetworkModify](#networkmodify) - Modify a network
* [JobsList](#jobslist) - Get a list of jobs
* [OptionsJobs](#optionsjobs)
* [JobGet](#jobget) - Get the properties of a job
* [JobDelete](#jobdelete) - Delete a job
* [OptionsJobsJob](#optionsjobsjob)
* [OptionsJobsJobResumeSuspendState](#optionsjobsjobresumesuspendstate)
* [JobResumeSuspend](#jobresumesuspend) - Resume or Suspend a job
* [MediaList](#medialist) - Get a list of media
* [OptionsMedia](#optionsmedia)
* [MediaGet](#mediaget) - Get the properties of a media
* [MediaDelete](#mediadelete) - Delete a media
* [OptionsMediaMedia](#optionsmediamedia)
* [MediaModify](#mediamodify) - modify a media properties
* [MediaAssignLegacy](#mediaassignlegacy) - Assign media
* [OptionsMediaMediaAssign](#optionsmediamediaassign)
* [MediaAssign](#mediaassign) - Assign media
* [MediaUnassignLegacy](#mediaunassignlegacy) - Unassign media
* [OptionsMediaMediaUnassign](#optionsmediamediaunassign)
* [MediaUnassign](#mediaunassign) - Unassign media
* [OptionsMediaMediaProfile](#optionsmediamediaprofile)
* [MediaProfileModify](#mediaprofilemodify) - Modify a media profile
* [MediaDrain](#mediadrain) - Media drain
* [OptionsMediaMediaDrain](#optionsmediamediadrain)
* [MediaDiagnose](#mediadiagnose) - Media diagnose
* [OptionsMediaMediaTenantDiagnose](#optionsmediamediatenantdiagnose)
* [VirtualMediaList](#virtualmedialist) - Get a list of virtual media
* [VirtualMediaCreate](#virtualmediacreate) - Create (virtual) media
* [OptionsVirtualmedia](#optionsvirtualmedia)
* [VirtualMediaDelete](#virtualmediadelete) - Delete virtual media
* [OptionsVirtualmediaMedia](#optionsvirtualmediamedia)
* [ConnectivitiesList](#connectivitieslist) - Get a list of connectivities
* [ConnectivityCreate](#connectivitycreate) - Create a new connectivity
* [OptionsConnectivities](#optionsconnectivities)
* [ConnectivityGet](#connectivityget) - Get the properties of a connectivity
* [ConnectivityDelete](#connectivitydelete) - Delete a connectivity
* [OptionsConnectivitiesConnectivity](#optionsconnectivitiesconnectivity)
* [ConnectivityModify](#connectivitymodify) - Modify a connectivity
* [CapacityGroupGet](#capacitygroupget) - Get available capacity groups
* [OptionsCapacitygroups](#optionscapacitygroups)
* [UserCreate](#usercreate) - Create Tenant's first user
* [OptionsSignup](#optionssignup)
* [Signout](#signout) - Signs out user from all devices
* [OptionsSignout](#optionssignout)
* [SignIn](#signin) - SIO Sign In
* [OptionsSignin](#optionssignin)
* [Tenanthostdelete](#tenanthostdelete) - Delete a tenant host
* [OptionsTenantTenantIDTenanthostsTenantHost](#optionstenanttenantidtenanthoststenanthost)
* [TenanthostAccessCredentials](#tenanthostaccesscredentials) - Given the tenant's host token (Refresh Token), get the ID Token for this machine
* [OptionsTenantTenanthostCredentials](#optionstenanttenanthostcredentials)
* [Tenanthostget](#tenanthostget) - Get a tenant host
* [OptionsTenantTenanthost](#optionstenanttenanthost)
* [TenantIDget](#tenantidget) - Get a Tenant ID from Tenant's Token
* [OptionsTenantTenantid](#optionstenanttenantid)
* [TenantRefreshToken](#tenantrefreshtoken) - Get the tenant's Refresh Token
* [OptionsTenantRefreshtoken](#optionstenantrefreshtoken)
* [TenantAPIAccessRefreshToken](#tenantapiaccessrefreshtoken) - Given the Tenant's API Access Refresh Token, return new Access Token, Identity Token and a newer Refresh Token
* [OptionsTenantApiaccessCredentialsRefresh](#optionstenantapiaccesscredentialsrefresh)
* [TenantToken](#tenanttoken) - Get the Tenant's Token
* [OptionsTenantToken](#optionstenanttoken)
* [Tenantuserget](#tenantuserget) - Get tenant user
* [TenantUserCreate](#tenantusercreate) - Create Tenant's additional user
* [OptionsTenantUser](#optionstenantuser)
* [UserConfirm](#userconfirm) - Confirm user signup
* [OptionsTenantUserConfirmation](#optionstenantuserconfirmation)
* [RequestChangePassword](#requestchangepassword) - Request Change clients password
* [OptionsTenantUserRequestchangepassword](#optionstenantuserrequestchangepassword)
* [ChangePasswordLoggedIn](#changepasswordloggedin) - Change clients password when user is logged in
* [OptionsTenantUserChangepasswordloggedin](#optionstenantuserchangepasswordloggedin)
* [ChangePassword](#changepassword) - Change clients password
* [OptionsTenantUserChangepassword](#optionstenantuserchangepassword)
* [VersionGet](#versionget) - Get version of sio
* [OptionsVersion](#optionsversion)
* [GetSubscription](#getsubscription) - get subscription
* [OptionsAzuremarketplaceSubscription](#optionsazuremarketplacesubscription)
* [AlertsList](#alertslist) - Get a list of alerts
* [OptionsAlerts](#optionsalerts)
* [AlertAcknowledge](#alertacknowledge) - Acknowledge a pending alert
* [OptionsAlertsAlertAcknowledge](#optionsalertsalertacknowledge)
* [GetMachineInfo](#getmachineinfo) - get system info
* [OptionsSystemMachineinfo](#optionssystemmachineinfo)
* [OptionsCharts](#optionscharts)
* [OptionsChartsProxyPlus](#optionschartsproxyplus)
* [GetLogTenantidLogfile](#getlogtenantidlogfile)
* [OptionsLogTenantidLogfile](#optionslogtenantidlogfile)

## VolumesList

Get a list of volumes

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
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

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `capacity`                                               | **bool*                                                  | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.VolumesListResponse](../../models/operations/volumeslistresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## VolumeCreate

Create a new volume

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.VolumeCreate(ctx, components.VolumeInput{
        Name: "<value>",
        Type: components.TypeFile,
        Size: 843324,
        Policy: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [components.VolumeInput](../../models/components/volumeinput.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*operations.VolumeCreateResponse](../../models/operations/volumecreateresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsVolumes

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsVolumes(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsVolumesResponse](../../models/operations/optionsvolumesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## BatchVolumesPlan

check if volumes can be created

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.BatchVolumesPlan(ctx, operations.BatchVolumesPlanRequestBody{
        Volumes: []operations.Volumes{
            operations.Volumes{
                Size: 10,
                Zone: volumezsdk.String("us-east-1a"),
            },
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.VolumePlanOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `requestBody`                                                                                    | [operations.BatchVolumesPlanRequestBody](../../models/operations/batchvolumesplanrequestbody.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `verbose`                                                                                        | **bool*                                                                                          | :heavy_minus_sign:                                                                               | if true will return the volume plan if false will omit the plan from the response                |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.BatchVolumesPlanResponse](../../models/operations/batchvolumesplanresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsVolumesPlan

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsVolumesPlan(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsVolumesPlanResponse](../../models/operations/optionsvolumesplanresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ConsistencyGroupSnapshotCreate

Create a new snapshot for given consistency group

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.ConsistencyGroupSnapshotCreate(ctx, operations.ConsistencyGroupSnapshotCreateRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.ConsistencyGroupSnapshotCreateRequestBody](../../models/operations/consistencygroupsnapshotcreaterequestbody.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.ConsistencyGroupSnapshotCreateResponse](../../models/operations/consistencygroupsnapshotcreateresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsVolumesSnapshot

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsVolumesSnapshot(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsVolumesSnapshotResponse](../../models/operations/optionsvolumessnapshotresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ConsistencyGroupGet

List of snapshots group

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.ConsistencyGroupGet(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Snapshots != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `snapshotGroupName`                                      | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ConsistencyGroupGetResponse](../../models/operations/consistencygroupgetresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsVolumesSnapshotSnapshotGroupName

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsVolumesSnapshotSnapshotGroupName(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `snapshotGroupName`                                      | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsVolumesSnapshotSnapshotGroupNameResponse](../../models/operations/optionsvolumessnapshotsnapshotgroupnameresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## VolumeGet

Get the properties of a volume

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.VolumeGet(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Volume != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `volume`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.VolumeGetResponse](../../models/operations/volumegetresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## VolumeDelete

Delete a volume

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.VolumeDelete(ctx, "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `volume`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `force`                                                  | **bool*                                                  | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.VolumeDeleteResponse](../../models/operations/volumedeleteresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsVolumesVolume

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsVolumesVolume(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `volume`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsVolumesVolumeResponse](../../models/operations/optionsvolumesvolumeresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## VolumeModify

Modify a volume

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.VolumeModify(ctx, "<value>", components.VolumeInput{
        Name: "<value>",
        Type: components.TypeBlock,
        Size: 446700,
        Policy: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `volumePathParameter`                                            | *string*                                                         | :heavy_check_mark:                                               | N/A                                                              |
| `volume1`                                                        | [components.VolumeInput](../../models/components/volumeinput.md) | :heavy_check_mark:                                               | A Volume object                                                  |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*operations.VolumeModifyResponse](../../models/operations/volumemodifyresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## VolumeDescribe

describe existing volume 

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.VolumeDescribe(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.VolumeGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `volume`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.VolumeDescribeResponse](../../models/operations/volumedescriberesponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsVolumesVolumeDescribe

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsVolumesVolumeDescribe(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `volume`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsVolumesVolumeDescribeResponse](../../models/operations/optionsvolumesvolumedescriberesponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## VolumeRecoverInitiate

Initiate recover on volume

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.VolumeRecoverInitiate(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `volume`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.VolumeRecoverInitiateResponse](../../models/operations/volumerecoverinitiateresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsVolumesVolumeRecover

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsVolumesVolumeRecover(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `volume`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsVolumesVolumeRecoverResponse](../../models/operations/optionsvolumesvolumerecoverresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## SnapshotsList

Get a list of snapshots

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.SnapshotsList(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Snapshots != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `volume`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SnapshotsListResponse](../../models/operations/snapshotslistresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## SnapshotCreate

Create a new snapshot

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.SnapshotCreate(ctx, "<value>", components.SnapshotInput{
        Name: "<value>",
        Consistency: components.ConsistencyCrash,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `volume`                                                             | *string*                                                             | :heavy_check_mark:                                                   | N/A                                                                  |
| `snapshot`                                                           | [components.SnapshotInput](../../models/components/snapshotinput.md) | :heavy_check_mark:                                                   | A Snapshot object                                                    |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.SnapshotCreateResponse](../../models/operations/snapshotcreateresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsVolumesVolumeSnapshots

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsVolumesVolumeSnapshots(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `volume`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsVolumesVolumeSnapshotsResponse](../../models/operations/optionsvolumesvolumesnapshotsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## SnapshotGet

Get the properties of a snapshot

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.SnapshotGet(ctx, "<value>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Snapshot != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `volume`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `snapshot`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SnapshotGetResponse](../../models/operations/snapshotgetresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## SnapshotDelete

Delete a snapshot

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.SnapshotDelete(ctx, "<value>", "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `volume`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `snapshot`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `force`                                                  | **bool*                                                  | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SnapshotDeleteResponse](../../models/operations/snapshotdeleteresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsVolumesVolumeSnapshotsSnapshot

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsVolumesVolumeSnapshotsSnapshot(ctx, "<value>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `volume`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `snapshot`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsVolumesVolumeSnapshotsSnapshotResponse](../../models/operations/optionsvolumesvolumesnapshotssnapshotresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## SnapshotModify

Modify a snapshot

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.SnapshotModify(ctx, "<value>", "<value>", components.SnapshotInput{
        Name: "<value>",
        Consistency: components.ConsistencyCrash,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `volume`                                                             | *string*                                                             | :heavy_check_mark:                                                   | N/A                                                                  |
| `snapshotPathParameter`                                              | *string*                                                             | :heavy_check_mark:                                                   | N/A                                                                  |
| `snapshot1`                                                          | [components.SnapshotInput](../../models/components/snapshotinput.md) | :heavy_check_mark:                                                   | A Snapshot object                                                    |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.SnapshotModifyResponse](../../models/operations/snapshotmodifyresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsVolumesVolumeSnapshotsSnapshotRollback

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsVolumesVolumeSnapshotsSnapshotRollback(ctx, "<value>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `volume`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `snapshot`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsVolumesVolumeSnapshotsSnapshotRollbackResponse](../../models/operations/optionsvolumesvolumesnapshotssnapshotrollbackresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## SnapshotRollback

Roll back to snapshot

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.SnapshotRollback(ctx, "<value>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `volume`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `snapshot`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SnapshotRollbackResponse](../../models/operations/snapshotrollbackresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## AssociationsList

Get a list of associations

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.AssociationsList(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Associations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `startfrom`                                              | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `count`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.AssociationsListResponse](../../models/operations/associationslistresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## AssociationCreate

Create association

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.AssociationCreate(ctx, components.AssociationCreate{})
    if err != nil {
        log.Fatal(err)
    }
    if res.SuccessJobResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [components.AssociationCreate](../../models/components/associationcreate.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.AssociationCreateResponse](../../models/operations/associationcreateresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsAssociations

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsAssociations(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsAssociationsResponse](../../models/operations/optionsassociationsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## AssociationDelete

Delete an association

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.AssociationDelete(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `association`                                            | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.AssociationDeleteResponse](../../models/operations/associationdeleteresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsAssociationsAssociation

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsAssociationsAssociation(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `association`                                            | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsAssociationsAssociationResponse](../../models/operations/optionsassociationsassociationresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## AssociationModify

Association modify

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.AssociationModify(ctx, "<value>", components.AssociationModify{})
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `association`                                                                | *string*                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |
| `associationModify`                                                          | [components.AssociationModify](../../models/components/associationmodify.md) | :heavy_check_mark:                                                           | An association object                                                        |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.AssociationModifyResponse](../../models/operations/associationmodifyresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ExportsList

Get a list of associations

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.ExportsList(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Exports != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `startfrom`                                              | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `count`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ExportsListResponse](../../models/operations/exportslistresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ExportCreate

Create export

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.ExportCreate(ctx, components.ExportCreate{})
    if err != nil {
        log.Fatal(err)
    }
    if res.SuccessJobResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [components.ExportCreate](../../models/components/exportcreate.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.ExportCreateResponse](../../models/operations/exportcreateresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsExports

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsExports(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsExportsResponse](../../models/operations/optionsexportsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ExportDelete

Delete an export

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.ExportDelete(ctx, "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SuccessJobResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `export`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `force`                                                  | **bool*                                                  | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ExportDeleteResponse](../../models/operations/exportdeleteresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsExportsExport

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsExportsExport(ctx, "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `export`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `force`                                                  | **bool*                                                  | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsExportsExportResponse](../../models/operations/optionsexportsexportresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## AutoProvisionVolumes

Create a new auto provisioned volume

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.AutoProvisionVolumes(ctx, components.AutoProvisionVolume{
        Volume: components.VolumeInput{
            Name: "<value>",
            Type: components.TypeBlock,
            Size: 488662,
            Policy: "<value>",
        },
        CloudProvider: "<value>",
        AccountID: "<id>",
        Region: "<value>",
        AvailabilityZones: []string{
            "<value>",
        },
        Subnets: []string{
            "<value>",
        },
        OsType: components.AutoProvisionVolumeOsTypeRhel,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.AutoProvisionVolume](../../models/components/autoprovisionvolume.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.AutoProvisionVolumesResponse](../../models/operations/autoprovisionvolumesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsAutoprovisionvolumes

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsAutoprovisionvolumes(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsAutoprovisionvolumesResponse](../../models/operations/optionsautoprovisionvolumesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## SnapshotsListAll

Get a list of all snapshots

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.SnapshotsListAll(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Snapshots != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `capacity`                                               | **bool*                                                  | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SnapshotsListAllResponse](../../models/operations/snapshotslistallresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsSnapshots

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsSnapshots(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsSnapshotsResponse](../../models/operations/optionssnapshotsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## AttachmentsListAll

Get a list of all attachments

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.AttachmentsListAll(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Attachments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.AttachmentsListAllResponse](../../models/operations/attachmentslistallresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsAttachments

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsAttachments(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsAttachmentsResponse](../../models/operations/optionsattachmentsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## AttachmentsListForVolume

Get a list of attachments for a given volume

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.AttachmentsListForVolume(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Attachments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `volume`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.AttachmentsListForVolumeResponse](../../models/operations/attachmentslistforvolumeresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsVolumesVolumeAttachments

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsVolumesVolumeAttachments(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `volume`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsVolumesVolumeAttachmentsResponse](../../models/operations/optionsvolumesvolumeattachmentsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## AttachmentsList

Get a list of attachments for a given volume and snapshot

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.AttachmentsList(ctx, "<value>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Attachments != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `volume`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `snapshot`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.AttachmentsListResponse](../../models/operations/attachmentslistresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## AttachmentCreate

Create a new attachment

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.AttachmentCreate(ctx, "<value>", "<value>", components.AttachmentInput{
        Node: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `volume`                                                                 | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `snapshot`                                                               | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `attachment`                                                             | [components.AttachmentInput](../../models/components/attachmentinput.md) | :heavy_check_mark:                                                       | An Attachment object                                                     |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.AttachmentCreateResponse](../../models/operations/attachmentcreateresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsVolumesVolumeSnapshotsSnapshotAttachments

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsVolumesVolumeSnapshotsSnapshotAttachments(ctx, "<value>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `volume`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `snapshot`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsVolumesVolumeSnapshotsSnapshotAttachmentsResponse](../../models/operations/optionsvolumesvolumesnapshotssnapshotattachmentsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## AttachmentGet

Get the properties of an attachment

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.AttachmentGet(ctx, "<value>", "<value>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Attachment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `volume`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `snapshot`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `node`                                                   | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.AttachmentGetResponse](../../models/operations/attachmentgetresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## AttachmentDelete

Delete an attachment

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.AttachmentDelete(ctx, "<value>", "<value>", "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `volume`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `snapshot`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `node`                                                   | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `force`                                                  | **bool*                                                  | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.AttachmentDeleteResponse](../../models/operations/attachmentdeleteresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsVolumesVolumeSnapshotsSnapshotAttachmentsNode

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsVolumesVolumeSnapshotsSnapshotAttachmentsNode(ctx, "<value>", "<value>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `volume`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `snapshot`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `node`                                                   | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsVolumesVolumeSnapshotsSnapshotAttachmentsNodeResponse](../../models/operations/optionsvolumesvolumesnapshotssnapshotattachmentsnoderesponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## AttachmentModify

Modify an attachment

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.AttachmentModify(ctx, "<value>", "<value>", "<value>", components.AttachmentInput{
        Node: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `volume`                                                                 | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `snapshot`                                                               | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `node`                                                                   | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `attachment`                                                             | [components.AttachmentInput](../../models/components/attachmentinput.md) | :heavy_check_mark:                                                       | An Attachment object                                                     |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.AttachmentModifyResponse](../../models/operations/attachmentmodifyresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## NodesList

Get a list of nodes

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.NodesList(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Nodes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.NodesListResponse](../../models/operations/nodeslistresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsNodes

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsNodes(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsNodesResponse](../../models/operations/optionsnodesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsNodesTagsNode

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsNodesTagsNode(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `node`                                                   | *string*                                                 | :heavy_check_mark:                                       | Name of node to return                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsNodesTagsNodeResponse](../../models/operations/optionsnodestagsnoderesponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## NodeSetTags

Set the tags of a node

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.NodeSetTags(ctx, "<value>", map[string]string{
        "key": "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `node`                                                   | *string*                                                 | :heavy_check_mark:                                       | Name of node to return                                   |
| `requestBody`                                            | map[string]*string*                                      | :heavy_check_mark:                                       | user defined tags                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.NodeSetTagsResponse](../../models/operations/nodesettagsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## NodeGet

Get the properties of a node

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.NodeGet(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Node != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `node`                                                   | *string*                                                 | :heavy_check_mark:                                       | Name of node to return                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.NodeGetResponse](../../models/operations/nodegetresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## NodeDelete

Delete a node

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.NodeDelete(ctx, "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `node`                                                   | *string*                                                 | :heavy_check_mark:                                       | Name of node to return                                   |
| `force`                                                  | **bool*                                                  | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.NodeDeleteResponse](../../models/operations/nodedeleteresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsNodesNode

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsNodesNode(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `node`                                                   | *string*                                                 | :heavy_check_mark:                                       | Name of node to return                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsNodesNodeResponse](../../models/operations/optionsnodesnoderesponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## NodeDrain

Node drain

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.NodeDrain(ctx, "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `node`                                                   | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `cleanup`                                                | **bool*                                                  | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.NodeDrainResponse](../../models/operations/nodedrainresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsNodesNodeDrain

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsNodesNodeDrain(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `node`                                                   | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsNodesNodeDrainResponse](../../models/operations/optionsnodesnodedrainresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## NodeDescribe

Node describe

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.NodeDescribe(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.NodeDescribeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `node`                                                   | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.NodeDescribeResponse](../../models/operations/nodedescriberesponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsNodesNodeDescribe

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsNodesNodeDescribe(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `node`                                                   | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsNodesNodeDescribeResponse](../../models/operations/optionsnodesnodedescriberesponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## NodeCollectLogs

Node collect logs

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.NodeCollectLogs(ctx, "<value>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `node`                                                   | *string*                                                 | :heavy_check_mark:                                       | Name of node                                             |
| `tenant`                                                 | *string*                                                 | :heavy_check_mark:                                       | Tenant ID                                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.NodeCollectLogsResponse](../../models/operations/nodecollectlogsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsNodesLogsNodeTenant

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsNodesLogsNodeTenant(ctx, "<value>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `node`                                                   | *string*                                                 | :heavy_check_mark:                                       | Name of node                                             |
| `tenant`                                                 | *string*                                                 | :heavy_check_mark:                                       | Tenant ID                                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsNodesLogsNodeTenantResponse](../../models/operations/optionsnodeslogsnodetenantresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## NodeRepair

Execute commands on node for repair

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.NodeRepair(ctx, "<value>", "<value>", components.RepairCmds{
        Cmds: []string{
            "<value>",
            "<value>",
        },
        Checksum: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `node`                                                         | *string*                                                       | :heavy_check_mark:                                             | Name of node                                                   |
| `tenant`                                                       | *string*                                                       | :heavy_check_mark:                                             | Tenant ID                                                      |
| `repairCmds`                                                   | [components.RepairCmds](../../models/components/repaircmds.md) | :heavy_check_mark:                                             | A repair node object                                           |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*operations.NodeRepairResponse](../../models/operations/noderepairresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsNodesRepairNodeTenant

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsNodesRepairNodeTenant(ctx, "<value>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `node`                                                   | *string*                                                 | :heavy_check_mark:                                       | Name of node                                             |
| `tenant`                                                 | *string*                                                 | :heavy_check_mark:                                       | Tenant ID                                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsNodesRepairNodeTenantResponse](../../models/operations/optionsnodesrepairnodetenantresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## NodeUpgrade

performing node version upgrade

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.NodeUpgrade(ctx, "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                         | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ctx`                                                             | [context.Context](https://pkg.go.dev/context#Context)             | :heavy_check_mark:                                                | The context to use for the request.                               |
| `node`                                                            | *string*                                                          | :heavy_check_mark:                                                | Name of node to upgrade                                           |
| `nodeVersion`                                                     | [*components.NodeVersion](../../models/components/nodeversion.md) | :heavy_minus_sign:                                                | Connector Version                                                 |
| `opts`                                                            | [][operations.Option](../../models/operations/option.md)          | :heavy_minus_sign:                                                | The options for this request.                                     |

### Response

**[*operations.NodeUpgradeResponse](../../models/operations/nodeupgraderesponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsNodesUpgradeNode

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsNodesUpgradeNode(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `node`                                                   | *string*                                                 | :heavy_check_mark:                                       | Name of node to upgrade                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsNodesUpgradeNodeResponse](../../models/operations/optionsnodesupgradenoderesponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## NodeUpgradeForSupport

performing node version upgrade

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.NodeUpgradeForSupport(ctx, "<value>", "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                         | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ctx`                                                             | [context.Context](https://pkg.go.dev/context#Context)             | :heavy_check_mark:                                                | The context to use for the request.                               |
| `tenant`                                                          | *string*                                                          | :heavy_check_mark:                                                | Name of tenant to upgrade                                         |
| `node`                                                            | *string*                                                          | :heavy_check_mark:                                                | Name of node to upgrade                                           |
| `nodeVersion`                                                     | [*components.NodeVersion](../../models/components/nodeversion.md) | :heavy_minus_sign:                                                | Connector Version                                                 |
| `opts`                                                            | [][operations.Option](../../models/operations/option.md)          | :heavy_minus_sign:                                                | The options for this request.                                     |

### Response

**[*operations.NodeUpgradeForSupportResponse](../../models/operations/nodeupgradeforsupportresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsNodesUpgradeNodeTenantTenant

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsNodesUpgradeNodeTenantTenant(ctx, "<value>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `tenant`                                                 | *string*                                                 | :heavy_check_mark:                                       | Name of tenant to upgrade                                |
| `node`                                                   | *string*                                                 | :heavy_check_mark:                                       | Name of node to upgrade                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsNodesUpgradeNodeTenantTenantResponse](../../models/operations/optionsnodesupgradenodetenanttenantresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## PoliciesList

Get a list of policies

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.PoliciesList(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Policies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PoliciesListResponse](../../models/operations/policieslistresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## PolicyCreate

Create a new policy

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.PolicyCreate(ctx, components.PolicyInput{
        Name: "<value>",
        Capacityoptimization: components.CapacityoptimizationBalanced,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [components.PolicyInput](../../models/components/policyinput.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*operations.PolicyCreateResponse](../../models/operations/policycreateresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsPolicies

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsPolicies(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsPoliciesResponse](../../models/operations/optionspoliciesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## PolicyGet

Get the properties of a policy

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.PolicyGet(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Policy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `policy`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PolicyGetResponse](../../models/operations/policygetresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## PolicyDelete

Delete a policy

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.PolicyDelete(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `policy`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PolicyDeleteResponse](../../models/operations/policydeleteresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsPoliciesPolicy

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsPoliciesPolicy(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `policy`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsPoliciesPolicyResponse](../../models/operations/optionspoliciespolicyresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## PolicyModify

Modify a policy

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.PolicyModify(ctx, "<value>", components.PolicyInput{
        Name: "<value>",
        Capacityoptimization: components.CapacityoptimizationBalanced,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `policyPathParameter`                                            | *string*                                                         | :heavy_check_mark:                                               | N/A                                                              |
| `policy1`                                                        | [components.PolicyInput](../../models/components/policyinput.md) | :heavy_check_mark:                                               | A Policy object                                                  |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*operations.PolicyModifyResponse](../../models/operations/policymodifyresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## PolicyGetVolumes

Get the properties of a policy

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.PolicyGetVolumes(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Volumes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `policy`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PolicyGetVolumesResponse](../../models/operations/policygetvolumesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsPoliciesPolicyVolumes

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsPoliciesPolicyVolumes(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `policy`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsPoliciesPolicyVolumesResponse](../../models/operations/optionspoliciespolicyvolumesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## PolicyPlan

Show policy volume create plan

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.PolicyPlan(ctx, "<value>", 456938, "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Plan != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `policy`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `size`                                                   | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |
| `zone`                                                   | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `capacityGroup`                                          | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PolicyPlanResponse](../../models/operations/policyplanresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsPoliciesPolicySizeSizeZoneZone

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsPoliciesPolicySizeSizeZoneZone(ctx, "<value>", 922490, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `policy`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `size`                                                   | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |
| `zone`                                                   | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsPoliciesPolicySizeSizeZoneZoneResponse](../../models/operations/optionspoliciespolicysizesizezonezoneresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## NetworksList

Get a list of networks

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.NetworksList(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Networks != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.NetworksListResponse](../../models/operations/networkslistresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## NetworkCreate

Create a new network

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.NetworkCreate(ctx, components.Network{
        Name: "<value>",
        Type: components.NetworkTypeManagement,
        Ipstart: "<value>",
        Ipend: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [components.Network](../../models/components/network.md) | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.NetworkCreateResponse](../../models/operations/networkcreateresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsNetworks

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsNetworks(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsNetworksResponse](../../models/operations/optionsnetworksresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## NetworkGet

Get the properties of a network

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.NetworkGet(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Network != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `network`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.NetworkGetResponse](../../models/operations/networkgetresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## NetworkDelete

Delete a network

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.NetworkDelete(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `network`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.NetworkDeleteResponse](../../models/operations/networkdeleteresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsNetworksNetwork

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsNetworksNetwork(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `network`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsNetworksNetworkResponse](../../models/operations/optionsnetworksnetworkresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## NetworkModify

Modify a network

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.NetworkModify(ctx, "<value>", components.Network{
        Name: "<value>",
        Type: components.NetworkTypeStorage,
        Ipstart: "<value>",
        Ipend: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `networkPathParameter`                                   | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `network1`                                               | [components.Network](../../models/components/network.md) | :heavy_check_mark:                                       | A Network object                                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.NetworkModifyResponse](../../models/operations/networkmodifyresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## JobsList

Get a list of jobs

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.JobsList(ctx, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Jobs != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `internal`                                               | **bool*                                                  | :heavy_minus_sign:                                       | N/A                                                      |
| `page`                                                   | **int64*                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `count`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.JobsListResponse](../../models/operations/jobslistresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsJobs

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsJobs(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsJobsResponse](../../models/operations/optionsjobsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## JobGet

Get the properties of a job

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.JobGet(ctx, 615165)
    if err != nil {
        log.Fatal(err)
    }
    if res.Job != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `job`                                                    | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.JobGetResponse](../../models/operations/jobgetresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## JobDelete

Delete a job

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.JobDelete(ctx, 410406)
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `job`                                                    | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.JobDeleteResponse](../../models/operations/jobdeleteresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsJobsJob

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsJobsJob(ctx, 249106)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `job`                                                    | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsJobsJobResponse](../../models/operations/optionsjobsjobresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsJobsJobResumeSuspendState

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsJobsJobResumeSuspendState(ctx, 376221, 878301)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `job`                                                    | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |
| `state`                                                  | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsJobsJobResumeSuspendStateResponse](../../models/operations/optionsjobsjobresumesuspendstateresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## JobResumeSuspend

Resume or Suspend a job

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.JobResumeSuspend(ctx, 31568, 468286)
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `job`                                                    | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |
| `state`                                                  | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.JobResumeSuspendResponse](../../models/operations/jobresumesuspendresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## MediaList

Get a list of media

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.MediaList(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Media != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.MediaListResponse](../../models/operations/medialistresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsMedia

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsMedia(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsMediaResponse](../../models/operations/optionsmediaresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## MediaGet

Get the properties of a media

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.MediaGet(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Media != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `media`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.MediaGetResponse](../../models/operations/mediagetresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## MediaDelete

Delete a media

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.MediaDelete(ctx, "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `media`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `force`                                                  | **bool*                                                  | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.MediaDeleteResponse](../../models/operations/mediadeleteresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsMediaMedia

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsMediaMedia(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `media`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsMediaMediaResponse](../../models/operations/optionsmediamediaresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## MediaModify

modify a media properties

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.MediaModify(ctx, "<value>", components.MediaModify{})
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `media`                                                          | *string*                                                         | :heavy_check_mark:                                               | N/A                                                              |
| `mediaModify`                                                    | [components.MediaModify](../../models/components/mediamodify.md) | :heavy_check_mark:                                               | A Media Modify Object                                            |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*operations.MediaModifyResponse](../../models/operations/mediamodifyresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## MediaAssignLegacy

Assign media

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.MediaAssignLegacy(ctx, "<value>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `media`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `capacityGroup`                                          | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `reprofile`                                              | **bool*                                                  | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.MediaAssignLegacyResponse](../../models/operations/mediaassignlegacyresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsMediaMediaAssign

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsMediaMediaAssign(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `media`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsMediaMediaAssignResponse](../../models/operations/optionsmediamediaassignresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## MediaAssign

Assign media

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.MediaAssign(ctx, "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                         | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ctx`                                                             | [context.Context](https://pkg.go.dev/context#Context)             | :heavy_check_mark:                                                | The context to use for the request.                               |
| `media`                                                           | *string*                                                          | :heavy_check_mark:                                                | N/A                                                               |
| `mediaAssign`                                                     | [*components.MediaAssign](../../models/components/mediaassign.md) | :heavy_minus_sign:                                                | N/A                                                               |
| `opts`                                                            | [][operations.Option](../../models/operations/option.md)          | :heavy_minus_sign:                                                | The options for this request.                                     |

### Response

**[*operations.MediaAssignResponse](../../models/operations/mediaassignresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## MediaUnassignLegacy

Unassign media

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.MediaUnassignLegacy(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `media`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.MediaUnassignLegacyResponse](../../models/operations/mediaunassignlegacyresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsMediaMediaUnassign

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsMediaMediaUnassign(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `media`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsMediaMediaUnassignResponse](../../models/operations/optionsmediamediaunassignresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## MediaUnassign

Unassign media

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.MediaUnassign(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `media`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.MediaUnassignResponse](../../models/operations/mediaunassignresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsMediaMediaProfile

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsMediaMediaProfile(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `media`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsMediaMediaProfileResponse](../../models/operations/optionsmediamediaprofileresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## MediaProfileModify

Modify a media profile

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.MediaProfileModify(ctx, "<value>", components.MediaProfile{})
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `media`                                                            | *string*                                                           | :heavy_check_mark:                                                 | N/A                                                                |
| `mediaProfile`                                                     | [components.MediaProfile](../../models/components/mediaprofile.md) | :heavy_check_mark:                                                 | A Media Profile object                                             |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.MediaProfileModifyResponse](../../models/operations/mediaprofilemodifyresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## MediaDrain

Media drain

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.MediaDrain(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `media`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.MediaDrainResponse](../../models/operations/mediadrainresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsMediaMediaDrain

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsMediaMediaDrain(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `media`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsMediaMediaDrainResponse](../../models/operations/optionsmediamediadrainresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## MediaDiagnose

Media diagnose

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.MediaDiagnose(ctx, "<value>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `media`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `tenant`                                                 | *string*                                                 | :heavy_check_mark:                                       | Tenant ID                                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.MediaDiagnoseResponse](../../models/operations/mediadiagnoseresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsMediaMediaTenantDiagnose

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsMediaMediaTenantDiagnose(ctx, "<value>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `media`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `tenant`                                                 | *string*                                                 | :heavy_check_mark:                                       | Tenant ID                                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsMediaMediaTenantDiagnoseResponse](../../models/operations/optionsmediamediatenantdiagnoseresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## VirtualMediaList

Get a list of virtual media

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.VirtualMediaList(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.VirtualMedias != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `startfrom`                                              | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `count`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.VirtualMediaListResponse](../../models/operations/virtualmedialistresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## VirtualMediaCreate

Create (virtual) media

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.VirtualMediaCreate(ctx, components.VirtualMediaCreate{})
    if err != nil {
        log.Fatal(err)
    }
    if res.SuccessJobResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.VirtualMediaCreate](../../models/components/virtualmediacreate.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.VirtualMediaCreateResponse](../../models/operations/virtualmediacreateresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsVirtualmedia

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsVirtualmedia(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsVirtualmediaResponse](../../models/operations/optionsvirtualmediaresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## VirtualMediaDelete

Delete virtual media

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.VirtualMediaDelete(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `media`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.VirtualMediaDeleteResponse](../../models/operations/virtualmediadeleteresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsVirtualmediaMedia

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsVirtualmediaMedia(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `media`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsVirtualmediaMediaResponse](../../models/operations/optionsvirtualmediamediaresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ConnectivitiesList

Get a list of connectivities

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.ConnectivitiesList(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Connectivities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ConnectivitiesListResponse](../../models/operations/connectivitieslistresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ConnectivityCreate

Create a new connectivity

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.ConnectivityCreate(ctx, components.Connectivity{
        Name: "<value>",
        Zones1: "<value>",
        Systemtypes1: "<value>",
        Zones2: "<value>",
        Systemtypes2: "<value>",
        Mediaprotocol: "<value>",
        Replicationprotocol: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [components.Connectivity](../../models/components/connectivity.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.ConnectivityCreateResponse](../../models/operations/connectivitycreateresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsConnectivities

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsConnectivities(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsConnectivitiesResponse](../../models/operations/optionsconnectivitiesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ConnectivityGet

Get the properties of a connectivity

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.ConnectivityGet(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Connectivity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `connectivity`                                           | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ConnectivityGetResponse](../../models/operations/connectivitygetresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ConnectivityDelete

Delete a connectivity

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.ConnectivityDelete(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `connectivity`                                           | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ConnectivityDeleteResponse](../../models/operations/connectivitydeleteresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsConnectivitiesConnectivity

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsConnectivitiesConnectivity(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `connectivity`                                           | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsConnectivitiesConnectivityResponse](../../models/operations/optionsconnectivitiesconnectivityresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ConnectivityModify

Modify a connectivity

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.ConnectivityModify(ctx, "<value>", components.Connectivity{
        Name: "<value>",
        Zones1: "<value>",
        Systemtypes1: "<value>",
        Zones2: "<value>",
        Systemtypes2: "<value>",
        Mediaprotocol: "<value>",
        Replicationprotocol: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `connectivityPathParameter`                                        | *string*                                                           | :heavy_check_mark:                                                 | N/A                                                                |
| `connectivity1`                                                    | [components.Connectivity](../../models/components/connectivity.md) | :heavy_check_mark:                                                 | A Connectivity object                                              |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.ConnectivityModifyResponse](../../models/operations/connectivitymodifyresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## CapacityGroupGet

Get available capacity groups

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.CapacityGroupGet(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.CapacityGroups != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CapacityGroupGetResponse](../../models/operations/capacitygroupgetresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsCapacitygroups

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsCapacitygroups(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsCapacitygroupsResponse](../../models/operations/optionscapacitygroupsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## UserCreate

Create Tenant's first user

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.UserCreate(ctx, components.TenantCreateAdminUserRequest{
        Email: "Kenton_Roob@hotmail.com",
        Password: "pj1nHuhXMsEv8nr",
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SignUpResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [components.TenantCreateAdminUserRequest](../../models/components/tenantcreateadminuserrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UserCreateResponse](../../models/operations/usercreateresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| apierrors.SignUpResponse | 500                      | application/json         |
| apierrors.APIError       | 4XX, 5XX                 | \*/\*                    |

## OptionsSignup

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsSignup(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsSignupResponse](../../models/operations/optionssignupresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Signout

Signs out user from all devices

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.Signout(ctx, components.SignoutRequest{
        AccessToken: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [components.SignoutRequest](../../models/components/signoutrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.SignoutResponse](../../models/operations/signoutresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsSignout

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsSignout(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsSignoutResponse](../../models/operations/optionssignoutresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## SignIn

SIO Sign In

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.SignIn(ctx, components.SignIn{
        Email: "Melvin_Gleason@gmail.com",
        Password: "xx1YiyzOqOQfF59",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SignInResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [components.SignIn](../../models/components/signin.md)   | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SignInResponse](../../models/operations/signinresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsSignin

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsSignin(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsSigninResponse](../../models/operations/optionssigninresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Tenanthostdelete

Delete a tenant host

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.Tenanthostdelete(ctx, "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.TenantHostDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `tenantID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `tenantHost`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.TenanthostdeleteResponse](../../models/operations/tenanthostdeleteresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.TenantHostDeleteResponse | 500                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

## OptionsTenantTenantIDTenanthostsTenantHost

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsTenantTenantIDTenanthostsTenantHost(ctx, "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `tenantID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `tenantHost`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsTenantTenantIDTenanthostsTenantHostResponse](../../models/operations/optionstenanttenantidtenanthoststenanthostresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## TenanthostAccessCredentials

Given the tenant's host token (Refresh Token), get the ID Token for this machine

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.TenanthostAccessCredentials(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.RefreshTokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `refreshtoken`                                           | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.TenanthostAccessCredentialsResponse](../../models/operations/tenanthostaccesscredentialsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsTenantTenanthostCredentials

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsTenantTenanthostCredentials(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsTenantTenanthostCredentialsResponse](../../models/operations/optionstenanttenanthostcredentialsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Tenanthostget

Get a tenant host

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.Tenanthostget(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.GetTenantHostResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `tenanthosttoken`                                        | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.TenanthostgetResponse](../../models/operations/tenanthostgetresponse.md), error**

### Errors

| Error Type                      | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| apierrors.GetTenantHostResponse | 500                             | application/json                |
| apierrors.APIError              | 4XX, 5XX                        | \*/\*                           |

## OptionsTenantTenanthost

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsTenantTenanthost(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsTenantTenanthostResponse](../../models/operations/optionstenanttenanthostresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## TenantIDget

Get a Tenant ID from Tenant's Token

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.TenantIDget(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.GetTenantIDResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `tenanttoken`                                            | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.TenantIDgetResponse](../../models/operations/tenantidgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.GetTenantIDResponse | 500                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## OptionsTenantTenantid

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsTenantTenantid(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsTenantTenantidResponse](../../models/operations/optionstenanttenantidresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## TenantRefreshToken

Get the tenant's Refresh Token

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.TenantRefreshToken(ctx, components.GetTenantRefreshTokenRequest{
        Accesstoken: "<value>",
        Hostname: "immediate-independence.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RefreshToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [components.GetTenantRefreshTokenRequest](../../models/components/gettenantrefreshtokenrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.TenantRefreshTokenResponse](../../models/operations/tenantrefreshtokenresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsTenantRefreshtoken

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsTenantRefreshtoken(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsTenantRefreshtokenResponse](../../models/operations/optionstenantrefreshtokenresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## TenantAPIAccessRefreshToken

Given the Tenant's API Access Refresh Token, return new Access Token, Identity Token and a newer Refresh Token

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.TenantAPIAccessRefreshToken(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.RefreshTokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `refreshtoken`                                           | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.TenantAPIAccessRefreshTokenResponse](../../models/operations/tenantapiaccessrefreshtokenresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsTenantApiaccessCredentialsRefresh

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsTenantApiaccessCredentialsRefresh(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `refreshtoken`                                           | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsTenantApiaccessCredentialsRefreshResponse](../../models/operations/optionstenantapiaccesscredentialsrefreshresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## TenantToken

Get the Tenant's Token

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.TenantToken(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.TenantTokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.TenantTokenResponse](../../models/operations/tenanttokenresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsTenantToken

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsTenantToken(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsTenantTokenResponse](../../models/operations/optionstenanttokenresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Tenantuserget

Get tenant user

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.Tenantuserget(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetTenanUserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.TenantusergetResponse](../../models/operations/tenantusergetresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.GetTenanUserResponse | 500                            | application/json               |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |

## TenantUserCreate

Create Tenant's additional user

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.TenantUserCreate(ctx, components.TenantCreateUserRequest{
        Email: "Mohammad.Goodwin@gmail.com",
        Password: "wkOTuM8dPrk9EXK",
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [components.TenantCreateUserRequest](../../models/components/tenantcreateuserrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.TenantUserCreateResponse](../../models/operations/tenantusercreateresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsTenantUser

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsTenantUser(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsTenantUserResponse](../../models/operations/optionstenantuserresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## UserConfirm

Confirm user signup

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.UserConfirm(ctx, "<id>", "Tom57", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `clientID`                                               | *string*                                                 | :heavy_check_mark:                                       | Cognito Client ID                                        |
| `userName`                                               | *string*                                                 | :heavy_check_mark:                                       | Cognito User Name                                        |
| `confirmationCode`                                       | *string*                                                 | :heavy_check_mark:                                       | Cognito Signup Confirmation Code                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.UserConfirmResponse](../../models/operations/userconfirmresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsTenantUserConfirmation

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsTenantUserConfirmation(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsTenantUserConfirmationResponse](../../models/operations/optionstenantuserconfirmationresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## RequestChangePassword

Request Change clients password

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.RequestChangePassword(ctx, components.RequestChangePasswordRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [components.RequestChangePasswordRequest](../../models/components/requestchangepasswordrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.RequestChangePasswordResponse](../../models/operations/requestchangepasswordresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsTenantUserRequestchangepassword

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsTenantUserRequestchangepassword(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsTenantUserRequestchangepasswordResponse](../../models/operations/optionstenantuserrequestchangepasswordresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ChangePasswordLoggedIn

Change clients password when user is logged in

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.ChangePasswordLoggedIn(ctx, components.ChangePasswordRequestLoggedIn{})
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [components.ChangePasswordRequestLoggedIn](../../models/components/changepasswordrequestloggedin.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ChangePasswordLoggedInResponse](../../models/operations/changepasswordloggedinresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsTenantUserChangepasswordloggedin

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsTenantUserChangepasswordloggedin(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsTenantUserChangepasswordloggedinResponse](../../models/operations/optionstenantuserchangepasswordloggedinresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ChangePassword

Change clients password

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"github.com/avivlevitski-vlz/volumez-sdk/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.ChangePassword(ctx, components.ChangePasswordRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.ChangePasswordRequest](../../models/components/changepasswordrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ChangePasswordResponse](../../models/operations/changepasswordresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsTenantUserChangepassword

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsTenantUserChangepassword(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsTenantUserChangepasswordResponse](../../models/operations/optionstenantuserchangepasswordresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## VersionGet

Get version of sio

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.VersionGet(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.VersionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.VersionGetResponse](../../models/operations/versiongetresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsVersion

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsVersion(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsVersionResponse](../../models/operations/optionsversionresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetSubscription

get subscription

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.GetSubscription(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `token`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSubscriptionResponse](../../models/operations/getsubscriptionresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsAzuremarketplaceSubscription

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsAzuremarketplaceSubscription(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsAzuremarketplaceSubscriptionResponse](../../models/operations/optionsazuremarketplacesubscriptionresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## AlertsList

Get a list of alerts

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.AlertsList(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Alerts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `capacity`                                               | **bool*                                                  | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.AlertsListResponse](../../models/operations/alertslistresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsAlerts

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsAlerts(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsAlertsResponse](../../models/operations/optionsalertsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## AlertAcknowledge

Acknowledge a pending alert

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.AlertAcknowledge(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.RegularResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `alert`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.AlertAcknowledgeResponse](../../models/operations/alertacknowledgeresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsAlertsAlertAcknowledge

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsAlertsAlertAcknowledge(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `alert`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsAlertsAlertAcknowledgeResponse](../../models/operations/optionsalertsalertacknowledgeresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetMachineInfo

get system info

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.GetMachineInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.MachineInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetMachineInfoResponse](../../models/operations/getmachineinforesponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 500           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsSystemMachineinfo

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsSystemMachineinfo(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsSystemMachineinfoResponse](../../models/operations/optionssystemmachineinforesponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 409, 500      | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## OptionsCharts

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsCharts(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsChartsResponse](../../models/operations/optionschartsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## OptionsChartsProxyPlus

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsChartsProxyPlus(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsChartsProxyPlusResponse](../../models/operations/optionschartsproxyplusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetLogTenantidLogfile

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.GetLogTenantidLogfile(ctx, "<value>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `logfile`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `tenantid`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetLogTenantidLogfileResponse](../../models/operations/getlogtenantidlogfileresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## OptionsLogTenantidLogfile

### Example Usage

```go
package main

import(
	"context"
	"os"
	volumezsdk "github.com/avivlevitski-vlz/volumez-sdk"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := volumezsdk.New(
        volumezsdk.WithSecurity(os.Getenv("VOLUMEZSDK_STORAGE_IO_AUTHORIZER")),
    )

    res, err := s.OptionsLogTenantidLogfile(ctx, "<value>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `logfile`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `tenantid`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OptionsLogTenantidLogfileResponse](../../models/operations/optionslogtenantidlogfileresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |