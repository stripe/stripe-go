package stripe

import (
	"encoding/json"
	"testing"

	"github.com/max-cape/stripe-go-test/form"
	assert "github.com/stretchr/testify/require"
)

func TestSourceParams_AppendTo(t *testing.T) {
	{
		params := &PaymentSourceSourceParams{Token: String("tok_123")}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"tok_123"}, body.Get("source"))
	}

	{
		params := &PaymentSourceSourceParams{Card: &CardParams{Number: String("4242424242424242")}}
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
			Type: PaymentSourceTypeCard,
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
			Type: PaymentSourceTypeBankAccount,
			ID:   id,
			BankAccount: &BankAccount{
				ID:                id,
				AccountHolderName: name,
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
		assert.Equal(t, unmarshalled.BankAccount.AccountHolderName, name)
	}
}

func TestPaymentSource_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v PaymentSource
		err := json.Unmarshal([]byte(`"ba_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "ba_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		// We build the JSON object manually here because it's key that the
		// `object` field is included so that the source knows what type to
		// decode
		data := []byte(`{"id":"ba_123", "object":"bank_account"}`)

		var v PaymentSource
		err := json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, PaymentSourceTypeBankAccount, v.Type)

		// The payment source has a field for each possible type, so the bank
		// account is located one level down
		assert.Equal(t, "ba_123", v.BankAccount.ID)
	}
}
