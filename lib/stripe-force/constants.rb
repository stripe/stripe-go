# typed: strict
# frozen_string_literal: true

module StripeForce
  module Constants
    # application constants
    POLL_FREQUENCY = T.let(3 * 60, Integer)
    MAX_STRIPE_PRICE_PRECISION = 12
    MAX_SF_RETRY_ATTEMPTS = 8

    SF_ORDER = 'Order'
    SF_ORDER_ITEM = 'OrderItem'
    SF_PRODUCT = 'Product2'
    SF_ACCOUNT = 'Account'
    SF_CONTACT = 'Contact'
    SF_OPPORTUNITY = 'Opportunity'
    SF_PRICEBOOK = 'Pricebook2'
    SF_PRICEBOOK_ENTRY = 'PricebookEntry'
    SF_CONSUMPTION_SCHEDULE = 'ConsumptionSchedule'
    SF_PRODUCT_CONSUMPTION_SCHEDULE = 'ProductConsumptionSchedule'
    SF_CONSUMPTION_RATE = 'ConsumptionRate'
    SF_CONTRACT = 'Contract'

    # custom coupon objects
    QUOTE_SF_STRIPE_COUPON = 'Quote_Stripe_Coupon__c'
    QUOTE_SF_STRIPE_COUPON_ASSOCIATION = 'Quote_Stripe_Coupon_Association__c'
    QUOTE_LINE_SF_STRIPE_COUPON_ASSOCIATION = 'Quote_Line_Stripe_Coupon_Association__c'
    ORDER_SF_STRIPE_COUPON = 'Order_Stripe_Coupon__c'

    SF_ID = 'Id'
    SF_LAST_MODIFIED_DATE = 'LastModifiedDate'
    SF_ORDER_ACCOUNT = 'AccountId'
    SF_ORDER_CONTRACTED = 'SBQQ__Contracted__c'
    SF_ORDER_QUOTE = 'SBQQ__Quote__c'

    SF_CURRENCY_ISO_CODE = 'CurrencyIsoCode'

    SF_ORDER_ITEM_REVISED_ORDER_PRODUCT = 'SBQQ__RevisedOrderProduct__c'

    SF_CONTRACT_ORDER_ID = 'SBQQ__Order__c'
    SF_CONTRACT_QUOTE_ID = 'SBQQ__Quote__c'

    CPQ_QUOTE = 'SBQQ__Quote__c'
    CPQ_QUOTE_LINE = 'SBQQ__QuoteLine__c'
    CPQ_QUOTE_LINE_PRODUCT = 'SBQQ__Product__c'
    CPQ_QUOTE_LINE_PRICEBOOK_ENTRY = 'SBQQ__PricebookEntryId__c'
    CPQ_CONSUMPTION_SCHEDULE = 'SBQQ__OrderItemConsumptionSchedule__c'
    CPQ_CONSUMPTION_RATE = 'SBQQ__OrderItemConsumptionRate__c'
    CPQ_AMENDED_CONTRACT = 'SBQQ__AmendedContract__c'

    CPQ_QUOTE_PRIMARY_CONTACT = 'SBQQ__PrimaryContact__c'
    CPQ_QUOTE_PRICEBOOK = "SBQQ__PricebookId__c"
    CPQ_QUOTE_OPPORTUNITY = 'SBQQ__Opportunity2__c'
    CPQ_QUOTE_ORDERED = 'SBQQ__Ordered__c'
    CPQ_QUOTE_PRIMARY = 'SBQQ__Primary__c'
    CPQ_QUOTE_SUBSCRIPTION_START_DATE = 'SBQQ__StartDate__c'
    CPQ_QUOTE_SUBSCRIPTION_TERM = 'SBQQ__SubscriptionTerm__c'
    CPQ_QUOTE_SUBSCRIPTION_PRICING = 'SBQQ__SubscriptionPricing__c'
    CPQ_QUOTE_QUANTITY = 'SBQQ__Quantity__c'

    CPQ_QUOTE_BILLING_FREQUENCY = 'SBQQ__BillingFrequency__c'
    class CPQBillingFrequencyOptions < T::Enum
      enums do
        MONTHLY = new("Monthly")
        QUARTERLY = new("Quarterly")
        SEMIANNUAL = new("Semiannual")
        ANNUAL = new("Annual")
        INVOICE_PLAN = new("Invoice Plan")
      end
    end

    CPQ_PRODUCT_SUBSCRIPTION_TYPE = 'SBQQ__SubscriptionType__c'
    class CPQProductSubscriptionTypeOptions < T::Enum
      enums do
        RENEWABLE = new("Renewable")
        ONE_TIME = new("")
      end
    end

    CPQ_PRODUCT_CHARGE_TYPE = 'SBQQ__ChargeType__c'

    # https://help.salesforce.com/s/articleView?id=sf.cpq_subscription_fields.htm&type=5
    CPQ_PRODUCT_BILLING_TYPE = 'SBQQ__BillingType__c'
    class CPQProductBillingTypeOptions < T::Enum
      enums do
        # licensed
        ADVANCE = new("Advance")

        # metered
        ARREARS = new("Arrears")
      end
    end

    # custom fields added by our package
    GENERIC_STRIPE_ID = 'Stripe_ID__c'
    ORDER_LINE_SKIP = 'Skip_Line_Item__c'
    ORDER_INVOICE_PAYMENT_LINK = 'Stripe_Invoice_Payment_Link__c'
    ORDER_SUBSCRIPTION_PAYMENT_LINK = 'Stripe_Subscription_Payment_Link__c'

    SYNC_RECORD = 'Sync_Record__c'

    class SalesforceStripeCouponFields < T::Enum
     enums do
        NAME = new('Name__c')
        PERCENT_OFF = new('Percent_Off__c')
        AMOUNT_OFF = new('Amount_Off__c')
        DURATION = new('Duration__c')
        DURATION_IN_MONTHS = new('Duration_In_Months__c')
        MAX_REDEMPTIONS = new('Max_Redemptions__c')
        CURRENCY_ISO_CODE = new(StripeForce::Constants::SF_CURRENCY_ISO_CODE)
     end
    end

    class SyncRecordFields < T::Enum
      enums do
        COMPOUND_ID = new('Compound_ID__c')

        PRIMARY_RECORD_ID = new("Primary_Record_ID__c")
        PRIMARY_OBJECT_TYPE = new("Primary_Object_Type__c")

        SECONDARY_RECORD_ID = new("Secondary_Record_ID__c")
        SECONDARY_OBJECT_TYPE = new("Secondary_Object_Type__c")

        RESOLUTION_MESSAGE = new("Resolution_Message__c")
        RESOLUTION_STATUS = new("Resolution_Status__c")
      end
    end

    class SyncRecordResolutionStatuses < T::Enum
      enums do
        SUCCESS = new('Success')
        ERROR = new('Error')
        RESOLVED = new('Resolved')
      end
    end

    # TODO these types are heavily customized by our users and should not be relied upon in production
    # https://developer.salesforce.com/docs/atlas.en-us.packagingGuide.meta/packagingGuide/coa_order_type.htm
    class OrderTypeOptions < T::Enum
      enums do
        NEW = new("New")
        AMENDMENT = new("Amendment")
        RENEWAL = new("Renewal")
      end
    end

    class OrderStatusOptions < T::Enum
      enums do
        ACTIVATED = new("Activated")
      end
    end

    # non-cpq constants

    SALESFORCE_ACCOUNT_ID_HEADER = 'Salesforce-Account-Id'
    SALESFORCE_KEY_HEADER = 'Salesforce-Key'
    SALESFORCE_PACKAGE_NAMESPACE_HEADER = "Salesforce-Package-Namespace"
    SALESFORCE_INSTANCE_TYPE_HEADER = 'Salesforce-Type'
    SALESFORCE_PACKAGE_ID_HEADER = 'Salesforce-Package-Id'

    CONNECTOR_SETTING_SALESFORCE_NAMESPACE = "salesforce_namespace"
    CONNECTOR_SETTING_SALESFORCE_INSTANCE_TYPE = 'salesforce_instance_type'
    CONNECTOR_SETTING_CPQ_TERM_UNIT = 'cpq_term_unit'
    CONNECTOR_SETTING_MULTICURRENCY_ENABLED = 'multicurrency_enabled'

    class SalesforceNamespaceOptions < T::Enum
      enums do
        QA = new("QaStripeConnect")
        PRODUCTION = new("stripeConnector")
        NONE = new("c")
      end
    end

    class SFInstanceTypes < T::Enum
      enums do
        SANDBOX = new("SANDBOX")
        SCRATCH_ORG = new("SCRATCH_ORG")
        TRIAL = new("TRIAL")
        PRODUCTION = new("PRODUCTION")
      end
    end

    class FeatureFlags < T::Enum
      enums do
        IGNORE_WEBHOOKS = new('ignore_webhooks')
        REJECT_WEBHOOKS = new('reject_webhooks')
        LOUD_SANDBOX_LOGGING = new('loud_sandbox_logging')
        TEST_CLOCKS = new('test_clocks')
        SF_CACHING = new('sf_caching')
        CATCH_ALL_ERRORS = new('catch_all_errors')
        UPDATE_CUSTOMER_ON_ORDER_TRANSLATION = new('update_customer_on_order_creation')
        ACCOUNT_POLLING = new('account_polling')
        COUPONS = new('coupons')
        TERMINATED_ORDER_ITEM_CREDIT = new('terminated_order_item_credit')
        AUTO_ADVANCE_PRORATION_INVOICE = new('auto_advance_proration_invoices')
      end
    end

    class MetadataKeys < T::Enum
      enums do
        PRORATION = new('proration')
        CREDIT = new('credit')
        DUPLICATE_PRICE = new('duplicate')
        AUTO_ARCHIVE_PRICE = new('auto_archive')
        ORIGINAL_PRICE_ID = new('original_stripe_price_id')
        PRORATION_INVOICE = new('proration_invoice')
        BACKEND_PRORATION = new('backend_proration')
      end
    end

  end
end
