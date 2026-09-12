### 🎉 Introducing new Stripe Client
Starting with v82.1, the new `stripe.Client` type is replacing `client.API` to provide a more ergonomic, consistent, and less error-prone experience. You create the former using `stripe.NewClient(stripeKey)`. It’s almost a drop-in replacement, except for the differences listed below.

1. Service method names now align with Stripe API docs. The `stripe.Client` uses `Create`, `Retrieve`, `Update`, and `Delete` (instead of `New`, `Get`, `Update`, and `Del`).
2. The first argument of each service method is a `context.Context`.
3. Parameter objects are now method-specific. For example, `CustomerCreateParams` and `CustomerDeleteParams` instead of simply `CustomerParams`. This allows us to put the right fields in the right methods at compile time.
4. Services are all version-namespaced for symmetry. E.g. `stripeClient.V1BillingMeterEvents` and `stripeClient.V2BillingMeterEvents`.
5. `List` methods return an `iter.Seq2`, so they can be ranged over without explicit calls to `Next`, `Current`, and `Err`.

  ### 🎉 Native support in Go for V2 APIs
```go
params := &stripe.V2CoreEventListParams{ObjectID: stripe.String("mtr_123")}
for event, err := range sc.V2CoreEvents.List(context.TODO(), params) {
  // handle err
  // process event object
}
```

More details can be found at https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
