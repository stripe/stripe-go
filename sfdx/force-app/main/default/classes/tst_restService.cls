@isTest
public with sharing class tst_restService {
     
    @isTest
    public static void test_restServiceEvents() {
        //insert all the records we will be checking the JSON response for
        Account account = new Account (
            Name='SFDC Account'
        );
        insert account;

        SBQQ__Quote__c quote = new SBQQ__Quote__c(
            SBQQ__Account__c = account.id,
            SBQQ__Status__c = 'Approved',
            SBQQ__StartDate__c = date.parse('06/07/2025'),
            SBQQ__EndDate__c = date.parse('06/10/2026')
        );
        insert quote;
        
        Opportunity opportunity = new Opportunity(
            Name = 'testOpportunity',
            StageName = 'prospecting',
            CloseDate = date.parse('06/07/2020'),
            SBQQ__PrimaryQuote__c = quote.id
        );
        insert opportunity;

        Contract contract = new Contract(
            Name = 'testContract',
            SBQQ__Quote__c = quote.id,
            AccountId = account.id
        );
        insert contract;

        Contact contact = new Contact(
            FirstName = 'testContact',
            LastName = 'testContact',
            Email = 'testContact@testContact.com',
            AccountId = account.id
        );
        insert contact;


        Order order = new Order(
            Name = 'testOrder',
            AccountId = account.id,
            BillToContactId = contact.id,
            ContractId = contract.id,
            OpportunityId = opportunity.id,
            SBQQ__Quote__c = quote.id,
            status = 'Activated'
        );
        insert order;

        OrderItem orderItem = new OrderItem(
            OrderId = order.id,
            SBQQ__Contract__c = contact.id
        );
        insert orderItem;

        //make mock request
        //list<String> requestListOfOrderIds = new List<String>{(String)order.id};
        RestRequest request = new RestRequest();
        request.requestBody = Blob.valueOf(JSON.serialize(requestListOfOrderIds));
        request.httpMethod = 'Get';
        RestContext.request = request;
        RestContext.response = new RestResponse();
        RestResponse response = RestContext.response;

        Test.startTest(); 
        //HttpCalloutMock results = new mockCallout();
        //Test.setMock(HttpCalloutMock.class, results);
        //restService.restServiceEvents();
        Test.stopTest();

        //System.debug(response);
        //System.debug(response.getBody());

    }

}