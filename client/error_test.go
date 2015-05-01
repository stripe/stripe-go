package client

import (
	"testing"

	. "github.com/getbread/stripe-go"
	"github.com/getbread/stripe-go/currency"
	"github.com/getbread/stripe-go/utils"
)

func TestErrors(t *testing.T) {
	c := &API{}
	c.Init("bad_key", nil)

	_, err := c.Account.Get()

	if err == nil {
		t.Errorf("Expected an error")
	}

	stripeErr := err.(*Error)

	if stripeErr.Type != InvalidRequest {
		t.Errorf("Type %v does not match expected type\n", stripeErr.Type)
	}

	if stripeErr.HTTPStatusCode != 401 {
		t.Errorf("HTTPStatusCode %q does not match expected value of \"401\"", stripeErr.HTTPStatusCode)
	}
}

func TestErrorsExtra(t *testing.T) {
	apiKey := utils.GetTestKey()

	c := &API{}
	c.Init(apiKey, nil)

	cardParams := &CardParams{
		CVC:    "888",
		Month:  "05",
		Year:   "2019",
		Number: "4000000000000341",
	}
	token, err := c.Tokens.New(&TokenParams{
		Card: cardParams,
	})

	chargeParams := &ChargeParams{
		Amount:   10000,
		Currency: currency.USD,
	}
	chargeParams.SetSource(token.ID)

	_, err = c.Charges.New(chargeParams)

	if err == nil {
		t.Errorf("Expected an error")
	}

	stripeErr := err.(*Error)

	if len(stripeErr.Extra) == 0 {
		t.Error("Expecting `Extra` field to be filled.")
	}

	chargeId, ok := stripeErr.Extra["charge"]
	if !ok {
		t.Error("Expecting `charge` key in `Extra` field to be present.")
	}

	if chargeId == "" {
		t.Error("`charge` field should not be empty.")
	}

}
