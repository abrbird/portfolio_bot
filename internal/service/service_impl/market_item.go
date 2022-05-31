package service_impl

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository"
)

type MarketItemService struct{}

func (m MarketItemService) RetrieveById(ctx context.Context, marketItemId int64, repo repository.MarketItemRepository) domain.MarketItemRetrieve {
	return repo.RetrieveById(ctx, marketItemId)
}

func (m MarketItemService) Retrieve(ctx context.Context, code string, type_ string, repo repository.MarketItemRepository) domain.MarketItemRetrieve {
	return repo.Retrieve(ctx, code, type_)
}

func (m MarketItemService) RetrieveByType(ctx context.Context, codes []string, type_ string, repo repository.MarketItemRepository) domain.MarketItemsRetrieve {
	return *repo.RetrieveByType(ctx, codes, type_)
}

func (m MarketItemService) RetrieveMany(
	ctx context.Context,
	marketItems []domain.MarketItem,
	repo repository.MarketItemRepository,
) domain.MarketItemsRetrieve {
	marketItemsTypeCodesMap := make(map[string][]string, 0)
	for _, mi := range marketItems {
		if _, ok := marketItemsTypeCodesMap[mi.Type]; !ok {
			marketItemsTypeCodesMap[mi.Type] = make([]string, 0)
		}
		marketItemsTypeCodesMap[mi.Type] = append(marketItemsTypeCodesMap[mi.Type], mi.Code)
	}

	availableMarketItems := make([]domain.MarketItem, 0)
	for type_, codes := range marketItemsTypeCodesMap {
		marketItemsRetrieve := repo.RetrieveByType(ctx, codes, type_)
		if marketItemsRetrieve.Error != nil {
			return domain.MarketItemsRetrieve{MarketItems: nil, Error: marketItemsRetrieve.Error}
		}
		availableMarketItems = append(availableMarketItems, marketItemsRetrieve.MarketItems...)
	}

	return domain.MarketItemsRetrieve{MarketItems: availableMarketItems, Error: nil}
}
