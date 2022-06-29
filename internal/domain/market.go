package domain

import pb "github.com/abrbird/portfolio_bot/pkg/api"

const MarketItemStockType = "Stock"
const MarketItemETFType = "ETF"
const MarketItemCryptoCurrencyType = "CryptoCurrency"

type MarketItem struct {
	Id    int64
	Code  string
	Type  string
	Title string
}

type MarketItemRetrieve struct {
	MarketItem *MarketItem
	Error      error
}

type MarketItemsRetrieve struct {
	MarketItems []MarketItem
	Error       error
}

type MarketPrice struct {
	MarketItemId int64
	Price        float64
	Timestamp    int64
}

type MarketPriceRetrieve struct {
	MarketPrice *MarketPrice
	Error       error
}

type MarketPricesRetrieve struct {
	MarketPrices []MarketPrice
	Error        error
}

func (mir *MarketItemsRetrieve) GetPBItems() []*pb.MarketItem {
	mis := make([]*pb.MarketItem, len(mir.MarketItems))
	for i, item := range mir.MarketItems {
		mis[i] = &pb.MarketItem{
			Id:    item.Id,
			Code:  item.Code,
			Type:  item.Type,
			Title: item.Title,
		}
	}
	return mis
}

func (mpr *MarketPricesRetrieve) GetPBItems() []*pb.MarketPrice {
	mps := make([]*pb.MarketPrice, len(mpr.MarketPrices))
	for i, item := range mpr.MarketPrices {
		mps[i] = &pb.MarketPrice{
			MarketItemId: item.MarketItemId,
			Price:        item.Price,
			Timestamp:    item.Timestamp,
		}
	}
	return mps
}
