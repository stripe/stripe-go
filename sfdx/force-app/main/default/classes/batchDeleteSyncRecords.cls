// runs via a cron/scheduled job to deleted records beyond the user-configurable retention
global with sharing class batchDeleteSyncRecords implements Database.Batchable<SObject> {
    public batchDeleteSyncRecords() {}

    public List<SObject> start(Database.BatchableContext BC) {
        checkAccess();

        List<Sync_Record__c> allSuccessAndResolvedSyncRecords = [
            SELECT Id, Resolution_Status__c, CreatedDate
            FROM Sync_Record__c
            WHERE (
                Resolution_Status__c = 'Resolved'
                OR Resolution_Status__c = 'Success'
            )
            WITH SECURITY_ENFORCED
            ORDER BY CreatedDate ASC
        ];

        Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();

        if(stripeConnectRec.Id != null && stripeConnectRec.Sync_Record_Retention__c != null &&
        (allSuccessAndResolvedSyncRecords.size() > integer.valueof(stripeConnectRec.Sync_Record_Retention__c))) {
            return allSuccessAndResolvedSyncRecords;
        }
        return new List<SObject>();
    }

    public void execute(Database.BatchableContext c, List<SObject> scope) {
        try {
            delete scope;
        } catch(Exception e) {
            errorLogger.create('Batch Delete Sync Records', e);
        } 
    }
    public void finish(Database.BatchableContext BC) {}

    public static void checkAccess() {
        dmlManager.checkAccess(new Map<String, Map<String, Set<String>>> {
            'Sync_Record__c' => new Map<String, Set<String>> {
                'Delete' => new Set<String> {}
            }
        });
    }
}
