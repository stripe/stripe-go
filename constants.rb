module StripeForce
  module Constants
    CPQ_QUOTE = 'SBQQ__Quote__c'.freeze
    CPQ_QUOTE_PRIMARY_CONTACT = 'SBQQ__PrimaryContact__c'.freeze
    CPQ_QUOTE_OPPORTUNITY = 'SBQQ__Opportunity2__c'.freeze
    CPQ_QUOTE_ORDERED = 'SBQQ__Ordered__c'.freeze
    CPQ_QUOTE_PRIMARY = 'SBQQ__Primary__c'.freeze
    CPQ_QUOTE_LINE = 'SBQQ__QuoteLine__c'.freeze
    CPQ_QUOTE_LINE_PRODUCT = 'SBQQ__Product__c'.freeze
    CPQ_QUOTE_LINE_PRICEBOOK_ENTRY = 'SBQQ__PricebookEntryId__c'.freeze

    OPPORTUNITY_STRIPE_ID = 'Stripe_Customer_ID__c'.freeze
    ORDER_STRIPE_ID = 'Stripe_Transaction_ID__c'.freeze
    PRICE_BOOK_STRIPE_ID = 'Stripe_Price_ID__c'.freeze
  end
end