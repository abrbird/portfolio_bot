package sql_repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/abrbird/portfolio_bot/internal/domain"
	"strings"
)

type SQLMarketItemRepository struct {
	store *SQLRepository
}

type SQLMarketItem struct {
	Id    int64
	Code  string
	Type  sql.NullString
	Title sql.NullString
}

func (r SQLMarketItemRepository) RetrieveById(ctx context.Context, marketItemId int64) domain.MarketItemRetrieve {
	const query = `
		SELECT 
		    id,
    		code,
    		type,
    		title
		FROM market_item
		WHERE id = $1
	`

	sqlMarketItem := &SQLMarketItem{}
	if err := r.store.db.QueryRowContext(
		ctx,
		query,
		marketItemId,
	).Scan(
		&sqlMarketItem.Id,
		&sqlMarketItem.Code,
		&sqlMarketItem.Type,
		&sqlMarketItem.Title,
	); err != nil {
		return domain.MarketItemRetrieve{MarketItem: nil, Error: domain.NotFoundError}
	}
	marketItem := &domain.MarketItem{
		Id:    sqlMarketItem.Id,
		Code:  sqlMarketItem.Code,
		Type:  sqlMarketItem.Type.String,
		Title: sqlMarketItem.Title.String,
	}
	return domain.MarketItemRetrieve{MarketItem: marketItem, Error: nil}
}

func (r SQLMarketItemRepository) Retrieve(ctx context.Context, code string, type_ string) domain.MarketItemRetrieve {
	const query = `
		SELECT 
		    id,
    		code,
    		type,
    		title
		FROM market_item
		WHERE code = $1 AND type = $2
	`

	sqlMarketItem := &SQLMarketItem{}
	if err := r.store.db.QueryRowContext(
		ctx,
		query,
		code,
		type_,
	).Scan(
		&sqlMarketItem.Id,
		&sqlMarketItem.Code,
		&sqlMarketItem.Type,
		&sqlMarketItem.Title,
	); err != nil {
		return domain.MarketItemRetrieve{MarketItem: nil, Error: domain.NotFoundError}
	}
	marketItem := &domain.MarketItem{
		Id:    sqlMarketItem.Id,
		Code:  sqlMarketItem.Code,
		Type:  sqlMarketItem.Type.String,
		Title: sqlMarketItem.Title.String,
	}
	return domain.MarketItemRetrieve{MarketItem: marketItem, Error: nil}
}

func (r SQLMarketItemRepository) RetrieveByType(ctx context.Context, codes []string, type_ string) *domain.MarketItemsRetrieve {

	var placeholders []string
	var values []interface{}

	for _, code := range codes {
		placeholders = append(
			placeholders,
			fmt.Sprintf("$%d", len(placeholders)+1),
		)
		values = append(values, code)
	}
	values = append(values, type_)

	query := fmt.Sprintf(`
		SELECT 
			id,
    		code,
    		type,
    		title
		FROM market_item
		WHERE code IN (%s) AND type = $%d`,
		strings.Join(placeholders, ","),
		len(placeholders)+1,
	)

	items := make([]domain.MarketItem, 0)
	rows, err := r.store.db.QueryContext(
		ctx,
		query,
		values...,
	)
	if err != nil {
		return &domain.MarketItemsRetrieve{MarketItems: nil, Error: domain.NotFoundError}
	}

	for rows.Next() {
		var sqlMarketItem SQLMarketItem
		if err := rows.Scan(
			&sqlMarketItem.Id,
			&sqlMarketItem.Code,
			&sqlMarketItem.Type,
			&sqlMarketItem.Title,
		); err != nil {
			return &domain.MarketItemsRetrieve{MarketItems: nil, Error: domain.NotFoundError}
		}
		items = append(
			items,
			domain.MarketItem{
				Id:    sqlMarketItem.Id,
				Code:  sqlMarketItem.Code,
				Type:  sqlMarketItem.Type.String,
				Title: sqlMarketItem.Title.String,
			})
	}

	return &domain.MarketItemsRetrieve{MarketItems: items, Error: err}
}
