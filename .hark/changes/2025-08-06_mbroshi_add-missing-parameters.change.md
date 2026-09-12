---
title: Add missing parameters to Card and BankAccount services
pr_url: https://github.com/stripe/stripe-go/pull/2102
released_in_version: 82.4.1
---

* Fixes bugs in `V1Cards` and `V1BankAccounts` services: ensures the new `stripe.Client` pattern supports all parameters previously available in `client.API` for those two services.
