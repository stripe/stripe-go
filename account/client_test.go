package account

import (
	"testing"

	stripe "github.com/stripe-internal/stripe-go"
	"github.com/stripe-internal/stripe-go/token"
	. "github.com/stripe-internal/stripe-go/utils"
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
		TOSAcceptance: &stripe.TOSAcceptanceParams{
			IP:        "127.0.0.1",
			Date:      1437578361,
			UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/600.7.12 (KHTML, like Gecko) Version/8.0.7 Safari/600.7.12",
		},
	}

	_, err := New(params)
	if err != nil {
		t.Error(err)
	}
}

func TestAccountDelete(t *testing.T) {
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

	acct, err := New(params)
	if err != nil {
		t.Error(err)
	}
	err = Del(acct.ID)
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

func TestAccountUpdateWithBankAccount(t *testing.T) {
	params := &stripe.AccountParams{
		Managed: true,
		Country: "CA",
	}

	acct, _ := New(params)

	params = &stripe.AccountParams{
		ExternalAccount: &stripe.BankAccountParams{
			Country: "US",
			Routing: "110000000",
			Account: "000123456789",
		},
	}

	_, err := Update(acct.ID, params)
	if err != nil {
		t.Error(err)
	}
}

func TestAccountUpdateWithToken(t *testing.T) {
	params := &stripe.AccountParams{
		Managed: true,
		Country: "CA",
	}

	acct, _ := New(params)

	tokenParams := &stripe.TokenParams{
		Bank: &stripe.BankAccountParams{
			Country: "US",
			Routing: "110000000",
			Account: "000123456789",
		},
	}

	tok, _ := token.New(tokenParams)

	params = &stripe.AccountParams{
		ExternalAccount: &stripe.BankAccountParams{
			Token: tok.ID,
		},
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
