package stripe

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/form"
)

func TestBankAccountParams_AppendToAsSourceOrExternalAccount(t *testing.T) {
	// We should add more tests for all the various corner cases here ...

	// Includes account_holder_name
	{
		params := &BankAccountParams{AccountHolderName: "Tyrion"}
		body := &form.Values{}
		params.AppendToAsSourceOrExternalAccount(body)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"Tyrion"}, body.Get("external_account[account_holder_name]"))
	}

	// Does not include account_holder_name if empty
	{
		params := &BankAccountParams{}
		body := &form.Values{}
		params.AppendToAsSourceOrExternalAccount(body)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string(nil), body.Get("external_account[account_holder_name]"))
	}

	// Includes account_holder_name
	{
		params := &BankAccountParams{AccountHolderType: "individual"}
		body := &form.Values{}
		params.AppendToAsSourceOrExternalAccount(body)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"individual"}, body.Get("external_account[account_holder_type]"))
	}

	// Does not include account_holder_name if empty
	{
		params := &BankAccountParams{}
		body := &form.Values{}
		params.AppendToAsSourceOrExternalAccount(body)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string(nil), body.Get("external_account[account_holder_type]"))
	}
}
