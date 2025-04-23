package maxtests

import (
	"cape.buf.dev/gen/go/cape/pt-apis/grpc/go/privatetech/auth/v3/authv3grpc"
	numberinguserv1grpc "cape.buf.dev/gen/go/cape/pt-apis/grpc/go/privatetech/numbering/user/v1alpha/userv1alphagrpc"
	authv3 "cape.buf.dev/gen/go/cape/pt-apis/protocolbuffers/go/privatetech/auth/v3"
	typev1 "cape.buf.dev/gen/go/cape/pt-apis/protocolbuffers/go/privatetech/common/type/v1"
	numberinguserv1 "cape.buf.dev/gen/go/cape/pt-apis/protocolbuffers/go/privatetech/numbering/user/v1alpha"
	pkiv1 "cape.buf.dev/gen/go/cape/pt-apis/protocolbuffers/go/privatetech/pki/v1"
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/max-cape/stripe-go-test/paymentintent"
	gcpki "github.com/private-tech-inc/golang-common/pkg/pki"
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"math/big"
	"testing"

	"cape.buf.dev/gen/go/cape/pt-apis/grpc/go/privatetech/product/v1/productv1grpc"
	productv1 "cape.buf.dev/gen/go/cape/pt-apis/protocolbuffers/go/privatetech/product/v1"

	"cape.buf.dev/gen/go/cape/pt-apis/grpc/go/privatetech/checkout/v2alpha/checkoutv2alphagrpc"
	checkoutv2alpha "cape.buf.dev/gen/go/cape/pt-apis/protocolbuffers/go/privatetech/checkout/v2alpha"
	"github.com/google/uuid"
	"github.com/max-cape/stripe-go-test"

	"cape.buf.dev/gen/go/cape/pt-apis/grpc/go/privatetech/esim/v3/esimv3grpc"
	esimv3 "cape.buf.dev/gen/go/cape/pt-apis/protocolbuffers/go/privatetech/esim/v3"
)

const (
	nucleusDevAddress         = "dev.d-qqb2l2.us-east-1.api.cape.ninja"
	metadataIdempotencyKeyKey = "idempotency-key"
)

func TestOnboarding(t *testing.T) {
	// Generate long-lived keypair
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	require.NoError(t, err)

	require.NoError(t, WritePrivateKeyToPTDirectory("my-test.pem", privKey))

	// Publish public key against nuc-dev
	unauthedConn, err := grpc.NewClient(nucleusDevAddress, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	require.NoError(t, err)
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(unauthedConn)

	authClient := authv3grpc.NewAuthAPIClient(unauthedConn)

	publicKeyProto := &pkiv1.PublicKey{
		PublicKey: &pkiv1.PublicKey_EcPublicKey{
			EcPublicKey: &pkiv1.EllipticCurvePublicKey{
				Type: &pkiv1.EllipticCurvePublicKey_Ed25519{
					Ed25519: &pkiv1.Ed25519PublicKey{
						Rfc8032Encoded: pubKey,
					},
				},
			},
		},
	}

	_, err = authClient.PublishPublicKey(context.Background(), &authv3.PublishPublicKeyRequest{
		PublicKey: publicKeyProto,
	})

	// Perform challenge/response against nuc-dev to get access token
	request := &authv3.GetAuthenticationChallengeRequest{
		PublicKey: publicKeyProto,
	}
	authChallengeResp, err := authClient.GetAuthenticationChallenge(context.Background(), request)
	require.NoError(t, err)
	nonce := new(big.Int).SetBytes(authChallengeResp.GetNonce().GetValue())

	accessToken, _, err := gcpki.VerifyChallengeAndGetTokens(context.Background(), authClient, publicKeyProto, privKey, nonce, false)
	require.NoError(t, err)

	credsInstance := &creds{AccessToken: accessToken}

	// Redial nuc-dev to use the access token
	authedConn, err := grpc.NewClient(
		nucleusDevAddress,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		grpc.WithPerRPCCredentials(credsInstance),
	)
	require.NoError(t, err)

	// Reserve msisdn against numbering
	numberingClient := numberinguserv1grpc.NewMSISDNUserAPIClient(authedConn)
	_, err = numberingClient.ReserveMSISDNUserByNPA(context.Background(), &numberinguserv1.ReserveMSISDNUserByNPARequest{
		Npa: "404",
	})
	require.NoError(t, err)

	// Fetch the list of available products and select the Provisioning product
	catalogClient := productv1grpc.NewCatalogAPIClient(authedConn)
	response, err := catalogClient.ListProducts(context.Background(), &productv1.ListProductsRequest{})
	require.NoError(t, err)
	var productId string
	for _, item := range response.GetProductCatalog().GetItems() {
		if item.GetProduct().GetProductName() == "Privacy Unlimited" {
			productId = item.GetProduct().GetProductId()
			break
		}
	}

	// Make a checkout session against nuc-dev
	idempotencyKey := uuid.New().String()
	checkoutClient := checkoutv2alphagrpc.NewCheckoutAPIClient(authedConn)
	cart := &checkoutv2alpha.Cart{
		Products: []*productv1.Product{{ProductId: productId}},
	}
	checkoutSessionRequest := &checkoutv2alpha.CreateCheckoutSessionRequest{
		Cart: cart,
		PurchaserInformation: &checkoutv2alpha.PurchaserInformation{
			PostalAddress: &typev1.PostalAddress{
				PostalCode: "55555",
				RegionCode: "US",
			},
		},
	}
	createCheckoutSessionResp, err := checkoutClient.CreateCheckoutSession(IdempotencyKey(context.Background(), idempotencyKey), checkoutSessionRequest)
	require.NoError(t, err)

	stripeSession := createCheckoutSessionResp.GetCheckoutSession().GetStripeSession()
	clientSecret := stripeSession.GetClientSecret()

	splitSecret := strings.Split(clientSecret, "_")

	paymentIntentID := splitSecret[0] + "_" + splitSecret[1]

	//errOnRequiresAction := true
	paymentMethod := "pm_card_amex_threeDSecureNotSupported"
	paymentMethodType := "card"
	paymentIntentConfirmParams := &stripe.PaymentIntentConfirmParams{
		Params: stripe.Params{},
		//CaptureMethod:         "automatic",
		//ConfirmationToken:     nil,
		// This means that the intent will
		//ErrorOnRequiresAction: &errOnRequiresAction,
		//Expand:                nil,
		//Mandate:               nil,
		//MandateData:           nil,
		//OffSession:            nil,
		PaymentMethod: &paymentMethod,
		//PaymentMethodData:    nil,
		//PaymentMethodOptions: nil,
		PaymentMethodTypes: []*string{&paymentMethodType},
		//RadarOptions:       nil,
		//ReceiptEmail:       nil,
		ReturnURL: nil,
		//SetupFutureUsage: nil,
		//Shipping:         nil,
		UseStripeSDK: nil,
		ClientSecret: &clientSecret,
	}
	stripe.Key = stripeSession.GetPublishableKey()

	intent, err := paymentintent.Confirm(paymentIntentID, paymentIntentConfirmParams)
	require.NoError(t, err)
	fmt.Println(intent)

	// Use simservice/getLPAForSubscriber to figure out when our esim is available
	simServiceClient := esimv3grpc.NewSimServiceAPIClient(authedConn)

	counter := 0
outerFor:
	for {
		counter += 1
		lpaInfoResp, err := simServiceClient.GetLPAInfoForSubscriber(context.Background(), &esimv3.GetLPAInfoForSubscriberRequest{})
		require.NoError(t, err)
		for _, info := range lpaInfoResp.GetLpaInfo() {
			if info.GetSimStatus() == esimv3.SIMStatus_SIM_STATUS_RELEASED {
				break outerFor
			}
		}
		// Bail after some attempts, and assume the sim isn't released.
		require.Less(t, counter, 30, "Sim hasn't been released for subscriber")
		time.Sleep(3 * time.Second)
	}
	fmt.Println("Successfully released Esim from nuc-dev")
}

type creds struct {
	AccessToken string
}

// GetRequestMetadata returns the metadata to be embedded inside a RPC.
func (c *creds) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("no access token provided")
	}
	return map[string]string{
		"authorization": "bearer " + c.AccessToken,
	}, nil
}

// RequireTransportSecurity returns whether or not transport security is required for this RPC.
func (c *creds) RequireTransportSecurity() bool {
	return false
}

// IdempotencyKey adds an idempotency-key value to gRPC metadata.
func IdempotencyKey(ctx context.Context, idempotencyKey string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, metadataIdempotencyKeyKey, idempotencyKey)
}

func WritePrivateKeyToPTDirectory(path string, key any) error {
	dir, err := GetTestDirectory()
	if err != nil {
		return err
	}

	relativePath := filepath.Join(dir, path)

	pkcs8Bytes, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		return err
	}
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: pkcs8Bytes,
	}
	pemBytes := pem.EncodeToMemory(block)
	if err := os.WriteFile(relativePath, pemBytes, 0700); err != nil {
		return err
	}

	return nil
}

func GetTestDirectory() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(homeDir, ".pt-test"), nil
}

// WriteToPTDirectory will write a file to disk relative to the Private Tech directory.
func WriteToPTDirectory(bytes []byte, name string) (string, error) {
	path, err := joinDirectory(name)
	if err != nil {
		return "", err
	}

	if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
		return "", err
	}
	if err = os.WriteFile(path, bytes, 0o400); err != nil {
		return "", err
	}
	return path, nil
}
