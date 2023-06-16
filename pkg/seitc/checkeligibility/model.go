package checkeligibility

import "errors"

var (
	ErrObjNil = errors.New("object is nil")
)

type CheckEligibilityRequest struct {
	ConsumerID       string     `json:"consumerId,omitempty"`
	ConversationID   string     `json:"conversationId,omitempty"`
	EnrollDevice     bool       `json:"enrollDevice,omitempty"`
	DeviceData       DeviceData `json:"deviceData,omitempty"`
	EncData          string     `json:"encData,omitempty"`
	MessageID        string     `json:"messageId,omitempty"`
	TokenLocation    string     `json:"tokenLocation,omitempty"`
	PresentationMode string     `json:"presentationMode,omitempty"`
	TokenRequestorID string     `json:"tokenRequestorId,omitempty"`
	PanSource        string     `json:"panSource,omitempty"`
}

type CheckEligibilityResponse struct {
	CardMetaData     CardMetaData `json:"cardMetaData"`
	ConversationID   string       `json:"conversationId"`
	EnrollmentID     string       `json:"enrollmentId"`
	MessageID        string       `json:"messageId"`
	PanReferenceID   string       `json:"panReferenceId"`
	StatusCode       string       `json:"statusCode"`
	TokenRequestorID string       `json:"tokenRequestorId"`
	StatusMessage    *string      `json:"statusMessage,omitempty"`
}

type DeviceData struct {
	DeviceID                  string  `json:"deviceId"`
	DeviceName                string  `json:"deviceName"`
	DevicePlatform            string  `json:"devicePlatform"`
	DeviceType                string  `json:"deviceType"`
	Language                  string  `json:"language"`
	Locale                    string  `json:"locale"`
	NfcCapable                bool    `json:"nfcCapable"`
	OsVersion                 string  `json:"osVersion"`
	WalletID                  string  `json:"walletId"`
	AndroidId                 *string `json:"androidId,omitempty"`
	Brand                     *string `json:"brand,omitempty"`
	DeviceBluetoothMac        *string `json:"deviceBluetoothMac,omitempty"`
	DeviceIp                  *string `json:"deviceIp,omitempty"`
	DeviceLocation            *string `json:"deviceLocation,omitempty"`
	DeviceLocationExtended    *string `json:"deviceLocationExtended,omitempty"`
	DevicePhoneNumber         *string `json:"devicePhoneNumber,omitempty"`
	DeviceProfile             *string `json:"deviceProfile,omitempty"`
	DeviceTimeZone            *string `json:"deviceTimeZone,omitempty"`
	DeviceZoneManager         *string `json:"deviceZoneManager,omitempty"`
	HostDeviceId              *string `json:"hostDeviceId,omitempty"`
	ImeiNumber                *string `json:"imeiNumber,omitempty"`
	Manufacturer              *string `json:"manufacturer,omitempty"`
	Model                     *string `json:"model,omitempty"`
	NetworkType               *string `json:"networkType,omitempty"`
	OsBuildId                 *string `json:"osBuildId,omitempty"`
	OsId                      *string `json:"osId,omitempty"`
	Seid                      *string `json:"seid,omitempty"`
	SerialNumber              *string `json:"serialNumber,omitempty"`
	SimSerialNumber           *string `json:"simSerialNumber,omitempty"`
	WalletAccountEmailAddress *string `json:"walletAccountEmailAddress,omitempty"`
}

type CardMetaData struct {
	TermsAndConditionsID  string           `json:"termsAndConditionsId"`
	CardAssets            []*CardAsset     `json:"cardAssets,omitempty"`
	CardBackgroundColor   *string          `json:"cardBackgroundColor,omitempty"`
	CardForegroundColor   *string          `json:"cardForegroundColor,omitempty"`
	Coresidence           *Coresidence     `json:"coresidence,omitempty"`
	EnabledServices       *EnabledServices `json:"enabledServices,omitempty"`
	IssuerAppData         *IssuerAppData   `json:"issuerAppData,omitempty"`
	IssuerData            *IssuerData      `json:"issuerData,omitempty"`
	LabelColor            *string          `json:"labelColor,omitempty"`
	LongDescription       *string          `json:"longDescription,omitempty"`
	ShortDescription      *string          `json:"shortDescription,omitempty"`
	TermsAndConditionsUrl *string          `json:"termsAndConditionsUrl,omitempty"`
}

type Coresidence struct {
	AlternateNetworks []*AlternateNetwork `json:"alternateNetworks,omitempty"`
	CardImageModel    *string             `json:"cardImageModel,omitempty"`
}
type IssuerData struct {
	CustomerServiceTelephone *string `json:"customerServiceTelephone,omitempty"`
	IssuerEmail              *string `json:"issuerEmail,omitempty"`
	IssuerName               *string `json:"issuerName,omitempty"`
	IssuerPrivacyUrl         *string `json:"issuerPrivacyUrl,omitempty"`
	IssuerWebsite            *string `json:"issuerWebsite,omitempty"`
	OnlineBankingUrl         *string `json:"onlineBankingUrl,omitempty"`
}

type IssuerAppData struct {
	IssuerAppAction         *string `json:"issuerAppAction,omitempty"`
	IssuerAppAddress        *string `json:"issuerAppAddress,omitempty"`
	IssuerAppExtraTextValue *string `json:"issuerAppExtraTextValue,omitempty"`
	IssuerAppName           *string `json:"issuerAppName,omitempty"`
}

type EnabledServices struct {
	MerchantPresentedQR bool `json:"merchantPresentedQR"`
}

type AlternateNetwork struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CardAsset struct {
	AssetType   string  `json:"assetType"`
	Guid        string  `json:"guid"`
	MimeType    *string `json:"mimeType,omitempty"`
	PixelHeight *string `json:"pixelHeight,omitempty"`
	PixelWidth  *string `json:"pixelWidth,omitempty"`
}

type EncDataOfCheckEligibilityRequest struct {
	EncPanData        PanData         `json:"encPanData"`
	EncCardholderData *CardholderData `json:"encCardholderData,omitempty"`
}

type PanData struct {
	Pan            string  `json:"pan"`
	PanExpDate     string  `json:"panExpDate"`
	Cvv            string  `json:"cvv"`
	PanReferenceId *string `json:"panReferenceId,omitempty"`
}

type CardholderData struct {
	CardholderName *string         `json:"cardholderName,omitempty"`
	BillingAddress *BillingAddress `json:"billingAddress,omitempty"`
}

type BillingAddress struct {
	Line1       *string `json:"line1,omitempty"`
	Line2       *string `json:"line2,omitempty"`
	Line3       *string `json:"line3,omitempty"`
	Line4       *string `json:"line4,omitempty"`
	Line5       *string `json:"line5,omitempty"`
	City        *string `json:"city,omitempty"`
	State       *string `json:"state,omitempty"`
	CountryCode *string `json:"countryCode,omitempty"`
	PostalCode  *string `json:"postalCode,omitempty"`
}
