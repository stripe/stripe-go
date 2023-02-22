public with sharing class OrderCouponTriggerHandler {
    public class CouponException extends Sentry_Exception {}

    private Map<Id, Id> quoteIdsToOrderIds = null;
    private Map<Id, SBQQ__QuoteLine__c> quoteLines = null;
    // the order coupons that are bulk persisted at the end of the trigger execution
    private List<Order_Stripe_Coupon__c> orderCouponsToSave = null;

    private static final String FIELD_CURRENCY_ISO_CODE = constants.FIELD_CURRENCY_ISO_CODE;

    public OrderCouponTriggerHandler(Map<Id, Id> quoteIdsToOrderIds) {
        this.refreshState(quoteIdsToOrderIds);
    }

    public void refreshState(Map<Id, Id> quoteIdsToOrderIds) {
        if (quoteIdsToOrderIds == null) {
            this.quoteIdsToOrderIds = new Map<Id, Id>();
        } else {
            this.quoteIdsToOrderIds = quoteIdsToOrderIds;
        }

        orderCouponsToSave = new List<Order_Stripe_Coupon__c>();
        quoteLines = new Map<Id, SBQQ__QuoteLine__c>();
    }

    public void process() {
        removePreviouslyProcessedOrders();
        processQuoteCoupons();
        processQuoteLineCoupons();
    }

    // in case the user has workflows that would cause the trigger to fire twice, let's ensure we don't create duplicate Order Stripe Coupons
    public void removePreviouslyProcessedOrders() {
        // create the reverse map of order id -> quote id
        Map<Id, Id> orderIdsToQuoteIds = new Map<Id, Id>();
        for (Id quoteId : quoteIdsToOrderIds.keySet()) {
            orderIdsToQuoteIds.put(quoteIdsToOrderIds.get(quoteId), quoteId);
        }

        Map<Id, OrderItem> orderItemsById = new Map<Id, OrderItem>([SELECT Id, OrderId FROM OrderItem WHERE OrderId = :quoteIdsToOrderIds.values()]);

        // query for all Order Stripe Coupons that are tied to these orders or order items
        List<Order_Stripe_Coupon__c> coupons = [
            SELECT
                Order__c,
                Order_Item__c
            FROM Order_Stripe_Coupon__c
            WHERE
                Order__c = :quoteIdsToOrderIds.values()
                OR Order_Item__c = :orderItemsById.keySet()
        ];

        for (Order_Stripe_Coupon__c coupon : coupons) {
            Id orderId = null;
            if (coupon.Order_Item__c != null) {
                orderId = orderItemsById.get(coupon.Order_Item__c).OrderId;
            } else if (coupon.Order__c != null) {
                orderId = coupon.Order__c;
            }

            // this should really never happen since an Order Stripe Coupon 
            // should always be created with either Order__c or Order_Item__c fields set
            if (orderId == null) {
                continue;
            }

            // remove any quote we have already processed
            Id quoteId = orderIdsToQuoteIds.get(orderId);
            if (quoteIdsToOrderIds.containsKey(quoteId)) {
                quoteIdsToOrderIds.remove(quoteId);
            }
        }
    }

    public void processQuoteCoupons() {
        // SECTION: fetch the quote->order coupon data
        List<Quote_Stripe_Coupon_Association__c> qscAssociations = getQuoteCouponAssociations(quoteIdsToOrderIds.keySet());

        if (qscAssociations.isEmpty()) {
            return;
        }

        for (Quote_Stripe_Coupon_Association__c assoc : qscAssociations) {
            Quote_Stripe_Coupon__c quoteCoupon = assoc.Quote_Stripe_Coupon__r;

            // if the coupon is null, log an error and move on, as other orders may be fine.
            if (quoteCoupon == null || quoteCoupon.Id == null) {
                CouponException ex = (CouponException) Sentry_ExceptionFactory.build(CouponException.class);
                ex.setMessage('No Quote_Stripe_Coupon__c found for Quote_Stripe_Coupon_Association__c');
                ex.context.put('MapOfQuoteIdsToOrderIds', quoteIdsToOrderIds);
                ex.context.put('Quote_Stripe_Coupon_Association__c.Id', assoc.Id);
                ex.context.put('SBQQ__Quote__c.Id', assoc.Quote__c);
                ex.context.put('Order.Id', quoteIdsToOrderIds.get(assoc.Quote__c));
                errorLogger.create('OrderCouponTriggerHandler-CouponNotFound', ex);
                continue;
            }

            // clone the quote coupon
            Id orderId = quoteIdsToOrderIds.get(assoc.Quote__c);
            Order_Stripe_Coupon__c orderCoupon = createOrderCoupon(quoteCoupon);
            orderCoupon.Order__c = orderId;

            orderCouponsToSave.add(orderCoupon);
        }
    }

    public void processQuoteLineCoupons() {
        // SECTION: fetch the quote line -> order line coupon data
        quoteLines = getQuoteLinesByIds(quoteIdsToOrderIds.keySet());

        if (quoteLines.isEmpty()) {
            return;
        }

        Set<Id> quoteLineIds = quoteLines.keySet();
        List<Quote_Line_Stripe_Coupon_Association__c> qlscAssociations = getQuoteLineCouponAssociations(quoteLineIds);
        
        if (qlscAssociations.isEmpty()) {
            return;
        }

        Map<Id, Id> orderItemByQuoteLineId = getOrderItemIdByQuoteLineIds(quoteLineIds);

        for (Quote_Line_Stripe_Coupon_Association__c assoc : qlscAssociations) {
            Quote_Stripe_Coupon__c quoteCoupon = assoc.Quote_Stripe_Coupon__r;

            // if the coupon is null, log an error and move on, as other order lines may be fine.
            if (quoteCoupon == null || quoteCoupon.Id == null) {
                CouponException ex = (CouponException) Sentry_ExceptionFactory.build(CouponException.class);
                ex.setMessage('No Quote_Stripe_Coupon__c found for Quote_Line_Stripe_Coupon_Association__c');
                ex.context.put('MapOfQuoteIdsToOrderIds', quoteIdsToOrderIds);
                ex.context.put('Quote_Line_Stripe_Coupon_Association__c.Id', assoc.Id);
                ex.context.put('SBQQ__QuoteLine__c', assoc.Quote_Line__c);
                SBQQ__QuoteLine__c ql = quoteLines.get(assoc.Quote_Line__c);
                ex.context.put('SBQQ__Quote__c', (ql != null) ? ql.SBQQ__Quote__c : null);
                ex.context.put('Order.Id', (ql != null) ? quoteIdsToOrderIds.get(ql.Id) : null);
                errorLogger.create('OrderCouponTriggerHandler-CouponNotFound', ex);
                continue;
            }

            // clone the quote coupon
            Order_Stripe_Coupon__c orderCoupon = createOrderCoupon(quoteCoupon);
            orderCoupon.Order_Item__c = orderItemByQuoteLineId.get(assoc.Quote_Line__c);

            orderCouponsToSave.add(orderCoupon);
        }
    }

    public void persistChanges() {
        // SECTION: save the order and order item coupons
        Database.insert(orderCouponsToSave);
    }

    public static Order_Stripe_Coupon__c createOrderCoupon(Quote_Stripe_Coupon__c quoteCoupon) {
        // clone the Quote Coupon on the quote, it will have a different Id
        Order_Stripe_Coupon__c osc = new Order_Stripe_Coupon__c(
            Amount_Off__c = quoteCoupon.Amount_Off__c,
            Duration__c = quoteCoupon.Duration__c,
            Duration_In_Months__c = quoteCoupon.Duration_In_Months__c,
            Max_Redemptions__c = quoteCoupon.Max_Redemptions__c,
            Name__c = quoteCoupon.Name__c,
            Percent_Off__c = quoteCoupon.Percent_Off__c,
            Quote_Stripe_Coupon_Id__c = quoteCoupon.Id
        );

        if (UserInfo.isMultiCurrencyOrganization()) {
            osc.put(FIELD_CURRENCY_ISO_CODE, quoteCoupon.get(FIELD_CURRENCY_ISO_CODE));
        }

        return osc;
    }

    private List<Quote_Stripe_Coupon_Association__c> getQuoteCouponAssociations(Set<Id> quoteIds) {
        List<String> fields = new List<String> {
            'Quote__c',
            'Quote_Stripe_Coupon__c',
            'Quote_Stripe_Coupon__r.Id',
            'Quote_Stripe_Coupon__r.Amount_Off__c',
            'Quote_Stripe_Coupon__r.Duration__c',
            'Quote_Stripe_Coupon__r.Duration_In_Months__c',
            'Quote_Stripe_Coupon__r.Max_Redemptions__c',
            'Quote_Stripe_Coupon__r.Name__c',
            'Quote_Stripe_Coupon__r.Percent_Off__c'
        };

        List<String> multiCurrFields = new List<String> {
            'Quote_Stripe_Coupon__r.' + FIELD_CURRENCY_ISO_CODE,
            FIELD_CURRENCY_ISO_CODE
        };

        return utilities.execMultiCurrLookupQuery(
                Quote_Stripe_Coupon_Association__c.SObjectType,
                'Quote__c',
                quoteIds,
                fields,
                multiCurrFields
        );
    }

    private Map<Id, SBQQ__QuoteLine__c> getQuoteLinesByIds(Set<Id> quoteIds) {
        return new Map<Id, SBQQ__QuoteLine__c>([
            SELECT
                Id,
                SBQQ__Quote__c,
                (
                    SELECT
                        Id,
                        SBQQ__QuoteLine__c,
                        OrderId
                    FROM SBQQ__OrderProducts__r
                )
            FROM SBQQ__QuoteLine__c
            WHERE SBQQ__Quote__c = :quoteIds
        ]);
    }

    private List<Quote_Line_Stripe_Coupon_Association__c> getQuoteLineCouponAssociations(Set<Id> quoteLineIds) {
        List<String> fields = new List<String> {
            'Id',
            'Quote_Line__c',
            'Quote_Line__r.Id',
            'Quote_Line__r.SBQQ__Quote__c',
            'Quote_Stripe_Coupon__c',
            'Quote_Stripe_Coupon__r.Id',
            'Quote_Stripe_Coupon__r.Amount_Off__c',
            'Quote_Stripe_Coupon__r.Duration__c',
            'Quote_Stripe_Coupon__r.Duration_In_Months__c',
            'Quote_Stripe_Coupon__r.Max_Redemptions__c',
            'Quote_Stripe_Coupon__r.Name__c',
            'Quote_Stripe_Coupon__r.Percent_Off__c'
        };

        List<String> multiCurrFields = new List<String> {
            'Quote_Stripe_Coupon__r.' + FIELD_CURRENCY_ISO_CODE,
            FIELD_CURRENCY_ISO_CODE
        };

        return utilities.execMultiCurrLookupQuery(
                Quote_Line_Stripe_Coupon_Association__c.SObjectType,
                'Quote_Line__c',
                quoteLineIds,
                fields,
                multiCurrFields
        );
    }

     private Map<Id, Id> getOrderItemIdByQuoteLineIds(Set<Id> quoteLineIds) {
        List<OrderItem> orderItems = [
            SELECT Id, SBQQ__QuoteLine__c
            FROM OrderItem
            WHERE SBQQ__QuoteLine__c IN :quoteLineIds
        ];

         Map<Id, Id> orderItemsByQuoteLineId = new Map<Id, Id>();
         for (OrderItem orderItem: orderItems)
         {
            orderItemsByQuoteLineId.put(orderItem.SBQQ__QuoteLine__c, orderItem.Id);
         }
         return orderItemsByQuoteLineId;
    }      
}