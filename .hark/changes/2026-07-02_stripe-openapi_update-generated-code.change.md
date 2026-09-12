---
title: Update generated code for private-preview
pr_url: https://github.com/stripe/stripe-go/pull/2381
is_breaking: true
is_stripe_api_change: true
released_in_version: 86.2.0-alpha.2
---

* Add support for new resources `CryptoCustomerConsumerWallet`, `CryptoCustomerPaymentToken`, `CryptoCustomer`, `CryptoOnrampSession`, and `CryptoOnrampTransactionLimits`
* Add support for `Get` and `List` methods on resource `CryptoCustomer`
* Add support for `Checkout`, `Get`, `List`, `New`, and `Quote` methods on resource `CryptoOnrampSession`
* Add support for `Get` method on resource `CryptoOnrampTransactionLimits`
* Add support for `ElectronicCommerceIndicator` on `ChargePaymentMethodDetailsCard`
* Add support for `AmountReceived` and `AmountRequested` on `ChargePaymentMethodDetailsCrypto`, `PaymentAttemptRecordPaymentMethodDetailsCrypto`, and `PaymentRecordPaymentMethodDetailsCrypto`
* Add support for `Fingerprint` on `ChargePaymentMethodDetailsGiftCard`, `PaymentAttemptRecordPaymentMethodDetailsGiftCard`, and `PaymentRecordPaymentMethodDetailsGiftCard`
* Add support for `AddressCollectionPrecision` on `CheckoutSessionAutomaticTaxParams`
* Add support for `Subscription` on `CheckoutSessionItem`
* ⚠️  Remove support for `Deactivation` on `GiftCardOperation`
* ⚠️  Remove support for value `deactivation` from enum `GiftCardOperation.Type`
* Add support for `MerchantAmountExchangeRate` on `IssuingAuthorization` and `IssuingTransaction`
* Add support for `DeviceID` on `IssuingAuthorizationTokenDetailsNetworkDataDevice` and `IssuingTokenNetworkDataDevice`
* Add support for `Program` on `IssuingCard`
* Add support for `PaymentMethodDetails` on `PaymentAttemptRecordReportFailedParams` and `PaymentRecordReportPaymentAttemptFailedParams`
* Add support for `Reason` on `PaymentAttemptRecordReportRefundParams` and `PaymentRecordReportRefundParams`
* Add support for `AmountReconciliation` on `PaymentIntentConfirmPaymentMethodOptionsCryptoParams`, `PaymentIntentPaymentMethodOptionsCryptoParams`, and `PaymentIntentPaymentMethodOptionsCrypto`
* Add support for `ConnectPermissions` and `Permissions` on `V2IamApiKeyParams` and `V2IamApiKey`
* Add support for `Credit` on `V2MoneyManagementFinancialAccount`
* Add support for new value `credit` on enum `V2MoneyManagementFinancialAccount.Type`
* Add support for new value `currency_required` on enum `V2MoneyManagementPayoutIntentNextActionHandleFailure.FailureReason`
* Add support for new values `issuing_authorization`, `issuing_transaction`, and `platform_funded_credit_transaction` on enums `V2MoneyManagementTransaction.Category` and `V2MoneyManagementTransactionEntryTransactionDetails.Category`
* Add support for `Account`, `IssuingAuthorization`, `IssuingDispute`, and `IssuingTransaction` on `V2MoneyManagementTransactionEntryTransactionDetailsFlow` and `V2MoneyManagementTransactionFlow`
* Add support for new values `issuing_authorization`, `issuing_dispute`, and `issuing_transaction` on enums `V2MoneyManagementTransactionEntryTransactionDetailsFlow.Type` and `V2MoneyManagementTransactionFlow.Type`
* Change type of `V2MoneyManagementFinancialAccountParams.Type` from `literal('storage')` to `enum('credit'|'storage')`
* Add support for `ExpiresAt` on `V2IamApiKeyParams`
