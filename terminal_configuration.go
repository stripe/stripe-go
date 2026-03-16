//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Security type of the WiFi network. The hash with the corresponding name contains the credentials for this security type.
type TerminalConfigurationWifiType string

// List of values that TerminalConfigurationWifiType can take
const (
	TerminalConfigurationWifiTypeEnterpriseEapPeap TerminalConfigurationWifiType = "enterprise_eap_peap"
	TerminalConfigurationWifiTypeEnterpriseEapTLS  TerminalConfigurationWifiType = "enterprise_eap_tls"
	TerminalConfigurationWifiTypePersonalPsk       TerminalConfigurationWifiType = "personal_psk"
)

// Deletes a Configuration object.
type TerminalConfigurationParams struct {
	Params `form:"*"`
	// An object containing device type specific settings for BBPOS WisePad 3 readers.
	BBPOSWisePad3 *TerminalConfigurationBBPOSWisePad3Params `form:"bbpos_wisepad3"`
	// An object containing device type specific settings for BBPOS WisePOS E readers.
	BBPOSWisePOSE *TerminalConfigurationBBPOSWisePOSEParams `form:"bbpos_wisepos_e"`
	// Configuration for cellular connectivity.
	Cellular *TerminalConfigurationCellularParams `form:"cellular"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Name of the configuration
	Name *string `form:"name"`
	// Configurations for collecting transactions offline.
	Offline *TerminalConfigurationOfflineParams `form:"offline"`
	// Reboot time settings for readers. that support customized reboot time configuration.
	RebootWindow *TerminalConfigurationRebootWindowParams `form:"reboot_window"`
	// An object containing device type specific settings for Stripe S700 readers.
	StripeS700 *TerminalConfigurationStripeS700Params `form:"stripe_s700"`
	// An object containing device type specific settings for Stripe S710 readers.
	StripeS710 *TerminalConfigurationStripeS710Params `form:"stripe_s710"`
	// Tipping configurations for readers that support on-reader tips.
	Tipping *TerminalConfigurationTippingParams `form:"tipping"`
	// An object containing device type specific settings for Verifone P400 readers.
	VerifoneP400 *TerminalConfigurationVerifoneP400Params `form:"verifone_p400"`
	// Configurations for connecting to a WiFi network.
	Wifi *TerminalConfigurationWifiParams `form:"wifi"`
}

// AddExpand appends a new field to expand.
func (p *TerminalConfigurationParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// An object containing device type specific settings for BBPOS WisePad 3 readers.
type TerminalConfigurationBBPOSWisePad3Params struct {
	// A File ID representing an image you want to display on the reader.
	Splashscreen *string `form:"splashscreen"`
}

// An object containing device type specific settings for BBPOS WisePOS E readers.
type TerminalConfigurationBBPOSWisePOSEParams struct {
	// A File ID representing an image to display on the reader
	Splashscreen *string `form:"splashscreen"`
}

// Configuration for cellular connectivity.
type TerminalConfigurationCellularParams struct {
	// Determines whether to allow the reader to connect to a cellular network. Defaults to false.
	Enabled *bool `form:"enabled"`
}

// Configurations for collecting transactions offline.
type TerminalConfigurationOfflineParams struct {
	// Determines whether to allow transactions to be collected while reader is offline. Defaults to false.
	Enabled *bool `form:"enabled"`
}

// Reboot time settings for readers. that support customized reboot time configuration.
type TerminalConfigurationRebootWindowParams struct {
	// Integer between 0 to 23 that represents the end hour of the reboot time window. The value must be different than the start_hour.
	EndHour *int64 `form:"end_hour"`
	// Integer between 0 to 23 that represents the start hour of the reboot time window.
	StartHour *int64 `form:"start_hour"`
}

// An object containing device type specific settings for Stripe S700 readers.
type TerminalConfigurationStripeS700Params struct {
	// A File ID representing an image you want to display on the reader.
	Splashscreen *string `form:"splashscreen"`
}

// An object containing device type specific settings for Stripe S710 readers.
type TerminalConfigurationStripeS710Params struct {
	// A File ID representing an image you want to display on the reader.
	Splashscreen *string `form:"splashscreen"`
}

// Tipping configuration for AED
type TerminalConfigurationTippingAedParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for AUD
type TerminalConfigurationTippingAUDParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for CAD
type TerminalConfigurationTippingCADParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for CHF
type TerminalConfigurationTippingCHFParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for CZK
type TerminalConfigurationTippingCZKParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for DKK
type TerminalConfigurationTippingDKKParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for EUR
type TerminalConfigurationTippingEURParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for GBP
type TerminalConfigurationTippingGBPParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for GIP
type TerminalConfigurationTippingGipParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for HKD
type TerminalConfigurationTippingHKDParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for HUF
type TerminalConfigurationTippingHufParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for JPY
type TerminalConfigurationTippingJPYParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for MXN
type TerminalConfigurationTippingMxnParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for MYR
type TerminalConfigurationTippingMYRParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for NOK
type TerminalConfigurationTippingNOKParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for NZD
type TerminalConfigurationTippingNZDParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for PLN
type TerminalConfigurationTippingPLNParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for RON
type TerminalConfigurationTippingRonParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for SEK
type TerminalConfigurationTippingSEKParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for SGD
type TerminalConfigurationTippingSGDParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for USD
type TerminalConfigurationTippingUSDParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configurations for readers that support on-reader tips.
type TerminalConfigurationTippingParams struct {
	// Tipping configuration for AED
	Aed *TerminalConfigurationTippingAedParams `form:"aed"`
	// Tipping configuration for AUD
	AUD *TerminalConfigurationTippingAUDParams `form:"aud"`
	// Tipping configuration for CAD
	CAD *TerminalConfigurationTippingCADParams `form:"cad"`
	// Tipping configuration for CHF
	CHF *TerminalConfigurationTippingCHFParams `form:"chf"`
	// Tipping configuration for CZK
	CZK *TerminalConfigurationTippingCZKParams `form:"czk"`
	// Tipping configuration for DKK
	DKK *TerminalConfigurationTippingDKKParams `form:"dkk"`
	// Tipping configuration for EUR
	EUR *TerminalConfigurationTippingEURParams `form:"eur"`
	// Tipping configuration for GBP
	GBP *TerminalConfigurationTippingGBPParams `form:"gbp"`
	// Tipping configuration for GIP
	Gip *TerminalConfigurationTippingGipParams `form:"gip"`
	// Tipping configuration for HKD
	HKD *TerminalConfigurationTippingHKDParams `form:"hkd"`
	// Tipping configuration for HUF
	Huf *TerminalConfigurationTippingHufParams `form:"huf"`
	// Tipping configuration for JPY
	JPY *TerminalConfigurationTippingJPYParams `form:"jpy"`
	// Tipping configuration for MXN
	Mxn *TerminalConfigurationTippingMxnParams `form:"mxn"`
	// Tipping configuration for MYR
	MYR *TerminalConfigurationTippingMYRParams `form:"myr"`
	// Tipping configuration for NOK
	NOK *TerminalConfigurationTippingNOKParams `form:"nok"`
	// Tipping configuration for NZD
	NZD *TerminalConfigurationTippingNZDParams `form:"nzd"`
	// Tipping configuration for PLN
	PLN *TerminalConfigurationTippingPLNParams `form:"pln"`
	// Tipping configuration for RON
	Ron *TerminalConfigurationTippingRonParams `form:"ron"`
	// Tipping configuration for SEK
	SEK *TerminalConfigurationTippingSEKParams `form:"sek"`
	// Tipping configuration for SGD
	SGD *TerminalConfigurationTippingSGDParams `form:"sgd"`
	// Tipping configuration for USD
	USD *TerminalConfigurationTippingUSDParams `form:"usd"`
}

// An object containing device type specific settings for Verifone P400 readers.
type TerminalConfigurationVerifoneP400Params struct {
	// A File ID representing an image you want to display on the reader.
	Splashscreen *string `form:"splashscreen"`
}

// Credentials for a WPA-Enterprise WiFi network using the EAP-PEAP authentication method.
type TerminalConfigurationWifiEnterpriseEapPeapParams struct {
	// A File ID representing a PEM file containing the server certificate
	CaCertificateFile *string `form:"ca_certificate_file"`
	// Password for connecting to the WiFi network
	Password *string `form:"password"`
	// Name of the WiFi network
	Ssid *string `form:"ssid"`
	// Username for connecting to the WiFi network
	Username *string `form:"username"`
}

// Credentials for a WPA-Enterprise WiFi network using the EAP-TLS authentication method.
type TerminalConfigurationWifiEnterpriseEapTLSParams struct {
	// A File ID representing a PEM file containing the server certificate
	CaCertificateFile *string `form:"ca_certificate_file"`
	// A File ID representing a PEM file containing the client certificate
	ClientCertificateFile *string `form:"client_certificate_file"`
	// A File ID representing a PEM file containing the client RSA private key
	PrivateKeyFile *string `form:"private_key_file"`
	// Password for the private key file
	PrivateKeyFilePassword *string `form:"private_key_file_password"`
	// Name of the WiFi network
	Ssid *string `form:"ssid"`
}

// Credentials for a WPA-Personal WiFi network.
type TerminalConfigurationWifiPersonalPskParams struct {
	// Password for connecting to the WiFi network
	Password *string `form:"password"`
	// Name of the WiFi network
	Ssid *string `form:"ssid"`
}

// Configurations for connecting to a WiFi network.
type TerminalConfigurationWifiParams struct {
	// Credentials for a WPA-Enterprise WiFi network using the EAP-PEAP authentication method.
	EnterpriseEapPeap *TerminalConfigurationWifiEnterpriseEapPeapParams `form:"enterprise_eap_peap"`
	// Credentials for a WPA-Enterprise WiFi network using the EAP-TLS authentication method.
	EnterpriseEapTLS *TerminalConfigurationWifiEnterpriseEapTLSParams `form:"enterprise_eap_tls"`
	// Credentials for a WPA-Personal WiFi network.
	PersonalPsk *TerminalConfigurationWifiPersonalPskParams `form:"personal_psk"`
	// Security type of the WiFi network. Fill out the hash with the corresponding name to provide the set of credentials for this security type.
	Type *string `form:"type"`
}

// Returns a list of Configuration objects.
type TerminalConfigurationListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// if present, only return the account default or non-default configurations.
	IsAccountDefault *bool `form:"is_account_default"`
}

// AddExpand appends a new field to expand.
func (p *TerminalConfigurationListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Deletes a Configuration object.
type TerminalConfigurationDeleteParams struct {
	Params `form:"*"`
}

// Retrieves a Configuration object.
type TerminalConfigurationRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *TerminalConfigurationRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// An object containing device type specific settings for BBPOS WisePad 3 readers.
type TerminalConfigurationUpdateBBPOSWisePad3Params struct {
	// A File ID representing an image you want to display on the reader.
	Splashscreen *string `form:"splashscreen"`
}

// An object containing device type specific settings for BBPOS WisePOS E readers.
type TerminalConfigurationUpdateBBPOSWisePOSEParams struct {
	// A File ID representing an image to display on the reader
	Splashscreen *string `form:"splashscreen"`
}

// Configuration for cellular connectivity.
type TerminalConfigurationUpdateCellularParams struct {
	// Determines whether to allow the reader to connect to a cellular network. Defaults to false.
	Enabled *bool `form:"enabled"`
}

// Configurations for collecting transactions offline.
type TerminalConfigurationUpdateOfflineParams struct {
	// Determines whether to allow transactions to be collected while reader is offline. Defaults to false.
	Enabled *bool `form:"enabled"`
}

// Reboot time settings for readers. that support customized reboot time configuration.
type TerminalConfigurationUpdateRebootWindowParams struct {
	// Integer between 0 to 23 that represents the end hour of the reboot time window. The value must be different than the start_hour.
	EndHour *int64 `form:"end_hour"`
	// Integer between 0 to 23 that represents the start hour of the reboot time window.
	StartHour *int64 `form:"start_hour"`
}

// An object containing device type specific settings for Stripe S700 readers.
type TerminalConfigurationUpdateStripeS700Params struct {
	// A File ID representing an image you want to display on the reader.
	Splashscreen *string `form:"splashscreen"`
}

// An object containing device type specific settings for Stripe S710 readers.
type TerminalConfigurationUpdateStripeS710Params struct {
	// A File ID representing an image you want to display on the reader.
	Splashscreen *string `form:"splashscreen"`
}

// Tipping configuration for AED
type TerminalConfigurationUpdateTippingAedParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for AUD
type TerminalConfigurationUpdateTippingAUDParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for CAD
type TerminalConfigurationUpdateTippingCADParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for CHF
type TerminalConfigurationUpdateTippingCHFParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for CZK
type TerminalConfigurationUpdateTippingCZKParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for DKK
type TerminalConfigurationUpdateTippingDKKParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for EUR
type TerminalConfigurationUpdateTippingEURParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for GBP
type TerminalConfigurationUpdateTippingGBPParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for GIP
type TerminalConfigurationUpdateTippingGipParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for HKD
type TerminalConfigurationUpdateTippingHKDParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for HUF
type TerminalConfigurationUpdateTippingHufParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for JPY
type TerminalConfigurationUpdateTippingJPYParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for MXN
type TerminalConfigurationUpdateTippingMxnParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for MYR
type TerminalConfigurationUpdateTippingMYRParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for NOK
type TerminalConfigurationUpdateTippingNOKParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for NZD
type TerminalConfigurationUpdateTippingNZDParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for PLN
type TerminalConfigurationUpdateTippingPLNParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for RON
type TerminalConfigurationUpdateTippingRonParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for SEK
type TerminalConfigurationUpdateTippingSEKParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for SGD
type TerminalConfigurationUpdateTippingSGDParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for USD
type TerminalConfigurationUpdateTippingUSDParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configurations for readers that support on-reader tips.
type TerminalConfigurationUpdateTippingParams struct {
	// Tipping configuration for AED
	Aed *TerminalConfigurationUpdateTippingAedParams `form:"aed"`
	// Tipping configuration for AUD
	AUD *TerminalConfigurationUpdateTippingAUDParams `form:"aud"`
	// Tipping configuration for CAD
	CAD *TerminalConfigurationUpdateTippingCADParams `form:"cad"`
	// Tipping configuration for CHF
	CHF *TerminalConfigurationUpdateTippingCHFParams `form:"chf"`
	// Tipping configuration for CZK
	CZK *TerminalConfigurationUpdateTippingCZKParams `form:"czk"`
	// Tipping configuration for DKK
	DKK *TerminalConfigurationUpdateTippingDKKParams `form:"dkk"`
	// Tipping configuration for EUR
	EUR *TerminalConfigurationUpdateTippingEURParams `form:"eur"`
	// Tipping configuration for GBP
	GBP *TerminalConfigurationUpdateTippingGBPParams `form:"gbp"`
	// Tipping configuration for GIP
	Gip *TerminalConfigurationUpdateTippingGipParams `form:"gip"`
	// Tipping configuration for HKD
	HKD *TerminalConfigurationUpdateTippingHKDParams `form:"hkd"`
	// Tipping configuration for HUF
	Huf *TerminalConfigurationUpdateTippingHufParams `form:"huf"`
	// Tipping configuration for JPY
	JPY *TerminalConfigurationUpdateTippingJPYParams `form:"jpy"`
	// Tipping configuration for MXN
	Mxn *TerminalConfigurationUpdateTippingMxnParams `form:"mxn"`
	// Tipping configuration for MYR
	MYR *TerminalConfigurationUpdateTippingMYRParams `form:"myr"`
	// Tipping configuration for NOK
	NOK *TerminalConfigurationUpdateTippingNOKParams `form:"nok"`
	// Tipping configuration for NZD
	NZD *TerminalConfigurationUpdateTippingNZDParams `form:"nzd"`
	// Tipping configuration for PLN
	PLN *TerminalConfigurationUpdateTippingPLNParams `form:"pln"`
	// Tipping configuration for RON
	Ron *TerminalConfigurationUpdateTippingRonParams `form:"ron"`
	// Tipping configuration for SEK
	SEK *TerminalConfigurationUpdateTippingSEKParams `form:"sek"`
	// Tipping configuration for SGD
	SGD *TerminalConfigurationUpdateTippingSGDParams `form:"sgd"`
	// Tipping configuration for USD
	USD *TerminalConfigurationUpdateTippingUSDParams `form:"usd"`
}

// An object containing device type specific settings for Verifone P400 readers.
type TerminalConfigurationUpdateVerifoneP400Params struct {
	// A File ID representing an image you want to display on the reader.
	Splashscreen *string `form:"splashscreen"`
}

// Credentials for a WPA-Enterprise WiFi network using the EAP-PEAP authentication method.
type TerminalConfigurationUpdateWifiEnterpriseEapPeapParams struct {
	// A File ID representing a PEM file containing the server certificate
	CaCertificateFile *string `form:"ca_certificate_file"`
	// Password for connecting to the WiFi network
	Password *string `form:"password"`
	// Name of the WiFi network
	Ssid *string `form:"ssid"`
	// Username for connecting to the WiFi network
	Username *string `form:"username"`
}

// Credentials for a WPA-Enterprise WiFi network using the EAP-TLS authentication method.
type TerminalConfigurationUpdateWifiEnterpriseEapTLSParams struct {
	// A File ID representing a PEM file containing the server certificate
	CaCertificateFile *string `form:"ca_certificate_file"`
	// A File ID representing a PEM file containing the client certificate
	ClientCertificateFile *string `form:"client_certificate_file"`
	// A File ID representing a PEM file containing the client RSA private key
	PrivateKeyFile *string `form:"private_key_file"`
	// Password for the private key file
	PrivateKeyFilePassword *string `form:"private_key_file_password"`
	// Name of the WiFi network
	Ssid *string `form:"ssid"`
}

// Credentials for a WPA-Personal WiFi network.
type TerminalConfigurationUpdateWifiPersonalPskParams struct {
	// Password for connecting to the WiFi network
	Password *string `form:"password"`
	// Name of the WiFi network
	Ssid *string `form:"ssid"`
}

// Configurations for connecting to a WiFi network.
type TerminalConfigurationUpdateWifiParams struct {
	// Credentials for a WPA-Enterprise WiFi network using the EAP-PEAP authentication method.
	EnterpriseEapPeap *TerminalConfigurationUpdateWifiEnterpriseEapPeapParams `form:"enterprise_eap_peap"`
	// Credentials for a WPA-Enterprise WiFi network using the EAP-TLS authentication method.
	EnterpriseEapTLS *TerminalConfigurationUpdateWifiEnterpriseEapTLSParams `form:"enterprise_eap_tls"`
	// Credentials for a WPA-Personal WiFi network.
	PersonalPsk *TerminalConfigurationUpdateWifiPersonalPskParams `form:"personal_psk"`
	// Security type of the WiFi network. Fill out the hash with the corresponding name to provide the set of credentials for this security type.
	Type *string `form:"type"`
}

// Updates a new Configuration object.
type TerminalConfigurationUpdateParams struct {
	Params `form:"*"`
	// An object containing device type specific settings for BBPOS WisePad 3 readers.
	BBPOSWisePad3 *TerminalConfigurationUpdateBBPOSWisePad3Params `form:"bbpos_wisepad3"`
	// An object containing device type specific settings for BBPOS WisePOS E readers.
	BBPOSWisePOSE *TerminalConfigurationUpdateBBPOSWisePOSEParams `form:"bbpos_wisepos_e"`
	// Configuration for cellular connectivity.
	Cellular *TerminalConfigurationUpdateCellularParams `form:"cellular"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Name of the configuration
	Name *string `form:"name"`
	// Configurations for collecting transactions offline.
	Offline *TerminalConfigurationUpdateOfflineParams `form:"offline"`
	// Reboot time settings for readers. that support customized reboot time configuration.
	RebootWindow *TerminalConfigurationUpdateRebootWindowParams `form:"reboot_window"`
	// An object containing device type specific settings for Stripe S700 readers.
	StripeS700 *TerminalConfigurationUpdateStripeS700Params `form:"stripe_s700"`
	// An object containing device type specific settings for Stripe S710 readers.
	StripeS710 *TerminalConfigurationUpdateStripeS710Params `form:"stripe_s710"`
	// Tipping configurations for readers that support on-reader tips.
	Tipping *TerminalConfigurationUpdateTippingParams `form:"tipping"`
	// An object containing device type specific settings for Verifone P400 readers.
	VerifoneP400 *TerminalConfigurationUpdateVerifoneP400Params `form:"verifone_p400"`
	// Configurations for connecting to a WiFi network.
	Wifi *TerminalConfigurationUpdateWifiParams `form:"wifi"`
}

// AddExpand appends a new field to expand.
func (p *TerminalConfigurationUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// An object containing device type specific settings for BBPOS WisePad 3 readers.
type TerminalConfigurationCreateBBPOSWisePad3Params struct {
	// A File ID representing an image you want to display on the reader.
	Splashscreen *string `form:"splashscreen"`
}

// An object containing device type specific settings for BBPOS WisePOS E readers.
type TerminalConfigurationCreateBBPOSWisePOSEParams struct {
	// A File ID representing an image to display on the reader
	Splashscreen *string `form:"splashscreen"`
}

// Configuration for cellular connectivity.
type TerminalConfigurationCreateCellularParams struct {
	// Determines whether to allow the reader to connect to a cellular network. Defaults to false.
	Enabled *bool `form:"enabled"`
}

// Configurations for collecting transactions offline.
type TerminalConfigurationCreateOfflineParams struct {
	// Determines whether to allow transactions to be collected while reader is offline. Defaults to false.
	Enabled *bool `form:"enabled"`
}

// Reboot time settings for readers. that support customized reboot time configuration.
type TerminalConfigurationCreateRebootWindowParams struct {
	// Integer between 0 to 23 that represents the end hour of the reboot time window. The value must be different than the start_hour.
	EndHour *int64 `form:"end_hour"`
	// Integer between 0 to 23 that represents the start hour of the reboot time window.
	StartHour *int64 `form:"start_hour"`
}

// An object containing device type specific settings for Stripe S700 readers.
type TerminalConfigurationCreateStripeS700Params struct {
	// A File ID representing an image you want to display on the reader.
	Splashscreen *string `form:"splashscreen"`
}

// An object containing device type specific settings for Stripe S710 readers.
type TerminalConfigurationCreateStripeS710Params struct {
	// A File ID representing an image you want to display on the reader.
	Splashscreen *string `form:"splashscreen"`
}

// Tipping configuration for AED
type TerminalConfigurationCreateTippingAedParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for AUD
type TerminalConfigurationCreateTippingAUDParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for CAD
type TerminalConfigurationCreateTippingCADParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for CHF
type TerminalConfigurationCreateTippingCHFParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for CZK
type TerminalConfigurationCreateTippingCZKParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for DKK
type TerminalConfigurationCreateTippingDKKParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for EUR
type TerminalConfigurationCreateTippingEURParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for GBP
type TerminalConfigurationCreateTippingGBPParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for GIP
type TerminalConfigurationCreateTippingGipParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for HKD
type TerminalConfigurationCreateTippingHKDParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for HUF
type TerminalConfigurationCreateTippingHufParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for JPY
type TerminalConfigurationCreateTippingJPYParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for MXN
type TerminalConfigurationCreateTippingMxnParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for MYR
type TerminalConfigurationCreateTippingMYRParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for NOK
type TerminalConfigurationCreateTippingNOKParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for NZD
type TerminalConfigurationCreateTippingNZDParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for PLN
type TerminalConfigurationCreateTippingPLNParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for RON
type TerminalConfigurationCreateTippingRonParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for SEK
type TerminalConfigurationCreateTippingSEKParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for SGD
type TerminalConfigurationCreateTippingSGDParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configuration for USD
type TerminalConfigurationCreateTippingUSDParams struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []*int64 `form:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []*int64 `form:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold *int64 `form:"smart_tip_threshold"`
}

// Tipping configurations for readers that support on-reader tips.
type TerminalConfigurationCreateTippingParams struct {
	// Tipping configuration for AED
	Aed *TerminalConfigurationCreateTippingAedParams `form:"aed"`
	// Tipping configuration for AUD
	AUD *TerminalConfigurationCreateTippingAUDParams `form:"aud"`
	// Tipping configuration for CAD
	CAD *TerminalConfigurationCreateTippingCADParams `form:"cad"`
	// Tipping configuration for CHF
	CHF *TerminalConfigurationCreateTippingCHFParams `form:"chf"`
	// Tipping configuration for CZK
	CZK *TerminalConfigurationCreateTippingCZKParams `form:"czk"`
	// Tipping configuration for DKK
	DKK *TerminalConfigurationCreateTippingDKKParams `form:"dkk"`
	// Tipping configuration for EUR
	EUR *TerminalConfigurationCreateTippingEURParams `form:"eur"`
	// Tipping configuration for GBP
	GBP *TerminalConfigurationCreateTippingGBPParams `form:"gbp"`
	// Tipping configuration for GIP
	Gip *TerminalConfigurationCreateTippingGipParams `form:"gip"`
	// Tipping configuration for HKD
	HKD *TerminalConfigurationCreateTippingHKDParams `form:"hkd"`
	// Tipping configuration for HUF
	Huf *TerminalConfigurationCreateTippingHufParams `form:"huf"`
	// Tipping configuration for JPY
	JPY *TerminalConfigurationCreateTippingJPYParams `form:"jpy"`
	// Tipping configuration for MXN
	Mxn *TerminalConfigurationCreateTippingMxnParams `form:"mxn"`
	// Tipping configuration for MYR
	MYR *TerminalConfigurationCreateTippingMYRParams `form:"myr"`
	// Tipping configuration for NOK
	NOK *TerminalConfigurationCreateTippingNOKParams `form:"nok"`
	// Tipping configuration for NZD
	NZD *TerminalConfigurationCreateTippingNZDParams `form:"nzd"`
	// Tipping configuration for PLN
	PLN *TerminalConfigurationCreateTippingPLNParams `form:"pln"`
	// Tipping configuration for RON
	Ron *TerminalConfigurationCreateTippingRonParams `form:"ron"`
	// Tipping configuration for SEK
	SEK *TerminalConfigurationCreateTippingSEKParams `form:"sek"`
	// Tipping configuration for SGD
	SGD *TerminalConfigurationCreateTippingSGDParams `form:"sgd"`
	// Tipping configuration for USD
	USD *TerminalConfigurationCreateTippingUSDParams `form:"usd"`
}

// An object containing device type specific settings for Verifone P400 readers.
type TerminalConfigurationCreateVerifoneP400Params struct {
	// A File ID representing an image you want to display on the reader.
	Splashscreen *string `form:"splashscreen"`
}

// Credentials for a WPA-Enterprise WiFi network using the EAP-PEAP authentication method.
type TerminalConfigurationCreateWifiEnterpriseEapPeapParams struct {
	// A File ID representing a PEM file containing the server certificate
	CaCertificateFile *string `form:"ca_certificate_file"`
	// Password for connecting to the WiFi network
	Password *string `form:"password"`
	// Name of the WiFi network
	Ssid *string `form:"ssid"`
	// Username for connecting to the WiFi network
	Username *string `form:"username"`
}

// Credentials for a WPA-Enterprise WiFi network using the EAP-TLS authentication method.
type TerminalConfigurationCreateWifiEnterpriseEapTLSParams struct {
	// A File ID representing a PEM file containing the server certificate
	CaCertificateFile *string `form:"ca_certificate_file"`
	// A File ID representing a PEM file containing the client certificate
	ClientCertificateFile *string `form:"client_certificate_file"`
	// A File ID representing a PEM file containing the client RSA private key
	PrivateKeyFile *string `form:"private_key_file"`
	// Password for the private key file
	PrivateKeyFilePassword *string `form:"private_key_file_password"`
	// Name of the WiFi network
	Ssid *string `form:"ssid"`
}

// Credentials for a WPA-Personal WiFi network.
type TerminalConfigurationCreateWifiPersonalPskParams struct {
	// Password for connecting to the WiFi network
	Password *string `form:"password"`
	// Name of the WiFi network
	Ssid *string `form:"ssid"`
}

// Configurations for connecting to a WiFi network.
type TerminalConfigurationCreateWifiParams struct {
	// Credentials for a WPA-Enterprise WiFi network using the EAP-PEAP authentication method.
	EnterpriseEapPeap *TerminalConfigurationCreateWifiEnterpriseEapPeapParams `form:"enterprise_eap_peap"`
	// Credentials for a WPA-Enterprise WiFi network using the EAP-TLS authentication method.
	EnterpriseEapTLS *TerminalConfigurationCreateWifiEnterpriseEapTLSParams `form:"enterprise_eap_tls"`
	// Credentials for a WPA-Personal WiFi network.
	PersonalPsk *TerminalConfigurationCreateWifiPersonalPskParams `form:"personal_psk"`
	// Security type of the WiFi network. Fill out the hash with the corresponding name to provide the set of credentials for this security type.
	Type *string `form:"type"`
}

// Creates a new Configuration object.
type TerminalConfigurationCreateParams struct {
	Params `form:"*"`
	// An object containing device type specific settings for BBPOS WisePad 3 readers.
	BBPOSWisePad3 *TerminalConfigurationCreateBBPOSWisePad3Params `form:"bbpos_wisepad3"`
	// An object containing device type specific settings for BBPOS WisePOS E readers.
	BBPOSWisePOSE *TerminalConfigurationCreateBBPOSWisePOSEParams `form:"bbpos_wisepos_e"`
	// Configuration for cellular connectivity.
	Cellular *TerminalConfigurationCreateCellularParams `form:"cellular"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Name of the configuration
	Name *string `form:"name"`
	// Configurations for collecting transactions offline.
	Offline *TerminalConfigurationCreateOfflineParams `form:"offline"`
	// Reboot time settings for readers. that support customized reboot time configuration.
	RebootWindow *TerminalConfigurationCreateRebootWindowParams `form:"reboot_window"`
	// An object containing device type specific settings for Stripe S700 readers.
	StripeS700 *TerminalConfigurationCreateStripeS700Params `form:"stripe_s700"`
	// An object containing device type specific settings for Stripe S710 readers.
	StripeS710 *TerminalConfigurationCreateStripeS710Params `form:"stripe_s710"`
	// Tipping configurations for readers that support on-reader tips.
	Tipping *TerminalConfigurationCreateTippingParams `form:"tipping"`
	// An object containing device type specific settings for Verifone P400 readers.
	VerifoneP400 *TerminalConfigurationCreateVerifoneP400Params `form:"verifone_p400"`
	// Configurations for connecting to a WiFi network.
	Wifi *TerminalConfigurationCreateWifiParams `form:"wifi"`
}

// AddExpand appends a new field to expand.
func (p *TerminalConfigurationCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type TerminalConfigurationBBPOSWisePad3 struct {
	// A File ID representing an image to display on the reader
	Splashscreen *File `json:"splashscreen"`
}
type TerminalConfigurationBBPOSWisePOSE struct {
	// A File ID representing an image to display on the reader
	Splashscreen *File `json:"splashscreen"`
}
type TerminalConfigurationCellular struct {
	// Whether a cellular-capable reader can connect to the internet over cellular.
	Enabled bool `json:"enabled"`
}
type TerminalConfigurationOffline struct {
	// Determines whether to allow transactions to be collected while reader is offline. Defaults to false.
	Enabled bool `json:"enabled"`
}
type TerminalConfigurationRebootWindow struct {
	// Integer between 0 to 23 that represents the end hour of the reboot time window. The value must be different than the start_hour.
	EndHour int64 `json:"end_hour"`
	// Integer between 0 to 23 that represents the start hour of the reboot time window.
	StartHour int64 `json:"start_hour"`
}
type TerminalConfigurationStripeS700 struct {
	// A File ID representing an image to display on the reader
	Splashscreen *File `json:"splashscreen"`
}
type TerminalConfigurationStripeS710 struct {
	// A File ID representing an image to display on the reader
	Splashscreen *File `json:"splashscreen"`
}
type TerminalConfigurationTippingAed struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingAUD struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingCAD struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingCHF struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingCZK struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingDKK struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingEUR struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingGBP struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingGip struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingHKD struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingHuf struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingJPY struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingMxn struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingMYR struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingNOK struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingNZD struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingPLN struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingRon struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingSEK struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingSGD struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTippingUSD struct {
	// Fixed amounts displayed when collecting a tip
	FixedAmounts []int64 `json:"fixed_amounts"`
	// Percentages displayed when collecting a tip
	Percentages []int64 `json:"percentages"`
	// Below this amount, fixed amounts will be displayed; above it, percentages will be displayed
	SmartTipThreshold int64 `json:"smart_tip_threshold"`
}
type TerminalConfigurationTipping struct {
	Aed *TerminalConfigurationTippingAed `json:"aed"`
	AUD *TerminalConfigurationTippingAUD `json:"aud"`
	CAD *TerminalConfigurationTippingCAD `json:"cad"`
	CHF *TerminalConfigurationTippingCHF `json:"chf"`
	CZK *TerminalConfigurationTippingCZK `json:"czk"`
	DKK *TerminalConfigurationTippingDKK `json:"dkk"`
	EUR *TerminalConfigurationTippingEUR `json:"eur"`
	GBP *TerminalConfigurationTippingGBP `json:"gbp"`
	Gip *TerminalConfigurationTippingGip `json:"gip"`
	HKD *TerminalConfigurationTippingHKD `json:"hkd"`
	Huf *TerminalConfigurationTippingHuf `json:"huf"`
	JPY *TerminalConfigurationTippingJPY `json:"jpy"`
	Mxn *TerminalConfigurationTippingMxn `json:"mxn"`
	MYR *TerminalConfigurationTippingMYR `json:"myr"`
	NOK *TerminalConfigurationTippingNOK `json:"nok"`
	NZD *TerminalConfigurationTippingNZD `json:"nzd"`
	PLN *TerminalConfigurationTippingPLN `json:"pln"`
	Ron *TerminalConfigurationTippingRon `json:"ron"`
	SEK *TerminalConfigurationTippingSEK `json:"sek"`
	SGD *TerminalConfigurationTippingSGD `json:"sgd"`
	USD *TerminalConfigurationTippingUSD `json:"usd"`
}
type TerminalConfigurationVerifoneP400 struct {
	// A File ID representing an image to display on the reader
	Splashscreen *File `json:"splashscreen"`
}
type TerminalConfigurationWifiEnterpriseEapPeap struct {
	// A File ID representing a PEM file containing the server certificate
	CaCertificateFile string `json:"ca_certificate_file"`
	// Password for connecting to the WiFi network
	Password string `json:"password"`
	// Name of the WiFi network
	Ssid string `json:"ssid"`
	// Username for connecting to the WiFi network
	Username string `json:"username"`
}
type TerminalConfigurationWifiEnterpriseEapTLS struct {
	// A File ID representing a PEM file containing the server certificate
	CaCertificateFile string `json:"ca_certificate_file"`
	// A File ID representing a PEM file containing the client certificate
	ClientCertificateFile string `json:"client_certificate_file"`
	// A File ID representing a PEM file containing the client RSA private key
	PrivateKeyFile string `json:"private_key_file"`
	// Password for the private key file
	PrivateKeyFilePassword string `json:"private_key_file_password"`
	// Name of the WiFi network
	Ssid string `json:"ssid"`
}
type TerminalConfigurationWifiPersonalPsk struct {
	// Password for connecting to the WiFi network
	Password string `json:"password"`
	// Name of the WiFi network
	Ssid string `json:"ssid"`
}
type TerminalConfigurationWifi struct {
	EnterpriseEapPeap *TerminalConfigurationWifiEnterpriseEapPeap `json:"enterprise_eap_peap"`
	EnterpriseEapTLS  *TerminalConfigurationWifiEnterpriseEapTLS  `json:"enterprise_eap_tls"`
	PersonalPsk       *TerminalConfigurationWifiPersonalPsk       `json:"personal_psk"`
	// Security type of the WiFi network. The hash with the corresponding name contains the credentials for this security type.
	Type TerminalConfigurationWifiType `json:"type"`
}

// A Configurations object represents how features should be configured for terminal readers.
// For information about how to use it, see the [Terminal configurations documentation](https://docs.stripe.com/terminal/fleet/configurations-overview).
type TerminalConfiguration struct {
	APIResource
	BBPOSWisePad3 *TerminalConfigurationBBPOSWisePad3 `json:"bbpos_wisepad3"`
	BBPOSWisePOSE *TerminalConfigurationBBPOSWisePOSE `json:"bbpos_wisepos_e"`
	Cellular      *TerminalConfigurationCellular      `json:"cellular"`
	Deleted       bool                                `json:"deleted"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Whether this Configuration is the default for your account
	IsAccountDefault bool `json:"is_account_default"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String indicating the name of the Configuration object, set by the user
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value.
	Object       string                             `json:"object"`
	Offline      *TerminalConfigurationOffline      `json:"offline"`
	RebootWindow *TerminalConfigurationRebootWindow `json:"reboot_window"`
	StripeS700   *TerminalConfigurationStripeS700   `json:"stripe_s700"`
	StripeS710   *TerminalConfigurationStripeS710   `json:"stripe_s710"`
	Tipping      *TerminalConfigurationTipping      `json:"tipping"`
	VerifoneP400 *TerminalConfigurationVerifoneP400 `json:"verifone_p400"`
	Wifi         *TerminalConfigurationWifi         `json:"wifi"`
}

// TerminalConfigurationList is a list of Configurations as retrieved from a list endpoint.
type TerminalConfigurationList struct {
	APIResource
	ListMeta
	Data []*TerminalConfiguration `json:"data"`
}
