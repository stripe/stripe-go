/*you can hit this endpoint  per org by authenticating into the org and makeing a GET request to 
https://{org domain}.my.salesforce.com/services/apexrest/batchApi*/
@RestResource(urlMapping='/batchApi')
global with sharing class restService {
    //custom exception
    public class NotOrdersFoundException extends Exception {}

     //collects errors from service
    global static List<Object> responseJsonPayloadErrors = new List<Object>();

    //builds json paylod 
    global static List<Object> responseJsonPayloadRecords = new List<Object>();

    @HttpGet
    global static void restServiceEvents() {
        RestRequest restRequest = RestContext.request;
        RestResponse response = RestContext.response;
        try { 
            //get body of request and cast to map
            Map<String, Object> parsedReqBody = (Map<String, Object>)JSON.deserializeUntyped(restRequest.requestBody.toString());

            //check if order_ids were sent if not throw error
            List<Object> requestedListOfOrderIds;
            if(parsedReqBody.get('order_ids') != null) {
                requestedListOfOrderIds = (List<Object>)parsedReqBody.get('order_ids');                  
            } else {
                throw new NotOrdersFoundException('There were no order Ids sent with this request.');
            }

            //initialize lists and check if product ids, order fields or product fields were sent 
            List<Object> requestListOfProduct2Ids = new List<Object>();
            List<Object> requestListOfProductFields = new List<Object>();
            List<Object> requestListOfOrderFields = new List<Object>();
            if(parsedReqBody.get('product2_ids') != null) requestListOfProduct2Ids = (List<Object>)parsedReqBody.get('product2_ids'); 
            if(parsedReqBody.get('product2_fields') != null) requestListOfProductFields = (List<Object>)parsedReqBody.get('product2_fields'); 
            if(parsedReqBody.get('order_fields') != null) requestListOfOrderFields = (List<Object>)parsedReqBody.get('order_fields'); 

            //converting object lists to sets to remove dupes if any exist and allow binding in dynamic soql query
            Set<String> listOfOrderIds = new Set<String>();
            Set<String> listOfProduct2Ids = new Set<String>();
            for(Object requestedOrderId : requestedListOfOrderIds) {
                listOfOrderIds.add(String.valueOf(requestedOrderId));
            }
            for(Object requestedProduct2Id : requestListOfProduct2Ids) {
                listOfProduct2Ids.add(String.valueOf(requestedProduct2Id));
            }
            Set<String> listOfOrderFields = new Set<String>();
            Set<String> listOfProductFields = new Set<String>();
            for(Object requestedOrderField : requestListOfOrderFields) {
                listOfOrderFields.add(String.valueOf(requestedOrderField));
            }
            if(parsedReqBody.get('product2_fields') != null)requestlistOfProductFields = (List<Object>)parsedReqBody.get('product2_fields');
            for(Object requestedProductField : requestListOfProductFields) {
                listOfProductFields.add(String.valueOf(requestedProductField));
            }

            //if any of the field lists are empty get all fields 
            if(listOfOrderFields.isEmpty())listOfOrderFields = getAccessableFieldsByObject('Order');
            if(listOfProductFields.isEmpty())listOfProductFields = getAccessableFieldsByObject('Product2');

            //get all the order records based on list order ids requested
            String orderQuery = 'SELECT '+ String.join(((Iterable<String>)listOfOrderFields), ',') + ' FROM Order WHERE Id in :listOfOrderIds  WITH SECURITY_ENFORCED';
            List<Order> listOfOrderRecords = Database.query(orderQuery);

            /*if order records are found in query get  related object records if look up fields have values*/
            if(!listOfOrderRecords.isEmpty()) {
                for(SObject orderRecord : listOfOrderRecords) {
                    //converting order to a map so we can add the related orderItems Id before adding the record to payload
                    Map<String, Object> orderRecordMap = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(orderRecord));
                    
                    //get related order records based on look up 
                    Set<String> listOfOrderItemFields = getAccessableFieldsByObject('OrderItem'); 
                    orderRecordMap = addRelatedOrderRecordsToPayload(listOfOrderFields,listOfOrderItemFields, orderRecordMap);
         
                    //get related order item records and their related records By Id
                    orderRecordMap = getRelatedOrderItemRecordsById(orderRecordMap,listOfOrderItemFields); 
                     
                    //adding modified order record to payload 
                    if(orderRecordMap != null || !orderRecordMap.isEmpty())responseJsonPayloadRecords.add((Object)orderRecordMap);
                }                          
            } else {
                throw new NotOrdersFoundException('There were no orders found based on the list of order Ids sent. Here are the order Ids that were sent: ' + listOfOrderIds);
            }
            /*create response payload with a records object containing a list of objects of all the records
            fetched and an error object containing and errors of in-accessable fields */
            response.addHeader('Content-Type', 'application/json');
            response.statusCode = 200;
            Map<String, Object> resBody = new Map<String, Object>();
            resBody.put('errors', responseJsonPayloadErrors);
            Map<String, List<Object>> orderDetails = new Map<String,List<Object>>();
            resBody.put('records',(Object)responseJsonPayloadRecords);
            response.responseBody = Blob.valueOf(JSON.serialize(resBody));
        } catch(DmlException e) {
            //Catches any CRUD SOQL failure and adds to error response map and continues
            Map<String,Object> DmlExceptionMap = new Map<String,Object>{
                'error_type' => 'DmlException',
                'error_message' => e.getMessage() + 'Line Number: '+string.valueOf(e.getLineNumber())
            };
            responseJsonPayloadErrors.add((Object)DmlExceptionMap);

        } catch(SObjectException e) {
             //Catches any SObject access failure and addes to error response map and continues
            Map<String,Object> SObjectExceptionMap = new Map<String,Object>{
                'error_type' => 'SObjectException',
                'error_message' => e.getMessage() + 'Line Number: '+string.valueOf(e.getLineNumber())
            };
            responseJsonPayloadErrors.add((Object)SObjectExceptionMap); 
        } catch(Exception e) {
            /*Catches any other exception and ends operation. returning the error message 
            and line number as the message param of response*/
            response.statusCode = 500;
            Map<String, Object> resBody = new Map<String, Object>();
            resBody.put('error_type', 'APEX_ERROR');
            resBody.put('error_message', e.getLineNumber() + ': ' + e.getMessage());
            response.responseBody = Blob.valueOf(JSON.serialize(resBody));
            return;
        }
    }

    //fetches all related object records that we care about based on the order record passed in 
    private static Map<String,Object> addRelatedOrderRecordsToPayload(Set<String> listOfOrderFields, Set<String> listOfOrderItemFields, Map<String,Object> orderRecordMap) {

        if(listOfOrderFields.contains('ContractId')) {
            //gets related contract record based the contract this order looks up to 
            orderRecordMap = addRelatedRecordToPayload(orderRecordMap,'Contract','ContractId');
            String orderId = (String)orderRecordMap.get('Id');

            //gets related orders for the contract this order looks up to 
            String contractId = string.valueOf(orderRecordMap.get('ContractId'));
            if(contractId != null) {
                String contractOrderQuery = 'SELECT '+ String.join(((Iterable<String>)listOfOrderFields), ',') + ' FROM Order WHERE ContractId = :contractId AND Id != :orderId WITH SECURITY_ENFORCED';
                List<Order> listOfContractOrderRecords = Database.query(contractOrderQuery);
                Map<String, Object> contractOrderRecordMap = new Map<String,Object>();
                for(SObject contractOrderRecord : listOfContractOrderRecords) {
                    contractOrderRecordMap = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(contractOrderRecord));
                    contractOrderRecordMap = getRelatedOrderItemRecordsById(contractOrderRecordMap,listOfOrderItemFields); 
                    
                    /*recursively call this function to get all related data for each order related to this contract 
                    while explicitly exluding the order which called this funciton */
                    contractOrderRecordMap = addRelatedOrderRecordsToPayload(listOfOrderFields, listOfOrderItemFields, contractOrderRecordMap);
                }
                if(contractOrderRecordMap != null || !contractOrderRecordMap.isEmpty())responseJsonPayloadRecords.add((Object)contractOrderRecordMap); 
            }
        }
        //gets related account record based the account this order looks up to 
        if(listOfOrderFields.contains('AccountId')) {
            orderRecordMap = addRelatedRecordToPayload(orderRecordMap,'Account','AccountId');
        }
        //gets related contact record based the contact this order looks up to 
        if(listOfOrderFields.contains('BillToContactId')) {
            orderRecordMap = addRelatedRecordToPayload(orderRecordMap,'Contact','BillToContactId');
        }
        //gets related opportunity record based the opportunity this order looks up to 
        if(listOfOrderFields.contains('OpportunityId')) {
            orderRecordMap = addRelatedRecordToPayload(orderRecordMap,'Opportunity','OpportunityId');
        }
        //gets related pricebook record based the pricebook this order looks up to 
        if(listOfOrderFields.contains('Pricebook2Id')) {
            orderRecordMap = addRelatedRecordToPayload(orderRecordMap,'Pricebook2','Pricebook2Id');
        }

        //gets related quote record and specific quote child records based on the quote this order looks up to 
        if(listOfOrderFields.contains('SBQQ__Quote__c')) {
            String quoteId = string.valueOf(orderRecordMap.get('SBQQ__Quote__c'));
            if(quoteId != null) {
                //get related quote record to order record
                //orderRecordMap = addRelatedRecordToPayload(orderRecordMap,'SBQQ__Quote__c','SBQQ__Quote__c');
               

                //get related records based on the lookup to quote from the order
                Set<String> listOfQuoteFields = getAccessableFieldsByObject('SBQQ__Quote__c');
                String quoteQuery = 'SELECT '+ String.join(((Iterable<String>)listOfQuoteFields), ',')+' FROM SBQQ__Quote__c WHERE Id = :quoteId WITH SECURITY_ENFORCED LIMIT 1 ';
                List<SBQQ__Quote__c> listOfQuoteRecords = Database.query(quoteQuery);

                if(!listOfQuoteRecords.isEmpty()) {
                    SObject quoteRecord = listOfQuoteRecords[0];
                    Map<String, Object> quoteRecordMap = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(quoteRecord));

                    //adds list of the related quote ID to orderRecord 
                    List<Object> relatedQuoteRecordIdList = new List<Object>{quoteRecordMap.get('Id')};
                    orderRecordMap.put('Quotes',(Object)relatedQuoteRecordIdList);

                    //gets related primary contact based on the contact this quote looks up to 
                    if(listOfQuoteFields.contains('SBQQ__PrimaryContact__c')) {
                        quoteRecordMap = addRelatedRecordToPayload(quoteRecordMap,'Contact','SBQQ__PrimaryContact__c');
                    }

                    //gets related opportunity based on the opportunity this quote looks up to 
                    if(listOfQuoteFields.contains('SBQQ__Opportunity2__c')) {
                        quoteRecordMap = addRelatedRecordToPayload(quoteRecordMap,'Opportunity','SBQQ__Opportunity2__c');
                    }
                    /* Do we want the Accounts related to the Order Or Quote? -Questiion for Mike ANd Spencer
                    if(listOfQuoteFields.contains('SBQQ__Account__c')) {
                        quoteRecordMap = addRelatedRecordToPayload(quoteRecordMap,'Account','SBQQ__Account__c');
                    } */
                    
                    //add modified quote record with lookup ID lists to records param in payload
                    if(quoteRecordMap != null || !quoteRecordMap.isEmpty())responseJsonPayloadRecords.add((Object)quoteRecordMap);
                }
            }
        }
        return orderRecordMap;
    }

    //gets all related order item records and related child records that we care about
    private static Map<String,Object> getRelatedOrderItemRecordsById(Map<String,Object> orderRecordMap, Set<String> listOfOrderItemFields) {
        String orderRecordId = (String)orderRecordMap.get('Id');

        //get related order Items based on the Id of the order record passed in 
        String orderItemQuery = 'SELECT '+ String.join(((Iterable<String>)listOfOrderItemFields), ',') + ' FROM OrderItem WHERE OrderId = :orderRecordId  WITH SECURITY_ENFORCED LIMIT 1';
        List<OrderItem> listOfOrderItemData = Database.query(orderItemQuery);

        //create list to get all Order Item Ids that are related to the Order record passed in
        List<String> listOfOrderItemIds = new List<String>();
        for(SObject orderItem : listOfOrderItemData) {

            //add each related order item id to a list to put in record at the end
            listOfOrderItemIds.add(string.valueOf(orderItem.get('Id'))); 

            //convert order item record to map so we can add related record Ids to it if they exist
            Map<String, Object> orderItemMap = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(orderItem));

            //gets related product this order item looks up to 
            if(listOfOrderItemFields.contains('Product2Id')) {
                orderItemMap = addRelatedRecordToPayload(orderItemMap,'Product2','Product2Id');
            }
            //gets related price book entry this order item looks up to 
            if(listOfOrderItemFields.contains('PricebookEntryId')) {
                orderItemMap = addRelatedRecordToPayload(orderItemMap,'PricebookEntry','PricebookEntryId');
            }
            //gets related contract this order item looks up to 
            if(listOfOrderItemFields.contains('SBQQ__Contract__c')) {
                orderItemMap = addRelatedRecordToPayload(orderItemMap,'Contract','SBQQ__Contract__c');
            }
            //gets related contract line item this order item looks up to 
            if(listOfOrderItemFields.contains('SBQQSC__ContractLineItem__c')) {
                orderItemMap = addRelatedRecordToPayload(orderItemMap,'ContractLineItem','SBQQSC__ContractLineItem__c');
            }
            //gets related quote line this order item looks up to 
            if(listOfOrderItemFields.contains('SBQQ__QuoteLine__c')) {
                orderItemMap = addRelatedRecordToPayload(orderItemMap,'SBQQ__QuoteLine__c','SBQQ__QuoteLine__c');
            }
            //gets related subscribtion this order item looks up to 
            if(listOfOrderItemFields.contains('SBQQ__Subscription__c')) {
                orderItemMap = addRelatedRecordToPayload(orderItemMap,'SBQQ__Subscription__c','SBQQ__Subscription__c');
            }  
            //add order item record to record payload param with ids of related objects included       
            if(orderItemMap != null || !orderItemMap.isEmpty())responseJsonPayloadRecords.add((Object)orderItemMap);
        }
        orderRecordMap.put('OrderItems',(Object)listOfOrderItemIds);
        return orderRecordMap;
    }
    

    /* this function gets the related records based on the record, object API name and lookupfield passed in.
    We then add a list of Ids with the corresponding plural object name to the record passed in and return to 
    calling function for insert */
    private static Map<String, Object> addRelatedRecordToPayload(Map<String, Object> recordMap, String objectApiName, String LookupField) {
        
        //list to related record Ids to 
        List<Object> listOfRelatedRecordsIds = new List<Object>();

        //get the Id of the record that the recordMap looks up to 
        String objectId = string.valueOf(recordMap.get(LookupField));
        if(objectId != null) {
            Set<String> listOfObjectFields = getAccessableFieldsByObject(objectApiName);

            //query list of related records based on lookup field and object api name 
            String query = 'SELECT '+ String.join(((Iterable<String>)listOfObjectFields), ',') + ' FROM ' + objectApiName +' WHERE Id = :objectId  WITH SECURITY_ENFORCED';
            List<SObject> listOfRelatedRecordData = Database.query(query);
            
            /*add each related record to the records json response and add the Id to a list 
            to add back to parent record map after*/
            for(SObject relatedRecord : listOfRelatedRecordData) {                
                if(relatedRecord != null)responseJsonPayloadRecords.add((Object)relatedRecord);
                listOfRelatedRecordsIds.add(string.valueOf(relatedRecord.get('Id'))); 
            }
            if(!listOfRelatedRecordsIds.isEmpty()){
                /* before adding list of related record Ids to the parent 
                record switch the objectApiName to the  plural object name */
                switch on objectApiName {
                    when 'Pricebook2' {
                        objectApiName = 'PriceBooks';
                    }
                    when 'Account' {
                        objectApiName = 'Accounts';
                    }
                    when 'Opportunity' {
                        objectApiName = 'Opportunities';
                    }
                    when 'Product2' {
                        objectApiName = 'Products';
                    }
                    when 'PricebookEntry' {
                        objectApiName = 'PricebookEntries';
                    }
                    when 'Subscription' {
                        objectApiName = 'Subscriptions';
                    }
                    when 'Contract' {
                        objectApiName = 'Contracts';
                    }
                    when 'Contact' {
                        objectApiName = 'Contacts';
                    }
                    when 'OrderItem' {
                        objectApiName = 'OrderItems';
                    }
                }
                //add list of child ids to parent record 
                recordMap.put(objectApiName,(Object)listOfRelatedRecordsIds);
            }
        }
        return recordMap;
    }

    //get a set off all accesable fields to an object based on the objects api name passed in 
    private static Set<String> getAccessableFieldsByObject(String ObjectApiName) {
        Set<String> accessableFields = new Set<String>();
         //create error if we dont have access to this object 
         if (Schema.describeSObjects(new List<String>{ objectApiName })[0].isAccessible()) {
            Schema.DescribeSObjectResult describedObj = Schema.describeSObjects(new List<String>{ObjectApiName}, SObjectDescribeOptions.DEFERRED).get(0);
            for(Schema.SObjectField fieldSchema : describedObj.fields.getMap().values()) {
                Schema.DescribeFieldResult describedField = fieldSchema.getDescribe();

                //create an error in the rest errors list if a field is not accessable
                if(!describedField.isAccessible()) {
                    Map<String,Object> errorResponseMap = new Map<String,Object>{
                        'error_type' => 'FieldNotAccessable',
                        'error_message' => 'Field Not Accessable',
                        'sobject_type' => ObjectApiName,
                        'sobject_field' => describedField.getName()
                    };
                    responseJsonPayloadErrors.add((Object)errorResponseMap);
                    continue;
                }
                //add all accessable fields to a list 
            accessableFields.add(describedField.getName());
        }
    } else {
        Map<String,Object> errorResponseMap = new Map<String,Object>{
            'error_type' => 'ObjectNotAccessable',
            'error_message' => 'Object Not Accessable',
            'sobject_type' => ObjectApiName,
            'sobject_field' => null
        };
        responseJsonPayloadErrors.add((Object)errorResponseMap);
    }
       return accessableFields;
    }
}
