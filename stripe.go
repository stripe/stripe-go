// Package stripe provides a client for Stripe REST APIs.
package stripe

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// url is the public Stripe URL for APIs.
const defaultUrl = "https://api.stripe.com/v1"

// apiversion is the currently supported API version
const apiversion = "2014-08-20"

// clientversion is the binding version
const clientversion = "1.0"

// Backend is an interface for making calls against a Stripe service.
// This interface exists to enable mocking for during testing if needed.
type Backend interface {
	Call(method, path, token string, body *url.Values, v interface{}) error
}

// InternalBackend is the internal implementation for making HTTP calls to Stripe.
type InternalBackend struct {
	url        string
	httpClient *http.Client
}

func NewInternalBackend(httpClient *http.Client, url string) *InternalBackend {
	if url == "" {
		url = defaultUrl
	}

	return &InternalBackend{
		url:        url,
		httpClient: GetHttpClient(),
	}
}

// Key is the Stripe API key used globally in the binding.
var Key string

var debug bool
var backend Backend
var httpClient *http.Client

// SetDebug enables additional tracing globally.
// The method is designed for used during testing.
func SetDebug(value bool) {
	debug = value
}

// GetBackend returns the currently used backend in the binding.
func GetBackend() Backend {
	if backend == nil {
		backend = NewInternalBackend(GetHttpClient(), "")
	}

	return backend
}

// SetBackend sets the backend used in the binding.
func SetBackend(b Backend) {
	backend = b
}

// SetHttpClient sets the HTTP client used in the binding.
func SetHttpClient(c *http.Client) {
	httpClient = c
}

// GetHttpClient returns the HTTP client used in the binding.
func GetHttpClient() *http.Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return httpClient
}

// Call is the Backend.Call implementation for invoking Stripe APIs.
func (s *InternalBackend) Call(method, path, token string, body *url.Values, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = s.url + path

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
