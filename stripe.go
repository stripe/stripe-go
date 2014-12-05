// Package stripe provides the binding for Stripe REST APIs.
package stripe

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// defaultURL is the public Stripe URL for APIs.
const defaultURL = "https://api.stripe.com/v1"

// apiversion is the currently supported API version
const apiversion = "2014-12-08"

// clientversion is the binding version
const clientversion = "3.0.0"

// Backend is an interface for making calls against a Stripe service.
// This interface exists to enable mocking for during testing if needed.
type Backend interface {
	Call(method, path, key string, body *url.Values, v interface{}) error
}

// InternalBackend is the internal implementation for making HTTP calls to Stripe.
type InternalBackend struct {
	url        string
	httpClient *http.Client
}

// NewInternalBackend returns a customized backend used for making calls in this binding.
// This method should be called in one of two scenarios:
//	1. You're running in a Google AppEngine environment where the http.DefaultClient is not available.
//  2. You're doing internal development at Stripe.
func NewInternalBackend(httpClient *http.Client, url string) *InternalBackend {
	if len(url) == 0 {
		url = defaultURL
	}

	return &InternalBackend{
		url:        url,
		httpClient: httpClient,
	}
}

// Key is the Stripe API key used globally in the binding.
var Key string

var debug bool
var backend Backend

// SetDebug enables additional tracing globally.
// The method is designed for used during testing.
func SetDebug(value bool) {
	debug = value
}

// GetBackend returns the currently used backend in the binding.
func GetBackend() Backend {
	if backend == nil {
		backend = NewInternalBackend(http.DefaultClient, "")
	}

	return backend
}

// SetBackend sets the backend used in the binding.
func SetBackend(b Backend) {
	backend = b
}

// Call is the Backend.Call implementation for invoking Stripe APIs.
func (s *InternalBackend) Call(method, path, key string, form *url.Values, v interface{}) error {
	req, err := s.NewRequest(method, path, key, form)
	if err != nil {
		return err
	}

	if err := s.Do(req, v); err != nil {
		return err
	}

	return nil
}

// NewRequest is used by Call to generate an http.Request. It handles encoding
// parameters and attaching the Authorization header.
func (s *InternalBackend) NewRequest(method, path, key string, form *url.Values) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = s.url + path

	var body io.Reader
	if form != nil && len(*form) > 0 {
		data := form.Encode()
		if strings.ToUpper(method) == "GET" {
			path += "?" + data
		} else {
			body = bytes.NewBufferString(data)
		}
	}

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		log.Printf("Cannot create Stripe request: %v\n", err)
		return nil, err
	}

	req.SetBasicAuth(key, "")
	req.Header.Add("Stripe-Version", apiversion)
	req.Header.Add("User-Agent", "Stripe/v1 GoBindings/"+clientversion)
	if body != nil {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	return req, nil
}

// Do is used by Call to execute an API request and parse the response. It uses
// the backend's HTTP client to execute the request and unmarshals the response
// into v. It also handles unmarshaling errors returned by the API.
func (s *InternalBackend) Do(req *http.Request, v interface{}) error {
	log.Printf("Requesting %v %q\n", req.Method, req.URL.Path)
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
				Type:           ErrorType(root["type"].(string)),
				Msg:            root["message"].(string),
				HTTPStatusCode: res.StatusCode,
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
