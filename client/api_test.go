package client

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestAPIInit(t *testing.T) {
	api := API{}
	api.Init("sk_test_123", nil)
	assert.Equal(t, "sk_test_123", api.Charges.Key)
}

func TestAPINew(t *testing.T) {
	api := New("sk_test_123", nil)
	assert.Equal(t, "sk_test_123", api.Charges.Key)
}
