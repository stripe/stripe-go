# Changelog

## Added

- `CheckoutSessionSetupIntentDataParams.Metadata`
- Invoice `UpcomingLines` method

## Changed

- Renamed files
  | Old name | New name |
  | --- | --- |
  | fee.go | applicationfee.go |
  | fee/client.go | applicationfee/client.go |
  | sub.go | subscription.go |
  | sub/client.go | subscription/client.go |
  | subitem.go | subscriptionitem.go |
  | subitem/client.go | subscriptionitem/client.go |
  | subschedule.go | subscriptionschedule.go |
  | subschedule/client.go | subscriptionschedule/client.go |
  | reversal.go | transferreversal.go |
  | reversal/client.go | transferreversal/client.go |

- Renamed structs, fields, enums, and methods to be consistent with the library's naming conventions and with the other Stripe SDKs.

  - `Ach` to `ACH`
    | Resource | Old name | New name |
    | --- | --- | --- |
    | Account | `AccountCapabilitiesUSBankAccountAchPayments` | `AccountCapabilitiesUSBankAccountACHPayments` |
    | Account | `AccountCapabilitiesUSBankAccountAchPaymentsActive` | `AccountCapabilitiesUSBankAccountACHPaymentsActive` |
    | Account | `AccountCapabilitiesUSBankAccountAchPaymentsInactive` | `AccountCapabilitiesUSBankAccountACHPaymentsInactive` |
    | Account | `AccountCapabilitiesUSBankAccountAchPaymentsPending` | `AccountCapabilitiesUSBankAccountACHPaymentsPending` |
    | Account | `AccountCapabilitiesUSBankAccountAchPaymentsParams` | `AccountCapabilitiesUSBankAccountACHPaymentsParams` |
    | Account | `AccountCapabilitiesParams.USBankAccountAchPayments` | `AccountCapabilitiesParams.USBankAccountACHPayments` |
    | Account | `AccountCapabilities.USBankAccountAchPayments` | `AccountCapabilities.USBankAccountACHPayments` |
    | Charge | `ChargePaymentMethodDetailsTypeAchCreditTransfer` | `ChargePaymentMethodDetailsTypeACHCreditTransfer` |
    | Charge | `ChargePaymentMethodDetailsTypeAchDebit` | `ChargePaymentMethodDetailsTypeACHDebit` |
    | Invoice | `InvoicePaymentSettingsPaymentMethodTypeAchCreditTransfer` | `InvoicePaymentSettingsPaymentMethodTypeACHhCreditTransfer` |
    | Invoice | `InvoicePaymentSettingsPaymentMethodTypeAchDebit` | `InvoicePaymentSettingsPaymentMethodTypeACHDebit` |
  
  - `Acss` to `ACSS`
    | Resource | Old name | New name |
    | --- | --- | --- |
    | FundingInstructions | `FundingInstructionsBankTransferFinancialAddressIban.Bic` | `FundingInstructionsBankTransferFinancialAddressIBAN.BIC` |
  
  - `Amex` to `AmEx`
    | Resource | Old name | New name |
    | --- | --- | --- |
    | Charge | `ChargePaymentMethodDetailsCardWallet.AmexExpressCheckout` | `ChargePaymentMethodDetailsCardWallet.AmExExpressCheckout` |
    | Charge | `ChargePaymentMethodDetailsCardWalletAmexExpressCheckout` | `ChargePaymentMethodDetailsCardWalletAmExExpressCheckout` |

  - `Bic` to `BIC`
    | Resource | Old name | New name |
    | --- | --- | --- |
    | Charge | `ChargePaymentMethodDetailsTypeAcssDebit` | `ChargePaymentMethodDetailsTypeACSSDebit` |
  
  - `Eps` to `EPS`
    | Resource | Old name | New name |
    | --- | --- | --- |
    | Charge | `ChargePaymentMethodDetailsTypeEps` | `ChargePaymentMethodDetailsTypeEPS` |

  - `FEDEX` to `FedEx`
    | Resource | Old name | New name |
    | --- | --- | --- |
    | Issuing Card | `IssuingCardShippingCarrierFEDEX` | `IssuingCardShippingCarrierFedEx` |
  
  - `Iban` to `IBAN`
    | Resource | Old name | New name |
    | --- | --- | --- |
    | FundingInstructions | `FundingInstructionsBankTransferFinancialAddress.Iban` | `FundingInstructionsBankTransferFinancialAddress.IBAN` |
    | FundingInstructions | `FundingInstructionsBankTransferFinancialAddressIban` | `FundingInstructionsBankTransferFinancialAddressIBAN` |
    | FundingInstructions | `FundingInstructionsBankTransferFinancialAddressIban.Iban` | `FundingInstructionsBankTransferFinancialAddressIBAN.IBAN` |
    | FundingInstructions | `FundingInstructionsBankTransferFinancialAddressTypeIban` | `FundingInstructionsBankTransferFinancialAddressTypeIBAN` |
  
  - `Ideal` to `IDEAL`
    | Resource | Old name | New name |
    | --- | --- | --- |
    | Account | `AccountCapabilitiesIdealPaymentsParams` | `AccountCapabilitiesIDEALPaymentsParams` |
    | Account | `AccountCapabilitiesParams.IdealPayments` | `AccountCapabilitiesParams.IDEALPayments` |
    | Account | `AccountCapabilities.IdealPayments` | `AccountCapabilities.IDEALPayments` |
    | Charge | `ChargePaymentMethodDetailsTypeIdeal` | `ChargePaymentMethodDetailsTypeIDEAL` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptions.Ideal` | `CheckoutSessionPaymentMethodOptions.IDEAL` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptionsIdeal` | `CheckoutSessionPaymentMethodOptionsIDEAL` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptionsIdealSetupFutureUsage` | `CheckoutSessionPaymentMethodOptionsIDEALSetupFutureUsage` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptionsIdealSetupFutureUsageNone` | `CheckoutSessionPaymentMethodOptionsIDEALSetupFutureUsageNone` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptionsIdealParams` | `CheckoutSessionPaymentMethodOptionsIDEALParams` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptionsParams.Ideal` | `CheckoutSessionPaymentMethodOptionsParams.IDEAL` |
    | Invoice | `InvoicePaymentSettingsPaymentMethodTypeIdeal` | `InvoicePaymentSettingsPaymentMethodTypeIDEAL` |
  
  - `Sepa` to `SEPA`
    | Resource | Old name | New name |
    | --- | --- | --- |
    | Charge | `ChargePaymentMethodDetailsTypeSepaDebit` | `ChargePaymentMethodDetailsTypeSEPAaDebit` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptions.SepaDebit` | `CheckoutSessionPaymentMethodOptions.SEPADebit` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptionsSepaDebit` | `CheckoutSessionPaymentMethodOptionsSEPADebit` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptionsSepaDebitSetupFutureUsage` | `CheckoutSessionPaymentMethodOptionsSEPADebitSetupFutureUsage` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptionsSepaDebitSetupFutureUsageNone` | `CheckoutSessionPaymentMethodOptionsSEPADebitSetupFutureUsageNone` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptionsSepaDebitSetupFutureUsageOffSession` | `CheckoutSessionPaymentMethodOptionsSEPADebitSetupFutureUsageOffSession` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptionsSepaDebitSetupFutureUsageOnSession` | `CheckoutSessionPaymentMethodOptionsSEPADebitSetupFutureUsageOnSession` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptionsSepaDebitParams` | `CheckoutSessionPaymentMethodOptionsSEPADebitParams` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptionsParams.SepaDebit` | `CheckoutSessionPaymentMethodOptionsParams.SEPADebit` |
    | FundingInstructions | `FundingInstructionsBankTransferFinancialAddressSupportedNetworkSepa` | `FundingInstructionsBankTransferFinancialAddressSupportedNetworkSEPA` |
    | Invoice | `InvoicePaymentSettingsPaymentMethodTypeSepaCreditTransfer` | `InvoicePaymentSettingsPaymentMethodTypeSEPACreditTransfer` |
    | Invoice | `InvoicePaymentSettingsPaymentMethodTypeSepaDebit` | `InvoicePaymentSettingsPaymentMethodTypeSEPADebit` |

  - `ExternalAccount` to `AccountExternalAccount`
    | Resource | Old name | New name |
    | --- | --- | --- |
    | Account | `ExternalAccount` | `AccountExternalAccount` |
    | Account | `ExternalAccountType` | `AccountExternalAccountType` |
    | Account | `ExternalAccountTypeBankAccount` | `AccountExternalAccountTypeBankAccount` |
    | Account | `ExternalAccountTypeCard` | `AccountExternalAccountTypeCard` |
    | Account | `AccountExternalAccountParams` | `AccountExternalAccountParams` |
    | Account | `ExternalAccountList` | `AccountExternalAccountList` |

  - `InvoiceLine` to `InvoiceLineItem`
    | Resource | Old name | New name |
    | --- | --- | --- |
    | Invoice | `InvoiceLine` | `InvoiceLineItem` |
    | Invoice | `InvoiceLineList` | `InvoiceLineItemList` |
    | Invoice | `InvoiceLineType` | `InvoiceLineItemType` |
    | Invoice | `InvoiceLineTypeInvoiceItem` | `InvoiceLineItemTypeInvoiceItem` |
    | Invoice | `InvoiceLineTypeSubscription` | `InvoiceLinItemeTypeSubscription` |
    | Invoice | `InvoiceLineDiscountAmount` | `InvoiceLineItemDiscountAmount` |
    | Invoice | `InvoiceLineProrationDetails` | `InvoiceLineItemProrationDetails` |
    | Invoice | `InvoiceLineProrationDetailsCreditedItems` | `InvoiceLineItemProrationDetailsCreditedItems` |

  - Other
    | Resource | Old name | New name |
    | --- | --- | --- |
    | Account | `PayoutInterval` | `AccountSettingsPayoutsScheduleInterval` |
    | Account | `PayoutIntervalDaily` | `AccountSettingsPayoutsScheduleIntervalDaily` |
    | Account | `PayoutIntervalManual` | `AccountSettingsPayoutsScheduleIntervalManual` |
    | Account | `PayoutIntervalMonthly` | `AccountSettingsPayoutsScheduleIntervalMonthly` |
    | Account | `PayoutIntervalWeekly` | `AccountSettingsPayoutsScheduleIntervalWeekly` |
    | Account | `AccountDocumentsParams.CompanyMemorandumOfAssocation` | `AccountDocumentsParams.CompanyMemorandumOfAssociation` |
    | Account | `AccountDeclineSettingsParams` | `AccountSettingsCardPaymentsDeclineOnParams` |
    | Account | `PayoutScheduleParams` | `AccountSettingsPayoutsScheduleParams` |
    | Account | `AccountPayoutSchedule` | `AccountSettingsPayoutsSchedule` |
    | Balance | `Amount.Value` | `Amount.Amount` |
    | Balance | `BalanceDetails` | `BalanceIssuing` |
    | BalanceTransaction | `BalanceTransactionSourceTypeReversal` | `BalanceTransactionSourceTypeTransferReversal` (this is a fix, the previous value of the enum was incorrect) |
    | BalanceTransaction | `BalanceTransactionTypeIssuingAuthorizationDispute` | `BalanceTransactionTypeIssuingDispute` |
    | BalanceTransaction | `BalanceTransactionTypeIssuingAuthorizationTransaction` | `BalanceTransactionTypeIssuingTransaction` |
    | BalanceTransaction | `BalanceTransactionFee` | `BalanceTransactionFeeDetail` |
    | BalanceTransaction | `BalanceTransactionSource.Reversal` | `BalanceTransactionSource.TransferReversal` |
    | Card | `CardBrandAmex` | `CardBrandAmericanExpress` |
    | Card | `TokenizationMethodAndroidPay` | `CardTokenizationMethodAndroidPay` |
    | Card | `TokenizationMethodApplePay` | `CardTokenizationMethodApplePay` |
    | Charge | `DestinationParams` | `ChargeDestinationParams` |
    | Charge | `ChargeLevel3LineItemsParams` | `ChargeLevel3LineItemParams` |
    | Charge | `FraudDetailsParams` | `ChargeFraudDetailsParams` |
    | Charge | `CaptureParams` | `ChargeCaptureParams` |
    | Charge | `ChargeTransferDataParams` | `ChargeCaptureTransferDataParams` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDs` | `CheckoutSessionCustomerDetailsTaxID` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsType` | `CheckoutSessionCustomerDetailsTaxIDType` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeAETRN` | `CheckoutSessionCustomerDetailsTaxIDTypeAETRN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeAUABN` | `CheckoutSessionCustomerDetailsTaxIDTypeAUABN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeAUARN` | `CheckoutSessionCustomerDetailsTaxIDTypeAUARN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeBGUIC` | `CheckoutSessionCustomerDetailsTaxIDTypeBGUIC` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeBRCNPJ` | `CheckoutSessionCustomerDetailsTaxIDTypeBRCNPJ` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeBRCPF` | `CheckoutSessionCustomerDetailsTaxIDTypeBRCPF` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeCABN` | `CheckoutSessionCustomerDetailsTaxIDTypeCABN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeCAGSTHST` | `CheckoutSessionCustomerDetailsTaxIDTypeCAGSTHST` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeCAPSTBC` | `CheckoutSessionCustomerDetailsTaxIDTypeCAPSTBC` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeCAPSTMB` | `CheckoutSessionCustomerDetailsTaxIDTypeCAPSTMB` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeCAPSTSK` | `CheckoutSessionCustomerDetailsTaxIDTypeCAPSTSK` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeCAQST` | `CheckoutSessionCustomerDetailsTaxIDTypeCAQST` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeCHVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeCHVAT` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeCLTIN` | `CheckoutSessionCustomerDetailsTaxIDTypeCLTIN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeESCIF` | `CheckoutSessionCustomerDetailsTaxIDTypeESCIF` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeEUOSSVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeEUOSSVAT` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeEUVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeEUVAT` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeGBVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeGBVAT` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeGEVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeGEVAT` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeHKBR` | `CheckoutSessionCustomerDetailsTaxIDTypeHKBR` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeHUTIN` | `CheckoutSessionCustomerDetailsTaxIDTypeHUTIN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeIDNPWP` | `CheckoutSessionCustomerDetailsTaxIDTypeIDNPWP` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeILVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeILVAT` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeINGST` | `CheckoutSessionCustomerDetailsTaxIDTypeINGST` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeISVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeISVAT` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeJPCN` | `CheckoutSessionCustomerDetailsTaxIDTypeJPCN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeJPRN` | `CheckoutSessionCustomerDetailsTaxIDTypeJPRN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeKRBRN` | `CheckoutSessionCustomerDetailsTaxIDTypeKRBRN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeLIUID` | `CheckoutSessionCustomerDetailsTaxIDTypeLIUID` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeMXRFC` | `CheckoutSessionCustomerDetailsTaxIDTypeMXRFC` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeMYFRP` | `CheckoutSessionCustomerDetailsTaxIDTypeMYFRP` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeMYITN` | `CheckoutSessionCustomerDetailsTaxIDTypeMYITN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeMYSST` | `CheckoutSessionCustomerDetailsTaxIDTypeMYSST` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeNOVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeNOVAT` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeNZGST` | `CheckoutSessionCustomerDetailsTaxIDTypeNZGST` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeRUINN` | `CheckoutSessionCustomerDetailsTaxIDTypeRUINN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeRUKPP` | `CheckoutSessionCustomerDetailsTaxIDTypeRUKPP` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeSAVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeSAVAT` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeSGGST` | `CheckoutSessionCustomerDetailsTaxIDTypeSGGST` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeSGUEN` | `CheckoutSessionCustomerDetailsTaxIDTypeSGUEN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeSITIN` | `CheckoutSessionCustomerDetailsTaxIDTypeSITIN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeTHVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeTHVAT` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeTWVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeTWVAT` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeUAVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeUAVAT` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeUnknown` | `CheckoutSessionCustomerDetailsTaxIDTypeUnknown` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeUSEIN` | `CheckoutSessionCustomerDetailsTaxIDTypeUSEIN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeZAVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeZAVAT` |
    | CheckoutSessions | `CheckoutSessionSubscriptionDataItemsParams` | `CheckoutSessionSubscriptionDataItemParams` |
    | CreditNote | `CreditNoteVoidParams` | `CreditNoteVoidCreditNoteParams` |
    | CreditNote | `CreditNoteLineItemListPreviewParams` | `CreditNotePreviewLinesParams` |
    | CreditNote | `CreditNoteLineItemListParams` | `CreditNoteListLinesParams` |
    | Customer | `CustomerInvoiceCustomField` | `CustomerInvoiceSettingsCustomField`|
    | Customer | `CustomerInvoiceCustomFieldParams` | `Customer dInvoiceSettingsCustomFieldParams`|
    | Customer | `CustomerShippingDetails` | `CustomerShipping` |
    | Customer | `CustomerShippingDetailsParams` | `CustomerShippingParams` |
    | Dispute | `EvidenceDetails` | `DisputeEvidenceDetails` |
    | FeeRefund | `FeeRefundListParams.ApplicationFee` | `FeeRefundListParams.ID` |
    | Invoice | `InvoiceUpcomingCustomerDetailsShippingParams` | `InvoiceUpcomingLinesCustomerDetailsShippingParams` |
    | Invoice | `InvoiceUpcomingCustomerDetailsTaxParams` | `InvoiceUpcomingLinesCustomerDetailsTaxParams` |
    | Invoice | `InvoiceUpcomingCustomerDetailsTaxIDParams` | `InvoiceUpcomingLinesCustomerDetailsTaxIDParams` |
    | Invoice | `InvoiceUpcomingCustomerDetailsParams` | `InvoiceUpcomingLinesCustomerDetailsParams` |
    | Invoice | `InvoiceSendParams` | `InvoiceSendInvoiceParams` |
    | Invoice | `InvoiceVoidParams` | `InvoiceVoidInvoiceParams` |
    | Invoice | `InvoiceLineList` | `InvoiceLineItemList` |
    | Invoice | `InvoiceLineListParams` | `InvoiceListLinesParams` |
    | Invoice | `InvoiceLineListParams.ID` | `InvoiceListLinesParams.Invoice` |
    | Invoice | `InvoiceDiscountAmount` | `InvoiceTotalDiscountAmount` |
    | Invoice | `InvoiceTaxAmount` | `InvoiceTotalTaxAmount` |
    | Invoice | `Invoice.ThreasholdReason` | `Invoice.ThresholdReason` (closes #1244) |
    | Invoice | `GetNext` | `Upcoming` |
    | Issuing | `IssuingAuthorizationWalletType` | `IssuingAuthorizationWallet` |
    | Issuing | `IssuingAuthorizationWalletTypeApplePay` | `IssuingAuthorizationWalletApplePay` |
    | Issuing | `IssuingAuthorizationWalletTypeGooglePay` | `IssuingAuthorizationWalletGooglePay` |
    | Issuing | `IssuingAuthorizationWalletTypeSamsungPlay` | `IssuingAuthorizationWalletSamsungPlay` |

- Replace `AccountAddressParams` with `AccountCompanyAddressKanaParams` and `AccountCompanyAddressKanjiParams`
- Change type:
  | Resource | Target | Old type | New type | Rationale |
  | --- | --- | --- | --- | --- |
  | Account | `AccountCapabilities.AffirmPayments` | `AccountCapabilitiesAffirmPayments` | `AccountCapabilityStatus` | Be consistent about using `AccountCapabilityStatus` for all of `AccountCapabilities` (all the same enum values) |
  | Account | `AccountCapabilities.BankTransferPayments` | `AccountCapabilitiesBankTransferPayments` | `AccountCapabilityStatus` | Be consistent about using `AccountCapabilityStatus` for all of `AccountCapabilities` (all the same enum values) |
  | Account | `AccountCapabilities.KonbiniPayments` | `AccountCapabilitiesKonbiniPayments` | `AccountCapabilityStatus` | Be consistent about using `AccountCapabilityStatus` for all of `AccountCapabilities` (all the same enum values) |
  | Account | `AccountCapabilities.LinkPayments` | `AccountCapabilitiesLinkPayments` | `AccountCapabilityStatus` | Be consistent about using `AccountCapabilityStatus` for all of `AccountCapabilities` (all the same enum values) |
  | Account | `AccountCapabilities.PayNowPayments` | `AccountCapabilitiesPayNowPayments` | `AccountCapabilityStatus` | Be consistent about using `AccountCapabilityStatus` for all of `AccountCapabilities` (all the same enum values) |
  | Account | `AccountCapabilities.PromptPayPayments` | `AccountCapabilitiesPromptPayPayments` | `AccountCapabilityStatus` | Be consistent about using `AccountCapabilityStatus` for all of `AccountCapabilities` (all the same enum values) |
  | Account | `AccountCapabilities.Treasury` | `AccountCapabilitiesTreasury` | `AccountCapabilityStatus` | Be consistent about using `AccountCapabilityStatus` for all of `AccountCapabilities` (all the same enum values) |
  | Account | `AccountCapabilities.USBankAccountAchPayments` | `AccountCapabilitiesUSBankAccountAchPayments` | `AccountCapabilityStatus` | Be consistent about using `AccountCapabilityStatus` for all of `AccountCapabilities` (all the same enum values) |
  | Account | `AccountCompanyParams.Address` | `*AccountAddressParams` | `*AddressParams` | `*AccountAddressParams` had `Town` field that is not applicable to `AccountCompanyParams.Address` |
  | Account | `AccountCompanyParams.AddressKana` | `*AccountAddressParams` | `*AccountCompanyAddressKanaParams` | `AccountAddressParams` has been split into separate `AccountCompanyAddressKanaParams` and `AccountCompanyAddressKanjiParams` |
  | Account | `AccountCompanyParams.AddressKanji` | `*AccountAddressParams` | `*AccountCompanyAddressKanjiParams` | `AccountAddressParams` has been split into separate `AccountCompanyAddressKanaParams` and `AccountCompanyAddressKanjiParams` |
  | Account | `AccountSettingsCardIssuingParams.TOSAcceptance` | `*AccountTOSAcceptanceParams` | `*AccountSettingsCardIssuingTOSAcceptanceParams` | `AccountTOSAcceptanceParams` has `ServiceAgreement` field that is not applicable to `AccountSettingsCardIssuingParams.TOSAcceptance` |
  | ApplicationFee | `ApplicationFee.Application` | `string` | `*Application` | `Application` is expandable |
  | BalanceTransaction | `BalanceTransactionSource.IssuingTransaction` | `*IssuingAuthorization` | `*IssuingTransaction` | `ApplicationFee.Application` is expandable |
  | BillingPortalConfiguration | `BillingPortalConfiguration.Application` | `string` | `*Application` | `Application` is expandable |
  | Card | `Card.AddressLine1Check` | `CardVerification` | `CardAddressLine1Check` | Same values. `CardVerification` was replaced with separate definitions for `AddressLine1Check`, `AddressZipCheck`, and `CVCCheck` |
  | Card | `Card.AddressZipCheck` | `CardVerification` | `CardAddressZipCheck` | Same values. `CardVerification` was replaced with separate definitions for `AddressLine1Check`, `AddressZipCheck`, and `CVCCheck` |
  | Card | `Card.CVCCheck` | `CardVerification` | `CardCVCCheck` | Same values. `CardVerification` was replaced with separate definitions for `AddressLine1Check`, `AddressZipCheck`, and `CVCCheck` |
  | Card | `Card.ExpMonth` | `uint8` | `int64` | Be consistent with `int` types in library |
  | Card | `Card.ExpYear` | `uint16` | `int64` | Be consistent with `int` types in library |
  | Charge | `Charge.BillingDetails` | `*BillingDetails` | `ChargeBillingDetails` | Unshared with PaymentMethod |
  | Charge | `Charge.Level3` | `ChargeLevel3` | `*ChargeLevel3` | `ChargeLevel3` is expandable |
  | Charge | `ChargePaymentMethodDetailsCardChecks.AddressLine1Check` | `CardVerification` | `ChargePaymentMethodDetailsCardChecksAddressLine1Check` | Same values. `CardVerification` was replaced with resource- and field-specific enums (with the same values) |
  | Charge | `ChargePaymentMethodDetailsCardChecks.AddressPostalCodeCheck` | `CardVerification` | `ChargePaymentMethodDetailsCardChecksAddressPostalCodeCheck` | Same values. `CardVerification` was replaced with resource- and field-specific enums (with the same values) |
  | Charge | `ChargePaymentMethodDetailsCardChecks.CVCCheck` | `CardVerification` | `ChargePaymentMethodDetailsCardChecksCVCCheck` | Same values. `CardVerification` was replaced with resource- and field-specific enums (with the same values) |
  | Customer | `Customer.Address` | `Address` | `*Address` | `Address` is nullable |
  | Customer | `CustomerParams.Source` | `*SourceParams` | `*string` | `CustomerParams.Source` should be the ID of a source |
  | Discount | `Discount.CheckoutSession` | `*CheckoutSession` | `string` | `discount.checkout_session` is not expandable |
  | Discount | `Discount.Customer` | `string` | `*Customer` | `discount.customer` is expandable |
  | Invoice | `InvoiceCustomerTaxID.Type` | `TaxIDType` | `*TaxIDType` | `invoice.customer_tax_ids.type` is nullable |
  | Invoice | `Invoice.Type` | `*CollectionMethod` | `CollectionMethod` | `invoice.collection_method` is not nullable |
  | Invoice | `Invoice.CustomerName` | `*string` | `string` | `invoice.customer_name` is not nullable |
  | Invoice | `Invoice.CustomerPhone` | `*string` | `string` | `invoice.customer_phone` is not nullable |
  | Invoice | `Invoice.CustomerTaxExempt` | `CustomerTaxExempt` | `*CustomerTaxExempt` | `invoice.customer_tax_exempt` is nullable |
  | Invoice | `Invoice.StatusTransitions` | `InvoiceStatusTransitions` | `*InvoiceStatusTransitions` | `invoice.status_transitions` is nullable |
  | Issuing Card | `IssuingCardShippingParams.Name` | `string` | `*string` | Params fields should be pointers |

- Moved `BalanceTransaction` iterator from `balance.go` to `balancetransaction.go`
- Fixed `BalanceTransactionSource` `UnmarshalJSON` for when `BalanceTransactionSource.Type == "transfer_reversal"` (previously, we were checking if `Type == "reversal"`, which was always false)
- For BankAccount and Card client methods, check that exactly one of `params.Account` and `params.Customer` is set (previously they could both be set, but only one would be used, and it was different between BankAccount and Card)
- Replace `CardVerification` with field-specific enums (with the same values)
- Move `Del` from `discount/client.go` to `customer/client.go` and rename to `DeleteDiscount`
- Move `DelSub` from `discount/client.go` to `subscription/client.go` and rename to `DeleteDiscount`
- Add separate parameter struct for CreditNote `ListPreviewLines` (renamed to `PreviewLines`) method (`[CreditNoteLineItemListPreviewParams -> CreditNotePreviewParams].Lines` `CreditNoteLineParams` -> `CreditNotePreviewLineParams`)
- Replace `FeeRefundParams.ApplicationFee` with `FeeRefundParams.Fee` and `FeeRefundParams.ID`
- Add separate parameter struct for Invoice `GetNext` (renamed to `Upcoming`) method (`InvoiceUpcomingParams`, and nested params `InvoiceUpcomingLinesInvoiceItemPriceDataParams`, `InvoiceUpcomingLinesInvoiceItemDiscountParams`, `InvoiceUpcomingLinesDiscountParams`, `InvoiceUpcomingLinesInvoiceItemPeriodParams`). `Upcoming`-only fields `Coupon`, `CustomerDetails`, `InvoiceItems`, `Subscription`, `SubscriptionBillingCycleAnchor`, `Schedule`, `SubscriptionBillingCycleAnchor`, `SubscriptionBillingCycleAnchorNow`, `SubscriptionBillingCycleAnchorUnchanged`, `SubscriptionCancelAt`, `SubscriptionCancelAtPeriodEnd`, `SubscriptionCancelNow`, `SubscriptionDefaultTaxRates`, `SubscriptionItems`, `SubscriptionProrationBehavior`, `SubscriptionProrationDate`, `SubscriptionStartDate`, `SubscriptionTrialEnd`, `SubscriptionTrialEndNow`, and `SubscriptionTrialFromPlan` are removed from `InvoiceParams`.

## Deprecated

- SKU TODO

## Removed
- `UnmarshalJSON` for resources that are not expandable: `BillingPortalSession`, `Capability`, `CheckoutSession`, `FileLink`, `InvoiceItem`
- `AccountRejectReason` (was only referenced in `account/client_test.go`, actual `AccountRejectParams.Reason` is `*string`)
- `AccountParams.RequestedCapabilities` (use Capabilities instead: https://stripe.com/docs/connect/account-capabilities)

  ```go
  // Before
  params := &stripe.AccountParams{
    RequestedCapabilities: ["card_payments"],
  }

  // After
  params := &stripe.AccountParams{
    Capabilities: &stripe.AccountCapabilitiesParams{
        CardPayments: &stripe.AccountCapabilitiesCardPaymentsParams{
            Requested: stripe.Bool(true),
        },
    },
  }
  ```

- ``AccountSettingsParams.Dashboard` and `AccountSettingsDashboardParams` (Note: `Dashboard` are still available on `AccountSettings`, but it's not available as parameters for any of the methods)
- `AccountCompany.RegistrationNumber` (Note: `RegistrationNumber` is still available on `AccountCompanyParams`, but is not returned in the response)
- `BalanceTransactionStatus`. It was meant to be an enum, but none of the enum values were defined, so it was just an alias for string.
- `CardParams.AccountType`. `AccountType` does not exist on any client method for Card. It does on BankAccount, which is similar.
- `id` param from CheckoutSessions `ListLineItems`. Use `CheckoutSessionListLineItemsParams.Session` instead.

  ```go
  // Before
  params := &stripe.CheckoutSessionListLineItemsParams{}
  i := ListLineItems("cs_123", params)

  // After
  params := &stripe.CheckoutSessionListLineItemsParams{Session: stripe.String("cs_123")}
  i := ListLineItems(params)
  ```

- `CheckoutSessionLineItemPriceDataRecurringParams.AggregateUsage`, `CheckoutSessionLineItemPriceDataRecurringParams.TrialPeriodDays`, and `CheckoutSessionLineItemPriceDataRecurringParams.UsageType`
- `CheckoutSessionPaymentIntentDataParams.Params`, `CheckoutSessionSetupIntentDataParams.Params`, `CheckoutSessionSubscriptionDataParams.Params`- TODO not top level
- `CheckoutSessionTotalDetailsBreakdownTax.TaxRate`. Use `CheckoutSessionTotalDetailsBreakdownTax.Rate`
- `CheckoutSessionTotalDetailsBreakdownTax.Deleted`. TODO
- `CustomerParams.Token`
- `Discount` `APIResource` embed
- `FilePurposeFoundersStockDocument` (`"founders_stock_document"` option for `File.Purpose`)
- `InvoiceParams.Paid`. Use `invoice.status` to check for status. `invoice.status` is a read-only field.
- `InvoiceParams.SubscriptionPlan` and `InvoiceParams.SubscriptionQuantity` (note: these would have been on `InvoiceUpcomingParams`)
- `InvoiceListLinesParams.Customer` and `InvoiceListLinesParams.Subscription` (these are not available for Invoice `ListLines`, but are available for `List`)
- `IssuingAuthorizationRequestHistoryViolatedAuthorizationControlEntity` and `IssuingAuthorizationRequestHistoryViolatedAuthorizationControlName` (unused enums)
- `IssuingCardSpendingControlsParams.SpendingLimitsCurrency`. `issuing_card` has `currency`, and `issuing_card.spending_controls.spending_limits.amount` will use that currency