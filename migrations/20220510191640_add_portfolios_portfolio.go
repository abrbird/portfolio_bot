package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddPortfoliosPortfolio, downAddPortfoliosPortfolio)
}

func upAddPortfoliosPortfolio(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE portfolios_portfolio (
		    ID bigserial PRIMARY KEY,
		    UserId bigint NOT NULL,
		    BaseCurrencyCode VARCHAR(8) NOT NULL,
			UNIQUE (UserId, BaseCurrencyCode),
		    CONSTRAINT portfolios_portfolio_fk_users_user 
			    FOREIGN KEY(UserId)
			        REFERENCES users_user(id) ON DELETE CASCADE,
			CONSTRAINT portfolios_portfolio_fk_currencies_currency 
			    FOREIGN KEY(BaseCurrencyCode)
			        REFERENCES currencies_currency(code) ON DELETE CASCADE
		);
	`)
	if err != nil {
		return err
	}

	return nil
}

func downAddPortfoliosPortfolio(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE portfolios_portfolio;")
	if err != nil {
		return err
	}
	return nil
}
