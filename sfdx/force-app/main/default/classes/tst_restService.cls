@isTest
public with sharing class tst_restService {
     
    @isTest 
    public static void test_restServiceEvents() {
        //insert all the records we will be checking the JSON response for
        Id pricebookId = Test.getStandardPricebookId();
        Account account = new Account (
            Name='SFDC Account',
            SBQQ__RenewalModel__c = 'Contract Based'
        );
        insert account;
        Opportunity opportunity = new Opportunity(
            Name = 'testOpportunity',
            AccountId = account.id,
            StageName = 'prospecting',
            CloseDate = system.today() + 4,
            Pricebook2Id = pricebookId
        );
        insert opportunity;
        
        SBQQ__Quote__c quote = new SBQQ__Quote__c(
            SBQQ__Account__c = account.id,
            SBQQ__Status__c = 'Approved',
            SBQQ__StartDate__c = date.parse('06/07/2025'),
            SBQQ__EndDate__c = date.parse('06/10/2026'),
            SBQQ__PricebookId__c = pricebookId,
            SBQQ__Opportunity2__c = opportunity.id,
            SBQQ__SubscriptionTerm__c = 2
        );
        insert quote;
       
        Product2 product = new Product2(
            Name = ' testProduct',
            Description='testProductEntry ',
            productCode = 'ABC',
            isActive = true
        );
        insert product;
        PricebookEntry standardPrice = new PricebookEntry(
            Pricebook2Id = pricebookId,
            Product2Id = product.Id,
            UnitPrice = 1,
            IsActive = true,
            UseStandardPrice = false
        );
        insert standardPrice;
        SBQQ__QuoteLine__c quoteLine = new SBQQ__QuoteLine__c(
            SBQQ__Quote__c = quote.id,
            SBQQ__Product__c = product.id,
            SBQQ__PricebookEntryId__c = standardPrice.id,
            SBQQ__Quantity__c = 10
        );
        insert quoteLine;
          
        OpportunityLineItem opportunityLine = new OpportunityLineItem(
            Product2Id = product.id,
            PricebookEntryId = standardPrice.id,
            OpportunityId = opportunity.id,
            SBQQ__QuoteLine__c = quoteLine.id,
            TotalPrice = 100.00,
            Quantity = 10
        );
        insert opportunityLine; 
        Contract contract = new Contract(
            Name = 'testContract',
            SBQQ__Quote__c = quote.id,
            AccountId = account.id,
            SBQQ__Opportunity__c = opportunity.id,
            ContractTerm = 2,
            StartDate = system.today(),
            SBQQ__RenewalPricebookId__c = 'testRenew',
            SBQQ__AmendmentPricebookId__c = 'testAmmend'
        );
        insert contract;
        opportunity.SBQQ__PrimaryQuote__c = quote.id;
        opportunity.SBQQ__RenewedContract__c = contract.id;
        opportunity.ContractId = contract.id;
        update opportunity;
        Contact contact = new Contact(
            FirstName = 'testContact',
            LastName = 'testContact',
            Email = 'testContact@testContact.com',
            AccountId = account.id
        );
        insert contact;
        SBQQ__Subscription__c subscription = new SBQQ__Subscription__c(
            SBQQ__Account__c = account.id,
            SBQQ__Contract__c = contract.id,
            SBQQ__Product__c = product.id,
            SBQQ__Quantity__c = 10
        );
        insert subscription;
        Order order = new Order(
            Name = 'testOrder',
            AccountId = account.id,
            BillToContactId = contact.id,
            ContractId = contract.id,
            OpportunityId = opportunity.id,
            //SBQQ__Quote__c = quote.id, causing validation error need to dig on this Arnold 
            status = 'Draft',
            EffectiveDate = system.today(),
            EndDate = system.today() + 40,
            Pricebook2Id =  pricebookId 
        );
        insert order;
        Order order2 = new Order(
            Name = 'testOrder2',
            AccountId = account.id,
            BillToContactId = contact.id,
            ContractId = contract.id,
            OpportunityId = opportunity.id,
            //SBQQ__Quote__c = quote.id, causing validation error need to dig on this Arnold 
            status = 'Draft',
            EffectiveDate = system.today(),
            EndDate = system.today() + 40,
            Pricebook2Id =  pricebookId 
        );
        insert order2;
        OrderItem orderItem = new OrderItem(
            OrderId = order.id,
            Quantity = 24,
            UnitPrice = 240,
            Product2id = product.id,
            PricebookEntryId = standardPrice.id,
            ServiceDate = system.today() + 4,
            EndDate = system.today() + 14
        );
        insert orderItem;
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
        Test.startTest(); 
        restService.restServiceEvents();
        Test.stopTest();
        Map<String, Object> parsedReqBody = (Map<String, Object>)JSON.deserializeUntyped(response.responseBody.toString());
        List<Object> records = (List<Object>)parsedReqBody.get('records');
        
        for(Object record : records) {
            Map<String, Object> recordMap = (Map<String, Object>)record;
            if(recordMap.get('attributes') != null) {
                
                Map<String, Object> attributes = (Map<String, Object>)recordMap.get('attributes');
                String objectType = (String)attributes.get('type');
                
                switch on objectType {
                    when 'Contract' {
                        System.assertEquals((String)recordMap.get('Id'), (String)contract.id);
                        System.assertEquals((String)recordMap.get('SBQQ__RenewalPricebookId__c'), 'testRenew');
                        System.assertEquals((String)recordMap.get('SBQQ__AmendmentPricebookId__c'), 'testAmmend');
                        System.assertEquals((String)recordMap.get('SBQQ__Quote__c'), (String)quote.id);
                        System.assertEquals(recordMap.get('ContractTerm'), 2);
                        System.assertEquals((String)recordMap.get('StartDate'), string.valueOf(system.today()));
                        System.assertEquals((String)recordMap.get('AccountId'), (String)account.id);
                        System.assertEquals((String)recordMap.get('SBQQ__Opportunity__c'), (String)opportunity.id);
                    }
                    when 'Account' {
                        System.assertEquals((String)recordMap.get('Id'), (String)account.id);
                        System.assertEquals((String)recordMap.get('Name'), 'SFDC Account');
                        System.assertEquals((String)recordMap.get('SBQQ__RenewalModel__c'), 'Contract Based');
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
                        System.assertEquals((String)recordMap.get('Id'), (String)orderItem.id);
                        System.assertEquals((String)recordMap.get('OrderId'), (String)order.id);
                        System.assertEquals((String)recordMap.get('PricebookEntryId'), (String)standardPrice.id);
                        System.assertEquals(recordMap.get('Quantity'), 24);
                        System.assertEquals(recordMap.get('UnitPrice'), 240);
                        System.assertEquals((String)recordMap.get('ServiceDate'), string.valueOf(system.today() + 4));
                        System.assertEquals((String)recordMap.get('EndDate'), string.valueOf(system.today() + 14));
                        List<Object> Products = (List<Object>)recordMap.get('Products');
                        List<Object> PricebookEntries = (List<Object>)recordMap.get('PricebookEntries');
                        System.assertEquals((String)Products[0], (String)product.id);
                        System.assertEquals((String)PricebookEntries[0], (String)standardPrice.id);
                    }
                    when 'Order' {
                        if((String)recordMap.get('Id') == order.id){
                            System.assertEquals((String)recordMap.get('Id'), (String)order.id);
                            System.assertEquals((String)recordMap.get('AccountId'), (String)account.id);
                            System.assertEquals((String)recordMap.get('BillToContactId'), (String)contact.id);
                            System.assertEquals((String)recordMap.get('ContractId'), (String)contract.id);
                            System.assertEquals((String)recordMap.get('Pricebook2Id'), (String)pricebookId);
                            System.assertEquals((String)recordMap.get('EffectiveDate'), string.valueOf(system.today()));
                            System.assertEquals((String)recordMap.get('EndDate'), string.valueOf(system.today() + 40));
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
                            System.assertEquals((String)recordMap.get('EffectiveDate'), string.valueOf(system.today()));
                            System.assertEquals((String)recordMap.get('EndDate'), string.valueOf(system.today() + 40));
                            System.assertEquals((String)recordMap.get('Name'), 'testOrder2');
                            List<Object> Contracts = (List<Object>)recordMap.get('Contracts');
                            List<Object> Accounts = (List<Object>)recordMap.get('Accounts');
                            List<Object> Contacts = (List<Object>)recordMap.get('Contacts');
                            List<Object> OrderItems = (List<Object>)recordMap.get('OrderItems');
                            System.assertEquals((String)Contracts[0], (String)contract.id);
                            System.assertEquals((String)Accounts[0], (String)account.id);
                            System.assertEquals((String)Contacts[0], (String)contact.id);
                            System.assertEquals((String)OrderItems[0], (String)orderItem.id);
                        }
 
                       /*   */
                        
                    }
                    when 'Opportunity' {
                        System.assertEquals((String)recordMap.get('Id'), (String)opportunity.id);
                        System.assertEquals((String)recordMap.get('AccountId'), (String)account.id);
                        System.assertEquals((String)recordMap.get('Pricebook2Id'), (String)pricebookId);
                        System.assertEquals((String)recordMap.get('StageName'), 'prospecting');
                        System.assertEquals((String)recordMap.get('Name'), 'testOpportunity');
                        System.assertEquals((String)recordMap.get('CloseDate'), string.valueOf(system.today() + 4));
                    }
                }
            } 
        }
    }
}