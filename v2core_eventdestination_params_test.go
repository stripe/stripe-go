package stripe

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/v83/form"
)

func TestV2CoreEventDestinationParams_AppendTo(t *testing.T) {
	params := &V2CoreEventDestinationParams{
		Include: []*string{
			String("foo"),
			String("bar"),
		},
	}
	body := &form.Values{}
	form.AppendTo(body, params)
	assert.Equal(t, []string{"foo", "bar"}, body.Get("include"))
	assert.Equal(t, "include=foo&include=bar", body.Encode())
}
