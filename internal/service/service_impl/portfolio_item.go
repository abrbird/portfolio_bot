package service_impl

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository"
)

type PortfolioItemService struct{}

func (p PortfolioItemService) RetrieveMany(ctx context.Context, portfolioId int64, repo repository.PortfolioItemRepository) *domain.PortfolioItemsRetrieve {
	return repo.RetrievePortfolioItems(ctx, portfolioId)
}

func (p PortfolioItemService) RetrieveOrCreate(ctx context.Context, portfolioItemData domain.PortfolioItemCreate, repo repository.PortfolioItemRepository) domain.PortfolioItemRetrieve {
	return repo.RetrieveOrCreate(ctx, portfolioItemData)
}

func (p PortfolioItemService) Delete(ctx context.Context, portfolioItemId int64, repo repository.PortfolioItemRepository) error {
	return repo.Delete(ctx, portfolioItemId)
}
