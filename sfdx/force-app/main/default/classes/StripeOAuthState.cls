public with sharing class StripeOAuthState {

    public class OAuthInvalidStateException extends Sentry_Exception {}

    private static Map<String, String> JWTHeader = new Map<String, String>{
        'alg' => 'HS256',
        'typ' => 'JWT'
    };
    private static String JWTAlgo = 'hmacSHA256';

    // maps to line ~89 in state_encryption_algo.rb
    private static final String FIELD_SALESFORCE_ACCOUNT_ID = 'org_id';
    private static final String FIELD_SALESFORCE_NAMESPACE = 'sf_ns';
    private static final String FIELD_SALESFORCE_INSTANCE_TYPE = 'sf_it';
    private static final String FIELD_SALESFORCE_INSTANCE_SUBDOMAIN = 'sf_sd';
    private static final String FIELD_SALESFORCE_INSTANCE_URL = 'sf_iu';
    private static final String FIELD_USER_ID = 'uid';
    private static final String FIELD_CSAC_ID = 'csac';
    private static final String FIELD_CSRF = 'csrf';
    private static final String FIELD_STRIPE_ACCOUNT_ID = 'st_id';
    private static final String FIELD_STRIPE_ACCOUNT_LIVEMODE = 'st_lm';
    private static final String FIELD_STRIPE_ACCOUNT_NAME = 'st_na';
    private static final String FIELD_PRIMARY_STRIPE_ACCOUNT_ID = 'pst_id';
    private static final String FIELD_PRIMARY_STRIPE_ACCOUNT_LIVEMODE = 'pst_lm';
    private static final String FIELD_LIVEMODE = 'livemode';
    private static final String FIELD_ORG_KEY = 'org_key';


    public Map<String, Object> data;

    public String org_id {
        get { return (String) data.get(FIELD_SALESFORCE_ACCOUNT_ID); }
        set { data.put(FIELD_SALESFORCE_ACCOUNT_ID, value); }
    }

    public String namespace {
        get { return (String) data.get(FIELD_SALESFORCE_NAMESPACE); }
        set { data.put(FIELD_SALESFORCE_NAMESPACE, value); }
    }

    public String instance_type {
        get { return (String) data.get(FIELD_SALESFORCE_INSTANCE_TYPE); }
        set { data.put(FIELD_SALESFORCE_INSTANCE_TYPE, value); }
    }

    public String instance_subdomain {
        get { return (String) data.get(FIELD_SALESFORCE_INSTANCE_SUBDOMAIN); }
        set { data.put(FIELD_SALESFORCE_INSTANCE_SUBDOMAIN, value); }
    }

    public String user_id {
        get { return (String) data.get(FIELD_USER_ID); }
        set { data.put(FIELD_USER_ID, value); }
    }

    public String csac_id {
        get { return (String) data.get(FIELD_CSAC_ID); }
        set { data.put(FIELD_CSAC_ID, value); }
    }

    public String csrf {
        get { return (String) data.get(FIELD_CSRF); }
        set { data.put(FIELD_CSRF, value); }
    }

    public String stripe_account_id {
        get { return (String) data.get(FIELD_STRIPE_ACCOUNT_ID); }
        set { data.put(FIELD_STRIPE_ACCOUNT_ID, value); }
    }

    public String stripe_account_livemode {
        get { return (String) data.get(FIELD_STRIPE_ACCOUNT_LIVEMODE); }
        set { data.put(FIELD_STRIPE_ACCOUNT_LIVEMODE, value); }
    }

    public String stripe_account_name {
        get { return (String) data.get(FIELD_STRIPE_ACCOUNT_NAME); }
        set { data.put(FIELD_STRIPE_ACCOUNT_NAME, value); }
    }

    public String primary_stripe_account_id {
        get { return (String) data.get(FIELD_PRIMARY_STRIPE_ACCOUNT_ID); }
        set { data.put(FIELD_PRIMARY_STRIPE_ACCOUNT_ID, value); }
    }

    public String primary_stripe_account_livemode {
        get { return (String) data.get(FIELD_PRIMARY_STRIPE_ACCOUNT_LIVEMODE); }
        set { data.put(FIELD_PRIMARY_STRIPE_ACCOUNT_LIVEMODE, value); }
    }

    public String salesforce_instance_url {
        get { return (String) data.get(FIELD_SALESFORCE_INSTANCE_URL); }
        set { data.put(FIELD_SALESFORCE_INSTANCE_URL, value); }
    }

    public Boolean auth_live_mode {
        get { return data.get(FIELD_LIVEMODE) != null ? Boolean.valueOf(data.get(FIELD_LIVEMODE)) : null; }
        set { if (value != null) { data.put(FIELD_LIVEMODE, Boolean.valueOf(value) ? 'true' : 'false'); } }
    }

    public String salesforce_package_key {
        get { return (String) data.get(FIELD_ORG_KEY); }
        set { data.put(FIELD_ORG_KEY, value); }
    }

    public static StripeOAuthState create() {
        Stripe_Connection__c conn = utilities.getStripeConnection();
        Stripe_Account__c acct = utilities.getPrimaryStripeAccount();
        Url instanceUrl = new Url(Url.getOrgDomainUrl().toExternalForm());
        String subdomain = instanceUrl.getHost().split('\\.')[0];
        return new StripeOAuthState(
                UserInfo.getOrganizationId(),
                EncodingUtil.convertToHex(Crypto.generateAesKey(256)),
                subdomain,
                acct,
                conn);
    }

    public static StripeOAuthState from_encrypted_state(String encrypted_state) {
        Map<String, Object> data = StripeOAuthState.decrypt(encrypted_state);
        if (data == null) {
            OAuthInvalidStateException ex = (OAuthInvalidStateException) Sentry_ExceptionFactory.build(OAuthInvalidStateException.class);
            ex.setMessage('Invalid OAuth State');
            ex.context.put('method', 'from_encrypted_state');
            ex.context.put('encrypted_state', encrypted_state);

            throw ex;
        }
        return new StripeOAuthState(data);
    }

    private StripeOAuthState() {
    }

    public StripeOAuthState(Map<String, Object> data) {
        this.data = data;
    }

    public StripeOAuthState(String org_id, String csrf, String salesforce_instance_subdomain, Stripe_Account__c primaryAccount, Stripe_Connection__c conn) {
        this.data = new Map<String, Object>();
        this.org_id = org_id;
        this.csrf = csrf;
        this.namespace = constants.NAMESPACE;
        this.instance_type = Sentry_Environment.getInstanceType().name();
        this.instance_subdomain = salesforce_instance_subdomain;
        this.primary_stripe_account_id = primary_stripe_account_id;
        this.primary_stripe_account_livemode = primary_stripe_account_livemode;
        this.instance_subdomain = salesforce_instance_subdomain;
        this.primary_stripe_account_id = (primaryAccount != null) ? primaryAccount.Stripe_ID__c : null;
        this.primary_stripe_account_livemode = (primaryAccount != null && primaryAccount.Is_Live_Mode__c) ? 'true' : 'false';
        this.csac_id = (primaryAccount != null) ? primaryAccount.CSAC_ID__c : null;
        this.salesforce_instance_url = Url.getOrgDomainUrl().toExternalForm();

        // On first load, there will be no API_Key__C on the Stripe_Connection__c
        // object. We need to generate a new key and store it on the account config in V2 auth.
        // Then in addStripeAccount in setupAssistant, we will store the new package key on the conn
        if (String.isBlank(conn.API_Key__c)) {
            Blob blb = Crypto.generateAesKey(128);
            String packageKey  = EncodingUtil.convertToHex(blb);
            this.salesforce_package_key = packageKey;
        } else {
            // User has already gone through setup and has a package key, use the same key for future auth
            this.salesforce_package_key = conn.API_Key__c;
        }
    }

    override public String toString() {
        return StripeOAuthState.encrypt(data);
    }

    public String toJSON() {
        return JSON.serialize(data);
    }

    // jwt encoding reffed from
    // https://help.salesforce.com/s/articleView?id=sfdo.EL_Create_a_JWT.htm&type=5

    public static String toSanitizedBase64(Blob b) {
        String base64 = EncodingUtil.base64Encode(b);
        base64 = base64.replaceAll('=', '');
        base64 = base64.replaceAll('\\+', '-');
        base64 = base64.replaceAll('\\/', '_');

        return base64;
    }

    public static Blob fromSanitizedBase64(String s) {
        String base64 = s.replaceAll('-', '+');
        base64 = base64.replaceAll('_', '/');
        // '=' * (4 - b64.length % 4)
        base64 = base64 + '='.repeat(4 - Math.mod(base64.length(), 4));
        return EncodingUtil.base64Decode(base64);
    }

    public static String encrypt(Map<String, Object> state) {
        try {
            String header_b64 = toSanitizedBase64(Blob.valueOf(JSON.serialize(JWTHeader)));
            String state_b64 = toSanitizedBase64(Blob.valueOf(JSON.serialize(state)));
            Blob hash;

            if (utilities.isUsingV2Middleware()) {
                String str_b64 = header_b64 + '.' + state_b64;
                hash = Crypto.generateMac(JWTAlgo, Blob.valueOf(str_b64), utilities.getOAuthStateSigningKey());            
            } else {
                hash = Crypto.generateMac(JWTAlgo, Blob.valueOf(state_b64), utilities.getOAuthStateSigningKey());
            }

            String hash_b64 = toSanitizedBase64(hash);
            return header_b64 + '.' + state_b64 + '.' + hash_b64;
        } catch (Exception e) {
            errorLogger.create('StripeOAuthState.encrypt failed', e);
            return null;
        }
    }

    public static Map<String, Object> decrypt(String st) {
        List<String> parts;
        String input;
        try {
            if (utilities.isUsingV2Middleware()) {
                parts = st.split('\\.');
                String header = parts[0];
                String payload = parts[1];
                input = header + '.' + payload;
            } else {
                parts = st.split('\\.');
                input = parts[1];
            }

            Blob hash = fromSanitizedBase64(parts[2]);
            Boolean verified = Crypto.verifyHmac(JWTAlgo, Blob.valueOf(input), utilities.getOAuthStateSigningKey(), hash);
            if (!verified) {
                OAuthInvalidStateException ex = (OAuthInvalidStateException) Sentry_ExceptionFactory.build(OAuthInvalidStateException.class);
                ex.setMessage('Invalid OAuth State In Decrypt');
                ex.context.put('method', 'decrypt');
                ex.context.put('hash', toSanitizedBase64(hash));
                ex.context.put('encrypted_state', st);
                throw ex;
            }
            return (Map<String, Object>) JSON.deserializeUntyped(fromSanitizedBase64(parts[1]).toString());
        } catch (Exception e) {
            errorLogger.create('StripeOAuthState.decrypt failed', e);
            return null;
        }
    }
}