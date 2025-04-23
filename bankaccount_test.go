package stripe

import (
	"encoding/json"
	"testing"

	"github.com/max-cape/stripe-go-test/form"
	assert "github.com/stretchr/testify/require"
)

func TestBankAccount_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v BankAccount
		err := json.Unmarshal([]byte(`"ba_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "ba_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := BankAccount{ID: "ba_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "ba_123", v.ID)
	}
}

func TestBankAccountListParams_AppendTo(t *testing.T) {
	// Adds `object` for account (this will hit the customer sources endpoint)
	{
		params := &BankAccountListParams{Account: String("acct_123")}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"bank_account"}, body.Get("object"))
	}

	// Adds `object` for customer (this will hit the external accounts endpoint)
	{
		params := &BankAccountListParams{Customer: String("cus_123")}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"bank_account"}, body.Get("object"))
	}
}

func TestBankAccountParams_AppendToAsSourceOrExternalAccount(t *testing.T) {
	// We should add more tests for all the various corner cases here ...

	// Includes account_holder_name
	{
		params := &BankAccountParams{AccountHolderName: String("Tyrion")}
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
		params := &BankAccountParams{AccountHolderType: String(string(BankAccountAccountHolderTypeIndividual))}
		body := &form.Values{}
		params.AppendToAsSourceOrExternalAccount(body)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"individual"}, body.Get("external_account[account_holder_type]"))
	}

	// Includes Params
	{
		params := &BankAccountParams{
			Params: Params{
				Metadata: map[string]string{
					"foo": "bar",
				},
			},
		}
		body := &form.Values{}
		params.AppendToAsSourceOrExternalAccount(body)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"bar"}, body.Get("metadata[foo]"))
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
