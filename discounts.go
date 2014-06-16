package stripe

import (
	"fmt"
)

// Discount is the resource representing a Stripe discount.
// For more details see https://stripe.com/docs/api#discounts.
type Discount struct {
	Coupon   *Coupon `json:"coupon"`
	Customer string  `json:"customer"`
	Start    int64   `json:"start"`
	End      int64   `json:"end"`
	Sub      string  `json:"subscription"`
}

// DiscountClient is the client used to invoke discount-related APIs.
type DiscountClient struct {
	api   Api
	token string
}

// Delete removes a discount from a customer.
// For more details see https://stripe.com/docs/api#delete_discount.
func (c *DiscountClient) Delete(customerId string) error {
	return c.api.Call("DELETE", fmt.Sprintf("/customers/%v/discount", customerId), c.token, nil, nil)
}

// DeleteSubscription removes a discount from a customer's subscription.
// For more details see https://stripe.com/docs/api#delete_subscription_discount.
func (c *DiscountClient) DeleteSubscription(customerId, subscriptionid string) error {
	return c.api.Call("DELETE", fmt.Sprintf("/customers/%v/subscriptions/%v/discount", customerId, subscriptionid), c.token, nil, nil)
}
