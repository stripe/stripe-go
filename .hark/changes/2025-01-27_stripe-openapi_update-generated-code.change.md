---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1965
is_stripe_api_change: true
released_in_version: 81.3.0
---

* Add support for `Close` method on resource `Treasury.FinancialAccount`
* Add support for `PayByBankPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for `DirectorshipDeclaration` and `OwnershipExemptionReason` on `AccountCompanyParams`, `AccountCompany`, and `TokenAccountCompanyParams`
* Add support for `ProofOfUltimateBeneficialOwnership` on `AccountDocumentsParams`
* Add support for `FinancialAccount` on `AccountSessionComponentsParams`, `AccountSessionComponents`, and `TreasuryOutboundTransferDestinationPaymentMethodDetails`
* Add support for `FinancialAccountTransactions`, `IssuingCard`, and `IssuingCardsList` on `AccountSessionComponentsParams` and `AccountSessionComponents`
* Add support for `AdviceCode` on `ChargeOutcome`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, and `StripeError`
* Add support for `PayByBank` on `ChargePaymentMethodDetails`, `CheckoutSessionPaymentMethodOptionsParams`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `Country` on `ChargePaymentMethodDetailsPaypal`, `ConfirmationTokenPaymentMethodPreviewPaypal`, and `PaymentMethodPaypal`
* Add support for `Discounts` on `CheckoutSession`
* Add support for new value `SD` on enums `CheckoutSessionShippingAddressCollectionAllowedCountries` and `PaymentLinkShippingAddressCollectionAllowedCountries`
* Add support for new value `pay_by_bank` on enums `ConfirmationTokenPaymentMethodPreviewType` and `PaymentMethodType`
* Add support for `PhoneNumberCollection` on `PaymentLinkParams`
* Add support for new value `pay_by_bank` on enum `PaymentLinkPaymentMethodTypes`
* Add support for `Jpy` on `TerminalConfigurationTippingParams` and `TerminalConfigurationTipping`
* Add support for `Nickname` on `TreasuryFinancialAccountParams` and `TreasuryFinancialAccount`
* Add support for `ForwardingSettings` on `TreasuryFinancialAccountParams`
* Add support for `IsDefault` on `TreasuryFinancialAccount`
* Add support for `DestinationPaymentMethodData` on `TreasuryOutboundTransferParams`
* Change type of `TreasuryOutboundTransferDestinationPaymentMethodDetailsType` from `literal('us_bank_account')` to `enum('financial_account'|'us_bank_account')`
* Add support for `OutboundTransfer` on `TreasuryReceivedCreditLinkedFlowsSourceFlowDetails`
* Add support for new value `outbound_transfer` on enum `TreasuryReceivedCreditLinkedFlowsSourceFlowDetailsType`
