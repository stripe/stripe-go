@isTest
public with sharing class test_restService {     
    @isTest 
    public static void test_restServiceEvents() {
        Boolean isCpqInstalled = utilities.isCpqEnabled();

        //insert all the records we will be checking the JSON response for
        Id pricebookId = Test.getStandardPricebookId();
        Map<String, Schema.SObjectType> accountObject = Schema.getGlobalDescribe();
        sObject account = accountObject.get('account').newSObject();
        account.put('name','SFDC Account');
        if(isCpqInstalled) {
            account.put('SBQQ__RenewalModel__c','Contract Based');
            account.put('SBQQ__ContractCoTermination__c','Never');
            account.put('SBQQ__CoTerminationEvent__c','Renewal');
            account.put('SBQQ__RenewalPricingMethod__c','Same');
            account.put('SBQQ__PriceHoldEnd__c', system.today().AddMonths(1));
            account.put('SBQQ__AssetQuantitiesCombined__c',false);
            account.put('SBQQ__CoTermedContractsCombined__c',false);
            account.put('SBQQ__PreserveBundle__c',true);
        }
        insert account;

        Map<String, Schema.SObjectType> opportunityObject = Schema.getGlobalDescribe();
        sObject opportunity = opportunityObject.get('Opportunity').newSObject();
        opportunity.put('Name', 'testOpportunity');
        opportunity.put('AccountId', account.id);
        opportunity.put('StageName','prospecting');
        opportunity.put('CloseDate', system.today() + 4);
        opportunity.put('Pricebook2Id', pricebookId);
        if(isCpqInstalled) {
            opportunity.put('SBQQ__Contracted__c', true);
        }
        insert opportunity;
        
        sObject quote;
        if(isCpqInstalled) {
            Map<String, Schema.SObjectType> quoteObject = Schema.getGlobalDescribe();
            quote = quoteObject.get('SBQQ__Quote__c').newSObject();
            quote.put('SBQQ__Account__c',account.id);
            quote.put('SBQQ__Status__c','Approved');
            quote.put('SBQQ__StartDate__c',date.today());
            quote.put('SBQQ__PricebookId__c',pricebookId);
            quote.put('SBQQ__Opportunity2__c',opportunity.id);
            quote.put('SBQQ__SubscriptionTerm__c',12);
            quote.put('SBQQ__ContractingMethod__c','By Subscription End Date');
            quote.put('SBQQ__Primary__c',true);
            quote.put('SBQQ__Ordered__c',true);
            insert quote;
            opportunity.put('SBQQ__PrimaryQuote__c', quote.Id);
            upsert opportunity;
        }

        sObject quoteLineGroup;
        if(isCpqInstalled) {
            Map<String, Schema.SObjectType> quoteLineGroupObject = Schema.getGlobalDescribe();
            quoteLineGroup = quoteLineGroupObject.get('SBQQ__QuoteLineGroup__c').newSObject();
            quoteLineGroup.put('SBQQ__Account__c',account.id);
            quoteLineGroup.put('SBQQ__Quote__c',quote.id);
            quoteLineGroup.put('SBQQ__StartDate__c',system.today());
            quoteLineGroup.put('SBQQ__BillingFrequency__c','Monthly');
            insert quoteLineGroup;
        }    
    
        Map<String, Schema.SObjectType> productObject = Schema.getGlobalDescribe();
        sObject product = productObject.get('Product2').newSObject();
        product.put('Name','testProduct');
        product.put('Description','testProductEntry');
        product.put('productCode','ABC');
        if(isCpqInstalled)product.put('SBQQ__DefaultQuantity__c',12.000);
        product.put('isActive',true);
        insert product;

        sObject cost;
        if(isCpqInstalled) {
            Map<String, Schema.SObjectType> costObject = Schema.getGlobalDescribe();
            cost = costObject.get('SBQQ__Cost__c').newSObject();
            cost.put('SBQQ__Active__c',true);
            cost.put('SBQQ__Product__c',product.id);
            cost.put('SBQQ__UnitCost__c',10.00);
            insert cost;
        }

        Map<String, Schema.SObjectType> standardPriceObject = Schema.getGlobalDescribe();
        sObject standardPrice = standardPriceObject.get('PricebookEntry').newSObject();
        standardPrice.put('Pricebook2Id',pricebookId);
        standardPrice.put('Product2Id',product.Id);
        standardPrice.put('UnitPrice',1);
        standardPrice.put('isActive',true);
        standardPrice.put('UseStandardPrice',false);
        insert standardPrice;

        sObject quoteLine;
        if(isCpqInstalled) {
            Map<String, Schema.SObjectType> quoteLineObject = Schema.getGlobalDescribe();
            quoteLine = quoteLineObject.get('SBQQ__QuoteLine__c').newSObject();
            quoteLine.put('SBQQ__Quote__c',quote.id);
            quoteLine.put('SBQQ__Product__c',product.id);
            quoteLine.put('SBQQ__PricebookEntryId__c',standardPrice.id);
            quoteLine.put('SBQQ__Quantity__c',10); 
            insert quoteLine;
        }
        
        Map<String, Schema.SObjectType> opportunityLineObject = Schema.getGlobalDescribe();
        sObject opportunityLine = opportunityLineObject.get('OpportunityLineItem').newSObject();
        opportunityLine.put('Product2Id',product.id);
        opportunityLine.put('PricebookEntryId',standardPrice.id);
        opportunityLine.put('OpportunityId',opportunity.id);
        opportunityLine.put('TotalPrice',100.00);
        opportunityLine.put('Quantity',10);
        if(isCpqInstalled)opportunityLine.put('SBQQ__QuoteLine__c',quoteLine.id);
        insert opportunityLine;

        Map<String, Schema.SObjectType> contractObject = Schema.getGlobalDescribe();
        sObject contract = contractObject.get('Contract').newSObject();
        contract.put('AccountId',account.id);
        contract.put('ContractTerm',12);
        contract.put('Status','Draft');
        contract.put('StartDate', system.today());
        if(isCpqInstalled) {
            contract.put('SBQQ__Quote__c',quote.id);
            contract.put('SBQQ__Opportunity__c',opportunity.id);
            contract.put('SBQQ__RenewalPricebookId__c',pricebookId);
            contract.put('SBQQ__AmendmentRenewalBehavior__c','Latest End Date');
            contract.put('SBQQ__AmendmentStartDate__c',system.today().AddMonths(1));
            contract.put('SBQQ__AmendmentPricebookId__c',pricebookId);
            contract.put('SBQQ__PreserveBundleStructureUponRenewals__c',true);
        }
        insert contract;

        opportunity.put('ContractId', contract.id);
        if(isCpqInstalled) {
            opportunity.put('SBQQ__PrimaryQuote__c', quote.id);
            opportunity.put('SBQQ__RenewedContract__c', contract.id);
        }
        upsert opportunity;

        Map<String, Schema.SObjectType> contactObject = Schema.getGlobalDescribe();
        sObject contact = contactObject.get('Contact').newSObject();
        contact.put('FirstName','testContact');
        contact.put('LastName','testContact');
        contact.put('Email','testContact@testContact.com');
        contact.put('AccountId',account.id);
        insert contact;

        sObject subscription;
        if(isCpqInstalled) {
            Map<String, Schema.SObjectType> subscriptionObject = Schema.getGlobalDescribe();
            subscription = subscriptionObject.get('SBQQ__Subscription__c').newSObject();
            subscription.put('SBQQ__Account__c',account.id);
            subscription.put('SBQQ__Contract__c',contract.id);
            subscription.put('SBQQ__Product__c',product.id);
            subscription.put('SBQQ__Quantity__c',10);
            subscription.put('SBQQ__QuoteLine__c',quoteLine.Id);
            insert subscription;
            quoteLine.put('SBQQ__RenewedSubscription__c',subscription.id);
            upsert subscription;
        }

        sObject contractedPrice;
        if(isCpqInstalled) {
            Map<String, Schema.SObjectType> contractedPriceObject = Schema.getGlobalDescribe();
            contractedPrice = contractedPriceObject.get('SBQQ__ContractedPrice__c').newSObject();
            contractedPrice.put('SBQQ__Account__c', account.Id);
            contractedPrice.put('SBQQ__Contract__c',contract.Id);
            contractedPrice.put('SBQQ__Description__c','Test Contract Price');
            contractedPrice.put('SBQQ__EffectiveDate__c',system.today());
            contractedPrice.put('SBQQ__ExpirationDate__c',system.today().AddMonths(12));
            contractedPrice.put('SBQQ__FilterField__c','Product');
            contractedPrice.put('SBQQ__NonDiscountable__c',false);
            contractedPrice.put('SBQQ__Operator__c','	greater or equals');
            contractedPrice.put('SBQQ__OriginalQuoteLine__c',quoteLine.Id);
            contractedPrice.put('SBQQ__Price__c',100.00);

            insert contractedPrice;
        }
        
        //Contract Ammendment 
        Map<String, Schema.SObjectType> amendedOpportunityObject = Schema.getGlobalDescribe();
        sObject amendedOpportunity = amendedOpportunityObject.get('Opportunity').newSObject();
        amendedOpportunity.put('Name', 'Amendment for contract #' + contract.get('Id'));
        amendedOpportunity.put('AccountId', account.id);
        amendedOpportunity.put('StageName','prospecting');
        amendedOpportunity.put('CloseDate', contract.get('StartDate'));
        amendedOpportunity.put('Pricebook2Id', pricebookId);
        if(isCpqInstalled) {
            amendedOpportunity.put('SBQQ__Contracted__c', true);
        }
        insert amendedOpportunity;

        sObject amendedQuote;
        sObject amendedQuoteLine;
        if(isCpqInstalled) {
            Map<String, Schema.SObjectType> amendedQuoteObject = Schema.getGlobalDescribe();
            amendedQuote = amendedQuoteObject.get('SBQQ__Quote__c').newSObject();
            amendedQuote.put('SBQQ__Account__c',account.id);
            amendedQuote.put('SBQQ__Status__c','Approved');
            amendedQuote.put('SBQQ__StartDate__c',contract.get('StartDate'));
            amendedQuote.put('SBQQ__EndDate__c',contract.get('SBQQ__ExpirationDate__c'));
            amendedQuote.put('SBQQ__PricebookId__c',pricebookId);
            amendedQuote.put('SBQQ__Opportunity2__c',opportunity.id);
            amendedQuote.put('SBQQ__SubscriptionTerm__c',12);
            amendedQuote.put('SBQQ__ContractingMethod__c','By Subscription End Date');
            amendedQuote.put('SBQQ__Primary__c',true);
            amendedQuote.put('SBQQ__Ordered__c',true);
            insert amendedQuote;
            opportunity.put('SBQQ__PrimaryQuote__c', amendedQuote.get('Id'));
            upsert opportunity;

            Map<String, Schema.SObjectType> amendedQuoteLineObject = Schema.getGlobalDescribe();
            amendedQuoteLine = amendedQuoteLineObject.get('SBQQ__QuoteLine__c').newSObject();
            amendedQuoteLine.put('SBQQ__Quote__c',amendedQuote.id);
            amendedQuoteLine.put('SBQQ__Product__c',product.id);
            amendedQuoteLine.put('SBQQ__PricebookEntryId__c',standardPrice.id);
            amendedQuoteLine.put('SBQQ__Quantity__c',5); 
            insert amendedQuoteLine;

            subscription.put('SBQQ__QuoteLine__c',amendedQuoteLine.Id);
            upsert subscription;

            contractedPrice.put('SBQQ__OriginalQuoteLine__c',amendedQuoteLine.Id);
            upsert contractedPrice;
        }
        //End Contract Ammendment

        Map<String, Schema.SObjectType> orderObject = Schema.getGlobalDescribe();
        sObject order = orderObject.get('Order').newSObject();
        order.put('Name','testOrder');
        order.put('AccountId',account.id);
        order.put('BillToContactId',contact.id);
        order.put('ContractId',contract.id);
        order.put('OpportunityId',opportunity.id);
        order.put('status','Draft');
        order.put('EffectiveDate',system.today());
        order.put('EndDate',system.today() + 40);
        order.put('Pricebook2Id',pricebookId);
        if(isCpqInstalled)order.put('SBQQ__Quote__c',amendedQuote.id);
        insert order;

        Map<String, Schema.SObjectType> order2Object = Schema.getGlobalDescribe();
        sObject order2 = order2Object.get('Order').newSObject();
        order2.put('Name','testOrder2');
        order2.put('AccountId',account.id);
        order2.put('BillToContactId',contact.id);
        order2.put('ContractId',contract.id);
        order2.put('OpportunityId',opportunity.id);
        order2.put('status','Draft');
        order2.put('EffectiveDate',system.today());
        order2.put('EndDate',system.today() + 40);
        order2.put('Pricebook2Id',pricebookId);
        if(isCpqInstalled)order2.put('SBQQ__Quote__c',amendedQuote.id);
        insert order2;

        Map<String, Schema.SObjectType> orderItemObject = Schema.getGlobalDescribe();
        sObject orderItem = orderItemObject.get('OrderItem').newSObject();
        orderItem.put('OrderId',order.id);
        orderItem.put('Quantity',24);
        orderItem.put('UnitPrice',240);
        orderItem.put('Product2id',product.id);
        orderItem.put('PricebookEntryId',standardPrice.id);
        orderItem.put('ServiceDate',system.today() + 4);
        orderItem.put('EndDate',system.today() + 14);
        insert orderItem;

        Map<String, Schema.SObjectType> orderItem2Object = Schema.getGlobalDescribe();
        sObject orderItem2 = orderItem2Object.get('OrderItem').newSObject();
        orderItem2.put('OrderId',order2.id);
        orderItem2.put('Quantity',24);
        orderItem2.put('UnitPrice',240);
        orderItem2.put('Product2id',product.id);
        orderItem2.put('PricebookEntryId',standardPrice.id);
        orderItem2.put('ServiceDate',system.today() + 4);
        orderItem2.put('EndDate',system.today() + 14);
        insert orderItem2;

        contract.put('Status','Activated');
        upsert contract;

        orderItem2.put('Quantity',80);
        orderItem2.put('UnitPrice', 500);
        upsert orderItem2;  
        
        //make mock request
        List<String> listOfOrderIds = new List<String>{(String)order.id};
        Map<String,Object> requestOrderMap = new Map<String,Object> {
            'order_ids' => (Object)listOfOrderIds
            };
        RestRequest request = new RestRequest();
        request.requestBody = Blob.valueOf(JSON.serialize(requestOrderMap));
        request.httpMethod = 'Get';
        RestContext.request = request;
        RestContext.response = new RestResponse();
        RestResponse response = RestContext.response;

        restService.restServiceEvents();

        Map<String, Object> parsedReqBody = (Map<String, Object>)JSON.deserializeUntyped(response.responseBody.toString());
        List<Object> records = (List<Object>)parsedReqBody.get('records');
        System.debug(JSON.serialize(parsedReqBody.get('records')));
        
        for(Object record : records) {
            Map<String, Object> recordMap = (Map<String, Object>)record;
            if(recordMap.get('attributes') != null) {
                
                Map<String, Object> attributes = (Map<String, Object>)recordMap.get('attributes');
                String objectType = (String)attributes.get('type');
                
                switch on objectType {
                    when 'Contract' {
                        System.assertEquals((String)recordMap.get('Id'), (String)contract.id);
                        System.assertEquals(recordMap.get('ContractTerm'), 12);
                        System.assertEquals((String)recordMap.get('StartDate'), String.valueOf(system.today()));
                        System.assertEquals((String)recordMap.get('AccountId'), (String)account.id);
                        if(isCpqInstalled) {
                            System.assertEquals((String)recordMap.get('SBQQ__RenewalPricebookId__c'), pricebookId);
                            System.assertEquals((String)recordMap.get('SBQQ__AmendmentPricebookId__c'), pricebookId);
                            System.assertEquals((String)recordMap.get('SBQQ__Quote__c'), (String)quote.id);
                            System.assertEquals((String)recordMap.get('SBQQ__Opportunity__c'), (String)opportunity.id);
                        }
                    }
                    when 'Account' {
                        System.assertEquals((String)recordMap.get('Id'), (String)account.id);
                        System.assertEquals((String)recordMap.get('Name'), 'SFDC Account');
                        if(isCpqInstalled) System.assertEquals((String)recordMap.get('SBQQ__RenewalModel__c'), 'Contract Based');
                    }
                    when 'Contact' {
                        System.assertEquals((String)recordMap.get('Id'), (String)contact.id);
                        System.assertEquals((String)recordMap.get('FirstName'), 'testContact');
                        System.assertEquals((String)recordMap.get('LastName'), 'testContact');
                        System.assertEquals((String)recordMap.get('Email'),'testcontact@testcontact.com');
                        System.assertEquals((String)recordMap.get('AccountId'), (String)account.id);
                    }
                    when 'Product2' {
                        System.assertEquals((String)recordMap.get('Id'), (String)product.id);
                        System.assertEquals((String)recordMap.get('Name'), 'testProduct');
                        System.assertEquals((String)recordMap.get('Description'), 'testProductEntry');
                    }
                    when 'PricebookEntry' {
                        System.assertEquals((String)recordMap.get('Id'), (String)standardPrice.id);
                        System.assertEquals((String)recordMap.get('Pricebook2Id'), (String)pricebookId);
                        System.assertEquals((String)recordMap.get('Product2Id'), (String)product.id);
                        System.assertEquals(recordMap.get('UnitPrice'), 1);
                    }
                    when 'OrderItem' {
                        if((String)recordMap.get('Id') == orderItem.id) {
                            System.assertEquals((String)recordMap.get('Id'), (String)orderItem.id);
                            System.assertEquals((String)recordMap.get('OrderId'), (String)order.id);
                            System.assertEquals((String)recordMap.get('PricebookEntryId'), (String)standardPrice.id);
                            System.assertEquals(recordMap.get('Quantity'), 24);
                            System.assertEquals(recordMap.get('UnitPrice'), 240);
                            System.assertEquals((String)recordMap.get('ServiceDate'), String.valueOf(system.today() + 4));
                            System.assertEquals((String)recordMap.get('EndDate'), String.valueOf(system.today() + 14));
                            List<Object> Products = (List<Object>)recordMap.get('Products');
                            List<Object> PricebookEntries = (List<Object>)recordMap.get('PricebookEntries');
                            System.assertEquals((String)Products[0], (String)product.id);
                            System.assertEquals((String)PricebookEntries[0], (String)standardPrice.id);
                        } else {
                            System.assertEquals((String)recordMap.get('Id'), (String)orderItem2.id);
                            System.assertEquals((String)recordMap.get('OrderId'), (String)order2.id);
                            System.assertEquals((String)recordMap.get('PricebookEntryId'), (String)standardPrice.id);
                            //checking that we get the updated quantity and price 
                            System.assertEquals(recordMap.get('Quantity'), 80);
                            System.assertEquals(recordMap.get('UnitPrice'), 500);
                            System.assertEquals((String)recordMap.get('ServiceDate'), String.valueOf(system.today() + 4));
                            System.assertEquals((String)recordMap.get('EndDate'), String.valueOf(system.today() + 14));
                            List<Object> Products = (List<Object>)recordMap.get('Products');
                            List<Object> PricebookEntries = (List<Object>)recordMap.get('PricebookEntries');
                            System.assertEquals((String)Products[0], (String)product.id);
                            System.assertEquals((String)PricebookEntries[0], (String)standardPrice.id);
                        }
                    }
                    when 'Order' {
                        if((String)recordMap.get('Id') == order.id) {
                            System.assertEquals((String)recordMap.get('Id'), (String)order.id);
                            System.assertEquals((String)recordMap.get('AccountId'), (String)account.id);
                            System.assertEquals((String)recordMap.get('BillToContactId'), (String)contact.id);
                            System.assertEquals((String)recordMap.get('ContractId'), (String)contract.id);
                            System.assertEquals((String)recordMap.get('Pricebook2Id'), (String)pricebookId);
                            System.assertEquals((String)recordMap.get('EffectiveDate'), String.valueOf(system.today()));
                            System.assertEquals((String)recordMap.get('EndDate'), String.valueOf(system.today() + 40));
                            System.assertEquals((String)recordMap.get('Name'), 'testOrder');
                            List<Object> Contracts = (List<Object>)recordMap.get('Contracts');
                            List<Object> Accounts = (List<Object>)recordMap.get('Accounts');
                            List<Object> Contacts = (List<Object>)recordMap.get('Contacts');
                            List<Object> OrderItems = (List<Object>)recordMap.get('OrderItems');
                            List<Object> ContractOrders = (List<Object>)recordMap.get('ContractOrders');
                            System.assertEquals((String)Contracts[0], (String)contract.id);
                            System.assertEquals((String)Accounts[0], (String)account.id);
                            System.assertEquals((String)Contacts[0], (String)contact.id);
                            System.assertEquals((String)OrderItems[0], (String)orderItem.id);
                            System.assertEquals((String)ContractOrders[0], (String)order2.id);
                        } else {
                            System.assertEquals((String)recordMap.get('Id'), (String)order2.id);
                            System.assertEquals((String)recordMap.get('AccountId'), (String)account.id);
                            System.assertEquals((String)recordMap.get('BillToContactId'), (String)contact.id);
                            System.assertEquals((String)recordMap.get('ContractId'), (String)contract.id);
                            System.assertEquals((String)recordMap.get('Pricebook2Id'), (String)pricebookId);
                            System.assertEquals((String)recordMap.get('EffectiveDate'), String.valueOf(system.today()));
                            System.assertEquals((String)recordMap.get('EndDate'), String.valueOf(system.today() + 40));
                            System.assertEquals((String)recordMap.get('Name'), 'testOrder2');
                            List<Object> Contracts = (List<Object>)recordMap.get('Contracts');
                            List<Object> Accounts = (List<Object>)recordMap.get('Accounts');
                            List<Object> Contacts = (List<Object>)recordMap.get('Contacts');
                            List<Object> OrderItems = (List<Object>)recordMap.get('OrderItems');
                            System.assertEquals((String)Contracts[0], (String)contract.id);
                            System.assertEquals((String)Accounts[0], (String)account.id);
                            System.assertEquals((String)Contacts[0], (String)contact.id);
                            System.assertEquals((String)OrderItems[0], (String)orderItem2.id);
                        }
                    }
                    when 'Opportunity' {
                        System.assertEquals((String)recordMap.get('Id'), (String)opportunity.id);
                        System.assertEquals((String)recordMap.get('AccountId'), (String)account.id);
                        System.assertEquals((String)recordMap.get('Pricebook2Id'), (String)pricebookId);
                        System.assertEquals((String)recordMap.get('StageName'), 'Prospecting');
                        System.assertEquals((String)recordMap.get('Name'), 'testOpportunity');
                        System.assertEquals((String)recordMap.get('CloseDate'), String.valueOf(system.today() + 4));
                    }
                }
            } 
        }
    }
}