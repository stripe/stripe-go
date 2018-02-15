package stripe

import (
	"encoding/json"
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

func TestPaymentSource_MarshalJSON(t *testing.T) {
	{
		id := "card_123"
		name := "alice cooper"
		paymentSource := &PaymentSource{
			Type: PaymentSourceCard,
			ID:   id,
			Card: &Card{
				ID:   id,
				Name: name,
			},
		}

		d, err := json.Marshal(paymentSource)
		assert.NoError(t, err)
		assert.NotNil(t, d)

		unmarshalled := &PaymentSource{}
		err = json.Unmarshal(d, unmarshalled)
		assert.NoError(t, err)

		assert.Equal(t, unmarshalled.ID, id)
		assert.NotNil(t, unmarshalled.Card)
		assert.Equal(t, unmarshalled.Card.ID, id)
		assert.Equal(t, unmarshalled.Card.Name, name)
	}

	{
		id := "ba_123"
		name := "big bank"
		paymentSource := &PaymentSource{
			Type: PaymentSourceBankAccount,
			ID:   id,
			BankAccount: &BankAccount{
				ID:   id,
				Name: name,
			},
		}

		d, err := json.Marshal(paymentSource)
		assert.NoError(t, err)
		assert.NotNil(t, d)

		unmarshalled := &PaymentSource{}
		err = json.Unmarshal(d, unmarshalled)
		assert.NoError(t, err)

		assert.Equal(t, unmarshalled.ID, id)
		assert.NotNil(t, unmarshalled.BankAccount)
		assert.Equal(t, unmarshalled.BankAccount.ID, id)
		assert.Equal(t, unmarshalled.BankAccount.Name, name)
	}
}
