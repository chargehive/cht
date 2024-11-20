package capability

type Capability uint8

const (
	TransmitPan           Capability = 1 // Card Holder Data
	TransmitNetworkToken             = 2
	ConnectorTokenization            = 5 // Connector exchanges payment details for its own token

	Authorization        = 10
	CaptureAuthorization = 11
	Capture              = 12
	PartialCapture       = 13
	Refund               = 14
	PartialRefund        = 15
	Void                 = 16
	VerifyInstrument     = 17

	DisputeNotifications = 30 // Early warning dispute e.g. RDR
	Dispute              = 35 // Chargeback
	DisputeReversal      = 39 // Dispute Challenge / Won

	Reconciliation = 40

	AcceptSchemeTransactionID  = 50
	ProvideSchemeTransactionID = 51

	SchemeMerchantAdviceCodes = 52
	MappedMerchantAdviceCodes = 53

	RecurringPayments = 60
	External3DS       = 65 // Allows external 3ds

	RiskAnalysis = 70
)
