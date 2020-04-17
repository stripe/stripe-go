# V71 Migration Guide

Some of the major changes included in version 71 are breaking, but in relatively minor ways. Read on for details.

## Go Modules

The library now fully supports [Go Modules](https://github.com/golang/go/wiki/Modules), Go's preferred packaging system which is built into the language. Requires that previously looked like:

``` go
import (
    "github.com/stripe/stripe-go"
    "github.com/stripe/stripe-go/customer"
)
```

Should now also include a major version:

``` go
import (
    "github.com/stripe/stripe-go/v71"
    "github.com/stripe/stripe-go/v71/customer"
)
```

This convention is recognized by the Go toolchain. After adding a major version to paths, just run `go build`, `go install`, `go test`, etc. and it will automatically fetch stripe-go and do the right thing.

Go Modules are supported on every version of Go supported by Stripe and basic awareness for them has been backported as far as Go 1.9.

Upgrading to stripe-go v71 will be difficult for projects using [Dep](https://github.com/golang/dep). There have been attempts to [contribute minimal Go Module awareness](https://github.com/golang/dep/pull/1963) to the project, but no support has been able to make it in. As Dep is now unmaintained for all intents and purposes, if you're still using it, we recommend moving your project to Go Modules instead. Because Go Modules are Dep-aware (even if the vice versa is not true), this is likely easier than you'd intuitively think because the tooling supports it automatically. See the official [post on migrating to Go Modules](https://blog.golang.org/migrating-to-go-modules) for more details.

## API response information

All API resource structs now have a `LastResponse` that allows access to useful information on the API request that generated them. For example:

``` go
customer, err := customer.New(params)
fmt.Printf("request ID = %s\n", customer.LastResponse.RequestID)
```

Note that although `LastResponse` is available on any struct that might be returned from a API call, it will only be populated for the top-level API resource that was received. e.g. When fetching an invoice from `GET /invoices/:id`, the invoice's `LastResponse` is populated, but that of its invoice items is not. However, when fetching an invoice item from `GET /invoiceitems/:id`, `LastResponse` for that invoice item is populated.

The full struct in `LastResponse` makes these properties available:

``` go
// APIResponse encapsulates some common features of a response from the
// Stripe API.
type APIResponse struct {
	// Header contain a map of all HTTP header keys to values. Its behavior and
	// caveats are identical to that of http.Header.
	Header http.Header

	// IdempotencyKey contains the idempotency key used with this request.
	// Idempotency keys are a Stripe-specific concept that helps guarantee that
	// requests that fail and need to be retried are not duplicated.
	IdempotencyKey string

	// RawJSON contains the response body as raw bytes.
	RawJSON []byte

	// RequestID contains a string that uniquely identifies the Stripe request.
	// Used for debugging or support purposes.
	RequestID string

	// Status is a status code and message. e.g. "200 OK"
	Status string

	// StatusCode is a status code as integer. e.g. 200
	StatusCode int
}
```

## `BackendConfig` now takes pointers

The types for primitives on `BackendConfig` now take pointers:

* `EnableTelemetry` changes from `bool` to `*bool`.
* `MaxNetworkRetries` changes from `int` to `*int64`.
* `URL` changes from `string` to `*string`.

We made the change so that we can distinguish between a value that was let unset by the user versus one which has defaulted to empty (i.e. `MaxNetworkRetries` as a vanilla `int` left unset is indistinguishable from one that's been set to `0` explicitly).

Use the standard `stripe.*` class of parameter helpers to set these easily:

``` go
config := &stripe.BackendConfig{
    MaxNetworkRetries: stripe.Int64(2),
}
```

## `MaxNetworkRetries` now defaults to 2

The library now retries intermittently failed requests by default (i.e. if the feature has not been explicitly configured by the user) up to **two times**. This can be changed back to the previous behavior of no retries using `MaxNetworkRetries` on `BackendConfig`:

``` go
config := &stripe.BackendConfig{
    MaxNetworkRetries: stripe.Int64(0), // Zero retries
}
```

## `LeveledLogger` now default

Deprecated logging infrastructure has been removed, notably the top-level variables `stripe.LogLevel` and `stripe.Logger`. Use `DefaultLeveledLogger` instead:

``` go
stripe.DefaultLeveledLogger = &stripe.LeveledLogger{
    Level: stripe.LevelDebug,
}
```

Or set `LeveledLogger` explicitly on a `BackendConfig`:

``` go
config := &stripe.BackendConfig{
    LeveledLogger: &stripe.LeveledLogger{
        Level: stripe.LevelDebug,
    },
}
```

## Only errors are logged by default

If no logging has been configured, stripe-go will now show **only errors by default**. Previously, it logged both errors and informational messages. This behavior can be reverted by setting the default logger to informational:

``` go
stripe.DefaultLeveledLogger = &stripe.LeveledLogger{
    Level: stripe.LevelInfo,
}
```

## Other API changes

* [#1052](https://github.com/stripe/stripe-go/pull/1052) Beta features have been removed from Issuing APIs
* [#1068](https://github.com/stripe/stripe-go/pull/1068) Other breaking API changes
    * `PaymentIntent` is now expandable on `Charge`
    * `Percentage` was removed as a filter when listing `TaxRate`
    * Removed `RenewalInterval` on `SubscriptionSchedule`
    * Removed `Country` and `RoutingNumber` from `ChargePaymentMethodDetailsAcssDebit`
