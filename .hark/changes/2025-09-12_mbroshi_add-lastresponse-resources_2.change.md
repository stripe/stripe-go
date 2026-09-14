---
title: Add LastResponse to resources returned in List and Search APIs
pr_url: https://github.com/stripe/stripe-go/pull/2117
released_in_version: 82.5.1
---

  Add a `LastResponse` to each resource returned from either a `List` or `Search` API call using `stripe.Client`. The `RawJSON` is the JSON corresponding to just that item. This is useful for accessing fields not exposed in the SDK.

```go
for cust, err := range sc.V1Customers.List(context.TODO(), &stripe.CustomerListParams{}) {
    if err != nil {
        return err
    }
    customerJSON := cust.LastResponse.RawJSON
    log.Printf("Customer JSON: %s", customerJSON) // {"id":"cus_123",...}
}
```
