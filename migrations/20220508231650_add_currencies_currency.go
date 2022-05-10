package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddCurrencies, downAddCurrencies)
}

func upAddCurrencies(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE currencies_currency (
    		Code VARCHAR(8) PRIMARY KEY,
			Type VARCHAR(8),
			Title VARCHAR(128)
		);
	`)
	if err != nil {
		return err
	}

	return nil
}

func downAddCurrencies(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE currencies_currency;")
	if err != nil {
		return err
	}
	return nil
}
