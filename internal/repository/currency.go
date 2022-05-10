package repository

import "gitlab.ozon.dev/zBlur/homework-2/internal/domain"

type CurrencyRepository interface {
	Retrieve(currencyCode string) domain.CurrencyRetrieve

	//Create(currency *domain.Currency) error
	//Update(currency *domain.Currency) error
	//Delete(currencyCode string) error
}
