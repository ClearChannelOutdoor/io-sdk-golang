package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/clearchanneloutdoor/io-sdk-golang/internal"
	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/accounts"
	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/api"
	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/bookings"
	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/clients"
	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/contracts"
	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/creatives"
	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/customers"
	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/displays"
	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/geopath"
	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/markets"
	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/networks"
	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/orders"
	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/photos"
	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/products"
	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/renewals"
	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/structures"
	"github.com/clearchanneloutdoor/io-sdk-golang/pkg/taxa"
	"golang.org/x/oauth2/clientcredentials"
)

type command struct {
	api     string
	childID string
	id      string
	opts    *api.Options
	method  string
	server  string
}

func parseArgs() command {
	cmd := command{
		opts: api.EmptyOptions(),
	}

	i := 1
	l := len(os.Args)
	for i < l {
		a := os.Args[i]

		// check for API
		if (a == "-a" || a == "--api") && i < l-1 {
			i++
			cmd.api = os.Args[i]
		}

		// check for filters
		if (a == "-f" || a == "--filter") && i < l-1 {
			i++
			// split the filter (field1:value1,field2:value2)
			terms := strings.Split(os.Args[i], ",")
			for _, t := range terms {
				// split the term (field:value)
				parts := strings.Split(t, ":")
				if len(parts) != 2 {
					continue
				}

				// add the filter
				cmd.opts.AddFilter(parts[0], parts[1])
			}
		}

		// check for help
		if a == "-h" || a == "--help" || a == "-?" {
			printUsageAndExit()
			break
		}

		// check for method
		if (a == "-m" || a == "--method") && i < l-1 {
			i++
			cmd.method = os.Args[i]

			// check for ID
			requiresID := cmd.method == "delete" ||
				cmd.method == "get" ||
				cmd.method == "patch" ||
				cmd.method == "update"

			if requiresID && i < l-1 {
				i++
				cmd.id = os.Args[i]

				// check for a child ID
				if i < l-1 && os.Args[i+1][0] != '-' {
					i++
					cmd.childID = os.Args[i]
				}
			}

			if cmd.method == "search" && i < l-1 && os.Args[i+1][0] != '-' {
				i++
				cmd.id = os.Args[i]
			}
		}

		// check for server override
		if (a == "-s" || a == "--server") && i < l-1 {
			i++
			cmd.server = os.Args[i]
		}

		// check for version
		if a == "-v" || a == "--version" {
			printVersionAndExit()
			break
		}

		i++
	}

	return cmd
}

func printUsageAndExit(exitCode ...int) {
	ec := 0
	if len(exitCode) > 0 {
		ec = exitCode[0]
	}

	fmt.Println("Usage: io -a <api> -m <method> [-f <filter>]")
	fmt.Println("Version:", api.Version)
	fmt.Println()
	fmt.Println("\t-a, --api\t\tThe API to use")
	fmt.Println("\t-f, --filter\t\tA filter to apply to the request")
	fmt.Println("\t-h, --help\t\tPrint this help")
	fmt.Println("\t-m, --method\t\tThe method to call")
	fmt.Println()
	fmt.Println("Example: Get a display with the ID \"abc123\":")
	fmt.Println()
	fmt.Println("\tio -a displays -m get -f \"abc123\"")
	fmt.Println()
	fmt.Println("Example: Get any displays with the external ID \"quattro:123\":")
	fmt.Println()
	fmt.Println("\tio -a displays -m search -f \"externalIDs:quattro:*\"")
	fmt.Println()
	fmt.Println("Example: Get all digital displays that are 1080p:")
	fmt.Println()
	fmt.Println("\tio -a displays -m search -f \"mediaProducts.type:Digital\" -f \"digital.width:1080\"")
	fmt.Println()
	os.Exit(ec)
}

func printVersionAndExit() {
	fmt.Println("Version:", api.Version)
	os.Exit(0)
}

func runClientCommand[T any](client func() (*clients.Client[T], error), cmd command) {
	cl, err := client()
	if err != nil {
		panic(err)
	}

	var res *T
	var sr api.SearchResult[T]

	switch cmd.method {
	case "delete":
		err = cl.Delete(cmd.id)
	case "get":
		res, err = cl.Get(cmd.id)
	case "patch":
		panic(errors.New("not implemented"))
	case "search":
		sr, err = cl.Search(cmd.opts)
	case "update":
		panic(errors.New("not implemented"))
	}
	if err != nil {
		panic(err)
	}

	// write out the response if there is one
	if res != nil {
		jsn, _ := json.MarshalIndent(res, "", "\t")
		os.Stdout.Write(jsn)
	}

	// write out the search result if there is one
	if sr.Total != 0 {
		jsn, _ := json.MarshalIndent(sr, "", "\t")
		os.Stdout.Write(jsn)
	}

	fmt.Println()
}

func runChildClientCommand[T any](client func() (*clients.ChildClient[T], error), cmd command) {
	cl, err := client()
	if err != nil {
		panic(err)
	}

	var res *T
	var sr api.SearchResult[T]

	switch cmd.method {
	case "delete":
		err = cl.Delete(cmd.id, cmd.childID)
	case "get":
		res, err = cl.Get(cmd.id, cmd.childID)
	case "patch":
		panic(errors.New("not implemented"))
	case "search":
		sr, err = cl.Search(cmd.id, cmd.opts)
	case "update":
		panic(errors.New("not implemented"))
	}
	if err != nil {
		panic(err)
	}

	// write out the response if there is one
	if res != nil {
		jsn, _ := json.MarshalIndent(res, "", "\t")
		os.Stdout.Write(jsn)
	}

	// write out the search result if there is one
	if sr.Total != 0 {
		jsn, _ := json.MarshalIndent(sr, "", "\t")
		os.Stdout.Write(jsn)
	}

	fmt.Println()
}

func main() {
	// override the envName if it is a known environment
	envName := os.Getenv("CCO_ENV")

	// load the access settings for the environment
	as, err := internal.LoadAccessSettings(envName)
	if err != nil {
		panic(err)
	}
	cc := &clientcredentials.Config{
		ClientID:     as.ClientID,
		ClientSecret: as.ClientSecret,
		Scopes:       as.Scopes,
		TokenURL:     as.TokenURL,
	}

	//opts := api.EmptyOptions()
	cmd := parseArgs()

	if cmd.api == "" {
		fmt.Printf("no API specified\n\n")
		printUsageAndExit(1)
	}

	if cmd.method == "" {
		fmt.Printf("no method specified\n\n")
		printUsageAndExit(1)
	}

	ctx := context.Background()

	switch cmd.api {
	case "accounts":
		runClientCommand(func() (*clients.Client[accounts.Account], error) {
			return accounts.NewClient(ctx, cc, cmd.server)
		}, cmd)
	case "bookings":
		runClientCommand(func() (*clients.Client[bookings.Booking], error) {
			return bookings.NewClient(ctx, cc, cmd.server)
		}, cmd)
	case "contracts":
		runClientCommand(func() (*clients.Client[contracts.Contract], error) {
			return contracts.NewClient(ctx, cc, cmd.server)
		}, cmd)
	case "creatives-v1":
		runClientCommand(func() (*clients.Client[creatives.Creative], error) {
			return creatives.NewCreativeV1Client(ctx, cc, cmd.server)
		}, cmd)
	case "creatives-v2":
		runClientCommand(func() (*clients.Client[creatives.AdCreative], error) {
			return creatives.NewCreativeV2Client(ctx, cc, cmd.server)
		}, cmd)
	case "customers":
		runClientCommand(func() (*clients.Client[customers.Customer], error) {
			return customers.NewClient(ctx, cc, cmd.server)
		}, cmd)
	case "displays":
		runClientCommand(func() (*clients.Client[displays.Display], error) {
			return displays.NewClient(ctx, cc, cmd.server)
		}, cmd)
	case "geopath-construction-classifications":
		runClientCommand(func() (*clients.Client[geopath.ConstructionClassification], error) {
			return geopath.NewConstructionClassificationClient(ctx, cc, cmd.server)
		}, cmd)
	case "geopath-construction-placements":
		runClientCommand(func() (*clients.Client[geopath.ConstructionPlacement], error) {
			return geopath.NewConstructionPlacementClient(ctx, cc, cmd.server)
		}, cmd)
	case "geopath-construction-types":
		runClientCommand(func() (*clients.Client[geopath.ConstructionType], error) {
			return geopath.NewConstructionTypeClient(ctx, cc, cmd.server)
		}, cmd)
	case "geopath-frames":
		runClientCommand(func() (*clients.Client[geopath.Frame], error) {
			return geopath.NewFrameClient(ctx, cc, cmd.server)
		}, cmd)
	case "geopath-frames-history":
		runChildClientCommand(func() (*clients.ChildClient[geopath.Measure], error) {
			return geopath.NewFrameHistoryClient(ctx, cc, cmd.server)
		}, cmd)
	case "geopath-frames-measures":
		runChildClientCommand(func() (*clients.ChildClient[geopath.Measure], error) {
			return geopath.NewFrameMeasuresClient(ctx, cc, cmd.server)
		}, cmd)
	case "geopath-illumination-types":
		runClientCommand(func() (*clients.Client[geopath.IlluminationType], error) {
			return geopath.NewIlluminationTypeClient(ctx, cc, cmd.server)
		}, cmd)
	case "geopath-location-types":
		runClientCommand(func() (*clients.Client[geopath.LocationType], error) {
			return geopath.NewLocationTypeClient(ctx, cc, cmd.server)
		}, cmd)
	case "geopath-measures":
		runClientCommand(func() (*clients.Client[geopath.Measure], error) {
			return geopath.NewMeasuresClient(ctx, cc, cmd.server)
		}, cmd)
	case "geopath-media-types":
		runClientCommand(func() (*clients.Client[geopath.MediaType], error) {
			return geopath.NewMediaTypeClient(ctx, cc, cmd.server)
		}, cmd)
	case "geopath-segment-ids":
		runClientCommand(func() (*clients.Client[geopath.SegmentID], error) {
			return geopath.NewSegmentIDClient(ctx, cc, cmd.server)
		}, cmd)
	case "geopath-segment-names":
		runClientCommand(func() (*clients.Client[geopath.SegmentName], error) {
			return geopath.NewSegmentNameClient(ctx, cc, cmd.server)
		}, cmd)
	case "markets":
		runClientCommand(func() (*clients.Client[markets.Market], error) {
			return markets.NewClient(ctx, cc, cmd.server)
		}, cmd)
	case "networks":
		runClientCommand(func() (*clients.Client[networks.Network], error) {
			return networks.NewClient(ctx, cc, cmd.server)
		}, cmd)
	case "network-displays":
		runChildClientCommand(func() (*clients.ChildClient[networks.NetworkDisplay], error) {
			return networks.NewDisplayClient(ctx, cc, cmd.server)
		}, cmd)
	case "orders":
		runClientCommand(func() (*clients.Client[orders.Order], error) {
			return orders.NewOrdersClient(ctx, cc, cmd.server)
		}, cmd)
	case "orderlines":
		runClientCommand(func() (*clients.Client[orders.OrderLine], error) {
			return orders.NewOrderLinesClient(ctx, cc, cmd.server)
		}, cmd)
	case "photos":
		runClientCommand(func() (*clients.Client[photos.Photo], error) {
			return photos.NewClient(ctx, cc, cmd.server)
		}, cmd)
	case "products":
		runClientCommand(func() (*clients.Client[products.Product], error) {
			return products.NewClient(ctx, cc, cmd.server)
		}, cmd)
	case "renewals":
		runClientCommand(func() (*clients.Client[renewals.Relationship], error) {
			return renewals.NewClient(ctx, cc, cmd.server)
		}, cmd)
	case "structures":
		runClientCommand(func() (*clients.Client[structures.Structure], error) {
			return structures.NewClient(ctx, cc, cmd.server)
		}, cmd)
	case "taxa-cco":
		runClientCommand(func() (*clients.Client[taxa.CCOCode], error) {
			return taxa.NewCCOCodeClient(ctx, cc, cmd.server)
		}, cmd)
	case "taxa-iab-v1":
		runClientCommand(func() (*clients.Client[taxa.IABV1Taxonomy], error) {
			return taxa.NewIABV1Client(ctx, cc, cmd.server)
		}, cmd)
	case "taxa-iab-v2":
		runClientCommand(func() (*clients.Client[taxa.IABV2Taxonomy], error) {
			return taxa.NewIABV2Client(ctx, cc, cmd.server)
		}, cmd)
	}
}
