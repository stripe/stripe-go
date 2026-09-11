---
title: Fix all `Get` methods to support standardized parameter structs + remove some deprecated functions
pr_link: https://github.com/stripe/stripe-go/pull/589
released_in_version: 35.0.0
---

* `IssuerFraudRecordListParams` now uses `*string` for `Charge` (set it using `stripe.String` like elsewhere)
* `event.Get` now takes `stripe.EventParams` instead of `Params` for consistency
* The `Get` method for `countryspec`, `exchangerate`, `issuerfraudrecord` now take an extra params struct parameter to be consistent and allow setting a connected account (use `stripe.CountrySpecParams`, `stripe.ExchangeRateParams`, and `IssuerFraudRecordParams`)
* `charge.MarkFraudulent` and `charge.MarkSafe` have been removed; use `charge.Update` instead
* `charge.CloseDispute` and `charge.UpdateDispute` have been removed; use `dispute.Update` or `dispute.Close` instead
* `loginlink.New` now properly passes its params struct into its API call
