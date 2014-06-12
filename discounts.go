package stripe

import (
	"fmt"
)

type Discount struct {
	Coupon   *Coupon `json:"coupon"`
	Customer string  `json:customer"`
	Start    int64   `json:"start"`
	End      int64   `json:end"`
	Sub      string  `json:subscription"`
}

type DiscountClient struct {
	api   Api
	token string
}

func (c *DiscountClient) Delete(customerId string) error {
	_, err := c.api.Call("DELETE", fmt.Sprintf("/customers/%v/discount", customerId), c.token, nil)
	return err
}

func (c *DiscountClient) DeleteSubscription(customerId, subscriptionid string) error {
	_, err := c.api.Call("DELETE", fmt.Sprintf("/customers/%v/subscriptions/%v/discount", customerId, subscriptionid), c.token, nil)
	return err
}
