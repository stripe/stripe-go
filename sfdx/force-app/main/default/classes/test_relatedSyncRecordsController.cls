@isTest
public with sharing class test_relatedSyncRecordsController {
    @isTest
    static public void test_getRelatedSyncRecords() {
        Stripe_Connection__c stripeConnectRec = new Stripe_Connection__c(); 
        stripeConnectRec.Stripe_Connected__c = true;
        stripeConnectRec.Salesforce_Connected__c = true;
        stripeConnectRec.API_Key__c = 'testPackageGUID';
        upsert stripeConnectRec;

        Sync_Record__c testSyncRecord = new Sync_Record__c (
            Primary_Record_ID__c = '1234',
            Primary_Object_Type__c = 'Order',
            Secondary_Record_ID__c = '4321',
            Secondary_Object_Type__c = 'OrderItem',
            Resolution_Status__c = 'Error',
            Compound_ID__c = '12344321'
        );
        insert testSyncRecord;

        testSyncRecord.Primary_Record_ID__c = testSyncRecord.Id;
        testSyncRecord.Secondary_Record_ID__c = testSyncRecord.Id;
        upsert testSyncRecord;

        Test.startTest();
        String response = relatedSyncRecordsController.getRelatedSyncRecords(testSyncRecord.Id);
        Test.stopTest();

        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        
        List<Object> syncRecordList = (List<Object>)resultsMap.get('syncRecordList');
        Map<String, Object> syncRecordListMap = (Map<String, Object>)syncRecordList[0];

        System.assertEquals((String)syncRecordListMap.get('Id'), testSyncRecord.Id);
        System.assertEquals((String)syncRecordListMap.get('Primary_Record_ID__c'), testSyncRecord.Id);
        System.assertEquals((String)syncRecordListMap.get('Secondary_Record_ID__c'), testSyncRecord.Id);
        System.assertEquals((String)syncRecordListMap.get('Primary_Object_Type__c'), testSyncRecord.Primary_Object_Type__c);
        System.assertEquals((String)syncRecordListMap.get('Secondary_Object_Type__c'), testSyncRecord.Secondary_Object_Type__c);
        System.assertEquals((String)syncRecordListMap.get('Resolution_Status__c'), testSyncRecord.Resolution_Status__c);
    }
}
