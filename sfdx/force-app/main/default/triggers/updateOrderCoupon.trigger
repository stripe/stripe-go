trigger updateOrderCoupon on Order (after insert, after update) {
  public class CouponException extends Exception {}
  
  try {
    // for all new Orders, check if the corresponding quote has coupons and duplicate/copy to the corresponding order
    for (Order order : Trigger.new) {
      if (order.SBQQ__Quote__c != null && order.Status == 'Activated' && (Trigger.oldMap.get(order.Id) == null || Trigger.oldMap.get(order.Id).Status != 'Activated')) {
        // get the corresponding quote for this order
        List<Order> orders = [
          SELECT Id, SBQQ__Quote__c
          FROM Order
          WHERE Id = :order.Id
        ];
 
        if (orders.isEmpty()) {
          continue;
        }

        Id quoteId = orders.get(0).SBQQ__Quote__c;
        if (quoteId == null) {
          throw new CouponException('Order does not contain SBQQ__Quote__c field'); 
        }

        // fetch the Stripe Coupon Quote Associations for this quote
        List<Stripe_Coupon_Quote_Association__c> stripeCouponQuoteAssociations = [
          SELECT Id, Stripe_Coupon__c
          FROM Stripe_Coupon_Quote_Association__c
          WHERE Quote__c = :quoteId
        ];

        // Check if this Quote has associated coupons
        if (stripeCouponQuoteAssociations.isEmpty())
        {
          continue;
        }

        // for each Stripe Coupon Quote Association
        for (Stripe_Coupon_Quote_Association__c stripeCouponQuoteAssociation: stripeCouponQuoteAssociations)
        {
          List<Stripe_Coupon__c> quoteCoupons = [
            SELECT Id, Amount_Off__c, Duration__c, Duration_In_Months__c, Max_Redemptions__c, Name__c, Percent_Off__c
            FROM Stripe_Coupon__c
            WHERE Id = :stripeCouponQuoteAssociation.Stripe_Coupon__c
          ];
          
          if (quoteCoupons.isEmpty())
          {
            throw new CouponException('no stripeCoupon found for stripeCouponQuoteAssociation: ' + stripeCouponQuoteAssociation);
          }
          
          Stripe_Coupon__c quoteCoupon = quoteCoupons.get(0);

          // clone the Stripe Coupon on the quote, it will have a different Id
          Stripe_Coupon_Serialized__c clonedCoupon = new Stripe_Coupon_Serialized__c();
          clonedCoupon.Amount_Off__c = quoteCoupon.Amount_Off__c;
          clonedCoupon.Duration__c = quoteCoupon.Duration__c;
          clonedCoupon.Duration_In_Months__c = quoteCoupon.Duration_In_Months__c;
          clonedCoupon.Max_Redemptions__c = quoteCoupon.Max_Redemptions__c;
          clonedCoupon.Name__c = quoteCoupon.Name__c;
          clonedCoupon.Percent_Off__c = quoteCoupon.Percent_Off__c;
          clonedCoupon.Original_Stripe_Coupon_Id__c = quoteCoupon.Id;
     
          // insert the cloned Stripe coupon
          Database.SaveResult result = Database.insertImmediate((sObject)clonedCoupon);
          if (!result.isSuccess()) { 
              // operation failed, so get all errors                
              for(Database.Error err : result.getErrors()) {
                  System.debug('The following error has occurred.');                    
                  System.debug(err.getStatusCode() + ': ' + err.getMessage());
                  System.debug('Fields that affected this error: ' + err.getFields());
              }
            }
          
          // create a Stripe Coupon Order Association junction object
          Stripe_Coupon_Order_Association__c orderStripeCouponAssociation = new Stripe_Coupon_Order_Association__c(
            Stripe_Coupon__c = clonedCoupon.Id,
            Order__c = order.Id
          );

          // insert this record
          Database.insertImmediate((sObject)orderStripeCouponAssociation);
        }
      }
    }
  } catch (Exception e) {
    errorLogger.create('updateOrderCouponsTrigger', e);
  }
}