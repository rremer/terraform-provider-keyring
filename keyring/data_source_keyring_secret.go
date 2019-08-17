package keyring

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func keyringSecretCommonSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"keyring": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"service": {
			Type:         schema.TypeString,
			Default:      defaultService,
			Optional:     true,
			ValidateFunc: ValidateKeyringService,
		},
		"name": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: ValidateKeyringEntry,
		},
	}
}

func dataSourceKeyringSecret() *schema.Resource {
	s := keyringSecretCommonSchema()

	s["secret"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  true,
		Sensitive: true,
	}

	return &schema.Resource{
		Read:   resourceKeyringSecretRead,
		Schema: s,
	}
}
