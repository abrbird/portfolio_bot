package domain

const MarketItemStockType = "Stock"
const MarketItemETFType = "ETF"
const MarketItemCryptoCurrencyType = "CryptoCurrency"

type MarketItem struct {
	Id    int64
	Code  string
	Type  string
	Title string
}

type MarketItemRetrieve struct {
	MarketItem *MarketItem
	Error      error
}

type MarketItemsRetrieve struct {
	MarketItems []MarketItem
	Error       error
}

type MarketPrice struct {
	MarketItemId int64
	Price        float64
	Timestamp    int64
}

type MarketPriceRetrieve struct {
	MarketPrice *MarketPrice
	Error       error
}

type MarketPricesRetrieve struct {
	MarketPrices []MarketPrice
	Error        error
}
