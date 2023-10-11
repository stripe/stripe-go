//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The channel through which the applicant has submitted their application.
type IssuingCreditUnderwritingRecordApplicationApplicationMethod string

// List of values that IssuingCreditUnderwritingRecordApplicationApplicationMethod can take
const (
	IssuingCreditUnderwritingRecordApplicationApplicationMethodInPerson IssuingCreditUnderwritingRecordApplicationApplicationMethod = "in_person"
	IssuingCreditUnderwritingRecordApplicationApplicationMethodMail     IssuingCreditUnderwritingRecordApplicationApplicationMethod = "mail"
	IssuingCreditUnderwritingRecordApplicationApplicationMethodOnline   IssuingCreditUnderwritingRecordApplicationApplicationMethod = "online"
	IssuingCreditUnderwritingRecordApplicationApplicationMethodPhone    IssuingCreditUnderwritingRecordApplicationApplicationMethod = "phone"
)

// Scope of demand made by the applicant.
type IssuingCreditUnderwritingRecordApplicationPurpose string

// List of values that IssuingCreditUnderwritingRecordApplicationPurpose can take
const (
	IssuingCreditUnderwritingRecordApplicationPurposeCreditLimitIncrease IssuingCreditUnderwritingRecordApplicationPurpose = "credit_limit_increase"
	IssuingCreditUnderwritingRecordApplicationPurposeCreditLineOpening   IssuingCreditUnderwritingRecordApplicationPurpose = "credit_line_opening"
)

// The event that triggered the underwriting.
type IssuingCreditUnderwritingRecordCreatedFrom string

// List of values that IssuingCreditUnderwritingRecordCreatedFrom can take
const (
	IssuingCreditUnderwritingRecordCreatedFromApplication     IssuingCreditUnderwritingRecordCreatedFrom = "application"
	IssuingCreditUnderwritingRecordCreatedFromProactiveReview IssuingCreditUnderwritingRecordCreatedFrom = "proactive_review"
)

// List of reasons why the application was rejected up to 4 reasons, in order of importance.
type IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason string

// List of values that IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason can take
const (
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonApplicantTooYoung                      IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "applicant_too_young"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonApplicationIsNotBeneficialOwner        IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "application_is_not_beneficial_owner"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonBankruptcy                             IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "bankruptcy"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonBusinessSizeTooSmall                   IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "business_size_too_small"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonChangeInFinancialState                 IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "change_in_financial_state"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonChangeInUtilizationOfCreditLine        IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "change_in_utilization_of_credit_line"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonCustomerAlreadyExists                  IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "customer_already_exists"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonDebtToCashBalanceRatioTooHigh          IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "debt_to_cash_balance_ratio_too_high"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonDebtToEquityRatioTooHigh               IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "debt_to_equity_ratio_too_high"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonDecreaseInIncomeToExpenseRatio         IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "decrease_in_income_to_expense_ratio"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonDecreaseInSocialMediaPerformance       IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "decrease_in_social_media_performance"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonDelinquentCreditObligations            IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "delinquent_credit_obligations"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonDurationOfResidence                    IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "duration_of_residence"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonExceedsAcceptablePlatformExposure      IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "exceeds_acceptable_platform_exposure"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonExcessiveIncomeOrRevenueObligations    IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "excessive_income_or_revenue_obligations"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonExpensesToCashBalanceRatioTooHigh      IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "expenses_to_cash_balance_ratio_too_high"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonForeclosureOrRepossession              IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "foreclosure_or_repossession"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonFrozenFileAtCreditBureau               IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "frozen_file_at_credit_bureau"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonGarnishmentOrAttachment                IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "garnishment_or_attachment"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonGovernmentLoanProgramCriteria          IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "government_loan_program_criteria"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonHasRecentCreditLimitIncrease           IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "has_recent_credit_limit_increase"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonHighConcentrationOfClients             IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "high_concentration_of_clients"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonIncompleteApplication                  IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "incomplete_application"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonInconsistentMonthlyRevenues            IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "inconsistent_monthly_revenues"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonInsufficientAccountHistoryWithPlatform IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "insufficient_account_history_with_platform"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonInsufficientBankAccountHistory         IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "insufficient_bank_account_history"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonInsufficientCashBalance                IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "insufficient_cash_balance"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonInsufficientCashFlow                   IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "insufficient_cash_flow"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonInsufficientCollateral                 IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "insufficient_collateral"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonInsufficientCreditExperience           IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "insufficient_credit_experience"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonInsufficientCreditUtilization          IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "insufficient_credit_utilization"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonInsufficientDeposits                   IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "insufficient_deposits"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonInsufficientIncome                     IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "insufficient_income"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonInsufficientPeriodInOperation          IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "insufficient_period_in_operation"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonInsufficientRevenue                    IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "insufficient_revenue"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonInsufficientSocialMediaPerformance     IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "insufficient_social_media_performance"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonInsufficientTradeCreditInsurance       IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "insufficient_trade_credit_insurance"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonInsufficientUsageAsQualifiedExpenses   IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "insufficient_usage_as_qualified_expenses"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonLatePaymentHistoryReportedToBureau     IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "late_payment_history_reported_to_bureau"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonLienCollectionActionOrJudgement        IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "lien_collection_action_or_judgement"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonNegativePublicInformation              IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "negative_public_information"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonNoCreditFile                           IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "no_credit_file"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonOther                                  IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "other"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonOutsideSupportedCountry                IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "outside_supported_country"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonOutsideSupportedState                  IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "outside_supported_state"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonPoorPaymentHistoryWithPlatform         IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "poor_payment_history_with_platform"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonPriorOrCurrentLegalAction              IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "prior_or_current_legal_action"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonProhibitedIndustry                     IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "prohibited_industry"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonRateOfCashBalanceFluctuationTooHigh    IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "rate_of_cash_balance_fluctuation_too_high"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonRecentInquiriesOnBusinessCreditReport  IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "recent_inquiries_on_business_credit_report"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonRemovalOfBankAccountConnection         IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "removal_of_bank_account_connection"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonRevenueDiscrepancy                     IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "revenue_discrepancy"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonRunwayTooShort                         IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "runway_too_short"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonSuspectedFraud                         IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "suspected_fraud"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonTooManyNonSufficientFundsOrOverdrafts  IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "too_many_non_sufficient_funds_or_overdrafts"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonUnableToVerifyAddress                  IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "unable_to_verify_address"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonUnableToVerifyIdentity                 IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "unable_to_verify_identity"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonUnableToVerifyIncomeOrRevenue          IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "unable_to_verify_income_or_revenue"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonUnprofitable                           IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "unprofitable"
	IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasonUnsupportableBusinessType              IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason = "unsupportable_business_type"
)

// List of reasons why the existing credit was decreased, up to 4 reasons, in order of importance.
type IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason string

// List of values that IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason can take
const (
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonApplicantTooYoung                      IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "applicant_too_young"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonApplicationIsNotBeneficialOwner        IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "application_is_not_beneficial_owner"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonBankruptcy                             IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "bankruptcy"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonBusinessSizeTooSmall                   IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "business_size_too_small"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonChangeInFinancialState                 IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "change_in_financial_state"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonChangeInUtilizationOfCreditLine        IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "change_in_utilization_of_credit_line"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonCustomerAlreadyExists                  IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "customer_already_exists"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonDebtToCashBalanceRatioTooHigh          IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "debt_to_cash_balance_ratio_too_high"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonDebtToEquityRatioTooHigh               IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "debt_to_equity_ratio_too_high"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonDecreaseInIncomeToExpenseRatio         IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "decrease_in_income_to_expense_ratio"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonDecreaseInSocialMediaPerformance       IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "decrease_in_social_media_performance"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonDelinquentCreditObligations            IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "delinquent_credit_obligations"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonDurationOfResidence                    IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "duration_of_residence"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonExceedsAcceptablePlatformExposure      IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "exceeds_acceptable_platform_exposure"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonExcessiveIncomeOrRevenueObligations    IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "excessive_income_or_revenue_obligations"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonExpensesToCashBalanceRatioTooHigh      IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "expenses_to_cash_balance_ratio_too_high"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonForeclosureOrRepossession              IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "foreclosure_or_repossession"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonFrozenFileAtCreditBureau               IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "frozen_file_at_credit_bureau"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonGarnishmentOrAttachment                IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "garnishment_or_attachment"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonGovernmentLoanProgramCriteria          IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "government_loan_program_criteria"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonHasRecentCreditLimitIncrease           IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "has_recent_credit_limit_increase"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonHighConcentrationOfClients             IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "high_concentration_of_clients"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonIncompleteApplication                  IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "incomplete_application"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonInconsistentMonthlyRevenues            IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "inconsistent_monthly_revenues"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonInsufficientAccountHistoryWithPlatform IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "insufficient_account_history_with_platform"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonInsufficientBankAccountHistory         IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "insufficient_bank_account_history"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonInsufficientCashBalance                IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "insufficient_cash_balance"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonInsufficientCashFlow                   IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "insufficient_cash_flow"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonInsufficientCollateral                 IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "insufficient_collateral"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonInsufficientCreditExperience           IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "insufficient_credit_experience"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonInsufficientCreditUtilization          IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "insufficient_credit_utilization"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonInsufficientDeposits                   IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "insufficient_deposits"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonInsufficientIncome                     IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "insufficient_income"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonInsufficientPeriodInOperation          IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "insufficient_period_in_operation"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonInsufficientRevenue                    IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "insufficient_revenue"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonInsufficientSocialMediaPerformance     IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "insufficient_social_media_performance"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonInsufficientTradeCreditInsurance       IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "insufficient_trade_credit_insurance"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonInsufficientUsageAsQualifiedExpenses   IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "insufficient_usage_as_qualified_expenses"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonLatePaymentHistoryReportedToBureau     IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "late_payment_history_reported_to_bureau"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonLienCollectionActionOrJudgement        IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "lien_collection_action_or_judgement"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonNegativePublicInformation              IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "negative_public_information"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonNoCreditFile                           IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "no_credit_file"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonOther                                  IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "other"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonOutsideSupportedCountry                IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "outside_supported_country"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonOutsideSupportedState                  IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "outside_supported_state"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonPoorPaymentHistoryWithPlatform         IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "poor_payment_history_with_platform"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonPriorOrCurrentLegalAction              IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "prior_or_current_legal_action"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonProhibitedIndustry                     IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "prohibited_industry"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonRateOfCashBalanceFluctuationTooHigh    IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "rate_of_cash_balance_fluctuation_too_high"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonRecentInquiriesOnBusinessCreditReport  IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "recent_inquiries_on_business_credit_report"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonRemovalOfBankAccountConnection         IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "removal_of_bank_account_connection"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonRevenueDiscrepancy                     IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "revenue_discrepancy"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonRunwayTooShort                         IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "runway_too_short"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonSuspectedFraud                         IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "suspected_fraud"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonTooManyNonSufficientFundsOrOverdrafts  IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "too_many_non_sufficient_funds_or_overdrafts"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonUnableToVerifyAddress                  IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "unable_to_verify_address"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonUnableToVerifyIdentity                 IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "unable_to_verify_identity"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonUnableToVerifyIncomeOrRevenue          IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "unable_to_verify_income_or_revenue"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonUnprofitable                           IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "unprofitable"
	IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasonUnsupportableBusinessType              IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason = "unsupportable_business_type"
)

// List of reasons why the existing account was closed, up to 4 reasons, in order of importance.
type IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason string

// List of values that IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason can take
const (
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonApplicantTooYoung                      IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "applicant_too_young"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonApplicationIsNotBeneficialOwner        IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "application_is_not_beneficial_owner"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonBankruptcy                             IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "bankruptcy"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonBusinessSizeTooSmall                   IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "business_size_too_small"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonChangeInFinancialState                 IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "change_in_financial_state"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonChangeInUtilizationOfCreditLine        IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "change_in_utilization_of_credit_line"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonCustomerAlreadyExists                  IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "customer_already_exists"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonDebtToCashBalanceRatioTooHigh          IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "debt_to_cash_balance_ratio_too_high"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonDebtToEquityRatioTooHigh               IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "debt_to_equity_ratio_too_high"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonDecreaseInIncomeToExpenseRatio         IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "decrease_in_income_to_expense_ratio"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonDecreaseInSocialMediaPerformance       IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "decrease_in_social_media_performance"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonDelinquentCreditObligations            IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "delinquent_credit_obligations"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonDurationOfResidence                    IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "duration_of_residence"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonExceedsAcceptablePlatformExposure      IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "exceeds_acceptable_platform_exposure"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonExcessiveIncomeOrRevenueObligations    IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "excessive_income_or_revenue_obligations"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonExpensesToCashBalanceRatioTooHigh      IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "expenses_to_cash_balance_ratio_too_high"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonForeclosureOrRepossession              IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "foreclosure_or_repossession"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonFrozenFileAtCreditBureau               IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "frozen_file_at_credit_bureau"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonGarnishmentOrAttachment                IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "garnishment_or_attachment"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonGovernmentLoanProgramCriteria          IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "government_loan_program_criteria"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonHasRecentCreditLimitIncrease           IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "has_recent_credit_limit_increase"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonHighConcentrationOfClients             IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "high_concentration_of_clients"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonIncompleteApplication                  IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "incomplete_application"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonInconsistentMonthlyRevenues            IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "inconsistent_monthly_revenues"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonInsufficientAccountHistoryWithPlatform IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "insufficient_account_history_with_platform"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonInsufficientBankAccountHistory         IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "insufficient_bank_account_history"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonInsufficientCashBalance                IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "insufficient_cash_balance"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonInsufficientCashFlow                   IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "insufficient_cash_flow"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonInsufficientCollateral                 IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "insufficient_collateral"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonInsufficientCreditExperience           IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "insufficient_credit_experience"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonInsufficientCreditUtilization          IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "insufficient_credit_utilization"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonInsufficientDeposits                   IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "insufficient_deposits"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonInsufficientIncome                     IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "insufficient_income"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonInsufficientPeriodInOperation          IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "insufficient_period_in_operation"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonInsufficientRevenue                    IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "insufficient_revenue"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonInsufficientSocialMediaPerformance     IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "insufficient_social_media_performance"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonInsufficientTradeCreditInsurance       IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "insufficient_trade_credit_insurance"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonInsufficientUsageAsQualifiedExpenses   IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "insufficient_usage_as_qualified_expenses"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonLatePaymentHistoryReportedToBureau     IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "late_payment_history_reported_to_bureau"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonLienCollectionActionOrJudgement        IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "lien_collection_action_or_judgement"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonNegativePublicInformation              IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "negative_public_information"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonNoCreditFile                           IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "no_credit_file"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonOther                                  IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "other"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonOutsideSupportedCountry                IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "outside_supported_country"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonOutsideSupportedState                  IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "outside_supported_state"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonPoorPaymentHistoryWithPlatform         IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "poor_payment_history_with_platform"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonPriorOrCurrentLegalAction              IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "prior_or_current_legal_action"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonProhibitedIndustry                     IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "prohibited_industry"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonRateOfCashBalanceFluctuationTooHigh    IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "rate_of_cash_balance_fluctuation_too_high"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonRecentInquiriesOnBusinessCreditReport  IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "recent_inquiries_on_business_credit_report"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonRemovalOfBankAccountConnection         IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "removal_of_bank_account_connection"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonRevenueDiscrepancy                     IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "revenue_discrepancy"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonRunwayTooShort                         IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "runway_too_short"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonSuspectedFraud                         IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "suspected_fraud"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonTooManyNonSufficientFundsOrOverdrafts  IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "too_many_non_sufficient_funds_or_overdrafts"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonUnableToVerifyAddress                  IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "unable_to_verify_address"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonUnableToVerifyIdentity                 IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "unable_to_verify_identity"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonUnableToVerifyIncomeOrRevenue          IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "unable_to_verify_income_or_revenue"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonUnprofitable                           IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "unprofitable"
	IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasonUnsupportableBusinessType              IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason = "unsupportable_business_type"
)

// Outcome of the decision.
type IssuingCreditUnderwritingRecordDecisionType string

// List of values that IssuingCreditUnderwritingRecordDecisionType can take
const (
	IssuingCreditUnderwritingRecordDecisionTypeAdditionalInformationRequested IssuingCreditUnderwritingRecordDecisionType = "additional_information_requested"
	IssuingCreditUnderwritingRecordDecisionTypeApplicationRejected            IssuingCreditUnderwritingRecordDecisionType = "application_rejected"
	IssuingCreditUnderwritingRecordDecisionTypeCreditLimitApproved            IssuingCreditUnderwritingRecordDecisionType = "credit_limit_approved"
	IssuingCreditUnderwritingRecordDecisionTypeCreditLimitDecreased           IssuingCreditUnderwritingRecordDecisionType = "credit_limit_decreased"
	IssuingCreditUnderwritingRecordDecisionTypeCreditLineClosed               IssuingCreditUnderwritingRecordDecisionType = "credit_line_closed"
	IssuingCreditUnderwritingRecordDecisionTypeNoChanges                      IssuingCreditUnderwritingRecordDecisionType = "no_changes"
	IssuingCreditUnderwritingRecordDecisionTypeWithdrawnByApplicant           IssuingCreditUnderwritingRecordDecisionType = "withdrawn_by_applicant"
)

// The decision before the exception was applied.
type IssuingCreditUnderwritingRecordUnderwritingExceptionOriginalDecisionType string

// List of values that IssuingCreditUnderwritingRecordUnderwritingExceptionOriginalDecisionType can take
const (
	IssuingCreditUnderwritingRecordUnderwritingExceptionOriginalDecisionTypeAdditionalInformationRequested IssuingCreditUnderwritingRecordUnderwritingExceptionOriginalDecisionType = "additional_information_requested"
	IssuingCreditUnderwritingRecordUnderwritingExceptionOriginalDecisionTypeApplicationRejected            IssuingCreditUnderwritingRecordUnderwritingExceptionOriginalDecisionType = "application_rejected"
	IssuingCreditUnderwritingRecordUnderwritingExceptionOriginalDecisionTypeCreditLimitApproved            IssuingCreditUnderwritingRecordUnderwritingExceptionOriginalDecisionType = "credit_limit_approved"
	IssuingCreditUnderwritingRecordUnderwritingExceptionOriginalDecisionTypeCreditLimitDecreased           IssuingCreditUnderwritingRecordUnderwritingExceptionOriginalDecisionType = "credit_limit_decreased"
	IssuingCreditUnderwritingRecordUnderwritingExceptionOriginalDecisionTypeCreditLineClosed               IssuingCreditUnderwritingRecordUnderwritingExceptionOriginalDecisionType = "credit_line_closed"
	IssuingCreditUnderwritingRecordUnderwritingExceptionOriginalDecisionTypeNoChanges                      IssuingCreditUnderwritingRecordUnderwritingExceptionOriginalDecisionType = "no_changes"
	IssuingCreditUnderwritingRecordUnderwritingExceptionOriginalDecisionTypeWithdrawnByApplicant           IssuingCreditUnderwritingRecordUnderwritingExceptionOriginalDecisionType = "withdrawn_by_applicant"
)

// Retrieves a CreditUnderwritingRecord object.
type IssuingCreditUnderwritingRecordParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *IssuingCreditUnderwritingRecordParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves a list of CreditUnderwritingRecord objects. The objects are sorted in descending order by creation date, with the most-recently-created object appearing first.
type IssuingCreditUnderwritingRecordListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *IssuingCreditUnderwritingRecordListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Details about the application submission.
type IssuingCreditUnderwritingRecordCreateFromApplicationApplicationParams struct {
	// The channel through which the applicant has submitted their application. Defaults to `online`.
	ApplicationMethod *string `form:"application_method"`
	// Scope of demand made by the applicant.
	Purpose *string `form:"purpose"`
	// Date when the applicant submitted their application.
	SubmittedAt *int64 `form:"submitted_at"`
}

// Information about the company or person applying or holding the account.
type IssuingCreditUnderwritingRecordCreateFromApplicationCreditUserParams struct {
	// Email of the applicant or accountholder.
	Email *string `form:"email"`
	// Full name of the company or person.
	Name *string `form:"name"`
}

// Creates a CreditUnderwritingRecord object with information about a credit application submission
type IssuingCreditUnderwritingRecordCreateFromApplicationParams struct {
	Params `form:"*"`
	// Details about the application submission.
	Application *IssuingCreditUnderwritingRecordCreateFromApplicationApplicationParams `form:"application"`
	// Information about the company or person applying or holding the account.
	CreditUser *IssuingCreditUnderwritingRecordCreateFromApplicationCreditUserParams `form:"credit_user"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
}

// AddExpand appends a new field to expand.
func (p *IssuingCreditUnderwritingRecordCreateFromApplicationParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *IssuingCreditUnderwritingRecordCreateFromApplicationParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Information about the company or person applying or holding the account.
type IssuingCreditUnderwritingRecordCreateFromProactiveReviewCreditUserParams struct {
	// Email of the applicant or accountholder.
	Email *string `form:"email"`
	// Full name of the company or person.
	Name *string `form:"name"`
}

// Details about the credit limit approved. An approved credit limit is required before you can set a `credit_limit_amount` in the [CreditPolicy API](https://stripe.com/docs/api/issuing/credit_policy/)
type IssuingCreditUnderwritingRecordCreateFromProactiveReviewDecisionCreditLimitApprovedParams struct {
	// The credit approved, in the currency of the account and [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
	Amount *int64 `form:"amount"`
	// The currency of the credit approved, will default to the Account's Issuing currency.
	Currency *string `form:"currency"`
}

// Details about the credit limit decreased.
type IssuingCreditUnderwritingRecordCreateFromProactiveReviewDecisionCreditLimitDecreasedParams struct {
	// The credit approved, in the currency of the account and [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
	Amount *int64 `form:"amount"`
	// The currency of the credit approved, will default to the Account's Issuing currency.
	Currency *string `form:"currency"`
	// Details about the `reasons.other` when present.
	ReasonOtherExplanation *string `form:"reason_other_explanation"`
	// List of reasons why the existing credit was decreased, up to 4 reasons, in order of importance.
	Reasons []*string `form:"reasons"`
}

// Details about the credit line closed.
type IssuingCreditUnderwritingRecordCreateFromProactiveReviewDecisionCreditLineClosedParams struct {
	// Details about the `reasons.other` when present.
	ReasonOtherExplanation *string `form:"reason_other_explanation"`
	// List of reasons why the credit line was closed, up to 4 reasons, in order of importance.
	Reasons []*string `form:"reasons"`
}

// Details about the decision.
type IssuingCreditUnderwritingRecordCreateFromProactiveReviewDecisionParams struct {
	// Details about the credit limit approved. An approved credit limit is required before you can set a `credit_limit_amount` in the [CreditPolicy API](https://stripe.com/docs/api/issuing/credit_policy/)
	CreditLimitApproved *IssuingCreditUnderwritingRecordCreateFromProactiveReviewDecisionCreditLimitApprovedParams `form:"credit_limit_approved"`
	// Details about the credit limit decreased.
	CreditLimitDecreased *IssuingCreditUnderwritingRecordCreateFromProactiveReviewDecisionCreditLimitDecreasedParams `form:"credit_limit_decreased"`
	// Details about the credit line closed.
	CreditLineClosed *IssuingCreditUnderwritingRecordCreateFromProactiveReviewDecisionCreditLineClosedParams `form:"credit_line_closed"`
	// Outcome of the decision.
	Type *string `form:"type"`
}

// If an exception to the usual underwriting criteria was made for this decision, details about the exception must be provided. Exceptions should only be granted in rare circumstances, in consultation with Stripe Compliance.
type IssuingCreditUnderwritingRecordCreateFromProactiveReviewUnderwritingExceptionParams struct {
	// Written explanation for the exception.
	Explanation *string `form:"explanation"`
	// The decision before the exception was applied.
	OriginalDecisionType *string `form:"original_decision_type"`
}

// Creates a CreditUnderwritingRecord object from an underwriting decision coming from a proactive review of an existing accountholder
type IssuingCreditUnderwritingRecordCreateFromProactiveReviewParams struct {
	Params `form:"*"`
	// Information about the company or person applying or holding the account.
	CreditUser *IssuingCreditUnderwritingRecordCreateFromProactiveReviewCreditUserParams `form:"credit_user"`
	// Date when a decision was made.
	DecidedAt *int64 `form:"decided_at"`
	// Details about the decision.
	Decision *IssuingCreditUnderwritingRecordCreateFromProactiveReviewDecisionParams `form:"decision"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// If an exception to the usual underwriting criteria was made for this decision, details about the exception must be provided. Exceptions should only be granted in rare circumstances, in consultation with Stripe Compliance.
	UnderwritingException *IssuingCreditUnderwritingRecordCreateFromProactiveReviewUnderwritingExceptionParams `form:"underwriting_exception"`
}

// AddExpand appends a new field to expand.
func (p *IssuingCreditUnderwritingRecordCreateFromProactiveReviewParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *IssuingCreditUnderwritingRecordCreateFromProactiveReviewParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Details about the application rejection.
type IssuingCreditUnderwritingRecordReportDecisionDecisionApplicationRejectedParams struct {
	// Details about the `reasons.other` when present.
	ReasonOtherExplanation *string `form:"reason_other_explanation"`
	// List of reasons why the application was rejected, up to 4 reasons, in order of importance.
	Reasons []*string `form:"reasons"`
}

// Details about the credit limit approved. An approved credit limit is required before you can set a `credit_limit_amount` in the [CreditPolicy API](https://stripe.com/docs/api/issuing/credit_policy/)
type IssuingCreditUnderwritingRecordReportDecisionDecisionCreditLimitApprovedParams struct {
	// The credit approved, in the currency of the account and [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
	Amount *int64 `form:"amount"`
	// The currency of the credit approved, will default to the Account's Issuing currency.
	Currency *string `form:"currency"`
}

// Details about the decision.
type IssuingCreditUnderwritingRecordReportDecisionDecisionParams struct {
	// Details about the application rejection.
	ApplicationRejected *IssuingCreditUnderwritingRecordReportDecisionDecisionApplicationRejectedParams `form:"application_rejected"`
	// Details about the credit limit approved. An approved credit limit is required before you can set a `credit_limit_amount` in the [CreditPolicy API](https://stripe.com/docs/api/issuing/credit_policy/)
	CreditLimitApproved *IssuingCreditUnderwritingRecordReportDecisionDecisionCreditLimitApprovedParams `form:"credit_limit_approved"`
	// Outcome of the decision.
	Type *string `form:"type"`
}

// If an exception to the usual underwriting criteria was made for this decision, details about the exception must be provided. Exceptions should only be granted in rare circumstances, in consultation with Stripe Compliance.
type IssuingCreditUnderwritingRecordReportDecisionUnderwritingExceptionParams struct {
	// Written explanation for the exception.
	Explanation *string `form:"explanation"`
	// The decision before the exception was applied.
	OriginalDecisionType *string `form:"original_decision_type"`
}

// Update a CreditUnderwritingRecord object from an decision made on a credit application
type IssuingCreditUnderwritingRecordReportDecisionParams struct {
	Params `form:"*"`
	// Date when a decision was made.
	DecidedAt *int64 `form:"decided_at"`
	// Details about the decision.
	Decision *IssuingCreditUnderwritingRecordReportDecisionDecisionParams `form:"decision"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// If an exception to the usual underwriting criteria was made for this decision, details about the exception must be provided. Exceptions should only be granted in rare circumstances, in consultation with Stripe Compliance.
	UnderwritingException *IssuingCreditUnderwritingRecordReportDecisionUnderwritingExceptionParams `form:"underwriting_exception"`
}

// AddExpand appends a new field to expand.
func (p *IssuingCreditUnderwritingRecordReportDecisionParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *IssuingCreditUnderwritingRecordReportDecisionParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Details about the application submission.
type IssuingCreditUnderwritingRecordCorrectApplicationParams struct {
	// The channel through which the applicant has submitted their application. Defaults to `online`.
	ApplicationMethod *string `form:"application_method"`
	// Scope of demand made by the applicant.
	Purpose *string `form:"purpose"`
	// Date when the applicant submitted their application.
	SubmittedAt *int64 `form:"submitted_at"`
}

// Information about the company or person applying or holding the account.
type IssuingCreditUnderwritingRecordCorrectCreditUserParams struct {
	// Email of the applicant or accountholder.
	Email *string `form:"email"`
	// Full name of the company or person.
	Name *string `form:"name"`
}

// Details about the application rejection.
type IssuingCreditUnderwritingRecordCorrectDecisionApplicationRejectedParams struct {
	// Details about the `reasons.other` when present.
	ReasonOtherExplanation *string `form:"reason_other_explanation"`
	// List of reasons why the application was rejected, up to 4 reasons, in order of importance.
	Reasons []*string `form:"reasons"`
}

// Details about the credit limit approved. An approved credit limit is required before you can set a `credit_limit_amount` in the [CreditPolicy API](https://stripe.com/docs/api/issuing/credit_policy/)
type IssuingCreditUnderwritingRecordCorrectDecisionCreditLimitApprovedParams struct {
	// The credit approved, in the currency of the account and [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
	Amount *int64 `form:"amount"`
	// The currency of the credit approved, will default to the Account's Issuing currency.
	Currency *string `form:"currency"`
}

// Details about the credit limit decreased.
type IssuingCreditUnderwritingRecordCorrectDecisionCreditLimitDecreasedParams struct {
	// The credit approved, in the currency of the account and [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
	Amount *int64 `form:"amount"`
	// The currency of the credit approved, will default to the Account's Issuing currency.
	Currency *string `form:"currency"`
	// Details about the `reasons.other` when present.
	ReasonOtherExplanation *string `form:"reason_other_explanation"`
	// List of reasons why the existing credit was decreased, up to 4 reasons, in order of importance.
	Reasons []*string `form:"reasons"`
}

// Details about the credit line closed.
type IssuingCreditUnderwritingRecordCorrectDecisionCreditLineClosedParams struct {
	// Details about the `reasons.other` when present.
	ReasonOtherExplanation *string `form:"reason_other_explanation"`
	// List of reasons why the credit line was closed, up to 4 reasons, in order of importance.
	Reasons []*string `form:"reasons"`
}

// Details about the decision.
type IssuingCreditUnderwritingRecordCorrectDecisionParams struct {
	// Details about the application rejection.
	ApplicationRejected *IssuingCreditUnderwritingRecordCorrectDecisionApplicationRejectedParams `form:"application_rejected"`
	// Details about the credit limit approved. An approved credit limit is required before you can set a `credit_limit_amount` in the [CreditPolicy API](https://stripe.com/docs/api/issuing/credit_policy/)
	CreditLimitApproved *IssuingCreditUnderwritingRecordCorrectDecisionCreditLimitApprovedParams `form:"credit_limit_approved"`
	// Details about the credit limit decreased.
	CreditLimitDecreased *IssuingCreditUnderwritingRecordCorrectDecisionCreditLimitDecreasedParams `form:"credit_limit_decreased"`
	// Details about the credit line closed.
	CreditLineClosed *IssuingCreditUnderwritingRecordCorrectDecisionCreditLineClosedParams `form:"credit_line_closed"`
	// Outcome of the decision.
	Type *string `form:"type"`
}

// If an exception to the usual underwriting criteria was made for this decision, details about the exception must be provided. Exceptions should only be granted in rare circumstances, in consultation with Stripe Compliance.
type IssuingCreditUnderwritingRecordCorrectUnderwritingExceptionParams struct {
	// Written explanation for the exception.
	Explanation *string `form:"explanation"`
	// The decision before the exception was applied.
	OriginalDecisionType *string `form:"original_decision_type"`
}

// Update a CreditUnderwritingRecord object to correct mistakes
type IssuingCreditUnderwritingRecordCorrectParams struct {
	Params `form:"*"`
	// Details about the application submission.
	Application *IssuingCreditUnderwritingRecordCorrectApplicationParams `form:"application"`
	// Information about the company or person applying or holding the account.
	CreditUser *IssuingCreditUnderwritingRecordCorrectCreditUserParams `form:"credit_user"`
	// Date when a decision was made.
	DecidedAt *int64 `form:"decided_at"`
	// Details about the decision.
	Decision *IssuingCreditUnderwritingRecordCorrectDecisionParams `form:"decision"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// If an exception to the usual underwriting criteria was made for this decision, details about the exception must be provided. Exceptions should only be granted in rare circumstances, in consultation with Stripe Compliance.
	UnderwritingException *IssuingCreditUnderwritingRecordCorrectUnderwritingExceptionParams `form:"underwriting_exception"`
}

// AddExpand appends a new field to expand.
func (p *IssuingCreditUnderwritingRecordCorrectParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *IssuingCreditUnderwritingRecordCorrectParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// For decisions triggered by an application, details about the submission.
type IssuingCreditUnderwritingRecordApplication struct {
	// The channel through which the applicant has submitted their application.
	ApplicationMethod IssuingCreditUnderwritingRecordApplicationApplicationMethod `json:"application_method"`
	// Scope of demand made by the applicant.
	Purpose IssuingCreditUnderwritingRecordApplicationPurpose `json:"purpose"`
	// Date when the applicant submitted their application.
	SubmittedAt int64 `json:"submitted_at"`
}
type IssuingCreditUnderwritingRecordCreditUser struct {
	// Email of the applicant or accountholder.
	Email string `json:"email"`
	// Full name of the company or person.
	Name string `json:"name"`
}

// Details about a decision application_rejected.
type IssuingCreditUnderwritingRecordDecisionApplicationRejected struct {
	// Details about the `reasons.other` when present.
	ReasonOtherExplanation string `json:"reason_other_explanation"`
	// List of reasons why the application was rejected up to 4 reasons, in order of importance.
	Reasons []IssuingCreditUnderwritingRecordDecisionApplicationRejectedReason `json:"reasons"`
}

// Details about a decision credit_limit_approved.
type IssuingCreditUnderwritingRecordDecisionCreditLimitApproved struct {
	// Credit amount approved. An approved credit limit is required before you can set a amount in the [CreditPolicy API](https://stripe.com/docs/api/issuing/credit_policy).
	Amount int64 `json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
}

// Details about a decision credit_limit_decreased.
type IssuingCreditUnderwritingRecordDecisionCreditLimitDecreased struct {
	// Credit amount approved after decrease. An approved credit limit is required before you can set a amount in the [CreditPolicy API](https://stripe.com/docs/api/issuing/credit_policy).
	Amount int64 `json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Details about the `reasons.other` when present.
	ReasonOtherExplanation string `json:"reason_other_explanation"`
	// List of reasons why the existing credit was decreased, up to 4 reasons, in order of importance.
	Reasons []IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReason `json:"reasons"`
}

// Details about a decision credit_line_closed.
type IssuingCreditUnderwritingRecordDecisionCreditLineClosed struct {
	// Details about the `reasons.other` when present.
	ReasonOtherExplanation string `json:"reason_other_explanation"`
	// List of reasons why the existing account was closed, up to 4 reasons, in order of importance.
	Reasons []IssuingCreditUnderwritingRecordDecisionCreditLineClosedReason `json:"reasons"`
}

// Details about the decision.
type IssuingCreditUnderwritingRecordDecision struct {
	// Details about a decision application_rejected.
	ApplicationRejected *IssuingCreditUnderwritingRecordDecisionApplicationRejected `json:"application_rejected"`
	// Details about a decision credit_limit_approved.
	CreditLimitApproved *IssuingCreditUnderwritingRecordDecisionCreditLimitApproved `json:"credit_limit_approved"`
	// Details about a decision credit_limit_decreased.
	CreditLimitDecreased *IssuingCreditUnderwritingRecordDecisionCreditLimitDecreased `json:"credit_limit_decreased"`
	// Details about a decision credit_line_closed.
	CreditLineClosed *IssuingCreditUnderwritingRecordDecisionCreditLineClosed `json:"credit_line_closed"`
	// Outcome of the decision.
	Type IssuingCreditUnderwritingRecordDecisionType `json:"type"`
}

// If an exception to the usual underwriting criteria was made for this application, details about the exception must be provided. Exceptions should only be granted in rare circumstances, in consultation with Stripe Compliance.
type IssuingCreditUnderwritingRecordUnderwritingException struct {
	// Written explanation for the exception.
	Explanation string `json:"explanation"`
	// The decision before the exception was applied.
	OriginalDecisionType IssuingCreditUnderwritingRecordUnderwritingExceptionOriginalDecisionType `json:"original_decision_type"`
}

// Every time an applicant submits an application for a Charge Card product your Platform offers, or every time your Platform takes a proactive credit decision on an existing account, you must record the decision by creating a new CreditUnderwritingRecord object on a Connected account.
//
// [Follow the guide](https://stripe.com/docs/issuing/coming_soon) to learn about your requirements as a Platform.
type IssuingCreditUnderwritingRecord struct {
	APIResource
	// For decisions triggered by an application, details about the submission.
	Application *IssuingCreditUnderwritingRecordApplication `json:"application"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The event that triggered the underwriting.
	CreatedFrom IssuingCreditUnderwritingRecordCreatedFrom `json:"created_from"`
	CreditUser  *IssuingCreditUnderwritingRecordCreditUser `json:"credit_user"`
	// Date when a decision was made.
	DecidedAt int64 `json:"decided_at"`
	// Details about the decision.
	Decision *IssuingCreditUnderwritingRecordDecision `json:"decision"`
	// For underwriting initiated by an application, a decision must be taken 30 days after the submission.
	DecisionDeadline int64 `json:"decision_deadline"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// If an exception to the usual underwriting criteria was made for this application, details about the exception must be provided. Exceptions should only be granted in rare circumstances, in consultation with Stripe Compliance.
	UnderwritingException *IssuingCreditUnderwritingRecordUnderwritingException `json:"underwriting_exception"`
}

// IssuingCreditUnderwritingRecordList is a list of CreditUnderwritingRecords as retrieved from a list endpoint.
type IssuingCreditUnderwritingRecordList struct {
	APIResource
	ListMeta
	Data []*IssuingCreditUnderwritingRecord `json:"data"`
}
