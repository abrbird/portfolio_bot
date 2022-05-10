package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddMarketItem, downAddMarketItem)
}

func upAddMarketItem(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE market_item (
		    ID serial PRIMARY KEY,
    		Code VARCHAR(16) NOT NULL,
			Type VARCHAR(16) NOT NULL,
			Title VARCHAR(256),
			UNIQUE (Code, Type)
		);
	`)
	if err != nil {
		return err
	}

	return nil
}

func downAddMarketItem(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE market_item;")
	if err != nil {
		return err
	}

	return nil
}
