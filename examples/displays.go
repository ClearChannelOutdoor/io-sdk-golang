package main

import (
	"fmt"

	"cco.dev/io/pkg/api"
	"cco.dev/io/pkg/displays"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
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
