# IO SDK

The IO SDK enables easy consumption of the CCO.IO Automated Direct and Programmatic endpoints in Go based applications and from the command line. By default, the SDK is configured to communicate via the `https://direct.cco.io` gateway. The SDK enables custom server configuration to support additional sandbox development and testing scenarios.

## Usage

```bash
go get -u github.com/clearchanneloutdoor/io-sdk-golang
```

The SDK can be leveraged in applications directly, as shown in the example below that queries all digital billboards that are 1080p:

```go
package main

import (
	"fmt"

	"cco.dev/io/pkg/api"
	"cco.dev/io/pkg/displays"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	// create a new client for retrieving displays
	client, err := displays.NewClient(api.ProductionEnvironment, &clientcredentials.Config{
		ClientID:     "replace-with-your-client-id",
		ClientSecret: "replace-with-your-client-secret",
		TokenURL:     "https://direct.cco.io/v2/token",
	})
	if err != nil {
		panic(err)
	}

	// get all digital displays that are 1080p
	res, err := client.Search(
		api.EmptyOptions().
			AddFilter("mediaProducts.type", "Digital").
			AddFilter("digital.width", 1080))
	if err != nil {
		panic(err)
	}

	// print the results
	for _, d := range res.Data {
		fmt.Printf("%s\t%s\n", d.ID, d.Title)
	}
}
```
