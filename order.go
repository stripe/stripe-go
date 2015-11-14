package stripe

import "encoding/json"

type OrderParams struct {
	Params
	Currency string
	Customer string
	Email    string
	Items    []OrderItemParams
	Shipping Shipping
}

type OrderItemParams struct {
	Amount      int64
	Currency    string
	Description string
	Parent      string
	Quantity    *int64
	Type        string
}

type Shipping struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
	Phone   string  `json:"phone"`
}

type ShippingMethod struct {
	ID          string `json:"id"`
	Amount      int64  `json:"amount"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
}

type Order struct {
	ID                     string            `json:"id"`
	Amount                 int64             `json:"amount"`
	Application            string            `json:"application"`
	ApplicationFee         int64             `json:"application_fee"`
	Charge                 Charge            `json:"charge"`
	Created                int64             `json:"created"`
	Currency               string            `json:"currency"`
	Customer               Customer          `json:"customer"`
	Email                  string            `json:"email"`
	Items                  []OrderItem       `json:"items"`
	Meta                   map[string]string `json:"metadata"`
	SelectedShippingMethod *string           `json:"selected_shipping_method"`
	Shipping               Shipping          `json:"shipping"`
	ShippingMethods        []ShippingMethod  `json:"shipping_methods"`
	Status                 string            `json:"status"`
	Updated                int64             `json:"updated"`
}

type OrderItem struct {
	Amount      int64  `json:"amount"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
	Parent      string `json:"parent"`
	Quantity    int64  `json:"quantity"`
	Type        string `json:"type"`
}

// UnmarshalJSON handles deserialization of an Order.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (o *Order) UnmarshalJSON(data []byte) error {
	type order Order
	var oo order
	err := json.Unmarshal(data, &oo)
	if err == nil {
		*o = Order(oo)
	} else {
		// the id is surrounded by "\" characters, so strip them
		o.ID = string(data[1 : len(data)-1])
	}

	return nil
}
