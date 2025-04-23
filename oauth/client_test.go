package oauth

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestAuthorizeURL(t *testing.T) {
	url := AuthorizeURL(&stripe.AuthorizeURLParams{})

	assert.Equal(t, url, "https://connect.stripe.com/oauth/authorize?")
}

func TestAuthorizeURLWithOptionalArgs(t *testing.T) {
	url := AuthorizeURL(&stripe.AuthorizeURLParams{
		ClientID:              stripe.String("ca_123"),
		State:                 stripe.String("test-state"),
		Scope:                 stripe.String("read_only"),
		RedirectURI:           stripe.String("https://t.example.com"),
		ResponseType:          stripe.String("test-code"),
		StripeLanding:         stripe.String("register"),
		AlwaysPrompt:          stripe.Bool(true),
		Express:               stripe.Bool(true),
		SuggestedCapabilities: stripe.StringSlice([]string{"card_payments"}),
	})

	assert.Contains(t, url, "https://connect.stripe.com/express/oauth/authorize?")
	assert.Contains(t, url, "client_id=ca_123")
	assert.Contains(t, url, "redirect_uri=https%3A%2F%2Ft.example.com")
	assert.Contains(t, url, "response_type=test-code")
	assert.Contains(t, url, "scope=read_only")
	assert.Contains(t, url, "state=test-state")
	assert.Contains(t, url, "stripe_landing=register")
	assert.Contains(t, url, "always_prompt=true")
	assert.Contains(t, url, "suggested_capabilities[0]=card_payments")
}

func TestAuthorizeURLWithStripeUser(t *testing.T) {
	url := AuthorizeURL(&stripe.AuthorizeURLParams{
		ResponseType: stripe.String("test-code"),
		StripeUser: &stripe.OAuthStripeUserParams{
			BlockKana:          stripe.String("block-kana"),
			BlockKanji:         stripe.String("block-kanji"),
			BuildingKana:       stripe.String("building-kana"),
			BuildingKanji:      stripe.String("building-kanji"),
			BusinessName:       stripe.String("b-name"),
			BusinessType:       stripe.String("llc"),
			City:               stripe.String("Elko"),
			Country:            stripe.String("US"),
			Currency:           stripe.String("USD"),
			DOBDay:             stripe.Int64(15),
			DOBMonth:           stripe.Int64(10),
			DOBYear:            stripe.Int64(2019),
			Email:              stripe.String("test@example.com"),
			FirstName:          stripe.String("first-name"),
			FirstNameKana:      stripe.String("first-name-kana"),
			FirstNameKanji:     stripe.String("first-name-kanji"),
			Gender:             stripe.String("female"),
			LastName:           stripe.String("last-name"),
			LastNameKana:       stripe.String("last-name-kana"),
			LastNameKanji:      stripe.String("last-name-kanji"),
			PhoneNumber:        stripe.String("999-999-9999"),
			PhysicalProduct:    stripe.Bool(false),
			ProductDescription: stripe.String("product-description"),
			State:              stripe.String("NV"),
			StreetAddress:      stripe.String("123 main"),
			URL:                stripe.String("http://example.com"),
			Zip:                stripe.String("12345"),
		},
	})

	assert.Contains(t, url, "response_type=test-code")
	assert.Contains(t, url, "stripe_user[block_kana]=block-kana")
	assert.Contains(t, url, "stripe_user[block_kanji]=block-kanji")
	assert.Contains(t, url, "stripe_user[building_kana]=building-kana")
	assert.Contains(t, url, "stripe_user[building_kanji]=building-kanji")
	assert.Contains(t, url, "stripe_user[business_name]=b-name")
	assert.Contains(t, url, "stripe_user[business_type]=llc")
	assert.Contains(t, url, "stripe_user[city]=Elko")
	assert.Contains(t, url, "stripe_user[country]=US")
	assert.Contains(t, url, "stripe_user[currency]=USD")
	assert.Contains(t, url, "stripe_user[dob_day]=15")
	assert.Contains(t, url, "stripe_user[dob_month]=10")
	assert.Contains(t, url, "stripe_user[dob_year]=2019")
	assert.Contains(t, url, "stripe_user[email]=test%40example.com")
	assert.Contains(t, url, "stripe_user[first_name]=first-name")
	assert.Contains(t, url, "stripe_user[first_name_kana]=first-name-kana")
	assert.Contains(t, url, "stripe_user[first_name_kanji]=first-name-kanji")
	assert.Contains(t, url, "stripe_user[gender]=female")
	assert.Contains(t, url, "stripe_user[last_name]=last-name")
	assert.Contains(t, url, "stripe_user[last_name_kana]=last-name-kana")
	assert.Contains(t, url, "stripe_user[last_name_kanji]=last-name-kanji")
	assert.Contains(t, url, "stripe_user[phone_number]=999-999-9999")
	assert.Contains(t, url, "stripe_user[physical_product]=false")
	assert.Contains(t, url, "stripe_user[product_description]=product-description")
	assert.Contains(t, url, "stripe_user[state]=NV")
	assert.Contains(t, url, "stripe_user[street_address]=123+main")
	assert.Contains(t, url, "stripe_user[url]=http%3A%2F%2Fexample.com")
	assert.Contains(t, url, "stripe_user[zip]=12345")
}

type roundTripFunc func(req *http.Request) *http.Response

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func newTestClient(fn roundTripFunc) *http.Client {
	return &http.Client{
		Transport: roundTripFunc(fn),
	}
}

func stubConnectBackend(httpClient *http.Client) {
	mockBackend := stripe.GetBackendWithConfig(
		stripe.ConnectBackend,
		&stripe.BackendConfig{
			URL:        stripe.String("https://localhost:12113"),
			HTTPClient: httpClient,
		},
	)
	stripe.SetBackend(stripe.ConnectBackend, mockBackend)
}

func TestNewOAuthToken(t *testing.T) {
	stripe.Key = "sk_123"

	// stripe-mock doesn't support connect URLs so this stubs out the server.
	httpClient := newTestClient(func(req *http.Request) *http.Response {
		buf := new(bytes.Buffer)
		buf.ReadFrom(req.Body)
		reqBody := buf.String()

		assert.Contains(t, req.URL.String(), "https://localhost:12113/oauth/token")
		assert.Contains(t, reqBody, "client_secret=sk_123")
		assert.Contains(t, reqBody, "grant_type=authorization_code")
		assert.Contains(t, reqBody, "code=code")
		assert.Contains(t, reqBody, "assert_capabilities[0]=card_payments")

		responseBody := `{
      "access_token":"sk_123",
      "livemode":false,
      "refresh_token":"rt_123",
      "token_type":"bearer",
      "stripe_publishable_key":"pk_123",
      "stripe_user_id":"acct_123",
      "scope":"read_write"
    }`
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(responseBody)),
			Header:     make(http.Header),
		}
	})
	stubConnectBackend(httpClient)

	token, err := New(&stripe.OAuthTokenParams{
		GrantType:          stripe.String("authorization_code"),
		Code:               stripe.String("code"),
		AssertCapabilities: stripe.StringSlice([]string{"card_payments"}),
	})
	assert.Nil(t, err)
	assert.NotNil(t, token)
	assert.Equal(t, token.AccessToken, "sk_123")
	assert.Equal(t, token.Livemode, false)
	assert.Equal(t, token.RefreshToken, "rt_123")
	assert.Equal(t, token.TokenType, stripe.OAuthTokenTypeBearer)
	assert.Equal(t, token.StripePublishableKey, "pk_123")
	assert.Equal(t, token.StripeUserID, "acct_123")
	assert.Equal(t, token.Scope, stripe.OAuthScopeTypeReadWrite)
}

func TestNewOAuthTokenWithCustomKey(t *testing.T) {
	stripe.Key = "sk_123"
	// stripe-mock doesn't support connect URLs so this stubs out the server.
	httpClient := newTestClient(func(req *http.Request) *http.Response {
		buf := new(bytes.Buffer)
		buf.ReadFrom(req.Body)
		reqBody := buf.String()
		assert.Contains(t, reqBody, "client_secret=sk_999")

		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`{}`)),
			Header:     make(http.Header),
		}
	})
	stubConnectBackend(httpClient)

	token, err := New(&stripe.OAuthTokenParams{
		ClientSecret: stripe.String("sk_999"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestNewOAuthTokenWithError(t *testing.T) {
	stripe.Key = "sk_123"
	// stripe-mock doesn't support connect URLs so this stubs out the server.

	responseBody := `{"error":"invalid_grant","error_description": "Authorization code does not exist"}`
	httpClient := newTestClient(func(req *http.Request) *http.Response {
		buf := new(bytes.Buffer)
		buf.ReadFrom(req.Body)
		reqBody := buf.String()
		assert.Contains(t, reqBody, "client_secret=sk_999")

		return &http.Response{
			StatusCode: 400,
			Body:       ioutil.NopCloser(bytes.NewBufferString(responseBody)),
			Header:     make(http.Header),
		}
	})
	stubConnectBackend(httpClient)

	token, err := New(&stripe.OAuthTokenParams{
		ClientSecret: stripe.String("sk_999"),
	})

	assert.NotNil(t, token)
	assert.NotNil(t, err)

	stripeErr := err.(*stripe.Error)
	assert.Equal(t, 400, stripeErr.HTTPStatusCode)
	assert.Equal(t, "Authorization code does not exist", stripeErr.OAuthErrorDescription)
	assert.Equal(t, "invalid_grant", stripeErr.OAuthError)
}

func TestDeauthorize(t *testing.T) {
	stripe.Key = "sk_123"

	// stripe-mock doesn't support connect URLs so this stubs out the server.
	httpClient := newTestClient(func(req *http.Request) *http.Response {
		buf := new(bytes.Buffer)
		buf.ReadFrom(req.Body)
		reqBody := buf.String()
		assert.Contains(t, req.URL.String(), "https://localhost:12113/oauth/deauthorize")
		assert.Contains(t, reqBody, "client_id=sk_999")
		assert.Contains(t, reqBody, "stripe_user_id=acct_123")

		resBody := `{"stripe_user_id": "acct_123"}`
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(resBody)),
			Header:     make(http.Header),
		}
	})
	stubConnectBackend(httpClient)

	deauthorization, err := Del(&stripe.DeauthorizeParams{
		ClientID:     stripe.String("sk_999"),
		StripeUserID: stripe.String("acct_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, deauthorization)
	assert.Equal(t, deauthorization.StripeUserID, "acct_123")
}
