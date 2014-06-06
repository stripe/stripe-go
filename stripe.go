package stripe

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Api interface {
	Call(method, path, token string, body *url.Values) ([]byte, error)
}

type Client struct {
	Token   string
	api     Api
	Charges *ChargeClient
}

type s struct{}

const uri = "https://api.stripe.com/v1"

var httpClient = &http.Client{}

func (c *Client) Init(token string, api Api) {
	if api == nil {
		api = &s{}
	}

	c.api = api
	c.Token = token

	c.Charges = &ChargeClient{api: c.api, token: c.Token}
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

	res, err := httpClient.Do(req)
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

	//log.Printf("Response %q", ret)

	return ret, nil
}
