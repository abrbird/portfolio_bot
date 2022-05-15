package service

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository"
)

type PortfolioService interface {
	RetrieveOrCreate(ctx context.Context, portfolioData domain.PortfolioCreate, repo repository.PortfolioRepository) domain.PortfolioRetrieve
}

type PortfolioItemService interface {
	RetrieveMany(ctx context.Context, portfolioId int64, repo repository.PortfolioItemRepository) *domain.PortfolioItemsRetrieve
	RetrieveOrCreate(ctx context.Context, portfolioItemData domain.PortfolioItemCreate, repo repository.PortfolioItemRepository) domain.PortfolioItemRetrieve
	Delete(ctx context.Context, portfolioItemId int64, repo repository.PortfolioItemRepository) error
}
