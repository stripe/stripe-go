// Description: delete test records to look as if the account is empty
// Usage: sfdx force:apex:execute -f bin/delete_test_record_data.apex -u mbianco+standardcpq@stripe.com
// NOTE: do NOT include a log level definition when executing, it seems to mess with which account the apex is run on
// note: need to delete this entitlement record manually in each org in order to delete all account records:

// https://help.salesforce.com/s/articleView?id=000358234&type=1
// https://appiphony15-dev-ed.lightning.force.com/lightning/setup/CompanyResourceDisk/home

public static void deleteData(String ObjectApiName) {
    // max that can be used with deleteImmediately
    Integer batchSize = 500;

    if(Limits.getLimitCpuTime() - Limits.getCpuTime() < 4000) {
        System.debug('Close to limit, quitting!');
        return;
    }

    try {
        System.debug('ObjectApiName: ' + ObjectApiName);

        // magic constants can't be used in a `from` clause
        String query = 'SELECT Id FROM ' + ObjectApiName +
                        ' WITH SECURITY_ENFORCED' +
                       ' LIMIT ' + batchSize;

        List<SObject> listOfRecordsToDelete = Database.query(query);
        System.debug('number of records: ' + listOfRecordsToDelete.size());

        if(!ListOfRecordsToDelete.isEmpty()) {
            System.debug('deleting records');

            // `false` second params ensure that if there are some errors, we don't fail completely
            // there will always be some errors due to weird data issues in SF
            Database.delete(listOfRecordsToDelete, false);

            if(ListOfRecordsToDelete.size() == batchSize) {
                deleteData(ObjectApiName);
            }
        } else {
            System.debug('deleted all records');
        }
    } catch (QueryException queryExcept) {
        // CPQ object does not exist
        // TODO why could this ever happen?
        System.debug('query exception:');
        System.debug(string.valueOf(queryExcept.getLineNumber()) + queryExcept.getMessage());
    } catch (Exception excep) {
        System.debug('general exception:');
        System.debug(string.valueOf(excep.getLineNumber()) + excep.getMessage());
    }
}

public static void deactivateOrders() {
    List<Order> activatedOrderList = [
        SELECT Id, Status, SBQQ__Contracted__c
        FROM Order
        WHERE Status = 'Activated'
        LIMIT 2000
    ];

    System.debug('Updating order (' + activatedOrderList.size() + ') state to enable deletion...');

    // NOTE updates are run sync to enable this script to run faster when it is run mu;tiple times in a row

    if(!activatedOrderList.isEmpty()) {
        try {
            // before 'deactivating' an order, we have to disable contracting
            for(Order activatedOrder : activatedOrderList) {
                activatedOrder.SBQQ__Contracted__c = false;
            }
            Database.updateImmediate((List<SObject>)activatedOrderList);

            // now, with the order non-contracted, we can deactivate it
            for(Order activatedOrder : activatedOrderList) {
                activatedOrder.status = 'Draft';
            }
            Database.updateImmediate((List<SObject>)activatedOrderList);
        } catch (Exception excep) {
            System.debug('order cleanup exception:');
            System.debug(string.valueOf(excep.getLineNumber()) + excep.getMessage());
        }
    }
}

public static string determinePackagePrefix() {
    Integer stripeConnectorQaPackCount = [
        SELECT COUNT()
        FROM PackageLicense
        WHERE NamespacePrefix LIKE 'QaStripeConnect%'
    ];

    if(stripeConnectorQaPackCount > 0) {
        return 'QaStripeConnect__';
    }

    Integer stripeConnectorProdPackCount = [
        SELECT COUNT()
        FROM PackageLicense
        WHERE NamespacePrefix
        LIKE 'stripeConnector%'
    ];

    if(stripeConnectorProdPackCount > 0) {
        return 'stripeConnector__';
    }

    return '';
}

public static boolean isCPQInstalled() {
    Integer sbqqPackageCount = [
        SELECT COUNT()
        FROM PackageLicense
        WHERE NamespacePrefix
        LIKE 'SBQQ%'
    ];

    return sbqqPackageCount > 0;
}

public static void deleteCPQRecordJobs() {
    if(!isCPQInstalled()) {
        System.debug('CPQ not installed, no record jobs to delete');
        return;
    }

    Boolean isRecordJobsUpdated = false;
    List<SBQQ__RecordJob__c> recordJobsToDelete = [
        SELECT SBQQ__JobStatus__c, Id
        FROM SBQQ__RecordJob__c
        LIMIT 10000
    ];

    System.debug('listOfRecordJobs.size: ' + recordJobsToDelete.size());

    if (!recordJobsToDelete.isEmpty()) {
        for (SBQQ__RecordJob__c recordJob : recordJobsToDelete) {
            if(recordJob.SBQQ__JobStatus__c != 'Completed') {
                recordJob.SBQQ__JobStatus__c = 'Completed';
                isRecordJobsUpdated = true;
            }
        }

        if (isRecordJobsUpdated) {
            Database.updateImmediate((List<SObject>)recordJobsToDelete);
        }

        Database.deleteImmediate((List<SObject>)recordJobsToDelete);
    }
}

deleteCPQRecordJobs();
deactivateOrders();

// determine the packaged namespace in this particular account
String syncRecordApiName = determinePackagePrefix() + 'Sync_Record__c';

List<String> listOfObjectRecordsToDelete = new List<String> {
    'Contract',
    'SBQQ__OrderItemConsumptionSchedule__c',
    'SBQQ__OrderItemConsumptionRate__c',
    'Order',
    'OrderItem',
    'SBQQ__QuoteLine__c',
    'SBQQ__Quote__c',
    'Opportunity',
    'Case',
    'Account',
    'Contact',
    'ConsumptionSchedule',
    'ConsumptionRate',
     syncRecordApiName,
    'PricebookEntry',
    'Product2'
};

for (String objectApiName : listOfObjectRecordsToDelete) {
    deleteData(objectApiName);
}
