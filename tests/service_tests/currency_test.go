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

func TestCurrencyRetrieve(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	code := "USD"
	typeF := "fiat"

	mockRepo := mock_repository.NewCurrencyRepositoryMock(mc)
	mockRepo.RetrieveMock.Expect(
		context.Background(),
		code,
	).Return(
		domain.CurrencyRetrieve{
			Currency: &domain.Currency{
				Code:  code,
				Type:  typeF,
				Title: "Title",
			},
			Error: nil,
		},
	)
	currencyService := service_impl.CurrencyService{}
	currencyRetrieve := currencyService.Retrieve(context.Background(), code, mockRepo)

	assert.Nil(t, currencyRetrieve.Error)
	assert.NotNil(t, currencyRetrieve.Currency)
	assert.Equal(t, currencyRetrieve.Currency.Code, code)
	assert.Equal(t, currencyRetrieve.Currency.Type, typeF)
}
