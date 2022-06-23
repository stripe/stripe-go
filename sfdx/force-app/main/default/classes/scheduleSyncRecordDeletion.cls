global with sharing class scheduleSyncRecordDeletion implements Schedulable {
    //schedules recurring batch job to run at 10pm daily
    global static String scheduleJob() {
        String scheduledTime = '0 0 22 * * ? *';
        scheduleSyncRecordDeletion scheduledSyncDeleteJob = new scheduleSyncRecordDeletion();
        return System.schedule('Stripe Connector Sync Record Deletion Batch Job', scheduledTime, scheduledSyncDeleteJob);
    }

    global void execute(SchedulableContext sc) {
        try {
            batchDeleteSyncRecords deleteSyncRecordsJob = new batchDeleteSyncRecords();
            ID batchprocessid = Database.executeBatch(deleteSyncRecordsJob, 500);
        } catch (Exception e) {
            errorLogger.create('Scheduling Batch Delete Sync Records Job', e);
        }
    }

    // https://developer.salesforce.com/forums/?id=906F0000000BSU0IAO
    @TestVisible private static void abortStripeCronJobs () {
        if(!Test.isRunningTest()) {
            return;
        }

        List<CronTrigger> scheduledBatchJobList = [
            SELECT id, CronJobDetail.Name
            FROM CronTrigger
            WHERE CronJobDetail.Name LIKE 'Stripe%'
        ];

        if(!scheduledBatchJobList.isEmpty()) {
            for (CronTrigger scheduledBatchJob : scheduledBatchJobList){
                System.abortJob(scheduledBatchJob.id);
            }
        }
    }
}
