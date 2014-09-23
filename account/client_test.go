package account

import (
	"testing"

	stripe "github.com/stripe/stripe-go"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	stripe.Key = GetTestKey()
}

func TestAccountGet(t *testing.T) {
	target, err := Get()

	if err != nil {
		t.Error(err)
	}

	if len(target.ID) == 0 {
		t.Errorf("Account is missing id\n")
	}

	if len(target.Country) == 0 {
		t.Errorf("Account is missing country\n")
	}

	if len(target.Currencies) == 0 {
		t.Errorf("Account is missing currencies\n")
	}

	if len(target.DefaultCurrency) == 0 {
		t.Errorf("Account is missing default currency\n")
	}

	if len(target.Name) == 0 {
		t.Errorf("Account is missing name\n")
	}

	if len(target.Email) == 0 {
		t.Errorf("Account is missing email\n")
	}

	if len(target.Timezone) == 0 {
		t.Errorf("Account is missing timezone\n")
	}
}
