# IO SDK

The IO SDK enables easy consumption of the CCO.IO Automated Direct and Programmatic endpoints in Go based applications and from the command line. 

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

## CLI

This SDK can also be installed as a command line tool, which may be useful for troubleshooting and validation:

```bash
$ io --help
Usage: io -a <api> -m <method> [-f <filter>]

  -a, --api	        The API to use
  -f, --filter	    A filter to apply to the request
  -h, --help	      Print this help
  -m, --method	    The method to call

Example: Get a display with the ID "abc123":

	io -a displays -m get -f "abc123"

Example: Get any displays with the external ID "quattro:123":

	io -a displays -m search -f "externalIDs:quattro:*"

Example: Get all digital displays that are 1080p:

	io -a displays -m search -f "mediaProducts.type:Digital" -f "digital.width:1080"
```

### CLI Installation

