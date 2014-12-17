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

const (
	apiURL     = "https://api.stripe.com/v1"
	uploadsURL = "https://uploads.stripe.com/v1"
)

// apiversion is the currently supported API version
const apiversion = "2014-12-08"

// clientversion is the binding version
const clientversion = "4.0.0"

// defaultHTTPTimeout is the default timeout on the http.Client used by the bindings.
// This is chosen to be consistent with the other Stripe language bindings and
// to coordinate with other timeouts configured in the Stripe infrastructure.
const defaultHTTPTimeout = 80 * time.Second

// Totalbackends is the total number of Stripe API endpoints supported by the binding.
const TotalBackends = 2

// Backend is an interface for making calls against a Stripe service.
// This interface exists to enable mocking for during testing if needed.
type Backend interface {
	Call(method, path, key string, body *url.Values, v interface{}) error
	CallMultipart(method, path, key, boundary string, body io.Reader, v interface{}) error
}

// BackendConfiguration is the internal implementation for making HTTP calls to Stripe.
type BackendConfiguration struct {
	Type       SupportedBackend
	URL        string
	HTTPClient *http.Client
}

// SupportedBackend is an enumeration of supported Stripe endpoints.
// Currently supported valuesa are "api" and "uploads".
type SupportedBackend string

const (
	APIBackend     SupportedBackend = "api"
	UploadsBackend SupportedBackend = "uploads"
)

type StripeBackends struct {
	API, Uploads Backend
}

// Key is the Stripe API key used globally in the binding.
var Key string

var debug bool
var httpClient = &http.Client{Timeout: defaultHTTPTimeout}
var backends StripeBackends

// SetHTTPClient overrides the default HTTP client.
// This is useful if you're running in a Google AppEngine environment
// where the http.DefaultClient is not available.
func SetHTTPClient(client *http.Client) {
	httpClient = client
}

// GetBackend returns the currently used backend in the binding.
func GetBackend(backend SupportedBackend) Backend {
	var ret Backend
	switch backend {
	case APIBackend:
		if backends.API == nil {
			backends.API = BackendConfiguration{backend, apiURL, httpClient}
		}

		ret = backends.API
	case UploadsBackend:
		if backends.Uploads == nil {
			backends.Uploads = BackendConfiguration{backend, uploadsURL, httpClient}
		}
		ret = backends.Uploads
	}

	return ret
}

// SetBackend sets the backend used in the binding.
func SetBackend(backend SupportedBackend, b Backend) {
	switch backend {
	case APIBackend:
		backends.API = b
	case UploadsBackend:
		backends.Uploads = b
	}
}

// Call is the Backend.Call implementation for invoking Stripe APIs.
func (s BackendConfiguration) Call(method, path, key string, form *url.Values, v interface{}) error {
	var body io.Reader
	if form != nil && len(*form) > 0 {
		data := form.Encode()
		if strings.ToUpper(method) == "GET" {
			path += "?" + data
		} else {
			body = bytes.NewBufferString(data)
		}
	}

	req, err := s.NewRequest(method, path, key, "application/x-www-form-urlencoded", body)
	if err != nil {
		return err
	}

	if err := s.Do(req, v); err != nil {
		return err
	}

	return nil
}

// CallMultipart is the Backend.CallMultipart implementation for invoking Stripe APIs.
func (s BackendConfiguration) CallMultipart(method, path, key, boudary string, body io.Reader, v interface{}) error {
	contentType := "multipart/form-data; boundary=" + boudary

	req, err := s.NewRequest(method, path, key, contentType, body)
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
func (s *BackendConfiguration) NewRequest(method, path, key, contentType string, body io.Reader) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = s.URL + path

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		log.Printf("Cannot create Stripe request: %v\n", err)
		return nil, err
	}

	req.SetBasicAuth(key, "")
	req.Header.Add("Stripe-Version", apiversion)
	req.Header.Add("User-Agent", "Stripe/v1 GoBindings/"+clientversion)
	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// Do is used by Call to execute an API request and parse the response. It uses
// the backend's HTTP client to execute the request and unmarshals the response
// into v. It also handles unmarshaling errors returned by the API.
func (s *BackendConfiguration) Do(req *http.Request, v interface{}) error {
	log.Printf("Requesting %v %v%v\n", req.Method, req.URL.Host, req.URL.Path)
	start := time.Now()

	res, err := s.HTTPClient.Do(req)

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

// SetDebug enables additional tracing globally.
// The method is designed for used during testing.
func SetDebug(value bool) {
	debug = value
}
