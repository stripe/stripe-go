package testing

import (
	"net/http"
	"os"

	stripe "github.com/stripe/stripe-go"
)

// This file should contain any testing helpers that should be commonly
// available across all tests in the Stripe package.
//
// There's not much in here because it' a relatively recent addition to the
// package, but should be used as appropriate for any new changes.

const (
	// TestMerchantID is a token that can be used to represent a merchant ID in
	// simple tests.
	TestMerchantID = "acct_xxxx"
)

func init() {
	port := os.Getenv("STRIPE_STUB_PORT")
	if port == "" {
		port = "12111"
	}

	stripe.Key = "sk_test_myTestKey"
	stripe.SetBackend("api", stripe.BackendConfiguration{
		stripe.APIBackend,
		"http://localhost:" + port + "/v1",
		&http.Client{},
	})

	// TODO: sample connection and version check
}
