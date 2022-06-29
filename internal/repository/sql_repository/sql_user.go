package sql_repository

import (
	"context"
	"database/sql"
	"github.com/abrbird/portfolio_bot/internal/domain"
)

type SQLUserRepository struct {
	store *SQLRepository
}

type SQLUser struct {
	Id        int64
	UserName  sql.NullString
	FirstName sql.NullString
	LastName  sql.NullString
}

func (r *SQLUserRepository) Create(ctx context.Context, user *domain.User) error {

	const query = `
		INSERT INTO users_user (
			id,
	  		username,
			firstname,
		   	lastname
	    ) VALUES (
			$1, $2, $3, $4
	  	)
	  	RETURNING id
	`

	err := r.store.db.QueryRowContext(
		ctx,
		query,
		user.Id,
		user.UserName,
		user.FirstName,
		user.LastName,
	).Scan(&user.Id)
	if err != nil {
		return domain.UnknownError
	}
	return nil
}

func (r *SQLUserRepository) Retrieve(ctx context.Context, userId domain.UserId) domain.UserRetrieve {
	const query = `
		SELECT 
    		id,
    		username,
    		firstname,
    		lastname
		FROM users_user
		WHERE id = $1
	`

	sqlUser := &SQLUser{}
	if err := r.store.db.QueryRowContext(
		ctx,
		query,
		userId,
	).Scan(
		&sqlUser.Id,
		&sqlUser.UserName,
		&sqlUser.FirstName,
		&sqlUser.LastName,
	); err != nil {
		return domain.UserRetrieve{User: nil, Error: domain.NotFoundError}
	}
	user := &domain.User{
		Id:        domain.UserId(sqlUser.Id),
		UserName:  sqlUser.UserName.String,
		FirstName: sqlUser.FirstName.String,
		LastName:  sqlUser.LastName.String,
	}
	return domain.UserRetrieve{User: user, Error: nil}
}

func (r *SQLUserRepository) RetrieveOrCreate(ctx context.Context, user *domain.User) domain.UserRetrieve {
	const query = `
		INSERT INTO users_user (
			id,
	  		username,
			firstname,
		   	lastname
	    ) VALUES (
			$1, $2, $3, $4
	  	) 
	  	ON CONFLICT ON CONSTRAINT users_user_pkey
		DO UPDATE SET (
			username,
			firstname,
			lastname
		) = (
			$2, $3, $4
		) WHERE users_user.id = $1
		RETURNING id, username, firstname, lastname
	`

	sqlUser := &SQLUser{}
	if err := r.store.db.QueryRowContext(
		ctx,
		query,
		user.Id,
		user.UserName,
		user.FirstName,
		user.LastName,
	).Scan(
		&sqlUser.Id,
		&sqlUser.UserName,
		&sqlUser.FirstName,
		&sqlUser.LastName,
	); err != nil {
		return domain.UserRetrieve{User: nil, Error: domain.UnknownError}
	}
	userR := &domain.User{
		Id:        domain.UserId(sqlUser.Id),
		UserName:  sqlUser.UserName.String,
		FirstName: sqlUser.FirstName.String,
		LastName:  sqlUser.LastName.String,
	}
	return domain.UserRetrieve{User: userR, Error: nil}
}

func (r *SQLUserRepository) Update(ctx context.Context, user *domain.User) error {
	const query = `
		UPDATE users_user SET (
			username,
		    firstname,
			lastname
		) = (
			$2, $3, $4
		)
		WHERE id = $1
	`

	res, err := r.store.db.ExecContext(
		ctx,
		query,
		user.Id,
		user.UserName,
		user.FirstName,
		user.LastName,
	)
	if err != nil {
		return domain.NotFoundError
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return domain.NotFoundError
	}
	if rows == 0 {
		return domain.NotFoundError
	}
	return nil
}

func (r *SQLUserRepository) Delete(ctx context.Context, userId domain.UserId) error {
	const query = `
		DELETE FROM users_user 
	    WHERE id = $1
	`

	res, err := r.store.db.ExecContext(
		ctx,
		query,
		userId,
	)
	if err != nil {
		return domain.NotFoundError
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return domain.NotFoundError
	}
	if rows == 0 {
		return domain.NotFoundError
	}
	return nil
}
