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

// Delete removes a discount from a customer.
// For more details see https://stripe.com/docs/api#delete_discount.
func Delete(customerId string) error {
	return getC().Delete(customerId)
}

func (c Client) Delete(customerId string) error {
	return c.B.Call("DELETE", fmt.Sprintf("/customers/%v/discount", customerId), c.Key, nil, nil)
}

// DeleteSubscription removes a discount from a customer's subscription.
// For more details see https://stripe.com/docs/api#delete_subscription_discount.
func DeleteSubscription(customerId, subscriptionId string) error {
	return getC().DeleteSubscription(customerId, subscriptionId)
}

func (c Client) DeleteSubscription(customerId, subscriptionId string) error {
	return c.B.Call("DELETE", fmt.Sprintf("/customers/%v/subscriptions/%v/discount", customerId, subscriptionId), c.Key, nil, nil)
}

func getC() Client {
	return Client{GetBackend(), Key}
}
