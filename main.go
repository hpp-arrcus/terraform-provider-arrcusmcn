package main

import (
	"context"
	"log"

	"github.com/Arrcus/terraform-provider-arrcusmcn/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	// plugin.Serve(&plugin.ServeOpts{
	// 	ProviderFunc: func() *schema.Provider {
	// 		return provider.Provider()
	// 	},
	// })
	opts := providerserver.ServeOpts{
		// TODO: Update this string with the published name of your provider.
		Address: "registry.terraform.io/hashicorp/arrcusmcn",
	}

	err := providerserver.Serve(context.Background(), provider.New("1.0.0"), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
