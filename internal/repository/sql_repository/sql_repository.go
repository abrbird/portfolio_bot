package sql_repository

import (
	"database/sql"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository"
)

type SQLRepository struct {
	db             *sql.DB
	userRepository *SQLUserRepository
}

func (s *SQLRepository) User() repository.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &SQLUserRepository{store: s}

	return s.userRepository
}

func New(db *sql.DB) *SQLRepository {
	return &SQLRepository{
		db: db,
	}
}
