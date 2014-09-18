Go Stripe (BETA) [![Build Status](https://travis-ci.org/stripe/stripe-go.svg?branch=master)](https://travis-ci.org/stripe/stripe-go)
========

## Summary

The official [Stripe](https://stripe.com) Go client library.

**NOTE**: This library is still in beta.  We may make significant API changes between
minor versions until the 1.0 release.  We will not make breaking changes in patch versions.

## Versioning

Each revision of the binding is tagged and the version is updated accordingly.

Given Go's lack of built-in versioning, it is highly recommended you use a
[package mangement tool](https://code.google.com/p/go-wiki/wiki/PackageManagementTools) in order
to ensure a newer version of the binding does not affect backwards compatibility.

To see the list of past versions, run `git tag`. To manually get an older
version of the client, clone this repo, checkout the specific tag and build the
library:

```sh
git clone https://github.com/stripe/stripe-go.git
cd stripe
git checkout api_version_tag
make build
```

For more details on changes between versions, see the [binding changelog](CHANGELOG)
and [API changelog](https://stripe.com/docs/upgrades).

## Installation

```sh
go get github.com/stripe/stripe-go
```

### Development

Pull requests from the community are welcome. If you submit one, please keep
the following guidelines in mind:

1. Code should be `go fmt` compliant.
2. All types, structs and funcs should be documented.
3. Ensure that `make checkin` succeeds.

## Test

For running additional tests, follow the steps below:

Set the STRIPE_KEY environment variable to match your test private key:
```sh
export STRIPE_KEY=YOUR_API_KEY
```

Then run:
```sh
make test
```

## Usage

While some resources may contain more/less APIs, the following pattern is
applied throughout the library for a given `resource`:

### Without a Client

If you're only dealing with a single key, you can simply import the packages
required for the resources you're interacting with without the need to create a
client.

```go
import (
  "github.com/stripe/stripe-go"
  "github.com/stripe/stripe-go/resource"
)

// Setup
stripe.Key = "sk_key"

stripe.SetBackend(backend) // optional, useful for mocking

// Create
resource, err := resource.New(ResourceParams)

// Get
resource, err := resource.Get(id, ResourceParams)

// Update
resource, err := resource.Update(ResourceParams)

// Delete
err := resource.Del(id)

// List
i := resource.List(ResourceListParams)
for !i.Stop() {
  resource, err := i.Next()
}
```

### With a Client

If you're dealing with multiple keys, it is recommended you use the
`client.Api`.  This allows you to create as many clients as needed, each with
their own individual key.

```go
import (
  "github.com/stripe/stripe-go"
  "github.com/stripe/stripe-go/client"
)

// Setup
stripe := &client.Api{}
stripe.Init("sk_key", nil)
// the second parameter represents the Backend used by the client. It can be
// useful to set one explicitly to either get a custom http.Client or mock it
// entirely in tests.

// Create
resource, err := stripe.Resources.New(ResourceParams)

// Get
resource, err := stripe.Resources.Get(id, ResourceParams)

// Update
resource, err := stripe.Resources.Update(ResourceParams)

// Delete
err := stripe.Resources.Del(id)

// List
i := stripe.Resources.List(ResourceListParams)
for !i.Stop() {
  resource, err := i.Next()
}
```

### Connect Flows

If you're using an `access token` you will need to use a client. Simply pass
the `access token` value as the `tok` when initalizing the client.

```go

import (
  "github.com/stripe/stripe-go"
  "github.com/stripe/stripe-go/client"
)

stripe := &client.Api{}
stripe.Init("access_token", nil)
```

## Documentation

Below are a few simple examples. For details on all the functionality in this
library, see the [GoDoc](http://godoc.org/github.com/stripe/stripe-go) documentation.

For more details about Stripe, check out the [documentation](https://stripe.com/docs).

### Customers

```go
params := &stripe.CustomerParams{
		Balance: -123,
		Card: &stripe.CardParams{
			Name:   "Go Stripe",
			Number: "378282246310005",
			Month:  "06",
			Year:   "15",
		},
		Desc:  "Stripe Developer",
		Email: "gostripe@stripe.com",
	}

customer, err := customer.New(params)
```

### Charges

```go
params := &stripe.ChargeListParams{Customer: customer.Id}
params.Filters.AddFilter("include", "", "total_count")

i := charge.List(params)
for !i.Stop() {
  c, err := i.Next()
  // perform an action on each charge
}
```
### Events

```go
i := event.List(nil)
for !i.Stop() {
  e, err := i.Next()

  // access event data via e.GetObjValue("resource_name_based_on_type", "resource_property_name")
  // alternatively you can access values via e.Data.Obj["resource_name_based_on_type"].(map[string]interface{})["resource_property_name"]

  // access previous attributes via e.GetPrevValue("resource_name_based_on_type", "resource_property_name")
  // alternatively you can access values via e.Data.Prev["resource_name_based_on_type"].(map[string]interface{})["resource_property_name"]
}
```

For any requests, bug or comments, please [open an issue](https://github.com/stripe/stripe-go/issues/new).
