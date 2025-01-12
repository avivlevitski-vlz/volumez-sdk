# AlertsListResponse


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `HTTPMeta`                                                         | [components.HTTPMetadata](../../models/components/httpmetadata.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `Alerts`                                                           | [][components.Alert](../../models/components/alert.md)             | :heavy_minus_sign:                                                 | List of alerts                                                     |
| `Headers`                                                          | map[string][]*string*                                              | :heavy_check_mark:                                                 | N/A                                                                |