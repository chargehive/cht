package cht

type InstrumentType string

const (
	Card     InstrumentType = "card"
	Wallet   InstrumentType = "wallet"
	Bank     InstrumentType = "bank"
	Voucher  InstrumentType = "voucher"
	BNPL     InstrumentType = "bnpl"
	Phone    InstrumentType = "phone"
	Redirect InstrumentType = "redirect"
)

type InstrumentSource string

const (

	// Cards

	Pan          InstrumentSource = "pan"
	NetworkToken InstrumentSource = "network-token"

	// Wallets

	GooglePay  InstrumentSource = "google-pay"
	ApplePay   InstrumentSource = "apple-pay"
	SamsungPay InstrumentSource = "samsung-pay"
	AmazonPay  InstrumentSource = "amazon-pay"
	RevolutPay InstrumentSource = "revolut-pay"
	WeChatPay  InstrumentSource = "wechat-pay"
	AliPay     InstrumentSource = "alipay"
	PayPal     InstrumentSource = "paypal"

	// Phone Bill Payments

	TextToPay InstrumentSource = "text-to-pay"

	// Bank Debits

	Bacs         InstrumentSource = "bacs"
	SEPA         InstrumentSource = "sepa"
	ACH          InstrumentSource = "ach"
	BankTransfer InstrumentSource = "bank-transfer" // Generic bank transfer - transaction reconcile

	// Bank Redirects

	//BanContact InstrumentSource = "bancontact"
	//EPS        InstrumentSource = "eps"
	//IDEAL      InstrumentSource = "ideal"
	//Przelewy24 InstrumentSource = "p24"
	//TWINT      InstrumentSource = "twint"
	//Sofort     InstrumentSource = "sofort"

	// Vouchers & Gift Cards

	//MultiBanco InstrumentSource = "multibanco"

	// Buy Now Pay Later

	Klarna   InstrumentSource = "klarna"
	Afterpay InstrumentSource = "afterpay"
)
