---
title: Make more files codegen-able
is_stripe_api_change: true
released_in_version: 72.74.0
---

- Add support for `acss_debit`, `au_becs_debit`, `bacs_debit`, and `sepa_debit` on `SetupAttemptPaymentMethodDetails`
- Add support for `setup_intent` on `SetupAttempt`
- Add support for `duplicate` option for `SetupIntentCancellationReason`
- Add support for `challenge_only` option for `SetupIntentPaymentMethodOptionsCardRequestThreeDSecure`
- Add support for `sepa_debit` on `SetupIntentPaymentMethodOptionsParams` and `SetupIntentPaymentMethodOptions`
- Add support for `client_secret` on `SetupIntentParams`
