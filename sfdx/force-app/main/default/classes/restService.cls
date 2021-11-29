@RestResource(urlMapping='/batchApi')
global with sharing class restService {
    public class NotOrdersFoundException extends Exception {}

    @HttpGet
    global static void restServiceEvents() {
        RestRequest restRequest = RestContext.request;
        RestResponse response = RestContext.response;
        try { 
            Map<String, Object> parsedReqBody = (Map<String, Object>)JSON.deserializeUntyped(restRequest.requestBody.toString());
            List<Object> orderIdsResponse = (List<Object>)parsedReqBody.get('order_ids');                    
            List<String> orderIds = new List<String> ();
            for(Object orderId: orderIdsResponse){
                orderIds.add(String.escapeSingleQuotes(String.valueOf(orderId)));
            }

            List<Object> product2IdsResponse = (List<Object>)parsedReqBody.get('product2_ids'); 
            List<String> product2Ids = new List<String> ();
            for(Object productId: product2IdsResponse){
                product2Ids.add(String.escapeSingleQuotes(String.valueOf(productId)));
            }

            List<Object> orderFieldsResponse = (List<Object>)parsedReqBody.get('order_fields');
            List<String> orderFields = new List<String> ();
            for(Object orderField: orderFieldsResponse){
                orderFields.add(String.escapeSingleQuotes(String.valueOf(orderField)));
            }

            List<Object> product2FieldsResponse = (List<Object>)parsedReqBody.get('product2_fields');
            List<String> product2Fields = new List<String> ();
            for(Object product2Field: product2FieldsResponse){
                product2Fields.add(String.escapeSingleQuotes(String.valueOf(product2Field)));
            }

            String orderQuery = 'SELECT Id,'+ String.join((orderFields), ',')+', (SELECT Id,'+ String.join((product2Fields), ',')+' FROM OrderItems WHERE Id in :product2Ids) FROM Order WHERE Id in :orderIds  WITH SECURITY_ENFORCED';
            List<Order> orderData = Database.query(orderQuery);

            //Map<Object, Object> responseOrderRecrodsMap = new Map<Object, Object>();
            List<Object> responseOrderRecrods = new List<Object>();

            if(!orderData.isEmpty()) {
                for(SObject orderRec : orderData){
                    //List<Object> responseOrderRecrods = new List<Object>();
                    if(orderRec.getSObjects('OrderItems') != null) {
                        List<SObject> orderItemsData = orderRec.getSObjects('OrderItems');
                        for(SObject orderItemRec : orderItemsData){
                            //responseOrderRecrods.add(orderItemRec);
                            responseOrderRecrods.add((Object)orderRec + '' + orderItemRec);
                        }
                        //responseOrderRecrodsMap.add((Object)orderRec + ',' + responseOrderRecrods);
                    } else {
                        responseOrderRecrods.add((Object)orderRec + 'There were no order items found based on product2Ids sent. Here are the product2Ids that were sent: ' + product2Ids); 
                    }
                }                          
            } else {
                throw new NotOrdersFoundException('There were no orders found based on the list of order Ids sent. Here are the order Ids that were sent: ' + orderIds);
            }

            response.addHeader('Content-Type', 'application/json');
            response.statusCode = 200;
            Map<String, Object> resBody = new Map<String, Object>();
            Map<String, List<Object>> orderDetails = new Map<String,List<Object>>();
            orderDetails.put('ORDER_UID',(List<Object>)orderData);

            resBody.put('orders',(Object)responseOrderRecrods);
            response.responseBody = Blob.valueOf(JSON.serialize(resBody));

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
}
