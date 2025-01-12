# VolumeCreateResponse


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `HTTPMeta`                                                                | [components.HTTPMetadata](../../models/components/httpmetadata.md)        | :heavy_check_mark:                                                        | N/A                                                                       |
| `RegularResponse`                                                         | [*components.RegularResponse](../../models/components/regularresponse.md) | :heavy_minus_sign:                                                        | Volume has been created successfully                                      |
| `ErrorResponse`                                                           | [*components.ErrorResponse](../../models/components/errorresponse.md)     | :heavy_minus_sign:                                                        | 409 response                                                              |
| `Headers`                                                                 | map[string][]*string*                                                     | :heavy_check_mark:                                                        | N/A                                                                       |