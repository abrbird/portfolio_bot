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

func TestPortfolioItemRetrieveMany(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	portfolioId := int64(1)
	portfolioItems := []domain.PortfolioItem{
		{
			Id:           1,
			PortfolioId:  portfolioId,
			Price:        10.01,
			MarketItemId: 10,
			Volume:       1000,
		},
		{
			Id:           5,
			PortfolioId:  portfolioId,
			Price:        202.02001,
			MarketItemId: 103,
			Volume:       243.151,
		},
	}

	mockRepo := mock_repository.NewPortfolioItemRepositoryMock(mc)
	mockRepo.RetrievePortfolioItemsMock.Expect(
		context.Background(),
		portfolioId,
	).Return(
		&domain.PortfolioItemsRetrieve{
			PortfolioItems: portfolioItems,
			Error:          nil,
		},
	)

	portfolioItemService := service_impl.PortfolioItemService{}
	portfolioItemsRetrieve := portfolioItemService.RetrieveMany(context.Background(), portfolioId, mockRepo)

	assert.Nil(t, portfolioItemsRetrieve.Error)
	assert.NotNil(t, portfolioItemsRetrieve.PortfolioItems)
	assert.Equal(t, len(portfolioItemsRetrieve.PortfolioItems), len(portfolioItems))
	for _, pi := range portfolioItemsRetrieve.PortfolioItems {
		if pi.PortfolioId != portfolioId {
			t.Errorf("PortfolioItem.PortfolioId must be equal to portfolioId: %v != %v", pi.PortfolioId, portfolioId)
		}
	}
}

func TestPortfolioItemRetrieveOrCreate(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	portfolioId := int64(1)
	marketItemId := int64(5)
	portfolioItemCreate := domain.PortfolioItemCreate{
		PortfolioId:  portfolioId,
		MarketItemId: marketItemId,
		Price:        10.010,
		Volume:       100,
	}

	mockRepo := mock_repository.NewPortfolioItemRepositoryMock(mc)
	mockRepo.RetrieveOrCreateMock.Expect(
		context.Background(),
		portfolioItemCreate,
	).Return(
		domain.PortfolioItemRetrieve{
			PortfolioItem: &domain.PortfolioItem{
				Id:           1,
				PortfolioId:  portfolioId,
				MarketItemId: marketItemId,
				Price:        portfolioItemCreate.Price,
				Volume:       portfolioItemCreate.Volume,
			},
			Error: nil,
		},
	)

	portfolioItemService := service_impl.PortfolioItemService{}
	portfolioItemRetrieve := portfolioItemService.RetrieveOrCreate(context.Background(), portfolioItemCreate, mockRepo)

	assert.Nil(t, portfolioItemRetrieve.Error)
	assert.NotNil(t, portfolioItemRetrieve.PortfolioItem)
	assert.Equal(t, portfolioItemRetrieve.PortfolioItem.PortfolioId, portfolioItemCreate.PortfolioId)
	assert.Equal(t, portfolioItemRetrieve.PortfolioItem.MarketItemId, portfolioItemCreate.MarketItemId)
	assert.Equal(t, portfolioItemRetrieve.PortfolioItem.Price, portfolioItemCreate.Price)
	assert.Equal(t, portfolioItemRetrieve.PortfolioItem.Volume, portfolioItemCreate.Volume)
}

func TestPortfolioItemDelete(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	portfolioItemId := int64(1)

	mockRepo := mock_repository.NewPortfolioItemRepositoryMock(mc)
	mockRepo.DeleteMock.Expect(
		context.Background(),
		portfolioItemId,
	).Return(
		nil,
	)

	portfolioItemService := service_impl.PortfolioItemService{}
	err := portfolioItemService.Delete(context.Background(), portfolioItemId, mockRepo)

	assert.Nil(t, err)
}
