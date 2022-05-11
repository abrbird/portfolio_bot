package server

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-2/pkg/api"
)

func (t tserver) AvailableMarketItems(context.Context, *api.Empty) (*api.MarketItemsResponse, error) {
	configMarketItems := t.conf.Application.GetDomainMarketItems()
	availableMarketItemsRetrieve := t.serv.MarketItem().RetrieveMany(configMarketItems, t.repo.MarketItem())

	if availableMarketItemsRetrieve.Error != nil {
		return nil, availableMarketItemsRetrieve.Error
	}

	return &api.MarketItemsResponse{MarketItems: availableMarketItemsRetrieve.GetPBItems()}, nil
}

func (t tserver) MarketItemsPrices(ctx context.Context, req *api.MarketItemPricesRequest) (*api.MarketItemPricesResponse, error) {
	marketPricesRetrieve := t.serv.MarketPrice().GetMarketItemPrices(
		req.GetMarketItemId(),
		req.GetStartTimestamp(),
		req.GetEndTimestamp(),
		req.GetInterval(),
		t.repo.MarketPrice(),
	)
	if marketPricesRetrieve.Error != nil {
		return nil, marketPricesRetrieve.Error
	}
	return &api.MarketItemPricesResponse{MarketPrices: marketPricesRetrieve.GetPBItems()}, nil
}
