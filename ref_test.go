package stripe

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

// refTestObject is a minimal struct used as the type parameter in Ref tests.
// It embeds APIResource so that *refTestObject satisfies LastResponseSetter,
// matching what all real Stripe API types do.
type refTestObject struct {
	APIResource
	ID   string `json:"id"`
	Type string `json:"type"`
}

// refMockBackend is a minimal Backend implementation for testing Ref.Fetch.
type refMockBackend struct {
	// callMethod, callPath, and callKey record the arguments passed to Call.
	callMethod string
	callPath   string
	callKey    string
	// callResult is JSON-unmarshaled into the v argument when Call is invoked.
	callResult []byte
	// callErr is returned by Call if non-nil.
	callErr error
}

func (m *refMockBackend) Call(method, path, key string, params ParamsContainer, v LastResponseSetter) error {
	m.callMethod = method
	m.callPath = path
	m.callKey = key
	if m.callErr != nil {
		return m.callErr
	}
	if m.callResult != nil {
		return json.Unmarshal(m.callResult, v)
	}
	return nil
}

func (m *refMockBackend) CallStreaming(method, path, key string, params ParamsContainer, v StreamingLastResponseSetter) error {
	return nil
}

func (m *refMockBackend) CallRaw(method, path, key string, body []byte, params *Params, v LastResponseSetter) error {
	return nil
}

func (m *refMockBackend) CallMultipart(method, path, key, boundary string, body *bytes.Buffer, params *Params, v LastResponseSetter) error {
	return nil
}

func (m *refMockBackend) SetMaxNetworkRetries(maxNetworkRetries int64) {}

func TestRefJSONUnmarshal(t *testing.T) {
	raw := `{"id":"obj_123","type":"v2.billing.meter","url":"/v2/billing/meters/obj_123"}`

	var ref Ref[refTestObject]
	err := json.Unmarshal([]byte(raw), &ref)

	assert.NoError(t, err)
	assert.Equal(t, "obj_123", ref.ID)
	assert.Equal(t, "v2.billing.meter", ref.Type)
	assert.Equal(t, "/v2/billing/meters/obj_123", ref.URL)
	// backend is not set via JSON; it must remain nil
	assert.Nil(t, ref.backend)
}

func TestRefFetchNilBackend(t *testing.T) {
	ref := &Ref[refTestObject]{
		ID:   "obj_123",
		Type: "v2.billing.meter",
		URL:  "/v2/billing/meters/obj_123",
	}

	result, err := ref.Fetch(context.Background())

	assert.Nil(t, result)
	assert.EqualError(t, err, "stripe: Ref.Fetch called on a Ref without an attached backend; use stripe.Client to obtain Ref values")
}

func TestRefSetBackend(t *testing.T) {
	ref := &Ref[refTestObject]{
		ID:   "obj_123",
		Type: "v2.billing.meter",
		URL:  "/v2/billing/meters/obj_123",
	}

	mock := &refMockBackend{}
	ref.SetBackend(mock, "sk_test_abc")

	assert.Equal(t, mock, ref.backend)
	assert.Equal(t, "sk_test_abc", ref.key)
}

func TestRefFetchCallsBackend(t *testing.T) {
	resultJSON := `{"id":"obj_123","type":"v2.billing.meter"}`
	mock := &refMockBackend{callResult: []byte(resultJSON)}

	ref := &Ref[refTestObject]{
		ID:   "obj_123",
		Type: "v2.billing.meter",
		URL:  "/v2/billing/meters/obj_123",
	}
	ref.SetBackend(mock, "sk_test_abc")

	obj, err := ref.Fetch(context.Background())

	assert.NoError(t, err)
	assert.NotNil(t, obj)
	assert.Equal(t, "obj_123", obj.ID)
	assert.Equal(t, "GET", mock.callMethod)
	assert.Equal(t, "/v2/billing/meters/obj_123", mock.callPath)
	assert.Equal(t, "sk_test_abc", mock.callKey)
}
