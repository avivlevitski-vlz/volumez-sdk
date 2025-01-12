# NetworkModifyResponse


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `HTTPMeta`                                                                | [components.HTTPMetadata](../../models/components/httpmetadata.md)        | :heavy_check_mark:                                                        | N/A                                                                       |
| `RegularResponse`                                                         | [*components.RegularResponse](../../models/components/regularresponse.md) | :heavy_minus_sign:                                                        | A network was updated successfully                                        |
| `Headers`                                                                 | map[string][]*string*                                                     | :heavy_check_mark:                                                        | N/A                                                                       |