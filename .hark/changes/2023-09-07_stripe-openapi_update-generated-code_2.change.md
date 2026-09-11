---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/1725
is_stripe_api_change: true
released_in_version: 75.5.0-beta.1
---

* Remove support for `SubmitCard` test helper method on resource `Issuing.Card`
* Add support for `TaxForms` on `AccountSettingsParams` and `AccountSettings`
* Add support for `CardDesign` on `IssuingCardParams`
* Remove support for value `submitted` from enum `IssuingCardShippingStatus`
* Add support for new value `platform_default` on enum `IssuingCardDesignPreference`
