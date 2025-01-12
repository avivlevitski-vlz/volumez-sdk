# NodeUpgradeForSupportRequest


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Tenant`                                                          | *string*                                                          | :heavy_check_mark:                                                | Name of tenant to upgrade                                         |
| `Node`                                                            | *string*                                                          | :heavy_check_mark:                                                | Name of node to upgrade                                           |
| `NodeVersion`                                                     | [*components.NodeVersion](../../models/components/nodeversion.md) | :heavy_minus_sign:                                                | Connector Version                                                 |