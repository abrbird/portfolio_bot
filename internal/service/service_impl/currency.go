package service_impl

import (
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository"
)

type CurrencyService struct{}

//func (c CurrencyService) Retrieve(code string, repo repository.CurrencyRepository) <-chan domain.CurrencyRetrieve {
//	channel := make(chan domain.CurrencyRetrieve)
//
//	go func() {
//		channel <- repo.Retrieve(code)
//		close(channel)
//	}()
//
//	return channel
//}

func (c CurrencyService) Retrieve(code string, repo repository.CurrencyRepository) domain.CurrencyRetrieve {
	return repo.Retrieve(code)
}
