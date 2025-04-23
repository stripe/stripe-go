package testing

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	"github.com/max-cape/stripe-go-test/form"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/http2"
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
	MockMinimumVersion = "0.109.0"

	// TestMerchantID is a token that can be used to represent a merchant ID in
	// simple tests.
	TestMerchantID = "acct_123"

	// TestAPIKey is an API key that can be used against stripe-mock in tests.
	TestAPIKey = "sk_test_myTestKey"
)

func init() {
	// Enable strict mode on form encoding so that we'll panic if any kind of
	// malformed param struct is detected
	form.Strict = true

	port := os.Getenv("STRIPE_MOCK_PORT")
	if port == "" {
		port = "12112"
	}

	// stripe-mock's certificate for localhost is self-signed so configure a
	// specialized client that skips the certificate authority check.
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	// Go can often enable HTTP/2 automatically if it's supported, but
	// confusingly, if you set `TLSClientConfig`, it disables it and you have
	// to explicitly invoke http2's `ConfigureTransport` to get it back.
	//
	// See the incorrectly closed bug report here:
	//
	//     https://github.com/golang/go/issues/20645
	//
	err := http2.ConfigureTransport(transport)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize HTTP/2 transport: %v\n", err)
		os.Exit(1)
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	resp, err := httpClient.Get("https://localhost:" + port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't reach stripe-mock at `localhost:%s` (%v). Is "+
			"it running? Please see README for setup instructions.\n", port, err)
		os.Exit(1)
	}
	version := resp.Header.Get("Stripe-Mock-Version")
	if version != "master" && compareVersions(version, MockMinimumVersion) > 0 {
		fmt.Fprintf(os.Stderr, "Your version of stripe-mock (%s) is too old. The "+
			"minimum version to run this test suite is %s. Please see its "+
			"repository for upgrade instructions.\n", version, MockMinimumVersion)
		os.Exit(1)
	}

	stripe.Key = TestAPIKey

	// Configure a backend for stripe-mock and set it for both the API and
	// Uploads (unlike the real Stripe API, stripe-mock supports both these
	// backends).
	stripeMockBackend := stripe.GetBackendWithConfig(
		stripe.APIBackend,
		&stripe.BackendConfig{
			URL:           stripe.String("https://localhost:" + port),
			HTTPClient:    httpClient,
			LeveledLogger: stripe.DefaultLeveledLogger,
		},
	)
	stripe.SetBackend(stripe.APIBackend, stripeMockBackend)
	stripe.SetBackend(stripe.UploadsBackend, stripeMockBackend)
}

func MockServer(t *testing.T, method, path string, params interface{}, resp string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, method)
		assert.Equal(t, r.URL.Path, path)
		assert.Equal(t, r.Header.Get("Authorization"), "Bearer "+TestAPIKey)

		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		if len(body) > 0 {
			assert.NoError(t, json.Unmarshal(body, params))
		}
		data := []byte(resp)
		_, err = w.Write(data)
		assert.NoError(t, err)
	}))
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
