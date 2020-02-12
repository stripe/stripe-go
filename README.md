# Go Stripe

[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/stripe/stripe-go)
[![Build Status](https://travis-ci.org/stripe/stripe-go.svg?branch=master)](https://travis-ci.org/stripe/stripe-go)
[![Coverage Status](https://coveralls.io/repos/github/stripe/stripe-go/badge.svg?branch=master)](https://coveralls.io/github/stripe/stripe-go?branch=master)

The official [Stripe][stripe] Go client library.

## Installation

Install stripe-go with:

```sh
go get -u github.com/stripe/stripe-go
```

Then, import it using:

``` go
import (
    "github.com/stripe/stripe-go"
    "github.com/stripe/stripe-go/customer"
)
```

### Go Module Support

The library currently *does not* ship with first-class support for Go
modules. We put in support for it before, but ran into compatibility problems
for existing installations using Dep (see discussion in [closer to the bottom
of this thread][gomodvsdep]), and [reverted support][gomodrevert]. Our current
plan is to wait for better module compatibility in Dep (see a [preliminary
patch here][depgomodsupport]), give the release a little grace time to become
more widely distributed, then bring support back.

For now, require stripe-go in `go.mod` with a version but without a *version
suffix* in the path like so:

``` go
module github.com/my/package

require (
    github.com/stripe/stripe-go v68.20.0
)
```

And use the same style of import paths as above:

``` go
import (
    "github.com/stripe/stripe-go"
    "github.com/stripe/stripe-go/customer"
)
```

## Documentation

For a comprehensive list of examples, check out the [API
documentation][api-docs].

For details on all the functionality in this library, see the [GoDoc][godoc]
documentation.

Below are a few simple examples:

### Customers

```go
params := &stripe.CustomerParams{
	Description: stripe.String("Stripe Developer"),
	Email:       stripe.String("gostripe@stripe.com"),
}

customer, err := customer.New(params)
```

### PaymentIntents

```go
params := &stripe.PaymentIntentListParams{
    Customer: stripe.String(customer.ID),
}

// set this so you can easily retry your request in case of a timeout
params.Params.IdempotencyKey = stripe.NewIdempotencyKey()

i := paymentintent.List(params)
for i.Next() {
	pi := i.PaymentIntent()
}

if err := i.Err(); err != nil {
	// handle
}
```

### Events

```go
i := event.List(nil)
for i.Next() {
	e := i.Event()

	// access event data via e.GetObjectValue("resource_name_based_on_type", "resource_property_name")
	// alternatively you can access values via e.Data.Object["resource_name_based_on_type"].(map[string]interface{})["resource_property_name"]

	// access previous attributes via e.GetPreviousValue("resource_name_based_on_type", "resource_property_name")
	// alternatively you can access values via e.Data.PrevPreviousAttributes["resource_name_based_on_type"].(map[string]interface{})["resource_property_name"]
}
```

Alternatively, you can use the `event.Data.Raw` property to unmarshal to the
appropriate struct.

### Authentication with Connect

There are two ways of authenticating requests when performing actions on behalf
of a connected account, one that uses the `Stripe-Account` header containing an
account's ID, and one that uses the account's keys. Usually the former is the
recommended approach. [See the documentation for more information][connect].

To use the `Stripe-Account` approach, use `SetStripeAccount()` on a `ListParams`
or `Params` class. For example:

```go
// For a list request
listParams := &stripe.CustomerListParams{}
listParams.SetStripeAccount("acct_123")
```

```go
// For any other kind of request
params := &stripe.CustomerParams{}
params.SetStripeAccount("acct_123")
```

To use a key, pass it to `API`'s `Init` function:

```go

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
)

stripe := &client.API{}
stripe.Init("access_token", nil)
```

### Google AppEngine

If you're running the client in a Google AppEngine environment, you'll need to
create a per-request Stripe client since the `http.DefaultClient` is not
available. Here's a sample handler:

```go
import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
)

func handler(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        httpClient := urlfetch.Client(c)

        sc := stripeClient.New("sk_test_123", stripe.NewBackends(httpClient))

        params := &stripe.CustomerParams{
            Description: stripe.String("Stripe Developer"),
            Email:       stripe.String("gostripe@stripe.com"),
        }
        customer, err := sc.Customers.New(params)
        if err != nil {
            fmt.Fprintf(w, "Could not create customer: %v", err)
        }
        fmt.Fprintf(w, "Customer created: %v", customer.ID)
}
```

## Usage

While some resources may contain more/less APIs, the following pattern is
applied throughout the library for a given `$resource$`:

### Without a Client

If you're only dealing with a single key, you can simply import the packages
required for the resources you're interacting with without the need to create a
client.

```go
import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/$resource$"
)

// Setup
stripe.Key = "sk_key"

stripe.SetBackend("api", backend) // optional, useful for mocking

// Create
$resource$, err := $resource$.New(stripe.$Resource$Params)

// Get
$resource$, err := $resource$.Get(id, stripe.$Resource$Params)

// Update
$resource$, err := $resource$.Update(stripe.$Resource$Params)

// Delete
resourceDeleted, err := $resource$.Del(id, stripe.$Resource$Params)

// List
i := $resource$.List(stripe.$Resource$ListParams)
for i.Next() {
	$resource$ := i.$Resource$()
}

if err := i.Err(); err != nil {
	// handle
}
```

### With a Client

If you're dealing with multiple keys, it is recommended you use `client.API`.
This allows you to create as many clients as needed, each with their own
individual key.

```go
import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
)

// Setup
sc := &client.API{}
sc.Init("sk_key", nil) // the second parameter overrides the backends used if needed for mocking

// Create
$resource$, err := sc.$Resource$s.New(stripe.$Resource$Params)

// Get
$resource$, err := sc.$Resource$s.Get(id, stripe.$Resource$Params)

// Update
$resource$, err := sc.$Resource$s.Update(stripe.$Resource$Params)

// Delete
resourceDeleted, err := sc.$Resource$s.Del(id, stripe.$Resource$Params)

// List
i := sc.$Resource$s.List(stripe.$Resource$ListParams)
for i.Next() {
	resource := i.$Resource$()
}

if err := i.Err(); err != nil {
	// handle
}
```

### Configuring Automatic Retries

You can enable automatic retries on requests that fail due to a transient
problem by configuring the maximum number of retries:

```go
import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
)

config := &stripe.BackendConfig{
    MaxNetworkRetries: 2,
}

sc := &client.API{}
sc.Init("sk_key", &stripe.Backends{
    API:     stripe.GetBackendWithConfig(stripe.APIBackend, config),
    Uploads: stripe.GetBackendWithConfig(stripe.UploadsBackend, config),
})

coupon, err := sc.Coupons.New(...)
```

Various errors can trigger a retry, like a connection error or a timeout, and
also certain API responses like HTTP status `409 Conflict`.

[Idempotency keys][idempotency-keys] are added to requests to guarantee that
retries are safe.

### Configuring Logging

Configure logging using the global `DefaultLeveledLogger` variable:

```go
stripe.DefaultLeveledLogger = &stripe.LeveledLogger{
    Level: stripe.LevelInfo,
}
```

Or on a per-backend basis:

```go
config := &stripe.BackendConfig{
    LeveledLogger: &stripe.LeveledLogger{
        Level: stripe.LevelInfo,
    },
}
```

It's possible to use non-Stripe leveled loggers as well. Stripe expects loggers
to comply to the following interface:

```go
type LeveledLoggerInterface interface {
	Debugf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
}
```

Some loggers like [Logrus][logrus] and Zap's [SugaredLogger][sugaredlogger]
support this interface out-of-the-box so it's possible to set
`DefaultLeveledLogger` to a `*logrus.Logger` or `*zap.SugaredLogger` directly.
For others it may be necessary to write a thin shim layer to support them.

### Writing a Plugin

If you're writing a plugin that uses the library, we'd appreciate it if you
identified using `stripe.SetAppInfo`:

```go
stripe.SetAppInfo(&stripe.AppInfo{
    Name:    "MyAwesomePlugin",
    URL:     "https://myawesomeplugin.info",
    Version: "1.2.34",
})
```

This information is passed along when the library makes calls to the Stripe
API. Note that while `Name` is always required, `URL` and `Version` are
optional.

### Request latency telemetry

By default, the library sends request latency telemetry to Stripe. These
numbers help Stripe improve the overall latency of its API for all users.

You can disable this behavior if you prefer:

```go
config := &stripe.BackendConfig{
	EnableTelemetry: false,
}
```

## Development

Pull requests from the community are welcome. If you submit one, please keep
the following guidelines in mind:

1. Code must be `go fmt` compliant.
2. All types, structs and funcs should be documented.
3. Ensure that `make test` succeeds.

## Test

The test suite needs testify's `require` package to run:

    github.com/stretchr/testify/require

Before running the tests, make sure to grab all of the package's dependencies:

    go get -t -v

It also depends on [stripe-mock][stripe-mock], so make sure to fetch and run it from a
background terminal ([stripe-mock's README][stripe-mock-usage] also contains
instructions for installing via Homebrew and other methods):

    go get -u github.com/stripe/stripe-mock
    stripe-mock

Run all tests:

    make test

Run tests for one package:

    go test ./invoice

Run a single test:

    go test ./invoice -run TestInvoiceGet

For any requests, bug or comments, please [open an issue][issues] or [submit a
pull request][pulls].

[api-docs]: https://stripe.com/docs/api/go
[api-changelog]: https://stripe.com/docs/upgrades
[connect]: https://stripe.com/docs/connect/authentication
[depgomodsupport]: https://github.com/golang/dep/pull/1963
[godoc]: http://godoc.org/github.com/stripe/stripe-go
[gomodrevert]: https://github.com/stripe/stripe-go/pull/774
[gomodvsdep]: https://github.com/stripe/stripe-go/pull/712
[idempotency-keys]: https://stripe.com/docs/api/ruby#idempotent_requests
[issues]: https://github.com/stripe/stripe-go/issues/new
[logrus]: https://github.com/sirupsen/logrus/
[modules]: https://github.com/golang/go/wiki/Modules
[package-management]: https://code.google.com/p/go-wiki/wiki/PackageManagementTools
[pulls]: https://github.com/stripe/stripe-go/pulls
[stripe]: https://stripe.com
[stripe-mock]: https://github.com/stripe/stripe-mock
[stripe-mock-usage]: https://github.com/stripe/stripe-mock#usage

<!--
# vim: set tw=79:
-->
