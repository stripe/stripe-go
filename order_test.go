package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/form"
)

func TestOrder_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v Order
		err := json.Unmarshal([]byte(`"or_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "or_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := Order{ID: "or_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "or_123", v.ID)
	}
}

func TestOrderItem_UnmarshalJSON(t *testing.T) {
	// Try unmarshaling a SKU order item
	{
		orderItemData := map[string]interface{}{
			"object": "order_item",
			"amount": 123,
			"parent": "TEST-SKU-123",
			"type":   "sku",
		}
		bytes, err := json.Marshal(&orderItemData)
		assert.NoError(t, err)

		var orderItem OrderItem
		err = json.Unmarshal(bytes, &orderItem)
		assert.NoError(t, err)
		assert.Equal(t, "TEST-SKU-123", orderItem.Parent.ID)
	}

	// Try unmarshaling a SKU order item with parent expanded
	{
		orderItemData := map[string]interface{}{
			"object": "order_item",
			"amount": 123,
			"parent": map[string]interface{}{
				"id":     "TEST-SKU-123",
				"object": "sku",
			},
			"type": "sku",
		}
		bytes, err := json.Marshal(&orderItemData)
		assert.NoError(t, err)

		var orderItem OrderItem
		err = json.Unmarshal(bytes, &orderItem)
		assert.NoError(t, err)
		assert.Equal(t, "TEST-SKU-123", orderItem.Parent.ID)
		assert.Equal(t, OrderItemParentTypeSKU, orderItem.Parent.Type)
		assert.Equal(t, "TEST-SKU-123", orderItem.Parent.SKU.ID)
	}

	// Try unmarshaling a coupon order item
	{
		orderItemData := map[string]interface{}{
			"object": "order_item",
			"amount": 0,
			"parent": "TEST-COUPON-123",
			"type":   "coupon",
		}
		bytes, err := json.Marshal(&orderItemData)
		assert.NoError(t, err)

		var orderItem OrderItem
		err = json.Unmarshal(bytes, &orderItem)
		assert.NoError(t, err)
		assert.Equal(t, "TEST-COUPON-123", orderItem.Parent.ID)
	}

	// Try unmarshaling a shipping order item
	{
		orderItemData := map[string]interface{}{
			"object": "order_item",
			"amount": 1000,
			"parent": "ship_MZmIpV7v14QZLlRR",
			"type":   "shipping",
		}
		bytes, err := json.Marshal(&orderItemData)
		assert.NoError(t, err)

		var orderItem OrderItem
		err = json.Unmarshal(bytes, &orderItem)
		assert.NoError(t, err)
		assert.Equal(t, "ship_MZmIpV7v14QZLlRR", orderItem.Parent.ID)
	}
}

func TestOrderUpdateParams_AppendTo(t *testing.T) {
	{
		params := &OrderUpdateParams{
			Status: String("fulfilled"),
			Shipping: &OrderUpdateShippingParams{
				Carrier:        String("USPS"),
				TrackingNumber: String("123"),
			},
		}
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
			Address: &Address{
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
