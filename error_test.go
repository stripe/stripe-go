package stripe

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestErrorError(t *testing.T) {
	err := &Error{Type: "foo", Msg: "bar"}
	assert.Equal(t, `{"message":"bar","type":"foo"}`, err.Error())
}

func TestErrorResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Request-Id", "req_123")
		w.WriteHeader(http.StatusUnauthorized)
		_, err := fmt.Fprintln(w, `{"error":{"message":"bar","type":"`+ErrorTypeInvalidRequest+`"}}`)
		assert.NoError(t, err)
	}))
	defer ts.Close()

	backend := GetBackendWithConfig(APIBackend, &BackendConfig{
		// Suppress error log output to make a verbose run of this test less
		// alarming (because we're testing specifically for an error).
		LeveledLogger: &LeveledLogger{Level: LevelNull},

		URL: String(ts.URL),
	})

	err := backend.Call(http.MethodGet, "/v1/account", "sk_test_badKey", nil, nil)
	assert.Error(t, err)

	stripeErr := err.(*Error)
	assert.Equal(t, ErrorTypeInvalidRequest, stripeErr.Type)
	assert.Equal(t, "req_123", stripeErr.RequestID)
	assert.Equal(t, 401, stripeErr.HTTPStatusCode)
	var invalidRequestErr *InvalidRequestError
	assert.True(t, errors.As(err, &invalidRequestErr))
}

func TestErrorNetworkAdviceCode(t *testing.T) {
	raw := `{"error":{"type":"card_error","network_advice_code":"02"}}`
	var wrapper struct {
		Error *Error `json:"error"`
	}
	err := json.Unmarshal([]byte(raw), &wrapper)
	assert.NoError(t, err)
	assert.Equal(t, "02", wrapper.Error.NetworkAdviceCode)
}

func TestErrorPaymentIntent(t *testing.T) {
	raw := `{"error":{"type":"card_error","payment_intent":{"id":"pi_123","amount":1000}}}`
	var wrapper struct {
		Error *Error `json:"error"`
	}
	err := json.Unmarshal([]byte(raw), &wrapper)
	assert.NoError(t, err)
	assert.NotNil(t, wrapper.Error.PaymentIntent)
	assert.Equal(t, "pi_123", wrapper.Error.PaymentIntent.ID)
	assert.Equal(t, int64(1000), wrapper.Error.PaymentIntent.Amount)
}

func TestErrorRedact(t *testing.T) {
	pi := &PaymentIntent{Amount: int64(400), ClientSecret: "foo"}
	si := &SetupIntent{Description: "keepme", ClientSecret: "foo"}

	t.Run("BothIntentObjects", func(t *testing.T) {
		err := &Error{PaymentIntent: pi, SetupIntent: si}
		redacted := err.redact()
		assert.Equal(t, int64(400), err.PaymentIntent.Amount)
		assert.Equal(t, int64(400), redacted.(*Error).PaymentIntent.Amount)
		assert.Equal(t, "keepme", err.SetupIntent.Description)
		assert.Equal(t, "keepme", redacted.(*Error).SetupIntent.Description)
		assert.Equal(t, "foo", err.PaymentIntent.ClientSecret)
		assert.Equal(t, "foo", err.SetupIntent.ClientSecret)
		assert.Equal(t, "foo", pi.ClientSecret)
		assert.Equal(t, "foo", si.ClientSecret)
		assert.Equal(t, "REDACTED", redacted.(*Error).PaymentIntent.ClientSecret)
		assert.Equal(t, "REDACTED", redacted.(*Error).SetupIntent.ClientSecret)
	})

	t.Run("NeitherIntentObject", func(t *testing.T) {
		err := Error{PaymentIntent: nil, SetupIntent: nil}
		redacted := err.redact()
		assert.Nil(t, err.PaymentIntent)
		assert.Nil(t, redacted.(*Error).PaymentIntent)
	})

	t.Run("PaymentIntentAlone", func(t *testing.T) {
		err := &Error{PaymentIntent: pi}
		redacted := err.redact()
		assert.Equal(t, int64(400), err.PaymentIntent.Amount)
		assert.Equal(t, int64(400), redacted.(*Error).PaymentIntent.Amount)
		assert.Equal(t, "foo", err.PaymentIntent.ClientSecret)
		assert.Equal(t, "foo", pi.ClientSecret)
		assert.Equal(t, "REDACTED", redacted.(*Error).PaymentIntent.ClientSecret)
	})

	t.Run("SetupIntentAlone", func(t *testing.T) {
		err := &Error{SetupIntent: si}
		redacted := err.redact()
		assert.Equal(t, "keepme", err.SetupIntent.Description)
		assert.Equal(t, "keepme", redacted.(*Error).SetupIntent.Description)
		assert.Equal(t, "foo", err.SetupIntent.ClientSecret)
		assert.Equal(t, "foo", si.ClientSecret)
		assert.Equal(t, "REDACTED", redacted.(*Error).SetupIntent.ClientSecret)
	})
}
