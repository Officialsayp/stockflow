package domain

type PaymentMethod string

const (
	PaymentMethodDebitCard   PaymentMethod = "debit_card"
	PaymentMethodCreditCard  PaymentMethod = "credit_card"
	PaymentMethodSBP         PaymentMethod = "sbp"
	PaymentMethodUponReceipt PaymentMethod = "upon_receipt"
)
