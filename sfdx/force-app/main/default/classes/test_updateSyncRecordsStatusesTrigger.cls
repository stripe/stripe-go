@isTest
public with sharing class test_updateSyncRecordsStatusesTrigger {
    // TODO should add to general test helpers
    // https://salesforce.stackexchange.com/questions/21137/creating-unit-tests-without-interacting-with-the-database-creating-fake-ids
    static Integer s_num = 1;
    public static String getFakeId(Schema.SObjectType sot) {
        String result = String.valueOf(s_num++);
        return sot.getDescribe().getKeyPrefix() + '0'.repeat(12-result.length()) + result;
    }
    
    @isTest
    static public void testUpdateSyncRecordsStatuses() {
        setupSyncRecordRetention();

        Sync_Record__c parentSyncRecord = generatePrimarySyncRecord('Error');
        insert parentSyncRecord;

        // adding secondary linked sync records which should be auto-resolved when the parent is updated
        List<Sync_Record__c> syncRecordsToInsert = new List<Sync_Record__c>();
        for(Integer i = 0; i < 50; i++) {
            String secondaryRecordId = getFakeId(Schema.OrderItem.SObjectType);          
            Sync_Record__c syncRec = new Sync_Record__c(
                Primary_Record_ID__c = parentSyncRecord.Primary_Record_ID__c ,
                Secondary_Record_ID__c = secondaryRecordId,
                Resolution_Status__c = 'Error',
                Compound_ID__c = parentSyncRecord.Primary_Record_ID__c + '-' + secondaryRecordId
            );
            syncRecordsToInsert.add(syncRec);
        }
        insert syncRecordsToInsert;

        Test.startTest();
        // Now we translate the top level Order as a success, the child objects above should be marked as Resolved.
        parentSyncRecord.Resolution_Status__c = 'Success';
        update parentSyncRecord;
        Test.stopTest();

        List<Sync_Record__c> updatedSyncRecords = getUpdatedSyncRecords(parentSyncRecord.Primary_Record_ID__c);

        System.assertEquals(50, updatedSyncRecords.size());  
    }

    // this occurs when a parent (order) record is added which resolves children (order lines)
    @isTest
    static public void testInsertSyncRecordsStatuses() {
        setupSyncRecordRetention();

        Sync_Record__c parentSyncRecord = generatePrimarySyncRecord('Success');

        // adding secondary linked sync records which should be auto-resolved when the parent is updated
        List<Sync_Record__c> syncRecordsToInsert = new List<Sync_Record__c>();
        for(Integer i = 0; i < 50; i++) {
            String secondaryRecordId = getFakeId(Schema.OrderItem.SObjectType);          
            Sync_Record__c syncRec = new Sync_Record__c(
                Primary_Record_ID__c = parentSyncRecord.Primary_Record_ID__c ,
                Secondary_Record_ID__c = secondaryRecordId,
                Resolution_Status__c = 'Error',
                Compound_ID__c = parentSyncRecord.Primary_Record_ID__c + '-' + secondaryRecordId
            );
            syncRecordsToInsert.add(syncRec);
        }
        insert syncRecordsToInsert;

        // although the trigger runs async, the `Test` helpers ensure all async operations/futures complete before assertions
        Test.startTest();
        // Now we translate the top level Order as a success, the child objects above should be marked as Resolved.
        insert parentSyncRecord;
        Test.stopTest();

        List<Sync_Record__c> updatedSyncRecords = getUpdatedSyncRecords(parentSyncRecord.Primary_Record_ID__c);

        System.assertEquals(50, updatedSyncRecords.size());  
    }
    
    static public  List<Sync_Record__c> getUpdatedSyncRecords(String parentId) {
        return [
            SELECT Id
            FROM Sync_Record__c
            WHERE Primary_Record_ID__c = :parentId
            AND Resolution_Status__c = 'Resolved'
        ];
    }

    static public Sync_Record__c generatePrimarySyncRecord(String status) {
        String primaryObjectId = getFakeId(Schema.Order.SObjectType);
        return new Sync_Record__c (
            Primary_Record_ID__c = primaryObjectId,
            Primary_Object_Type__c = 'Order',
            Secondary_Record_ID__c = primaryObjectId,
            Secondary_Object_Type__c = 'Order',
            Resolution_Status__c = status,
            Compound_ID__c = primaryObjectId + '-' + primaryObjectId
        );        
    }

    // TODO should make sure a consistent method is used anytime we generate connector state for testing
    static public void setupSyncRecordRetention() {
        Stripe_Connection__c stripeConnectRec = new Stripe_Connection__c(); 
        stripeConnectRec.Stripe_Connected__c = true;
        stripeConnectRec.Salesforce_Connected__c = true;
        stripeConnectRec.API_Key__c = 'testPackageGUID';
        // this connector state is just to define the retention
        stripeConnectRec.Sync_Record_Retention__c = '51';
        upsert stripeConnectRec;
    }
}
