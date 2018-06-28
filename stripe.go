// Package stripe provides the binding for Stripe REST APIs.
package stripe

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/stripe/stripe-go/form"
)

//
// Public constants
//

const (
	// APIBackend is a constant representing the API service backend.
	APIBackend SupportedBackend = "api"

	// APIURL is the URL of the API service backend.
	APIURL string = "https://api.stripe.com/v1"

	// UnknownPlatform is the string returned as the system name if we couldn't get
	// one from `uname`.
	UnknownPlatform string = "unknown platform"

	// UploadsBackend is a constant representing the uploads service backend.
	UploadsBackend SupportedBackend = "uploads"

	// UploadsURL is the URL of the uploads service backend.
	UploadsURL string = "https://uploads.stripe.com/v1"
)

//
// Public variables
//

// Key is the Stripe API key used globally in the binding.
var Key string

// LogLevel is the logging level for this library.
// 0: no logging
// 1: errors only
// 2: errors + informational (default)
// 3: errors + informational + debug
var LogLevel = 2

// Logger controls how stripe performs logging at a package level. It is useful
// to customise if you need it prefixed for your application to meet other
// requirements
var Logger Printfer

//
// Public types
//

// AppInfo contains information about the "app" which this integration belongs
// to. This should be reserved for plugins that wish to identify themselves
// with Stripe.
type AppInfo struct {
	Name      string `json:"name"`
	PartnerID string `json:"partner_id"`
	URL       string `json:"url"`
	Version   string `json:"version"`
}

// formatUserAgent formats an AppInfo in a way that's suitable to be appended
// to a User-Agent string. Note that this format is shared between all
// libraries so if it's changed, it should be changed everywhere.
func (a *AppInfo) formatUserAgent() string {
	str := a.Name
	if a.Version != "" {
		str += "/" + a.Version
	}
	if a.URL != "" {
		str += " (" + a.URL + ")"
	}
	return str
}

// Backend is an interface for making calls against a Stripe service.
// This interface exists to enable mocking for during testing if needed.
type Backend interface {
	Call(method, path, key string, params ParamsContainer, v interface{}) error
	CallRaw(method, path, key string, body *form.Values, params *Params, v interface{}) error
	CallMultipart(method, path, key, boundary string, body io.Reader, params *Params, v interface{}) error
	SetMaxNetworkRetries(maxNetworkRetries int)
}

// BackendConfiguration is the internal implementation for making HTTP calls to Stripe.
type BackendConfiguration struct {
	Type              SupportedBackend
	URL               string
	HTTPClient        *http.Client
	MaxNetworkRetries int
}

// Call is the Backend.Call implementation for invoking Stripe APIs.
func (s *BackendConfiguration) Call(method, path, key string, params ParamsContainer, v interface{}) error {
	var body *form.Values
	var commonParams *Params

	if params != nil {
		// This is a little unfortunate, but Go makes it impossible to compare
		// an interface value to nil without the use of the reflect package and
		// its true disciples insist that this is a feature and not a bug.
		//
		// Here we do invoke reflect because (1) we have to reflect anyway to
		// use encode with the form package, and (2) the corresponding removal
		// of boilerplate that this enables makes the small performance penalty
		// worth it.
		reflectValue := reflect.ValueOf(params)

		if reflectValue.Kind() == reflect.Ptr && !reflectValue.IsNil() {
			commonParams = params.GetParams()
			body = &form.Values{}
			form.AppendTo(body, params)
		}
	}

	return s.CallRaw(method, path, key, body, commonParams, v)
}

// CallMultipart is the Backend.CallMultipart implementation for invoking Stripe APIs.
func (s *BackendConfiguration) CallMultipart(method, path, key, boundary string, body io.Reader, params *Params, v interface{}) error {
	contentType := "multipart/form-data; boundary=" + boundary

	req, err := s.NewRequest(method, path, key, contentType, body, params)
	if err != nil {
		return err
	}

	if err := s.Do(req, v); err != nil {
		return err
	}

	return nil
}

func (s *BackendConfiguration) CallRaw(method, path, key string, form *form.Values, params *Params, v interface{}) error {
	var body io.Reader
	if form != nil && !form.Empty() {
		data := form.Encode()
		if method == http.MethodGet {
			path += "?" + data
		} else {
			body = bytes.NewBufferString(data)
		}
	}

	req, err := s.NewRequest(method, path, key, "application/x-www-form-urlencoded", body, params)
	if err != nil {
		return err
	}

	if err := s.Do(req, v); err != nil {
		return err
	}

	return nil
}

// NewRequest is used by Call to generate an http.Request. It handles encoding
// parameters and attaching the appropriate headers.
func (s *BackendConfiguration) NewRequest(method, path, key, contentType string, body io.Reader, params *Params) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = s.URL + path

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		if LogLevel > 0 {
			Logger.Printf("Cannot create Stripe request: %v\n", err)
		}
		return nil, err
	}

	authorization := "Bearer " + key

	req.Header.Add("Authorization", authorization)
	req.Header.Add("Stripe-Version", apiversion)
	req.Header.Add("User-Agent", encodedUserAgent)
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("X-Stripe-Client-User-Agent", encodedStripeUserAgent)

	if params != nil {
		if params.Context != nil {
			req = req.WithContext(params.Context)
		}

		if params.IdempotencyKey != nil {
			idempotencyKey := strings.TrimSpace(*params.IdempotencyKey)
			if len(idempotencyKey) > 255 {
				return nil, errors.New("Cannot use an idempotency key longer than 255 characters.")
			}

			req.Header.Add("Idempotency-Key", idempotencyKey)
		} else if isHTTPWriteMethod(method) {
			req.Header.Add("Idempotency-Key", NewIdempotencyKey())
		}

		if params.StripeAccount != nil {
			req.Header.Add("Stripe-Account", strings.TrimSpace(*params.StripeAccount))
		}

		for k, v := range params.Headers {
			for _, line := range v {
				req.Header.Add(k, line)
			}
		}
	}

	return req, nil
}

// Do is used by Call to execute an API request and parse the response. It uses
// the backend's HTTP client to execute the request and unmarshals the response
// into v. It also handles unmarshaling errors returned by the API.
func (s *BackendConfiguration) Do(req *http.Request, v interface{}) error {
	if LogLevel > 1 {
		Logger.Printf("Requesting %v %v%v\n", req.Method, req.URL.Host, req.URL.Path)
	}

	start := time.Now()

	var res *http.Response
	var err error
	for retry := 0; ; {
		res, err = s.HTTPClient.Do(req)
		if s.shouldRetry(err, res, retry) {
			if LogLevel > 0 {
				Logger.Printf("Request failed with error: %v. Response: %v\n", err, res)
			}
			sleep := sleepTime(retry)
			time.Sleep(sleep)
			retry++
			if LogLevel > 1 {
				Logger.Printf("Retry request %v %v time.\n", req.URL, retry)
			}
		} else {
			break
		}
	}

	if LogLevel > 2 {
		Logger.Printf("Completed in %v\n", time.Since(start))
	}

	if err != nil {
		if LogLevel > 0 {
			Logger.Printf("Request to Stripe failed: %v\n", err)
		}
		return err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		if LogLevel > 0 {
			Logger.Printf("Cannot parse Stripe response: %v\n", err)
		}
		return err
	}

	if res.StatusCode >= 400 {
		return s.ResponseToError(res, resBody)
	}

	if LogLevel > 2 {
		Logger.Printf("Stripe Response: %q\n", resBody)
	}

	if v != nil {
		return json.Unmarshal(resBody, v)
	}

	return nil
}

func (s *BackendConfiguration) ResponseToError(res *http.Response, resBody []byte) error {
	// for some odd reason, the Erro structure doesn't unmarshal
	// initially I thought it was because it's a struct inside of a struct
	// but even after trying that, it still didn't work
	// so unmarshalling to a map for now and parsing the results manually
	// but should investigate later
	var errMap map[string]interface{}

	err := json.Unmarshal(resBody, &errMap)
	if err != nil {
		return err
	}

	e, ok := errMap["error"]
	if !ok {
		err := errors.New(string(resBody))
		if LogLevel > 0 {
			Logger.Printf("Unparsable error returned from Stripe: %v\n", err)
		}
		return err
	}

	root := e.(map[string]interface{})

	stripeErr := &Error{
		Type:           ErrorType(root["type"].(string)),
		Msg:            root["message"].(string),
		HTTPStatusCode: res.StatusCode,
		RequestID:      res.Header.Get("Request-Id"),
	}

	if code, ok := root["code"]; ok {
		stripeErr.Code = ErrorCode(code.(string))
	}

	if param, ok := root["param"]; ok {
		stripeErr.Param = param.(string)
	}

	if charge, ok := root["charge"]; ok {
		stripeErr.ChargeID = charge.(string)
	}

	switch stripeErr.Type {
	case ErrorTypeAPI:
		stripeErr.Err = &APIError{stripeErr: stripeErr}

	case ErrorTypeAPIConnection:
		stripeErr.Err = &APIConnectionError{stripeErr: stripeErr}

	case ErrorTypeAuthentication:
		stripeErr.Err = &AuthenticationError{stripeErr: stripeErr}

	case ErrorTypeCard:
		cardErr := &CardError{stripeErr: stripeErr}
		stripeErr.Err = cardErr

		if declineCode, ok := root["decline_code"]; ok {
			cardErr.DeclineCode = declineCode.(string)
		}

	case ErrorTypeInvalidRequest:
		stripeErr.Err = &InvalidRequestError{stripeErr: stripeErr}

	case ErrorTypePermission:
		stripeErr.Err = &PermissionError{stripeErr: stripeErr}

	case ErrorTypeRateLimit:
		stripeErr.Err = &RateLimitError{stripeErr: stripeErr}
	}

	if LogLevel > 0 {
		Logger.Printf("Error encountered from Stripe: %v\n", stripeErr)
	}

	return stripeErr
}

// SetMaxNetworkRetries sets max number of retries on failed requests
func (s *BackendConfiguration) SetMaxNetworkRetries(maxNetworkRetries int) {
	s.MaxNetworkRetries = maxNetworkRetries
}

// Checks if an error is a problem that we should retry on. This includes both
// socket errors that may represent an intermittent problem and some special
// HTTP statuses.
func (s *BackendConfiguration) shouldRetry(err error, resp *http.Response, numRetries int) bool {
	if numRetries >= s.MaxNetworkRetries {
		return false
	}

	if err != nil {
		return true
	}

	if resp.StatusCode == http.StatusConflict {
		return true
	}
	return false
}

// Backends are the currently supported endpoints.
type Backends struct {
	API, Uploads Backend
	mu           sync.RWMutex
}

// Printfer is an interface to be implemented by Logger.
type Printfer interface {
	Printf(format string, v ...interface{})
}

// SupportedBackend is an enumeration of supported Stripe endpoints.
// Currently supported values are "api" and "uploads".
type SupportedBackend string

//
// Public functions
//

// Bool returns a pointer to the bool value passed in.
func Bool(v bool) *bool {
	return &v
}

// BoolValue returns the value of the bool pointer passed in or
// false if the pointer is nil.
func BoolValue(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}

// Float64 returns a pointer to the float64 value passed in.
func Float64(v float64) *float64 {
	return &v
}

// Float64Value returns the value of the float64 pointer passed in or
// 0 if the pointer is nil.
func Float64Value(v *float64) float64 {
	if v != nil {
		return *v
	}
	return 0
}

// FormatURLPath takes a format string (of the kind used in the fmt package)
// representing a URL path with a number of parameters that belong in the path
// and returns a formatted string.
//
// This is mostly a pass through to Sprintf. It exists to make it
// it impossible to accidentally provide a parameter type that would be
// formatted improperly; for example, a string pointer instead of a string.
//
// It also URL-escapes every given parameter. This usually isn't necessary for
// a standard Stripe ID, but is needed in places where user-provided IDs are
// allowed, like in coupons or plans. We apply it broadly for extra safety.
func FormatURLPath(format string, params ...string) string {
	// Convert parameters to interface{} and URL-escape them
	untypedParams := make([]interface{}, len(params))
	for i, param := range params {
		untypedParams[i] = interface{}(url.QueryEscape(param))
	}

	return fmt.Sprintf(format, untypedParams...)
}

// GetBackend returns the currently used backend in the binding.
func GetBackend(backend SupportedBackend) Backend {
	switch backend {
	case APIBackend:
		backends.mu.RLock()
		ret := backends.API
		backends.mu.RUnlock()
		if ret != nil {
			return ret
		}
		backends.mu.Lock()
		defer backends.mu.Unlock()
		backends.API = &BackendConfiguration{Type: backend, URL: apiURL, HTTPClient: httpClient}
		return backends.API

	case UploadsBackend:
		backends.mu.RLock()
		ret := backends.Uploads
		backends.mu.RUnlock()
		if ret != nil {
			return ret
		}
		backends.mu.Lock()
		defer backends.mu.Unlock()
		backends.Uploads = &BackendConfiguration{Type: backend, URL: uploadsURL, HTTPClient: httpClient}
		return backends.Uploads
	}

	return nil
}

// Int64 returns a pointer to the int64 value passed in.
func Int64(v int64) *int64 {
	return &v
}

// Int64Value returns the value of the int64 pointer passed in or
// 0 if the pointer is nil.
func Int64Value(v *int64) int64 {
	if v != nil {
		return *v
	}
	return 0
}

// NewBackends creates a new set of backends with the given HTTP client. You
// should only need to use this for testing purposes or on App Engine.
func NewBackends(httpClient *http.Client) *Backends {
	return &Backends{
		API: &BackendConfiguration{
			Type: APIBackend, URL: APIURL, HTTPClient: httpClient},
		Uploads: &BackendConfiguration{
			Type: UploadsBackend, URL: UploadsURL, HTTPClient: httpClient},
	}
}

// ParseID attempts to parse a string scalar from a given JSON value which is
// still encoded as []byte. If the value was a string, it returns the string
// along with true as the second return value. If not, false is returned as the
// second return value.
//
// The purpose of this function is to detect whether a given value in a
// response from the Stripe API is a string ID or an expanded object.
func ParseID(data []byte) (string, bool) {
	s := string(data)

	if !strings.HasPrefix(s, "\"") {
		return "", false
	}

	if !strings.HasSuffix(s, "\"") {
		return "", false
	}

	return s[1 : len(s)-1], true
}

// SetAppInfo sets app information. See AppInfo.
func SetAppInfo(info *AppInfo) {
	if info != nil && info.Name == "" {
		panic(fmt.Errorf("App info name cannot be empty"))
	}
	appInfo = info

	// This is run in init, but we need to reinitialize it now that we have
	// some app info.
	initUserAgent()
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

// SetHTTPClient overrides the default HTTP client.
// This is useful if you're running in a Google AppEngine environment
// where the http.DefaultClient is not available.
func SetHTTPClient(client *http.Client) {
	httpClient = client
}

// String returns a pointer to the string value passed in.
func String(v string) *string {
	return &v
}

// StringValue returns the value of the string pointer passed in or
// "" if the pointer is nil.
func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

//
// Private constants
//

const apiURL = "https://api.stripe.com/v1"

// apiversion is the currently supported API version
const apiversion = "2018-02-06"

// clientversion is the binding version
const clientversion = "35.7.0"

// defaultHTTPTimeout is the default timeout on the http.Client used by the library.
// This is chosen to be consistent with the other Stripe language libraries and
// to coordinate with other timeouts configured in the Stripe infrastructure.
const defaultHTTPTimeout = 80 * time.Second

// maxNetworkRetriesDelay and minNetworkRetriesDelay defines sleep time in milliseconds between
// tries to send HTTP request again after network failure.
const maxNetworkRetriesDelay = 5000 * time.Millisecond
const minNetworkRetriesDelay = 500 * time.Millisecond

const uploadsURL = "https://uploads.stripe.com/v1"

//
// Private types
//

// stripeClientUserAgent contains information about the current runtime which
// is serialized and sent in the `X-Stripe-Client-User-Agent` as additional
// debugging information.
type stripeClientUserAgent struct {
	Application     *AppInfo `json:"application"`
	BindingsVersion string   `json:"bindings_version"`
	Language        string   `json:"lang"`
	LanguageVersion string   `json:"lang_version"`
	Publisher       string   `json:"publisher"`
	Uname           string   `json:"uname"`
}

//
// Private variables
//

var appInfo *AppInfo
var backends Backends
var encodedStripeUserAgent string
var encodedUserAgent string
var httpClient = &http.Client{Timeout: defaultHTTPTimeout}

//
// Private functions
//

// getUname tries to get a uname from the system, but not that hard. It tries
// to execute `uname -a`, but swallows any errors in case that didn't work
// (i.e. non-Unix non-Mac system or some other reason).
func getUname() string {
	path, err := exec.LookPath("uname")
	if err != nil {
		return UnknownPlatform
	}

	cmd := exec.Command(path, "-a")
	var out bytes.Buffer
	cmd.Stderr = nil // goes to os.DevNull
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return UnknownPlatform
	}

	return out.String()
}

func init() {
	Logger = log.New(os.Stderr, "", log.LstdFlags)
	initUserAgent()
}

func initUserAgent() {
	encodedUserAgent = "Stripe/v1 GoBindings/" + clientversion
	if appInfo != nil {
		encodedUserAgent += " " + appInfo.formatUserAgent()
	}

	stripeUserAgent := &stripeClientUserAgent{
		Application:     appInfo,
		BindingsVersion: clientversion,
		Language:        "go",
		LanguageVersion: runtime.Version(),
		Publisher:       "stripe",
		Uname:           getUname(),
	}
	marshaled, err := json.Marshal(stripeUserAgent)
	// Encoding this struct should never be a problem, so we're okay to panic
	// in case it is for some reason.
	if err != nil {
		panic(err)
	}
	encodedStripeUserAgent = string(marshaled)
}

func isHTTPWriteMethod(method string) bool {
	return method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch || method == http.MethodDelete
}

// sleepTime calculates sleeping/delay time in milliseconds between failure and a new one request.
func sleepTime(numRetries int) time.Duration {
	// Apply exponential backoff with minNetworkRetriesDelay on the
	// number of num_retries so far as inputs.
	delay := minNetworkRetriesDelay + minNetworkRetriesDelay*time.Duration(numRetries*numRetries)

	// Do not allow the number to exceed maxNetworkRetriesDelay.
	if delay > maxNetworkRetriesDelay {
		delay = maxNetworkRetriesDelay
	}

	// Apply some jitter by randomizing the value in the range of 75%-100%.
	jitter := rand.Int63n(int64(delay / 4))
	delay -= time.Duration(jitter)

	// But never sleep less than the base sleep seconds.
	if delay < minNetworkRetriesDelay {
		delay = minNetworkRetriesDelay
	}

	return delay
}
