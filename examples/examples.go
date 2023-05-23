package main

import (
	"fmt"

	"cco.dev/io/pkg/api"
	"cco.dev/io/pkg/displays"
	"cco.dev/io/pkg/networks"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	// perform a search for displays
	GetDisplays()

	// perform a search for network displays
	GetNetworkdisplays()
}

func GetDisplays() {
	client, err := displays.NewClient(api.ProductionEnvironment, &clientcredentials.Config{
		ClientID:     "my-client-id",
		ClientSecret: "my-client-secret",
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

func GetNetworkdisplays() {
	client, err := networks.NewDisplayClient(api.ProductionEnvironment, &clientcredentials.Config{
		ClientID:     "my-client-id",
		ClientSecret: "my-client-secret",
		TokenURL:     "https://direct.cco.io/v2/token",
	})
	if err != nil {
		panic(err)
	}

	// get all digital displays that are 1080p
	res, err := client.Search(
		"",
		api.EmptyOptions())
	if err != nil {
		panic(err)
	}

	// print the results
	for _, d := range res.Data {
		fmt.Printf("network: %s, displays: %+v\n", d.NetworkID, d.ExternalIDs)
	}
}