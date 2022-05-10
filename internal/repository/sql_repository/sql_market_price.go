package sql_repository

import (
	"database/sql"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
)

type SQLMarketPriceRepository struct {
	store *SQLRepository
}

type SQLMarketPrice struct {
	MarketItemId int64
	Price        float64
	Timestamp    sql.NullTime
}

func (r SQLMarketPriceRepository) RetrieveLast(marketItemId int64) domain.MarketPriceRetrieve {
	const query = `
		SELECT 
    		market_item_id,
    		price,
    		ts
		FROM market_price
		WHERE market_item_id = $1 AND ts = (SELECT MAX(ts) from market_price WHERE market_item_id = $1)
	`

	sqlMarketPrice := &SQLMarketPrice{}
	if err := r.store.db.QueryRow(
		query,
		marketItemId,
	).Scan(
		&sqlMarketPrice.MarketItemId,
		&sqlMarketPrice.Price,
		&sqlMarketPrice.Timestamp,
	); err != nil {
		return domain.MarketPriceRetrieve{MarketPrice: nil, Error: err}
	}
	marketPrice := &domain.MarketPrice{
		MarketItemId: sqlMarketPrice.MarketItemId,
		Price:        sqlMarketPrice.Price,
		Timestamp:    sqlMarketPrice.Timestamp.Time.Unix(),
	}
	return domain.MarketPriceRetrieve{MarketPrice: marketPrice, Error: nil}
}

func (r SQLMarketPriceRepository) RetrieveInterval(
	marketItemId int64,
	start int64,
	end int64,
) *domain.MarketPricesRetrieve {

	const query = `
		SELECT 
    		market_item_id,
    		price,
    		ts
		FROM market_price
		WHERE market_item_id = $1 AND ts >= to_timestamp($2) AND ts <= to_timestamp($3)
	`

	prices := make([]domain.MarketPrice, 0)
	rows, err := r.store.db.Query(
		query,
		marketItemId,
		start,
		end,
	)
	if err != nil {
		return &domain.MarketPricesRetrieve{MarketPrices: nil, Error: err}
	}

	for rows.Next() {
		var sqlMarketPrice SQLMarketPrice

		if err := rows.Scan(
			&sqlMarketPrice.MarketItemId,
			&sqlMarketPrice.Price,
			&sqlMarketPrice.Timestamp,
		); err != nil {
			return &domain.MarketPricesRetrieve{MarketPrices: nil, Error: err}
		}
		prices = append(
			prices,
			domain.MarketPrice{
				MarketItemId: sqlMarketPrice.MarketItemId,
				Price:        sqlMarketPrice.Price,
				Timestamp:    sqlMarketPrice.Timestamp.Time.Unix(),
			})
	}

	return &domain.MarketPricesRetrieve{MarketPrices: prices, Error: err}
}

func (r SQLMarketPriceRepository) Create(marketPrice *domain.MarketPrice) error {
	const query = `
		INSERT INTO market_price (
			market_item_id, 
		  	price,
		  	ts
	    ) VALUES (
			$1, $2, to_timestamp($3)
	  	)
	  	ON CONFLICT ON CONSTRAINT market_price_market_item_id_ts_key
		DO UPDATE SET (
			market_item_id,
			ts
		) = (
			$1, to_timestamp($3)
		) WHERE market_price.market_item_id = $1 AND market_price.ts = to_timestamp($3)
		RETURNING market_item_id, price, ts;
	`

	var sqlMarketPrice SQLMarketPrice
	return r.store.db.QueryRow(
		query,
		marketPrice.MarketItemId,
		marketPrice.Price,
		marketPrice.Timestamp,
	).Scan(
		&sqlMarketPrice.MarketItemId,
		&sqlMarketPrice.Price,
		&sqlMarketPrice.Timestamp,
	)
}
