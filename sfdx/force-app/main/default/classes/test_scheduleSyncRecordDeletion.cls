@isTest
public with sharing class test_scheduleSyncRecordDeletion {
    @isTest

    static public void test_scheduleSyncRecordDeletion() {
        scheduleSyncRecordDeletion.abortStripeCronJobs();

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
        scheduleSyncRecordDeletion.scheduleJob();
        Test.stopTest();

        List<CronTrigger> scheduledJobsList = [SELECT Id, NextFireTime, CronJobDetail.Name FROM CronTrigger];
        System.assertNotEquals(scheduledJobsList[0], Null);
    }
}
