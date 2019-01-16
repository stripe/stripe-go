package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/form"
)

func TestRecipientParams_AppendTo(t *testing.T) {
	{
		params := &RecipientParams{Token: String("card_123")}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"card_123"}, body.Get("card"))
	}

	{
		params := &RecipientParams{Card: &CardParams{Name: String("A Card")}}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"A Card"}, body.Get("card[name]"))
	}

	{
		params := &RecipientParams{BankAccount: &BankAccountParams{Token: String("ba_123")}}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"ba_123"}, body.Get("bank_account"))
	}

	{
		params := &RecipientParams{BankAccount: &BankAccountParams{AccountNumber: String("123")}}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"123"}, body.Get("bank_account[account_number]"))
	}
}

func TestRecipient_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v Recipient
		err := json.Unmarshal([]byte(`"rp_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "rp_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := Recipient{ID: "rp_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "rp_123", v.ID)
	}
}
