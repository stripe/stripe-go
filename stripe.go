// Package stripe provides a client for Stripe REST APIs.
package stripe

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// uri is the public Stripe URL for APIs.
const uri = "https://api.stripe.com/v1"

// apiversion is the currently supported API version
const apiversion = "2014-08-20"

// clientversion is the binding version
const clientversion = "1.0"

// debug is a global variable that enables additional tracing
// to help with troubleshooting while testing.
var debug bool

// Api is an interface for making calls against a Stripe service.
// This interface exists to enable mocking for during testing if needed.
type Api interface {
	Call(method, path, token string, body *url.Values, v interface{}) error
}

// Client is the Stripe client. It contains all the different resources available.
type Client struct {
	// Token is the secret key used for authentication.
	Token string
	// api is the Api implementation used to invoke Stripe APIs.
	api Api
	// Charges is the client used to invoke /charges APIs.
	// For more details see https://stripe.com/docs/api/#charges.
	Charges *chargeClient
	// Customers is the client used to invoke /customers APIs.
	// For more details see https://stripe.com/docs/api/#customers.
	Customers *customerClient
	// Cards is the client used to invoke /cards APIs.
	// For more details see https://stripe.com/docs/api/#cards.
	Cards *cardClient
	// Subs is the client used to invoke /subscriptions APIs.
	// For more details see https://stripe.com/docs/api/#subscriptions.
	Subs *subscriptionClient
	// Plans is the client used to invoke /plans APIs.
	// For more details see https://stripe.com/docs/api/#plans.
	Plans *planClient
	// Coupons is the client used to invoke /coupons APIs.
	// For more details see https://stripe.com/docs/api/#coupons.
	Coupons *couponClient
	// Discounts is the client used to invoke discount-related APIs.
	// For mode details see https://stripe.com/docs/api/#discounts.
	Discounts *discountClient
	// Invoices is the client used to invoke /invoices APIs.
	// For more details see https://stripe.com/docs/api/#invoices.
	Invoices *invoiceClient
	// InvoiceItems is the client used to invoke /invoiceitems APIs.
	// For more details see https://stripe.com/docs/api/#invoiceitems.
	InvoiceItems *invoiceItemClient
	// Disputes is the client used to invoke dispute-related APIs.
	// For more details see https://stripe.com/docs/api/#disputes.
	Disputes *disputeClient
	// Transfers is the client used to invoke /transfers APIs.
	// For more details see https://stripe.com/docs/api/#transfers.
	Transfers *transferClient
	// Recipients is the client used to invoke /recipients APIs.
	// For more details see https://stripe.com/docs/api/#recipients.
	Recipients *recipientClient
	// Refunds is the client used to invoke /refunds APIs.
	// For more details see https://stripe.com/docs/api#refunds.
	Refunds *refundClient
	// Fees is the client used to invoke /application_fees APIs.
	// For more details see https://stripe.com/docs/api/#application_fees.
	Fees *appFeeClient
	// FeeRefunds is the client used to invoke /application_fees/refunds APIs.
	// For more details see https://stripe.com/docs/api/#fee_refundss.
	FeeRefunds *feeRefundClient
	// Account is the client used to invoke /account APIs.
	// For more details see https://stripe.com/docs/api/#account.
	Account *accountClient
	// Balance is the client used to invoke /balance and transaction-related APIs.
	// For more details see https://stripe.com/docs/api/#balance.
	Balance *balanceClient
	// Events is the client used to invoke /events APIs.
	// For more details see https://stripe.com/docs/api#events.
	Events *eventClient
	// Tokens is the client used to invoke /tokens APIs.
	// For more details see https://stripe.com/docs/api/#tokens.
	Tokens *tokenClient
}

// s is the internal implementation for making HTTP calls to Stripe.
type s struct {
	httpClient *http.Client
}

// Filters is a structure that contains a collection of filters for list-related APIs.
type Filters struct {
	f []*filter
}

// filter is the structure that contains a filter for list-related APIs.
// It ends up passing query string parameters in the format key[op]=value.
type filter struct {
	Key, Op, Val string
}

// ListResponse is the structure that contains the common properties
// of any *List structure. The Count property is only populated if the
// total_count include option is passed in (see tests for example).
type ListResponse struct {
	Count uint16 `json:"total_count"`
	More  bool   `json:"has_more"`
	Url   string `json:"url"`
}

// Init initializes the Stripe client with the appropriate token secret key
// as well as providing the ability to override the HTTP client and api used.
func (c *Client) Init(token string, client *http.Client, api Api) {
	if client == nil {
		client = http.DefaultClient
	}

	if api == nil {
		api = &s{httpClient: client}
	}
	c.api = api
	c.Token = token

	c.Charges = &chargeClient{api: c.api, token: c.Token}
	c.Customers = &customerClient{api: c.api, token: c.Token}
	c.Cards = &cardClient{api: c.api, token: c.Token}
	c.Subs = &subscriptionClient{api: c.api, token: c.Token}
	c.Plans = &planClient{api: c.api, token: c.Token}
	c.Coupons = &couponClient{api: c.api, token: c.Token}
	c.Discounts = &discountClient{api: c.api, token: c.Token}
	c.Invoices = &invoiceClient{api: c.api, token: c.Token}
	c.InvoiceItems = &invoiceItemClient{api: c.api, token: c.Token}
	c.Disputes = &disputeClient{api: c.api, token: c.Token}
	c.Transfers = &transferClient{api: c.api, token: c.Token}
	c.Recipients = &recipientClient{api: c.api, token: c.Token}
	c.Refunds = &refundClient{api: c.api, token: c.Token}
	c.Fees = &appFeeClient{api: c.api, token: c.Token}
	c.FeeRefunds = &feeRefundClient{api: c.api, token: c.Token}
	c.Account = &accountClient{api: c.api, token: c.Token}
	c.Balance = &balanceClient{api: c.api, token: c.Token}
	c.Events = &eventClient{api: c.api, token: c.Token}
	c.Tokens = &tokenClient{api: c.api, token: c.Token}
}

// SetDebug enables additional tracing globally.
// The method is designed for used during testing.
func (c *Client) SetDebug(value bool) {
	debug = value
}

// Call is the Api.Call implementation for invoking Stripe APIs.
func (s *s) Call(method, path, token string, body *url.Values, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = uri + path

	if body != nil && len(*body) > 0 {
		path += "?" + body.Encode()
	}

	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		log.Printf("Cannot create Stripe request: %v\n", err)
		return err
	}

	req.SetBasicAuth(token, "")
	req.Header.Add("Stripe-Version", apiversion)
	req.Header.Add("X-Stripe-Client-User-Agent", "Stripe.Go-"+clientversion)

	log.Printf("Requesting %v %q\n", method, path)
	start := time.Now()

	res, err := s.httpClient.Do(req)

	if debug {
		log.Printf("Completed in %v\n", time.Since(start))
	}

	if err != nil {
		log.Printf("Request to Stripe failed: %v\n", err)
		return err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Cannot parse Stripe response: %v\n", err)
		return err
	}

	if res.StatusCode >= 400 {
		// for some odd reason, the Erro structure doesn't unmarshal
		// initially I thought it was because it's a struct inside of a struct
		// but even after trying that, it still didn't work
		// so unmarshalling to a map for now and parsing the results manually
		// but should investigate later
		var errMap map[string]interface{}
		json.Unmarshal(resBody, &errMap)

		if e, found := errMap["error"]; !found {
			err := errors.New(string(resBody))
			log.Printf("Unparsable error returned from Stripe: %v\n", err)
			return err
		} else {
			root := e.(map[string]interface{})
			err := &Error{
				Type: ErrorType(root["type"].(string)),
				Msg:  root["message"].(string),
			}

			if code, found := root["code"]; found {
				err.Code = ErrorCode(code.(string))
			}

			if param, found := root["param"]; found {
				err.Param = param.(string)
			}

			log.Printf("Error encountered from Stripe: %v\n", err)
			return err
		}
	}

	if debug {
		log.Printf("Stripe Response: %q\n", resBody)
	}

	if v != nil {
		return json.Unmarshal(resBody, v)
	}

	return nil
}

// AddFilter adds a new filter with a given key, op and value.
func (f *Filters) AddFilter(key, op, value string) {
	filter := &filter{Key: key, Op: op, Val: value}
	f.f = append(f.f, filter)
}

// appendTo adds the list of filters to the query string values.
func (f *Filters) appendTo(values *url.Values) {
	for _, v := range f.f {
		if len(v.Op) > 0 {
			values.Add(fmt.Sprintf("%v[%v]", v.Key, v.Op), v.Val)
		} else {
			values.Add(v.Key, v.Val)
		}
	}
}
