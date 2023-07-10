/*** Stripe Toolbelt v2.0.1 ***/
public with sharing class constants {

    public String ORG_DOMAIN { get; set; }


    // Namespaces
    //fetches the namespace of the package in case we need to prefix it any where in apex for CRUD or accessing anything that will be namespaced in a package
    private static final List<String> RAW_NAMESPACE = String.valueOf(constants.class).split('[.]', 2);
    public static final String NAMESPACE = RAW_NAMESPACE.size() > 1 ? RAW_NAMESPACE[0] : 'c';

    // if empty we are outside of a package
    public static final String NAMESPACE_API = NAMESPACE == 'c' ? '' : NAMESPACE + '__';

    // Sandbox identifcation
    // TODO ORG_ID generation should move to where the instance type stuff is managed and should be passed to sentry
    private static final Organization ORG = [SELECT Id FROM Organization LIMIT 1];
    public static final String ORG_ID = ORG.Id;
    private static final List<Sentry_Environment.InstanceType> SANDBOX_INSTANCE_TYPES = new List<Sentry_Environment.InstanceType>{
        Sentry_Environment.InstanceType.SANDBOX,
        Sentry_Environment.InstanceType.SCRATCH_ORG
    };
    public static final Boolean IS_SANDBOX = SANDBOX_INSTANCE_TYPES.contains(Sentry_Environment.getInstanceType());

    // oAuth callout urls
    public static final String SALESFORCE_INSTANCE_URI = System.Url.getSalesforceBaseUrl().toExternalForm();
    public static final String SALESFORCE_BASE_URI = 'https://' + (IS_SANDBOX ? 'test' : 'login') + '.salesforce.com';
    public static final String SALESFORCE_OAUTH_URI = SALESFORCE_BASE_URI + '/services/oauth2';

    //Stripe Ruby URIS
    public static final String RUBY_SERVICE_BASE_URI = 'https://salesforce.suitesync.io';
    public static final String RUBY_SERVICE_CONFIG_URI = '/v1/configuration';
    public static final String RUBY_SERVICE_TRANSLATE_URI = '/v1/translate';
    public static final String RUBY_SERVICE_TRANSLATE_ALL_URI = '/v1/translate_all';

    //Setup Data Record Name
    public static final String SETUP_DATA_RECORD_NAME = 'SetupData';
    public static final String PACKAGED_PERMISSION_SET_NAME = 'Stripe Connector Integration User';

    public static final String FIELD_CURRENCY_ISO_CODE = 'CurrencyIsoCode';

    public static final String RESOLUTION_STATUS_ERROR = 'Error';

    public static final String SYSTEM_STRIPE = 'stripe';
    public static final String SYSTEM_SALESFORCE = 'salesforce';

    public constants() {
        // empty
    }

    // this seemingly-useless code is used for the bootstrap process in the setup.page
    public String getNamespace() {
        return constants.NAMESPACE;
    }
}
