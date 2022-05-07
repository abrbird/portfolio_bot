package server

import (
	"database/sql"
	_ "github.com/lib/pq"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository"
	"gitlab.ozon.dev/zBlur/homework-2/internal/service"
	"gitlab.ozon.dev/zBlur/homework-2/pkg/api"
)

type tserver struct {
	repo repository.Repository
	serv service.Service
	api.UnimplementedUserServiceServer
}

func NewServer(repo repository.Repository, serv service.Service) *tserver {
	return &tserver{
		repo: repo,
		serv: serv,
	}
}

func NewDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
