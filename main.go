package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/rremer/terraform-provider-keyring/keyring"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: keyring.Provider})
}
