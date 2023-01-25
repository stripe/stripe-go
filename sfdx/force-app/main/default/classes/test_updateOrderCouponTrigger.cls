@IsTest
public with sharing class test_updateOrderCouponTrigger {
    private static final String COUPON_NAME = '55 percent off coupon';
    private static final Integer COUPON_NAME_LENGTH = COUPON_NAME.length();
    private static final Integer COUPON_PERCENT_OFF = 55;
    private static final String COUPON_DURATION = 'once';
    private static final String ORDER_STATUS_ACTIVATED = 'Activated';

    @IsTest
    static public void testUpdateOrderCouponTriggerOnQuoteCoupons() {
        Boolean isCpqInstalled = utilities.isCpqEnabled();
        if(!isCpqInstalled) {
            return;
        }

        // Setup test data
        SBQQ__Quote__c quote = createQuote();
        Quote_Stripe_Coupon__c coupon = createStripeCoupon();
        createStripeCouponQuoteAssociation(quote.Id, coupon.Id);

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
        List<Order_Stripe_Coupon__c> orderStripeCoupons = getOrderStripeCoupon(quoteCoupon.Id);
        System.assertEquals(1, orderStripeCoupons.size());
        
        Order_Stripe_Coupon__c orderStripeCoupon = orderStripeCoupons.get(0);
        System.assertEquals(COUPON_NAME, orderStripeCoupon.Name__c.substring(0, COUPON_NAME_LENGTH));
    }

    @IsTest
    static public void testUpdateOrderCouponTriggerOnQuoteLineCoupons() {
        Boolean isCpqInstalled = utilities.isCpqEnabled();
        if(!isCpqInstalled) {
            return;
        }

        // Setup test data
        SBQQ__Quote__c quote = createQuote();
        Quote_Stripe_Coupon__c coupon = createStripeCoupon();
        SBQQ__QuoteLine__c quoteLine = [SELECT Id FROM SBQQ__QuoteLine__c WHERE SBQQ__Quote__c = :quote.Id];
        createStripeCouponQuoteLineAssociation(quoteLine.Id, coupon.Id);

        Test.startTest();

        Order order = activateOrderFromQuote(quote);

        Test.stopTest();

        List<SBQQ__QuoteLine__c> quoteLines = [SELECT Id FROM SBQQ__QuoteLine__c WHERE SBQQ__Quote__c = :quote.Id];
        System.assertEquals(1, quoteLines.size());

        // sanity check that the Quote Stripe Coupon Association was created
        List<Quote_Line_Stripe_Coupon_Association__c> stripeCouponQuoteLineAssociations = getStripeCouponQuoteLineAssociation(quoteLines.get(0).Id);
        System.assertEquals(1, stripeCouponQuoteLineAssociations.size());

        // sanity check that the Quote Stripe Coupon has the correct data
        List<Quote_Stripe_Coupon__c> quoteStripeCoupons = getStripeCoupons(stripeCouponQuoteLineAssociations.get(0).Quote_Stripe_Coupon__c);
        System.assertEquals(1, quoteStripeCoupons.size());

        Quote_Stripe_Coupon__c quoteCoupon = quoteStripeCoupons.get(0);
        System.assertEquals(COUPON_NAME, quoteCoupon.Name__c.substring(0, COUPON_NAME_LENGTH));
        System.assertEquals(COUPON_PERCENT_OFF, quoteCoupon.Percent_Off__c);
        System.assertEquals(COUPON_DURATION, quoteCoupon.Duration__c);

        // trigger should have created an Order Stripe Coupon
        List<Order_Stripe_Coupon__c> orderStripeCoupons = getOrderStripeCoupon(quoteCoupon.Id);
        System.assertEquals(1, orderStripeCoupons.size());

        Order_Stripe_Coupon__c orderStripeCoupon = orderStripeCoupons.get(0);
        System.assertEquals(COUPON_NAME, orderStripeCoupon.Name__c.substring(0, COUPON_NAME_LENGTH));

        List<OrderItem> orderItems = [SELECT Id FROM OrderItem WHERE OrderId = :order.Id];
        System.assertEquals(1, orderItems.size());
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
        Database.insert(account);

        Product2 product = new Product2();
        product.Name = 'testProduct ' + randomNumber;
        product.Description = 'testProductEntry';
        product.ProductCode = 'ABC' + randomNumber;
        product.SBQQ__DefaultQuantity__c = 12.000;
        product.IsActive = true;
        Database.insert(product);

        SBQQ__Cost__c cost = new SBQQ__Cost__c();
        cost.SBQQ__Active__c = true;
        cost.SBQQ__Product__c = product.Id;
        cost.SBQQ__UnitCost__c = 10.00;
        Database.insert(cost);
        
        PricebookEntry standardPrice = new PricebookEntry();
        standardPrice.Pricebook2Id = pricebookId;
        standardPrice.Product2Id = product.Id;
        standardPrice.UnitPrice = 1;
        standardPrice.IsActive = true;
        standardPrice.UseStandardPrice = false;
        Database.insert(standardPrice);

        Opportunity opp = new Opportunity();
        opp.Name = 'testOpportunity ' + randomNumber;
        opp.AccountId = account.Id;
        // this may need to eventually do some magic to pull the stage from the env
        // some customers may not have a 'prospecting' stage defined/allowed for the default record type.
        opp.StageName = 'prospecting';
        opp.Pricebook2Id = pricebookId;
        opp.CloseDate = Date.today();
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
        Database.insert(quote);

        SBQQ__QuoteLine__c quoteLine = new SBQQ__QuoteLine__c();
        quoteLine.SBQQ__Quote__c = quote.Id;
        quoteLine.SBQQ__Product__c = product.Id;
        quoteLine.SBQQ__PricebookEntryId__c = standardPrice.Id;
        quoteLine.SBQQ__Quantity__c = 10; 
        Database.insert(quoteLine);

        // moved setting the quote to approved to after the quote line was recorded.
        // If you don't follow normal lifecycle stuff in tests, things can get weird.
        quote.SBQQ__Status__c = 'Approved';
        update quote;

        return quote; 
    }

    static public Quote_Stripe_Coupon__c createStripeCoupon() {
        Integer randomNumber = Integer.valueOf((Math.random() * 1000));

        Quote_Stripe_Coupon__c salesforceStripeCoupon =  new Quote_Stripe_Coupon__c();
        salesforceStripeCoupon.Name__c = COUPON_NAME + ' ' + randomNumber;
        salesforceStripeCoupon.Percent_Off__c = COUPON_PERCENT_OFF;
        salesforceStripeCoupon.Duration__c = COUPON_DURATION;
        Database.insert(salesforceStripeCoupon);
        return salesforceStripeCoupon;
    }

    static public void createStripeCouponQuoteAssociation(Id quoteId, Id couponId) {
        Quote_Stripe_Coupon_Association__c salesforceStripeCouponQuoteAssociation =  new Quote_Stripe_Coupon_Association__c();
        salesforceStripeCouponQuoteAssociation.Quote__c = quoteId;
        salesforceStripeCouponQuoteAssociation.Quote_Stripe_Coupon__c = couponId;
        Database.insert(salesforceStripeCouponQuoteAssociation);
    }

    static public void createStripeCouponQuoteLineAssociation(Id quoteLineId, Id couponId) {
        Quote_Line_Stripe_Coupon_Association__c assoc =  new Quote_Line_Stripe_Coupon_Association__c();
        assoc.Quote_Line__c = quoteLineId;
        assoc.Quote_Stripe_Coupon__c = couponId;
        Database.insert(assoc);
    }

    static public List<Quote_Stripe_Coupon__c> getStripeCoupons(Id couponId) {
        return [
            SELECT Id, Amount_Off__c, Duration__c, Duration_In_Months__c, Max_Redemptions__c, Name__c, Percent_Off__c
            FROM Quote_Stripe_Coupon__c
            WHERE Id = :couponId
        ];
    }

    static public List<Quote_Stripe_Coupon_Association__c> getStripeCouponQuoteAssociation(Id quoteId) {
        return [
          SELECT Quote_Stripe_Coupon__c
          FROM Quote_Stripe_Coupon_Association__c
          WHERE Quote__c = :quoteId
        ];
    }

    static public List<Quote_Line_Stripe_Coupon_Association__c> getStripeCouponQuoteLineAssociation(Id quoteLineId) {
        return [
                SELECT Quote_Stripe_Coupon__c
                FROM Quote_Line_Stripe_Coupon_Association__c
                WHERE Quote_Line__c = :quoteLineId
        ];
    }

    static public List<Order_Stripe_Coupon__c> getOrderStripeCoupon(Id quoteCouponId) {
        return [
            SELECT Id, Name__c, Percent_Off__c, Duration__c, Quote_Stripe_Coupon_Id__c, Order__c
            FROM Order_Stripe_Coupon__c
            WHERE Quote_Stripe_Coupon_Id__c = :quoteCouponId];
    }
}
