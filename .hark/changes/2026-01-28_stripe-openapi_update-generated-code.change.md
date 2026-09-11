---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/2258
is_stripe_api_change: true
released_in_version: 84.3.0
---

* Add support for new resource `RadarPaymentEvaluation`
* Add support for `New` method on resource `RadarPaymentEvaluation`
* Add support for `AdjustableQuantity` on `LineItem`
* Add support for new value `risk_reserved` on enum `BalanceTransaction.BalanceType`
* Add support for new values `reserve_hold` and `reserve_release` on enum `BalanceTransaction.Type`
* Add support for new value `pl_nip` on enums `CheckoutSessionCustomerDetailsTaxIds.Type`, `TaxCalculationCustomerDetailsTaxId.Type`, `TaxId.Type`, and `TaxTransactionCustomerDetailsTaxId.Type`
* Add support for new value `adyen` on enums `ConfirmationTokenPaymentMethodPreviewIdeal.Bank`, `PaymentAttemptRecordPaymentMethodDetailsIdeal.Bank`, and `PaymentRecordPaymentMethodDetailsIdeal.Bank`
* Add support for new value `ADYBNL2A` on enums `ConfirmationTokenPaymentMethodPreviewIdeal.BIC`, `PaymentAttemptRecordPaymentMethodDetailsIdeal.BIC`, and `PaymentRecordPaymentMethodDetailsIdeal.BIC`
* Add support for `EnforceArithmeticValidation` on `PaymentIntentAmountDetailsParams`, `PaymentIntentCaptureAmountDetailsParams`, `PaymentIntentConfirmAmountDetailsParams`, and `PaymentIntentIncrementAuthorizationAmountDetailsParams`
* Add support for `Error` on `PaymentIntentAmountDetails`
* Remove support for `Bgn` on `TerminalConfigurationTippingParams` and `TerminalConfigurationTipping`
* Add support for `Topup` on `TreasuryReceivedDebitLinkedFlows`
* Add support for `ContactPhone` on `V2CoreAccountParams`, `V2CoreAccountTokenParams`, and `V2CoreAccount`
* Add support for `RegistrationDate` on `V2CoreAccountIdentityBusinessDetailsParams`, `V2CoreAccountIdentityBusinessDetails`, and `V2CoreAccountTokenIdentityBusinessDetailsParams`
* Add support for new value `gb_vat` on enum `V2CoreAccountIdentityBusinessDetailsIdNumber.Type`
* Add support for error code `request_blocked` on `Error`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, and `StripeError`
