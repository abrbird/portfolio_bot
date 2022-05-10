package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upMarketPrice, downMarketPrice)
}

func upMarketPrice(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE market_price (
    		market_item_id integer NOT NULL,
			price DECIMAL NOT NULL,
			ts TIMESTAMP NOT NULL,
			CONSTRAINT market_price_fk_market_item 
			    FOREIGN KEY(market_item_id)
			        REFERENCES market_item(id) ON DELETE CASCADE,
			UNIQUE (market_item_id, ts)
		);
	`)
	if err != nil {
		return err
	}

	return nil
}

func downMarketPrice(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE market_price;")
	if err != nil {
		return err
	}

	return nil
}
