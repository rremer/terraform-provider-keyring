package keyring_test

import (
	"github.com/hashicorp/terraform/helper/schema"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/rremer/terraform-provider-keyring/keyring"
)

var _ = Describe("data.keyring_secret", func() {

	Describe("Initializing the provider", func() {
		Context("with resource schema validation", func() {
			It("should succeed", func() {
				Expect(Provider().(*schema.Provider).InternalValidate()).Should(Succeed())
			})
		})

		Context("with default configuration", func() {
			It("should succeed", func() {
					err := Provider()
				Expect(err).ShouldNot(BeNil())
			})
		})
	})
})

func testAcceptanceDatasourceKeyringSecretConfig(resourceId, datasourceId string) string {
        return fmt.Sprintf(`
resource "keyring_secret" "%s" {
        entry = "terraform-acceptance-test"
        value = "Keyring provider acceptance test secret"
}

data "keyring_secret" "%s" {
        entry = "${keyring_secret.%s.resourceId}"
}

func TestAccDataSourceComputeAddress(t *testing.T) {
        t.Parallel()

        rsName := "foobar"
        rsFullName := fmt.Sprintf("google_compute_address.%s", rsName)
        dsName := "my_address"
        dsFullName := fmt.Sprintf("data.google_compute_address.%s", dsName)

        resource.Test(t, resource.TestCase{
                PreCheck:     func() { testAccPreCheck(t) },
                Providers:    testAccProviders,
                CheckDestroy: testAccCheckDataSourceComputeAddressDestroy(rsFullName),
                Steps: []resource.TestStep{
                        {
                                Config: testAccDataSourceComputeAddressConfig(rsName, dsName),
                                Check: resource.ComposeTestCheckFunc(
                                        testAccDataSourceComputeAddressCheck(dsFullName, rsFullName),
                                ),
                        },
                },
        })
}
