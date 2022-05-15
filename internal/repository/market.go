package repository

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
)

type MarketItemRepository interface {
	RetrieveById(ctx context.Context, marketItemId int64) domain.MarketItemRetrieve
	Retrieve(ctx context.Context, code string, type_ string) domain.MarketItemRetrieve
	RetrieveByType(ctx context.Context, codes []string, type_ string) *domain.MarketItemsRetrieve
}

type MarketPriceRepository interface {
	RetrieveLast(ctx context.Context, marketItemId int64) domain.MarketPriceRetrieve
	RetrieveInterval(ctx context.Context, marketItemId int64, start int64, end int64) *domain.MarketPricesRetrieve
	Create(ctx context.Context, marketPrice *domain.MarketPrice) (bool, error)
	BulkCreate(ctx context.Context, marketPrices *[]domain.MarketPrice) (int64, error)
}
