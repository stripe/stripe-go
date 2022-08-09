// Documentation on Triggers: https://trailhead.salesforce.com/content/learn/modules/apex_triggers/apex_triggers_intro
// https://docs.google.com/document/d/1_7kf3oUTuNqQuopfapTP30Di1XwpuY7VYS_APYiGuPY/edit#heading=h.xv8ify7m2448

trigger updateSyncRecordsStatuses on Sync_Record__c (after insert, after update) {

    try {
        // List to store triggered primary object ids
        Set<String> newlyResolvedSyncRecordPrimaryObjectIdList = new  Set<String>();

        // On new sync records for primary records, on Success or Resolved let's update all the children.
        if (Trigger.isInsert) {
            for (Sync_Record__c syncRecord: Trigger.new) { 
                if (
                    (String)syncRecord.Resolution_Status__c == 'Resolved' || (String)syncRecord.Resolution_Status__c == 'Success'
                    && syncRecord.Primary_Record_ID__c == syncRecord.Secondary_Record_ID__c
                ) {
                    newlyResolvedSyncRecordPrimaryObjectIdList.add(syncRecord.Primary_Record_ID__c);
                }
            }
        }

        // On updates to sync records for primary records, if it is moving into Success or Resolved, let's update all the children.
        if (Trigger.isUpdate) {
            for (Sync_Record__c syncRecord: Trigger.new) { 
                Sync_Record__c oldSyncRecord = trigger.oldMap.get(syncRecord.Id);
                Boolean isPrimarySyncRecord = syncRecord.Primary_Record_ID__c == syncRecord.Secondary_Record_ID__c;

                if (
                    isPrimarySyncRecord &&
                    (String)syncRecord.Resolution_Status__c == 'Resolved' || (String)syncRecord.Resolution_Status__c == 'Success' 
                    && (
                        oldSyncRecord.Resolution_Status__c == 'Error' ||
                        oldSyncRecord.Resolution_Status__c == 'Skipped'
                    )
                ) {
                    newlyResolvedSyncRecordPrimaryObjectIdList.add(syncRecord.Primary_Record_ID__c);
                }
            }
        }

        // Get all children records for the triggered primary records.
        List<Sync_Record__c> allUnresolvedSyncRecords = [
            SELECT Id, Primary_Record_ID__c, Resolution_Status__c
            FROM Sync_Record__c
            WHERE (
                Resolution_Status__c != 'Resolved'
                AND Resolution_Status__c != 'Success'
            )
            AND Primary_Record_ID__c IN :newlyResolvedSyncRecordPrimaryObjectIdList
            // don't update records which are causing this trigger to fire
            AND Id NOT IN :trigger.newMap.keySet()
            WITH SECURITY_ENFORCED
        ]; 

        // set all children sync records to resolved and persist.
        for (Sync_Record__c recordToMarkResolved : allUnresolvedSyncRecords) {
            recordToMarkResolved.Resolution_Status__c = 'Resolved';
        }

        Database.updateImmediate((List<SObject>)allUnresolvedSyncRecords);
        //update recordToMarkResolved;

    } catch (Exception e) {
        errorLogger.create('updateSyncRecordsStatusesTrigger', e);
    }
}