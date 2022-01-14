//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Returns a list of your order returns. The returns are returned sorted by creation date, with the most recently created return appearing first.
type OrderReturnListParams struct {
	ListParams `form:"*"`
	// Date this return was created.
	Created *int64 `form:"created"`
	// Date this return was created.
	CreatedRange *RangeQueryParams `form:"created"`
	// The order to retrieve returns for.
	Order *string `form:"order"`
}

// Retrieves the details of an existing order return. Supply the unique order ID from either an order return creation request or the order return list, and Stripe will return the corresponding order information.
type OrderReturnParams struct {
	Params `form:"*"`
	Items  []*OrderItemParams `form:"items"`
	Order  *string            `form:"-"` // Included in the URL
}

// A return represents the full or partial return of a number of [order items](https://stripe.com/docs/api#order_items).
// Returns always belong to an order, and may optionally contain a refund.
//
// Related guide: [Handling Returns](https://stripe.com/docs/orders/guide#handling-returns).
type OrderReturn struct {
	APIResource
	// A positive integer in the smallest currency unit (that is, 100 cents for $1.00, or 1 for ¥1, Japanese Yen being a zero-decimal currency) representing the total amount for the returned line item.
	Amount int64 `json:"amount"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The items included in this order return.
	Items []*OrderItem `json:"items"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The order that this return includes items from.
	Order *Order `json:"order"`
	// The ID of the refund issued for this return.
	Refund *Refund `json:"refund"`
}

// OrderReturnList is a list of OrderReturns as retrieved from a list endpoint.
type OrderReturnList struct {
	APIResource
	ListMeta
	Data []*OrderReturn `json:"data"`
}

// UnmarshalJSON handles deserialization of an OrderReturn.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (o *OrderReturn) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		o.ID = id
		return nil
	}

	type orderReturn OrderReturn
	var v orderReturn
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*o = OrderReturn(v)
	return nil
}
