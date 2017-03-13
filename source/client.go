package source

import (
	"fmt"
	"strconv"

	stripe "github.com/stripe/stripe-go"
)

// Client is used to invoke /sources APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs a new source.
// For more details see https://stripe.com/docs/api#create_source.
func New(params *stripe.SourceObjectParams) (*stripe.Source, error) {
	return getC().New(params)
}

// New POSTs a new source.
// For more details see https://stripe.com/docs/api#create_source.
func (c Client) New(params *stripe.SourceObjectParams) (*stripe.Source, error) {
	var body *stripe.RequestValues
	var commonParams *stripe.Params

	if params != nil {
		body = &stripe.RequestValues{}
		commonParams = &params.Params

		// Required fields
		body.Add("type", params.Type)

		// Optional fields
		if params.Currency != "" {
			body.Add("currency", string(params.Currency))
		}
		if params.Amount > 0 {
			body.Add("amount", strconv.FormatUint(params.Amount, 10))
		}
		if params.Flow != "" {
			body.Add("flow", string(params.Flow))
		}
		if params.Owner != nil {
			if params.Owner.Address != nil {
				if params.Owner.Address.Line1 != "" {
					body.Add("owner[address][line1]", params.Owner.Address.Line1)
				}
				if params.Owner.Address.Line2 != "" {
					body.Add("owner[address][line2]", params.Owner.Address.Line2)
				}
				if params.Owner.Address.City != "" {
					body.Add("owner[address][city]", params.Owner.Address.City)
				}
				if params.Owner.Address.State != "" {
					body.Add("owner[address][state]", params.Owner.Address.State)
				}
				if params.Owner.Address.PostalCode != "" {
					body.Add("owner[address][postal_code]", params.Owner.Address.PostalCode)
				}
				if params.Owner.Address.Country != "" {
					body.Add("owner[address][country]", params.Owner.Address.Country)
				}
			}
			if params.Owner.Email != "" {
				body.Add("owner[email]", params.Owner.Email)
			}
			if params.Owner.Name != "" {
				body.Add("owner[name]", params.Owner.Name)
			}
			if params.Owner.Phone != "" {
				body.Add("owner[phone]", params.Owner.Phone)
			}
		}

		if params.Redirect != nil {
			if params.Redirect.ReturnURL != "" {
				body.Add("redirect[return_url]", params.Redirect.ReturnURL)
			}
		}

		if params.Token != "" {
			body.Add("token", params.Token)
		}

		for k, v := range params.TypeData {
			body.Add(fmt.Sprintf("%s[%s]", params.Type, k), v)
		}

		params.AppendTo(body)
	}

	p := &stripe.Source{}
	err := c.B.Call("POST", "/sources", c.Key, body, commonParams, p)

	return p, err
}

// Get returns the details of a source
// For more details see https://stripe.com/docs/api#retrieve_source.
func Get(id string, params *stripe.SourceObjectParams) (*stripe.Source, error) {
	return getC().Get(id, params)
}

// Get returns the details of a source
// For more details see https://stripe.com/docs/api#retrieve_source.
func (c Client) Get(id string, params *stripe.SourceObjectParams) (*stripe.Source, error) {
	var body *stripe.RequestValues
	var commonParams *stripe.Params

	if params != nil {
		body = &stripe.RequestValues{}
		commonParams = &params.Params
		params.AppendTo(body)
	}

	source := &stripe.Source{}
	err := c.B.Call("GET", "/sources/"+id, c.Key, body, commonParams, source)
	return source, err
}

// Update updates a source's properties.
// For more details see	https://stripe.com/docs/api#update_source.
func Update(id string, params *stripe.SourceObjectParams) (*stripe.Source, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.SourceObjectParams) (*stripe.Source, error) {
	var body *stripe.RequestValues
	var commonParams *stripe.Params

	if params != nil {
		body = &stripe.RequestValues{}
		commonParams = &params.Params

		// Optional fields
		if params.Owner != nil {
			if params.Owner.Address != nil {
				if params.Owner.Address.Line1 != "" {
					body.Add("owner[address][line1]", params.Owner.Address.Line1)
				}
				if params.Owner.Address.Line2 != "" {
					body.Add("owner[address][line2]", params.Owner.Address.Line2)
				}
				if params.Owner.Address.City != "" {
					body.Add("owner[address][city]", params.Owner.Address.City)
				}
				if params.Owner.Address.State != "" {
					body.Add("owner[address][state]", params.Owner.Address.State)
				}
				if params.Owner.Address.PostalCode != "" {
					body.Add("owner[address][postal_code]", params.Owner.Address.PostalCode)
				}
				if params.Owner.Address.Country != "" {
					body.Add("owner[address][country]", params.Owner.Address.Country)
				}
			}
			if params.Owner.Email != "" {
				body.Add("owner[email]", params.Owner.Email)
			}
			if params.Owner.Name != "" {
				body.Add("owner[name]", params.Owner.Name)
			}
			if params.Owner.Phone != "" {
				body.Add("owner[phone]", params.Owner.Phone)
			}
		}

		params.AppendTo(body)
	}

	source := &stripe.Source{}
	err := c.B.Call("POST", "/sources/"+id, c.Key, body, commonParams, source)

	return source, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
