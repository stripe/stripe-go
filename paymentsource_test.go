package stripe

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/form"
)

func TestSourceParams_AppendTo(t *testing.T) {
	{
		params := &SourceParams{Token: "tok_123"}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"tok_123"}, body.Get("source"))
	}

	{
		params := &SourceParams{Card: &CardParams{Number: "4242424242424242"}}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"4242424242424242"}, body.Get("source[number]"))
		assert.Equal(t, []string{"card"}, body.Get("source[object]"))
	}
}
