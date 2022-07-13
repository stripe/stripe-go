// Description: delete test records to look as if the account is empty
// Usage: sfdx force:apex:execute -f bin/delete_test_record_data.apex -u mbianco+standardcpq@stripe.com
// NOTE: do NOT include a log level definition when executing, it seems to mess with which account the apex is run on
// note: need to delete this entitlement record manually in each org in order to delete all account records:

// https://help.salesforce.com/s/articleView?id=000358234&type=1
// https://appiphony15-dev-ed.lightning.force.com/lightning/setup/CompanyResourceDisk/home

public static void deleteData(String ObjectApiName) {
    Integer batchSize = 500;

    try {
        System.debug('ObjectApiName: ' + ObjectApiName);
        String query = 'SELECT Id FROM ' + ObjectApiName +
                       ' LIMIT ' + batchSize;

        List<SObject> listOfRecordsToDelete = Database.query(query);
        System.debug('number of records: ' + listOfRecordsToDelete.size());

        if(!ListOfRecordsToDelete.isEmpty()) {
            System.debug('deleting records');
            delete listOfRecordsToDelete;

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

List<Order> activatedOrderList = [
    SELECT Id, Status, SBQQ__Contracted__c
    FROM Order
    WHERE Status = 'Activated'
    LIMIT 10000
];

System.debug('Updating order (' + activatedOrderList.size() + ') state to enable deletion...');

if(!activatedOrderList.isEmpty()) {
    try {
        // before 'deactivating' an order, we have to disable contracting
        for(Order activatedOrder : activatedOrderList) {
            activatedOrder.SBQQ__Contracted__c = false;
        }
        upsert activatedOrderList;

        // now, with the order non-contracted, we can deactivate it
        for(Order activatedOrder : activatedOrderList) {
            activatedOrder.status = 'Draft';
        }
        upsert activatedOrderList;
    } catch (Exception excep) {
        System.debug('order cleanup exception:');
        System.debug(string.valueOf(excep.getLineNumber()) + excep.getMessage());
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

if(isCPQInstalled()) {
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
            upsert recordJobsToDelete;
        }
        delete recordJobsToDelete;
    }
}

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
