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
