---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/2216
is_breaking: true
is_stripe_api_change: true
released_in_version: 84.0.0
---

* ⚠️ Change the type of `V2CoreEventDestinationParams.Metadata` and `V2CoreEventDestinationUpdateParams.Metadata` to `map[string]*string` from `map[string]string`. This supports the ability to remove a key from a `Metadata` map by setting its value to `nil`.
* ⚠️ A corresponding change was made to the `V2CoreEventDestinationParams.AddMetadata` method to set its second argument to `*string` from `string`.
