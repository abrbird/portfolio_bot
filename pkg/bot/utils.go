package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/pkg/api"
	"strings"
)

const (
	CommandStart = "/start"
	CommandHelp  = "/help"
)

func GetUserName(user *api.User) string {
	userName := "anonymous"
	if len(user.FirstName) > 0 || len(user.LastName) > 0 {
		userName = fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	} else if len(user.UserName) > 0 {
		userName = user.UserName
	}
	userName = strings.TrimSpace(userName)

	return userName
}

func GetUser(message *tgbotapi.Chat) *domain.User {
	user := &domain.User{
		Id:        domain.UserId(message.ID),
		UserName:  message.UserName,
		FirstName: message.FirstName,
		LastName:  message.LastName,
	}
	return user
}
