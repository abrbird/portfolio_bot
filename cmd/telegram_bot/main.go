package main

import (
	"fmt"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/pkg/client"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"gitlab.ozon.dev/zBlur/homework-2/config"
)

func main() {
	configFile, err := config.ParseConfig("config/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	telegramApiKey := configFile.ExternalAPIKeys.Telegram
	fmt.Println("telegramAPiKey", telegramApiKey)

	tgBot, err := tgbotapi.NewBotAPI(telegramApiKey)
	if err != nil {
		log.Fatal(err)
	}
	tgBot.Debug = true
	log.Printf("Authorized on account %s", tgBot.Self.UserName)

	botClient := client.New()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tgBot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && update.Message.Chat.IsPrivate() { // handle only private chats
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			if update.Message.Text == "/start" {
				log.Println("started")
			}

			user := &domain.User{
				Id:        domain.UserId(update.Message.Chat.ID),
				UserName:  update.Message.Chat.UserName,
				FirstName: update.Message.Chat.FirstName,
				LastName:  update.Message.Chat.LastName,
			}

			user, err = botClient.GetOrCreateUser(user)
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "O-ops, something went wrong")
				_, _ = tgBot.Send(msg)
				log.Println(err)
			} else {
				userName := "anonymous"
				if len(user.FirstName) > 0 || len(user.LastName) > 0 {
					userName = fmt.Sprintf("%s %s", user.FirstName, user.LastName)
					userName = strings.TrimSpace(userName)
				} else if len(user.UserName) > 0 {
					userName = user.UserName
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Hi, %s! %s", userName, update.Message.Text))
				_, _ = tgBot.Send(msg)
			}
		}
	}
}
