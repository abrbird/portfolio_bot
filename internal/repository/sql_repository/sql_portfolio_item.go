package sql_repository

import "gitlab.ozon.dev/zBlur/homework-2/internal/domain"

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

func (r SQLPortfolioItemRepository) RetrieveOrCreate(portfolioData domain.PortfolioItemCreate) domain.PortfolioItemRetrieve {

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
			marketitemid
		) = (
			$1, $2
		) WHERE portfolios_item.portfolioid = $1 AND portfolios_item.marketitemid = $2
		RETURNING id, portfolioid, marketitemid, price, volume;
	`

	sqlPortfolioItem := &SQLPortfolioItem{}
	if err := r.store.db.QueryRow(
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
		return domain.PortfolioItemRetrieve{PortfolioItem: nil, Error: err}
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

func (r SQLPortfolioItemRepository) RetrievePortfolioItems(portfolioId int64) *domain.PortfolioItemsRetrieve {
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
	rows, err := r.store.db.Query(
		query,
		portfolioId,
	)
	if err != nil {
		return &domain.PortfolioItemsRetrieve{PortfolioItems: nil, Error: err}
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
			return &domain.PortfolioItemsRetrieve{PortfolioItems: nil, Error: err}
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

	return &domain.PortfolioItemsRetrieve{PortfolioItems: portfolioItem, Error: err}
}

func (r SQLPortfolioItemRepository) Delete(portfolioItemId int64) error {
	const query = `
		DELETE FROM portfolios_item
		WHERE id = $1
	`

	err := r.store.db.QueryRow(
		query,
		portfolioItemId,
	).Err()
	if err != nil {
		return err
	}
	return nil
}
