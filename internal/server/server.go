package server

import (
	"database/sql"
	"github.com/abrbird/portfolio_bot/config"
	"github.com/abrbird/portfolio_bot/internal/repository"
	"github.com/abrbird/portfolio_bot/internal/service"
	"github.com/abrbird/portfolio_bot/pkg/api"
	_ "github.com/lib/pq"
)

//var (
//	ErrorTimeOut = errors.New("timeout error")
//)

type tserver struct {
	conf config.Config
	repo repository.Repository
	serv service.Service
	api.UnimplementedUserPortfolioServiceServer
}

func NewServer(conf config.Config, repo repository.Repository, serv service.Service) *tserver {
	return &tserver{
		conf: conf,
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
