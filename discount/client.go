// package discount provides the discount-related APIs
package discount

import (
	"fmt"

	. "github.com/stripe/stripe-go"
)

// Client is used to invoke discount-related APIs.
type Client struct {
	B   Backend
	Key string
}

// Del removes a discount from a customer.
// For more details see https://stripe.com/docs/api#delete_discount.
func Del(customerId string) error {
	return getC().Del(customerId)
}

func (c Client) Del(customerId string) error {
	return c.B.Call("DELETE", fmt.Sprintf("/customers/%v/discount", customerId), c.Key, nil, nil)
}

// DelSub removes a discount from a customer's subscription.
// For more details see https://stripe.com/docs/api#delete_subscription_discount.
func DelSub(customerId, subscriptionId string) error {
	return getC().DelSub(customerId, subscriptionId)
}

func (c Client) DelSub(customerId, subscriptionId string) error {
	return c.B.Call("DELETE", fmt.Sprintf("/customers/%v/subscriptions/%v/discount", customerId, subscriptionId), c.Key, nil, nil)
}

func getC() Client {
	return Client{GetBackend(), Key}
}
