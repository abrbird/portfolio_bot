package service_impl

import (
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository"
)

type MarketItemService struct{}

//func (m MarketItemService) RetrieveById(marketItemId int64, repo repository.MarketItemRepository) <-chan domain.MarketItemRetrieve {
//	channel := make(chan domain.MarketItemRetrieve)
//
//	go func() {
//		channel <- repo.RetrieveById(marketItemId)
//		close(channel)
//	}()
//
//	return channel
//}
//
//func (m MarketItemService) Retrieve(code string, type_ string, repo repository.MarketItemRepository) <-chan domain.MarketItemRetrieve {
//	channel := make(chan domain.MarketItemRetrieve)
//
//	go func() {
//		channel <- repo.Retrieve(code, type_)
//		close(channel)
//	}()
//
//	return channel
//}
//
//func (m MarketItemService) RetrieveByType(codes []string, type_ string, repo repository.MarketItemRepository) <-chan domain.MarketItemsRetrieve {
//	channel := make(chan domain.MarketItemsRetrieve)
//
//	go func() {
//		channel <- *repo.RetrieveByType(codes, type_)
//		close(channel)
//	}()
//
//	return channel
//}

func (m MarketItemService) RetrieveById(marketItemId int64, repo repository.MarketItemRepository) domain.MarketItemRetrieve {
	return repo.RetrieveById(marketItemId)
}

func (m MarketItemService) Retrieve(code string, type_ string, repo repository.MarketItemRepository) domain.MarketItemRetrieve {
	return repo.Retrieve(code, type_)
}

func (m MarketItemService) RetrieveByType(codes []string, type_ string, repo repository.MarketItemRepository) domain.MarketItemsRetrieve {
	return *repo.RetrieveByType(codes, type_)
}
