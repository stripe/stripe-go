package applepaydomain

import (
	"testing"

	stripe "github.com/stripe/stripe-go"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	stripe.Key = GetTestKey()
}

func TestApplePayDomain(t *testing.T) {
	// Test creating an ApplePay Domain
	applePayDomainParams := &stripe.ApplePayDomainParams{
		DomainName: "test.com",
	}

	domain, err := New(applePayDomainParams)

	if err != nil {
		t.Error(err)
	}

	if domain.DomainName != applePayDomainParams.DomainName {
		t.Errorf("DomainName %v does not match expected DomainName %v\n", domain.DomainName, applePayDomainParams.DomainName)
	}

	// Test listing all ApplePayDomain(s)
	params := &stripe.ApplePayDomainListParams{}
	params.Filters.AddFilter("limit", "", "1")
	params.Single = true

	i := List(params)
	for i.Next() {
		if i.ApplePayDomain() == nil {
			t.Error("No nil values expected")
		}
	}
	if err := i.Err(); err != nil {
		t.Error(err)
	}

	// Test retrieving an ApplePayDomain
	domain2, err := Get(domain.ID, nil)

	if err != nil {
		t.Error(err)
	}

	if domain2.ID != domain.ID {
		t.Errorf("ID %q does not match expected id %q\n", domain2.ID, domain.ID)
	}

	// Testing to delete an ApplePayDomain
	domain3, err := Del(domain2.ID)
	if !domain3.Deleted {
		t.Errorf("Domain id %v expected to be marked as deleted on the returned resource\n", domain3.ID)
	}
}
