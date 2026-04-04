package stripe

import (
	"context"
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

// refTestObject is a minimal struct used as the type parameter in Ref tests.
type refTestObject struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func TestRefJSONUnmarshal(t *testing.T) {
	raw := `{"id":"obj_123","type":"v2.billing.meter","url":"/v2/billing/meters/obj_123"}`

	var ref Ref[refTestObject]
	err := json.Unmarshal([]byte(raw), &ref)

	assert.NoError(t, err)
	assert.Equal(t, "obj_123", ref.ID)
	assert.Equal(t, "v2.billing.meter", ref.Type)
	assert.Equal(t, "/v2/billing/meters/obj_123", ref.URL)
	// client is not set via JSON; it must remain nil
	assert.Nil(t, ref.client)
}

func TestRefFetchNilClient(t *testing.T) {
	ref := &Ref[refTestObject]{
		ID:   "obj_123",
		Type: "v2.billing.meter",
		URL:  "/v2/billing/meters/obj_123",
	}

	result, err := ref.Fetch(context.Background())

	assert.Nil(t, result)
	assert.EqualError(t, err, "stripe: Ref.Fetch called on a Ref without an attached client; use stripe.Client to obtain Ref values")
}
