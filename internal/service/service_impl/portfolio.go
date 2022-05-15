package service_impl

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository"
)

type PortfolioService struct{}

func (p PortfolioService) RetrieveOrCreate(ctx context.Context, portfolioData domain.PortfolioCreate, repo repository.PortfolioRepository) domain.PortfolioRetrieve {
	return repo.RetrieveOrCreate(ctx, portfolioData)
}
