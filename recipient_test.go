package stripe

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/form"
)

func TestRecipientParams_AppendTo(t *testing.T) {
	{
		params := &RecipientParams{Token: "card_123"}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"card_123"}, body.Get("card"))
	}

	{
		params := &RecipientParams{Card: &CardParams{Name: "A Card"}}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"A Card"}, body.Get("card[name]"))
	}

	{
		params := &RecipientParams{Bank: &BankAccountParams{Token: "ba_123"}}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"ba_123"}, body.Get("bank_account"))
	}

	{
		params := &RecipientParams{Bank: &BankAccountParams{Account: "123"}}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"123"}, body.Get("bank_account[account_number]"))
	}
}
