package service_tests

import (
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository/mock_repository"
	"gitlab.ozon.dev/zBlur/homework-2/internal/service/service_impl"
	"testing"
)

func TestRetrieveInterval(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := mock_repository.NewMarketPriceRepositoryMock(mc)
	mockRepo.RetrieveIntervalMock.When(
		context.Background(),
		1,
		1,
		12,
	).Then(
		&domain.MarketPricesRetrieve{
			MarketPrices: []domain.MarketPrice{
				{MarketItemId: 1, Price: 9.0, Timestamp: 1},
				{MarketItemId: 1, Price: 11.0, Timestamp: 3},
				{MarketItemId: 1, Price: 10.0, Timestamp: 4},
				{MarketItemId: 1, Price: 12.0, Timestamp: 5},
				{MarketItemId: 1, Price: 7.0, Timestamp: 6},
				{MarketItemId: 1, Price: 10.0, Timestamp: 11},
			},
		},
	)

	marketPriceService := service_impl.MarketPriceService{}
	marketPricesRetrieve := marketPriceService.RetrieveInterval(context.Background(), 1, 1, 12, mockRepo)

	assert.Nil(t, marketPricesRetrieve.Error)
	assert.NotNil(t, marketPricesRetrieve.MarketPrices)
	assert.LessOrEqual(t, len(marketPricesRetrieve.MarketPrices), 12)
}

func TestFillBlanks(t *testing.T) {

	intervals := []int64{1, 5, 10, 15}

	marketPriceService := service_impl.MarketPriceService{}
	marketPrices, blanksCount := marketPriceService.FillBlanks(
		context.Background(),
		[]domain.MarketPrice{
			{MarketItemId: 1, Price: 1, Timestamp: 2},
			{MarketItemId: 1, Price: 2, Timestamp: 4},
			{MarketItemId: 1, Price: 2, Timestamp: 5},
			{MarketItemId: 1, Price: 3, Timestamp: 8},
			{MarketItemId: 1, Price: 4, Timestamp: 10},
			{MarketItemId: 1, Price: 6, Timestamp: 16},
		},
		intervals,
	)

	assert.NotNil(t, marketPrices)
	assert.LessOrEqual(t, len(marketPrices), 4)
	assert.LessOrEqual(t, blanksCount, 2)
}

func TestRetrieveMarketItemPrices(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := mock_repository.NewMarketPriceRepositoryMock(mc)
	mockRepo.RetrieveIntervalMock.When(
		context.Background(),
		1,
		3,
		30,
	).Then(
		&domain.MarketPricesRetrieve{
			MarketPrices: []domain.MarketPrice{
				{MarketItemId: 1, Price: 9.0, Timestamp: 3},
				{MarketItemId: 1, Price: 11.0, Timestamp: 4},
				{MarketItemId: 1, Price: 10.0, Timestamp: 5},
				{MarketItemId: 1, Price: 12.0, Timestamp: 6},
				{MarketItemId: 1, Price: 7.0, Timestamp: 7},
				{MarketItemId: 1, Price: 10.0, Timestamp: 8},
				{MarketItemId: 1, Price: 9.0, Timestamp: 9},
				{MarketItemId: 1, Price: 11.0, Timestamp: 11},
				{MarketItemId: 1, Price: 10.0, Timestamp: 12},
				{MarketItemId: 1, Price: 12.0, Timestamp: 13},
				{MarketItemId: 1, Price: 7.0, Timestamp: 15},
				{MarketItemId: 1, Price: 10.0, Timestamp: 16},
				{MarketItemId: 1, Price: 9.0, Timestamp: 17},
				{MarketItemId: 1, Price: 11.0, Timestamp: 19},
				{MarketItemId: 1, Price: 10.0, Timestamp: 20},
				{MarketItemId: 1, Price: 12.0, Timestamp: 21},
				{MarketItemId: 1, Price: 7.0, Timestamp: 22},
				{MarketItemId: 1, Price: 10.0, Timestamp: 24},
				{MarketItemId: 1, Price: 9.0, Timestamp: 25},
				{MarketItemId: 1, Price: 11.0, Timestamp: 26},
				{MarketItemId: 1, Price: 10.0, Timestamp: 27},
			},
		},
	)

	marketPriceService := service_impl.MarketPriceService{}
	marketPricesRetrieve := marketPriceService.RetrieveMarketItemPrices(
		context.Background(),
		1,
		3,
		30,
		5,
		mockRepo)

	assert.Nil(t, marketPricesRetrieve.Error)
	assert.NotNil(t, marketPricesRetrieve.MarketPrices)
	assert.Equal(t, len(marketPricesRetrieve.MarketPrices), 6)
}

func TestRetrieveMarketItemPricesNegative(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := mock_repository.NewMarketPriceRepositoryMock(mc)
	mockRepo.RetrieveIntervalMock.When(
		context.Background(),
		1,
		3,
		30,
	).Then(
		&domain.MarketPricesRetrieve{
			MarketPrices: []domain.MarketPrice{
				{MarketItemId: 1, Price: 10.0, Timestamp: 5},
				{MarketItemId: 1, Price: 12.0, Timestamp: 6},
				{MarketItemId: 1, Price: 9.0, Timestamp: 9},
				{MarketItemId: 1, Price: 7.0, Timestamp: 15},
				{MarketItemId: 1, Price: 10.0, Timestamp: 16},
				{MarketItemId: 1, Price: 12.0, Timestamp: 21},
			},
		},
	)

	marketPriceService := service_impl.MarketPriceService{}
	marketPricesRetrieve := marketPriceService.RetrieveMarketItemPrices(
		context.Background(),
		1,
		3,
		30,
		2,
		mockRepo)

	assert.NotNil(t, marketPricesRetrieve.Error)
	assert.Nil(t, marketPricesRetrieve.MarketPrices)
}
