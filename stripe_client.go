package stripe

// Client is the Stripe client. It contains all the different services available.
type Client struct {
	// stripeClientStruct: The beginning of the section generated from our OpenAPI spec
	// stripeClientStruct: The end of the section generated from our OpenAPI spec

	V2CoreEvents *v2CoreEventService

	backend Backend
	key     string
}

// NewClient creates a new Stripe [Client] with the given API key.
func NewClient(key string, opts ...ClientOption) *Client {
	usage := []string{"stripe_client_new"}
	client := &Client{}
	cfg := clientConfig{key: key, usage: usage}
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		opt(&cfg)
	}
	initClient(client, cfg)
	return client
}

func initClient(client *Client, cfg clientConfig) {
	if cfg.backends == nil {
		cfg.backends = &Backends{
			API:         &UsageBackend{B: GetBackend(APIBackend), Usage: cfg.usage},
			Connect:     &UsageBackend{B: GetBackend(ConnectBackend), Usage: cfg.usage},
			Uploads:     &UsageBackend{B: GetBackend(UploadsBackend), Usage: cfg.usage},
			MeterEvents: &UsageBackend{B: GetBackend(MeterEventsBackend), Usage: cfg.usage},
		}
	}
	backends := cfg.backends
	key := cfg.key
	// enough information on the Client to make API calls
	client.backend = backends.API
	client.V2CoreEvents = &v2CoreEventService{B: backends.API, Key: key}
	client.key = key

	// stripeClientInit: The beginning of the section generated from our OpenAPI spec
	// stripeClientInit: The end of the section generated from our OpenAPI spec
}

type clientConfig struct {
	backends *Backends
	usage    []string
	key      string
}

// ClientOption allows for functional options to be passed to the NewClient constructor.
type ClientOption func(*clientConfig)

// WithBackends allows for setting a custom [*Backends] struct when creating a new client.
// This is useful for testing or when you want to use a different backend constructed
// from [NewBackendsWithConfig].
func WithBackends(backends *Backends) ClientOption {
	return func(c *clientConfig) {
		c.backends = backends
	}
}

// ParseEventNotification parses a Stripe event from the payload and verifies its signature.
// It returns a union of all possible event notification types that implement EventNotificationContainer.
func (c *Client) ParseEventNotification(payload []byte, header string, secret string, opts ...WebhookOption) (EventNotificationContainer, error) {
	if err := ValidatePayload(payload, header, secret, opts...); err != nil {
		return nil, err
	}

	return EventNotificationFromJSON(payload, *c)
}

// ConstructEvent initializes an Event object from a JSON webhook payload, validating
// the Stripe-Signature header using the specified signing secret. Returns an error
// if the body or Stripe-Signature header provided are unreadable, if the
// signature doesn't match, or if the timestamp for the signature is older than
// WebhookDefaultTolerance.
//
// NOTE: Stripe will only send Webhook signing headers after you have retrieved
// your signing secret from the Stripe dashboard:
// https://dashboard.stripe.com/webhooks
//
// This will return an error if the event API version does not match the
// APIVersion constant.
func (c *Client) ConstructEvent(payload []byte, header string, secret string, opts ...WebhookOption) (Event, error) {
	return ConstructEvent(payload, header, secret, opts...)
}
