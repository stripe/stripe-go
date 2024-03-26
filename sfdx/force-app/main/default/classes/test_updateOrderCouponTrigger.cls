@IsTest
public with sharing class test_updateOrderCouponTrigger {
    private static final String COUPON_NAME = '55 percent off coupon';
    private static final Integer COUPON_NAME_LENGTH = COUPON_NAME.length();
    private static final Integer COUPON_PERCENT_OFF = 55;
    private static final String COUPON_DURATION = 'once';
    public static final String ORDER_STATUS_ACTIVATED = 'Activated';
    public static final String ORDER_STATUS_DRAFT = 'Draft';
    private static final String FIELD_CURRENCY_ISO_CODE = constants.FIELD_CURRENCY_ISO_CODE;

    @IsTest
    static public void testUpdateOrderCouponTriggerOnQuoteCoupons() {
        Boolean isCpqInstalled = utilities.isCpqEnabled();
        if(!isCpqInstalled) {
            return;
        }

        // setup test data
        SBQQ__Quote__c quote = createQuote();
        Quote_Stripe_Coupon__c coupon = createStripeCoupon();
        createStripeCouponQuoteAssociation(quote.Id, coupon.Id);

        // fires trigger
        Test.startTest();
        Order order = activateOrderFromQuote(quote);
        Test.stopTest();

        // sanity check that the Quote Stripe Coupon Association was created
        List<Quote_Stripe_Coupon_Association__c> stripeCouponQuoteAssociations = getStripeCouponQuoteAssociation(quote.Id);
        System.assertEquals(1, stripeCouponQuoteAssociations.size());

        // sanity check that the Quote Stripe Coupon has the correct data
        List<Quote_Stripe_Coupon__c> quoteStripeCoupons = getStripeCoupons(stripeCouponQuoteAssociations.get(0).Quote_Stripe_Coupon__c);
        System.assertEquals(1, quoteStripeCoupons.size());
        
        Quote_Stripe_Coupon__c quoteCoupon = quoteStripeCoupons.get(0);
        System.assertEquals(COUPON_NAME, quoteCoupon.Name__c.substring(0, COUPON_NAME_LENGTH));
        System.assertEquals(COUPON_PERCENT_OFF, quoteCoupon.Percent_Off__c);
        System.assertEquals(COUPON_DURATION, quoteCoupon.Duration__c);

        // trigger should have created an Order Stripe Coupon
        List<Order_Stripe_Coupon__c> orderStripeCoupons = getOrderCouponFromQuoteCouponId(quoteCoupon.Id);
        System.assertEquals(1, orderStripeCoupons.size());
        
        Order_Stripe_Coupon__c orderStripeCoupon = orderStripeCoupons.get(0);
        System.assertEquals(COUPON_NAME, orderStripeCoupon.Name__c.substring(0, COUPON_NAME_LENGTH));
        System.assertEquals(order.Id, orderStripeCoupon.Order__c);
        System.assertEquals(null, orderStripeCoupon.Order_Item__c);
        System.assertEquals(quoteCoupon.Percent_Off__c, orderStripeCoupon.Percent_Off__c);
        System.assertEquals(quoteCoupon.Duration__c, orderStripeCoupon.Duration__c);
    }

    @IsTest
    static public void testUpdateOrderCouponTriggerOnQuoteLineCoupons() {
        Boolean isCpqInstalled = utilities.isCpqEnabled();
        if(!isCpqInstalled) {
            return;
        }

        String curr = UserInfo.getDefaultCurrency();

        // setup test data
        SBQQ__Quote__c quote = createQuote();
        Quote_Stripe_Coupon__c coupon = createStripeCoupon();
        SBQQ__QuoteLine__c quoteLine = [SELECT Id FROM SBQQ__QuoteLine__c WHERE SBQQ__Quote__c = :quote.Id];
        createStripeCouponQuoteLineAssociation(quoteLine.Id, coupon.Id);

        // fires trigger
        Test.startTest();
        Order order = activateOrderFromQuote(quote);
        Test.stopTest();

        List<SBQQ__QuoteLine__c> quoteLines = [SELECT Id FROM SBQQ__QuoteLine__c WHERE SBQQ__Quote__c = :quote.Id];
        System.assertEquals(1, quoteLines.size());

        // sanity check that the Quote Stripe Coupon Association was created
        List<Quote_Line_Stripe_Coupon_Association__c> stripeCouponQuoteLineAssociations = getStripeCouponQuoteLineAssociation(quoteLines.get(0).Id);
        System.assertEquals(1, stripeCouponQuoteLineAssociations.size());

        List<OrderItem> orderItems = [SELECT Id FROM OrderItem WHERE OrderId = :order.Id];
        System.assertEquals(1, orderItems.size());
        OrderItem orderItem = orderItems.get(0);

        // sanity check that the Quote Stripe Coupon has the correct data
        List<Quote_Stripe_Coupon__c> quoteStripeCoupons = getStripeCoupons(stripeCouponQuoteLineAssociations.get(0).Quote_Stripe_Coupon__c);
        System.assertEquals(1, quoteStripeCoupons.size());

        Quote_Stripe_Coupon__c quoteCoupon = quoteStripeCoupons.get(0);
        System.assertEquals(COUPON_NAME, quoteCoupon.Name__c.substring(0, COUPON_NAME_LENGTH));
        System.assertEquals(COUPON_PERCENT_OFF, quoteCoupon.Percent_Off__c);
        System.assertEquals(COUPON_DURATION, quoteCoupon.Duration__c);

        // trigger should have created an Order Stripe Coupon
        List<Order_Stripe_Coupon__c> orderStripeCoupons = getOrderCouponFromQuoteCouponId(quoteCoupon.Id);
        System.assertEquals(1, orderStripeCoupons.size());

        Order_Stripe_Coupon__c orderStripeCoupon = orderStripeCoupons.get(0);
        System.assertEquals(COUPON_NAME, orderStripeCoupon.Name__c.substring(0, COUPON_NAME_LENGTH));
        System.assertEquals(null, orderStripeCoupon.Order__c);
        System.assertEquals(orderItem.Id, orderStripeCoupon.Order_Item__c);
        System.assertEquals(quoteCoupon.Percent_Off__c, orderStripeCoupon.Percent_Off__c);
        System.assertEquals(quoteCoupon.Duration__c, orderStripeCoupon.Duration__c);
    }

    @IsTest
    static public void testDoubleFiringTrigger() {
        Boolean isCpqInstalled = utilities.isCpqEnabled();
        if(!isCpqInstalled) {
            return;
        }

        // setup test data
        SBQQ__Quote__c quote = createQuote();
        SBQQ__QuoteLine__c quoteLine = [SELECT Id FROM SBQQ__QuoteLine__c WHERE SBQQ__Quote__c = :quote.Id];
        
        // add a coupon on the quote
        Quote_Stripe_Coupon__c quoteCoupon = createStripeCoupon();
        createStripeCouponQuoteAssociation(quote.Id, quoteCoupon.Id);

        // add two coupons on the quote line
        Quote_Stripe_Coupon__c quoteLineCoupon = createStripeCoupon();
        createStripeCouponQuoteLineAssociation(quoteLine.Id, quoteLineCoupon.Id);
        createStripeCouponQuoteLineAssociation(quoteLine.Id, quoteLineCoupon.Id);

        // fire trigger by activating the quote
        Order order = activateOrderFromQuote(quote);
 
        // initial trigger should have created 3 order coupons (1 on the order and 2 on the order line)
        List<Order_Stripe_Coupon__c> orderStripeCoupons = getOrderCouponFromQuoteCouponId(quoteCoupon.Id);
        System.assertEquals(1, orderStripeCoupons.size());

        List<Order_Stripe_Coupon__c> orderLineStripeCoupons = getOrderCouponFromQuoteCouponId(quoteLineCoupon.Id);
        System.assertEquals(2, orderLineStripeCoupons.size());

        // now let's refire the updateCoupons trigger
        Test.startTest();
        refireTrigger(order);
        Test.stopTest();

        // confirm that no new order coupons were created due to the trigger firing again
        orderStripeCoupons = getOrderCouponFromQuoteCouponId(quoteCoupon.Id);
        System.assertEquals(1, orderStripeCoupons.size());
        orderLineStripeCoupons = getOrderCouponFromQuoteCouponId(quoteLineCoupon.Id);
        System.assertEquals(2, orderLineStripeCoupons.size());
    }

    @IsTest
    static public void testUpdateOrderCouponTriggerOnQuoteCouponsWithMultiCurrency() {
        Boolean isCpqInstalled = utilities.isCpqEnabled();

        if(!isCpqInstalled || UserInfo.isMultiCurrencyOrganization() == false) {
            return;
        }

        String isoCurrency = utilities.getCurrencyIsoCodePreferNonDefault();

        // setup test data
        SBQQ__Quote__c quote = createQuote(isoCurrency);
        Quote_Stripe_Coupon__c coupon = createStripeCoupon(isoCurrency);
        createStripeCouponQuoteAssociation(quote.Id, coupon.Id, isoCurrency);

        System.assertEquals(isoCurrency, quote.get(FIELD_CURRENCY_ISO_CODE));
        System.assertEquals(isoCurrency, coupon.get(FIELD_CURRENCY_ISO_CODE));

        // fires trigger
        Test.startTest();
        Order order = activateOrderFromQuote(quote);
        Test.stopTest();

        // sanity check that the Quote Stripe Coupon Association was created
        List<Quote_Stripe_Coupon_Association__c> stripeCouponQuoteAssociations = getStripeCouponQuoteAssociation(quote.Id);
        System.assertEquals(1, stripeCouponQuoteAssociations.size());
        Quote_Stripe_Coupon_Association__c stripeCouponQuoteAssociation = stripeCouponQuoteAssociations.get(0);
        System.assertEquals(isoCurrency, stripeCouponQuoteAssociation.get(FIELD_CURRENCY_ISO_CODE));

        // sanity check that the Quote Stripe Coupon has the correct data
        List<Quote_Stripe_Coupon__c> quoteStripeCoupons = getStripeCoupons(stripeCouponQuoteAssociations.get(0).Quote_Stripe_Coupon__c);
        System.assertEquals(1, quoteStripeCoupons.size());

        Quote_Stripe_Coupon__c quoteCoupon = quoteStripeCoupons.get(0);
        System.assertEquals(COUPON_NAME, quoteCoupon.Name__c.substring(0, COUPON_NAME_LENGTH));
        System.assertEquals(COUPON_PERCENT_OFF, quoteCoupon.Percent_Off__c);
        System.assertEquals(COUPON_DURATION, quoteCoupon.Duration__c);
        System.assertEquals(isoCurrency, quoteCoupon.get(FIELD_CURRENCY_ISO_CODE));

        // trigger should have created an Order Stripe Coupon
        List<Order_Stripe_Coupon__c> orderStripeCoupons = getOrderCouponFromQuoteCouponId(quoteCoupon.Id);
        System.assertEquals(1, orderStripeCoupons.size());

        Order_Stripe_Coupon__c orderStripeCoupon = orderStripeCoupons.get(0);
        System.assertEquals(COUPON_NAME, orderStripeCoupon.Name__c.substring(0, COUPON_NAME_LENGTH));
        System.assertEquals(isoCurrency, orderStripeCoupon.get(FIELD_CURRENCY_ISO_CODE));
    }

    static private Order activateOrderFromQuote(SBQQ__Quote__c quote) {
        // update quote to be ordered
        quote.SBQQ__Ordered__c = true;
        Database.SaveResult result = Database.update(quote, true);
        if (!result.isSuccess()) {
            // operation failed, so get all errors
            for(Database.Error err : result.getErrors()) {
                System.debug('The following error has occurred.');
                System.debug(err.getStatusCode() + ': ' + err.getMessage());
                System.debug('Fields that affected this error: ' + err.getFields());
            }
        }

        // mark the order as activated
        List<Order> orders = [
                SELECT Id
                FROM Order
                WHERE SBQQ__Quote__c = :quote.Id];
        System.assertEquals(1, orders.size());

        Order order = orders.get(0);
        order.Status = ORDER_STATUS_ACTIVATED;
        Database.update(order, true);

        return order;
    }

    static public SBQQ__Quote__c createQuote() {
        if (UserInfo.isMultiCurrencyOrganization()) {
            return createQuote(UserInfo.getDefaultCurrency());
        }

        return createQuote(null);
    }

    static public SBQQ__Quote__c createQuote(String currencyCode) {
        Id pricebookId = Test.getStandardPricebookId();

        Integer randomNumber = Integer.valueOf((Math.random() * 1000));

        // create an Account
        Account account = new Account();
        account.Name = 'SFDC Account ' + randomNumber;
        account.SBQQ__RenewalModel__c = 'Contract Based';
        account.SBQQ__ContractCoTermination__c = 'Never';
        account.SBQQ__CoTerminationEvent__c = 'Renewal';
        account.SBQQ__RenewalPricingMethod__c = 'Same';
        account.SBQQ__PriceHoldEnd__c = Date.today().addMonths(1);
        account.SBQQ__AssetQuantitiesCombined__c = false;
        account.SBQQ__CoTermedContractsCombined__c = false;
        account.SBQQ__PreserveBundle__c = true;
        conditionallyAddCurrencyCode(account, currencyCode);
        Database.insert(account);

        Product2 product = new Product2();
        product.Name = 'testProduct ' + randomNumber;
        product.Description = 'testProductEntry';
        product.ProductCode = 'ABC' + randomNumber;
        product.SBQQ__DefaultQuantity__c = 12.000;
        product.IsActive = true;
        conditionallyAddCurrencyCode(product, currencyCode);
        Database.insert(product);

        SBQQ__Cost__c cost = new SBQQ__Cost__c();
        cost.SBQQ__Active__c = true;
        cost.SBQQ__Product__c = product.Id;
        cost.SBQQ__UnitCost__c = 10.00;
        conditionallyAddCurrencyCode(cost, currencyCode);
        Database.insert(cost);
        
        PricebookEntry standardPrice = new PricebookEntry();
        standardPrice.Pricebook2Id = pricebookId;
        standardPrice.Product2Id = product.Id;
        standardPrice.UnitPrice = 1;
        standardPrice.IsActive = true;
        standardPrice.UseStandardPrice = false;
        conditionallyAddCurrencyCode(standardPrice, currencyCode);
        Database.insert(standardPrice);

        Opportunity opp = new Opportunity();
        opp.Name = 'testOpportunity ' + randomNumber;
        opp.AccountId = account.Id;
        // this may need to eventually do some magic to pull the stage from the env
        // some customers may not have a 'prospecting' stage defined/allowed for the default record type.
        opp.StageName = 'prospecting';
        opp.Pricebook2Id = pricebookId;
        opp.CloseDate = Date.today();
        conditionallyAddCurrencyCode(opp, currencyCode);
        Database.insert(opp);

        // Create the quote and quote line
        SBQQ__Quote__c quote = new SBQQ__Quote__c();
        quote.SBQQ__Account__c = account.Id;
        quote.SBQQ__Status__c = 'Draft';
        quote.SBQQ__StartDate__c = Date.today();
        quote.SBQQ__PricebookId__c = pricebookId;
        quote.SBQQ__Opportunity2__c = opp.Id;
        quote.SBQQ__SubscriptionTerm__c = 12;
        quote.SBQQ__ContractingMethod__c = 'By Subscription End Date';
        quote.SBQQ__Primary__c = true;
        conditionallyAddCurrencyCode(quote, currencyCode);
        Database.insert(quote);

        SBQQ__QuoteLine__c quoteLine = new SBQQ__QuoteLine__c();
        quoteLine.SBQQ__Quote__c = quote.Id;
        quoteLine.SBQQ__Product__c = product.Id;
        quoteLine.SBQQ__PricebookEntryId__c = standardPrice.Id;
        quoteLine.SBQQ__Quantity__c = 10;
        conditionallyAddCurrencyCode(quoteLine, currencyCode);
        Database.insert(quoteLine);

        // moved setting the quote to approved to after the quote line was recorded.
        // If you don't follow normal lifecycle stuff in tests, things can get weird.
        quote.SBQQ__Status__c = 'Approved';
        update quote;

        return quote; 
    }

    static public Quote_Stripe_Coupon__c createStripeCoupon() {
        if (UserInfo.isMultiCurrencyOrganization()) {
            return createStripeCoupon(UserInfo.getDefaultCurrency());
        }

        return createStripeCoupon(null);
    }
    static public Quote_Stripe_Coupon__c createStripeCoupon(String currencyCode) {
        Integer randomNumber = Integer.valueOf((Math.random() * 1000));

        Quote_Stripe_Coupon__c salesforceStripeCoupon =  new Quote_Stripe_Coupon__c();
        salesforceStripeCoupon.Name__c = COUPON_NAME + ' ' + randomNumber;
        salesforceStripeCoupon.Percent_Off__c = COUPON_PERCENT_OFF;
        salesforceStripeCoupon.Duration__c = COUPON_DURATION;
        conditionallyAddCurrencyCode(salesforceStripeCoupon, currencyCode);
        Database.insert(salesforceStripeCoupon);
        return salesforceStripeCoupon;
    }

    static public void createStripeCouponQuoteAssociation(Id quoteId, Id couponId) {
        if (UserInfo.isMultiCurrencyOrganization()) {
            createStripeCouponQuoteAssociation(quoteId, couponId, UserInfo.getDefaultCurrency());
            return;
        }

        createStripeCouponQuoteAssociation(quoteId, couponId, null);
    }
    static public void createStripeCouponQuoteAssociation(Id quoteId, Id couponId, String currencyCode) {
        Quote_Stripe_Coupon_Association__c salesforceStripeCouponQuoteAssociation =  new Quote_Stripe_Coupon_Association__c();
        salesforceStripeCouponQuoteAssociation.Quote__c = quoteId;
        salesforceStripeCouponQuoteAssociation.Quote_Stripe_Coupon__c = couponId;

        if (currencyCode != null) {
            salesforceStripeCouponQuoteAssociation.put(FIELD_CURRENCY_ISO_CODE, currencyCode);
        }

        Database.insert(salesforceStripeCouponQuoteAssociation);
    }

    static public void createStripeCouponQuoteLineAssociation(Id quoteLineId, Id couponId) {
        if (UserInfo.isMultiCurrencyOrganization()) {
            createStripeCouponQuoteLineAssociation(quoteLineId, couponId, UserInfo.getDefaultCurrency());
            return;
        }

        createStripeCouponQuoteLineAssociation(quoteLineId, couponId, null);
    }

    static public void createStripeCouponQuoteLineAssociation(Id quoteLineId, Id couponId, String currencyCode) {
        Quote_Line_Stripe_Coupon_Association__c assoc =  new Quote_Line_Stripe_Coupon_Association__c();
        assoc.Quote_Line__c = quoteLineId;
        assoc.Quote_Stripe_Coupon__c = couponId;

        if (currencyCode != null) {
            assoc.put(FIELD_CURRENCY_ISO_CODE, currencyCode);
        }

        Database.insert(assoc);
    }

    static public List<Quote_Stripe_Coupon__c> getStripeCoupons(Id couponId) {
        List<String> fields = new List<String> {
                'Id',
                'Amount_Off__c',
                'Duration__c',
                'Duration_In_Months__c',
                'Max_Redemptions__c',
                'Name__c',
                'Percent_Off__c',
                'Stripe_ID__c'
        };

        return utilities.execMultiCurrLookupQuery(
                Quote_Stripe_Coupon__c.SObjectType,
                'Id',
                new Set<Id> { couponId },
                fields,
                new List<String> { FIELD_CURRENCY_ISO_CODE }
        );
    }

    static public List<Quote_Stripe_Coupon_Association__c> getStripeCouponQuoteAssociation(Id quoteId) {
        return utilities.execMultiCurrLookupQuery(
                Quote_Stripe_Coupon_Association__c.SObjectType,
                'Quote__c',
                new Set<Id> { quoteId },
                new List<String> { 'Quote_Stripe_Coupon__c' },
                new List<String> { FIELD_CURRENCY_ISO_CODE }
        );
    }

    static public List<Quote_Line_Stripe_Coupon_Association__c> getStripeCouponQuoteLineAssociation(Id quoteLineId) {
        return utilities.execMultiCurrLookupQuery(
                Quote_Line_Stripe_Coupon_Association__c.SObjectType,
                'Quote_Line__c',
                new Set<Id> { quoteLineId },
                new List<String> { 'Quote_Stripe_Coupon__c' },
                new List<String> { FIELD_CURRENCY_ISO_CODE }
        );
    }

    static public List<Order_Stripe_Coupon__c> getOrderCouponFromQuoteCouponId(Id quoteCouponId) {
        List<String> fields = new List<String> {
            'Id',
            'Amount_Off__c',
            'Duration__c',
            'Duration_In_Months__c',
            'Max_Redemptions__c',
            'Name__c',
            'Order__c',
            'Order_Item__c',
            'Percent_Off__c',
            'Quote_Stripe_Coupon_Id__c',
            'Stripe_ID__c'
        };

        return utilities.execMultiCurrLookupQuery(
                Order_Stripe_Coupon__c.SObjectType,
                'Quote_Stripe_Coupon_Id__c',
                new Set<Id> { quoteCouponId },
                fields,
                new List<String> { FIELD_CURRENCY_ISO_CODE }
        );
    }

    // This helper is used only in test to cause the updateCoupons trigger to refire
    static private void refireTrigger(Order order)
    {
        order.Status = ORDER_STATUS_DRAFT;
        Database.update(order, true);
        order.Status = ORDER_STATUS_ACTIVATED;
        Database.update(order, true);
    }

    static private void conditionallyAddCurrencyCode(SObject obj, String currencyCode) {
        if (currencyCode != null) {
            obj.put(FIELD_CURRENCY_ISO_CODE, currencyCode);
        }
    }
}
