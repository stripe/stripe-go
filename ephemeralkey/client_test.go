package ephemeralkey

import (
	"bytes"
	"encoding/json"
	"testing"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/ephemeralkey"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	stripe.Key = GetTestKey()
}

func TestEphemeralKeyCreate(t *testing.T) {
	cust, err := customer.New(&stripe.CustomerParams{})
	if err != nil {
		t.Fatalf("%+v", err)
	}

	k, err := ephemeralkey.New(&stripe.EphemeralKeyParams{
		Customer:      cust.ID,
		StripeVersion: "2017-05-25",
	})
	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(k.AssociatedObjects) != 1 {
		t.Fatalf("Incorrect number of associated objects for ephkey", err)
	}

	if k.AssociatedObjects[0].ID != cust.ID {
		t.Fatalf("Incorrect associated object ID for ephkey", err)
	}

	if k.AssociatedObjects[0].Type != "customer" {
		t.Fatalf("Incorrect associated object Type for ephkey", err)
	}

	k, err = ephemeralkey.New(&stripe.EphemeralKeyParams{
		Customer: cust.ID,
	})
	if err.Error() != "params.StripeVersion must be specified" {
		t.Fatalf("Incorrect error %+v", err)
	}
}

func TestEphemeralKeyJSON(t *testing.T) {
	// Test that we can extract the json even if parsing fails because the
	// frontend and backend may be using different API versions
	invalidJson := []byte("{\"foo\":5}")
	k := &stripe.EphemeralKey{}

	err := json.Unmarshal(invalidJson, k)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	if k.ID != "" {
		t.Fatalf("Invalid JSON should not have been parsed: %+v", k)
	}

	if !bytes.Equal(k.RawJSON, invalidJson) {
		t.Fatalf("EphemeralKey did not preserve the raw JSON: %+v", k.RawJSON)
	}
}

func TestEphemeralKeyDelete(t *testing.T) {
	cust, err := customer.New(&stripe.CustomerParams{})
	if err != nil {
		t.Fatalf("%+v", err)
	}

	k, err := ephemeralkey.New(&stripe.EphemeralKeyParams{
		Customer:      cust.ID,
		StripeVersion: "2017-05-25",
	})
	if err != nil {
		t.Fatalf("%+v", err)
	}

	k, err = ephemeralkey.Del(k.ID)
	if err != nil {
		t.Fatalf("%+v", err)
	}
}
