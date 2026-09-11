---
title: Improved handling for enums in params
pr_link: https://github.com/stripe/stripe-go/pull/2022
released_in_version: 82.1.0
---

* You can now pass `string` enums into `stripe.String`. For example, `stripe.String(stripe.CurrencyUSD)` instead of `stripe.String(string(stripe.CurrencyUSD))`
