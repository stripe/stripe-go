//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Returns a list of your order returns. The returns are returned sorted by creation date, with the most recently created return appearing first.
type OrderReturnListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Order        *string           `form:"order"`
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
	Amount   int64        `json:"amount"`
	Created  int64        `json:"created"`
	Currency Currency     `json:"currency"`
	ID       string       `json:"id"`
	Items    []*OrderItem `json:"items"`
	Livemode bool         `json:"livemode"`
	Object   string       `json:"object"`
	Order    *Order       `json:"order"`
	Refund   *Refund      `json:"refund"`
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
