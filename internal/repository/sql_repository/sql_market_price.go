package sql_repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"strings"
)

type SQLMarketPriceRepository struct {
	store *SQLRepository
}

type SQLMarketPrice struct {
	MarketItemId int64
	Price        float64
	Timestamp    sql.NullTime
}

const MarketPriceFieldsNum = 3

func (r SQLMarketPriceRepository) RetrieveLast(ctx context.Context, marketItemId int64) domain.MarketPriceRetrieve {
	const query = `
		SELECT 
    		market_item_id,
    		price,
    		ts
		FROM market_price
		WHERE market_item_id = $1 AND ts = (SELECT MAX(ts) from market_price WHERE market_item_id = $1)
	`

	sqlMarketPrice := &SQLMarketPrice{}
	if err := r.store.db.QueryRowContext(
		ctx,
		query,
		marketItemId,
	).Scan(
		&sqlMarketPrice.MarketItemId,
		&sqlMarketPrice.Price,
		&sqlMarketPrice.Timestamp,
	); err != nil {
		return domain.MarketPriceRetrieve{MarketPrice: nil, Error: domain.NotFoundError}
	}
	marketPrice := &domain.MarketPrice{
		MarketItemId: sqlMarketPrice.MarketItemId,
		Price:        sqlMarketPrice.Price,
		Timestamp:    sqlMarketPrice.Timestamp.Time.Unix(),
	}
	return domain.MarketPriceRetrieve{MarketPrice: marketPrice, Error: nil}
}

func (r SQLMarketPriceRepository) RetrieveInterval(
	ctx context.Context,
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
		ORDER BY ts
	`

	prices := make([]domain.MarketPrice, 0)
	rows, err := r.store.db.QueryContext(
		ctx,
		query,
		marketItemId,
		start,
		end,
	)
	if err != nil {
		return &domain.MarketPricesRetrieve{MarketPrices: nil, Error: domain.NotFoundError}
	}

	for rows.Next() {
		var sqlMarketPrice SQLMarketPrice

		if err := rows.Scan(
			&sqlMarketPrice.MarketItemId,
			&sqlMarketPrice.Price,
			&sqlMarketPrice.Timestamp,
		); err != nil {
			return &domain.MarketPricesRetrieve{MarketPrices: nil, Error: domain.UnknownError}
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

func (r SQLMarketPriceRepository) Create(ctx context.Context, marketPrice *domain.MarketPrice) (bool, error) {
	const query = `
		INSERT INTO market_price (
			market_item_id, 
		  	price,
		  	ts
	    ) VALUES (
			$1, $2, to_timestamp($3)
	  	)
	  	ON CONFLICT ON CONSTRAINT market_price_market_item_id_ts_key
		DO NOTHING;
	`

	_, err := r.store.db.ExecContext(
		ctx,
		query,
		marketPrice.MarketItemId,
		marketPrice.Price,
		marketPrice.Timestamp,
	)
	if err != nil {
		return false, domain.UnknownError
	}
	return true, nil
}

func (r SQLMarketPriceRepository) BulkCreate(ctx context.Context, marketPrices *[]domain.MarketPrice) (int64, error) {
	var placeholders []string
	var values []interface{}

	for _, marketPrice := range *marketPrices {
		placeholders = append(
			placeholders,
			fmt.Sprintf("($%d,$%d,to_timestamp($%d))",
				len(placeholders)*MarketPriceFieldsNum+1,
				len(placeholders)*MarketPriceFieldsNum+2,
				len(placeholders)*MarketPriceFieldsNum+3,
			),
		)
		values = append(values, marketPrice.MarketItemId, marketPrice.Price, marketPrice.Timestamp)
	}

	insertQuery := fmt.Sprintf(`
		INSERT INTO market_price (
			market_item_id,
			price,
			ts
		) VALUES %s
		ON CONFLICT ON CONSTRAINT market_price_market_item_id_ts_key DO NOTHING;
		`,
		strings.Join(placeholders, ","),
	)
	res, err := r.store.db.ExecContext(ctx, insertQuery, values...)
	if err != nil {
		return 0, errors.Wrap(err, "failed to insert multiple records at once")
	}
	rowsInserted, err := res.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "failed to insert multiple records at once")
	}

	return rowsInserted, nil
}
