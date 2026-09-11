---
title: Bug fixes for `V1BankAccounts` and `V1Cards` services
pr_link: https://github.com/stripe/stripe-go/pull/2098
released_in_version: 82.4.1
---

* Fixes bugs in `Create` and `List` methods in `V1BankAccounts` and `V1Cards` services in the `stripe.Client` pattern, which were previously returning errors on any valid inputs.
