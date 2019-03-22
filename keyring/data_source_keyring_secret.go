package keyring

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/zalando/go-keyring"
	"log"
)

func dataSourceKeyringSecret() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKeyringSecretRead,
		Schema: map[string]*schema.Schema{
			"keyring": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service": {
				Type:         schema.TypeString,
				Default:      "terraform",
				Optional:     true,
				ValidateFunc: validateKeyringService,
			},
			"entry": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateKeyringEntry,
			},
			"value": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceKeyringSecretRead(d *schema.ResourceData, meta interface{}) error {

	service := d.Get("service").(string)
	entry := d.Get("entry").(string)
	keyringId := "login"
	log.Printf("[DEBUG] Fetching value from keyring path %s/%s/%s", keyringId, service, entry)

	value, err := keyring.Get(service, entry)
	if err != nil {
		return err
	}
	d.Set("value", value)
	d.Set("keyring", keyringId)
	//d.SetId(hash)
	return nil
}
