package service

import (
	"context"
	"github.com/abrbird/portfolio_bot/internal/domain"
	"github.com/abrbird/portfolio_bot/internal/repository"
)

type CurrencyService interface {
	Retrieve(ctx context.Context, code string, repo repository.CurrencyRepository) domain.CurrencyRetrieve
}
