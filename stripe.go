package stripe

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Api interface {
	Call(method, path, token string, body *url.Values) ([]byte, error)
}

type Client struct {
	Token     string
	api       Api
	Charges   *ChargeClient
	Customers *CustomerClient
	Cards     *CardClient
}

type s struct {
	httpClient *http.Client
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
}

func (s *s) Call(method, path, token string, body *url.Values) ([]byte, error) {
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
		return nil, err
	}

	req.SetBasicAuth(token, "")

	log.Printf("Requesting %v %q\n", method, path)
	start := time.Now()

	res, err := s.httpClient.Do(req)

	log.Printf("Completed in %v\n", time.Since(start))

	if err != nil {
		log.Printf("Request to Stripe failed: %v\n", err)
		return nil, err
	}
	defer res.Body.Close()

	ret, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Cannot parse Stripe response: %v\n", err)
		return nil, err
	}

	if res.StatusCode >= 400 {
		err = errors.New(string(ret))
		log.Printf("Error encountered from Stripe: %v\n", err)
		return nil, err
	}

	return ret, nil
}
