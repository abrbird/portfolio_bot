package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddUser, downAddUser)
}

func upAddUser(tx *sql.Tx) error {
	_, err := tx.Exec("CREATE TABLE users (ID bigint PRIMARY KEY, UserName VARCHAR, FirstName VARCHAR, LastName VARCHAR);")
	if err != nil {
		return err
	}
	return nil
}

func downAddUser(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE users;")
	if err != nil {
		return err
	}
	return nil
}
