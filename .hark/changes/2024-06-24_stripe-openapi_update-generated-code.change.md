---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1878
is_stripe_api_change: true
released_in_version: 79.0.0
---

This release changes the pinned API version to 2024-06-20. Please read the [API Changelog](https://docs.stripe.com/changelog/2024-06-20) and carefully review the API changes before upgrading.

### ⚠️ Breaking changes

  * Remove the unused resource `PlatformTaxFee`
  * Rename `VolumeDecimal` to `QuantityDecimal` on `IssuingTransactionPurchaseDetailsFuel`, `TestHelpersIssuingAuthorizationCapturePurchaseDetailsFuelParams`, `TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFuelParams`, and `TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFuelParams`

## Additions

* Add support for `FinalizeAmount` test helper method on resource `Issuing.Authorization`
* Add support for new value `ch_uid` on enums `CheckoutSessionCustomerDetailsTaxIdsType`, `InvoiceCustomerTaxIdsType`, `TaxCalculationCustomerDetailsTaxIdsType`, `TaxIdType`, and `TaxTransactionCustomerDetailsTaxIdsType`
* Add support for `Fleet` on `IssuingAuthorizationParams`, `IssuingAuthorization`, `IssuingTransactionPurchaseDetails`, `TestHelpersIssuingAuthorizationCapturePurchaseDetailsParams`, `TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsParams`, and `TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsParams`
* Add support for `Fuel` on `IssuingAuthorizationParams` and `IssuingAuthorization`
* Add support for `IndustryProductCode` and `QuantityDecimal` on `IssuingTransactionPurchaseDetailsFuel`, `TestHelpersIssuingAuthorizationCapturePurchaseDetailsFuelParams`, `TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsFuelParams`, and `TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsFuelParams`
* Add support for new values `card_canceled`, `card_expired`, `cardholder_blocked`, `insecure_authorization_method`, and `pin_blocked` on enum `IssuingAuthorizationRequestHistoryReason`
