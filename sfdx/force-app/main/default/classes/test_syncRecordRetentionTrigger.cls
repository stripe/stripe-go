@isTest
public with sharing class test_syncRecordRetentionTrigger {
    @isTest
    static public void test_syncRecordRetentionTrigger() {
        Stripe_Connection__c stripeConnectRec = new Stripe_Connection__c(); 
        stripeConnectRec.Stripe_Connected__c = true;
        stripeConnectRec.Salesforce_Connected__c = true;
        stripeConnectRec.API_Key__c = 'testPackageGUID';
        stripeConnectRec.Sync_Record_Retention__c = '100';
        upsert stripeConnectRec;

        List<Sync_Record__c> syncRecordsToInsert = new List<Sync_Record__c>();
        for(Integer i = 0; i < 500; i++) {
            Sync_Record__c syncRec = new Sync_Record__c(
                Primary_Record_ID__c = 'Order' + i,
                Resolution_Status__c = 'Resolved'
                );
            syncRecordsToInsert.add(syncRec);
        }
        
        Test.startTest();   
        insert syncRecordsToInsert;
        Test.stopTest();

        List<Sync_Record__c> syncRecordList = [SELECT Id, Resolution_Status__c, CreatedDate
                                                FROM Sync_Record__c
                                                WHERE Resolution_Status__c = 'Resolved'
                                                WITH SECURITY_ENFORCED
                                                ORDER BY CreatedDate];  
        System.assertEquals(syncRecordList.size(), 100);

    }
}
