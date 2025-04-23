package stripe

import (
	"encoding/json"
	"testing"

	"github.com/max-cape/stripe-go-test/form"
	assert "github.com/stretchr/testify/require"
)

func TestCardListParams_AppendTo(t *testing.T) {
	// Adds `object` for account (this will hit the external accounts endpoint)
	{
		params := &CardListParams{Account: String("acct_123")}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"card"}, body.Get("object"))
	}

	// Adds `object` for customer (this will hit the sources endpoint)
	{
		params := &CardListParams{Customer: String("cus_123")}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"card"}, body.Get("object"))
	}
}

func TestCard_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v Card
		err := json.Unmarshal([]byte(`"card_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "card_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := Card{ID: "card_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "card_123", v.ID)
	}
}

func TestCardParams_AppendToAsCardSourceOrExternalAccount(t *testing.T) {
	// We should add more tests for all the various corner cases here ...

	// Includes number and object
	{
		params := &CardParams{Number: String("1234")}
		body := &form.Values{}
		params.AppendToAsCardSourceOrExternalAccount(body, nil)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"1234"}, body.Get("source[number]"))
		assert.Equal(t, []string{"card"}, body.Get("source[object]"))
	}

	// Includes Params
	{
		params := &CardParams{
			Params: Params{
				Metadata: map[string]string{
					"foo": "bar",
				},
			},
		}
		body := &form.Values{}
		params.AppendToAsCardSourceOrExternalAccount(body, nil)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"bar"}, body.Get("metadata[foo]"))
	}

	// It takes key parts for deeper embedding
	{
		params := &CardParams{Number: String("1234")}
		body := &form.Values{}
		params.AppendToAsCardSourceOrExternalAccount(body, []string{"prefix1", "prefix2"})
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"1234"}, body.Get("prefix1[prefix2][source][number]"))
	}
}
