---
title: "[Breaking] Update `List` and `Search` methods with `stripe.Client` to return a `struct`"
pr_link: https://github.com/stripe/stripe-go/pull/2179
is_breaking: true
released_in_version: 85.0.0
---

- `List` and `Search` methods using `stripe.Client` now return a `struct` instead of `Seq2`. This is a backwards incompatible change, and you will need to add an additional call to `.All(ctx)` in your `for` loop. E.g.

```diff
-for c, err := range sc.V1Customers.List(ctx, nil) {
+for c, err := range sc.V1Customers.List(ctx, nil).All(ctx) {
	// handle err
	// do something
}
```
- For manual pagination use cases, you can access the API call's `error` by calling `list.Err()`, a page's data using `list.Data()`, and its metadata by calling `list.Meta()`.
