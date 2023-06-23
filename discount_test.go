package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestDiscount_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON object
	{
		data := "{\"coupon\":{\"id\":\"co_123\"}, \"id\":\"dp_123\"}\n"
		dataBytes := []byte(data)

		var v Discount
		err := json.Unmarshal(dataBytes, &v)
		assert.NoError(t, err)
		assert.Equal(t, "dp_123", v.ID)
		assert.Equal(t, "co_123", v.Coupon.ID)

		//stripe.accounts.delete("acct_123", nil)
		//DeletedResource:
		//{
		//	"id": "acct_123",
		//	"deleted": "true"
		//}
	}
}
