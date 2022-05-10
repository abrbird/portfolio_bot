package repository

import "gitlab.ozon.dev/zBlur/homework-2/internal/domain"

type MarketItemRepository interface {
	RetrieveById(marketItemId int64) domain.MarketItemRetrieve
	Retrieve(code string, type_ string) domain.MarketItemRetrieve
	RetrieveByType(codes []string, type_ string) *domain.MarketItemsRetrieve

	//Create(marketItem *domain.MarketItem) error
	//Update(marketItem *domain.MarketItem) error
	//Delete(marketItemId int64) error
}

type MarketPriceRepository interface {
	RetrieveLast(marketItemId int64) domain.MarketPriceRetrieve
	RetrieveInterval(marketItemId int64, start int64, end int64) *domain.MarketPricesRetrieve
	Create(marketPrice *domain.MarketPrice) error

	//Update(marketPrice *domain.MarketPrice) error
	//Delete(marketPriceId domain.UserId) error
}
