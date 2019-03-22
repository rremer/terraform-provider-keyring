package keyring

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform/helper/schema"
)

type regex struct {
	matcher     string
	description string
}

var (
	keyringServiceRegex regex = regex{"^[a-zA-Z][a-zA-Z0-9]{1,254}$", "be between 1 and 254 characters, start with a letter, and contain letters and numbers"}
	keyringEntryRegex   regex = regex{"^[a-zA-Z][a-zA-Z0-9]{1,254}$", "be between 1 and 254 characters, start with a letter, and contain letters and numbers"}
)

func validateRegex(re regex) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(string)
		if !regexp.MustCompile(re.matcher).MatchString(value) {
			errors = append(errors, fmt.Errorf(
				"%q (%q) must %q", k, value, re.description))
		}
		return
	}
}

func validateKeyringEntry(v interface{}, k string) (ws []string, errors []error) {
	return validateRegex(keyringEntryRegex)(v, k)
}

func validateKeyringService(v interface{}, k string) (ws []string, errors []error) {
	return validateRegex(keyringServiceRegex)(v, k)
}
