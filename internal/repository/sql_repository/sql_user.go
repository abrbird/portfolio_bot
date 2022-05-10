package sql_repository

import (
	"database/sql"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
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

func (r *SQLUserRepository) Create(user *domain.User) error {

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

	return r.store.db.QueryRow(
		query,
		user.Id,
		user.UserName,
		user.FirstName,
		user.LastName,
	).Scan(&user.Id)
}

func (r *SQLUserRepository) Retrieve(userId domain.UserId) domain.UserRetrieve {
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
	if err := r.store.db.QueryRow(
		query,
		userId,
	).Scan(
		&sqlUser.Id,
		&sqlUser.UserName,
		&sqlUser.FirstName,
		&sqlUser.LastName,
	); err != nil {
		return domain.UserRetrieve{User: nil, Error: err}
	}
	user := &domain.User{
		Id:        domain.UserId(sqlUser.Id),
		UserName:  sqlUser.UserName.String,
		FirstName: sqlUser.FirstName.String,
		LastName:  sqlUser.LastName.String,
	}
	return domain.UserRetrieve{User: user, Error: nil}
}

func (r *SQLUserRepository) RetrieveOrCreate(user *domain.User) domain.UserRetrieve {
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
	if err := r.store.db.QueryRow(
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
		return domain.UserRetrieve{User: nil, Error: err}
	}
	userR := &domain.User{
		Id:        domain.UserId(sqlUser.Id),
		UserName:  sqlUser.UserName.String,
		FirstName: sqlUser.FirstName.String,
		LastName:  sqlUser.LastName.String,
	}
	return domain.UserRetrieve{User: userR, Error: nil}
}

func (r *SQLUserRepository) Update(user *domain.User) error {
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

	err := r.store.db.QueryRow(
		query,
		user.Id,
		user.UserName,
		user.FirstName,
		user.LastName,
	).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *SQLUserRepository) Delete(userId domain.UserId) error {
	const query = `
		DELETE FROM users_user 
	    WHERE id = $1
	`

	err := r.store.db.QueryRow(
		query,
		userId,
	).Err()
	if err != nil {
		return err
	}
	return nil
}
