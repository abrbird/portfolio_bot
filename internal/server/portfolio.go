package server

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	pb "gitlab.ozon.dev/zBlur/homework-2/pkg/api"
)

func (t tserver) RetrieveOrCreatePortfolio(ctx context.Context, req *pb.CreatePortfolioRequest) (*pb.Portfolio, error) {
	portfolioRetrieved := t.repo.Portfolio().RetrieveOrCreate(
		domain.PortfolioCreate{
			UserId:           domain.UserId(req.GetUserId()),
			BaseCurrencyCode: t.conf.Application.BaseCurrency,
		},
	)
	if portfolioRetrieved.Error != nil {
		return nil, portfolioRetrieved.Error
	}

	portfolioItemsRetrieved := t.repo.PortfolioItem().RetrievePortfolioItems(portfolioRetrieved.Portfolio.Id)
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
