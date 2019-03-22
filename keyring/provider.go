package keyring

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"keyring_secret": dataSourceKeyringSecret(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"keyring_secret": resourceKeyringSecret(),
		},
	}
}
