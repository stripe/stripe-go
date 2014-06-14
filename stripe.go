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

var debug bool

type Api interface {
	Call(method, path, token string, body *url.Values, v interface{}) error
}

type Client struct {
	Token        string
	api          Api
	Charges      *ChargeClient
	Customers    *CustomerClient
	Cards        *CardClient
	Subs         *SubscriptionClient
	Plans        *PlanClient
	Coupons      *CouponClient
	Discounts    *DiscountClient
	Invoices     *InvoiceClient
	InvoiceItems *InvoiceItemClient
	Disputes     *DisputeClient
	Transfers    *TransferClient
	Recipients   *RecipientClient
	Fees         *AppFeeClient
	Account      *AccountClient
	Balance      *BalanceClient
	Tokens       *TokenClient
}

type s struct {
	httpClient *http.Client
}

type Filters struct {
	f []*filter
}

type filter struct {
	Key, Op, Val string
}

const uri = "https://api.stripe.com/v1"

func (c *Client) Init(token string, client *http.Client, api Api) {
	if client == nil {
		client = http.DefaultClient
	}

	if api == nil {
		api = &s{httpClient: client}
	}
	c.api = api

	c.Token = token

	c.Charges = &ChargeClient{api: c.api, token: c.Token}
	c.Customers = &CustomerClient{api: c.api, token: c.Token}
	c.Cards = &CardClient{api: c.api, token: c.Token}
	c.Subs = &SubscriptionClient{api: c.api, token: c.Token}
	c.Plans = &PlanClient{api: c.api, token: c.Token}
	c.Coupons = &CouponClient{api: c.api, token: c.Token}
	c.Discounts = &DiscountClient{api: c.api, token: c.Token}
	c.Invoices = &InvoiceClient{api: c.api, token: c.Token}
	c.InvoiceItems = &InvoiceItemClient{api: c.api, token: c.Token}
	c.Disputes = &DisputeClient{api: c.api, token: c.Token}
	c.Transfers = &TransferClient{api: c.api, token: c.Token}
	c.Recipients = &RecipientClient{api: c.api, token: c.Token}
	c.Fees = &AppFeeClient{api: c.api, token: c.Token}
	c.Account = &AccountClient{api: c.api, token: c.Token}
	c.Balance = &BalanceClient{api: c.api, token: c.Token}
	c.Tokens = &TokenClient{api: c.api, token: c.Token}
}

func (c *Client) SetDebug(value bool) {
	debug = value
}

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

func (f *Filters) AddFilter(key, op, value string) {
	filter := &filter{Key: key, Op: op, Val: value}
	f.f = append(f.f, filter)
}

func (f *Filters) appendTo(values *url.Values) {
	for _, v := range f.f {
		values.Add(fmt.Sprintf("%v[%v]", v.Key, v.Op), v.Val)
	}
}
