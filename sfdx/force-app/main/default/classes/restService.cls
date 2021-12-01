@RestResource(urlMapping='/batchApi')
global with sharing class restService {
    public class NotOrdersFoundException extends Exception {}
    global static List<Object> allRestErrors = new List<Object>();

    @HttpGet
    global static void restServiceEvents() {
        RestRequest restRequest = RestContext.request;
        RestResponse response = RestContext.response;
        try { 
            Map<String, Object> parsedReqBody = (Map<String, Object>)JSON.deserializeUntyped(restRequest.requestBody.toString());
            List<Object> orderIdsResponse = (List<Object>)parsedReqBody.get('order_ids'); 
            if(orderIdsResponse.isEmpty())throw new NotOrdersFoundException('There were no order Ids sent with this request.');                   
            List<String> orderIds = new List<String> ();
            for(Object orderId: orderIdsResponse) {
                orderIds.add(String.escapeSingleQuotes(String.valueOf(orderId)));
            }

            List<Object> product2IdsResponse = (List<Object>)parsedReqBody.get('product2_ids'); 
            List<String> product2Ids = new List<String> ();
            for(Object productId: product2IdsResponse) {
                product2Ids.add(String.escapeSingleQuotes(String.valueOf(productId)));
            }

            List<Object> orderFieldsResponse = (List<Object>)parsedReqBody.get('order_fields');
            Set<String> orderFields = new Set<String>();
            if(!orderFieldsResponse.isEmpty()) {
                for(Object orderField: orderFieldsResponse) {
                    orderFields.add(String.escapeSingleQuotes(String.valueOf(orderField)));
                }
            } else {
                orderFields = getAccessableFieldsByObject('Order');
            }

            List<Object> product2FieldsResponse = (List<Object>)parsedReqBody.get('product2_fields');
            Set<String> product2Fields = new Set<String> ();
            if(!product2FieldsResponse.isEmpty()) {
                for(Object product2Field: product2FieldsResponse) {
                    product2Fields.add(String.escapeSingleQuotes(String.valueOf(product2Field)));
                }
            } else {
                product2Fields = getAccessableFieldsByObject('OrderItem');
            }

            Set<String> quoteFields = getAccessableFieldsByObject('SBQQ__Quote__c');
            if(orderFields.contains('SBQQ__Quote__c')) { 
                for(String quoteField : quoteFields) {
                    orderFields.add('SBQQ__Quote__r.' + quoteField);
                }
            }

            Set<String> accountFields = getAccessableFieldsByObject('Account');
            if(orderFields.contains('AccountId')) { 
                for(String accountField : accountFields) {
                    orderFields.add('Account.' + accountField);
                }
            }

            Set<String> contractFields = getAccessableFieldsByObject('Contract');
            if(orderFields.contains('ContractId')) { 
                for(String contractField : contractFields) {
                    orderFields.add('Contract.' + contractField);
                }
            }

            Set<String> productFields = getAccessableFieldsByObject('Product2');
            if(product2Fields.contains('Product2Id')) { 
                for(String productField : productFields) {
                    product2Fields.add('PricebookEntry.Product2.' + productField);
                }
                //may be able to delete below Arnold 
                Set<String> priceBookFields = getAccessableFieldsByObject('PricebookEntry');
                for(String priceBookField : priceBookFields) {
                    product2Fields.add('PricebookEntry.' + priceBookField);
                }
            }

            String orderQuery;
            if(!product2Ids.isEmpty()){
                orderQuery = 'SELECT '+ String.join(((Iterable<String>)orderFields), ',')+', (SELECT '+ String.join(((Iterable<String>)product2Fields), ',')+' FROM OrderItems WHERE Id in :product2Ids) FROM Order WHERE Id in :orderIds  WITH SECURITY_ENFORCED';
            } else {
                orderQuery = 'SELECT '+ String.join(((Iterable<String>)orderFields), ',')+', (SELECT '+ String.join(((Iterable<String>)product2Fields), ',')+' FROM OrderItems) FROM Order WHERE Id in :orderIds  WITH SECURITY_ENFORCED';
            }
            System.debug(orderQuery);
            List<Order> orderData = Database.query(orderQuery);
            List<Object> responseOrderRecrods = new List<Object>();
            if(!orderData.isEmpty()) {
                for(SObject orderRec : orderData) {
                    List<Object> attibutes = new List<Object>();
                    String jsonQueryValue = JSON.serialize(orderRec);

                    if(orderRec.getSObjects('OrderItems') != null) {
                        List<SObject> orderItemsData = orderRec.getSObjects('OrderItems');
                        List<Object> orderItemsIds = new List<Object>();
                        for(SObject orderItemRec : orderItemsData) {
                            orderItemsIds.add(orderItemRec.get('Id'));
                        }
                        jsonQueryValue = jsonQueryValue + JSON.serialize('OrderLines') +': ' + JSON.serialize(orderItemsIds);
                    }

                    if(orderRec.get('ContractId') != null) {
                        String contractId = string.valueOf(orderRec.get('ContractId'));
                        String contractOrderQuery = 'SELECT '+ String.join(((Iterable<String>)orderFields), ',')+' FROM Order WHERE ContractId = :contractId  WITH SECURITY_ENFORCED';
                        List<Order> contractOrderData = Database.query(orderQuery);
                        if(!contractOrderData.isEmpty()) {
                            jsonQueryValue = jsonQueryValue+JSON.serialize(contractOrderData);
                            List<Object> contractOrderIds = new List<Object>();
                            for(SObject contractOrder : contractOrderData) {
                                contractOrderIds.add(contractOrder.get('Id'));
                            }
                            jsonQueryValue = jsonQueryValue + JSON.serialize('ContractOrders') +': ' + JSON.serialize(contractOrderIds);
                        }

                        //Can delete below most likey
                      /*String contractQuery = 'SELECT '+ String.join(((Iterable<String>)contractFields), ',')+' FROM Contract WHERE Id = :contractId  WITH SECURITY_ENFORCED';
                        List<Contract> contractData = Database.query(contractQuery);
                        if(!contractData.isEmpty()) {
                            jsonQueryValue = jsonQueryValue+JSON.serialize(contractData);
                            List<Object> contractIds = new List<Object>();
                            for(SObject contract : contractData) {
                                contractIds.add(contract.get('Id'));
                            }
                            jsonQueryValue = jsonQueryValue + JSON.serialize('Contracts') +': ' + JSON.serialize(contractIds);
                        } */
                    }
                    //Can delete below most likey
                    /*  if(orderRec.get('AccountId') != null) {
                        String accountId = string.valueOf(orderRec.get('AccountId'));
                        Set<String> accountFields = getAccessableFieldsByObject('Account');
                        String accountQuery = 'SELECT '+ String.join(((Iterable<String>)accountFields), ',')+' FROM Account WHERE Id = :accountId  WITH SECURITY_ENFORCED';
                        List<Account> accountData = Database.query(accountQuery);
                        if(!accountData.isEmpty()) {
                            jsonQueryValue = jsonQueryValue+JSON.serialize(accountData);
                        }
                    }  
                     if(orderRec.get('BillToContactId') != null) {
                        String contactId = string.valueOf(orderRec.get('BillToContactId'));
                        Set<String> contactFields = getAccessableFieldsByObject('Contact');
                        String contactQuery = 'SELECT '+ String.join(((Iterable<String>)contactFields), ',')+' FROM Contact WHERE Id = :contactId  WITH SECURITY_ENFORCED';
                        List<Contact> contactData = Database.query(contactQuery);
                        if(!contactData.isEmpty()) {
                            jsonQueryValue = jsonQueryValue+JSON.serialize(contactData);
                        }
                    } 
                    responseOrderRecrods.add(jsonQueryValue);      
                }*/
                    if(orderRec.get('SBQQ__Quote__c') != null) {
                        String quoteId = string.valueOf(orderRec.get('SBQQ__Quote__c'));
                        String quoteQuery = 'SELECT '+ String.join(((Iterable<String>)quoteFields), ',')+' FROM SBQQ__Quote__c WHERE Id = :quoteId  WITH SECURITY_ENFORCED';
                        List<SBQQ__Quote__c> quoteData = Database.query(quoteQuery);
                        System.debug(quoteQuery);
                        System.debug(quoteData);
                        if(!quoteData.isEmpty()) {
                            jsonQueryValue = jsonQueryValue+JSON.serialize(quoteData);
                            List<Object> quoteIds = new List<Object>();
                            for(SObject quote : quoteData) {
                                quoteIds.add(quote.get('Id'));
                                if(quote.get('SBQQ__Account__c') != null) {
                                    String accountId = string.valueOf(quote.get('SBQQ__Account__c'));
                                    String accountQuery = 'SELECT '+ String.join(((Iterable<String>)accountFields), ',')+' FROM Account WHERE Id = :accountId  WITH SECURITY_ENFORCED';
                                    List<Account> accountData = Database.query(accountQuery);
                                    if(!accountData.isEmpty()) {
                                        jsonQueryValue = jsonQueryValue+JSON.serialize(accountData);
                                        List<Object> accountIds = new List<Object>();
                                        for(SObject account : accountData) {
                                            accountIds.add(account.get('Id'));
                                        }
                                        jsonQueryValue = jsonQueryValue + JSON.serialize('SBQQ__Quote__cAccounts') +': ' + JSON.serialize(accountIds);
                                    }
                                } 

                                if(quote.get('SBQQ__Opportunity2__c') != null) {
                                    String opportunityId = string.valueOf(quote.get('SBQQ__Opportunity2__c'));
                                    Set<String> opportunityFields = getAccessableFieldsByObject('Opportunity');
                                    String opportunityQuery = 'SELECT '+ String.join(((Iterable<String>)opportunityFields), ',')+' FROM Opportunity WHERE Id = :opportunityId  WITH SECURITY_ENFORCED';
                                    List<Opportunity> opportunityData = Database.query(opportunityQuery);
                                    if(!opportunityData.isEmpty()) {
                                        jsonQueryValue = jsonQueryValue+JSON.serialize(opportunityData);
                                        List<Object> opportunityIds = new List<Object>();
                                        for(SObject opportunity : opportunityData) {
                                            opportunityIds.add(opportunity.get('Id'));
                                        }
                                        jsonQueryValue = jsonQueryValue + JSON.serialize('SBQQ__Quote__cOpportunities') +': ' + JSON.serialize(opportunityIds);
                                    }
                                }

                                if(quote.get('SBQQ__PrimaryContact__c') != null) {
                                    String contactId = string.valueOf(quote.get('SBQQ__PrimaryContact__c'));
                                    Set<String> contactFields = getAccessableFieldsByObject('Contact');
                                    String contactQuery = 'SELECT '+ String.join(((Iterable<String>)contactFields), ',')+' FROM Contact WHERE Id = :contactId  WITH SECURITY_ENFORCED';
                                    List<Contact> contactData = Database.query(contactQuery);
                                    if(!contactData.isEmpty()) {
                                        jsonQueryValue = jsonQueryValue+JSON.serialize(contactData);
                                        List<Object> contactIds = new List<Object>();
                                        for(SObject contact : contactData) {
                                            contactIds.add(contact.get('Id'));
                                        }
                                        jsonQueryValue = jsonQueryValue + JSON.serialize('SBQQ__Quote__cContacts') +': ' + JSON.serialize(contactIds);
                                    }
                                } 
                                
                            }
                            jsonQueryValue = jsonQueryValue + JSON.serialize('Quotes') +': ' + JSON.serialize(quoteIds);
                        }
                    } 
                    responseOrderRecrods.add(jsonQueryValue);      
                }                          
            } else {
                throw new NotOrdersFoundException('There were no orders found based on the list of order Ids sent. Here are the order Ids that were sent: ' + orderIds);
            }
            response.addHeader('Content-Type', 'application/json');
            response.statusCode = 200;
            Map<String, Object> resBody = new Map<String, Object>();
            resBody.put('errors', allRestErrors);
            Map<String, List<Object>> orderDetails = new Map<String,List<Object>>();
            resBody.put('records',(Object)responseOrderRecrods);
            response.responseBody = Blob.valueOf(JSON.serialize(resBody));

        } catch(DmlException e) {
            //errorLogger.create('Rest Service DmlException', e.getMessage(), String.valueOf(e.getLineNumber()), restRequest.requestBody.toString());
            allRestErrors.add('DmlException: ' + e.getMessage() + 'Line Number: '+string.valueOf(e.getLineNumber()));  
        } catch(SObjectException e) {
            //errorLogger.create('Rest Service SObjectException', e.getMessage(), String.valueOf(e.getLineNumber()), restRequest.requestBody.toString());
            allRestErrors.add('SObjectException: ' + e.getMessage() + 'Line Number: '+string.valueOf(e.getLineNumber()));    
        } catch(Exception e) {
            //errorLogger.create('Rest Service Event', e.getMessage(), String.valueOf(e.getLineNumber()), restRequest.requestBody.toString());
            response.statusCode = 500;
            Map<String, Object> resBody = new Map<String, Object>();
            resBody.put('error', 'APEX_ERROR');
            resBody.put('message', e.getLineNumber() + ': ' + e.getMessage());
            response.responseBody = Blob.valueOf(JSON.serialize(resBody));
            return;
        }
    }

    public static Set<String> getAccessableFieldsByObject(String ObjectName) {
        Schema.DescribeSObjectResult describedObj = Schema.describeSObjects(new List<String>{ObjectName}, SObjectDescribeOptions.DEFERRED).get(0);
		Set<String> accessableFields = new Set<String>();
        for(Schema.SObjectField fieldSchema : describedObj.fields.getMap().values()) {
            Schema.DescribeFieldResult describedField = fieldSchema.getDescribe();
            if(!describedField.isAccessible()) {
                allRestErrors.add('Object:'+ObjectName+' {Field Not Accessable:'+describedField.getName()+'}');
                continue;
            }
           accessableFields.add(describedField.getName());
       }
    return accessableFields;
    }
}
