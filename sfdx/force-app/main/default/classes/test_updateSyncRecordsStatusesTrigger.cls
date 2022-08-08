@isTest
public with sharing class test_updateSyncRecordsStatusesTrigger {
    @isTest
    static public void testUpdateSyncRecordsStatuses() {
        // TODO why is this being constructed manually here? Shouldn't there be a helper for this?
        Stripe_Connection__c stripeConnectRec = new Stripe_Connection__c(); 
        stripeConnectRec.Stripe_Connected__c = true;
        stripeConnectRec.Salesforce_Connected__c = true;
        stripeConnectRec.API_Key__c = 'testPackageGUID';
        stripeConnectRec.Sync_Record_Retention__c = '501';
        upsert stripeConnectRec;


        Sync_Record__c testSyncRecord1 = new Sync_Record__c (
            Primary_Record_ID__c = '1234',
            Primary_Object_Type__c = 'Order',
            Secondary_Record_ID__c = '4321',
            Secondary_Object_Type__c = 'OrderItem',
            Resolution_Status__c = 'Error',
            Compound_ID__c = '12344321'
        );
        insert testSyncRecord1;

        List<Sync_Record__c> syncRecordsToInsert = new List<Sync_Record__c>();
        for(Integer i = 0; i < 500; i++) {
            Sync_Record__c syncRec = new Sync_Record__c(
                Primary_Record_ID__c = testSyncRecord1.Primary_Record_ID__c ,
                Resolution_Status__c = 'Error'
                );
            syncRecordsToInsert.add(syncRec);
        }
        insert syncRecordsToInsert;

        Test.startTest();
        testSyncRecord1.Resolution_Status__c = 'Resolved';
        upsert testSyncRecord1;
        Test.stopTest();

        List<Sync_Record__c> updatedSyncRecords = [SELECT Id, Primary_Record_ID__c, Secondary_Record_ID__c, Resolution_Status__c, Compound_ID__c 
                                                   FROM Sync_Record__c
                                                   WHERE Primary_Record_ID__c = :testSyncRecord1.Primary_Record_ID__c
                                                   AND Resolution_Status__c = 'Resolved'];

        System.assertEquals(updatedSyncRecords.size(), 501);  
    }
}
