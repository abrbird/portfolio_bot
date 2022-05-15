package main

import (
	"fmt"
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

	//appClient := http_client.New(
	//	config_.ClientAPIKeys.AnyClient,
	//	fmt.Sprintf("http://%s:8090/v1", config_.Application.Host),
	//	5*time.Second,
	//)
	//appClient := grpc_client.New(config_.ClientAPIKeys.AnyClient, fmt.Sprintf("%s:8080", "0.0.0.0"))
	appClient := grpc_client.New(config_.ClientAPIKeys.AnyClient, fmt.Sprintf("%s:8080", config_.Application.Host))

	tgBot := bot.New(config_.ExternalAPIKeys.Telegram, appClient, false)
	log.Printf("Authorized on account %s \n", tgBot.GetSelf().UserName)

	updates := tgBot.GetUpdatesChan(60)
	for update := range updates {
		go tgBot.Handle(update)
	}
}
