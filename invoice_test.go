package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/v76/form"
)

func TestInvoiceParams_AppendTo(t *testing.T) {
	{
		params := &InvoiceUpcomingParams{SubscriptionBillingCycleAnchorNow: Bool(true)}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"now"}, body.Get("subscription_billing_cycle_anchor"))
	}

	{
		params := &InvoiceUpcomingParams{SubscriptionBillingCycleAnchorUnchanged: Bool(true)}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"unchanged"}, body.Get("subscription_billing_cycle_anchor"))
	}

	{
		params := &InvoiceUpcomingParams{SubscriptionTrialEndNow: Bool(true)}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"now"}, body.Get("subscription_trial_end"))
	}
}

func TestInvoice_Unmarshal(t *testing.T) {
	invoiceDataUnexpanded := map[string]interface{}{
		"id":     "in_123",
		"object": "invoice",
		"discounts": []string{
			"dis_123",
			"dis_abc",
		},
		"total_discount_amounts": []map[string]interface{}{
			{
				"amount":   123,
				"discount": "dis_123",
			},
			{
				"amount":   345,
				"discount": "dis_abc",
			},
		},
	}

	bytes, err := json.Marshal(&invoiceDataUnexpanded)
	assert.NoError(t, err)

	var invoiceUnexpanded Invoice
	err = json.Unmarshal(bytes, &invoiceUnexpanded)
	assert.NoError(t, err)

	assert.Equal(t, "in_123", invoiceUnexpanded.ID)
	assert.Equal(t, "invoice", invoiceUnexpanded.Object)

	assert.Equal(t, 2, len(invoiceUnexpanded.Discounts))
	assert.Equal(t, "dis_123", invoiceUnexpanded.Discounts[0].ID)
	assert.Equal(t, "dis_abc", invoiceUnexpanded.Discounts[1].ID)

	assert.Equal(t, 2, len(invoiceUnexpanded.TotalDiscountAmounts))
	assert.Equal(t, int64(123), invoiceUnexpanded.TotalDiscountAmounts[0].Amount)
	assert.Equal(t, "dis_123", invoiceUnexpanded.TotalDiscountAmounts[0].Discount.ID)
	assert.Equal(t, int64(345), invoiceUnexpanded.TotalDiscountAmounts[1].Amount)
	assert.Equal(t, "dis_abc", invoiceUnexpanded.TotalDiscountAmounts[1].Discount.ID)

	invoiceDataExpanded := map[string]interface{}{
		"id":     "in_123",
		"object": "invoice",
		"discounts": []map[string]interface{}{
			{
				"id":     "dis_123",
				"object": "discount",
				"coupon": map[string]interface{}{
					"id":          "co_123",
					"object":      "coupon",
					"currency":    "usd",
					"percent_off": 25.5,
				},
			},
			{
				"id":     "dis_abc",
				"object": "discount",
				"coupon": map[string]interface{}{
					"id":          "co_abc",
					"object":      "coupon",
					"currency":    "eur",
					"percent_off": 35.5,
				},
			},
		},
		"total_discount_amounts": []map[string]interface{}{
			{
				"amount": 123,
				"discount": map[string]interface{}{
					"id":     "dis_123",
					"object": "discount",
					"coupon": map[string]interface{}{
						"id":          "co_123",
						"object":      "coupon",
						"currency":    "usd",
						"percent_off": 25.5,
					},
				},
			},
			{
				"amount": 345,
				"discount": map[string]interface{}{
					"id":     "dis_abc",
					"object": "discount",
					"coupon": map[string]interface{}{
						"id":          "co_abc",
						"object":      "coupon",
						"currency":    "eur",
						"percent_off": 35.5,
					},
				},
			},
		},
	}

	bytes2, err2 := json.Marshal(&invoiceDataExpanded)
	assert.NoError(t, err2)

	var invoiceExpanded Invoice
	err = json.Unmarshal(bytes2, &invoiceExpanded)
	assert.NoError(t, err)

	assert.Equal(t, "in_123", invoiceExpanded.ID)
	assert.Equal(t, "invoice", invoiceExpanded.Object)

	assert.Equal(t, 2, len(invoiceExpanded.Discounts))
	assert.Equal(t, "dis_123", invoiceExpanded.Discounts[0].ID)
	assert.Equal(t, "discount", invoiceExpanded.Discounts[0].Object)
	assert.Equal(t, "co_123", invoiceExpanded.Discounts[0].Coupon.ID)
	assert.Equal(t, "coupon", invoiceExpanded.Discounts[0].Coupon.Object)
	assert.Equal(t, float64(25.5), invoiceExpanded.Discounts[0].Coupon.PercentOff)
	assert.Equal(t, "dis_abc", invoiceExpanded.Discounts[1].ID)
	assert.Equal(t, "discount", invoiceExpanded.Discounts[1].Object)
	assert.Equal(t, "co_abc", invoiceExpanded.Discounts[1].Coupon.ID)
	assert.Equal(t, "coupon", invoiceExpanded.Discounts[1].Coupon.Object)
	assert.Equal(t, float64(35.5), invoiceExpanded.Discounts[1].Coupon.PercentOff)

	assert.Equal(t, 2, len(invoiceExpanded.TotalDiscountAmounts))
	assert.Equal(t, int64(123), invoiceExpanded.TotalDiscountAmounts[0].Amount)
	assert.Equal(t, "dis_123", invoiceExpanded.TotalDiscountAmounts[0].Discount.ID)
	assert.Equal(t, "discount", invoiceExpanded.TotalDiscountAmounts[0].Discount.Object)
	assert.Equal(t, "co_123", invoiceExpanded.TotalDiscountAmounts[0].Discount.Coupon.ID)
	assert.Equal(t, "coupon", invoiceExpanded.TotalDiscountAmounts[0].Discount.Coupon.Object)
	assert.Equal(t, float64(25.5), invoiceExpanded.TotalDiscountAmounts[0].Discount.Coupon.PercentOff)
	assert.Equal(t, int64(345), invoiceExpanded.TotalDiscountAmounts[1].Amount)
	assert.Equal(t, "dis_abc", invoiceExpanded.TotalDiscountAmounts[1].Discount.ID)
	assert.Equal(t, "discount", invoiceExpanded.TotalDiscountAmounts[1].Discount.Object)
	assert.Equal(t, "co_abc", invoiceExpanded.TotalDiscountAmounts[1].Discount.Coupon.ID)
	assert.Equal(t, "coupon", invoiceExpanded.TotalDiscountAmounts[1].Discount.Coupon.Object)
	assert.Equal(t, float64(35.5), invoiceExpanded.TotalDiscountAmounts[1].Discount.Coupon.PercentOff)
}

func TestInvoice_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v Invoice
		err := json.Unmarshal([]byte(`"in_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "in_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := Invoice{ID: "in_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "in_123", v.ID)
	}
}
