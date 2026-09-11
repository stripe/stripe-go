---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/1732
is_stripe_api_change: true
released_in_version: 75.7.0-beta.1
---

* Add support for new resource `ConfirmationToken`
* Add support for `Get` method on resource `ConfirmationToken`
* Add support for `New` method on resource `Issuing.CardDesign`
* Add support for `RejectTestmode` test helper method on resource `Issuing.CardDesign`
* Add support for new value `issuing_card_design.rejected` on enum `EventType`
* Add support for `Features` on `IssuingCardBundle`
* Add support for `Preferences` on `IssuingCardDesignListParams`, `IssuingCardDesignParams`, and `IssuingCardDesign`
* Remove support for `Preference` on `IssuingCardDesignListParams`, `IssuingCardDesignParams`, and `IssuingCardDesign`
* Add support for `CardBundle` on `IssuingCardDesignParams`
* Add support for `CardLogo` and `CarrierText` on `IssuingCardDesignParams` and `IssuingCardDesign`
* Change type of `IssuingCardDesignLookupKeyParams` and `IssuingCardDesignNameParams` from `string` to `emptyStringable(string)`
* Add support for `RejectionReasons` on `IssuingCardDesign`
* Add support for `ConfirmationToken` on `PaymentIntentConfirmParams`, `PaymentIntentParams`, `SetupIntentConfirmParams`, and `SetupIntentParams`
