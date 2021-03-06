package service

import (
	"context"
	"github.com/abrbird/portfolio_bot/internal/domain"
	"github.com/abrbird/portfolio_bot/internal/repository"
)

type PortfolioService interface {
	RetrieveOrCreate(ctx context.Context, portfolioData domain.PortfolioCreate, repo repository.PortfolioRepository) domain.PortfolioRetrieve
}

type PortfolioItemService interface {
	RetrieveMany(ctx context.Context, portfolioId int64, repo repository.PortfolioItemRepository) *domain.PortfolioItemsRetrieve
	RetrieveOrCreate(ctx context.Context, portfolioItemData domain.PortfolioItemCreate, repo repository.PortfolioItemRepository) domain.PortfolioItemRetrieve
	Delete(ctx context.Context, portfolioItemId int64, repo repository.PortfolioItemRepository) error
}
