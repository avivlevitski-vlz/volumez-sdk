# NodeRepairRequest


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Node`                                                         | *string*                                                       | :heavy_check_mark:                                             | Name of node                                                   |
| `Tenant`                                                       | *string*                                                       | :heavy_check_mark:                                             | Tenant ID                                                      |
| `RepairCmds`                                                   | [components.RepairCmds](../../models/components/repaircmds.md) | :heavy_check_mark:                                             | A repair node object                                           |