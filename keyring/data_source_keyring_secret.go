package keyring

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceKeyringSecret() *schema.Resource {
	return &schema.Resource{
		Read: resourceKeyringSecretRead,
		Schema: map[string]*schema.Schema{
			"keyring": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service": {
				Type:         schema.TypeString,
				Default:      defaultService,
				Optional:     true,
				ValidateFunc: validateKeyringService,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateKeyringEntry,
			},
			"secret": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}
