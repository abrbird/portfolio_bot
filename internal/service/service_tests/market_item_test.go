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

func TestRetrieveByType(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	codeBTC := "BTC"
	codeAMZN := "AMZN"

	codes := []string{
		codeBTC,
		codeAMZN,
	}
	type_ := "crypto"

	mockRepo := mock_repository.NewMarketItemRepositoryMock(mc)
	mockRepo.RetrieveByTypeMock.Expect(
		context.Background(),
		codes,
		type_,
	).Return(
		&domain.MarketItemsRetrieve{
			MarketItems: []domain.MarketItem{
				{
					Id:    1,
					Title: "Title",
					Code:  codeAMZN,
					Type:  type_,
				},
				{
					Id:    1,
					Title: "Title",
					Code:  codeBTC,
					Type:  type_,
				},
			},
			Error: nil,
		},
	)
	marketItemService := service_impl.MarketItemService{}
	marketItemsRetrieve := marketItemService.RetrieveByType(context.Background(), codes, type_, mockRepo)

	assert.Nil(t, marketItemsRetrieve.Error)
	assert.NotNil(t, marketItemsRetrieve.MarketItems)
	assert.Equal(t, 2, len(marketItemsRetrieve.MarketItems))

	for _, mi := range marketItemsRetrieve.MarketItems {

		if mi.Code == codeBTC || mi.Code == codeAMZN {
			assert.Equal(t, mi.Type, type_)
		} else {
			t.Error("code is incorrect")
		}
	}
}

func TestRetrieveMany(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	marketItems := []domain.MarketItem{
		{Code: "BTC", Type: "CryptoCurrency"},
		{Code: "AMZN", Type: "Stock"},
		{Code: "AAAU", Type: "ETF"},
		{Code: "QTTT", Type: "ETF"},
	}

	mockRepo := mock_repository.NewMarketItemRepositoryMock(mc)
	mockRepo.RetrieveByTypeMock.When(
		context.Background(),
		[]string{"BTC"},
		"CryptoCurrency",
	).Then(
		&domain.MarketItemsRetrieve{
			MarketItems: []domain.MarketItem{
				{
					Id:    1,
					Title: "BTC title",
					Code:  "BTC",
					Type:  "CryptoCurrency",
				},
			},
			Error: nil,
		},
	)
	mockRepo.RetrieveByTypeMock.When(
		context.Background(),
		[]string{"AMZN"},
		"Stock",
	).Then(
		&domain.MarketItemsRetrieve{
			MarketItems: []domain.MarketItem{},
			Error:       nil,
		},
	)
	mockRepo.RetrieveByTypeMock.When(
		context.Background(),
		[]string{"AAAU", "QTTT"},
		"ETF",
	).Then(
		&domain.MarketItemsRetrieve{
			MarketItems: []domain.MarketItem{
				{
					Id:    2,
					Title: "ETF title",
					Code:  "AAAU",
					Type:  "ETF",
				},
				{
					Id:    3,
					Title: "QTTT title",
					Code:  "QTTT",
					Type:  "ETF",
				},
			},
			Error: nil,
		},
	)

	marketItemService := service_impl.MarketItemService{}
	marketItemsRetrieve := marketItemService.RetrieveMany(context.Background(), marketItems, mockRepo)

	assert.Nil(t, marketItemsRetrieve.Error)
	assert.NotNil(t, marketItemsRetrieve.MarketItems)
	assert.LessOrEqual(t, len(marketItemsRetrieve.MarketItems), 4)

	for _, mi := range marketItemsRetrieve.MarketItems {
		if mi.Code == "AMZN" || mi.Type == "Stock" {
			t.Errorf("there must not exist MarketItem: %v ", mi)
		}
	}
}
