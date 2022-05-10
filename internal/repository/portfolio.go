package repository

import "gitlab.ozon.dev/zBlur/homework-2/internal/domain"

type PortfolioRepository interface {
	Retrieve(userId domain.UserId) domain.PortfolioRetrieve
	RetrieveOrCreate(portfolioData domain.PortfolioCreate) domain.PortfolioRetrieve
}

type PortfolioItemRepository interface {
	RetrieveOrCreate(portfolioData domain.PortfolioItemCreate) domain.PortfolioItemRetrieve
	Delete(portfolioItemId int64) error
	RetrievePortfolioItems(portfolioId int64) *domain.PortfolioItemsRetrieve
}
