package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/sofixa/terraform-provider-solidfire/solidfire"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: solidfire.Provider,
	})
}
