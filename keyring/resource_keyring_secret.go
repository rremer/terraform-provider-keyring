package keyring

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/zalando/go-keyring"
	"log"
)

const (
	defaultKeyringId = "Login"
	defaultService   = "terraform"
	defaultUsername  = "terraform"
)

func resourceKeyringSecret() *schema.Resource {
	return &schema.Resource{
		Create: resourceKeyringSecretCreate,
		Read:   resourceKeyringSecretRead,
		Update: resourceKeyringSecretUpdate,
		Delete: resourceKeyringSecretDelete,
		Importer: &schema.ResourceImporter{
			State: resourceKeyringSecretImport,
		},
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
			"username": {
				Type:         schema.TypeString,
				Default:      defaultUsername,
				Optional:     true,
				ValidateFunc: validateKeyringEntry,
			},
			"secret": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
		},
	}
}
func resourceKeyringSecretCreate(d *schema.ResourceData, m interface{}) error {
	service := d.Get("service").(string)
	username := d.Get("username").(string)
	keyringId := defaultKeyringId
	secret := d.Get("secret").(string)

	log.Printf("[DEBUG] Setting secret for keyring path %s/%s/%s", keyringId, service, username)
	err := keyring.Set(service, username, secret)
	if err != nil {
		log.Printf("[ERROR] Failed to set secret for keyring path %s/%s/%s", keyringId, service, username)
		return err
	}
	return resourceKeyringSecretRead(d, m)
}

//TODO(rremer) set
func resourceKeyringSecretRead(d *schema.ResourceData, m interface{}) error {

	service := d.Get("service").(string)
	username := d.Get("username").(string)
	keyringId := defaultKeyringId
	log.Printf("[DEBUG] Fetching secret from keyring path %s/%s/%s", keyringId, service, username)

	secret, err := keyring.Get(service, username)
	if err != nil {
		log.Printf("[ERROR] Failed to fetch secret from keyring path %s/%s/%s", keyringId, service, username)
		return err
	}
	d.Set("secret", secret)
	d.Set("keyring", keyringId)
	d.SetId(keyringId + service + username)
	return nil
}

func resourceKeyringSecretUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceKeyringSecretDelete(d *schema.ResourceData, m interface{}) error {
	service := d.Get("service").(string)
	username := d.Get("username").(string)
	keyringId := defaultKeyringId

	log.Printf("[DEBUG] Deleting secret at keyring path %s/%s/%s", keyringId, service, username)
	err := keyring.Delete(service, username)
	if err != nil {
		log.Printf("[ERROR] Failed to delete secret at keyring path %s/%s/%s", keyringId, service, username)
		return err
	}
	return nil
}

func resourceKeyringSecretImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	return []*schema.ResourceData{}, nil
}
