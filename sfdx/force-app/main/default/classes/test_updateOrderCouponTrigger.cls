@isTest
public with sharing class test_updateOrderCouponTrigger {
    
    @isTest
    static public void testUpdateOrderCouponTrigger() {
        Boolean isCpqInstalled = utilities.isCpqEnabled();
        if(!isCpqInstalled) {
            return;
        }

        // Setup test Data
        SBQQ__Quote__c quote = generateQuote();
        Stripe_Coupon__c coupon = createStripeCoupon();
        createStripeCouponQuoteAssociation(quote.id, coupon.id);

        Test.startTest();
            // insert an ordered quote
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
        Test.stopTest();

        // Verify
        List<Order> orders = [
            SELECT Id
            FROM Order
            WHERE SBQQ__Quote__c = :quote.id];
        System.assertEquals(1, orders.size());

        // trigger should have fired and created a serialized Stripe Coupon
        List<Stripe_Coupon_Serialized__c> serializedStripeCoupons = getStripeCouponSerialized(coupon.Id);
        System.assertEquals(1, serializedStripeCoupons.size());
        
        Stripe_Coupon_Serialized__c serializedStripeCoupon = serializedStripeCoupons.get(0);
        System.assertEquals('55 percent off coupon', serializedStripeCoupon.Name__c);

        // should have create a Stripe Coupon Order Association__c as well
        List<Stripe_Coupon_Order_Association__c> stripeCouponOrderAssociations = getstripeCouponOrderAssociations(orders.get(0).id);
        System.assertEquals(1, stripeCouponOrderAssociations.size());
    }

    static public SBQQ__Quote__c generateQuote() {
        Id pricebookId = Test.getStandardPricebookId();

        Account account = new Account();
        account.name = 'SFDC Account';
        account.SBQQ__RenewalModel__c = 'Contract Based';
        account.SBQQ__ContractCoTermination__c = 'Never';
        account.SBQQ__CoTerminationEvent__c = 'Renewal';
        account.SBQQ__RenewalPricingMethod__c = 'Same';
        account.SBQQ__PriceHoldEnd__c = system.today().AddMonths(1);
        account.SBQQ__AssetQuantitiesCombined__c = false;
        account.SBQQ__CoTermedContractsCombined__c = false;
        account.SBQQ__PreserveBundle__c = true;
        Database.insertImmediate((sObject) account);

        Product2 product = new Product2();
        product.Name = 'testProduct';
        product.Description = 'testProductEntry';
        product.productCode = 'ABC';
        product.SBQQ__DefaultQuantity__c = 12.000;
        product.isActive = true;
        Database.insertImmediate((sObject)product);

        SBQQ__Cost__c cost = new SBQQ__Cost__c();
        cost.SBQQ__Active__c = true;
        cost.SBQQ__Product__c = product.id;
        cost.SBQQ__UnitCost__c = 10.00;
        Database.insertImmediate((sObject)cost);
        
        sObject standardPrice = Schema.getGlobalDescribe().get('PricebookEntry').newSObject();
        standardPrice.put('Pricebook2Id',pricebookId);
        standardPrice.put('Product2Id',product.Id);
        standardPrice.put('UnitPrice',1);
        standardPrice.put('isActive',true);
        standardPrice.put('UseStandardPrice',false);
        Database.insertImmediate(standardPrice);

        Opportunity opportunity = new Opportunity();
        opportunity.put('Name', 'testOpportunity');
        opportunity.put('AccountId', account.id);
        opportunity.put('StageName','prospecting');
        opportunity.put('CloseDate', system.today() + 4);
        opportunity.put('Pricebook2Id', pricebookId);
        opportunity.put('SBQQ__Contracted__c', true);
        Database.insertImmediate((sObject)opportunity);

        SBQQ__Quote__c quote = new SBQQ__Quote__c();
        quote.SBQQ__Account__c = account.id;
        quote.SBQQ__Status__c = 'Approved';
        quote.SBQQ__StartDate__c = date.today();
        quote.SBQQ__PricebookId__c = pricebookId;
        quote.SBQQ__Opportunity2__c = opportunity.id;
        quote.SBQQ__SubscriptionTerm__c = 12;
        quote.SBQQ__ContractingMethod__c = 'By Subscription End Date';
        quote.SBQQ__Primary__c = true;
        Database.insertImmediate((sObject)quote);  

        SBQQ__QuoteLine__c quoteLine = new SBQQ__QuoteLine__c();
        quoteLine.SBQQ__Quote__c = quote.id;
        quoteLine.SBQQ__Product__c = product.id;
        quoteLine.SBQQ__PricebookEntryId__c = standardPrice.id;
        quoteLine.SBQQ__Quantity__c = 10; 
        Database.insertImmediate((sObject)quoteLine);

        return (SBQQ__Quote__c) quote; 
    }

    static public Stripe_Coupon__c createStripeCoupon() {
        Stripe_Coupon__c salesforceStripeCoupon =  new Stripe_Coupon__c();
        salesforceStripeCoupon.Name__c = '55 percent off coupon';
        salesforceStripeCoupon.Percent_Off__c = 55;
        salesforceStripeCoupon.Duration__c = 'once';
        Database.insertImmediate((sObject) salesforceStripeCoupon);
        return salesforceStripeCoupon;
    }

    static public void createStripeCouponQuoteAssociation(String quoteId, String couponId) {
        Stripe_Coupon_Quote_Association__c salesforceStripeCouponQuoteAssociation =  new Stripe_Coupon_Quote_Association__c();
        salesforceStripeCouponQuoteAssociation.Quote__c = quoteId;
        salesforceStripeCouponQuoteAssociation.Stripe_Coupon__c = couponId;
        Database.insertImmediate((sObject)salesforceStripeCouponQuoteAssociation);
    }

    static public List<Stripe_Coupon_Serialized__c> getStripeCouponSerialized(Id originalCouponId) {
        return [
            SELECT Id, Name__c 
            FROM Stripe_Coupon_Serialized__c
            WHERE Original_Stripe_Coupon_Id__c = :originalCouponId];
    }

    static public List<Stripe_Coupon_Order_Association__c> getStripeCouponOrderAssociations(String orderId) {
        return [
            SELECT Id
            FROM Stripe_Coupon_Order_Association__c
            WHERE Order__c = :orderId];  
    }
}
