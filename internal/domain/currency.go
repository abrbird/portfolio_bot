package domain

const CurrencyFiatType = "fiat"
const CurrencyCryptoType = "crypto"

type Currency struct {
	Code  string
	Type  string
	Title string
}

type CurrencyRetrieve struct {
	Currency *Currency
	Error    error
}
