package tests

import (
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/zBlur/homework-2/internal/data_collector/yahoo_finance"
	"gitlab.ozon.dev/zBlur/homework-2/internal/data_collector/yahoo_finance/client"
	"gitlab.ozon.dev/zBlur/homework-2/internal/data_collector/yahoo_finance/client_mock"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository/mock_repository"
	"testing"
)

func TestYahooFinanceGetHistoricalMap(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	interval := yahoo_finance.Interval1H
	range_ := yahoo_finance.Range3MO
	marketItems := []domain.MarketItem{
		{
			Id:    int64(1),
			Code:  "BTC",
			Type:  "Crypto",
			Title: "BTC cc",
		},
		{
			Id:    int64(1),
			Code:  "EBAY",
			Type:  "Stock",
			Title: "Ebay stock",
		},
	}

	historicalMap := make(map[domain.MarketItem]yahoo_finance.Historical, len(marketItems))
	historicalMap[marketItems[0]] = yahoo_finance.Historical{
		PreviousClose:      nil,
		Symbol:             marketItems[0].Code,
		ChartPreviousClose: 0.5,
		Timestamp:          []int64{1, 3, 5, 6, 7, 10},
		Close:              []float64{0.1, 0.2, 0.3, 0.3, 0.2, 0.9},
		End:                nil,
		Start:              nil,
		DataGranularity:    1,
	}
	historicalMap[marketItems[1]] = yahoo_finance.Historical{
		PreviousClose:      nil,
		Symbol:             marketItems[1].Code,
		ChartPreviousClose: 8.5,
		Timestamp:          []int64{2, 4, 5, 7, 8},
		Close:              []float64{9.1, 7.2, 8.3, 8.2, 8.9},
		End:                nil,
		Start:              nil,
		DataGranularity:    1,
	}

	mockClient := client_mock.NewClientMock(mc)
	mockClient.GetHistoricalMapMock.Expect(
		marketItems,
		interval,
		range_,
	).Return(
		&historicalMap,
		nil,
	)

	mockRepo := mock_repository.NewMarketPriceRepositoryMock(mc)

	for _, marketItem := range marketItems {
		historical := historicalMap[marketItem]
		marketPrices := historical.ToMarketPriceArray(marketItem)

		mockRepo.BulkCreateMock.When(
			context.Background(),
			marketPrices,
		).Then(
			int64(len(*marketPrices)),
			nil,
		)
	}

	successCount, errorsCount, err := client.Collect(marketItems, interval, range_, mockRepo, mockClient)

	assert.Nil(t, err)
	assert.Equal(
		t,
		successCount,
		int64(len(historicalMap[marketItems[0]].Timestamp)+len(historicalMap[marketItems[1]].Timestamp)),
	)
	assert.Equal(t, errorsCount, int64(0))
}
