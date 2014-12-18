// Package discount provides the discount-related APIs
package discount

import (
	"fmt"

	stripe "github.com/stripe/stripe-go"
)

// Client is used to invoke discount-related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Del removes a discount from a customer.
// For more details see https://stripe.com/docs/api#delete_discount.
func Del(customerID string) error {
	return getC().Del(customerID)
}

func (c Client) Del(customerID string) error {
	return c.B.Call("DELETE", fmt.Sprintf("/customers/%v/discount", customerID), c.Key, nil, nil, nil)
}

// DelSub removes a discount from a customer's subscription.
// For more details see https://stripe.com/docs/api#delete_subscription_discount.
func DelSub(customerID, subscriptionID string) error {
	return getC().DelSub(customerID, subscriptionID)
}

func (c Client) DelSub(customerID, subscriptionID string) error {
	return c.B.Call("DELETE", fmt.Sprintf("/customers/%v/subscriptions/%v/discount", customerID, subscriptionID), c.Key, nil, nil, nil)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
