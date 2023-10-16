package setupattempt

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v75"
	_ "github.com/stripe/stripe-go/v75/testing"
)

func TestSetupAttemptList(t *testing.T) {
	params := &stripe.SetupAttemptListParams{
		SetupIntent: stripe.String("seti_123"),
	}
	i := List(params)

	// Verify that we can get at least one setup attempt
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.SetupAttempt())
	assert.NotNil(t, i.SetupAttemptList())
}
