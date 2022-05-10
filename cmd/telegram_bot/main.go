package main

import (
	"gitlab.ozon.dev/zBlur/homework-2/config"
	"gitlab.ozon.dev/zBlur/homework-2/pkg/bot"
	"gitlab.ozon.dev/zBlur/homework-2/pkg/client/grpc_client"
	"log"
)

func main() {
	config_, err := config.ParseConfig("config/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	//appClient := http_client.New(config_.ClientAPIKeys.AnyClient, "http://0.0.0.0:8090/v1", 5*time.Second)
	appClient := grpc_client.New(config_.ClientAPIKeys.AnyClient, "localhost:8080")

	tgBot := bot.New(config_.ExternalAPIKeys.Telegram, appClient)
	log.Printf("Authorized on account %s \n", tgBot.GetSelf().UserName)

	updates := tgBot.GetUpdatesChan(60)
	for update := range updates {
		go tgBot.Handle(update)
	}
}
