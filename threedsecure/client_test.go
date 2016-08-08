package threedsecure

import (
	"testing"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/currency"
	"github.com/stripe/stripe-go/customer"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	stripe.Key = GetTestKey()
}

func TestThreeDSecureNew(t *testing.T) {
	customerParams := &stripe.CustomerParams{}
	customerParams.SetSource(&stripe.CardParams{
		Name:   "Stripe Tester",
		Number: "378282246310005",
		Month:  "06",
		Year:   "20",
	})

	cust, _ := customer.New(customerParams)

	threeDSecureParams := &stripe.ThreeDSecureParams{
		Amount:    1000,
		Currency:  currency.USD,
		Customer:  cust.ID,
		Card:      cust.Sources.Values[0].Card.ID,
		ReturnURL: "https://test.com",
	}

	target, err := New(threeDSecureParams)

	if err != nil {
		t.Error(err)
	}

	if target.Amount != threeDSecureParams.Amount {
		t.Errorf("Amount %v does not match expected amount %v\n", target.Amount, threeDSecureParams.Amount)
	}

	if target.Currency != threeDSecureParams.Currency {
		t.Errorf("Currency %q does not match expected currency %q\n", target.Currency, threeDSecureParams.Currency)
	}

	if target.Card.ID != threeDSecureParams.Card {
		t.Errorf("Card ID %q does not match expected card ID %q\n", target.Card.ID, threeDSecureParams.Card)
	}

}
