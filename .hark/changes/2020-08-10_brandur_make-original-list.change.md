---
title: Make original list object accessible on iterators
pr_url: https://github.com/stripe/stripe-go/pull/1148
released_in_version: 71.44.0
---

* This change is technically breaking in that an exported type, `stripe.Query`, changes from `type Query func(*Params, *form.Values) ([]interface{}, ListMeta, error)` to `type Query func(*Params, *form.Values) ([]interface{}, ListContainer, error)`. We've opted to ship this as a minor version anyway because although exported, `Query` is meant for internal use in other stripe-go packages and the vast majority of users are unlikely to be referencing it. If you are, please refer to the diff in https://github.com/stripe/stripe-go/pull/1148 for how to update callsites accordingly. If you think there is a major use of `Query` that we've likely overlooked, please open an issue.
