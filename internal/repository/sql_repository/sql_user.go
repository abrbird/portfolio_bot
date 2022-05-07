package sql_repository

import (
	"database/sql"
	"errors"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
)

type SQLUserRepository struct {
	store *SQLRepository
}

type SQLUser struct {
	Id        sql.NullInt64
	UserName  sql.NullString
	FirstName sql.NullString
	LastName  sql.NullString
}

func (r *SQLUserRepository) Create(user *domain.User) error {
	userRetrieved := r.Retrieve(user.Id)
	if userRetrieved.User == nil {
		return r.store.db.QueryRow(
			"INSERT INTO users (id, username, firstname, lastname) VALUES ($1, $2, $3, $4) RETURNING id",
			user.Id,
			user.UserName,
			user.FirstName,
			user.LastName,
		).Scan(&user.Id)
	}
	return domain.ErrorAlreadyExists
}

func (r *SQLUserRepository) Retrieve(userId domain.UserId) domain.UserRetrieve {
	sqlUser := &SQLUser{}

	if err := r.store.db.QueryRow(
		"SELECT id, username, firstname, lastname FROM users WHERE id = $1",
		userId,
	).Scan(
		&sqlUser.Id,
		&sqlUser.UserName,
		&sqlUser.FirstName,
		&sqlUser.LastName,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.UserRetrieve{User: nil, Error: domain.ErrorNotFound}
		}
		return domain.UserRetrieve{User: nil, Error: err}
	}
	user := &domain.User{
		Id:        domain.UserId(sqlUser.Id.Int64),
		UserName:  sqlUser.UserName.String,
		FirstName: sqlUser.FirstName.String,
		LastName:  sqlUser.LastName.String,
	}

	return domain.UserRetrieve{User: user, Error: nil}
}

func (r *SQLUserRepository) Update(user *domain.User) error {
	userRetrieved := r.Retrieve(user.Id)
	if userRetrieved.Error != nil {
		return userRetrieved.Error
	}

	if userRetrieved.User != nil {
		_, err := r.store.db.Exec(
			"UPDATE users SET (username, firstname, lastname) = ($2, $3, $4) WHERE id = $1",
			user.Id,
			user.UserName,
			user.FirstName,
			user.LastName,
		)
		if err != nil {
			return err
		}
		return nil
	}
	return domain.UnknownError
}

func (r *SQLUserRepository) Delete(userId domain.UserId) error {
	if r, err := r.store.db.Exec("DELETE FROM users WHERE id = $1;", userId); err == nil {
		rows, err := r.RowsAffected()
		if err != nil {
			return err
		}
		if rows == 0 {
			return domain.ErrorNotFound
		}
		return nil
	} else {
		return err
	}
}
