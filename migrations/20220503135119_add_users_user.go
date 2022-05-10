package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddUser, downAddUser)
}

func upAddUser(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE users_user (
		    ID bigint PRIMARY KEY,
		 	UserName VARCHAR(32),
		 	FirstName VARCHAR(64),
		 	LastName VARCHAR(64)
		);
	`)
	if err != nil {
		return err
	}
	return nil
}

func downAddUser(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE users_user;")
	if err != nil {
		return err
	}
	return nil
}
