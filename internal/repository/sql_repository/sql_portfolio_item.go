package sql_repository

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
)

//portfolioRepository     *SQLPortfolioRepository
//portfolioItemRepository *

type SQLPortfolioItemRepository struct {
	store *SQLRepository
}

type SQLPortfolioItem struct {
	Id           int64
	PortfolioId  int64
	MarketItemId int64
	Price        float64
	Volume       float64
}

func (r SQLPortfolioItemRepository) RetrieveOrCreate(ctx context.Context, portfolioData domain.PortfolioItemCreate) domain.PortfolioItemRetrieve {
	const query = `
		INSERT INTO portfolios_item (
			portfolioid,
			marketitemid,
			price,
		 	volume
		) VALUES (
			$1, $2, $3, $4
		)
		ON CONFLICT ON CONSTRAINT portfolios_item_portfolioid_marketitemid_key
		DO UPDATE SET (
			portfolioid,
			marketitemid,
			price,
		    volume
		) = (
			$1, $2, $3, $4
		) WHERE portfolios_item.portfolioid = $1 AND portfolios_item.marketitemid = $2
		RETURNING id, portfolioid, marketitemid, price, volume;
	`

	sqlPortfolioItem := &SQLPortfolioItem{}
	if err := r.store.db.QueryRowContext(
		ctx,
		query,
		portfolioData.PortfolioId,
		portfolioData.MarketItemId,
		portfolioData.Price,
		portfolioData.Volume,
	).Scan(
		&sqlPortfolioItem.Id,
		&sqlPortfolioItem.PortfolioId,
		&sqlPortfolioItem.MarketItemId,
		&sqlPortfolioItem.Price,
		&sqlPortfolioItem.Volume,
	); err != nil {
		return domain.PortfolioItemRetrieve{PortfolioItem: nil, Error: domain.UnknownError}
	}
	portfolioItem := &domain.PortfolioItem{
		Id:           sqlPortfolioItem.Id,
		PortfolioId:  sqlPortfolioItem.PortfolioId,
		MarketItemId: sqlPortfolioItem.MarketItemId,
		Price:        sqlPortfolioItem.Price,
		Volume:       sqlPortfolioItem.Volume,
	}
	return domain.PortfolioItemRetrieve{PortfolioItem: portfolioItem, Error: nil}
}

func (r SQLPortfolioItemRepository) RetrievePortfolioItems(ctx context.Context, portfolioId int64) *domain.PortfolioItemsRetrieve {
	const query = `
		SELECT 
    		id,
    		portfolioid,
    		marketitemid,
    		price,
    		volume
		FROM portfolios_item
		WHERE portfolioId = $1
	`

	portfolioItem := make([]domain.PortfolioItem, 0)
	rows, err := r.store.db.QueryContext(
		ctx,
		query,
		portfolioId,
	)
	if err != nil {
		return &domain.PortfolioItemsRetrieve{PortfolioItems: nil, Error: domain.NotFoundError}
	}

	for rows.Next() {
		var sqlPortfolioItem SQLPortfolioItem

		if err := rows.Scan(
			&sqlPortfolioItem.Id,
			&sqlPortfolioItem.PortfolioId,
			&sqlPortfolioItem.MarketItemId,
			&sqlPortfolioItem.Price,
			&sqlPortfolioItem.Volume,
		); err != nil {
			return &domain.PortfolioItemsRetrieve{PortfolioItems: nil, Error: domain.UnknownError}
		}
		portfolioItem = append(
			portfolioItem,
			domain.PortfolioItem{
				Id:           sqlPortfolioItem.Id,
				PortfolioId:  sqlPortfolioItem.PortfolioId,
				MarketItemId: sqlPortfolioItem.MarketItemId,
				Price:        sqlPortfolioItem.Price,
				Volume:       sqlPortfolioItem.Volume,
			})
	}

	return &domain.PortfolioItemsRetrieve{PortfolioItems: portfolioItem, Error: nil}
}

func (r SQLPortfolioItemRepository) Delete(ctx context.Context, portfolioItemId int64) error {
	const query = `
		DELETE FROM portfolios_item
		WHERE id = $1
	`

	res, err := r.store.db.ExecContext(
		ctx,
		query,
		portfolioItemId,
	)
	if err != nil {
		return domain.NotFoundError
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return domain.NotFoundError
	}
	if rows == 0 {
		return domain.NotFoundError
	}
	return nil
}
