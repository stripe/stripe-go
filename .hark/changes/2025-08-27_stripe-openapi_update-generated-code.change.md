---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/2096
is_stripe_api_change: true
released_in_version: 82.6.0-beta.1
---

* Add support for `Get` and `List` methods on resource `InvoicePayment`
* Add support for `List` method on resource `Mandate`
* Add support for `Applied` on `V2CoreAccountConfigurationCustomerParams`, `V2CoreAccountConfigurationCustomer`, `V2CoreAccountConfigurationMerchantParams`, `V2CoreAccountConfigurationMerchant`, `V2CoreAccountConfigurationRecipientParams`, `V2CoreAccountConfigurationRecipient`, `V2CoreAccountConfigurationStorerParams`, and `V2CoreAccountConfigurationStorer`
* Add support for new values `ao_nif`, `az_tin`, `bd_etin`, `cr_cpj`, `cr_nite`, `do_rcn`, `gt_nit`, `kz_bin`, `mz_nuit`, `pe_ruc`, `pk_ntn`, `sa_crn`, and `sa_tin` on enum `V2CoreAccountIdentityBusinessDetailsIdNumbers.Type`
* Add support for new values `ao_nif`, `az_tin`, `bd_brc`, `bd_etin`, `bd_nid`, `cr_cpf`, `cr_dimex`, `cr_nite`, `do_rcn`, `gt_nit`, `kz_iin`, `mz_nuit`, `pe_dni`, `pk_cnic`, `pk_snic`, and `sa_tin` on enums `V2CoreAccountIdentityIndividualIdNumbers.Type` and `V2CorePersonIdNumbers.Type`
* Change type of `BillingAlertTriggered.Value` from `longInteger` to `decimal_string`
* Add support for `DisplayName` on `V2MoneyManagementFinancialAccountParams` and `V2MoneyManagementFinancialAccount`
* Add support for new value `currency_conversion` on enums `V2MoneyManagementTransaction.Category` and `V2MoneyManagementTransactionEntryTransactionDetails.Category`
* Add support for `CurrencyConversion` on `V2MoneyManagementTransactionEntryTransactionDetailsFlow` and `V2MoneyManagementTransactionFlow`
* Add support for new value `currency_conversion` on enums `V2MoneyManagementTransactionEntryTransactionDetailsFlow.Type` and `V2MoneyManagementTransactionFlow.Type`
* Add support for `Payments` on `BalanceSettingsParams` and `BalanceSettings`
* Remove support for `DebitNegativeBalances`, `Payouts`, and `SettlementTiming` on `BalanceSettingsParams` and `BalanceSettings`
* Add support for `Mandate` on `ChargePaymentMethodDetailsPix`, `PaymentAttemptRecordPaymentMethodDetailsPix`, and `PaymentRecordPaymentMethodDetailsPix`
* Add support for `CouponData` on `CheckoutSessionDiscountParams`
* Add support for `MandateOptions` on `CheckoutSessionPaymentMethodOptionsPixParams`, `CheckoutSessionPaymentMethodOptionsPix`, `PaymentIntentConfirmPaymentMethodOptionsPixParams`, `PaymentIntentPaymentMethodOptionsPixParams`, and `PaymentIntentPaymentMethodOptionsPix`
* Change type of `CheckoutSessionPaymentMethodOptionsPix.SetupFutureUsage`, `CheckoutSessionPaymentMethodOptionsPixParams.SetupFutureUsage`, `PaymentIntentConfirmPaymentMethodOptionsPixParams.SetupFutureUsage`, `PaymentIntentPaymentMethodOptionsPix.SetupFutureUsage`, and `PaymentIntentPaymentMethodOptionsPixParams.SetupFutureUsage` from `literal('none')` to `enum('none'|'off_session')`
* Add support for `Amount` on `MandateMultiUse`, `PaymentAttemptRecord`, and `PaymentRecord`
* Add support for `Currency` on `MandateMultiUse`
* Add support for `Pix` on `MandatePaymentMethodDetails`, `SetupAttemptPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodOptionsParams`, `SetupIntentPaymentMethodOptionsParams`, and `SetupIntentPaymentMethodOptions`
* Add support for `Limit` on `PaymentAttemptRecordListParams`
* Add support for `AmountAuthorized`, `AmountRefunded`, and `Application` on `PaymentAttemptRecord` and `PaymentRecord`
* Add support for `ProcessorDetails` on `PaymentAttemptRecord`, `PaymentRecordReportPaymentParams`, and `PaymentRecord`
* Remove support for `PaymentReference` on `PaymentAttemptRecord`, `PaymentRecordReportPaymentParams`, and `PaymentRecord`
* Add support for `Installments` on `PaymentAttemptRecordPaymentMethodDetailsAlma` and `PaymentRecordPaymentMethodDetailsAlma`
* Add support for `TransactionID` on `PaymentAttemptRecordPaymentMethodDetailsAlma`, `PaymentAttemptRecordPaymentMethodDetailsAmazonPay`, `PaymentAttemptRecordPaymentMethodDetailsBillie`, `PaymentAttemptRecordPaymentMethodDetailsKakaoPay`, `PaymentAttemptRecordPaymentMethodDetailsKrCard`, `PaymentAttemptRecordPaymentMethodDetailsNaverPay`, `PaymentAttemptRecordPaymentMethodDetailsPayco`, `PaymentAttemptRecordPaymentMethodDetailsRevolutPay`, `PaymentAttemptRecordPaymentMethodDetailsSamsungPay`, `PaymentAttemptRecordPaymentMethodDetailsSatispay`, `PaymentRecordPaymentMethodDetailsAlma`, `PaymentRecordPaymentMethodDetailsAmazonPay`, `PaymentRecordPaymentMethodDetailsBillie`, `PaymentRecordPaymentMethodDetailsKakaoPay`, `PaymentRecordPaymentMethodDetailsKrCard`, `PaymentRecordPaymentMethodDetailsNaverPay`, `PaymentRecordPaymentMethodDetailsPayco`, `PaymentRecordPaymentMethodDetailsRevolutPay`, `PaymentRecordPaymentMethodDetailsSamsungPay`, and `PaymentRecordPaymentMethodDetailsSatispay`
* Add support for `Location` and `Reader` on `PaymentAttemptRecordPaymentMethodDetailsPaynow` and `PaymentRecordPaymentMethodDetailsPaynow`
* Add support for `LatestActiveMandate` on `PaymentMethod`
* Add support for `Metadata` and `Period` on `QuotePreviewSubscriptionSchedulePhaseAddInvoiceItem`
* Add support for `PixDisplayQRCode` on `SetupIntentNextAction`
* Add support for `ReaderSecurity` on `TerminalConfigurationParams` and `TerminalConfiguration`
* Add support for error codes `customer_session_expired` and `india_recurring_payment_mandate_canceled` on `QuotePreviewInvoiceLastFinalizationError`
