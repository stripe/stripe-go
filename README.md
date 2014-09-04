Go Stripe (BETA) [![Build Status](https://travis-ci.org/stripe/stripe.svg?branch=master)](https://travis-ci.org/stripe/stripe)
========

## Summary

The official [Stripe](https://stripe.com) Go client library.

## Versioning

Each revision of the binding is tagged and the version is updated accordingly.

Given Go's lack of built-in versioning, it is highly recommended you use a
[package mangement tool](https://code.google.com/p/go-wiki/wiki/PackageManagementTools) in order
to ensure a newer version of the binding does not affect backwards compatibility.

To see the list of past versions, run `git tag`. To manually get an older
version of the client, clone this repo, checkout the specific tag and build the
library:

```sh
git clone https://github.com/stripe/stripe.git
cd stripe
git checkout api_version_tag
make build
```

For more details on changes between versions, see the [binding changelog](CHANGELOG) 
and [API changelog](https://stripe.com/docs/upgrades).

## Installation

```sh
go get github.com/stripe/stripe
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

First import the package into your code:
```go
import (
    "github.com/stripe/stripe"
)
```

To use the client, initialize it and before making any requests:
```go
stripe := &stripe.Client{}
stripe.Init(YOUR_API_KEY, nil, nil)
```

The second parameter can be used to set a different http.Client (by default,
http.DefaultClient is used).  The third parameter can be used to inject a mock
Api implementation so calls aren't actually made to Stripe. This can be useful
for writing your own unit tests.

## APIs

While some resources may contain more/less APIs, the following pattern is
applied throughout the library for a given `resource`:

```go
// Create 
resource, err := stripe.Resources.Create(ResourceParams)

// Get
resource, err := stripe.Resources.Get(id)

// Update
resource, err := stripe.Resources.Get(ResourceParams)

// Delete
err := stripe.Resources.Delete(id)

// List
resourceList, err := stripe.Resources.List(ResourceListParams)
```

## Documentation

Below are a few simple examples. For details on all the functionality in this
library, see the [GoDoc](http://godoc.org/github.com/stripe/stripe) documentation.

For more details about Stripe, check out the [documentation](https://stripe.com/docs).

### Customers

```go
params := &CustomerParams{
		Balance: -123,
		Card: &CardParams{
			Name:   "Go Stripe",
			Number: "378282246310005",
			Month:  "06",
			Year:   "15",
		},
		Desc:  "Stripe Developer",
		Email: "gostripe@stripe.com",
	}

customer, err := stripe.Customers.Create(params)
```

### Charges

```go
params := &ChargeListParams{Customer: customer.Id}
params.Filters.AddFilter("include", "", "total_count")

charges, err := stripe.Charges.List(params)

for _, charge := range(charges.Values) {
   // perform an action on each charge
}
```
### Events

```go
events, err := stripe.Events.List(nil)

for _, e := range(events.Values) {
  // access event data via e.GetObjValue("resource_name_based_on_type", "resource_property_name")
  // alternatively you can access values via e.Data.Obj["resource_name_based_on_type"].(map[string]interface{})["resource_property_name"]

  // access previous attributes via e.GetPrevValue("resource_name_based_on_type", "resource_property_name")
  // alternatively you can access values via e.Data.Prev["resource_name_based_on_type"].(map[string]interface{})["resource_property_name"]
}
```

For any requests, bug or comments, please [open an issue](https://github.com/stripe/stripe/issues/new). 
