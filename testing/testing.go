package testing

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// This file should contain any testing helpers that should be commonly
// available across all tests in the Stripe package.
//
// There's not much in here because it' a relatively recent addition to the
// package, but should be used as appropriate for any new changes.

const (
	// MockMinimumVersion is the minimum acceptable version for stripe-mock.
	// It's here so that if the library depends on new endpoints or features
	// added in a more recent version of stripe-mock, we can show people a
	// better error message instead of the test suite crashing with a bunch of
	// confusing 404 errors or the like.
	MockMinimumVersion = "0.4.0"

	// TestMerchantID is a token that can be used to represent a merchant ID in
	// simple tests.
	TestMerchantID = "acct_123"
)

func init() {
	// Enable strict mode on form encoding so that we'll panic if any kind of
	// malformed param struct is detected
	form.Strict = true

	port := os.Getenv("STRIPE_MOCK_PORT")
	if port == "" {
		port = "12111"
	}

	resp, err := http.Get("http://localhost:" + port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't reach stripe-mock at `localhost:%s`. Is "+
			"it running? Please see README for setup instructions.\n", port)
		os.Exit(1)
	}
	version := resp.Header.Get("Stripe-Mock-Version")
	if version != "master" && compareVersions(version, MockMinimumVersion) > 0 {
		fmt.Fprintf(os.Stderr, "Your version of stripe-mock (%s) is too old. The "+
			"minimum version to run this test suite is %s. Please see its "+
			"repository for upgrade instructions.\n", version, MockMinimumVersion)
		os.Exit(1)
	}

	stripe.Key = "sk_test_myTestKey"
	stripe.SetBackend("api", &stripe.BackendConfiguration{
		Type:       stripe.APIBackend,
		URL:        "http://localhost:" + port + "/v1",
		HTTPClient: &http.Client{},
	})
}

// compareVersions compares two semantic version strings. We need this because
// with more complex double-digit numbers, lexical comparison breaks down.
func compareVersions(a, b string) (ret int) {
	as := strings.Split(a, ".")
	bs := strings.Split(b, ".")

	loopMax := len(bs)
	if len(as) > len(bs) {
		loopMax = len(as)
	}

	for i := 0; i < loopMax; i++ {
		var x, y string
		if len(as) > i {
			x = as[i]
		}
		if len(bs) > i {
			y = bs[i]
		}

		xi, _ := strconv.Atoi(x)
		yi, _ := strconv.Atoi(y)

		if xi > yi {
			ret = -1
		} else if xi < yi {
			ret = 1
		}
		if ret != 0 {
			break
		}
	}
	return
}
