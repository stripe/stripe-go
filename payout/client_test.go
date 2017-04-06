package payout

import (
	"testing"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/account"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/currency"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	stripe.Key = GetTestKey()
}

func TestPayoutAllMethods(t *testing.T) {
	params := &stripe.AccountParams{
		Managed: true,
		Country: "US",
		ExternalAccount: &stripe.AccountExternalAccountParams{
			Country:  "US",
			Currency: "usd",
			Routing:  "110000000",
			Account:  "000123456789",
		},
		LegalEntity: &stripe.LegalEntity{
			Type: stripe.Individual,
			DOB: stripe.DOB{
				Day:   1,
				Month: 2,
				Year:  1990,
			},
			PersonalID: "123456789",
		},
		TOSAcceptance: &stripe.TOSAcceptanceParams{
			IP:        "127.0.0.1",
			Date:      1437578361,
			UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/600.7.12 (KHTML, like Gecko) Version/8.0.7 Safari/600.7.12",
		},
	}

	acc, err := account.New(params)
	if err != nil {
		t.Error(err)
	}

	chargeParams := &stripe.ChargeParams{
		Amount:   1000,
		Currency: currency.USD,
		Source: &stripe.SourceParams{
			Card: &stripe.CardParams{
				Number: "4000000000000077",
				Month:  "06",
				Year:   "20",
			},
		},
		Destination: &stripe.DestinationParams{
			Account: acc.ID,
		},
	}

	_, err = charge.New(chargeParams)
	if err != nil {
		t.Error(err)
	}

	payoutParams := &stripe.PayoutParams{
		Amount:     100,
		Currency:   currency.USD,
		SourceType: SourceCard,
	}
	payoutParams.Params.Account = acc.ID

	target, err := New(payoutParams)

	if err != nil {
		t.Error(err)
	}

	if target.Amount != payoutParams.Amount {
		t.Errorf("Amount %v does not match expected amount %v\n", target.Amount, payoutParams.Amount)
	}

	if target.Currency != payoutParams.Currency {
		t.Errorf("Curency %q does not match expected currency %q\n", target.Currency, payoutParams.Currency)
	}

	if target.Created == 0 {
		t.Errorf("Created date is not set\n")
	}

	if target.ArrivalDate == 0 {
		t.Errorf("Date is not set\n")
	}

	if target.SourceType != payoutParams.SourceType {
		t.Errorf("SourceType %q does not match expected SourceType %q\n", target.SourceType, payoutParams.SourceType)
	}

	getParams := &stripe.PayoutParams{}
	getParams.Params.Account = acc.ID

	payoutRetrieved, err := Get(target.ID, getParams)
	if err != nil {
		t.Error(err)
	}

	if payoutRetrieved.ID != target.ID {
		t.Errorf("ID %q does not match expected ID %q\n", payoutRetrieved.ID, target.ID)
	}

	updateParams := &stripe.PayoutParams{}
	updateParams.Params.Account = acc.ID
	updateParams.AddMeta("foo", "bar")

	payoutUpdated, err := Update(target.ID, updateParams)
	if err != nil {
		t.Error(err)
	}

	if payoutUpdated.Meta["foo"] != "bar" {
		t.Error("Payout metadata not updated")
	}

	multiplePayoutParams := &stripe.PayoutParams{
		Amount:   100,
		Currency: currency.USD,
	}
	multiplePayoutParams.Params.StripeAccount = acc.ID

	for i := 0; i < 3; i++ {
		New(multiplePayoutParams)
	}

	nbPayouts := 0
	listParams := &stripe.PayoutListParams{}
	listParams.StripeAccount = acc.ID
	i := List(listParams)
	for i.Next() {
		if i.Payout() == nil {
			t.Error("No nil values expected")
		}

		if i.Meta() == nil {
			t.Error("No metadata returned")
		}
		nbPayouts++
	}
	if err := i.Err(); err != nil {
		t.Error(err)
	}
	if nbPayouts != 4 {
		t.Errorf("Expected 4 payouts on %q but got %q\n", acc.ID, nbPayouts)
	}

	account.Del(acc.ID)
}
