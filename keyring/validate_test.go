package keyring_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	. "github.com/rremer/terraform-provider-keyring/keyring"
)

// StringValidationTestCase lifted from github.com/terraform-providers/terraform-provider-google/google under MPL2.0
type StringValidationTestCase struct {
	TestName    string
	Value       string
	ExpectError bool
}

// testStringValidation lifted from github.com/terraform-providers/terraform-provider-google/google under MPL2.0
func testStringValidation(testCase StringValidationTestCase, validationFunc schema.SchemaValidateFunc) []error {
	_, es := validationFunc(testCase.Value, testCase.TestName)
	if testCase.ExpectError {
		if len(es) > 0 {
			return nil
		} else {
			return []error{fmt.Errorf("Didn't see expected error in case \"%s\" with string \"%s\"", testCase.TestName, testCase.Value)}
		}
	}

	return es
}

// testStringValidationCases lifted from github.com/terraform-providers/terraform-provider-google/google under MPL2.0
func testStringValidationCases(cases []StringValidationTestCase, validationFunc schema.SchemaValidateFunc) []error {
	es := make([]error, 0)
	for _, c := range cases {
		es = append(es, testStringValidation(c, validationFunc)...)
	}

	return es
}

func TestValidateKeyringServiceRegex(t *testing.T) {
	s := []StringValidationTestCase{
		{TestName: "simple", Value: "terraform"},
		{TestName: "camelcase", Value: "terraformCapitalizedExample"},
		{TestName: "allowed numbers", Value: "terraform5432example1"},
		{TestName: "short", Value: "x"},
		{TestName: "long", Value: strings.Repeat("a", 254)},

		{TestName: "underscore", Value: "terraform_example", ExpectError: true},
		{TestName: "hyphen", Value: "terraform-example", ExpectError: true},
		{TestName: "null", Value: "", ExpectError: true},
		{TestName: "starting number", Value: "9neinneinnein", ExpectError: true},

		// this is OS-specific, and will pass on Linux generally
		//{TestName: "too long", Value: strings.Repeat("a", 255), ExpectError: true},
	}

	e := testStringValidationCases(s, ValidateKeyringService)
	if len(e) > 0 {
		t.Errorf("Failed validation: %v", e)
	}
}

func TestValidateKeyringEntryRegex(t *testing.T) {
	s := []StringValidationTestCase{
		{TestName: "simple", Value: "terraform"},
		{TestName: "camelcase", Value: "terraformCapitalizedExample"},
		{TestName: "allowed numbers", Value: "terraform5432example1"},
		{TestName: "short", Value: "x"},
		{TestName: "long", Value: strings.Repeat("a", 254)},

		{TestName: "underscore", Value: "terraform_example", ExpectError: true},
		{TestName: "hyphen", Value: "terraform-example", ExpectError: true},
		{TestName: "null", Value: "", ExpectError: true},
		{TestName: "starting number", Value: "9neinneinnein", ExpectError: true},

		// this is OS-specific, and will pass on Linux generally
		//{TestName: "too long", Value: strings.Repeat("a", 255), ExpectError: true},
	}

	e := testStringValidationCases(s, ValidateKeyringEntry)
	if len(e) > 0 {
		t.Errorf("Failed validation: %v", e)
	}
}
