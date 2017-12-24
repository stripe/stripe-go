package stripe_test

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"
	"runtime"
	"sync"
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	. "github.com/stripe/stripe-go/testing"
)

func TestBearerAuth(t *testing.T) {
	c := &stripe.BackendConfiguration{URL: stripe.APIURL}
	key := "apiKey"

	req, err := c.NewRequest("", "", key, "", nil, nil)
	assert.NoError(t, err)

	assert.Equal(t, "Bearer "+key, req.Header.Get("Authorization"))
}

func TestContext(t *testing.T) {
	c := &stripe.BackendConfiguration{URL: stripe.APIURL}
	p := &stripe.Params{Context: context.Background()}

	req, err := c.NewRequest("", "", "", "", nil, p)
	assert.NoError(t, err)

	assert.Equal(t, p.Context, req.Context())
}

func TestContext_Cancel(t *testing.T) {
	c := &stripe.BackendConfiguration{
		HTTPClient: &http.Client{},
		URL:        stripe.APIURL,
	}
	ctx, cancel := context.WithCancel(context.Background())
	p := &stripe.Params{Context: ctx}

	req, err := c.NewRequest("", "", "", "", nil, p)
	assert.NoError(t, err)

	assert.Equal(t, ctx, req.Context())

	// Cancel the context before we even try to start the request. This will
	// cause it to immediately return an error and also avoid any kind of race
	// condition.
	cancel()

	var v interface{}
	err = c.Do(req, &v)

	// Go 1.7 will produce an error message like:
	//
	//     Get https://api.stripe.com/v1/: net/http: request canceled while waiting for connection
	//
	// 1.8 and later produce something like:
	//
	//     Get https://api.stripe.com/v1/: context canceled
	//
	// When we drop support for 1.7 we can remove the first case of this
	// expression.
	assert.Regexp(t, regexp.MustCompile(`(request canceled|context canceled\z)`), err.Error())
}

// TestMultipleAPICalls will fail the test run if a race condition is thrown while running multiple NewRequest calls.
func TestMultipleAPICalls(t *testing.T) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c := &stripe.BackendConfiguration{URL: stripe.APIURL}
			key := "apiKey"

			req, err := c.NewRequest("", "", key, "", nil, nil)
			assert.NoError(t, err)

			assert.Equal(t, "Bearer "+key, req.Header.Get("Authorization"))
		}()
	}
	wg.Wait()
}

func TestIdempotencyKey(t *testing.T) {
	c := &stripe.BackendConfiguration{URL: stripe.APIURL}
	p := &stripe.Params{IdempotencyKey: "idempotency-key"}

	req, err := c.NewRequest("", "", "", "", nil, p)
	assert.NoError(t, err)

	assert.Equal(t, "idempotency-key", req.Header.Get("Idempotency-Key"))
}

func TestStripeAccount(t *testing.T) {
	c := &stripe.BackendConfiguration{URL: stripe.APIURL}
	p := &stripe.Params{StripeAccount: TestMerchantID}

	req, err := c.NewRequest("", "", "", "", nil, p)
	assert.NoError(t, err)

	assert.Equal(t, TestMerchantID, req.Header.Get("Stripe-Account"))

	// Also test the deprecated Account field for now as well. This should be
	// identical to the exercise above.
	p = &stripe.Params{Account: TestMerchantID}

	req, err = c.NewRequest("", "", "", "", nil, p)
	assert.NoError(t, err)

	assert.Equal(t, TestMerchantID, req.Header.Get("Stripe-Account"))
}

func TestUserAgent(t *testing.T) {
	c := &stripe.BackendConfiguration{URL: stripe.APIURL}

	req, err := c.NewRequest("", "", "", "", nil, nil)
	assert.NoError(t, err)

	// We keep out version constant private to the package, so use a regexp
	// match instead.
	expectedPattern := regexp.MustCompile(`^Stripe/v1 GoBindings/[1-9][0-9.]+[0-9]$`)

	match := expectedPattern.MatchString(req.Header.Get("User-Agent"))
	assert.True(t, match)
}

func TestUserAgentWithAppInfo(t *testing.T) {
	appInfo := &stripe.AppInfo{
		Name:    "MyAwesomePlugin",
		URL:     "https://myawesomeplugin.info",
		Version: "1.2.34",
	}
	stripe.SetAppInfo(appInfo)
	defer stripe.SetAppInfo(nil)

	c := &stripe.BackendConfiguration{URL: stripe.APIURL}

	req, err := c.NewRequest("", "", "", "", nil, nil)
	assert.NoError(t, err)

	// We keep out version constant private to the package, so use a regexp
	// match instead.
	expectedPattern := regexp.MustCompile(`^Stripe/v1 GoBindings/[1-9][0-9.]+[0-9] MyAwesomePlugin/1.2.34 \(https://myawesomeplugin.info\)$`)

	match := expectedPattern.MatchString(req.Header.Get("User-Agent"))
	assert.True(t, match)
}

func TestStripeClientUserAgent(t *testing.T) {
	c := &stripe.BackendConfiguration{URL: stripe.APIURL}

	req, err := c.NewRequest("", "", "", "", nil, nil)
	assert.NoError(t, err)

	encodedUserAgent := req.Header.Get("X-Stripe-Client-User-Agent")
	assert.NotEmpty(t, encodedUserAgent)

	var userAgent map[string]string
	err = json.Unmarshal([]byte(encodedUserAgent), &userAgent)
	assert.NoError(t, err)

	//
	// Just test a few headers that we know to be stable.
	//

	assert.Empty(t, userAgent["application"])
	assert.Equal(t, "go", userAgent["lang"])
	assert.Equal(t, runtime.Version(), userAgent["lang_version"])

	// Anywhere these tests are running can reasonable be expected to have a
	// `uname` to run, so do this basic check.
	assert.NotEqual(t, stripe.UnknownPlatform, userAgent["lang_version"])
}

func TestStripeClientUserAgentWithAppInfo(t *testing.T) {
	appInfo := &stripe.AppInfo{
		Name:    "MyAwesomePlugin",
		URL:     "https://myawesomeplugin.info",
		Version: "1.2.34",
	}
	stripe.SetAppInfo(appInfo)
	defer stripe.SetAppInfo(nil)

	c := &stripe.BackendConfiguration{URL: stripe.APIURL}

	req, err := c.NewRequest("", "", "", "", nil, nil)
	assert.NoError(t, err)

	encodedUserAgent := req.Header.Get("X-Stripe-Client-User-Agent")
	assert.NotEmpty(t, encodedUserAgent)

	var userAgent map[string]interface{}
	err = json.Unmarshal([]byte(encodedUserAgent), &userAgent)
	assert.NoError(t, err)

	decodedAppInfo := userAgent["application"].(map[string]interface{})
	assert.Equal(t, appInfo.Name, decodedAppInfo["name"])
	assert.Equal(t, appInfo.URL, decodedAppInfo["url"])
	assert.Equal(t, appInfo.Version, decodedAppInfo["version"])
}

func TestResponseToError(t *testing.T) {
	c := &stripe.BackendConfiguration{URL: stripe.APIURL}

	// A test response that includes a status code and request ID.
	res := &http.Response{
		Header: http.Header{
			"Request-Id": []string{"request-id"},
		},
		StatusCode: 402,
	}

	// An error that contains expected fields which we're going to serialize to
	// JSON and inject into our conversion function.
	expectedErr := &stripe.Error{
		Code:  stripe.Missing,
		Msg:   "That card was declined",
		Param: "expiry_date",
		Type:  stripe.ErrorTypeCard,
	}
	bytes, err := json.Marshal(expectedErr)
	assert.NoError(t, err)

	// Unpack the error that we just serialized so that we can inject a
	// type-specific field into it ("decline_code"). This will show up in a
	// field on a special stripe.CardError type which is attached to the common
	// stripe.Error.
	var raw map[string]string
	err = json.Unmarshal(bytes, &raw)
	assert.NoError(t, err)

	expectedDeclineCode := "decline-code"
	raw["decline_code"] = expectedDeclineCode
	bytes, err = json.Marshal(raw)
	assert.NoError(t, err)

	// A generic Golang error.
	err = c.ResponseToError(res, wrapError(bytes))

	// An error containing Stripe-specific fields that we cast back from the
	// generic Golang error.
	stripeErr := err.(*stripe.Error)

	assert.Equal(t, expectedErr.Code, stripeErr.Code)
	assert.Equal(t, expectedErr.Msg, stripeErr.Msg)
	assert.Equal(t, expectedErr.Param, stripeErr.Param)
	assert.Equal(t, res.Header.Get("Request-Id"), stripeErr.RequestID)
	assert.Equal(t, res.StatusCode, stripeErr.HTTPStatusCode)
	assert.Equal(t, expectedErr.Type, stripeErr.Type)

	// Just a bogus type coercion to demonstrate how this code might be
	// written. Because we've assigned ErrorTypeCard as the error's type, Err
	// should always come out as a CardError.
	_, ok := stripeErr.Err.(*stripe.InvalidRequestError)
	assert.False(t, ok)

	cardErr, ok := stripeErr.Err.(*stripe.CardError)
	assert.True(t, ok)
	assert.Equal(t, expectedDeclineCode, cardErr.DeclineCode)
}

//
// ---
//

// A simple function that allows us to represent an error response from Stripe
// which comes wrapper in a JSON object with a single field of "error".
func wrapError(serialized []byte) []byte {
	return []byte(`{"error":` + string(serialized) + `}`)
}
