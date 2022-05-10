package service

import (
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository"
)

type CurrencyService interface {
	//Retrieve(code string, repo repository.CurrencyRepository) <-chan domain.CurrencyRetrieve

	Retrieve(code string, repo repository.CurrencyRepository) domain.CurrencyRetrieve
}
