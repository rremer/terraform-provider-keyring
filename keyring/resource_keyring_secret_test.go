package keyring_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func testAcceptanceDatasourceKeyringSecretConfig(secret string) string {
	return fmt.Sprintf(`
resource "keyring_secret" "test" {
        name   = "test"
        secret = "%s"
}

data "keyring_secret" "test" {
        name = "${keyring_secret.test.name}"
}
`, secret)
}

func TestKeying_Resource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: testAcceptanceDatasourceKeyringSecretConfig("expectedSecret"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("keyring_secret.test", "name", "test"),
					resource.TestCheckResourceAttr("keyring_secret.test", "secret", "expectedSecret"),
					resource.TestCheckResourceAttr("data.keyring_secret.test", "name", "test"),
					resource.TestCheckResourceAttr("data.keyring_secret.test", "secret", "expectedSecret"),
				),
			},
		},
	})
}
