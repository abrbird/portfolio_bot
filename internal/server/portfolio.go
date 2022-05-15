package server

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	pb "gitlab.ozon.dev/zBlur/homework-2/pkg/api"
)

func (t tserver) RetrieveOrCreatePortfolio(ctx context.Context, req *pb.CreatePortfolioRequest) (*pb.Portfolio, error) {

	portfolioRetrieved := t.serv.Portfolio().RetrieveOrCreate(
		ctx,
		domain.PortfolioCreate{
			UserId:           domain.UserId(req.GetUserId()),
			BaseCurrencyCode: t.conf.Application.BaseCurrency,
		},
		t.repo.Portfolio(),
	)
	if portfolioRetrieved.Error != nil {
		return nil, portfolioRetrieved.Error
	}

	portfolioItemsRetrieved := t.serv.PortfolioItem().RetrieveMany(ctx, portfolioRetrieved.Portfolio.Id, t.repo.PortfolioItem())
	if portfolioItemsRetrieved.Error != nil {
		return nil, portfolioItemsRetrieved.Error
	}

	portfolio := pb.Portfolio{
		Id:               portfolioRetrieved.Portfolio.Id,
		UserId:           portfolioRetrieved.Portfolio.UserId.ToInt64(),
		BaseCurrencyCode: portfolioRetrieved.Portfolio.BaseCurrencyCode,
		Items:            portfolioItemsRetrieved.GetPBItems(),
	}

	return &portfolio, nil
}

func (t tserver) RetrieveOrCreatePortfolioItem(ctx context.Context, req *pb.CreatePortfolioItemRequest) (*pb.PortfolioItem, error) {
	portfolioItemRetrieved := t.serv.PortfolioItem().RetrieveOrCreate(
		ctx,
		domain.PortfolioItemCreate{
			PortfolioId:  req.GetPortfolioId(),
			MarketItemId: req.GetMarketItemId(),
			Price:        req.GetPrice(),
			Volume:       req.GetVolume(),
		},
		t.repo.PortfolioItem(),
	)
	if portfolioItemRetrieved.Error != nil {
		return nil, portfolioItemRetrieved.Error
	}
	portfolioItem := pb.PortfolioItem{
		Id:           portfolioItemRetrieved.PortfolioItem.Id,
		PortfolioId:  portfolioItemRetrieved.PortfolioItem.PortfolioId,
		MarketItemId: portfolioItemRetrieved.PortfolioItem.MarketItemId,
		Price:        portfolioItemRetrieved.PortfolioItem.Price,
		Volume:       portfolioItemRetrieved.PortfolioItem.Volume,
	}

	return &portfolioItem, nil
}

func (t tserver) DeletePortfolioItem(ctx context.Context, req *pb.DeletePortfolioItemRequest) (*pb.Empty, error) {
	err := t.serv.PortfolioItem().Delete(ctx, req.GetId(), t.repo.PortfolioItem())
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
