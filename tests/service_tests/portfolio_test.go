package service_tests

import (
	"context"
	"github.com/abrbird/portfolio_bot/internal/domain"
	"github.com/abrbird/portfolio_bot/internal/repository/mock_repository"
	"github.com/abrbird/portfolio_bot/internal/service/service_impl"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPortfolioRetrieveOrCreate(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	userId := 1
	baseCurrencyCode := "USD"
	portfolioCreate := domain.PortfolioCreate{
		UserId:           domain.UserId(userId),
		BaseCurrencyCode: baseCurrencyCode,
	}

	mockRepo := mock_repository.NewPortfolioRepositoryMock(mc)
	mockRepo.RetrieveOrCreateMock.Expect(
		context.Background(),
		portfolioCreate,
	).Return(
		domain.PortfolioRetrieve{
			Portfolio: &domain.Portfolio{
				Id:               1,
				UserId:           domain.UserId(userId),
				BaseCurrencyCode: baseCurrencyCode,
			},
			Error: nil,
		},
	)
	portfolioService := service_impl.PortfolioService{}
	portfolioRetrieve := portfolioService.RetrieveOrCreate(context.Background(), portfolioCreate, mockRepo)

	assert.Nil(t, portfolioRetrieve.Error)
	assert.NotNil(t, portfolioRetrieve.Portfolio)
	assert.Equal(t, portfolioRetrieve.Portfolio.UserId, domain.UserId(userId))
	assert.Equal(t, portfolioRetrieve.Portfolio.BaseCurrencyCode, baseCurrencyCode)
}
