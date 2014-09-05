package discount

import (
	"fmt"

	. "github.com/stripe/stripe-go"
)

// Client is used to invoke discount-related APIs.
type Client struct {
	B     Backend
	Token string
}

var c *Client

// Delete removes a discount from a customer.
// For more details see https://stripe.com/docs/api#delete_discount.
func Delete(customerId string) error {
	refresh()
	return c.Delete(customerId)
}

func (c *Client) Delete(customerId string) error {
	return c.B.Call("DELETE", fmt.Sprintf("/customers/%v/discount", customerId), c.Token, nil, nil)
}

// DeleteSubscription removes a discount from a customer's subscription.
// For more details see https://stripe.com/docs/api#delete_subscription_discount.
func DeleteSubscription(customerId, subscriptionId string) error {
	refresh()
	return c.DeleteSubscription(customerId, subscriptionId)
}

func (c *Client) DeleteSubscription(customerId, subscriptionId string) error {
	return c.B.Call("DELETE", fmt.Sprintf("/customers/%v/subscriptions/%v/discount", customerId, subscriptionId), c.Token, nil, nil)
}

func refresh() {
	if c == nil {
		c = &Client{B: GetBackend()}
	}

	c.Token = Key
}
