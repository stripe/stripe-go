@isTest
public with sharing class test_restService {   
    
    @isTest 
    public static void test_restServiceEventsAmendmentOrder() {

        Map<String,Object> testRecordsMap = makeData(true);
        Boolean isCpqInstalled = (Boolean)testRecordsMap.get('isCpqInstalled');
        Id pricebookId = (Id)testRecordsMap.get('pricebookId');
        Account account = (Account)testRecordsMap.get('account');
        Opportunity opportunity = (Opportunity)testRecordsMap.get('opportunity');
        Opportunity amendedOpportunity = (Opportunity)testRecordsMap.get('amendedOpportunity');
        sObject quote = (sObject)testRecordsMap.get('quote');
        sObject amendedQuote = (sObject)testRecordsMap.get('amendedQuote');
        Product2 product = (Product2)testRecordsMap.get('product');
        ConsumptionSchedule consumptionSchedule = (ConsumptionSchedule)testRecordsMap.get('consumptionSchedule');
        ConsumptionRate consumptionRate = (ConsumptionRate)testRecordsMap.get('consumptionRate');
        ProductConsumptionSchedule productConsumptionSchedule = (ProductConsumptionSchedule)testRecordsMap.get('productConsumptionSchedule');
        Contract contract = (Contract)testRecordsMap.get('contract');
        Order order = (Order)testRecordsMap.get('order');
        Contract amendedContract = (Contract)testRecordsMap.get('amendedContract');
        Order amendedOrder = (Order)testRecordsMap.get('amendedOrder');
        OrderItem orderItem = (OrderItem)testRecordsMap.get('orderItem');
        OrderItem orderItem2 = (OrderItem)testRecordsMap.get('orderItem2');
        PricebookEntry standardPrice = (PricebookEntry)testRecordsMap.get('standardPrice');
        sObject orderItemConsumptionRate = (sObject)testRecordsMap.get('orderItemConsumptionRate');
        sObject orderItemConsumptionSchedule = (sObject)testRecordsMap.get('orderItemConsumptionSchedule');
        Contact contact = (Contact)testRecordsMap.get('contact');


        Test.startTest();
        
        //make mock request
        List<String> listOfOrderIds = new List<String>{(String)order.id};
        List<String> listOfProductIds = new List<String>{(String)product.id};
        Map<String,Object> requestOrderMap = new Map<String,Object> {
            'order_ids' => (Object)listOfOrderIds,
            'product2_ids' => (Object)listOfProductIds
            };
        RestRequest request = new RestRequest();
        request.requestBody = Blob.valueOf(JSON.serialize(requestOrderMap));
        request.httpMethod = 'Post';
        RestContext.request = request;
        RestContext.response = new RestResponse();
        RestResponse response = RestContext.response;

        restService.executeBatchApi();
        Map<String, Object> parsedReqBody = (Map<String, Object>)JSON.deserializeUntyped(response.responseBody.toString());
        System.debug('AmendedOrder ' + JSON.serialize(parsedReqBody));
    
        Test.stopTest();

        Set<String> allObjectsToPullDataFrom = new Set<String>();
        List<Object> records = (List<Object>)parsedReqBody.get('records');
        
        for(Object record : records) {
            Map<String, Object> recordMap = (Map<String, Object>)record;
            if(recordMap.get('attributes') != null) {              
                Map<String, Object> attributes = (Map<String, Object>)recordMap.get('attributes');
                String objectType = (String)attributes.get('type');
                
                switch on objectType {
                    when 'Contract' {
                        allObjectsToPullDataFrom.add('Contract');
                        if (recordMap.get('Id') == (String)contract.id) {
                            System.assertEquals((String)recordMap.get('Id'), (String)contract.id);
                            System.assertEquals(recordMap.get('ContractTerm'), 12);
                            System.assertEquals((String)recordMap.get('StartDate'), String.valueOf(system.today()));
                            System.assertEquals((String)recordMap.get('AccountId'), (String)account.id);
                            if(isCpqInstalled) {
                                System.assertEquals((String)recordMap.get('SBQQ__RenewalPricebookId__c'), pricebookId);
                                System.assertEquals((String)recordMap.get('SBQQ__AmendmentPricebookId__c'), pricebookId);
                                System.assertEquals((String)recordMap.get('SBQQ__Opportunity__c'), (String)opportunity.id);
                            }
                        }
                    }
                    when 'Account' {
                        allObjectsToPullDataFrom.add('Account');
                        System.assertEquals((String)recordMap.get('Id'), (String)account.id);
                        System.assertEquals((String)recordMap.get('Name'), 'SFDC Account');
                        if(isCpqInstalled) System.assertEquals((String)recordMap.get('SBQQ__RenewalModel__c'), 'Contract Based');
                    }
                    when 'Contact' {
                        allObjectsToPullDataFrom.add('Contact');
                        System.assertEquals((String)recordMap.get('Id'), (String)contact.id);
                        System.assertEquals((String)recordMap.get('FirstName'), 'testContact');
                        System.assertEquals((String)recordMap.get('LastName'), 'testContact');
                        System.assertEquals((String)recordMap.get('Email'),'testcontact@testcontact.com');
                        System.assertEquals((String)recordMap.get('AccountId'), (String)account.id);
                    }
                    when 'Product2' {
                        allObjectsToPullDataFrom.add('Product2');
                        System.assertEquals((String)recordMap.get('Id'), (String)product.id);
                        System.assertEquals((String)recordMap.get('Name'), 'testProduct');
                        System.assertEquals((String)recordMap.get('Description'), 'testProductEntry');
                        if (recordMap.get('ProductConsumptionSchedules') != null ) {
                            List<Object> ProductConsumptionSchedulesList = (List<Object>)recordMap.get('ProductConsumptionSchedules');
                            List<Object> ConsumptionSchedulesList = (List<Object>)recordMap.get('ConsumptionSchedules');
                            System.assertEquals(ProductConsumptionSchedulesList[0], productConsumptionSchedule.Id);
                            System.assertEquals(ConsumptionSchedulesList[0], consumptionSchedule.Id);
                        }
                    }
                    when 'PricebookEntry' {
                        allObjectsToPullDataFrom.add('PricebookEntry');
                        System.assertEquals((String)recordMap.get('Id'), (String)standardPrice.id);
                        System.assertEquals((String)recordMap.get('Pricebook2Id'), (String)pricebookId);
                        System.assertEquals((String)recordMap.get('Product2Id'), (String)product.id);
                        System.assertEquals(recordMap.get('UnitPrice'), 1);
                    }
                    when 'OrderItem' {
                        allObjectsToPullDataFrom.add('OrderItem');
                        if((String)recordMap.get('Id') == orderItem.id) {
                            System.assertEquals((String)recordMap.get('Id'), (String)orderItem.id);
                            System.assertEquals((String)recordMap.get('OrderId'), (String)order.id);
                            System.assertEquals((String)recordMap.get('PricebookEntryId'), (String)standardPrice.id);
                            System.assertEquals(recordMap.get('Quantity'), 24);
                            System.assertEquals(recordMap.get('UnitPrice'), 240);
                            System.assertEquals((String)recordMap.get('ServiceDate'), String.valueOf(system.today() + 4));
                            System.assertEquals((String)recordMap.get('EndDate'), String.valueOf(system.today() + 14));
                            System.assertEquals((String)recordMap.get('Product2Id'), (String)product.id);
                            if (isCpqInstalled && recordMap.get('OrderItemConsumptionSchedules') != null ) {
                                List<Object> orderItemConsumptionSchedulesList = (List<Object>)recordMap.get('OrderItemConsumptionSchedules');
                                List<Object> ConsumptionSchedulesList = (List<Object>)recordMap.get('ConsumptionSchedules');
                                System.assertEquals(orderItemConsumptionSchedulesList[0], orderItemConsumptionSchedule.Id);
                            }
                        } else {
                            System.assertEquals((String)recordMap.get('Id'), (String)orderItem2.id);
                            System.assertEquals((String)recordMap.get('OrderId'), (String)amendedOrder.id);
                            System.assertEquals((String)recordMap.get('PricebookEntryId'), (String)standardPrice.id);
                            //checking that we get the updated quantity and price 
                            System.assertEquals(recordMap.get('Quantity'), 80);
                            System.assertEquals(recordMap.get('UnitPrice'), 500);
                            System.assertEquals((String)recordMap.get('ServiceDate'), String.valueOf(system.today() + 4));
                            System.assertEquals((String)recordMap.get('EndDate'), String.valueOf(system.today() + 14));
                            List<Object> Products = (List<Object>)recordMap.get('Products');
                            List<Object> PricebookEntries = (List<Object>)recordMap.get('PricebookEntries');
                        }
                    }
                    when 'Order' {
                        if(((String)recordMap.get('Id') == order.id)) {
                            allObjectsToPullDataFrom.add('Order');
                            System.assertEquals((String)recordMap.get('Id'), (String)order.id);
                            System.assertEquals((String)recordMap.get('AccountId'), (String)account.id);
                            System.assertEquals((String)recordMap.get('BillToContactId'), (String)contact.id);
                            System.assertEquals((String)recordMap.get('ContractId'), (String)contract.id);
                            System.assertEquals((String)recordMap.get('Pricebook2Id'), (String)pricebookId);
                            System.assertEquals((String)recordMap.get('EffectiveDate'), String.valueOf(system.today()));
                            System.assertEquals((String)recordMap.get('EndDate'), String.valueOf(system.today() + 40));
                            System.assertEquals((String)recordMap.get('Name'), 'testOrder');
                            List<Object> Contracts = (List<Object>)recordMap.get('Contracts');
                            List<Object> OrderItems = (List<Object>)recordMap.get('OrderItems');
                            List<Object> AmendedContracts = (List<Object>)recordMap.get('AmendedContracts');
                            List<Object> AmendedOrders = (List<Object>)recordMap.get('AmendedOrders');
                            List<Object> Accounts = (List<Object>)recordMap.get('Accounts');
                            List<Object> Contacts = (List<Object>)recordMap.get('Contacts');
                            List<Object> Quotes = (List<Object>)recordMap.get('Quotes');
                            System.assertEquals((String)Accounts[0], (String)account.id);
                            System.assertEquals((String)Contacts[0], (String)contact.id);
                            if(isCpqInstalled) {
                                System.assertNotEquals((String)Quotes[0], null);
                                System.assertEquals((String)OrderItems[0], (String)orderItem.id);
                                System.assertEquals((String)AmendedOrders[0], (String)amendedOrder.id);
                                System.assertEquals((String)AmendedContracts[0], (String)amendedContract.id);
                            } 
                        } else if((String)recordMap.get('Id') == amendedOrder.id) {
                            allObjectsToPullDataFrom.add('amendedOrder');
                            System.assertEquals((String)recordMap.get('Id'), (String)amendedOrder.id);
                            System.assertEquals((String)recordMap.get('AccountId'), (String)account.id);
                            System.assertEquals((String)recordMap.get('BillToContactId'), (String)contact.id);
                            System.assertEquals((String)recordMap.get('ContractId'), (String)amendedContract.id);
                            System.assertEquals((String)recordMap.get('Pricebook2Id'), (String)pricebookId);
                            System.assertEquals((String)recordMap.get('EffectiveDate'), String.valueOf(system.today()));
                            System.assertEquals((String)recordMap.get('EndDate'), String.valueOf(system.today() + 40));
                            System.assertEquals((String)recordMap.get('Name'), 'testamendedOrder');
                            List<Object> OrderItems = (List<Object>)recordMap.get('OrderItems');
                            List<Object> Contracts = (List<Object>)recordMap.get('Contracts');
                            List<Object> Accounts = (List<Object>)recordMap.get('Accounts');
                            List<Object> Contacts = (List<Object>)recordMap.get('Contacts');
                            System.assertEquals((String)OrderItems[0], (String)orderItem2.id);
                            System.assertEquals((String)Contracts[0], (String)amendedContract.id);
                            System.assertEquals((String)Accounts[0], (String)account.id);
                            System.assertEquals((String)Contacts[0], (String)contact.id);
                            System.assertEquals((String)OrderItems[0], (String)orderItem2.id);
                        }
                    }
                    when 'Opportunity' {
                        allObjectsToPullDataFrom.add('Opportunity');
                        if(((String)recordMap.get('Id') == opportunity.id)) {
                            System.assertEquals((String)recordMap.get('Id'), (String)opportunity.id);
                            System.assertEquals((String)recordMap.get('AccountId'), (String)account.id);
                            System.assertEquals((String)recordMap.get('Pricebook2Id'), (String)pricebookId);
                            System.assertEquals((String)recordMap.get('StageName'), 'Prospecting');
                            System.assertEquals((String)recordMap.get('Name'), 'testOpportunity');
                            System.assertEquals((String)recordMap.get('CloseDate'), String.valueOf(system.today() + 4));
                        } else {
                            System.assertEquals((String)recordMap.get('Id'), (String)amendedOpportunity.id);
                            System.assertEquals((String)recordMap.get('AccountId'), (String)account.id);
                            System.assertEquals((String)recordMap.get('Pricebook2Id'), (String)pricebookId);
                            System.assertEquals((String)recordMap.get('StageName'), 'Prospecting');
                            System.assertEquals((String)recordMap.get('Name'), 'testOpportunity');
                            System.assertEquals((String)recordMap.get('CloseDate'), String.valueOf(system.today() + 4));
                        }
                    } when 'ProductConsumptionSchedule' {
                        allObjectsToPullDataFrom.add('ProductConsumptionSchedule');
                        System.assertEquals((String)recordMap.get('Id'), (String)productConsumptionSchedule.id);
                        if (recordMap.get('ConsumptionScheduleId') != null ) {
                            System.assertEquals((String)recordMap.get('ConsumptionScheduleId'), consumptionSchedule.Id);
                        }
                    } when 'ConsumptionSchedule' {
                        allObjectsToPullDataFrom.add('ConsumptionSchedule');
                        System.assertEquals((String)recordMap.get('Id'), (String)consumptionSchedule.id);
                        if (recordMap.get('ConsumptionRates') != null ) {
                            List<Object> consumptionRatesList = (List<Object>)recordMap.get('ConsumptionRates');
                            System.assertEquals(consumptionRatesList[0], consumptionRate.Id);
                        }
                    } when 'ConsumptionRate' {
                        allObjectsToPullDataFrom.add('ConsumptionRate');
                        System.assertEquals((String)recordMap.get('Id'), (String)consumptionRate.id);
                    } when 'SBQQ__OrderItemConsumptionSchedule__c' {
                        allObjectsToPullDataFrom.add('SBQQ__OrderItemConsumptionSchedule__c');
                        System.assertEquals((String)recordMap.get('Id'), (String)orderItemConsumptionSchedule.id);
                        if (recordMap.get('OrderItemConsumptionRates') != null ) {
                            List<Object> orderItemConsumptionRatesList = (List<Object>)recordMap.get('OrderItemConsumptionRates');
                            System.assertEquals(orderItemConsumptionRatesList[0], orderItemConsumptionRate.Id);
                        }
                    } when 'SBQQ__OrderItemConsumptionRate__c' {
                        allObjectsToPullDataFrom.add('SBQQ__OrderItemConsumptionRate__c');
                        System.assertEquals((String)recordMap.get('Id'), (String)orderItemConsumptionRate.id);
                    }
                }
            } 
        }

        Set<String> allObjects = new Set<String>{
            'Contract',
            'Account',
            'Contact',
            'Product2',
            'PricebookEntry',
            'OrderItem',
            'Order',
            'ProductConsumptionSchedule',
            'ConsumptionSchedule'
        };

        if(isCpqInstalled) {
            allObjects.add('Opportunity');
            allObjects.add('amendedOrder');
        }

        Set<String> listOfMissingObjectNames = new Set<String>();
        for (String objectName : allObjects) {
            if (!allObjectsToPullDataFrom.contains(objectName)) {
                listOfMissingObjectNames.add(objectName);
            }
        }

        if (!listOfMissingObjectNames.isEmpty()) {
            system.assertEquals(null, listOfMissingObjectNames);
        }
    }

    @isTest 
    public static void test_restServiceEventsInitialOrder() {
        Map<String,Object> testRecordsMap = makeData(false);
        Boolean isCpqInstalled = (Boolean)testRecordsMap.get('isCpqInstalled');
        Id pricebookId = (Id)testRecordsMap.get('pricebookId');
        Account account = (Account)testRecordsMap.get('account');
        Opportunity opportunity = (Opportunity)testRecordsMap.get('opportunity');
        Opportunity amendedOpportunity = (Opportunity)testRecordsMap.get('amendedOpportunity');
        sObject quote = (sObject)testRecordsMap.get('quote');
        sObject amendedQuote = (sObject)testRecordsMap.get('amendedQuote');
        Product2 product = (Product2)testRecordsMap.get('product');
        ConsumptionSchedule consumptionSchedule = (ConsumptionSchedule)testRecordsMap.get('consumptionSchedule');
        ConsumptionRate consumptionRate = (ConsumptionRate)testRecordsMap.get('consumptionRate');
        ProductConsumptionSchedule productConsumptionSchedule = (ProductConsumptionSchedule)testRecordsMap.get('productConsumptionSchedule');
        Contract contract = (Contract)testRecordsMap.get('contract');
        Order initalOrder = (Order)testRecordsMap.get('order');
        Contract amendedContract = (Contract)testRecordsMap.get('amendedContract');
        Order amendedOrder = (Order)testRecordsMap.get('amendedOrder');
        OrderItem orderItem = (OrderItem)testRecordsMap.get('orderItem');
        OrderItem orderItem2 = (OrderItem)testRecordsMap.get('orderItem2');
        PricebookEntry standardPrice = (PricebookEntry)testRecordsMap.get('standardPrice');
        sObject orderItemConsumptionRate = (sObject)testRecordsMap.get('orderItemConsumptionRate');
        sObject orderItemConsumptionSchedule = (sObject)testRecordsMap.get('orderItemConsumptionSchedule');
        Contact contact = (Contact)testRecordsMap.get('contact');

        Test.startTest();
        //make mock request
        List<String> listOfOrderIds = new List<String>{(String)amendedOrder.id};
        List<String> listOfProductIds = new List<String>{(String)product.id};
        Map<String,Object> requestOrderMap = new Map<String,Object> {
            'order_ids' => (Object)listOfOrderIds,
            'product2_ids' => (Object)listOfProductIds
            };
        RestRequest request = new RestRequest();
        request.requestBody = Blob.valueOf(JSON.serialize(requestOrderMap));
        request.httpMethod = 'Post';
        RestContext.request = request;
        RestContext.response = new RestResponse();
        RestResponse response = RestContext.response;

        restService.executeBatchApi();
        Map<String, Object> parsedReqBody = (Map<String, Object>)JSON.deserializeUntyped(response.responseBody.toString());
        System.debug('InitialOrder ' +JSON.serialize(parsedReqBody));
    
        Test.stopTest();
        
        List<Object> records = (List<Object>)parsedReqBody.get('records');
       
        Set<String> allObjectsToPullDataFrom = new Set<String>();
        
        for(Object record : records) {
            Map<String, Object> recordMap = (Map<String, Object>)record;
            if(recordMap.get('attributes') != null) {              
                Map<String, Object> attributes = (Map<String, Object>)recordMap.get('attributes');
                String objectType = (String)attributes.get('type');
                
                switch on objectType {
                    when 'Contract' {
                        allObjectsToPullDataFrom.add('Contract');
                        System.assertEquals(recordMap.get('ContractTerm'), 12);
                        System.assertEquals((String)recordMap.get('StartDate'), String.valueOf(system.today()));
                        System.assertEquals((String)recordMap.get('AccountId'), (String)account.id);
                        if(isCpqInstalled) {
                            System.assertEquals((String)recordMap.get('SBQQ__RenewalPricebookId__c'), pricebookId);
                            System.assertEquals((String)recordMap.get('SBQQ__AmendmentPricebookId__c'), pricebookId);
                            System.assertEquals((String)recordMap.get('SBQQ__Opportunity__c'), (String)opportunity.id);
                        }
                    }
                    when 'Account' {
                        allObjectsToPullDataFrom.add('Account');
                        System.assertEquals((String)recordMap.get('Id'), (String)account.id);
                        System.assertEquals((String)recordMap.get('Name'), 'SFDC Account');
                        if(isCpqInstalled) System.assertEquals((String)recordMap.get('SBQQ__RenewalModel__c'), 'Contract Based');
                    }
                    when 'Contact' {
                        allObjectsToPullDataFrom.add('Contact');
                        System.assertEquals((String)recordMap.get('Id'), (String)contact.id);
                        System.assertEquals((String)recordMap.get('FirstName'), 'testContact');
                        System.assertEquals((String)recordMap.get('LastName'), 'testContact');
                        System.assertEquals((String)recordMap.get('Email'),'testcontact@testcontact.com');
                        System.assertEquals((String)recordMap.get('AccountId'), (String)account.id);
                    }
                    when 'Product2' {
                        allObjectsToPullDataFrom.add('Product2');
                        System.assertEquals((String)recordMap.get('Id'), (String)product.id);
                        System.assertEquals((String)recordMap.get('Name'), 'testProduct');
                        System.assertEquals((String)recordMap.get('Description'), 'testProductEntry');
                        if (recordMap.get('ProductConsumptionSchedules') != null ) {
                            List<Object> ProductConsumptionSchedulesList = (List<Object>)recordMap.get('ProductConsumptionSchedules');
                            List<Object> ConsumptionSchedulesList = (List<Object>)recordMap.get('ConsumptionSchedules');
                            System.assertEquals(ProductConsumptionSchedulesList[0], productConsumptionSchedule.Id);
                            System.assertEquals(ConsumptionSchedulesList[0], consumptionSchedule.Id);
                        }
                    }
                    when 'PricebookEntry' {
                        allObjectsToPullDataFrom.add('PricebookEntry');
                        System.assertEquals((String)recordMap.get('Id'), (String)standardPrice.id);
                        System.assertEquals((String)recordMap.get('Pricebook2Id'), (String)pricebookId);
                        System.assertEquals((String)recordMap.get('Product2Id'), (String)product.id);
                        System.assertEquals(recordMap.get('UnitPrice'), 1);
                    }
                    when 'OrderItem' {
                        allObjectsToPullDataFrom.add('OrderItem');
                        if((String)recordMap.get('Id') == orderItem.id) {
                            System.assertEquals((String)recordMap.get('Id'), (String)orderItem.id);
                            System.assertEquals((String)recordMap.get('OrderId'), (String)initalOrder.id);
                            System.assertEquals((String)recordMap.get('PricebookEntryId'), (String)standardPrice.id);
                            System.assertEquals(recordMap.get('Quantity'), 24);
                            System.assertEquals(recordMap.get('UnitPrice'), 240);
                            System.assertEquals((String)recordMap.get('ServiceDate'), String.valueOf(system.today() + 4));
                            System.assertEquals((String)recordMap.get('EndDate'), String.valueOf(system.today() + 14));
                            System.assertEquals((String)recordMap.get('Product2Id'), (String)product.id);
                            if (isCpqInstalled && recordMap.get('OrderItemConsumptionSchedules') != null ) {
                                List<Object> orderItemConsumptionSchedulesList = (List<Object>)recordMap.get('OrderItemConsumptionSchedules');
                                List<Object> ConsumptionSchedulesList = (List<Object>)recordMap.get('ConsumptionSchedules');
                                System.assertEquals(orderItemConsumptionSchedulesList[0], orderItemConsumptionSchedule.Id);
                            }
                        } else {
                            System.assertEquals((String)recordMap.get('Id'), (String)orderItem2.id);
                            System.assertEquals((String)recordMap.get('OrderId'), (String)amendedOrder.id);
                            System.assertEquals((String)recordMap.get('PricebookEntryId'), (String)standardPrice.id);
                            //checking that we get the updated quantity and price 
                            System.assertEquals(recordMap.get('Quantity'), 80);
                            System.assertEquals(recordMap.get('UnitPrice'), 500);
                            System.assertEquals((String)recordMap.get('ServiceDate'), String.valueOf(system.today() + 4));
                            System.assertEquals((String)recordMap.get('EndDate'), String.valueOf(system.today() + 14));
                            List<Object> Products = (List<Object>)recordMap.get('Products');
                            List<Object> PricebookEntries = (List<Object>)recordMap.get('PricebookEntries');
                        }
                    }
                    when 'Order' {
                        allObjectsToPullDataFrom.add('initialOrder');
                        if(((String)recordMap.get('Id') == amendedOrder.id)) {
                            System.assertEquals((String)recordMap.get('Id'), (String)amendedOrder.id);
                            System.assertEquals((String)recordMap.get('AccountId'), (String)account.id);
                            System.assertEquals((String)recordMap.get('BillToContactId'), (String)contact.id);
                            System.assertEquals((String)recordMap.get('ContractId'), (String)amendedContract.id);
                            System.assertEquals((String)recordMap.get('Pricebook2Id'), (String)pricebookId);
                            System.assertEquals((String)recordMap.get('EffectiveDate'), String.valueOf(system.today()));
                            System.assertEquals((String)recordMap.get('EndDate'), String.valueOf(system.today() + 40));
                            System.assertEquals((String)recordMap.get('Name'), 'testamendedOrder');
                            List<Object> Contracts = (List<Object>)recordMap.get('Contracts');
                            List<Object> Quotes = (List<Object>)recordMap.get('Quotes');
                            List<Object> OrderItems = (List<Object>)recordMap.get('OrderItems');
                            List<Object> InitialOrderQuotes = (List<Object>)recordMap.get('InitialOrderQuotes');
                            List<Object> InitialOrderIds = (List<Object>)recordMap.get('InitialOrderId');
                            List<Object> Accounts = (List<Object>)recordMap.get('Accounts');
                            List<Object> Contacts = (List<Object>)recordMap.get('Contacts');
                            List<Object> Opportunities = (List<Object>)recordMap.get('Opportunities');
                            System.assertEquals((String)Accounts[0], (String)account.id);
                            System.assertEquals((String)Contacts[0], (String)contact.id);
                            if(isCpqInstalled) {
                                System.assertNotEquals((String)Quotes[0], null);
                                System.assertEquals((String)OrderItems[0], (String)orderItem2.id);
                                System.assertEquals((String)InitialOrderIds[0], (String)initalOrder.id);
                                System.assertNotEquals((String)InitialOrderQuotes[0], null);
                                System.assertEquals((String)Opportunities[0], (String)amendedOpportunity.id);
                            } 
                        } else if((String)recordMap.get('Id') == initalOrder.id) {
                            allObjectsToPullDataFrom.add('Order');
                            System.assertEquals((String)recordMap.get('Id'), (String)initalOrder.id);
                            System.assertEquals((String)recordMap.get('AccountId'), (String)account.id);
                            System.assertEquals((String)recordMap.get('BillToContactId'), (String)contact.id);
                            System.assertEquals((String)recordMap.get('ContractId'), (String)contract.id);
                            System.assertEquals((String)recordMap.get('Pricebook2Id'), (String)pricebookId);
                            System.assertEquals((String)recordMap.get('EffectiveDate'), String.valueOf(system.today()));
                            System.assertEquals((String)recordMap.get('EndDate'), String.valueOf(system.today() + 40));
                            System.assertEquals((String)recordMap.get('Name'), 'testOrder');
                            List<Object> OrderItems = (List<Object>)recordMap.get('OrderItems');
                            List<Object> Contracts = (List<Object>)recordMap.get('Contracts');
                            List<Object> Accounts = (List<Object>)recordMap.get('Accounts');
                            List<Object> Contacts = (List<Object>)recordMap.get('Contacts');
                            List<Object> Opportunities = (List<Object>)recordMap.get('Opportunities');
                            System.assertEquals((String)Contracts[0], (String)contract.id);
                            System.assertEquals((String)Accounts[0], (String)account.id);
                            System.assertEquals((String)Contacts[0], (String)contact.id);
                            System.assertEquals((String)OrderItems[0], (String)orderItem.id);
                            System.assertEquals((String)Opportunities[0], (String)opportunity.id);
                        }
                    }
                    when 'Opportunity' {
                        allObjectsToPullDataFrom.add('Opportunity');
                        if(((String)recordMap.get('Id') == opportunity.id)) {
                            System.assertEquals((String)recordMap.get('Id'), (String)opportunity.id);
                            System.assertEquals((String)recordMap.get('AccountId'), (String)account.id);
                            System.assertEquals((String)recordMap.get('Pricebook2Id'), (String)pricebookId);
                            System.assertEquals((String)recordMap.get('StageName'), 'Prospecting');
                            System.assertEquals((String)recordMap.get('Name'), 'testOpportunity');
                            System.assertEquals((String)recordMap.get('CloseDate'), String.valueOf(system.today() + 4));
                        } else {
                            System.assertEquals((String)recordMap.get('Id'), (String)amendedOpportunity.id);
                            System.assertEquals((String)recordMap.get('AccountId'), (String)account.id);
                            System.assertEquals((String)recordMap.get('Pricebook2Id'), (String)pricebookId);
                            System.assertEquals((String)recordMap.get('StageName'), 'Prospecting');
                            System.assertEquals((String)recordMap.get('Name'), 'testOpportunity');
                            System.assertEquals((String)recordMap.get('CloseDate'), String.valueOf(system.today() + 4));
                        }
                    } when 'ProductConsumptionSchedule' {
                        allObjectsToPullDataFrom.add('ProductConsumptionSchedule');
                        System.assertEquals((String)recordMap.get('Id'), (String)productConsumptionSchedule.id);
                        if (recordMap.get('ConsumptionScheduleId') != null ) {
                            System.assertEquals((String)recordMap.get('ConsumptionScheduleId'), consumptionSchedule.Id);
                        }
                    } when 'ConsumptionSchedule' {
                        allObjectsToPullDataFrom.add('ConsumptionSchedule');
                        System.assertEquals((String)recordMap.get('Id'), (String)consumptionSchedule.id);
                        if (recordMap.get('ConsumptionRates') != null ) {
                            List<Object> consumptionRatesList = (List<Object>)recordMap.get('ConsumptionRates');
                            System.assertEquals(consumptionRatesList[0], consumptionRate.Id);
                        }
                    } when 'ConsumptionRate' {
                        allObjectsToPullDataFrom.add('ConsumptionRates');
                        System.assertEquals((String)recordMap.get('Id'), (String)consumptionRate.id);
                    } when 'SBQQ__OrderItemConsumptionSchedule__c' {
                        allObjectsToPullDataFrom.add('SBQQ__OrderItemConsumptionSchedule__c');
                        System.assertEquals((String)recordMap.get('Id'), (String)orderItemConsumptionSchedule.id);
                        if (recordMap.get('OrderItemConsumptionRates') != null ) {
                            List<Object> orderItemConsumptionRatesList = (List<Object>)recordMap.get('OrderItemConsumptionRates');
                            System.assertEquals(orderItemConsumptionRatesList[0], orderItemConsumptionRate.Id);
                        }
                    } when 'SBQQ__OrderItemConsumptionRate__c' {
                        allObjectsToPullDataFrom.add('SBQQ__OrderItemConsumptionRate__c');
                        System.assertEquals((String)recordMap.get('Id'), (String)orderItemConsumptionRate.id);
                    }
                }
            } 
        }

        Set<String> allObjects = new Set<String>{
            'Contract',
            'Account',
            'Contact',
            'Product2',
            'PricebookEntry',
            'OrderItem',
            'initialOrder',
            'ProductConsumptionSchedule',
            'ConsumptionSchedule'
        };
        
        if(isCpqInstalled) {
            allObjects.add('SBQQ__OrderItemConsumptionSchedule__c');
            allObjects.add('SBQQ__OrderItemConsumptionRate__c');
            allObjects.add('Opportunity');
            allObjects.add('Order');
        }

        Set<String> listOfMissingObjectNames = new Set<String>();
        for (String objectName : allObjects) {
            if (!allObjectsToPullDataFrom.contains(objectName)) {
                listOfMissingObjectNames.add(objectName);
            }
        }

        if (!listOfMissingObjectNames.isEmpty()) {
            system.assertEquals(null, listOfMissingObjectNames);
        }
    }

    private static Map<String,Object> makeData(Boolean isAmendedOrderTest) {
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


        sObject amendedOpportunity = opportunityObject.get('Opportunity').newSObject();
        amendedOpportunity.put('Name', 'testOpportunity');
        amendedOpportunity.put('AccountId', account.id);
        amendedOpportunity.put('StageName','prospecting');
        amendedOpportunity.put('CloseDate', system.today() + 4);
        amendedOpportunity.put('Pricebook2Id', pricebookId);
        if(isCpqInstalled) {
            amendedOpportunity.put('SBQQ__Contracted__c', true);
        }
        insert amendedOpportunity;
        
        sObject quote;
        sObject amendedQuote;
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

            Map<String, Schema.SObjectType> amendedQuoteObject = Schema.getGlobalDescribe();
            amendedQuote = amendedQuoteObject.get('SBQQ__Quote__c').newSObject();
            amendedQuote.put('SBQQ__Account__c',account.id);
            amendedQuote.put('SBQQ__Status__c','Approved');
            amendedQuote.put('SBQQ__StartDate__c',date.today());
            amendedQuote.put('SBQQ__PricebookId__c',pricebookId);
            amendedQuote.put('SBQQ__Opportunity2__c',amendedOpportunity.id);
            amendedQuote.put('SBQQ__SubscriptionTerm__c',12);
            amendedQuote.put('SBQQ__ContractingMethod__c','By Subscription End Date');
            amendedQuote.put('SBQQ__Primary__c',true);
            amendedQuote.put('SBQQ__Ordered__c',true);
            insert amendedQuote;

            opportunity.put('SBQQ__PrimaryQuote__c', quote.Id);
            upsert opportunity;

            amendedOpportunity.put('SBQQ__PrimaryQuote__c', amendedQuote.Id);
            upsert amendedOpportunity;
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

        Map<String, Schema.SObjectType> consumptionScheduleObject = Schema.getGlobalDescribe();
        sObject consumptionSchedule = consumptionScheduleObject.get('ConsumptionSchedule').newSObject();
        consumptionSchedule.put('Name','testConsumptionSchedule');
        consumptionSchedule.put('BillingTerm',5);
        consumptionSchedule.put('isActive',false);
        consumptionSchedule.put('BillingTermUnit','Month');
        consumptionSchedule.put('Type','Range');
        consumptionSchedule.put('RatingMethod','Tier');
        insert consumptionSchedule;

        Map<String, Schema.SObjectType> consumptionRateObject = Schema.getGlobalDescribe();
        sObject consumptionRate = consumptionRateObject.get('ConsumptionRate').newSObject();
        consumptionRate.put('Price',50.00);
        consumptionRate.put('ConsumptionScheduleId',consumptionSchedule.Id);
        consumptionRate.put('ProcessingOrder',1);
        consumptionRate.put('PricingMethod','PerUnit');
        consumptionRate.put('LowerBound',1);
        insert consumptionRate;

        consumptionSchedule.put('isActive',true);
        upsert consumptionSchedule;

        Map<String, Schema.SObjectType> productConsumptionScheduleObject = Schema.getGlobalDescribe();
        sObject productConsumptionSchedule = productConsumptionScheduleObject.get('ProductConsumptionSchedule').newSObject();
        productConsumptionSchedule.put('ConsumptionScheduleId',consumptionSchedule.Id);
        productConsumptionSchedule.put('ProductId',product.Id);
        insert productConsumptionSchedule;

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
        if(isCpqInstalled)order.put('SBQQ__Quote__c',quote.id);
        insert order;

        //this stays
        Map<String, Schema.SObjectType> contractObject2 = Schema.getGlobalDescribe();
        sObject amendedContract = contractObject2.get('Contract').newSObject();
        amendedContract.put('AccountId',account.id);
        amendedContract.put('ContractTerm',12);
        amendedContract.put('Status','Draft');
        amendedContract.put('StartDate', system.today());
        if(isCpqInstalled) {
            
            amendedContract.put('SBQQ__Opportunity__c',opportunity.id);
            amendedContract.put('SBQQ__RenewalPricebookId__c',pricebookId);
            amendedContract.put('SBQQ__AmendmentRenewalBehavior__c','Latest End Date');
            amendedContract.put('SBQQ__AmendmentStartDate__c',system.today().AddMonths(1));
            amendedContract.put('SBQQ__AmendmentPricebookId__c',pricebookId);
            amendedContract.put('SBQQ__PreserveBundleStructureUponRenewals__c',true);
            amendedContract.put('SBQQ__Quote__c',amendedQuote.id);
        }
        insert amendedContract;
        if(isCpqInstalled && isAmendedOrderTest) {
            order.put('SBQQ__Quote__c',amendedQuote.id);
            upsert order;
            amendedOpportunity.put('SBQQ__AmendedContract__c',amendedContract.id);
            upsert amendedOpportunity;
        } 

        Map<String, Schema.SObjectType> amendedOrderObject = Schema.getGlobalDescribe();
        sObject amendedOrder = amendedOrderObject.get('Order').newSObject();
        amendedOrder.put('Name','testamendedOrder');
        amendedOrder.put('AccountId',account.id);
        amendedOrder.put('BillToContactId',contact.id);
        amendedOrder.put('ContractId',amendedContract.id);
        amendedOrder.put('OpportunityId',amendedOpportunity.id);
        amendedOrder.put('status','Draft');
        amendedOrder.put('EffectiveDate',system.today());
        amendedOrder.put('EndDate',system.today() + 40);
        amendedOrder.put('Pricebook2Id',pricebookId);
        if(isCpqInstalled)amendedOrder.put('SBQQ__Quote__c',amendedQuote.id);
        insert amendedOrder;

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

        sObject orderItemConsumptionSchedule; 
        sObject orderItemConsumptionRate;
        if(isCpqInstalled) {
            amendedContract.put('SBQQ__Order__c',amendedOrder.id);
            upsert amendedContract;

            contract.put('SBQQ__Order__c',order.id);
            upsert contract;
            if(!isAmendedOrderTest) {
                String relatedOpportunityId = amendedOpportunity.id;
        
                String amendedContractOpportunityQuery = 'SELECT Id, SBQQ__AmendedContract__c FROM Opportunity WHERE Id = :relatedOpportunityId LIMIT 1';
                SObject contractOpportunityRecord  = Database.query(amendedContractOpportunityQuery);
                contractOpportunityRecord.put('SBQQ__AmendedContract__c', amendedContract.id);
                upsert contractOpportunityRecord;
    
                String relatedAmendedContractId = amendedContract.id;
                String quoteContractQuery = 'SELECT Id, SBQQ__Quote__c FROM Contract WHERE Id = :relatedAmendedContractId LIMIT 1';
                SObject quoteContractRecord  = Database.query(quoteContractQuery);
                quoteContractRecord.put('SBQQ__Quote__c', quote.id);
                upsert quoteContractRecord;
            } 
            Map<String, Schema.SObjectType> orderItemConsumptionScheduleObject = Schema.getGlobalDescribe();
            orderItemConsumptionSchedule = orderItemConsumptionScheduleObject.get('SBQQ__OrderItemConsumptionSchedule__c').newSObject();
            orderItemConsumptionSchedule.put('Name','testOrderItemConsumptionSchedule');
            orderItemConsumptionSchedule.put('SBQQ__BillingTerm__c',5);
            orderItemConsumptionSchedule.put('SBQQ__BillingTermUnit__c','Monthly');
            orderItemConsumptionSchedule.put('SBQQ__ConsumptionSchedule__c',consumptionSchedule.Id);
            orderItemConsumptionSchedule.put('SBQQ__OrderItem__c',orderItem.id);
            orderItemConsumptionSchedule.put('SBQQ__Type__c','Range');
            orderItemConsumptionSchedule.put('SBQQ__RatingMethod__c','Tier');
            orderItemConsumptionSchedule.put('SBQQ__Category__c','Rates');
            insert orderItemConsumptionSchedule;

            Map<String, Schema.SObjectType> orderItemConsumptionRateObject = Schema.getGlobalDescribe();
            orderItemConsumptionRate = orderItemConsumptionRateObject.get('SBQQ__OrderItemConsumptionRate__c').newSObject();
            orderItemConsumptionRate.put('Name','testOrderItemConsumptionRate');
            orderItemConsumptionRate.put('SBQQ__Price__c',50.00);
            orderItemConsumptionRate.put('SBQQ__OrderItemConsumptionSchedule__c',orderItemConsumptionSchedule.Id);
            orderItemConsumptionRate.put('SBQQ__ProcessingOrder__c',1);
            orderItemConsumptionRate.put('SBQQ__PricingMethod__c','Per Unit');
            orderItemConsumptionRate.put('SBQQ__LowerBound__c',1);
            insert orderItemConsumptionRate;
        }

        Map<String, Schema.SObjectType> orderItem2Object = Schema.getGlobalDescribe();
        sObject orderItem2 = orderItem2Object.get('OrderItem').newSObject();
        orderItem2.put('OrderId',amendedOrder.id);
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

        List<String> contractIdsList = new List<String> {
            contact.id,
            amendedContract.id    
        };

        List<Contract> listOfContractIdRecords = [SELECT Id, Status FROM Contract WHERE Id IN :contractIdsList];
        if(!listOfContractIdRecords.isEmpty()) {
            for(Contract contractRecord : listOfContractIdRecords) {
                contractRecord.Status = 'Activated';
            }
            upsert listOfContractIdRecords;
        }

        List<String> orderIdsList = new List<String> {
            order.id,
            amendedOrder.id    
        };

        List<Order> listOfOrderIdRecords = [SELECT Id, Status FROM Order WHERE Id IN :orderIdsList];
        if(!listOfOrderIdRecords.isEmpty()) {
            for(Order orderRecord : listOfOrderIdRecords) {
                orderRecord.Status = 'Activated';
            }
            upsert listOfOrderIdRecords;
        }
        
        return new Map<String,Object> {
            'order' => order,
            'product' => product,
            'contract' => contract,
            'pricebookId' => pricebookId,
            'account' => account,
            'opportunity' => opportunity,
            'contact' => contact,
            'productConsumptionSchedule' => productConsumptionSchedule,
            'consumptionSchedule' => consumptionSchedule,
            'standardPrice' => standardPrice,
            'orderItem' => orderItem,
            'orderItemConsumptionSchedule' => orderItemConsumptionSchedule,
            'orderItem2' => orderItem2,
            'amendedOrder' => amendedOrder,
            'amendedContract' => amendedContract,
            'amendedOpportunity' => amendedOpportunity,
            'productConsumptionSchedule' => productConsumptionSchedule,
            'consumptionSchedule' => consumptionSchedule,
            'consumptionRate' => consumptionRate,
            'orderItemConsumptionSchedule' => orderItemConsumptionSchedule,
            'orderItemConsumptionRate' => orderItemConsumptionRate,
            'isCpqInstalled' => isCpqInstalled
        };
    } 
}