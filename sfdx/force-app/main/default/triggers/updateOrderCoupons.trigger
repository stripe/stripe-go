trigger updateOrderCoupons on SBQQ__Quote__c (after update) {
  try {
    // for all newly ordered quotes, check if the quote has coupon and copy/duplicate to the corresponding order
    for(SBQQ__Quote__c quote : Trigger.new) {
      if (quote.SBQQ__Ordered__c == true && Trigger.oldMap.get(quote.Id).SBQQ__Ordered__c == false) {
        // get the corresponding order for this quote 
        List<Order> order = [
          SELECT Id
          FROM Order
          WHERE SBQQ__Quote__c = :quote.Id
        ];
        
        if (order == null || order.size() == 0)
        {
          continue;
        }

        Id orderId =  order.get(0).Id;

        // fetch the Stripe Coupon Quote Associations related to this quote
        List<sObject> stripeCouponQuoteAssociations = Utilities.fetchStripeCouponsFromAssociations('Stripe_Coupon_Quote_Association__c', 'Quote__c', quote.Id);
        if (stripeCouponQuoteAssociations == null)
        {
          continue;
        }

        // for each Stripe Coupon Quote Association
        for(sObject stripeCouponQuoteAssociation: stripeCouponQuoteAssociations)
        {
          String prefixedStripeCoupon = constants.NAMESPACE_API + 'Stripe_Coupon__c';
          String couponId = (String)stripeCouponQuoteAssociation.get(prefixedStripeCoupon);
          sObject quoteCoupon = Utilities.fetchStripeCoupon(couponId);
     
          // clone the Stripe Coupon on the quote, it will have a different Id
          sObject clonedCoupon = Utilities.cloneStripeCoupon(quoteCoupon);
          
          // insert the cloned Stripe coupon
          Database.insertImmediate(clonedCoupon);
          
          // create a Stripe Coupon Order Association junction object
          String prefixedOrderAssociationObjName = Utilities.getObjectName('Stripe_Coupon_Order_Association__c');
          sObject orderStripeCouponAssociation = Schema.getGlobalDescribe().get(prefixedOrderAssociationObjName).newSObject();
          orderStripeCouponAssociation.put((constants.NAMESPACE_API + 'Stripe_Coupon__c'), clonedCoupon.Id);
          orderStripeCouponAssociation.put((constants.NAMESPACE_API + 'Order__c'), orderId);

          // insert this record
          Database.insertImmediate(orderStripeCouponAssociation);
        }
      }
    }
  } catch (Exception e) {
      errorLogger.create('updateOrderCouponsTrigger', e);
  }
}