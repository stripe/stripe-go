package account

import (
	"testing"

	stripe "github.com/stripe/stripe-go"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	stripe.Key = GetTestKey()
}

func TestAccountNew(t *testing.T) {
	params := &stripe.AccountParams{
		Managed:              true,
		Country:              "CA",
		BusinessUrl:          "www.stripe.com",
		BusinessName:         "Stripe",
		BusinessPrimaryColor: "#ffffff",
		SupportEmail:         "foo@bar.com",
		SupportUrl:           "www.stripe.com",
		SupportPhone:         "4151234567",
		LegalEntity: &stripe.LegalEntity{
			Type:         stripe.Individual,
			BusinessName: "Stripe Go",
			DOB: stripe.DOB{
				Day:   1,
				Month: 2,
				Year:  1990,
			},
		},
	}

	_, err := New(params)
	if err != nil {
		t.Error(err)
	}
}

func TestAccountGetByID(t *testing.T) {
	params := &stripe.AccountParams{
		Managed: true,
		Country: "CA",
	}

	acct, _ := New(params)

	_, err := GetByID(acct.ID, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestAccountUpdate(t *testing.T) {
	params := &stripe.AccountParams{
		Managed: true,
		Country: "CA",
	}

	acct, _ := New(params)

	params = &stripe.AccountParams{
		Statement: "Stripe Go",
	}

	_, err := Update(acct.ID, params)
	if err != nil {
		t.Error(err)
	}
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

	if len(target.BusinessName) == 0 {
		t.Errorf("Account is missing business name\n")
	}

	if len(target.BusinessPrimaryColor) == 0 {
		t.Errorf("Account is missing business primary color\n")
	}

	if len(target.BusinessUrl) == 0 {
		t.Errorf("Account is missing business URL\n")
	}

	if len(target.SupportPhone) == 0 {
		t.Errorf("Account is missing support phone\n")
	}

	if len(target.SupportEmail) == 0 {
		t.Errorf("Account is missing support email\n")
	}

	if len(target.SupportUrl) == 0 {
		t.Errorf("Account is missing support URL\n")
	}
}
