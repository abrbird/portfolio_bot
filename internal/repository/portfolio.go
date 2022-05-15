package repository

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
)

type PortfolioRepository interface {
	Retrieve(ctx context.Context, userId domain.UserId) domain.PortfolioRetrieve
	RetrieveOrCreate(ctx context.Context, portfolioData domain.PortfolioCreate) domain.PortfolioRetrieve
}

type PortfolioItemRepository interface {
	RetrieveOrCreate(ctx context.Context, portfolioData domain.PortfolioItemCreate) domain.PortfolioItemRetrieve
	Delete(ctx context.Context, portfolioItemId int64) error
	RetrievePortfolioItems(ctx context.Context, portfolioId int64) *domain.PortfolioItemsRetrieve
}
