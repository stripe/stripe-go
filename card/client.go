package card

import (
	"errors"
	"fmt"
	"net/url"

	. "github.com/stripe/stripe-go"
)

// Client is used to invoke /cards APIs.
type Client struct {
	B   Backend
	Tok string
}

// Create POSTs new cards either for a customer or recipient.
// For more details see https://stripe.com/docs/api#create_card.
func Create(params *CardParams) (*Card, error) {
	return getC().Create(params)
}

func (c Client) Create(params *CardParams) (*Card, error) {
	body := &url.Values{}
	params.AppendDetails(body, true)
	params.AppendTo(body)

	card := &Card{}
	var err error

	if len(params.Customer) > 0 {
		err = c.B.Call("POST", fmt.Sprintf("/customers/%v/cards", params.Customer), c.Tok, body, card)
	} else if len(params.Recipient) > 0 {
		err = c.B.Call("POST", fmt.Sprintf("/recipients/%v/cards", params.Recipient), c.Tok, body, card)
	} else {
		err = errors.New("Invalid card params: either customer or recipient need to be set")
	}

	return card, err
}

// Get returns the details of a card.
// For more details see https://stripe.com/docs/api#retrieve_card.
func Get(id string, params *CardParams) (*Card, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *CardParams) (*Card, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}
		params.AppendTo(body)
	}

	card := &Card{}
	var err error

	if len(params.Customer) > 0 {
		err = c.B.Call("GET", fmt.Sprintf("/customers/%v/cards/%v", params.Customer, id), c.Tok, body, card)
	} else if len(params.Recipient) > 0 {
		err = c.B.Call("GET", fmt.Sprintf("/recipients/%v/cards/%v", params.Recipient, id), c.Tok, body, card)
	} else {
		err = errors.New("Invalid card params: either customer or recipient need to be set")
	}

	return card, err
}

// Update updates a card's properties.
// For more details see	https://stripe.com/docs/api#update_card.
func Update(id string, params *CardParams) (*Card, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *CardParams) (*Card, error) {
	body := &url.Values{}
	params.AppendDetails(body, false)
	params.AppendTo(body)

	card := &Card{}
	var err error

	if len(params.Customer) > 0 {
		err = c.B.Call("POST", fmt.Sprintf("/customers/%v/cards/%v", params.Customer, id), c.Tok, body, card)
	} else if len(params.Recipient) > 0 {
		err = c.B.Call("POST", fmt.Sprintf("/recipients/%v/cards/%v", params.Recipient, id), c.Tok, body, card)
	} else {
		err = errors.New("Invalid card params: either customer or recipient need to be set")
	}

	return card, err
}

// Delete remotes a card.
// For more details see https://stripe.com/docs/api#delete_card.
func Delete(id string, params *CardParams) error {
	return getC().Delete(id, params)
}

func (c Client) Delete(id string, params *CardParams) error {
	if len(params.Customer) > 0 {
		return c.B.Call("DELETE", fmt.Sprintf("/customers/%v/cards/%v", params.Customer, id), c.Tok, nil, nil)
	} else if len(params.Recipient) > 0 {
		return c.B.Call("DELETE", fmt.Sprintf("/recipients/%v/cards/%v", params.Recipient, id), c.Tok, nil, nil)
	} else {
		return errors.New("Invalid card params: either customer or recipient need to be set")
	}
}

// List returns a list of cards.
// For more details see https://stripe.com/docs/api#list_cards.
func List(params *CardListParams) (*CardList, error) {
	return getC().List(params)
}

func (c Client) List(params *CardListParams) (*CardList, error) {
	body := &url.Values{}

	params.AppendTo(body)

	list := &CardList{}
	var err error

	if len(params.Customer) > 0 {
		err = c.B.Call("GET", fmt.Sprintf("/customers/%v/cards", params.Customer), c.Tok, body, list)
	} else if len(params.Recipient) > 0 {
		err = c.B.Call("GET", fmt.Sprintf("/recipients/%v/cards", params.Recipient), c.Tok, body, list)
	} else {
		err = errors.New("Invalid card params: either customer or recipient need to be set")
	}

	return list, err
}

func getC() Client {
	return Client{GetBackend(), Key}
}
