package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"cco.dev/io/internal"
	"cco.dev/io/pkg/api"
	"cco.dev/io/pkg/clients"
	"cco.dev/io/pkg/displays"
	"cco.dev/io/pkg/markets"
	"cco.dev/io/pkg/networks"
	"golang.org/x/oauth2/clientcredentials"
)

type command struct {
	api    string
	id     string
	opts   *api.Options
	method string
}

func determineEnvironment() api.Environment {
	env := api.CustomEnvironment
	switch os.Getenv("CCO_ENV") {
	case "production":
		env = api.ProductionEnvironment
	case "develop":
		env = api.DevelopEnvironment
	case "staging":
		env = api.StagingEnvironment
	case "":
		env = api.ProductionEnvironment
	}

	return env
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
			hasID := cmd.method == "delete" ||
				cmd.method == "get" ||
				cmd.method == "patch" ||
				cmd.method == "update"

			if hasID && i < l-1 {
				i++
				cmd.id = os.Args[i]
			}
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

func runCommand[T any](client func() (*clients.Client[T], error), cmd command) {
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

func main() {
	// determine the environment
	env := determineEnvironment()
	envName := env.String()

	// override the envName if it is a known environment
	if env == api.CustomEnvironment {
		envName = os.Getenv("CCO_ENV")
	}

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

	switch cmd.api {
	case "displays":
		runCommand(func() (*clients.Client[displays.Display], error) {
			return displays.NewClient(env, cc)
		}, cmd)
	case "markets":
		runCommand(func() (*clients.Client[markets.Market], error) {
			return markets.NewClient(env, cc)
		}, cmd)
	case "networks":
		runCommand(func() (*clients.Client[networks.Network], error) {
			return networks.NewClient(env, cc)
		}, cmd)
	}
}
