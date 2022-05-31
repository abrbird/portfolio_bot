package bot

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"strings"
)

var (
	TimeoutError = errors.New("timeout error")
)

var (
	ErrorMessage            = "O-ops, something went wrong. Please try again later"
	UnknownCommandMessage   = fmt.Sprintf("Unknown command. May be this will help you /%s", CommandHelp)
	NoSufficientDataMessage = "No sufficient data"
	ParsingErrorMessage     = "Can not parse text"
)

var (
	PortfolioItemCreated = "Added: %s"
	PortfolioItemUpdated = "Updated: %s"
	PortfolioItemDeleted = "Deleted: %s"
)

const (
	CommandStart                    = "start"
	CommandHelp                     = "help"
	CommandMarket                   = "market"
	CommandPortfolio                = "portfolio"
	CommandPortfolioSummary         = "portfolio_summary"
	CommandAddOrUpdatePortfolioItem = "add_or_update_portfolio_item"
	CommandDeletePortfolioItem      = "delete_portfolio_item"
)

var HelpDescription = strings.Join(
	[]string{
		"Available commands:", "\n",
		"/", CommandMarket, "\n",
		"/", CommandPortfolio, "\n",
		"/", CommandPortfolioSummary, "\n",
		"/", CommandAddOrUpdatePortfolioItem, "\n",
		"/", CommandDeletePortfolioItem, "\n",
	},
	"",
)

var CommandDescriptions = map[string]string{
	CommandStart:                    "Hi, %s!",
	CommandHelp:                     HelpDescription,
	CommandMarket:                   `Available market symbols: %s. Choose one of them by typing it's name`,
	CommandPortfolio:                "Portfolio (%s): \n%s",
	CommandPortfolioSummary:         "Portfolio summary (%s): \n%s",
	CommandAddOrUpdatePortfolioItem: "Specify item's symbol, price and volume, e.g.\n'BTC 54321.09 0.1'",
	CommandDeletePortfolioItem:      "Choose item's id: \n%s",
}

func GetUserFromChat(message *tgbotapi.Chat) *domain.User {
	user := &domain.User{
		Id:        domain.UserId(message.ID),
		UserName:  message.UserName,
		FirstName: message.FirstName,
		LastName:  message.LastName,
	}
	return user
}
