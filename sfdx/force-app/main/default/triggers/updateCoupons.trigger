trigger updateCoupons on Order (after update) {
  public class CouponException extends Exception {}
  
  try {
    // for all new Orders, check if the corresponding quote has coupons and duplicate/copy to the corresponding order
    for (Order order : Trigger.new) {
      if (Trigger.oldMap.get(order.Id).Status == 'Draft' && order.Status == 'Activated' && order.SBQQ__Quote__c != null) {
        // get the corresponding quote for this order
        Id quoteId = order.SBQQ__Quote__c;

        // fetch any Quote_Stripe_Coupon_Association__c for this order
        List<Quote_Stripe_Coupon_Association__c> stripeCouponQuoteAssociations = [
          SELECT Quote_Stripe_Coupon__c
          FROM Quote_Stripe_Coupon_Association__c
          WHERE Quote__c = :quoteId
        ];

        // for each Quote_Stripe_Coupon_Association__c, create a corresponding Order_Stripe_Coupon_Association__c
        if (!stripeCouponQuoteAssociations.isEmpty())
        {
          for (Quote_Stripe_Coupon_Association__c stripeCouponQuoteAssociation: stripeCouponQuoteAssociations)
          {
            List<Quote_Stripe_Coupon__c> quoteCoupons = [
              SELECT Id, Amount_Off__c, Duration__c, Duration_In_Months__c, Max_Redemptions__c, Name__c, Percent_Off__c
              FROM Quote_Stripe_Coupon__c
              WHERE Id = :stripeCouponQuoteAssociation.Quote_Stripe_Coupon__c
            ];
            
            if (quoteCoupons.isEmpty())
            {
              throw new CouponException('no stripeCoupon found for stripeCouponQuoteAssociation: ' + stripeCouponQuoteAssociation.Id);
            }
            
            Quote_Stripe_Coupon__c quoteCoupon = quoteCoupons.get(0);
            Order_Stripe_Coupon__c clonedCoupon = Utilities.cloneStripeCoupon(quoteCoupon);
            
            // create a Stripe Coupon Order Association junction object
            Order_Stripe_Coupon_Association__c orderStripeCouponAssociation = new Order_Stripe_Coupon_Association__c();
            orderStripeCouponAssociation.Order_Stripe_Coupon__c = clonedCoupon.Id;
            orderStripeCouponAssociation.Order__c = order.Id;

            // insert this record
            Database.insert((sObject)orderStripeCouponAssociation);
          }
        }
      
        // fetch the related quote lines
        List<SBQQ__QuoteLine__c> quoteLines = [
          SELECT Id
          FROM SBQQ__QuoteLine__c
          WHERE SBQQ__Quote__c = :quoteId
        ];

        for (SBQQ__QuoteLine__c quoteLine : quoteLines) {   
          // fetch the Stripe Coupon Quote Line Associations for this quote line
          List<Quote_Line_Stripe_Coupon_Association__c> stripeCouponQuoteLineAssociations = [
            SELECT Id, Quote_Stripe_Coupon__c
            FROM Quote_Line_Stripe_Coupon_Association__c
            WHERE Quote_Line__c = :quoteLine.Id
          ];

          // for each Quote_Line_Stripe_Coupon_Association__c, create a corresponding Order_Item_Stripe_Coupon_Association__c
          if (!stripeCouponQuoteLineAssociations.isEmpty())
          {
            for (Quote_Line_Stripe_Coupon_Association__c stripeCouponQuoteLineAssociation: stripeCouponQuoteLineAssociations)
            {
              List<Quote_Stripe_Coupon__c> quoteLineCoupons = [
                SELECT Amount_Off__c, Duration__c, Duration_In_Months__c, Max_Redemptions__c, Name__c, Percent_Off__c
                FROM Quote_Stripe_Coupon__c
                WHERE Id = :stripeCouponQuoteLineAssociation.Quote_Stripe_Coupon__c
              ];
              
              if (quoteLineCoupons.isEmpty())
              {
                throw new CouponException('no stripeCoupon found for stripeCouponQuoteLineAssociation: ' + stripeCouponQuoteLineAssociation.Id);
              }
              
              Quote_Stripe_Coupon__c quoteLineCoupon = quoteLineCoupons.get(0);
              Order_Stripe_Coupon__c clonedCoupon = Utilities.cloneStripeCoupon(quoteLineCoupon);
            
              // get the corresponding order line for this quote line
              List<OrderItem> orderItem = [
                SELECT Id
                FROM OrderItem
                WHERE SBQQ__QuoteLine__c = :quoteLine.Id
              ];
         
              Id orderItemId =  orderItem.get(0).Id;

              // create the corresponding Order_Item_Stripe_Coupon_Association__c
              Order_Item_Stripe_Coupon_Association__c orderLineStripeCouponAssociation = new Order_Item_Stripe_Coupon_Association__c();
              orderLineStripeCouponAssociation.Order_Stripe_Coupon__c = clonedCoupon.Id;
              orderLineStripeCouponAssociation.Order_Item__c = orderItemId;

              // insert this record
              Database.insert((sObject)orderLineStripeCouponAssociation);
            }
          }
        }
      }
    }
  } catch (Exception e) {
    errorLogger.create('updateOrderCouponsTrigger', e);
  }
}