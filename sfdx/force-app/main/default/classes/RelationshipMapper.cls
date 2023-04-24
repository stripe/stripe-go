/**
 * Created by jmather-c on 3/14/23.
 */

/**
 * Handles taking in a root object, and a dot-separated path, and figuring out if it leads to a queryable field.
 */
public with sharing class RelationshipMapper {
    public class RelationshipMapperException extends Sentry_Exception {}

    public static Map<SObjectType, Map<String, DescribeFieldResult>> lookupFields = new Map<SObjectType, Map<String, DescribeFieldResult>>();

    public static Map<String, SObjectType> SupportedTypes = new Map<String, SObjectType> {
            'price' => PricebookEntry.getSObjectType(),
            'product' => Product2.getSObjectType(),
            'customer' => Account.getSObjectType(),
            'subscription' => Order.getSObjectType(),
            'price_order_item' => OrderItem.getSObjectType(),
            'subscription_item' => OrderItem.getSObjectType(),
            'subscription_schedule' => Order.getSObjectType()
    };

    public List<LookupChain> getValidChains(List<LookupChain> chains) {
        List<LookupChain> result = new List<LookupChain>();

        for (LookupChain chain : chains) {
            if (chain.isValid) {
                result.add(chain);
            }
        }

        return result;
    }

    public List<LookupChain> getValidChains(SObjectType type, String path) {
        return getValidChains(getLookupChains(type, path));
    }

    public List<LookupChain> getLookupChains(SObjectType type, String path) {
        List<LookupChain> chains = new List<LookupChain>();
        List<LookupChainItem> items = getLookupChainItems(type, path);

        if (items.isEmpty()) {
            return chains;
        }

        List<LookupChainItem> rawItems = items.clone();
        LookupChainItem firstItem = items.get(0);
        LookupChain firstChain = new LookupChain(firstItem);
        chains.add(firstChain);
        processTargets(chains, firstChain, firstItem, true);
        items.remove(0);

        while (items.isEmpty() == false ) {
            Boolean foundChangesThisCycle = false;

            for (LookupChain lc : chains) {
                LookupChainItem lastItem = lc.last();

                if (lastItem.hasNextTarget() == false) {
                    continue;
                }

                SObjectType nextTarget = lastItem.nextTarget();
                LookupChainItem nextItem = null;
                Integer nextItemIndex = 0;

                for (LookupChainItem item : items) {
                    if (item.root == nextTarget) {
                        nextItem = item;
                        break;
                    }

                    nextItemIndex++;
                }

                if (nextItem == null) {
                    continue;
                }

                foundChangesThisCycle = true;

                processTargets(chains, lc, nextItem, false);

                items.remove(nextItemIndex);
            }

            if (foundChangesThisCycle == false) {
                break;
            }
        }

        return chains;
    }

    public void processTargets(List<LookupChain> chains, LookupChain lc, LookupChainItem item, Boolean processingRoot) {
        if (item.hasMultipleNextTargets()) {
            for (Integer i = 1; i < item.next.size(); i++) {
                LookupChain lcClone = lc.copy();
                LookupChainItem itemClone = item.clone();
                itemClone.next = new List<SObjectType>{ item.next.get(i) };
                if (processingRoot) {
                    lcClone.root = itemClone;
                    lcClone.items.clear();
                }
                lcClone.add(itemClone);
                chains.add(lcClone);
            }
        }

        if (item.hasNextTarget()) {
            item.next = new List<SObjectType>{ item.next.get(0) };
        }

        if (processingRoot == false) {
            lc.add(item);
        }
    }

    public List<LookupChainItem> getLookupChainItems(SObjectType type, String path) {
        return getLookupChainItems(new List<LookupChainItem>(), new List<SObjectType> { type }, path);
    }

    public List<LookupChainItem> getLookupChainItems(List<LookupChainItem> result, List<SObjectType> objTypes, String path) {
        //System.debug('objTypes: ' + objTypes);
        //System.debug('path: ' + path);

        List<SObjectType> nextObjs = new List<SObjectType>();

        String fragment = '';
        String field = '';

        if (path.contains('.')) {
            fragment = path.substringAfter('.');
            field = path.substringBefore('.');
        } else {
            field = path;
        }

        //System.debug('fragment: ' + fragment);
        //System.debug('field: ' + field);

        for (SObjectType objType : objTypes) {
            DescribeSObjectResult objDesc = objType.getDescribe();
            Map<String, SObjectField> fields = objDesc.fields.getMap();
            LookupChainItem item = new LookupChainItem(objType, path);
            result.add(item);

            item.fragment = fragment;
            item.field = field;

            //System.debug('check key');

            if (fields.containsKey(item.field)) {
                //System.debug('fields key');
                item.fieldDesc = fields.get(item.field).getDescribe();
                //System.debug('fields key 2');
            } else {
                //System.debug('no fields key');
                item.fieldDesc = findLookupRelationship(item.root, item.field);
                //System.debug('no fields key 2');
            }

            if (item.fieldDesc == null) {
                RelationshipMapperException ex = (RelationshipMapperException) Sentry_ExceptionFactory.build(RelationshipMapperException.class);
                ex.setMessage('No field data found for path');
                ex.context.put('field', field);
                ex.context.put('path', path);
                ex.context.put('fragment', fragment);
                item.error = ex;
                item.isValid = false;
                continue;
            }

            // this is the end of the search
            if (item.fragment == '') {
                //System.debug('No fragment');
                continue;
            }

            if (String.isEmpty(item.fieldDesc.getRelationshipName())) {
                //System.debug('No relationship');
                continue;
            }

            //System.debug('get next');
            item.next = item.fieldDesc.getReferenceTo();
            if (item.next != null && item.next.size() > 1) {
                item.isPolymorphicReference = true;
            }
            //System.debug('add next');
            nextObjs.addAll(item.next);
            //System.debug('get relationship name');
            item.relationshipName = item.fieldDesc.getRelationshipName();

            item.debug();

            if (item.field == item.relationshipName) {
                item.field = findLookupRelationshipField(item.root, item.field);
            }

        }

        //System.debug('nextObjs: ' + nextObjs);

        if (nextObjs.isEmpty()) {
            return result;
        }

        return getLookupChainItems(result, nextObjs, fragment);
    }

    private DescribeFieldResult findLookupRelationship(SObjectType objType, String relationshipName) {
        //System.debug('find lookup relationship (' + objType + ', ' + relationshipName + ')');
        if (lookupFields.containsKey(objType) == false) {
            cacheLookupFields(objType);
        }

        return lookupFields.get(objType).get(relationshipName);
    }

    private String findLookupRelationshipField(SObjectType objType, String relationshipName) {
        if (lookupFields.containsKey(objType) == false) {
            cacheLookupFields(objType);
        }

        return lookupFields.get(objType).get(relationshipName).name;
    }

    private void cacheLookupFields(SObjectType objType) {
        Map<String, SObjectField> fields = objType.getDescribe().fields.getMap();
        Map<String, DescribeFieldResult> foundLookupFields = new Map<String, DescribeFieldResult>();
        lookupFields.put(objType, foundLookupFields);

        for (String fieldName : fields.keySet()) {
            SObjectField field = fields.get(fieldName);
            DescribeFieldResult fieldDesc = field.getDescribe();
            //System.debug('fr: ' + fieldDesc.relationshipName);
            if (String.isNotBlank(fieldDesc.relationshipName)) {
                foundLookupFields.put(fieldDesc.relationshipName, fieldDesc);
            }
        }
    }

    /**
     * Holds a list of the nodes for one interpretation of a path.
     *
     * If you look up the path 'Account.Name' on 'Opportunity', you would then have:
     * Lookup Chain 1:
     * - root: Opportunity
     *   items:
     *     - root: Opportunity
     *       field: Account
     *     - root: Account
     *       field: Name
     *
     * If you look up the path 'Account.Owner.Name' on 'Opportunity', you would then have:
     * Lookup Chain 1:
     * - root: Opportunity
     *   items:
     *     - root: Opportunity
     *       field: Account
     *     - root: Account
     *       field: Owner
     *     - root: User
     *       field: Name
     * Lookup Chain 2:
     * - root: Opportunity
     *   items:
     *     - root: Opportunity
     *       field: Account
     *     - root: Account
     *       field: Owner
     *     - root: Queue
     *       field: Name
     *
     */
    public class LookupChain {
        public List<LookupChainItem> items = new List<LookupChainItem>();
        public LookupChainItem root = null;
        public Boolean isValid = null;

        public LookupChain(LookupChainItem root) {
            this.root = root;
            items.add(root);
            isValid = root.isValid;
        }

        public LookupChain copy() {
            LookupChain cloned = new LookupChain(root);
            cloned.items = items.clone();
            return cloned;
        }

        public void add(LookupChainItem lci) {
            items.add(lci);
            isValid = lci.isValid;
        }

        public LookupChainItem last() {
            return items.get(items.size() - 1);
        }

        public Boolean isAccessible() {
            for (LookupChainItem item : items) {
                if (item.fieldDesc.isAccessible() == false) {
                    return false;
                }
            }

            return true;
        }

        public Boolean hasPolymorphicReference() {
            for (LookupChainItem item : items) {
                if (item.isPolymorphicReference) {
                    return true;
                }
            }

            return false;
        }

        public String getSOQLPath() {
            List<String> pathItems = new List<String>();

            for (LookupChainItem item : items) {
                if (item.relationshipName != null) {
                    pathItems.add(item.relationshipName);
                } else {
                    pathItems.add(item.field);
                }
            }

            return String.join(pathItems, '.');
        }

        public void debug() {
            System.debug('LookupChain:');
            System.debug('  root: ' + root);
            for (LookupChainItem item : items) {
                item.debug();
            }
            System.debug('');
        }
    }

    /**
     * Represents an independent step in resolving a path.
     */
    public class LookupChainItem {
        public RelationshipMapperException error = null;
        public Boolean isValid = true;
        public String path = null;
        public String field = null;
        public String fragment = null;
        public SObjectType root = null;
        public List<SObjectType> next = null;
        public DescribeFieldResult fieldDesc = null;
        public String relationshipName = null;
        public Boolean isPolymorphicReference = false;

        public LookupChainItem(SObjectType type, String path) {
            this.path = path;
            this.root = type;
        }

        public SObjectType nextTarget() {
            return (next == null || next.isEmpty() || next.size() > 1) ? null : next.get(0);
        }

        public Boolean hasNextTarget() {
            return next != null && next.isEmpty() == false;
        }

        public Boolean hasMultipleNextTargets() {
            return next != null && next.size() > 1;
        }

        public void debug() {
            System.debug('LookupChainItem:');
            System.debug('  root: ' + root);
            System.debug('  path: ' + path);
            System.debug('  field: ' + field);
            System.debug('  fragment: ' + fragment);
            System.debug('  relationshipName: ' + relationshipName);
            System.debug('  next: ' + next);
            System.debug('  nextTarget: ' + nextTarget());
        }
    }
}