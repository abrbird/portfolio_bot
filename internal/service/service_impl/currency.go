package service_impl

import (
	"context"
	"github.com/abrbird/portfolio_bot/internal/domain"
	"github.com/abrbird/portfolio_bot/internal/repository"
)

type CurrencyService struct{}

func (c CurrencyService) Retrieve(ctx context.Context, code string, repo repository.CurrencyRepository) domain.CurrencyRetrieve {
	return repo.Retrieve(ctx, code)
}
