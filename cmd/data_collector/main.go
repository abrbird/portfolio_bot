package main

import (
	"github.com/jasonlvhit/gocron"
	"gitlab.ozon.dev/zBlur/homework-2/config"
	"gitlab.ozon.dev/zBlur/homework-2/internal/data_collector"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository/sql_repository"
	"gitlab.ozon.dev/zBlur/homework-2/internal/server"
	"gitlab.ozon.dev/zBlur/homework-2/internal/service/service_impl"
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
