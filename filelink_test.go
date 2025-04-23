package stripe

import (
	"testing"

	"github.com/max-cape/stripe-go-test/form"
	assert "github.com/stretchr/testify/require"
)

func TestFileLinkParams_AppendTo(t *testing.T) {
	{
		params := &FileLinkParams{ExpiresAtNow: Bool(true)}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"now"}, body.Get("expires_at"))
	}
}
