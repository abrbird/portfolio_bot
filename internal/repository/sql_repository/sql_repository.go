package sql_repository

import (
	"database/sql"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository"
)

type SQLRepository struct {
	db                      *sql.DB
	userRepository          *SQLUserRepository
	currencyRepository      *SQLCurrencyRepository
	marketItemRepository    *SQLMarketItemRepository
	marketPriceRepository   *SQLMarketPriceRepository
	portfolioRepository     *SQLPortfolioRepository
	portfolioItemRepository *SQLPortfolioItemRepository
}

func New(db *sql.DB) *SQLRepository {
	return &SQLRepository{
		db: db,
	}
}

func (s *SQLRepository) User() repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &SQLUserRepository{store: s}
	}

	return s.userRepository
}

func (s *SQLRepository) Currency() repository.CurrencyRepository {
	if s.currencyRepository == nil {
		s.currencyRepository = &SQLCurrencyRepository{store: s}
	}

	return s.currencyRepository
}

func (s *SQLRepository) MarketItem() repository.MarketItemRepository {
	if s.marketItemRepository == nil {
		s.marketItemRepository = &SQLMarketItemRepository{store: s}
	}

	return s.marketItemRepository
}

func (s *SQLRepository) MarketPrice() repository.MarketPriceRepository {
	if s.marketPriceRepository == nil {
		s.marketPriceRepository = &SQLMarketPriceRepository{store: s}
	}

	return s.marketPriceRepository
}

func (s *SQLRepository) Portfolio() repository.PortfolioRepository {
	if s.portfolioRepository == nil {
		s.portfolioRepository = &SQLPortfolioRepository{store: s}
	}

	return s.portfolioRepository
}

func (s *SQLRepository) PortfolioItem() repository.PortfolioItemRepository {
	if s.portfolioItemRepository == nil {
		s.portfolioItemRepository = &SQLPortfolioItemRepository{store: s}
	}

	return s.portfolioItemRepository
}
