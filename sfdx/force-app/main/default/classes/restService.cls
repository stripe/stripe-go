/*
You can hit this endpoint per org by authenticating into the org and makeing a POST request to 
https://{orgDomain}.my.salesforce.com/services/apexrest/batchApi

https://developer.salesforce.com/docs/atlas.en-us.apexcode.meta/apexcode/apex_classes_annotation_rest_resource.htm
*/

@RestResource(urlMapping='/batchApi')
global with sharing class restService {
    /*
    NOTE we accidentally used GET for the route instead of POST, which made it hard to pass a message body in the ruby SF library
         in order to get around packaging issues, we are leaving this around for now until we figure out how to remove it completely. 
         Strangely enough, you are not supposed to be able to tag a @RestResource class with the same HTTP verb annotations twice, but we are, and it works.
    TODO https://github.com/stripe/stripe-salesforce/issues/383
     */
    @HttpGet
    @HttpPost
    @ReadOnly
    global static void restServiceEvents() {
        RestResponse response = RestContext.response;
        response.responseBody = Blob.valueOf('This route has been deprecated, use the POST route instead.');
        return;
    }

    //custom exception
    public class NoOrderIdsOrProductIdsFoundException extends Exception {
    }
    public class BatchRestServiceException extends Exception {
    }

    //collects errors from service
    global static Set<Object> responseJsonPayloadErrors = new Set<Object>();

    //builds json paylod 
    global static Set<Object> responseJsonPayloadRecords = new Set<Object>();

    //stores list of order and product fields
    global static Set<String> listOfOrderFields = new Set<String>();
    global static Set<String> listOfProductFields = new Set<String>();
    private static Rest_Service_Helper helper = new Rest_Service_Helper();

    private static RelationshipMapper_QueryHelper rmQueryHelper = new RelationshipMapper_QueryHelper();

    // Marked TestVisible so we can enable it in tests.
    @TestVisible
    private static Debug_Helper debugger = new Debug_Helper();

    //readonly annotation is used to increase the SOQL limit from 50k to a 100k 
    @ReadOnly

    // using POST here to make it easier to send a JSON message body (the SF ruby gem does not support GET with body out of the box)
    @HttpPost
    global static void executeBatchApi() {
        debugger.debug('restService.executeBatchApi', 'start');
        RestRequest restRequest = RestContext.request;
        RestResponse response = RestContext.response;
        try {
            Boolean isCpqInstalled = utilities.isCpqEnabled();

            //get body of request and cast to map
            PlatformPayload payload = (PlatformPayload) JSON.deserialize(restRequest.requestBody.toString(), PlatformPayload.class);
            debugger.debug('payload', payload);

            if ((payload.order_ids == null || payload.order_ids.isEmpty()) && (payload.product2_ids == null || payload.product2_ids.isEmpty())) {
                throw new NoOrderIdsOrProductIdsFoundException('There were no order or product Ids sent with this request.');
            }

            if (payload.field_mapping_data != null) {
                rmQueryHelper.initializeRelationshipChains(payload.field_mapping_data);
            }

            //converting object lists to sets to remove dupes if any exist and allow binding in dynamic soql query
            Set<String> listOfOrderIds = (payload.order_ids != null) ? payload.order_ids : new Set<String>();
            Set<String> listOfProduct2Ids = (payload.product2_ids != null) ? payload.product2_ids : new Set<String>();

            if (payload.order_fields != null) {
                listOfOrderFields = payload.order_fields;
            }

            if (payload.product2_fields != null) {
                listOfProductFields = payload.product2_fields;
            }

            //if any of the field lists are empty get all fields 
            if (listOfOrderFields.isEmpty())
                listOfOrderFields = getAccessibleFields('Order');

            if (listOfProductFields.isEmpty())
                listOfProductFields = getAccessibleFields('Product2');

            List<SObject> listOfOrderRecords = new List<SObject>();
            if (!listOfOrderIds.isEmpty()) {

                String opportunityAmendedContrctQuoteFieldName = 'Opportunity.SBQQ__AmendedContract__r.SBQQ__Quote__c';
                Boolean isFieldAccessible = Schema.sObjectType.Order.fields.OpportunityId.isAccessible();
                if (!isCpqInstalled || !isFieldAccessible) {
                    opportunityAmendedContrctQuoteFieldName = '';
                }

                //get all the order records based on list order ids requested
                String strOrderIds = helper.setToQueryString(listOfOrderIds);
                listOfOrderRecords = (List<Order>)query(
                        listOfOrderFields,
                        opportunityAmendedContrctQuoteFieldName,
                        'Order',
                        'Id IN ' + strOrderIds,
                        true);

                //Used to hold all order records after converting to map
                Set<Map<String, Object>> listOfOrderRecordMapRecords = new Set<Map<String, Object>>();

                //if order records are found in add to order map list
                if (!listOfOrderRecords.isEmpty()) {
                    for (SObject orderRecord : listOfOrderRecords) {
                        //converting order to a map so we can add the related orderItems Id before adding the record to payload
                        Map<String, Object> orderRecordMap = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(orderRecord));

                        //adding order records map to list
                        listOfOrderRecordMapRecords.add(orderRecordMap);
                    }

                    //get related order records based on specific lookups existing in the targeted object field list(s)  
                    Set<String> listOfOrderItemFields = getAccessibleFields('OrderItem');
                    listOfOrderRecordMapRecords = addRelatedOrderRecordsToPayload(listOfOrderFields, listOfOrderItemFields, listOfOrderRecordMapRecords, listOfOrderIds, isCpqInstalled);

                    //adding modified order record to payload 
                    for (Map<String, Object> modifiedOrderRec : listOfOrderRecordMapRecords) {
                        if (modifiedOrderRec != null && !modifiedOrderRec.isEmpty()) responseJsonPayloadRecords.add((Object) modifiedOrderRec);
                    }
                }
            }

            String strProductIds = helper.setToQueryString(listOfProduct2Ids);
            List<SObject> listOfProductRecords = new List<SObject>();
            if (!listOfProduct2Ids.isEmpty()) {
                //get all the product records based on list product ids requested
                listOfProductRecords = (List<Product2>) query(listOfProductFields, 'Product2', 'Id IN ' + strProductIds);

                //Used to hold all product records after converting to map
                Set<Map<String, Object>> listOfProduct2RecordMapRecords = new Set<Map<String, Object>>();
                if (!listOfProductRecords.isEmpty()) {
                    for (SObject productRecord : listOfProductRecords) {
                        //converting order to a map so we can add the related orderItems Id before adding the record to payload
                        Map<String, Object> productRecordMap = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(productRecord));

                        //adding product records map to list
                        listOfProduct2RecordMapRecords.add(productRecordMap);
                    }

                    listOfProduct2RecordMapRecords = addRelatedProductRecordsToPayload(listOfProductFields, listOfProduct2RecordMapRecords, listOfProduct2Ids, isCpqInstalled);

                    //adding modified order record to payload 
                    for (Map<String, Object> modifiedProductRec : listOfProduct2RecordMapRecords) {
                        if (modifiedProductRec != null && !modifiedProductRec.isEmpty()) responseJsonPayloadRecords.add((Object) modifiedProductRec);
                    }
                }
            }

            if (listOfProductRecords.isEmpty() && listOfOrderRecords.isEmpty()) {
                throw new NoOrderIdsOrProductIdsFoundException('There were no order or product records found based on the list of order and product Ids sent. Here are the order Ids that were sent: ' + listOfOrderIds + '. Here are the product Ids that were sent: ' + listOfProduct2Ids);
            }
            /*create response payload with a records object containing a list of objects of all the records
            fetched and an error object containing and errors of in-accessable fields */
            response.addHeader('Content-Type', 'application/json');
            response.statusCode = 200;
            Map<String, Object> resBody = new Map<String, Object>();
            resBody.put('errors', responseJsonPayloadErrors);
            Map<String, List<Object>> orderDetails = new Map<String, List<Object>>();
            resBody.put('records', (Object) responseJsonPayloadRecords);

            if (payload.field_mapping_data != null) {
                resBody.put('data_field_mapping', rmQueryHelper.getFieldMappings());
            }

            String respBody = JSON.serialize(resBody);
            debugger.debug('response(success)', respBody);
            response.responseBody = Blob.valueOf(respBody);
            return;
        } catch (DmlException e) {
            //Catches any CRUD SOQL failure and adds to error response map and continues
            Map<String, Object> DmlExceptionMap = new Map<String, Object>{
                    'error_type' => 'DmlException',
                    'error_message' => e.getMessage() + 'Line Number: ' + string.valueOf(e.getLineNumber())
            };
            responseJsonPayloadErrors.add((Object) DmlExceptionMap);

        } catch (SObjectException e) {
            //Catches any SObject access failure and addes to error response map and continues
            Map<String, Object> SObjectExceptionMap = new Map<String, Object>{
                    'error_type' => 'SObjectException',
                    'error_message' => e.getMessage() + 'Line Number: ' + string.valueOf(e.getLineNumber())
            };
            responseJsonPayloadErrors.add((Object) SObjectExceptionMap);
        } catch (RelationshipMapper_QueryHelper.QHException e) {
            // Catches the relationship mapper exception for either inaccessible fields or unsupported relationships.
            response.statusCode = 500;
            Map<String, Object> resBody = new Map<String, Object>();
            resBody.put('error_type', 'MAPPING_ERROR');
            resBody.put('error_message', e.getLineNumber() + ': ' + e.getMessage());
            if (e.getMessage() == RelationshipMapper_QueryHelper.EXCEPTION_FIELD_INACCESSIBLE) {
                resBody.put('failed_fields', e.context.get('failed_fields'));
            }
            String respBody = JSON.serialize(resBody);
            debugger.debug('response(exception)', respBody);
            response.responseBody = Blob.valueOf(respBody);
            return;
        } catch (Exception e) {
            /*Catches any other exception and ends operation. returning the error message 
            and line number as the message param of response*/
            response.statusCode = 500;
            Map<String, Object> resBody = new Map<String, Object>();
            resBody.put('error_type', 'APEX_ERROR');
            resBody.put('error_message', e.getLineNumber() + ': ' + e.getMessage());
            String respBody = JSON.serialize(resBody);
            debugger.debug('response(exception)', respBody);
            response.responseBody = Blob.valueOf(respBody);
            return;
        }

        if (responseJsonPayloadErrors.isEmpty()) {
            response.statusCode = 400;
            response.addHeader('Content-Type', 'application/json');
            response.responseBody = Blob.valueOf('"no results"');
            debugger.debug('response(unknown)', 'no results');
        } else {
            response.statusCode = 500;
            String respBody = JSON.serialize(responseJsonPayloadErrors);
            debugger.debug('response(errors)', respBody);
            response.addHeader('Content-Type', 'application/json');
            response.responseBody = Blob.valueOf(respBody);
        }
    }

    //fetches all related object records that we care about
    private static Set<Map<String, Object>> addRelatedProductRecordsToPayload(Set<String> listOfProductFields, Set<Map<String, Object>> listOfProduct2RecordMapRecords, Set<String> listOfProduct2Ids, Boolean isCpqInstalled) {
        debugger.debug('restService.addRelatedProductRecordsToPayload(listOfProduct2Ids)', listOfProduct2Ids);
        String strProductIds = helper.setToQueryString(listOfProduct2Ids);
        //get related pricebook entry records
        Set<String> listOfPricebookEntryFields = getAccessibleFields('PricebookEntry');
        if (listOfPricebookEntryFields.contains('Product2Id')) {
            List<SObject> listOfPricebookEntryRecords = query(listOfPricebookEntryFields, 'PricebookEntry', 'Product2Id IN ' + strProductIds);
            if (!listOfPricebookEntryRecords.isEmpty()) {
                for (Map<String, Object> productRecord : listOfProduct2RecordMapRecords) {
                    String productRecordId = (String) productRecord.get('Id');
                    Set<String> listPricebookEntryIds = new Set<String>();
                    for (SObject pricbookEntryRecord : listOfPricebookEntryRecords) {
                        //converting pricebookEntry record to a map so we can add the related pricebookEntry Id before adding the record to payload
                        Map<String, Object> pricbookEntryRecordMap = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(pricbookEntryRecord));
                        if (productRecordId == pricbookEntryRecordMap.get('Product2Id')) {
                            //add each related pricebookEntry item id to a list to put in product record at the end
                            listPricebookEntryIds.add(string.valueOf(pricbookEntryRecordMap.get('Product2Id')));
                            listOfProduct2RecordMapRecords.add(pricbookEntryRecordMap);
                        }
                    }
                    productRecord.put('PricebookEntries', (Object) listPricebookEntryIds);
                }
            }
        }

        //get related product consumption schedule records
        Set<String> listOfAllConsumptionScheduleIds = new Set<String>();
        Set<String> listOfProductConsumptionScheduleFields = getAccessibleFields('ProductConsumptionSchedule');
        debugger.debug('restService.addRelatedProductRecordsToPayload(listOfProductConsumptionScheduleFields)', listOfProductConsumptionScheduleFields);

        if (listOfProductConsumptionScheduleFields.contains('ProductId')) {
            debugger.debug('restService.addRelatedProductRecordsToPayload(listOfProductConsumptionScheduleFields)', 'found ProductId');
            List<SObject> listOfProductConsumptionScheduleRecords = query(listOfProductConsumptionScheduleFields, 'ProductConsumptionSchedule', 'ProductId IN ' + strProductIds);
            if (!listOfProductConsumptionScheduleRecords.isEmpty()) {
                for (Map<String, Object> productRecord : listOfProduct2RecordMapRecords) {
                    debugger.debug('restService.addRelatedProductRecordsToPayload(productRecord)', JSON.serialize(productRecord));
                    Set<String> listOfProductConsumptionScheduleIds = new Set<String>();
                    Set<String> listOfConsumptionScheduleIds = new Set<String>();
                    for (SObject productConsumptionScheduleRecord : listOfProductConsumptionScheduleRecords) {
                        //converting pricebookEntry record to a map so we can add the related pricebookEntry Id before adding the record to payload
                        Map<String, Object> productConsumptionScheduleMap = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(productConsumptionScheduleRecord));
                        debugger.debug('restService.addRelatedProductRecordsToPayload(ProductId)', productConsumptionScheduleMap.get('ProductId'));

                        if (listOfProduct2Ids.contains((String)productConsumptionScheduleMap.get('ProductId'))) {
                            //add each related pricebookEntry item id to a list to put in proct record at the end
                            listOfProductConsumptionScheduleIds.add(string.valueOf(productConsumptionScheduleMap.get('Id')));
                            listOfProduct2RecordMapRecords.add(productConsumptionScheduleMap);
                            if (listOfProductConsumptionScheduleFields.contains('ConsumptionScheduleId') && productConsumptionScheduleMap.get('ConsumptionScheduleId') != null) {
                                listOfAllConsumptionScheduleIds.add((String) productConsumptionScheduleMap.get('ConsumptionScheduleId'));
                                listOfConsumptionScheduleIds.add((String) productConsumptionScheduleMap.get('ConsumptionScheduleId'));
                            }
                        }
                    }
                    if (!listOfProductConsumptionScheduleIds.isEmpty()) {
                        productRecord.put('ProductConsumptionSchedules', (Object) listOfProductConsumptionScheduleIds);
                    }

                    if (!listOfConsumptionScheduleIds.isEmpty()) {
                        productRecord.put('ConsumptionSchedules', (Object) listOfConsumptionScheduleIds);
                    }
                }
            }
        }

        //get related consumption schedule and rate records 
        if (!listOfAllConsumptionScheduleIds.isEmpty()) {
            listOfProduct2RecordMapRecords.addAll(getRelatedConsumptionScheduleAndRateRecords(listOfProduct2RecordMapRecords, listOfAllConsumptionScheduleIds));
        }

        return listOfProduct2RecordMapRecords;
    }


    //fetches all related object records that we care about based on the order record passed in 
    private static Set<Map<String, Object>> addRelatedOrderRecordsToPayload(Set<String> listOfOrderFields, Set<String> listOfOrderItemFields, Set<Map<String, Object>> listOfOrderRecordMapRecords, Set<String> listOfOrderIds, Boolean isCpqInstalled) {
        //check if CPQ is installed before getting data from objects in the package 
        String orderIds = helper.setToQueryString(listOfOrderIds);

        if (isCpqInstalled) {

            //get list of all quote Ids that our orders look up to
            Set<String> listOfOrderRecordQuoteId = new Set<String>();
            String cpqQuoteObjectApiName = 'SBQQ__Quote__c';
            if (listOfOrderFields.contains(cpqQuoteObjectApiName)) {
                for (Map<String, Object> orderRecordMap : listOfOrderRecordMapRecords) {
                    if (orderRecordMap.get(cpqQuoteObjectApiName) != null) {
                        String quoteId = (String) orderRecordMap.get(cpqQuoteObjectApiName);
                        listOfOrderRecordQuoteId.add(quoteId);
                    }
                }

                //get all related initial orders and order amendments             
                Set<String> listOfInitialOrderRecordQuoteId = new Set<String>();
                Set<String> listOfInitialOrderRecordAmmendedContractId = new Set<String>();
                if (listOfOrderFields.contains(cpqQuoteObjectApiName)) {
                    Set<String> listOfAllOrderRelatedOpportunityQuoteIds = buildListOfAllOpportunityQuoteIds(listOfOrderRecordMapRecords);
                    if (listOfAllOrderRelatedOpportunityQuoteIds.size() > 0) {
                        String oroqIds = helper.setToQueryString(listOfAllOrderRelatedOpportunityQuoteIds);
                        List<SObject> listOfInitialOrderRecords = query(listOfOrderFields, 'Order', 'SBQQ__Quote__c IN ' + oroqIds + ' AND Id NOT IN ' + orderIds);

                        Map<String, Object> updatedInitalOrderAndQouteListMap = addInitialOrderToPayload(listOfInitialOrderRecords, listOfOrderRecordMapRecords, listOfInitialOrderRecordQuoteId, listOfOrderIds);
                        listOfOrderRecordMapRecords = (Set<Map<String, Object>>) updatedInitalOrderAndQouteListMap.get('listOfOrderRecordMapRecords');
                        listOfInitialOrderRecordQuoteId = (Set<String>) updatedInitalOrderAndQouteListMap.get('listOfInitialOrderRecordQuoteId');
                        listOfOrderIds = (Set<String>) updatedInitalOrderAndQouteListMap.get('listOfOrderIds');
                    }

                    Set<String> listOfRelatetedContractOrderRecordIds = buildListOfRelatedContractIds(listOfOrderIds, listOfOrderRecordQuoteId);

                    if (listOfRelatetedContractOrderRecordIds.size() > 0) {
                        String rcorIds = helper.setToQueryString(listOfRelatetedContractOrderRecordIds);
                        List<SObject> listOfamendedOrderRecords = query(
                                listOfOrderFields,
                                'Order',
                                'Opportunity.SBQQ__AmendedContract__c IN ' + rcorIds + ' AND Id NOT IN ' + orderIds);

                        Map<String, Object> updatedAmendedOrderAndQouteListMap = addOrderAmendmentsToPayload(listOfamendedOrderRecords, listOfOrderRecordMapRecords, listOfInitialOrderRecordAmmendedContractId, listOfOrderIds);
                        listOfOrderRecordMapRecords = (Set<Map<String, Object>>) updatedAmendedOrderAndQouteListMap.get('listOfOrderRecordMapRecords');
                        listOfInitialOrderRecordAmmendedContractId = (Set<String>) updatedAmendedOrderAndQouteListMap.get('listOfInitialOrderRecordAmmendedContractId');
                        listOfOrderIds = (Set<String>) updatedAmendedOrderAndQouteListMap.get('listOfOrderIds');                    }
                }
            }

            //get related records based on the lookup to quote from the order
            List<SObject> listOfQuoteRecords = new List<SObject>();
            Set<String> listOfQuoteFields = getAccessibleFields(cpqQuoteObjectApiName);
            if (!listOfOrderRecordQuoteId.isEmpty()) {
                String orqIds = helper.setToQueryString(listOfOrderRecordQuoteId);
                listOfQuoteRecords = query(listOfQuoteFields, cpqQuoteObjectApiName, 'Id IN ' + orqIds);
            }

            //gets related quote record and specific quote child records based on the quote this order looks up to 
            Set<Map<String, Object>> listOfQuoteRecordMaps = new Set<Map<String, Object>>();
            if (!listOfQuoteRecords.isEmpty()) {
                for (Map<String, Object> orderRecordMap : listOfOrderRecordMapRecords) {
                    List<Object> relatedQuoteRecordIdList = new List<Object>();
                    for (SObject quoteRecord : listOfQuoteRecords) {
                        if (orderRecordMap.get(cpqQuoteObjectApiName) == quoteRecord.get('Id')) {
                            Map<String, Object> quoteRecordMap = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(quoteRecord));
                            listOfQuoteRecordMaps.add(quoteRecordMap);
                            relatedQuoteRecordIdList.add(quoteRecordMap.get('Id'));
                        }
                    }

                    if (relatedQuoteRecordIdList.isEmpty()) {
                        continue;
                    }

                    //adds list of the related quote ID to orderRecords
                    orderRecordMap.put('Quotes', (Object) relatedQuoteRecordIdList);
                }
            }

            if (listOfQuoteRecordMaps.isEmpty()) {
                return listOfOrderRecordMapRecords;
            }
            //gets related primary contact records based on the contacts the list quote records looks up to 
            if (listOfQuoteFields.contains('SBQQ__PrimaryContact__c')) {
                listOfQuoteRecordMaps = addRelatedRecordToPayload(listOfQuoteRecordMaps, 'Contact', 'SBQQ__PrimaryContact__c');
            }

            //gets related opportunity records based on the opportunities the list quote records looks up to 
            if (listOfQuoteFields.contains('SBQQ__Opportunity2__c')) {
                listOfQuoteRecordMaps = addRelatedRecordToPayload(listOfQuoteRecordMaps, 'Opportunity', 'SBQQ__Opportunity2__c');
            }

            //gets related account records based on the accounts the list quote records looks up to 
            if (listOfQuoteFields.contains('SBQQ__Account__c')) {
                listOfQuoteRecordMaps = addRelatedRecordToPayload(listOfQuoteRecordMaps, 'Account', 'SBQQ__Account__c');
            }

            for (Map<String, Object> modifiedQuoteRecordMap : listOfQuoteRecordMaps) {
                //add modified quote record with lookup ID lists to records param in payload
                if (modifiedQuoteRecordMap != null && !modifiedQuoteRecordMap.isEmpty())responseJsonPayloadRecords.add((Object) modifiedQuoteRecordMap);
            }
        }

        if (listOfOrderFields.contains('ContractId')) {
            //pulls out all contract Ids from order records into seperate list
            Set<String> listOfContractRecordIds = new Set<String>();
            for (Map<String, Object> orderRecordMap : listOfOrderRecordMapRecords) {
                if (orderRecordMap.get('ContractId') != null)listOfContractRecordIds.add((String) orderRecordMap.get('ContractId'));
            }

            /*gets related orders for the contract that are not equal to our current list of order
            ids but lookup to any of the same contracts within our current order list*/
            if (!listOfContractRecordIds.isEmpty()) {
                String crIds = helper.setToQueryString(listOfContractRecordIds);
                List<Order> listOfContractOrderRecords = query(
                        listOfOrderFields,
                        'Order',
                        'ContractId IN ' + crIds + ' AND Id NOT IN ' + orderIds);

                /*convert any orders that lookup to the same contracts as the orders in listOfOrderRecordMapRecords. 
                Add the records to a new list to recurse on and get the Ids to add to our order record*/
                Set<Map<String, Object>> listOfContractOrderRecordMaps = new Set<Map<String, Object>>();
                Map<String, Object> contractOrderRecordMap = new Map<String, Object>();
                Set<String> listOfContractOrderRecordIds = new Set<String>();//delete
                for (Map<String, Object> orderRecordMap : listOfOrderRecordMapRecords) {
                    for (SObject contractOrderRecord : listOfContractOrderRecords) {
                        if (orderRecordMap.get('ContractId') != contractOrderRecord.get('ContractId')) {
                            continue;
                        }
                        contractOrderRecordMap = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(contractOrderRecord));
                        //adding all other orders that look up to the same contract to our list of order records to get related data below
                        listOfOrderRecordMapRecords.add(contractOrderRecordMap);

                        //creating seperate list so we can add back to each order record and the relation to the contract is shown
                        listOfContractOrderRecordIds.add((String) contractOrderRecordMap.get('Id'));

                        //adding new ids contract order ids to existing list so they will be excluded when we recurse
                        listOfOrderIds.add((String) contractOrderRecordMap.get('Id'));
                    }

                    if (listOfContractOrderRecordIds.isEmpty()) {
                        continue;
                    }
                    orderRecordMap.put('ContractOrders', (Object) listOfContractOrderRecordIds);
                }
                //gets related contracts record based the contracts these orders look up to
                listOfOrderRecordMapRecords = addRelatedRecordToPayload(listOfOrderRecordMapRecords, 'Contract', 'ContractId');
            }
        }
        if (!listOfOrderItemFields.isEmpty()) {
            //get all related order items based on the orders contained in list
            listOfOrderRecordMapRecords = getRelatedOrderItemRecordsById(listOfOrderRecordMapRecords, listOfOrderItemFields, listOfOrderIds, isCpqInstalled);
        }
        //gets related account records based on the accounts any order in our list looks up to 
        if (listOfOrderFields.contains('AccountId')) {
            listOfOrderRecordMapRecords = addRelatedRecordToPayload(listOfOrderRecordMapRecords, 'Account', 'AccountId');
        }
        //gets related contact records based on the contacts any order in our list looks up to 
        if (listOfOrderFields.contains('BillToContactId')) {
            listOfOrderRecordMapRecords = addRelatedRecordToPayload(listOfOrderRecordMapRecords, 'Contact', 'BillToContactId');
        }
        //gets related opportunity records based on the opportunities any order in our list looks up to 
        if (listOfOrderFields.contains('OpportunityId')) {
            listOfOrderRecordMapRecords = addRelatedRecordToPayload(listOfOrderRecordMapRecords, 'Opportunity', 'OpportunityId');
        }
        //gets related pricebook records based on the pricebooks sny order in our list looks up to 
        if (listOfOrderFields.contains('Pricebook2Id')) {
            listOfOrderRecordMapRecords = addRelatedRecordToPayload(listOfOrderRecordMapRecords, 'Pricebook2', 'Pricebook2Id');
        }

        return listOfOrderRecordMapRecords;
    }

    private static Set<String> buildListOfAllOpportunityQuoteIds(Set<Map<String, Object>> listOfOrderRecordMapRecords) {
        Set<String> listOfAllOrderRelatedOpportunityQuoteIds = new Set<String>();
        for (Map<String, Object> orderRecordMapRecord : listOfOrderRecordMapRecords) {
            if (orderRecordMapRecord.get('Opportunity') == null) {
                return listOfAllOrderRelatedOpportunityQuoteIds;
            }

            Map<String, Object> relatatedOrderOppertunityMap = (Map<String, Object>) orderRecordMapRecord.get('Opportunity');
            Set<String> listOfInitialOrderIds = new Set<String>();
            if (relatatedOrderOppertunityMap.get('SBQQ__AmendedContract__r') == null) {
                return listOfAllOrderRelatedOpportunityQuoteIds;
            }
            Map<String, Object> relatatedOrderContractMap = (Map<String, Object>) relatatedOrderOppertunityMap.get('SBQQ__AmendedContract__r');

            if (relatatedOrderContractMap.get('SBQQ__Quote__c') == null) {
                return listOfAllOrderRelatedOpportunityQuoteIds;
            }
            listOfAllOrderRelatedOpportunityQuoteIds.add(String.valueOf(relatatedOrderContractMap.get('SBQQ__Quote__c')));
        }
        return listOfAllOrderRelatedOpportunityQuoteIds;
    }

    private static Set<String> buildListOfRelatedContractIds(Set<String> listOfOrderIds, Set<String> listOfOrderRecordQuoteId) {
        Set<String> listOfRelatetedContractOrderRecordIds = new Set<String>();
        String orderIds = helper.setToQueryString(listOfOrderIds);
        String orqIds = helper.setToQueryString(listOfOrderRecordQuoteId);
        Set<String> contractFields = new Set<String>{'Id', 'SBQQ__Quote__c'};
        List<SObject> listOfAmendedOrderContractRecords = query(
                contractFields,
                'Contract',
                'SBQQ__Quote__c IN ' + orqIds + ' AND Id NOT IN ' + orderIds);

        if (listOfAmendedOrderContractRecords.isEmpty()) {
            return listOfRelatetedContractOrderRecordIds;
        }
        for (SObject relatedContractRecord : listOfAmendedOrderContractRecords) {
            Map<String, Object> relatedContractRecordMap = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(relatedContractRecord));
            if (relatedContractRecord.get('SBQQ__Quote__c') == null) {
                continue;
            }

            listOfRelatetedContractOrderRecordIds.add((String) relatedContractRecord.get('Id'));
        }
        return listOfRelatetedContractOrderRecordIds;
    }

    private static Map<String, Object> addInitialOrderToPayload(List<SObject> listOfInitialOrderRecords, Set<Map<String, Object>> listOfOrderRecordMapRecords, Set<String> listOfInitialOrderRecordQuoteId, Set<String> listOfOrderIds) {
        if (listOfInitialOrderRecords.isEmpty()) {
            return new Map<String, Object>{
                    'listOfOrderRecordMapRecords' => listOfOrderRecordMapRecords,
                    'listOfInitialOrderRecordQuoteId' => listOfInitialOrderRecordQuoteId,
                    'listOfOrderIds' => listOfOrderIds
            };
        }
        for (Map<String, Object> orderRecordMapRecord : listOfOrderRecordMapRecords) {
            Map<String, Object> relatatedOrderOppertunityMap = (Map<String, Object>) orderRecordMapRecord.get('Opportunity');

            Set<String> listOfInitialOrderIds = new Set<String>();

            if (relatatedOrderOppertunityMap == null) {
                continue;
            }

            Map<String, Object> relatatedOrderContractMap = (Map<String, Object>) relatatedOrderOppertunityMap.get('SBQQ__AmendedContract__r');
            if (relatatedOrderContractMap == null) {
                continue;
            }

            for (SObject initialOrderRecord : listOfInitialOrderRecords) {
                if (relatatedOrderContractMap.get('SBQQ__Quote__c') != initialOrderRecord.get('SBQQ__Quote__c')) {
                    continue;
                }
                Map<String, Object> initialOrderRecordMap = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(initialOrderRecord));
                listOfOrderRecordMapRecords.add(initialOrderRecordMap);

                listOfInitialOrderIds.add(String.valueOf(initialOrderRecord.get('Id')));
                listOfOrderIds.add(String.valueOf(initialOrderRecord.get('Id')));
                listOfInitialOrderRecordQuoteId.add(string.valueOf(initialOrderRecord.get('SBQQ__Quote__c')));
            }

            if (!listOfInitialOrderIds.isEmpty() && (listOfInitialOrderIds.size() > 1) && Test.isRunningTest()) {
                throw new BatchRestServiceException('More than one related initial order ID was found: ' + String.valueOf(listOfInitialOrderRecordQuoteId));
            }

            if (!listOfInitialOrderIds.isEmpty() && (listOfInitialOrderIds.size() > 1) && !Test.isRunningTest()) {
                try {
                    throw new BatchRestServiceException('More than one related initial order ID was found: ' + String.valueOf(listOfInitialOrderRecordQuoteId));
                } catch (BatchRestServiceException e) {
                    Sentry.record(e);
                }
            }

            if (!listOfInitialOrderIds.isEmpty()) {
                orderRecordMapRecord.put('InitialOrderId', (Object) listOfInitialOrderIds);
            }

            if (!listOfInitialOrderRecordQuoteId.isEmpty()) {
                orderRecordMapRecord.put('InitialOrderQuotes', (Object) listOfInitialOrderRecordQuoteId);
            }
        }

        return new Map<String, Object>{
                'listOfOrderRecordMapRecords' => listOfOrderRecordMapRecords,
                'listOfInitialOrderRecordQuoteId' => listOfInitialOrderRecordQuoteId,
                'listOfOrderIds' => listOfOrderIds
        };
    }

    private static Map<String, Object> addOrderAmendmentsToPayload(List<SObject> listOfamendedOrderRecords, Set<Map<String, Object>> listOfOrderRecordMapRecords, Set<String> listOfInitialOrderRecordAmmendedContractId, Set<String> listOfOrderIds) {
        if (listOfamendedOrderRecords.isEmpty()) {
            return new Map<String, Object>{
                    'listOfOrderRecordMapRecords' => listOfOrderRecordMapRecords,
                    'listOfInitialOrderRecordAmmendedContractId' => listOfInitialOrderRecordAmmendedContractId,
                    'listOfOrderIds' => listOfOrderIds
            };
        }

        for (Map<String, Object> orderRecordMapRecord : listOfOrderRecordMapRecords) {
            Map<String, Object> relatatedOrderOppertunityMap = (Map<String, Object>) orderRecordMapRecord.get('Opportunity');
            Map<String, Object> relatatedOrderAmendedContractMap = (Map<String, Object>) relatatedOrderOppertunityMap.get('SBQQ__AmendedContract__r');
            Set<String> listOfAmendedOrderRecordIds = new Set<String>();

            for (SObject amendedOrderRecord : listOfamendedOrderRecords) {
                Map<String, Object> amendedOrderRecordMap = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(amendedOrderRecord));
                listOfOrderRecordMapRecords.add(amendedOrderRecordMap);

                listOfInitialOrderRecordAmmendedContractId.add(string.valueOf(amendedOrderRecord.get('ContractId')));
                listOfAmendedOrderRecordIds.add(String.valueOf(amendedOrderRecord.get('Id')));
                listOfOrderIds.add(String.valueOf(amendedOrderRecord.get('Id')));

            }

            if (!listOfAmendedOrderRecordIds.isEmpty()) {
                orderRecordMapRecord.put('AmendedOrders', (Object) listOfAmendedOrderRecordIds);
            }

            if (!listOfInitialOrderRecordAmmendedContractId.isEmpty()) {
                orderRecordMapRecord.put('AmendedContracts', (Object) listOfInitialOrderRecordAmmendedContractId);
            }
        }

        return new Map<String, Object>{
                'listOfOrderRecordMapRecords' => listOfOrderRecordMapRecords,
                'listOfInitialOrderRecordAmmendedContractId' => listOfInitialOrderRecordAmmendedContractId,
                'listOfOrderIds' => listOfOrderIds
        };
    }


    //gets all related order item records and related child records that we care about
    private static Set<Map<String, Object>> getRelatedOrderItemRecordsById(Set<Map<String, Object>> listOfOrderRecordMapRecords, Set<String> listOfOrderItemFields, Set<String> listOfOrderIds, Boolean isCpqInstalled) {
        //get related order Items based on the Ids in the order record list
        String orderIds = helper.setToQueryString(listOfOrderIds);
        List<OrderItem> listOfOrderItemData = (List<OrderItem>) query(listOfOrderItemFields, 'OrderItem', 'OrderId IN ' + orderIds);

        /*used to build list of order items and add a list of lookup Ids to each order item record based on 
        each related record found*/
        Set<Map<String, Object>> listOfOrderItemRecordMapRecords = new Set<Map<String, Object>>();

        for (Map<String, Object> orderRecord : listOfOrderRecordMapRecords) {
            //create list to get all Order Item Ids that are related to the Order record passed in
            Set<String> listOfOrderItemIds = new Set<String>();

            String orderRecordId = (String) orderRecord.get('Id');

            for (SObject orderItem : listOfOrderItemData) {
                if (orderRecordId == orderItem.get('OrderId')) {
                    //convert order item record to map so we can add related record Ids to it if they exist
                    Map<String, Object> orderItemMap = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(orderItem));

                    //add each related order item id to a list to put in order record at the end
                    listOfOrderItemIds.add(string.valueOf(orderItemMap.get('Id')));

                    //add each related order item to list
                    listOfOrderItemRecordMapRecords.add(orderItemMap);
                }
            }
            orderRecord.put('OrderItems', (Object) listOfOrderItemIds);
        }

        //gets related product this order item looks up to 
        if (listOfOrderItemFields.contains('Product2Id')) {
            listOfOrderItemRecordMapRecords = addRelatedRecordToPayload(listOfOrderItemRecordMapRecords, 'Product2', 'Product2Id');
        }

        debugger.debug('restService.getRelatedOrderItemRecordsById(productIds)', 'begin');
        Set<String> productIds = new Set<String>();
        for (Map<String, Object> item : listOfOrderItemRecordMapRecords) {
            Map<String, Object> attrs = (Map<String, Object>) item.get('attributes');
            if (attrs == null || attrs.isEmpty()) {
                debugger.debug('restService.getRelatedOrderItemRecordsById(productIds)', 'no attributes');
                continue;
            }
            String objType = (String) attrs.get('type');
            debugger.debug('restService.getRelatedOrderItemRecordsById(productIds)', objType);
            if (objType == 'Product2') {
                productIds.add((String) item.get('Id'));
            }
            if (objType == 'OrderItem') {
                productIds.add((String) item.get('Product2Id'));
            }
        }

        debugger.debug('restService.getRelatedOrderItemRecordsById(productIds)', productIds);
        Set<String> productFields = helper.getAccessibleFields('Product2', responseJsonPayloadErrors);
        listOfOrderItemRecordMapRecords = addRelatedProductRecordsToPayload(productFields, listOfOrderItemRecordMapRecords, productIds, utilities.isCpqEnabled());

        //gets related price book entry this order item looks up to 
        if (listOfOrderItemFields.contains('PricebookEntryId')) {
            listOfOrderItemRecordMapRecords = addRelatedRecordToPayload(listOfOrderItemRecordMapRecords, 'PricebookEntry', 'PricebookEntryId');
        }

        //check if CPQ is installed before getting data from objects in the package 
        if (!isCpqInstalled) {
            //adds each modified order record to payload
            for (Map<String, Object> modifiedOrderItem : listOfOrderItemRecordMapRecords) {
                if (modifiedOrderItem != null && !modifiedOrderItem.isEmpty()) responseJsonPayloadRecords.add((Object) modifiedOrderItem);
            }

            return listOfOrderRecordMapRecords;
        }

        //get related order order item consumtption schedule records
        Set<String> listOfAllOrderItemConsumptionScheduleIds = new Set<String>();
        Set<String> listOfAllConsumptionScheduleIds = new Set<String>();
        Set<String> listOfOrderItemConsumptionScheduleFields = getAccessibleFields('SBQQ__OrderItemConsumptionSchedule__c');

        List<SObject> listOfOrderItemConsumptionScheduleRecords;
        if (listOfOrderItemConsumptionScheduleFields.contains('SBQQ__OrderItem__c')) {
            Set<String> allOrderItemIds = new Set<String>();
            for (Map<String, Object> orderRecord : listOfOrderRecordMapRecords) {
                Set<String> orderItemIds = (Set<String>) orderRecord.get('OrderItems');
                if (orderItemIds != null) {
                    allOrderItemIds.addAll(orderItemIds);
                }
            }
            String strAllOrderItemIds = helper.setToQueryString(allOrderItemIds);
            listOfOrderItemConsumptionScheduleRecords = query(listOfOrderItemConsumptionScheduleFields, 'SBQQ__OrderItemConsumptionSchedule__c', 'SBQQ__OrderItem__c IN ' + strAllOrderItemIds);

            if (!listOfOrderItemConsumptionScheduleRecords.isEmpty()) {
                for (Map<String, Object> orderItemRecord : listOfOrderItemRecordMapRecords) {
                    String orderItemRecordId = (String) orderItemRecord.get('Id');
                    Set<String> listOfConsumptionScheduleIds = new Set<String>();
                    for (SObject orderItemConsumptionScheduleRecord : listOfOrderItemConsumptionScheduleRecords) {
                        //converting orderItem record to a map so we can add the related orderItemConsumptionSchedule Id before adding the record to payload
                        Map<String, Object> orderItemConsumptionScheduleRecordMap = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(orderItemConsumptionScheduleRecord));
                        if (orderItemRecordId == orderItemConsumptionScheduleRecordMap.get('SBQQ__OrderItem__c')) {
                            //add each related orderItemConsumptionSchedule item id to a list to put in orderItem record at the end
                            listOfConsumptionScheduleIds.add(string.valueOf(orderItemConsumptionScheduleRecordMap.get('Id')));
                            listOfAllOrderItemConsumptionScheduleIds.add(string.valueOf(orderItemConsumptionScheduleRecordMap.get('Id')));
                            listOfAllConsumptionScheduleIds.add(string.valueOf(orderItemConsumptionScheduleRecordMap.get('SBQQ__ConsumptionSchedule__c')));
                        }
                    }
                    if (!listOfConsumptionScheduleIds.isEmpty()) {
                        orderItemRecord.put('OrderItemConsumptionSchedules', (Object) listOfConsumptionScheduleIds);
                    }
                }
            }
        }

        //get related consumtption schedule records
        if (!listOfAllOrderItemConsumptionScheduleIds.isEmpty()) {
            //Set<String> listOfOrderItemConsumptionScheduleFields
            String oicsIds = helper.setToQueryString(listOfAllOrderItemConsumptionScheduleIds);
            Set<String> listOfOrderItemConsumptionRateFields = getAccessibleFields('SBQQ__OrderItemConsumptionRate__c');
            if (listOfOrderItemConsumptionRateFields.contains('SBQQ__OrderItemConsumptionSchedule__c')) {
                List<SObject> listOfOrderItemConsumptionRateRecords = query(
                        listOfOrderItemConsumptionRateFields,
                        'SBQQ__OrderItemConsumptionRate__c',
                        'SBQQ__OrderItemConsumptionSchedule__c IN ' + oicsIds);

                if (!listOfOrderItemConsumptionRateRecords.isEmpty()) {
                    for (SObject orderItemConsumptionScheduleRecord : listOfOrderItemConsumptionScheduleRecords) {
                        Map<String, Object> orderItemConsumptionScheduleRecordMap = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(orderItemConsumptionScheduleRecord));
                        String orderItemConsumptionScheduleRecordId = (String) orderItemConsumptionScheduleRecordMap.get('Id');  //SBQQ__ConsumptionSchedule__c
                        Set<String> listOfConsumptionRateIds = new Set<String>();
                        for (SObject orderItemConsumptionRate : listOfOrderItemConsumptionRateRecords) {
                            //converting orderItem record to a map so we can add the related orderItemConsumptionSchedule Id before adding the record to payload

                            Map<String, Object> orderItemConsumptionRateMap = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(orderItemConsumptionRate));
                            if (orderItemConsumptionScheduleRecordId == orderItemConsumptionRateMap.get('SBQQ__OrderItemConsumptionSchedule__c')) {
                                //add each related orderItemConsumptionSchedule item id to a list to put in orderItem record at the end
                                listOfConsumptionRateIds.add(string.valueOf(orderItemConsumptionRateMap.get('Id')));
                            }
                            listOfOrderItemRecordMapRecords.add(orderItemConsumptionRateMap);

                        }
                        orderItemConsumptionScheduleRecordMap.put('OrderItemConsumptionRates', (Object) listOfConsumptionRateIds);
                        listOfOrderItemRecordMapRecords.add(orderItemConsumptionScheduleRecordMap);
                    }
                }
            }
            listOfOrderItemRecordMapRecords = getRelatedConsumptionScheduleAndRateRecords(listOfOrderItemRecordMapRecords, listOfAllConsumptionScheduleIds);
        }

        //gets related contract this order item looks up to 
        if (listOfOrderItemFields.contains('SBQQ__Contract__c')) {
            listOfOrderItemRecordMapRecords = addRelatedRecordToPayload(listOfOrderItemRecordMapRecords, 'Contract', 'SBQQ__Contract__c');
        }
        //gets related contract line item this order item looks up to 
        if (listOfOrderItemFields.contains('SBQQSC__ContractLineItem__c')) {
            listOfOrderItemRecordMapRecords = addRelatedRecordToPayload(listOfOrderItemRecordMapRecords, 'ContractLineItem', 'SBQQSC__ContractLineItem__c');
        }
        //gets related quote line this order item looks up to 
        if (listOfOrderItemFields.contains('SBQQ__QuoteLine__c')) {
            listOfOrderItemRecordMapRecords = addRelatedRecordToPayload(listOfOrderItemRecordMapRecords, 'SBQQ__QuoteLine__c', 'SBQQ__QuoteLine__c');
        }
        //gets related subscribtion this order item looks up to 
        if (listOfOrderItemFields.contains('SBQQ__Subscription__c')) {
            listOfOrderItemRecordMapRecords = addRelatedRecordToPayload(listOfOrderItemRecordMapRecords, 'SBQQ__Subscription__c', 'SBQQ__Subscription__c');
        }


        //adds each modified order record to payload
        for (Map<String, Object> modifiedOrderItem : listOfOrderItemRecordMapRecords) {
            if (modifiedOrderItem != null && !modifiedOrderItem.isEmpty()) responseJsonPayloadRecords.add((Object) modifiedOrderItem);
        }

        return listOfOrderRecordMapRecords;
    }

    private static Set<Map<String, Object>> getRelatedConsumptionScheduleAndRateRecords(Set<Map<String, Object>> listOfObjectRecordMapRecords, Set<String> listOfAllConsumptionScheduleIds) {
        List<SObject> listOfConsumptionScheduleRecords;
        String acsIds = helper.setToQueryString(listOfAllConsumptionScheduleIds);
        if (!listOfAllConsumptionScheduleIds.isEmpty()) {
            //get related consumption schedule records
            Set<String> listOfConsumptionScheduleFields = getAccessibleFields('ConsumptionSchedule');
            listOfConsumptionScheduleRecords = query(
                    listOfConsumptionScheduleFields,
                    'ConsumptionSchedule',
                    'Id IN ' + acsIds);

            //get related consumption rates records
            Set<String> listOfConsumptionRateFields = getAccessibleFields('ConsumptionRate');
            List<SObject> listOfConsumptionRateRecords = query(
                    listOfConsumptionRateFields,
                    'ConsumptionRate',
                    'ConsumptionScheduleId IN ' + acsIds);

            if (!listOfConsumptionRateRecords.isEmpty()) {
                for (SObject consumptionScheduleRecord : listOfConsumptionScheduleRecords) {
                    Map<String, Object> consumptionScheduleRecordMap = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(consumptionScheduleRecord));
                    Set<String> listOfConsumptionRateIds = new Set<String>();
                    for (SObject consumptionRateRecord : listOfConsumptionRateRecords) {
                        Map<String, Object> consumptionRateRecordMap = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(consumptionRateRecord));
                        listOfObjectRecordMapRecords.add(consumptionRateRecordMap);
                        if (consumptionScheduleRecordMap.get('Id') == consumptionRateRecordMap.get('ConsumptionScheduleId')) {
                            listOfConsumptionRateIds.add(String.valueOf(consumptionRateRecordMap.get('Id')));
                        }
                    }
                    consumptionScheduleRecordMap.put('ConsumptionRates', (Object) listOfConsumptionRateIds);
                    listOfObjectRecordMapRecords.add(consumptionScheduleRecordMap);
                }
            }
        }
        return listOfObjectRecordMapRecords;
    }

    /* this function gets the related records based on the list of records, object API name and lookupfield passed in.
    We then add a list of Ids with the corresponding plural object name to each record passed in and return to 
    calling function for insertion into the payload */
    private static Set<Map<String, Object>> addRelatedRecordToPayload(Set<Map<String, Object>> listOfOrderRecordMapRecords, String objectApiName, String LookupField) {
        //gets list of all fields we have access to for an object
        Set<String> listOfObjectFields = getAccessibleFields(objectApiName);

        if (listOfObjectFields.isEmpty()) {
            return listOfOrderRecordMapRecords;
        }
        //build list of lookup Ids based on lookup field name passed in from the list of records passed in
        Set<String> listOfOrderRecordIds = new Set<String>();
        for (Map<String, Object> orderRecord : listOfOrderRecordMapRecords) {
            if (orderRecord.get(LookupField) != null)listOfOrderRecordIds.add(String.valueOf(orderRecord.get(LookupField)));
        }

        if (listOfOrderRecordIds.isEmpty()) {
            return listOfOrderRecordMapRecords;
        }

        //query list of related records based on lookup field and object api name using list of lookup Ids from list of records
        String recordIds = helper.setToQueryString(listOfOrderRecordIds);
        List<SObject> listOfRelatedRecordData = query(listOfObjectFields, objectApiName, 'Id IN ' + recordIds);

        if (listOfRelatedRecordData.isEmpty()) {
            return listOfOrderRecordMapRecords;
        }

        /*add each related record to the records json response and add the Id to a list 
        to add back to parent record map after*/
        for (Map<String, Object> orderRecord : listOfOrderRecordMapRecords) {
            //list to related record Ids to populate and add per order record 
            Set<String> listOfRelatedRecordsIds = new Set<String> ();
            for (SObject relatedRecord : listOfRelatedRecordData) {
                String orderRecordLookupId = (String) orderRecord.get(LookupField);
                if (orderRecordLookupId == relatedRecord.get('Id')) {
                    Map<String, Object> relatedRecordMap = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(relatedRecord));
                    responseJsonPayloadRecords.add((Object) relatedRecord);
                    listOfRelatedRecordsIds.add(String.valueOf(relatedRecord.get('Id')));
                }
            }

            if (listOfRelatedRecordsIds.isEmpty()) {
                continue;
            }

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
            if (value == null) {
                orderRecord.put(objectApiName, (Object) listOfRelatedRecordsIds);
            }
        }
        return listOfOrderRecordMapRecords;
    }

    private static Set<String> getAccessibleFields(String objectApiName) {
        return helper.getAccessibleFields(objectApiName, responseJsonPayloadErrors);
    }

    private static List<SObject> query(Set<String> fields, String additionalField, String objectName, String whereClause, Boolean withSecurity) {
        if (additionalField != null && additionalField != '') {
            Set<String> allFields = new Set<String>(fields);
            allFields.add(additionalField);
            return query(allFields, objectName, whereClause, withSecurity);
        }

        return query(fields, objectName, whereClause, withSecurity);
    }

    private static List<SObject> query(Set<String> fields, Set<String> additionalFields, String objectName, String whereClause, Boolean withSecurity) {
        if (additionalFields != null && additionalFields.isEmpty() == false) {
            Set<String> allFields = new Set<String>(fields);
            allFields.addAll(additionalFields);
            return query(allFields, objectName, whereClause, withSecurity);
        }

        return query(fields, objectName, whereClause, withSecurity);
    }

    private static List<SObject> query(Set<String> fields, String objectName, String whereClause, Boolean withSecurity) {
        Set<String> relatedFields = rmQueryHelper.getRelatedFields(objectName);
        fields.addAll(relatedFields);
        String q = 'SELECT ' + helper.fieldsToQueryString(fields) + ' FROM ' + objectName + ' WHERE ' + whereClause;
        if (withSecurity) {
            q = q + ' WITH SECURITY_ENFORCED';
        }
        debugger.debug('restService.query(query)', q);
        return Database.query(q);
    }

    private static List<SObject> query(Set<String> fields, String objectName, String whereClause) {
        return query(fields, objectName, whereClause, true);
    }

    class PlatformPayload {
        public Set<String> order_ids = null;
        public Set<String> product2_ids = null;
        public Set<String> order_fields = null;
        public Set<String> product2_fields = null;
        public Map<String, Map<String, String>> field_mapping_data = null;
    }
}
