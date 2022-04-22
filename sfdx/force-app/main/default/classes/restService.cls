/*you can hit this endpoint  per org by authenticating into the org and makeing a GET request to 
https://{orgDomain}.my.salesforce.com/services/apexrest/batchApi*/
@RestResource(urlMapping='/batchApi')
global with sharing class restService {
    //custom exception
    public class NotOrdersFoundException extends Exception {}

     //collects errors from service
    global static Set<Object> responseJsonPayloadErrors = new Set<Object>();

    //builds json paylod 
    global static Set<Object> responseJsonPayloadRecords = new Set<Object>();

    //readonly annotation is used to increase the SOQL limit from 50k to a 100k 
    @ReadOnly
    
    // TODO only keeping the GET issue to avoid metadata issues in the package, we should remove
    @HttpGet

    // using POST here to make it easier to send a JSON message body (the SF ruby gem does not support GET with body out of the box)
    @HttpPost

    global static void restServiceEvents() {
        
        RestRequest restRequest = RestContext.request;
        RestResponse response = RestContext.response;
        try { 
            Boolean isCpqInstalled = utilities.isCpqEnabled();

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

            for(Object requestedProductField : requestListOfProductFields) {
                listOfProductFields.add(String.valueOf(requestedProductField));
            }

            //if any of the field lists are empty get all fields 
            if(listOfOrderFields.isEmpty())listOfOrderFields = getAccessableFieldsByObject('Order');
            if(listOfProductFields.isEmpty())listOfProductFields = getAccessableFieldsByObject('Product2');

            //get all the order records based on list order ids requested
            String orderQuery = 'SELECT '+ String.join(((Iterable<String>)listOfOrderFields), ',') + ' FROM Order WHERE Id IN :listOfOrderIds  WITH SECURITY_ENFORCED';
            List<Order> listOfOrderRecords = Database.query(orderQuery);

            //Used to hold all order records after converting to map
            Set<Map<String, Object>> listOfOrderRecordMapRecords = new Set<Map<String, Object>>();

            //if order records are found in add to order map list
            if(!listOfOrderRecords.isEmpty()) {
                for(SObject orderRecord : listOfOrderRecords) {
                    //converting order to a map so we can add the related orderItems Id before adding the record to payload
                    Map<String, Object> orderRecordMap = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(orderRecord));
                    
                    //adding order records map to list
                    listOfOrderRecordMapRecords.add(orderRecordMap);
                }         

                //get related order records based on specific lookups existing in the targeted object field list(s)  
                Set<String> listOfOrderItemFields = getAccessableFieldsByObject('OrderItem'); 
                listOfOrderRecordMapRecords = addRelatedOrderRecordsToPayload(listOfOrderFields,listOfOrderItemFields, listOfOrderRecordMapRecords, listOfOrderIds, isCpqInstalled);

                //adding modified order record to payload 
                for(Map<String, Object> modifiedOrderRec : listOfOrderRecordMapRecords) {
                    if(modifiedOrderRec != null && !modifiedOrderRec.isEmpty()) responseJsonPayloadRecords.add((Object)modifiedOrderRec);
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
    private static Set<Map<String,Object>> addRelatedOrderRecordsToPayload(Set<String> listOfOrderFields, Set<String> listOfOrderItemFields, Set<Map<String, Object>> listOfOrderRecordMapRecords, Set<String> listOfOrderIds, Boolean isCpqInstalled) {
        if(listOfOrderFields.contains('ContractId')) {
           //pulls out all contract Ids from order records into seperate list
           Set<String> listOfContractRecordIds = new Set<String>();
           for(Map<String,Object> orderRecordMap : listOfOrderRecordMapRecords) {
               if(orderRecordMap.get('ContractId') != null)listOfContractRecordIds.add((String)orderRecordMap.get('ContractId'));
           }

            /*gets related orders for the contract that are not equal to our current list of order
            ids but lookup to any of the same contracts within our current order list*/
            if(!listOfContractRecordIds.isEmpty()) {
                String contractOrderQuery = 'SELECT '+ String.join(((Iterable<String>)listOfOrderFields), ',') + ' FROM Order WHERE ContractId IN :listOfContractRecordIds AND Id NOT IN :listOfOrderIds WITH SECURITY_ENFORCED';
                List<Order> listOfContractOrderRecords = Database.query(contractOrderQuery);

                /*convert any orders that lookup to the same contracts as the orders in listOfOrderRecordMapRecords. 
                Add the records to a new list to recurse on and get the Ids to add to our order record*/
                Set<Map<String, Object>> listOfContractOrderRecordMaps = new Set<Map<String,Object>>();
                Map<String, Object> contractOrderRecordMap = new Map<String, Object>();
                Set<String> listOfContractOrderRecordIds = new Set<String>();//delete
                for(Map<String,Object> orderRecordMap : listOfOrderRecordMapRecords) {
                    for(SObject contractOrderRecord : listOfContractOrderRecords) {
                        if(orderRecordMap.get('ContractId') == contractOrderRecord.get('ContractId')) {
                            contractOrderRecordMap = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(contractOrderRecord));
                            //adding all other orders that look up to the same contract to our list of order records to get related data below
                            listOfOrderRecordMapRecords.add(contractOrderRecordMap);

                            //creating seperate list so we can add back to each order record and the relation to the contract is shown
                            listOfContractOrderRecordIds.add((String)contractOrderRecordMap.get('Id'));

                            //adding new ids contract order ids to existing list so they will be excluded when we recurse
                            listOfOrderIds.add((String)contractOrderRecordMap.get('Id'));
                        }
                    }
                    orderRecordMap.put('ContractOrders',(Object)listOfContractOrderRecordIds);
                }
                 //gets related contracts record based the contracts these orders look up to 
                listOfOrderRecordMapRecords = addRelatedRecordToPayload(listOfOrderRecordMapRecords,'Contract','ContractId');
            }
        }

        //get all related order items based on the orders contained in list
        listOfOrderRecordMapRecords = getRelatedOrderItemRecordsById(listOfOrderRecordMapRecords,listOfOrderItemFields, listOfOrderIds, isCpqInstalled);

        //gets related account records based on the accounts any order in our list looks up to 
        if(listOfOrderFields.contains('AccountId')) {
            listOfOrderRecordMapRecords = addRelatedRecordToPayload(listOfOrderRecordMapRecords,'Account','AccountId');
        }
        //gets related contact records based on the contacts any order in our list looks up to 
        if(listOfOrderFields.contains('BillToContactId')) {
            listOfOrderRecordMapRecords = addRelatedRecordToPayload(listOfOrderRecordMapRecords,'Contact','BillToContactId');
        }
        //gets related opportunity records based on the opportunities any order in our list looks up to 
        if(listOfOrderFields.contains('OpportunityId')) {
            listOfOrderRecordMapRecords = addRelatedRecordToPayload(listOfOrderRecordMapRecords,'Opportunity','OpportunityId');
        }
        //gets related pricebook records based on the pricebooks sny order in our list looks up to 
        if(listOfOrderFields.contains('Pricebook2Id')) {
            listOfOrderRecordMapRecords = addRelatedRecordToPayload(listOfOrderRecordMapRecords,'Pricebook2','Pricebook2Id');
        }

        //check if CPQ is installed before getting data from objects in the package 
        if(isCpqInstalled) {
            //get list of all quote Ids that our orders look up to
            Set<String> listOfOrderRecordQuoteId = new Set<String>();
            String cpqQuoteObjectApiName = 'SBQQ__Quote__c';
            if(listOfOrderFields.contains(cpqQuoteObjectApiName)) {
                for(Map<String, Object> orderRecordMap : listOfOrderRecordMapRecords) {
                    if(orderRecordMap.get(cpqQuoteObjectApiName) != null) {
                        String quoteId = (String)orderRecordMap.get(cpqQuoteObjectApiName);
                        listOfOrderRecordQuoteId.add(quoteId);
                    }
                }
            }

            //get related records based on the lookup to quote from the order
            List<SObject> listOfQuoteRecords = new List<SObject>();
            Set<String> listOfQuoteFields = getAccessableFieldsByObject(cpqQuoteObjectApiName);
            if(!listOfOrderRecordQuoteId.isEmpty()) {
                String quoteQuery = 'SELECT '+ String.join(((Iterable<String>)listOfQuoteFields), ',')+' FROM ' + cpqQuoteObjectApiName + ' WHERE Id IN :listOfOrderRecordQuoteId WITH SECURITY_ENFORCED';
                listOfQuoteRecords = Database.query(quoteQuery);
            }

            //gets related quote record and specific quote child records based on the quote this order looks up to 
            Set<Map<String, Object>> listOfQuoteRecordMaps = new Set<Map<String, Object>>();
            if(!listOfQuoteRecords.isEmpty() ) {
                for(Map<String, Object> orderRecordMap : listOfOrderRecordMapRecords) {
                    List<Object> relatedQuoteRecordIdList = new List<Object>();
                    for(SObject quoteRecord : listOfQuoteRecords) {
                        if(orderRecordMap.get(cpqQuoteObjectApiName) == quoteRecord.get('Id')) {
                            Map<String, Object> quoteRecordMap = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(quoteRecord));
                            listOfQuoteRecordMaps.add(quoteRecordMap); 
                            relatedQuoteRecordIdList.add(quoteRecordMap.get('Id'));
                        }
                    }
                    //adds list of the related quote ID to orderRecords
                    orderRecordMap.put('Quotes',(Object)relatedQuoteRecordIdList);
                }
            }
            
            if(!listOfQuoteRecordMaps.isEmpty()) {
                //gets related primary contact records based on the contacts the list quote records looks up to 
                if(listOfQuoteFields.contains('SBQQ__PrimaryContact__c')) {
                    listOfQuoteRecordMaps = addRelatedRecordToPayload(listOfQuoteRecordMaps,'Contact','SBQQ__PrimaryContact__c');
                }

                //gets related opportunity records based on the opportunities the list quote records looks up to 
                if(listOfQuoteFields.contains('SBQQ__Opportunity2__c')) {
                    listOfQuoteRecordMaps = addRelatedRecordToPayload(listOfQuoteRecordMaps,'Opportunity','SBQQ__Opportunity2__c');
                }

                //gets related account records based on the accounts the list quote records looks up to 
                if(listOfQuoteFields.contains('SBQQ__Account__c')) {
                    listOfQuoteRecordMaps = addRelatedRecordToPayload(listOfQuoteRecordMaps,'Account','SBQQ__Account__c');
                } 
                
                for(Map<String,Object> modifiedQuoteRecordMap : listOfQuoteRecordMaps) {
                    //add modified quote record with lookup ID lists to records param in payload
                    if(modifiedQuoteRecordMap != null && !modifiedQuoteRecordMap.isEmpty())responseJsonPayloadRecords.add((Object)modifiedQuoteRecordMap);
                }
            }
        }
        return listOfOrderRecordMapRecords;
    }

    //gets all related order item records and related child records that we care about
    private static Set<Map<String,Object>> getRelatedOrderItemRecordsById(Set<Map<String, Object>> listOfOrderRecordMapRecords, Set<String> listOfOrderItemFields, Set<String> listOfOrderIds, Boolean isCpqInstalled) {
        //get related order Items based on the Ids in the order record list
        String orderItemQuery = 'SELECT '+ String.join(((Iterable<String>)listOfOrderItemFields), ',') + ' FROM OrderItem WHERE OrderId IN :listOfOrderIds WITH SECURITY_ENFORCED';
        List<OrderItem> listOfOrderItemData = Database.query(orderItemQuery);

        /*used to build list of order items and add a list of lookup Ids to each order item record based on 
        each related record found*/
        Set<Map<String, Object>> listOfOrderItemRecordMapRecords = new Set<Map<String, Object>>();

        for(Map<String, Object> orderRecord : listOfOrderRecordMapRecords) {
            String orderRecordId = (String)orderRecord.get('Id');  
            //create list to get all Order Item Ids that are related to the Order record passed in
            Set<String> listOfOrderItemIds = new Set<String>();
            for(SObject orderItem : listOfOrderItemData) {
                if(orderRecordId == orderItem.get('OrderId')) {
                    //convert order item record to map so we can add related record Ids to it if they exist
                    Map<String, Object> orderItemMap = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(orderItem));

                    //add each related order item id to a list to put in order record at the end
                    listOfOrderItemIds.add(string.valueOf(orderItemMap.get('Id'))); 

                    //add each related order item to list
                    listOfOrderItemRecordMapRecords.add(orderItemMap);
                }
            }
            orderRecord.put('OrderItems',(Object)listOfOrderItemIds);
        }

        //gets related product this order item looks up to 
        if(listOfOrderItemFields.contains('Product2Id')) {
            listOfOrderItemRecordMapRecords = addRelatedRecordToPayload(listOfOrderItemRecordMapRecords,'Product2','Product2Id');
        }
        //gets related price book entry this order item looks up to 
        if(listOfOrderItemFields.contains('PricebookEntryId')) {
            listOfOrderItemRecordMapRecords = addRelatedRecordToPayload(listOfOrderItemRecordMapRecords,'PricebookEntry','PricebookEntryId');
        }
        //check if CPQ is installed before getting data from objects in the package 
        if(isCpqInstalled) {
            
            //gets related contract this order item looks up to 
            if(listOfOrderItemFields.contains('SBQQ__Contract__c')) {
                listOfOrderItemRecordMapRecords = addRelatedRecordToPayload(listOfOrderItemRecordMapRecords,'Contract','SBQQ__Contract__c');
            }
            //gets related contract line item this order item looks up to 
            if(listOfOrderItemFields.contains('SBQQSC__ContractLineItem__c')) {
                listOfOrderItemRecordMapRecords = addRelatedRecordToPayload(listOfOrderItemRecordMapRecords,'ContractLineItem','SBQQSC__ContractLineItem__c');
            }
            //gets related quote line this order item looks up to 
            if(listOfOrderItemFields.contains('SBQQ__QuoteLine__c')) {
                listOfOrderItemRecordMapRecords = addRelatedRecordToPayload(listOfOrderItemRecordMapRecords,'SBQQ__QuoteLine__c','SBQQ__QuoteLine__c');
            }
            //gets related subscribtion this order item looks up to 
            if(listOfOrderItemFields.contains('SBQQ__Subscription__c')) {
                listOfOrderItemRecordMapRecords = addRelatedRecordToPayload(listOfOrderItemRecordMapRecords,'SBQQ__Subscription__c','SBQQ__Subscription__c');
            }
        }

        //adds each modified order record to payload
        for(Map<String, Object> modifiedOrderItem : listOfOrderItemRecordMapRecords) {
            if(modifiedOrderItem != null && !modifiedOrderItem.isEmpty()) responseJsonPayloadRecords.add((Object)modifiedOrderItem);
        }

        return listOfOrderRecordMapRecords;
    }

    /* this function gets the related records based on the list of records, object API name and lookupfield passed in.
    We then add a list of Ids with the corresponding plural object name to each record passed in and return to 
    calling function for insertion into the payload */
    private static Set<Map<String, Object>> addRelatedRecordToPayload(Set<Map<String, Object>> listOfOrderRecordMapRecords, String objectApiName, String LookupField) {
         //gets list of all fields we have access to for an object
        Set<String> listOfObjectFields = getAccessableFieldsByObject(objectApiName);

        if(!listOfObjectFields.isEmpty()) {
            //build list of lookup Ids based on lookup field name passed in from the list of records passed in
            Set<String> listOfOrderRecordIds = new Set<String>();
            for(Map<String, Object> orderRecord : listOfOrderRecordMapRecords) {
                if(orderRecord.get(LookupField) != null)listOfOrderRecordIds.add(string.valueOf(orderRecord.get(LookupField)));
            }

            if(!listOfOrderRecordIds.isEmpty()) {
                //query list of related records based on lookup field and object api name using list of lookup Ids from list of records 
                String query = 'SELECT '+ String.join(((Iterable<String>)listOfObjectFields), ',') + ' FROM ' + objectApiName +' WHERE Id IN :listOfOrderRecordIds  WITH SECURITY_ENFORCED';
                List<SObject> listOfRelatedRecordData = Database.query(query);
                
                if(!listOfRelatedRecordData.isEmpty()) {
                    /*add each related record to the records json response and add the Id to a list 
                    to add back to parent record map after*/
                    for(Map<String, Object> orderRecord : listOfOrderRecordMapRecords) {
                        //list to related record Ids to populate and add per order record 
                        Set<String> listOfRelatedRecordsIds = new Set<String> ();
                        for(SObject relatedRecord : listOfRelatedRecordData) { 
                            String orderRecordLookupId = (String)orderRecord.get(LookupField);
                            if(orderRecordLookupId == relatedRecord.get('Id')){
                                if(relatedRecord != null)responseJsonPayloadRecords.add((Object)relatedRecord);
                                listOfRelatedRecordsIds.add(string.valueOf(relatedRecord.get('Id'))); 
                            }
                        }
                        if(!listOfRelatedRecordsIds.isEmpty()) {
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

                            //add list of child ids to parent record if it is not a direct lookup relationship
                            Object value = orderRecord.get(objectApiName);
                            if(value == null) {
                                orderRecord.put(objectApiName,(Object)listOfRelatedRecordsIds);
                            }
                        }
                    }
                }
            }
        }
        return listOfOrderRecordMapRecords;
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
