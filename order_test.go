package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/form"
)

func TestOrderUpdateParams_AppendTo(t *testing.T) {
	{
		params := &OrderUpdateParams{Status: "fulfilled", Shipping: &OrderUpdateShippingParams{Carrier: "USPS", TrackingNumber: "123"}}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"fulfilled"}, body.Get("status"))
		assert.Equal(t, []string{"USPS"}, body.Get("shipping[carrier]"))
		assert.Equal(t, []string{"123"}, body.Get("shipping[tracking_number]"))
	}
}

func TestShipping_MarshalJSON(t *testing.T) {
	{
		shipping := &Shipping{
			Name:           "name",
			Phone:          "phone",
			Carrier:        "USPS",
			TrackingNumber: "tracking.123",
			Address: Address{
				Line1:   "123 Market Street",
				City:    "San Francisco",
				State:   "CA",
				Country: "USA",
			},
		}

		d, err := json.Marshal(shipping)
		assert.NoError(t, err)
		assert.NotNil(t, d)

		unmarshalled := &Shipping{}
		err = json.Unmarshal(d, unmarshalled)
		assert.NoError(t, err)

		assert.Equal(t, unmarshalled.Name, shipping.Name)
		assert.Equal(t, unmarshalled.Phone, shipping.Phone)
		assert.Equal(t, unmarshalled.Carrier, shipping.Carrier)
		assert.Equal(t, unmarshalled.TrackingNumber, shipping.TrackingNumber)
		assert.NotNil(t, unmarshalled.Address)
		assert.Equal(t, unmarshalled.Address.Line1, shipping.Address.Line1)
		assert.Equal(t, unmarshalled.Address.City, shipping.Address.City)
		assert.Equal(t, unmarshalled.Address.State, shipping.Address.State)
		assert.Equal(t, unmarshalled.Address.Country, shipping.Address.Country)
	}
}
