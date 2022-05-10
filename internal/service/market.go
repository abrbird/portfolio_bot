package service

import (
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository"
)

type MarketItemService interface {
	//RetrieveById(marketItemId int64, repo repository.MarketItemRepository) <-chan domain.MarketItemRetrieve
	//Retrieve(code string, type_ string, repo repository.MarketItemRepository) <-chan domain.MarketItemRetrieve
	//RetrieveByType(codes []string, type_ string, repo repository.MarketItemRepository) <-chan domain.MarketItemsRetrieve

	RetrieveById(marketItemId int64, repo repository.MarketItemRepository) domain.MarketItemRetrieve
	Retrieve(code string, type_ string, repo repository.MarketItemRepository) domain.MarketItemRetrieve
	RetrieveByType(codes []string, type_ string, repo repository.MarketItemRepository) domain.MarketItemsRetrieve
}

type MarketPriceService interface {
	//RetrieveInterval(marketItemId int64, start int64, end int64, repo repository.MarketPriceRepository) <-chan domain.MarketPricesRetrieve
	//Create(marketPrice *domain.MarketPrice, repo repository.MarketPriceRepository) <-chan error

	RetrieveInterval(marketItemId int64, start int64, end int64, repo repository.MarketPriceRepository) domain.MarketPricesRetrieve
	Create(marketPrice *domain.MarketPrice, repo repository.MarketPriceRepository) error

	GetIntervals(startTimeStamp int64, endTimeStamp int64, interval int64) []int64
	FillBlanks(marketPrices []domain.MarketPrice, intervals []int64) ([]domain.MarketPrice, int)
	GetMarketItemPrices(
		marketItemId int64,
		startTimeStamp int64,
		endTimeStamp int64,
		interval int64,
		repo repository.MarketPriceRepository,
	) ([]domain.MarketPrice, error)
}
