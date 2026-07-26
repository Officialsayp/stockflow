package domain

type Money struct {
	Amount   int64
	Currency Currency
}

type Currency string

const (
	CurrencyUSD Currency = "USD"
	CurrencyEUR Currency = "EUR"
	CurrencyRUB Currency = "RUB"
	CurrencyRSD Currency = "RSD"
)
