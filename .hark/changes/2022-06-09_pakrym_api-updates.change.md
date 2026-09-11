---
title: API Updates
pr_link: https://github.com/stripe/stripe-go/pull/1473
is_stripe_api_change: true
released_in_version: 72.114.0
---

* Add support for `Treasury` on `AccountSettingsParams` and `AccountSettings`
* Add support for `RenderingOptions` on `CustomerInvoiceSettingsParams`
* Add support for `EUBankTransfer` on `CustomerCreateFundingInstructionsBankTransferParams`, `InvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferParams`, `InvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransfer`, `OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferParams`, `OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransfer`, `PaymentIntentConfirmPaymentMethodOptionsCustomerBalanceBankTransferParams`, `PaymentIntentPaymentMethodOptionsCustomerBalanceBankTransferParams`, `PaymentIntentPaymentMethodOptionsCustomerBalanceBankTransfer`, `SubscriptionPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferParams`, and `SubscriptionPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransfer`
* Change type of `CustomerCreateFundingInstructionsBankTransferRequestedAddressTypesParams` from `literal('zengin')` to `enum('iban'|'sort_code'|'spei'|'zengin')`
* Change type of `CustomerCreateFundingInstructionsBankTransferTypeParams`, `OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferTypeParams`, `OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferType`, `PaymentIntentConfirmPaymentMethodOptionsCustomerBalanceBankTransferTypeParams`, `PaymentIntentNextActionDisplayBankTransferInstructionsType`, `PaymentIntentPaymentMethodOptionsCustomerBalanceBankTransferTypeParams`, and `PaymentIntentPaymentMethodOptionsCustomerBalanceBankTransferType` from `literal('jp_bank_transfer')` to `enum('eu_bank_transfer'|'gb_bank_transfer'|'jp_bank_transfer'|'mx_bank_transfer')`
* Add support for `Iban`, `SortCode`, and `Spei` on `FundingInstructionsBankTransferFinancialAddresses` and `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddresses`
* Add support for new values `bacs`, `fps`, and `spei` on enums `FundingInstructionsBankTransferFinancialAddressesSupportedNetworks` and `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressesSupportedNetworks`
* Add support for new values `sort_code` and `spei` on enums `FundingInstructionsBankTransferFinancialAddressesType` and `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressesType`
* Change type of `OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypesParams`, `OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypes`, `PaymentIntentConfirmPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypesParams`, `PaymentIntentPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypesParams`, and `PaymentIntentPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypes` from `literal('zengin')` to `enum`
* Add support for `CustomUnitAmount` on `PriceParams` and `Price`
