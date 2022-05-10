package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/zBlur/homework-2/pkg/client"
	"log"
	"strings"
)

type TelegramHandler struct {
	apiKey string
	clnt   client.Client
	tgBot  *tgbotapi.BotAPI
}

func New(apiKey string, clnt client.Client) *TelegramHandler {
	tgBot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Fatal(err)
	}
	tgBot.Debug = true

	return &TelegramHandler{
		apiKey: apiKey,
		clnt:   clnt,
		tgBot:  tgBot,
	}
}

func (h *TelegramHandler) GetSelf() tgbotapi.User {
	return h.tgBot.Self
}

func (h *TelegramHandler) GetUpdatesChan(timeout int) tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeout
	return h.tgBot.GetUpdatesChan(u)
}

func (h *TelegramHandler) GetCommandFromMessage(message *tgbotapi.Message) (string, bool) {
	if strings.HasPrefix(message.Text, "/") {
		splitted := strings.Split(message.Text, " ")
		return splitted[0], true
	}
	return "", false
}

func (h *TelegramHandler) Send(chatId int64, text string) {
	msg := tgbotapi.NewMessage(chatId, text)
	_, err := h.tgBot.Send(msg)
	if err != nil {
		log.Println(err)
	}

}

func (h *TelegramHandler) Handle(update tgbotapi.Update) {
	if update.Message != nil && update.Message.Chat.IsPrivate() { // handle only private chats
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		user := GetUser(update.Message.Chat)
		userPB, err := h.clnt.GetOrCreateUser(user)

		if err != nil {
			h.SendError(update.Message.Chat.ID, err)
			return
		}

		if command, ok := h.GetCommandFromMessage(update.Message); ok {
			log.Println(command)

			if command == CommandStart {
				portfolioPB, err := h.clnt.GetOrCreatePortfolio(userPB.GetId())
				if err != nil {
					h.SendError(update.Message.Chat.ID, err)
					return
				}
				log.Println(portfolioPB)
				h.Send(update.Message.Chat.ID, fmt.Sprintf("Hi, %s!", GetUserName(userPB)))
			}
		}
	}
}

func (h *TelegramHandler) SendError(chatId int64, err error) {
	h.Send(chatId, "O-ops, something went wrong")
	log.Println(err)
}
