package service

// Business logic

type Service interface {
	User() UserService
	Currency() CurrencyService
	MarketItem() MarketItemService
	MarketPrice() MarketPriceService
}
