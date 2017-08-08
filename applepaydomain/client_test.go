package applepaydomain

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestApplePayDomainDel(t *testing.T) {
	domain, err := Del("apwc_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, domain)
}

func TestApplePayDomainGet(t *testing.T) {
	domain, err := Get("apwc_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, domain)
}

func TestApplePayDomainList(t *testing.T) {
	i := List(&stripe.ApplePayDomainListParams{})

	// Verify that we can get at least one domain
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.ApplePayDomain())
}

func TestApplePayDomainNew(t *testing.T) {
	domain, err := New(&stripe.ApplePayDomainParams{
		DomainName: "example.com",
	})
	assert.Nil(t, err)
	assert.NotNil(t, domain)
}
