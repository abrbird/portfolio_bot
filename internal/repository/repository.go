package repository

type Repository interface {
	User() UserRepository
	Currency() CurrencyRepository
	MarketItem() MarketItemRepository
	MarketPrice() MarketPriceRepository
	Portfolio() PortfolioRepository
	PortfolioItem() PortfolioItemRepository
}
