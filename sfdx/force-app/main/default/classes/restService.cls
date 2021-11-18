@RestResource(urlMapping='/batchApi')
global with sharing class restService {
    public class SetupIsNotCompleteException extends Exception {}
    public class NotConnectedToStripeException extends Exception {}
    public class NotOrdersFoundException extends Exception {}
    public class RequestAuthenticationFailedException extends Exception {}
    @HttpPost
    global static void restServiceEvents() {
        RestRequest restRequest = RestContext.request;
        RestResponse response = RestContext.response;
        try { 
            Map<String, Object> parsedReqBody = (Map<String, Object>)JSON.deserializeUntyped(restRequest.requestBody.toString());
            System.debug('parsedReqBody');
            System.debug(parsedReqBody);
            List<Object> orderIdsResponse = (List<Object>)parsedReqBody.get('order_ids');                    
            List<String> orderIds = new List<String> ();
            for(Object orderId: orderIdsResponse){
                orderIds.add(String.valueOf(orderId));
            }

            List<Object> product2IdsResponse = (List<Object>)parsedReqBody.get('product2_ids'); 
            List<String> product2Ids = new List<String> ();
            for(Object productId: product2IdsResponse){
                product2Ids.add(String.valueOf(productId));
            }

            List<Object> orderFields = (List<Object>)parsedReqBody.get('order_fields');
            List<Object> product2Fields = (List<Object>)parsedReqBody.get('product2_fields');  

            String orderQuery = 'SELECT Id,'+ String.join((orderFields), ',')+', (SELECT Id,'+ String.join((product2Fields), ',')+' FROM OrderItems WHERE Id in :product2Ids) FROM Order WHERE Id in :orderIds  WITH SECURITY_ENFORCED';
            List<Order> orderData = Database.query(orderQuery);

            //List<Object> responseOrderRecrods = new List<Object>();
            Map<Object, Object> responseOrderRecrodsMap = new Map<Object, Object>();

            if(!orderData.isEmpty()) {
                for(SObject orderRec : orderData){
                    List<Object> responseOrderRecrods = new List<Object>();
                    if(orderRec.getSObjects('OrderItems') != null) {
                        List<SObject> orderItemsData = orderRec.getSObjects('OrderItems');
                        if(!orderItemsData.isEmpty()) {
                            for(SObject orderItemRec : orderItemsData){
                                responseOrderRecrods.add(orderItemRec);
                            }
                            responseOrderRecrodsMap.put((Object)orderRec, responseOrderRecrods);
                        } else {
                            //what should we do if there are no related orderITems Spencer
                        }
                    } else {
                        responseOrderRecrodsMap.put((Object)orderRec, 'no order items found based on product2Ids passed in: ' + product2Ids); 
                    }
                }                          
            } else {
                throw new NotOrdersFoundException('There was no orders found based on the list of order Ids sent. value found. Order Ids sent: ' + JSON.serialize(orderIds));
            }

            response.addHeader('Content-Type', 'application/json');
            response.statusCode = 200;
            Map<String, Object> resBody = new Map<String, Object>();
            Map<String, List<Object>> orderDetails = new Map<String,List<Object>>();
            orderDetails.put('ORDER_UID',(List<Object>)orderData);

            resBody.put('orders',(Object)responseOrderRecrodsMap);
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
