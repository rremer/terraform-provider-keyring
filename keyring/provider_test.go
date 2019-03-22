package keyring_test

import (
	"github.com/hashicorp/terraform/helper/schema"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/rremer/terraform-provider-keyring/keyring"
	"testing"
)

func TestProvider(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Provider Suite")
}

var _ = Describe("Provider", func() {

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
