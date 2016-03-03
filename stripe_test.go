package stripe_test

import (
	"testing"

	stripe "github.com/stripe/stripe-go"
	. "github.com/stripe/stripe-go/testing"
)

func TestCheckinBackendConfigurationNewRequestWithStripeAccount(t *testing.T) {
	c := &stripe.BackendConfiguration{URL: stripe.APIURL}
	p := &stripe.Params{StripeAccount: TestMerchantID}

	req, err := c.NewRequest("", "", "", "", nil, p)
	if err != nil {
		t.Fatal(err)
	}

	if req.Header.Get("Stripe-Account") != TestMerchantID {
		t.Fatalf("Expected Stripe-Account %v but got %v.",
			TestMerchantID, req.Header.Get("Stripe-Account"))
	}

	// Also test the deprecated Account field for now as well. This should be
	// identical to the exercise above.
	p = &stripe.Params{Account: TestMerchantID}

	req, err = c.NewRequest("", "", "", "", nil, p)
	if err != nil {
		t.Fatal(err)
	}

	if req.Header.Get("Stripe-Account") != TestMerchantID {
		t.Fatalf("Expected Stripe-Account %v but got %v.",
			TestMerchantID, req.Header.Get("Stripe-Account"))
	}
}
