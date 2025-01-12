<!-- Start SDK Example Usage [usage] -->
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