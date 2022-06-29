package main

import (
	"github.com/abrbird/portfolio_bot/config"
	"github.com/abrbird/portfolio_bot/internal/data_collector"
	"github.com/abrbird/portfolio_bot/internal/repository/sql_repository"
	"github.com/abrbird/portfolio_bot/internal/server"
	"github.com/abrbird/portfolio_bot/internal/service/service_impl"
	"github.com/jasonlvhit/gocron"
	"log"
)

func main() {

	config_, err := config.ParseConfig("config/config.yml")
	if err != nil {
		log.Fatal(err)
	}
	db, err := server.NewDB(config_.Database.Uri())
	if err != nil {
		log.Fatalf("db connection failed: %v", err)
	}
	defer func() {
		err = db.Close()
		if err != nil {
			log.Fatalf("db connection failed: %v", err)
		}
	}()
	serv := service_impl.New()
	repo := sql_repository.New(db)

	err = gocron.Every(config_.Application.HistoryInterval).Seconds().From(gocron.NextTick()).Do(data_collector.Collect, config_, serv, repo)
	if err != nil {
		log.Fatal(err)
	}

	<-gocron.Start()
}
