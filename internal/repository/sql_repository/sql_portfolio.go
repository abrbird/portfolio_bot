package sql_repository

import (
	"context"
	"github.com/abrbird/portfolio_bot/internal/domain"
)

//portfolioItemRepository *SQLPortfolioItemRepository

type SQLPortfolioRepository struct {
	store *SQLRepository
}

type SQLPortfolio struct {
	Id               int64
	UserId           int64
	BaseCurrencyCode string
}

func (r SQLPortfolioRepository) Retrieve(ctx context.Context, userId domain.UserId) domain.PortfolioRetrieve {
	const query = `
		SELECT 
			id,
			userid,
			basecurrencycode
		FROM portfolios_portfolio
		WHERE userid = $1
	`

	sqlPortfolio := &SQLPortfolio{}
	if err := r.store.db.QueryRowContext(
		ctx,
		query,
		userId,
	).Scan(
		&sqlPortfolio.Id,
		&sqlPortfolio.UserId,
		&sqlPortfolio.BaseCurrencyCode,
	); err != nil {
		return domain.PortfolioRetrieve{Portfolio: nil, Error: domain.NotFoundError}
	}
	portfolio := &domain.Portfolio{
		Id:               sqlPortfolio.Id,
		UserId:           domain.UserId(sqlPortfolio.UserId),
		BaseCurrencyCode: sqlPortfolio.BaseCurrencyCode,
	}
	return domain.PortfolioRetrieve{Portfolio: portfolio, Error: nil}
}

func (r SQLPortfolioRepository) RetrieveOrCreate(ctx context.Context, portfolioData domain.PortfolioCreate) domain.PortfolioRetrieve {
	const query = `
		INSERT INTO portfolios_portfolio (
			userid,
			basecurrencycode
		) VALUES (
			$1, $2
		)
		ON CONFLICT ON CONSTRAINT portfolios_portfolio_userid_basecurrencycode_key
		DO UPDATE SET (
			userid,
			basecurrencycode
		) = (
			$1, $2
		) WHERE portfolios_portfolio.userid = $1 AND portfolios_portfolio.basecurrencycode = $2
		RETURNING id, userid, basecurrencycode;
	`

	sqlPortfolio := &SQLPortfolio{}
	if err := r.store.db.QueryRowContext(
		ctx,
		query,
		portfolioData.UserId,
		portfolioData.BaseCurrencyCode,
	).Scan(
		&sqlPortfolio.Id,
		&sqlPortfolio.UserId,
		&sqlPortfolio.BaseCurrencyCode,
	); err != nil {
		return domain.PortfolioRetrieve{Portfolio: nil, Error: domain.UnknownError}
	}
	portfolio := &domain.Portfolio{
		Id:               sqlPortfolio.Id,
		UserId:           domain.UserId(sqlPortfolio.UserId),
		BaseCurrencyCode: sqlPortfolio.BaseCurrencyCode,
	}
	return domain.PortfolioRetrieve{Portfolio: portfolio, Error: nil}
}
