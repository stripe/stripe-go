trigger updateOrderLineCoupons on SBQQ__Quote__c (after update) {
  try {
    // for all newly ordered quotes, check if the quote lines have coupon and duplicate to the corresponbding order lines
    for(SBQQ__Quote__c quote : Trigger.new) {
      if (quote.SBQQ__Ordered__c == true && Trigger.oldMap.get(quote.Id).SBQQ__Ordered__c == false) {
        // fetch the related quote lines
        List<SBQQ__QuoteLine__c> quoteLines = [
          SELECT Id
          FROM SBQQ__QuoteLine__c
          WHERE SBQQ__Quote__c = :quote.Id
        ];

        for (SBQQ__QuoteLine__c quoteLine : quoteLines) {   
          // get the corresponding order line for this quote line
          List<OrderItem> orderItem = [
            SELECT Id
            FROM OrderItem
            WHERE SBQQ__QuoteLine__c = :quoteLine.Id
          ];
         
          if (orderItem == null || orderItem.size() == 0)
          {
            continue;
          }
          Id orderItemId = orderItem.get(0).Id;

          // fetch the Stripe Coupon Quote Line Associations for this quote line
          List<sObject> stripeCouponQuoteLineAssociations = Utilities.fetchStripeCouponsFromAssociations('Stripe_Coupon_Quote_Line_Association__c', 'Quote_Line__c', quoteLine.Id);
          if (stripeCouponQuoteLineAssociations == null)
          {
            continue;
          }

          // for each Stripe Coupon Quote Line Association
          for (sObject stripeCouponQuoteLineAssociation: stripeCouponQuoteLineAssociations)
          {
            String prefixedStripeCoupon = constants.NAMESPACE_API + 'Stripe_Coupon__c';
            String couponId = (String)stripeCouponQuoteLineAssociation.get(prefixedStripeCoupon);
            sObject quoteLineCoupon = Utilities.fetchStripeCoupon(couponId);
            
            // clone the Stripe Coupon on the quote line, it will have a different Id
            sObject clonedCoupon = Utilities.cloneStripeCoupon(quoteLineCoupon);
          
            // insert the cloned Stripe coupon
            Database.insertImmediate(clonedCoupon);
            
            // create a Stripe Coupon Order Item Association
            String prefixedOrderItemAssociationObjName = Utilities.getObjectName('Stripe_Coupon_Order_Item_Association__c');
            sObject orderItemStripeCouponAssociation = Schema.getGlobalDescribe().get(prefixedOrderItemAssociationObjName).newSObject();
            orderItemStripeCouponAssociation.put((constants.NAMESPACE_API + 'Stripe_Coupon__c'), clonedCoupon.Id);
            orderItemStripeCouponAssociation.put((constants.NAMESPACE_API + 'Order_Item__c'), orderItemId);

            // insert this record
            Database.insertImmediate(orderItemStripeCouponAssociation);
          }
        }
      }
    }
  } catch (Exception e) {
      errorLogger.create('updateOrderLineCouponsTrigger', e);
  }
}