---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/1665
is_stripe_api_change: true
released_in_version: 74.22.0-beta.1
---

* Add support for `SubscriptionDetails` on `Invoice`
* Add support for new values `aba` and `swift` on enum `OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypes`
* Add support for new value `us_bank_transfer` on enum `OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferType`
* Add support for `SetPauseCollection` on `QuoteLine`, `QuoteLinesParams`, and `SubscriptionScheduleAmendAmendmentsParams`
* Add support for new value `pause_collection_start` on enums `QuoteSubscriptionDataBillOnAcceptanceBillFromType` and `QuoteSubscriptionDataOverridesBillOnAcceptanceBillFromType`
* Add support for `PauseCollection` on `SubscriptionSchedulePhasesParams` and `SubscriptionSchedulePhases`
* Add support for `LocalAmusementTax` on `TaxRegistrationCountryOptionsUsParams` and `TaxRegistrationCountryOptionsUs`
* Remove support for `Locations` on `TaxSettingsParams` and `TaxSettings`
