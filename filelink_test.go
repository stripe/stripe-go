package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/v72/form"
)

func TestFileLink_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v FileLink
		err := json.Unmarshal([]byte(`"link_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "link_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := FileLink{ID: "link_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "link_123", v.ID)
	}
}

func TestFileLinkParams_AppendTo(t *testing.T) {
	{
		params := &FileLinkParams{ExpiresAtNow: Bool(true)}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"now"}, body.Get("expires_at"))
	}
}
