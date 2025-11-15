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
	assert.Equal(t, "include[0]=foo&include[1]=bar", body.Encode())
}
