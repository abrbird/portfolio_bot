package service

// Business logic

type Service interface {
	User() UserService
	Currency() CurrencyService
	Portfolio() PortfolioService
	PortfolioItem() PortfolioItemService
	MarketItem() MarketItemService
	MarketPrice() MarketPriceService
}
