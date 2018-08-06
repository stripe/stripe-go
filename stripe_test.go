package stripe_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"runtime"
	"sync"
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	. "github.com/stripe/stripe-go/testing"
)

func TestBearerAuth(t *testing.T) {
	c := stripe.GetBackend(stripe.APIBackend).(*stripe.BackendImplementation)
	key := "apiKey"

	req, err := c.NewRequest("", "", key, "", nil)
	assert.NoError(t, err)

	assert.Equal(t, "Bearer "+key, req.Header.Get("Authorization"))
}

func TestContext(t *testing.T) {
	c := stripe.GetBackend(stripe.APIBackend).(*stripe.BackendImplementation)
	p := &stripe.Params{Context: context.Background()}

	req, err := c.NewRequest("", "", "", "", p)
	assert.NoError(t, err)

	assert.Equal(t, p.Context, req.Context())
}

func TestContext_Cancel(t *testing.T) {
	c := stripe.GetBackend(stripe.APIBackend).(*stripe.BackendImplementation)
	ctx, cancel := context.WithCancel(context.Background())
	p := &stripe.Params{Context: ctx}

	req, err := c.NewRequest("", "", "", "", p)
	assert.NoError(t, err)

	assert.Equal(t, ctx, req.Context())

	// Cancel the context before we even try to start the request. This will
	// cause it to immediately return an error and also avoid any kind of race
	// condition.
	cancel()

	var v interface{}
	err = c.Do(req, nil, &v)

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

// Tests client retries.
//
// You can get pretty good visibility into what's going on by running just this
// test on verbose:
//
//     go test . -run TestDo_Retry -test.v
//
func TestDo_Retry(t *testing.T) {
	type testServerResponse struct {
		Message string `json:"message"`
	}

	message := "Hello, client."
	requestNum := 0

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		assert.NoError(t, err)

		// The body should always be the same with every retry. We've
		// previously had regressions in this behavior as we switched to HTTP/2
		// and `Request` became non-reusable, so we want to check it with every
		// request.
		assert.Equal(t, "bar", r.Form.Get("foo"))

		switch requestNum {
		case 0:
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(`{"error":"Conflict (this should be retried)."}`))

		case 1:
			response := testServerResponse{Message: message}

			data, err := json.Marshal(response)
			assert.NoError(t, err)

			_, err = w.Write(data)
			assert.NoError(t, err)

		default:
			assert.Fail(t, "Should not have reached request %v", requestNum)
		}

		requestNum++
	}))
	defer testServer.Close()

	backend := stripe.GetBackendWithConfig(
		stripe.APIBackend,
		&stripe.BackendConfig{
			LogLevel:          3,
			MaxNetworkRetries: 5,
			URL:               testServer.URL,
		},
	).(*stripe.BackendImplementation)

	// Disable sleeping duration our tests.
	backend.SetNetworkRetriesSleep(false)

	request, err := backend.NewRequest(
		http.MethodPost,
		"/hello",
		"sk_test_123",
		"application/x-www-form-urlencoded",
		nil,
	)
	assert.NoError(t, err)

	bodyBuffer := bytes.NewBufferString("foo=bar")
	var response testServerResponse
	err = backend.Do(request, bodyBuffer, &response)

	assert.NoError(t, err)
	assert.Equal(t, message, response.Message)

	// We should have seen exactly two requests.
	assert.Equal(t, 2, requestNum)
}

func TestFormatURLPath(t *testing.T) {
	assert.Equal(t, "/resources/1/subresources/2",
		stripe.FormatURLPath("/resources/%s/subresources/%s", "1", "2"))

	// Tests that each parameter is escaped for use in URLs
	assert.Equal(t, "/resources/%25",
		stripe.FormatURLPath("/resources/%s", "%"))
}

func TestParseID(t *testing.T) {
	// JSON string
	{
		id, ok := stripe.ParseID([]byte(`"ch_123"`))
		assert.Equal(t, "ch_123", id)
		assert.True(t, ok)
	}

	// JSON object
	{
		id, ok := stripe.ParseID([]byte(`{"id":"ch_123"}`))
		assert.Equal(t, "", id)
		assert.False(t, ok)
	}

	// Other JSON scalar (this should never be used, but check the results anyway)
	{
		id, ok := stripe.ParseID([]byte(`123`))
		assert.Equal(t, "", id)
		assert.False(t, ok)
	}
}

// TestMultipleAPICalls will fail the test run if a race condition is thrown while running multiple NewRequest calls.
func TestMultipleAPICalls(t *testing.T) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c := stripe.GetBackend(stripe.APIBackend).(*stripe.BackendImplementation)
			key := "apiKey"

			req, err := c.NewRequest("", "", key, "", nil)
			assert.NoError(t, err)

			assert.Equal(t, "Bearer "+key, req.Header.Get("Authorization"))
		}()
	}
	wg.Wait()
}

func TestIdempotencyKey(t *testing.T) {
	c := stripe.GetBackend(stripe.APIBackend).(*stripe.BackendImplementation)
	p := &stripe.Params{IdempotencyKey: stripe.String("idempotency-key")}

	req, err := c.NewRequest("", "", "", "", p)
	assert.NoError(t, err)

	assert.Equal(t, "idempotency-key", req.Header.Get("Idempotency-Key"))
}

func TestNewBackends(t *testing.T) {
	httpClient := &http.Client{}
	backends := stripe.NewBackends(httpClient)
	assert.Equal(t, httpClient, backends.API.(*stripe.BackendImplementation).HTTPClient)
	assert.Equal(t, httpClient, backends.Uploads.(*stripe.BackendImplementation).HTTPClient)
}

func TestStripeAccount(t *testing.T) {
	c := stripe.GetBackend(stripe.APIBackend).(*stripe.BackendImplementation)
	p := &stripe.Params{}
	p.SetStripeAccount(TestMerchantID)

	req, err := c.NewRequest("", "", "", "", p)
	assert.NoError(t, err)

	assert.Equal(t, TestMerchantID, req.Header.Get("Stripe-Account"))
}

func TestUserAgent(t *testing.T) {
	c := stripe.GetBackend(stripe.APIBackend).(*stripe.BackendImplementation)

	req, err := c.NewRequest("", "", "", "", nil)
	assert.NoError(t, err)

	// We keep out version constant private to the package, so use a regexp
	// match instead.
	expectedPattern := regexp.MustCompile(`^Stripe/v1 GoBindings/[1-9][0-9.]+[0-9]$`)

	match := expectedPattern.MatchString(req.Header.Get("User-Agent"))
	assert.True(t, match)
}

func TestUserAgentWithAppInfo(t *testing.T) {
	appInfo := &stripe.AppInfo{
		Name:      "MyAwesomePlugin",
		PartnerID: "partner_1234",
		URL:       "https://myawesomeplugin.info",
		Version:   "1.2.34",
	}
	stripe.SetAppInfo(appInfo)
	defer stripe.SetAppInfo(nil)

	c := stripe.GetBackend(stripe.APIBackend).(*stripe.BackendImplementation)

	req, err := c.NewRequest("", "", "", "", nil)
	assert.NoError(t, err)

	//
	// User-Agent
	//

	// We keep out version constant private to the package, so use a regexp
	// match instead.
	expectedPattern := regexp.MustCompile(`^Stripe/v1 GoBindings/[1-9][0-9.]+[0-9] MyAwesomePlugin/1.2.34 \(https://myawesomeplugin.info\)$`)

	match := expectedPattern.MatchString(req.Header.Get("User-Agent"))
	assert.True(t, match)

	//
	// X-Stripe-Client-User-Agent
	//

	encodedUserAgent := req.Header.Get("X-Stripe-Client-User-Agent")
	assert.NotEmpty(t, encodedUserAgent)

	var userAgent map[string]interface{}
	err = json.Unmarshal([]byte(encodedUserAgent), &userAgent)
	assert.NoError(t, err)

	application := userAgent["application"].(map[string]interface{})

	assert.Equal(t, "MyAwesomePlugin", application["name"])
	assert.Equal(t, "partner_1234", application["partner_id"])
	assert.Equal(t, "https://myawesomeplugin.info", application["url"])
	assert.Equal(t, "1.2.34", application["version"])
}

func TestStripeClientUserAgent(t *testing.T) {
	c := stripe.GetBackend(stripe.APIBackend).(*stripe.BackendImplementation)

	req, err := c.NewRequest("", "", "", "", nil)
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

	c := stripe.GetBackend(stripe.APIBackend).(*stripe.BackendImplementation)

	req, err := c.NewRequest("", "", "", "", nil)
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
	c := stripe.GetBackend(stripe.APIBackend).(*stripe.BackendImplementation)

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
		Code:  stripe.ErrorCodeMissing,
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
