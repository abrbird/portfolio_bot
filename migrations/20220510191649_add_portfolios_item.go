package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddPortfoliosItems, downAddPortfoliosItems)
}

func upAddPortfoliosItems(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE portfolios_item (
		    ID bigserial PRIMARY KEY,
		    PortfolioID bigint NOT NULL,
		    MarketItemId integer NOT NULL,
		    price DECIMAL NOT NULL,
		    volume DECIMAL NOT NULL,
		    UNIQUE (PortfolioID, MarketItemId),
		    CONSTRAINT portfolios_item_fk_portfolios_portfolio 
			    FOREIGN KEY(PortfolioID)
			        REFERENCES portfolios_portfolio(id) ON DELETE CASCADE,
			CONSTRAINT portfolios_item_fk_market_item 
			    FOREIGN KEY(MarketItemId)
			        REFERENCES market_item(id) ON DELETE CASCADE
		);
	`)
	if err != nil {
		return err
	}
	return nil
}

func downAddPortfoliosItems(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE portfolios_item;")
	if err != nil {
		return err
	}
	return nil
}
