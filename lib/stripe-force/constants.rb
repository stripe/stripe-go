# typed: strict
# frozen_string_literal: true

module StripeForce
  module Constants
    SF_ORDER = 'Order'
    SF_PRODUCT = 'Product2'

    CPQ_QUOTE = 'SBQQ__Quote__c'

    CPQ_QUOTE_PRIMARY_CONTACT = 'SBQQ__PrimaryContact__c'
    CPQ_QUOTE_OPPORTUNITY = 'SBQQ__Opportunity2__c'
    CPQ_QUOTE_ORDERED = 'SBQQ__Ordered__c'
    CPQ_QUOTE_PRIMARY = 'SBQQ__Primary__c'
    CPQ_QUOTE_SUBSCRIPTION_START_DATE = 'SBQQ__StartDate__c'
    CPQ_QUOTE_SUBSCRIPTION_TERM = 'SBQQ__SubscriptionTerm__c'

    CPQ_QUOTE_LINE = 'SBQQ__QuoteLine__c'
    CPQ_QUOTE_LINE_PRODUCT = 'SBQQ__Product__c'
    CPQ_QUOTE_LINE_PRICEBOOK_ENTRY = 'SBQQ__PricebookEntryId__c'

    OPPORTUNITY_STRIPE_ID = 'Stripe_Customer_ID__c'
    ORDER_STRIPE_ID = 'Stripe_Transaction_ID__c'
    ORDER_INVOICE_PAYMENT_LINK = 'Stripe_Invoice_Payment_Link__c'
    PRICE_BOOK_STRIPE_ID = 'Stripe_Price_ID__c'
  end
end
