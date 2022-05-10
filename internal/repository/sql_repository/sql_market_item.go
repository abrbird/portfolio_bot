package sql_repository

import (
	"database/sql"
	"fmt"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
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

func (r SQLMarketItemRepository) RetrieveById(marketItemId int64) domain.MarketItemRetrieve {
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
	if err := r.store.db.QueryRow(
		query,
		marketItemId,
	).Scan(
		&sqlMarketItem.Id,
		&sqlMarketItem.Code,
		&sqlMarketItem.Type,
		&sqlMarketItem.Title,
	); err != nil {
		return domain.MarketItemRetrieve{MarketItem: nil, Error: err}
	}
	marketItem := &domain.MarketItem{
		Id:    sqlMarketItem.Id,
		Code:  sqlMarketItem.Code,
		Type:  sqlMarketItem.Type.String,
		Title: sqlMarketItem.Title.String,
	}
	return domain.MarketItemRetrieve{MarketItem: marketItem, Error: nil}
}

func (r SQLMarketItemRepository) Retrieve(code string, type_ string) domain.MarketItemRetrieve {
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
	if err := r.store.db.QueryRow(
		query,
		code,
		type_,
	).Scan(
		&sqlMarketItem.Id,
		&sqlMarketItem.Code,
		&sqlMarketItem.Type,
		&sqlMarketItem.Title,
	); err != nil {
		return domain.MarketItemRetrieve{MarketItem: nil, Error: err}
	}
	marketItem := &domain.MarketItem{
		Id:    sqlMarketItem.Id,
		Code:  sqlMarketItem.Code,
		Type:  sqlMarketItem.Type.String,
		Title: sqlMarketItem.Title.String,
	}
	return domain.MarketItemRetrieve{MarketItem: marketItem, Error: nil}
}

func (r SQLMarketItemRepository) RetrieveByType(codes []string, type_ string) *domain.MarketItemsRetrieve {

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
	rows, err := r.store.db.Query(
		query,
		values...,
	)
	if err != nil {
		return &domain.MarketItemsRetrieve{MarketItems: nil, Error: err}
	}

	for rows.Next() {
		var sqlMarketItem SQLMarketItem

		if err := rows.Scan(
			&sqlMarketItem.Id,
			&sqlMarketItem.Code,
			&sqlMarketItem.Type,
			&sqlMarketItem.Title,
		); err != nil {
			return &domain.MarketItemsRetrieve{MarketItems: nil, Error: err}
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
