@isTest
public with sharing class test_batchDeleteSyncRecords {
    @isTest
    
    static public void test_batchDeleteSyncRecords() {
        Stripe_Connection__c stripeConnectRec = new Stripe_Connection__c(); 
        stripeConnectRec.Stripe_Connected__c = true;
        stripeConnectRec.Salesforce_Connected__c = true;
        stripeConnectRec.API_Key__c = 'testPackageGUID';
        stripeConnectRec.Sync_Record_Retention__c = '0';
        upsert stripeConnectRec;

        Sync_Record__c testSyncRecord1 = new Sync_Record__c (
            Primary_Record_ID__c = '1234',
            Primary_Object_Type__c = 'Order',
            Secondary_Record_ID__c = '4321',
            Secondary_Object_Type__c = 'OrderItem',
            Resolution_Status__c = 'Resolved',
            Compound_ID__c = '12344321'
        );
        insert testSyncRecord1;

        Sync_Record__c testSyncRecord2 = new Sync_Record__c (
            Primary_Record_ID__c = '12345',
            Primary_Object_Type__c = 'Order',
            Secondary_Record_ID__c = '3333',
            Secondary_Object_Type__c = 'OrderItem',
            Resolution_Status__c = 'Resolved',
            Compound_ID__c = '123453333'
        );
        insert testSyncRecord2;

        Sync_Record__c testSyncRecord3 = new Sync_Record__c (
            Primary_Record_ID__c = '123456',
            Primary_Object_Type__c = 'Order',
            Secondary_Record_ID__c = '6666',
            Secondary_Object_Type__c = 'OrderItem',
            Resolution_Status__c = 'Resolved',
            Compound_ID__c = '1234566666'
        );
        insert testSyncRecord3;

        Test.startTest();   
        batchDeleteSyncRecords deleteSyncRecords = new batchDeleteSyncRecords();
        ID batchprocessid = Database.executeBatch(deleteSyncRecords);  
        Test.stopTest();

        System.assertNotEquals(batchprocessid, Null);

        List<Sync_Record__c> syncRecordList = [SELECT Id, Resolution_Status__c, CreatedDate
                                                FROM Sync_Record__c
                                                WHERE Resolution_Status__c = 'Resolved'
                                                WITH SECURITY_ENFORCED
                                                ORDER BY CreatedDate];  
        System.assert(syncRecordList.isEmpty());

    }
}
