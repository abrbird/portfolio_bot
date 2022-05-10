package sql_repository

import (
	"database/sql"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
)

type SQLCurrencyRepository struct {
	store *SQLRepository
}

type SQLCurrency struct {
	Code  string
	Type  sql.NullString
	Title sql.NullString
}

func (r SQLCurrencyRepository) Retrieve(currencyCode string) domain.CurrencyRetrieve {
	const query = `
		SELECT 
    		code,
    		type,
    		title
		FROM currencies_currency
		WHERE code = $1
	`

	sqlCurrency := &SQLCurrency{}
	if err := r.store.db.QueryRow(
		query,
		currencyCode,
	).Scan(
		&sqlCurrency.Code,
		&sqlCurrency.Type,
		&sqlCurrency.Title,
	); err != nil {
		return domain.CurrencyRetrieve{Currency: nil, Error: err}
	}
	currency := &domain.Currency{
		Code:  sqlCurrency.Code,
		Type:  sqlCurrency.Type.String,
		Title: sqlCurrency.Title.String,
	}
	return domain.CurrencyRetrieve{Currency: currency, Error: nil}
}
