// Description: delete test records to look as if the account is empty
// usage: sfdx force:apex:execute -f bin/delete_test_record_data.apex -u mbianco+standardcpq@stripe.com
// note: need to delete this entitlement record manually in each org in order to delete all account records:
// https://help.salesforce.com/s/articleView?id=000358234&type=1 


public static void deleteData(String ObjectApiName) {
    try {
        System.debug('ObjectApiName');
        System.debug(ObjectApiName);
        String query = 'SELECT Id FROM ' + ObjectApiName +
                       ' WITH SECURITY_ENFORCED LIMIT 10000';

        List<SObject> listOfRecordsToDelete =  Database.query(query);
        System.debug('listOfRecordsToDelete.size');
        System.debug(listOfRecordsToDelete.size());
        
        if(!ListOfRecordsToDelete.isEmpty()) {
          delete listOfRecordsToDelete;
          deleteData(ObjectApiName);
        }
    } catch (QueryException queryExcept) {
        //CPQ object does not exit
        System.debug('queryExcept');
        System.debug(string.valueOf(queryExcept.getLineNumber()) + queryExcept.getMessage());

    } catch (Exception excep) {
        System.debug('excep');
        System.debug(string.valueOf(excep.getLineNumber()) + excep.getMessage());
    }
}
List<Order> activatedOrderList = [Select Id, Status From Order Where Status = 'Activated'];

if(!activatedOrderList.isEmpty()) {
    for(Order activatedOrder : activatedOrderList) {
    activatedOrder.status = 'Draft';
    }
    upsert activatedOrderList;
}

String syncRecordApiName = 'Sync_Record__c';

Integer stripeConnectorQaPackCount = [SELECT COUNT() FROM PackageLicense WHERE NamespacePrefix LIKE 'QaStripeConnect%'];
if(stripeConnectorQaPackCount > 0) {
  syncRecordApiName = 'QaStripeConnect__Sync_Record__c';
} 

Integer stripeConnectorProdPackCount = [SELECT COUNT() FROM PackageLicense WHERE NamespacePrefix LIKE 'stripeConnector%'];
if(stripeConnectorProdPackCount > 0) {
  syncRecordApiName = 'stripeConnector__Sync_Record__c';
} 

List<String> listOfObjectRecordsToDelete = new List<String> {
    'Order',
    'OrderItem',
    'SBQQ__QuoteLine__c',
    'SBQQ__Quote__c',
    'Opportunity',
    'Case',
    'Account',
    'Contact',
    'PricebookEntry',
    'Product2',
     syncRecordApiName
};

for (String objectApiName : listOfObjectRecordsToDelete) {
    deleteData(objectApiName);
}
