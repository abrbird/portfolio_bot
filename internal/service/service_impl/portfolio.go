package service_impl

import (
	"context"
	"github.com/abrbird/portfolio_bot/internal/domain"
	"github.com/abrbird/portfolio_bot/internal/repository"
)

type PortfolioService struct{}

func (p PortfolioService) RetrieveOrCreate(ctx context.Context, portfolioData domain.PortfolioCreate, repo repository.PortfolioRepository) domain.PortfolioRetrieve {
	return repo.RetrieveOrCreate(ctx, portfolioData)
}
