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

func TestMarketPriceRetrieveInterval(t *testing.T) {
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
	assert.GreaterOrEqual(t, marketPricesRetrieve.MarketPrices[0].Timestamp, int64(1))
	assert.LessOrEqual(t, marketPricesRetrieve.MarketPrices[len(marketPricesRetrieve.MarketPrices)-1].Timestamp, int64(12))
}

func TestMarketPriceFillBlanks(t *testing.T) {

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

func TestMarketPriceRetrieveMarketItemPrices(t *testing.T) {
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

func TestMarketPriceRetrieveMarketItemPricesNegative(t *testing.T) {
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

func TestMarketPriceRetrieveLast(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	marketItemIds := []int64{1, 2, 4, 7}

	mockRepo := mock_repository.NewMarketPriceRepositoryMock(mc)
	mockRepo.RetrieveLastMock.When(
		context.Background(),
		1,
	).Then(
		domain.MarketPriceRetrieve{
			MarketPrice: &domain.MarketPrice{
				MarketItemId: 1,
				Price:        0.0,
				Timestamp:    10,
			},
			Error: nil,
		},
	)
	mockRepo.RetrieveLastMock.When(
		context.Background(),
		2,
	).Then(
		domain.MarketPriceRetrieve{
			MarketPrice: nil,
			Error:       nil,
		},
	)
	mockRepo.RetrieveLastMock.When(
		context.Background(),
		4,
	).Then(
		domain.MarketPriceRetrieve{
			MarketPrice: &domain.MarketPrice{
				MarketItemId: 4,
				Price:        10.0,
				Timestamp:    15,
			},
			Error: nil,
		},
	)
	mockRepo.RetrieveLastMock.When(
		context.Background(),
		7,
	).Then(
		domain.MarketPriceRetrieve{
			MarketPrice: &domain.MarketPrice{
				MarketItemId: 7,
				Price:        6.1,
				Timestamp:    9,
			},
			Error: nil,
		},
	)

	marketPriceService := service_impl.MarketPriceService{}
	marketLastPricesRetrieve := marketPriceService.RetrieveLast(context.Background(), marketItemIds, mockRepo)

	assert.Nil(t, marketLastPricesRetrieve.Error)
	assert.NotNil(t, marketLastPricesRetrieve.MarketPrices)
	assert.LessOrEqual(t, len(marketLastPricesRetrieve.MarketPrices), len(marketItemIds))
}

func TestMarketPriceCreate(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	marketPrice := domain.MarketPrice{
		MarketItemId: 1,
		Price:        0.0,
		Timestamp:    1,
	}

	mockRepo := mock_repository.NewMarketPriceRepositoryMock(mc)
	mockRepo.CreateMock.When(
		context.Background(),
		&marketPrice,
	).Then(
		true,
		nil,
	)

	marketPriceService := service_impl.MarketPriceService{}
	marketPriceCreated, err := marketPriceService.Create(context.Background(), &marketPrice, mockRepo)

	assert.Nil(t, err)
	assert.NotNil(t, marketPriceCreated)
	assert.True(t, marketPriceCreated)

	marketPrice_ := domain.MarketPrice{
		MarketItemId: 1,
		Price:        0.5,
		Timestamp:    1,
	}
	mockRepo.CreateMock.When(
		context.Background(),
		&marketPrice_,
	).Then(
		false,
		domain.UnknownError,
	)

	marketPriceCreated, err = marketPriceService.Create(context.Background(), &marketPrice_, mockRepo)
	assert.NotNil(t, err)
	assert.NotNil(t, marketPriceCreated)
	assert.False(t, marketPriceCreated)
}

func TestMarketPriceBulkCreate(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	marketPrices := []domain.MarketPrice{
		{MarketItemId: 1, Price: 1.5, Timestamp: 1},
		{MarketItemId: 1, Price: 2.5, Timestamp: 2},
		{MarketItemId: 1, Price: 3.5, Timestamp: 3},
		{MarketItemId: 1, Price: 4.5, Timestamp: 4},
		{MarketItemId: 1, Price: 5.5, Timestamp: 5},
		{MarketItemId: 1, Price: 6.5, Timestamp: 6},
		{MarketItemId: 1, Price: 7.5, Timestamp: 7},
		{MarketItemId: 1, Price: 8.5, Timestamp: 8},
		{MarketItemId: 1, Price: 9.5, Timestamp: 9},
		{MarketItemId: 1, Price: 10.5, Timestamp: 10},

		{MarketItemId: 2, Price: 11.5, Timestamp: 1},
		{MarketItemId: 2, Price: 12.5, Timestamp: 2},
		{MarketItemId: 2, Price: 13.5, Timestamp: 3},
		{MarketItemId: 2, Price: 14.5, Timestamp: 4},
		{MarketItemId: 2, Price: 15.5, Timestamp: 5},
		{MarketItemId: 2, Price: 16.5, Timestamp: 6},
		{MarketItemId: 2, Price: 17.5, Timestamp: 7},
		{MarketItemId: 2, Price: 18.5, Timestamp: 8},
		{MarketItemId: 2, Price: 19.5, Timestamp: 9},
		{MarketItemId: 2, Price: 15.5, Timestamp: 10},
	}

	mockRepo := mock_repository.NewMarketPriceRepositoryMock(mc)
	mockRepo.BulkCreateMock.Expect(
		context.Background(),
		&marketPrices,
	).Return(
		int64(len(marketPrices)),
		nil,
	)

	marketPriceService := service_impl.MarketPriceService{}
	marketPricesInserted, err := marketPriceService.BulkCreate(context.Background(), &marketPrices, mockRepo)

	assert.Nil(t, err)
	assert.NotNil(t, marketPricesInserted)
	assert.LessOrEqual(t, marketPricesInserted, int64(len(marketPrices)))
}
