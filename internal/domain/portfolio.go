package domain

import pb "gitlab.ozon.dev/zBlur/homework-2/pkg/api"

type Portfolio struct {
	Id               int64
	UserId           UserId
	BaseCurrencyCode string
}

type PortfolioCreate struct {
	UserId           UserId
	BaseCurrencyCode string
}

type PortfolioRetrieve struct {
	Portfolio *Portfolio
	Error     error
}

type PortfolioItem struct {
	Id           int64
	PortfolioId  int64
	MarketItemId int64
	Price        float64
	Volume       float64
}

type PortfolioItemCreate struct {
	PortfolioId  int64
	MarketItemId int64
	Price        float64
	Volume       float64
}

type PortfolioItemRetrieve struct {
	PortfolioItem *PortfolioItem
	Error         error
}

type PortfolioItemsRetrieve struct {
	PortfolioItems []PortfolioItem
	Error          error
}

func (pir *PortfolioItemsRetrieve) GetPBItems() []*pb.PortfolioItem {
	pis := make([]*pb.PortfolioItem, len(pir.PortfolioItems))
	for i, pi := range pir.PortfolioItems {
		pis[i] = &pb.PortfolioItem{
			Id:           pi.Id,
			PortfolioId:  pi.PortfolioId,
			MarketItemId: pi.MarketItemId,
			Price:        pi.Price,
			Volume:       pi.Volume,
		}
	}
	return pis
}
