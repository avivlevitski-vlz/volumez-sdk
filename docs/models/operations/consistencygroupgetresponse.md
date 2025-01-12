# ConsistencyGroupGetResponse


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `HTTPMeta`                                                         | [components.HTTPMetadata](../../models/components/httpmetadata.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `Snapshots`                                                        | [][components.Snapshot](../../models/components/snapshot.md)       | :heavy_minus_sign:                                                 | List of snapshots                                                  |
| `Headers`                                                          | map[string][]*string*                                              | :heavy_check_mark:                                                 | N/A                                                                |