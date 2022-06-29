package service

import (
	"context"
	"github.com/abrbird/portfolio_bot/internal/domain"
	"github.com/abrbird/portfolio_bot/internal/repository"
)

type MarketItemService interface {
	RetrieveById(ctx context.Context, marketItemId int64, repo repository.MarketItemRepository) domain.MarketItemRetrieve
	Retrieve(ctx context.Context, code string, type_ string, repo repository.MarketItemRepository) domain.MarketItemRetrieve
	RetrieveByType(ctx context.Context, codes []string, type_ string, repo repository.MarketItemRepository) domain.MarketItemsRetrieve
	RetrieveMany(ctx context.Context, marketItems []domain.MarketItem, repo repository.MarketItemRepository) domain.MarketItemsRetrieve
}

type MarketPriceService interface {
	Create(ctx context.Context, marketPrice *domain.MarketPrice, repo repository.MarketPriceRepository) (bool, error)
	BulkCreate(ctx context.Context, marketPrices *[]domain.MarketPrice, repo repository.MarketPriceRepository) (int64, error)
	RetrieveInterval(ctx context.Context, marketItemId int64, start int64, end int64, repo repository.MarketPriceRepository) *domain.MarketPricesRetrieve
	RetrieveLast(ctx context.Context, marketItemIds []int64, repo repository.MarketPriceRepository) *domain.MarketPricesRetrieve
	FillBlanks(ctx context.Context, marketPrices []domain.MarketPrice, intervals []int64) ([]domain.MarketPrice, int)
	RetrieveMarketItemPrices(
		ctx context.Context,
		marketItemId int64,
		startTimeStamp int64,
		endTimeStamp int64,
		interval int64,
		repo repository.MarketPriceRepository,
	) *domain.MarketPricesRetrieve
}
