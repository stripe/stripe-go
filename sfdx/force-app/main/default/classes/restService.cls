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

            String orderQuery;
            if(!product2Ids.isEmpty()){
                orderQuery = 'SELECT '+ String.join(((Iterable<String>)orderFields), ',')+', (SELECT '+ String.join(((Iterable<String>)product2Fields), ',')+' FROM OrderItems WHERE Id in :product2Ids) FROM Order WHERE Id in :orderIds  WITH SECURITY_ENFORCED';
            } else {
                orderQuery = 'SELECT '+ String.join(((Iterable<String>)orderFields), ',')+', (SELECT '+ String.join(((Iterable<String>)product2Fields), ',')+' FROM OrderItems) FROM Order WHERE Id in :orderIds  WITH SECURITY_ENFORCED';
            }
            List<Order> orderData = Database.query(orderQuery);
            List<Object> responseOrderRecrods = new List<Object>();
            if(!orderData.isEmpty()) {
                for(SObject orderRec : orderData) {
                    List<Object> attibutes = new List<Object>();
                    String jsonQueryValue = JSON.serialize(orderRec);

                    if(orderRec.get('ContractId') != null) {
                        String contractId = string.valueOf(orderRec.get('ContractId'));
                        Set<String> contractFields = getAccessableFieldsByObject('Contract');
                        String contractQuery = 'SELECT '+ String.join(((Iterable<String>)contractFields), ',')+' FROM Contract WHERE Id = :contractId  WITH SECURITY_ENFORCED';
                        List<Contract> contractData = Database.query(contractQuery);
                        if(!contractData.isEmpty()) {
                            jsonQueryValue = jsonQueryValue+JSON.serialize(contractData);
                        }
                        String contractOrderQuery = 'SELECT '+ String.join(((Iterable<String>)orderFields), ',')+' ROM Order WHERE ContractId = :contractId  WITH SECURITY_ENFORCED';
                        List<Order> contractOrderData = Database.query(orderQuery);
                        if(!contractOrderData.isEmpty()) {
                            jsonQueryValue = jsonQueryValue+JSON.serialize(contractOrderData);
                        }
                    }
                    if(orderRec.get('AccountId') != null) {
                        String accountId = string.valueOf(orderRec.get('AccountId'));
                        Set<String> accountFields = getAccessableFieldsByObject('Account');
                        String accountQuery = 'SELECT '+ String.join(((Iterable<String>)accountFields), ',')+' FROM Account WHERE Id = :accountId  WITH SECURITY_ENFORCED';
                        List<Account> accountData = Database.query(accountQuery);
                        if(!accountData.isEmpty()) {
                            jsonQueryValue = jsonQueryValue+JSON.serialize(accountData);
                        }
                    } 
                    if(orderRec.get('SBQQ__Quote__c') != null) {
                        String quoteId = string.valueOf(orderRec.get('SBQQ__Quote__c'));
                        Set<String> quoteFields = getAccessableFieldsByObject('SBQQ__Quote__c');
                        String quoteQuery = 'SELECT '+ String.join(((Iterable<String>)quoteFields), ',')+' FROM SBQQ__Quote__c WHERE Id = :quoteId  WITH SECURITY_ENFORCED';
                        List<SBQQ__Quote__c> quoteData = Database.query(quoteQuery);
                        if(!quoteData.isEmpty()) {
                            jsonQueryValue = jsonQueryValue+JSON.serialize(quoteData);
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
            errorLogger.create('Rest Service DmlException', e.getMessage(), String.valueOf(e.getLineNumber()), restRequest.requestBody.toString());
            allRestErrors.add('DmlException: ' + e.getMessage() + 'Line Number: '+string.valueOf(e.getLineNumber()));  
        } catch(SObjectException e) {
            errorLogger.create('Rest Service SObjectException', e.getMessage(), String.valueOf(e.getLineNumber()), restRequest.requestBody.toString());
            allRestErrors.add('SObjectException: ' + e.getMessage() + 'Line Number: '+string.valueOf(e.getLineNumber()));    
        } catch(Exception e) {
            errorLogger.create('Rest Service Event', e.getMessage(), String.valueOf(e.getLineNumber()), restRequest.requestBody.toString());
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
