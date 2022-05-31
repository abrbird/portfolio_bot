package bot

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/pkg/api"
	"gitlab.ozon.dev/zBlur/homework-2/pkg/bot/cache"
	"gitlab.ozon.dev/zBlur/homework-2/pkg/client"
	"gitlab.ozon.dev/zBlur/homework-2/pkg/graph"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

type TelegramHandler struct {
	apiKey    string
	clnt      client.Client
	tgBot     *tgbotapi.BotAPI
	plt       graph.Drawer
	userCache *cache.Cache
}

func New(apiKey string, clnt client.Client, debug bool) *TelegramHandler {
	tgBot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Fatal(err)
	}
	tgBot.Debug = debug

	return &TelegramHandler{
		apiKey:    apiKey,
		clnt:      clnt,
		tgBot:     tgBot,
		plt:       graph.New(),
		userCache: cache.New(60 * 5),
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
	return message.Command(), message.IsCommand()
}

func (h *TelegramHandler) Send(chatId int64, text string) {
	msg := tgbotapi.NewMessage(chatId, text)
	_, err := h.tgBot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func (h *TelegramHandler) SendFile(chatId int64, file tgbotapi.RequestFileData) {
	img := tgbotapi.NewInputMediaPhoto(file)
	mediaGroup := tgbotapi.NewMediaGroup(chatId, []interface{}{
		img,
	})

	_, err := h.tgBot.SendMediaGroup(mediaGroup)
	if err != nil {
		log.Println(err)
	}
}

func (h *TelegramHandler) Handle(update tgbotapi.Update) {
	// TODO: refactor this monster!!!

	if update.Message != nil && update.Message.Chat.IsPrivate() { // handle only private chats
		if h.tgBot.Debug {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		}

		user := GetUserFromChat(update.Message.Chat)
		userCache, err := h.GetUserCached(user.Id)
		if err != nil {
			userPB, err := h.clnt.GetOrCreateUser(user)
			if err != nil {
				h.Send(update.Message.Chat.ID, ErrorMessage)
				log.Println(err)
				return
			}
			userCache = &cache.UserCache{
				User:    userPB,
				Ts:      time.Now().UTC(),
				Command: CommandStart,
			}
			_, err = h.clnt.GetOrCreatePortfolio(userCache.User.GetId())
			if err != nil {
				h.Send(update.Message.Chat.ID, ErrorMessage)
				log.Println(err)
				return
			}
		}

		if command, ok := h.GetCommandFromMessage(update.Message); ok {

			switch command {
			case CommandStart:
				h.Send(
					update.Message.Chat.ID,
					fmt.Sprintf(
						CommandDescriptions[command],
						userCache.GetName(),
					),
				)
				userCache.Command = command
			case CommandHelp:
				h.Send(
					update.Message.Chat.ID,
					CommandDescriptions[command],
				)
				userCache.Command = command
			case CommandMarket:
				availableMarketItems, err := h.clnt.GetAvailableMarketItems()
				if err != nil {
					h.Send(update.Message.Chat.ID, ErrorMessage)
					log.Println(err)
					return
				}
				if len(availableMarketItems) == 0 {
					h.Send(
						update.Message.Chat.ID,
						NoSufficientDataMessage,
					)
					return
				}
				symbols := make([]string, len(availableMarketItems))
				for i, mi := range availableMarketItems {
					symbols[i] = mi.GetCode()
				}
				h.Send(
					update.Message.Chat.ID,
					fmt.Sprintf(
						CommandDescriptions[command],
						strings.Join(symbols, ", "),
					),
				)
				userCache.Command = command

			case CommandPortfolio:
				availableMarketItems, err := h.clnt.GetAvailableMarketItems()
				if err != nil {
					h.Send(update.Message.Chat.ID, ErrorMessage)
					log.Println(err)
					return
				}
				availableMarketItemsMap := make(map[int64]*api.MarketItem, 0)
				for _, mi := range availableMarketItems {
					availableMarketItemsMap[mi.GetId()] = mi
				}

				portfolioPB, err := h.clnt.GetOrCreatePortfolio(userCache.User.GetId())
				if err != nil {
					h.Send(update.Message.Chat.ID, ErrorMessage)
					log.Println(err)
					return
				}
				if len(portfolioPB.GetItems()) == 0 {
					h.Send(
						update.Message.Chat.ID,
						NoSufficientDataMessage,
					)
					return
				}

				portfolioItemsInfo := make([]string, 0)
				portfolioItemsInfo = append(
					portfolioItemsInfo,
					"#. symbol, price * volume = cost",
				)
				cost := float64(0)
				for i, pi := range portfolioPB.GetItems() {
					c := pi.GetPrice() * pi.GetVolume()
					cost += c
					portfolioItemsInfo = append(
						portfolioItemsInfo,
						fmt.Sprintf(
							"%d. %s, %.4f * %.4f = %.2f",
							i+1,
							availableMarketItemsMap[pi.GetMarketItemId()].GetCode(),
							pi.GetPrice(),
							pi.GetVolume(),
							c,
						),
					)
				}

				portfolioItemsInfo = append(
					portfolioItemsInfo,
					fmt.Sprintf("Total cost: %.2f", cost),
				)
				h.Send(
					update.Message.Chat.ID,
					fmt.Sprintf(
						CommandDescriptions[command],
						portfolioPB.GetBaseCurrencyCode(),
						strings.Join(portfolioItemsInfo, "\n"),
					),
				)
				userCache.Command = command

			case CommandPortfolioSummary:
				availableMarketItems, err := h.clnt.GetAvailableMarketItems()
				if err != nil {
					h.Send(update.Message.Chat.ID, ErrorMessage)
					log.Println(err)
					return
				}
				availableMarketItemsMap := make(map[int64]*api.MarketItem, 0)
				for _, mi := range availableMarketItems {
					availableMarketItemsMap[mi.GetId()] = mi
				}

				portfolioPB, err := h.clnt.GetOrCreatePortfolio(userCache.User.GetId())
				if err != nil {
					h.Send(update.Message.Chat.ID, ErrorMessage)
					log.Println(err)
					return
				}

				if len(portfolioPB.GetItems()) == 0 {
					h.Send(update.Message.Chat.ID, NoSufficientDataMessage)
					return
				}

				portfolioMarketItemIds := make([]int64, 0)
				for _, pi := range portfolioPB.GetItems() {
					portfolioMarketItemIds = append(portfolioMarketItemIds, pi.GetMarketItemId())
				}

				lastPrices, err := h.clnt.GetMarketLastPrices(portfolioMarketItemIds)
				if err != nil {
					h.Send(update.Message.Chat.ID, ErrorMessage)
					log.Println(err)
					return
				}

				if len(lastPrices) != len(portfolioMarketItemIds) {
					h.Send(update.Message.Chat.ID, NoSufficientDataMessage)
					return
				}

				portfolioItemsLastPricesMap := make(map[int64]*api.MarketPrice, len(lastPrices))
				lastTs := lastPrices[0].Timestamp
				for _, lmip := range lastPrices {
					portfolioItemsLastPricesMap[lmip.GetMarketItemId()] = lmip
					if lastTs > lmip.Timestamp {
						lastTs = lmip.Timestamp
					}
				}

				tsStart_ := time.Unix(lastTs, 0).UTC()
				differenceInDays := math.Floor(time.Now().UTC().Sub(tsStart_).Hours() / 24)

				portfolioItemsInfo := make([]string, 0)
				if differenceInDays > 0 {
					portfolioItemsInfo = append(
						portfolioItemsInfo,
						fmt.Sprintf("*Info: last updated %s*", tsStart_.Format(time.RFC822)),
					)
				}
				portfolioItemsInfo = append(
					portfolioItemsInfo,
					"#. symbol, price, cost, PNL% (profit)",
				)
				totalOldCost := float64(0)
				totalCurrentCost := float64(0)
				for i, pi := range portfolioPB.GetItems() {
					piPrice := portfolioItemsLastPricesMap[pi.GetMarketItemId()]
					oldCost := pi.GetVolume() * pi.GetPrice()
					currentCost := pi.GetVolume() * piPrice.GetPrice()
					diff := currentCost - oldCost
					diffRatio := diff / oldCost * 100
					totalOldCost += oldCost
					totalCurrentCost += currentCost
					portfolioItemsInfo = append(
						portfolioItemsInfo,
						fmt.Sprintf(
							"%d. %s, %.2f, %.2f, %.2f%% (%0.2f)",
							i+1,
							availableMarketItemsMap[pi.GetMarketItemId()].GetCode(),
							piPrice.GetPrice(),
							currentCost,
							diffRatio,
							diff,
						),
					)
				}
				totalDiff := totalCurrentCost - totalOldCost
				portfolioItemsInfo = append(
					portfolioItemsInfo,
					fmt.Sprintf(
						"Total cost: %.2f, %.2f%% (%0.2f)",
						totalCurrentCost,
						totalDiff/totalOldCost*100,
						totalDiff,
					),
				)

				interval := int64(60*60) * 12
				timeShift := int64(60 * 60 * 24 * 7)
				portfolioItemsPricesMap := make(map[int64][]*api.MarketPrice, 0)
				for _, pmi := range portfolioPB.GetItems() {
					marketItemPrices, err := h.clnt.GetMarketItemPrices(
						pmi.GetMarketItemId(),
						time.Now().Unix()-timeShift,
						time.Now().Unix(),
						interval,
					)
					if err != nil {
						break
					}
					portfolioItemsPricesMap[pmi.GetMarketItemId()] = marketItemPrices
				}

				h.Send(
					update.Message.Chat.ID,
					fmt.Sprintf(
						CommandDescriptions[command],
						portfolioPB.GetBaseCurrencyCode(),
						strings.Join(portfolioItemsInfo, "\n"),
					),
				)

				if len(portfolioItemsPricesMap) == len(portfolioPB.GetItems()) {
					buffer, err := h.plt.PortfolioSummary(totalOldCost, portfolioPB.GetItems(), portfolioItemsPricesMap)
					if err != nil {
						log.Println(err)
					} else {
						file := tgbotapi.FileBytes{
							Name:  "image.png",
							Bytes: buffer.Bytes(),
						}
						h.SendFile(update.Message.Chat.ID, file)
					}
				}
				userCache.Command = command

			case CommandAddOrUpdatePortfolioItem:
				h.Send(
					update.Message.Chat.ID,
					CommandDescriptions[command],
				)
				userCache.Command = command

			case CommandDeletePortfolioItem:
				availableMarketItems, err := h.clnt.GetAvailableMarketItems()
				if err != nil {
					h.Send(update.Message.Chat.ID, ErrorMessage)
					log.Println(err)
					return
				}
				availableMarketItemsMap := make(map[int64]*api.MarketItem, 0)
				for _, mi := range availableMarketItems {
					availableMarketItemsMap[mi.GetId()] = mi
				}

				portfolioPB, err := h.clnt.GetOrCreatePortfolio(userCache.User.GetId())
				if err != nil {
					h.Send(update.Message.Chat.ID, ErrorMessage)
					log.Println(err)
					return
				}
				if len(portfolioPB.GetItems()) == 0 {
					h.Send(
						update.Message.Chat.ID,
						NoSufficientDataMessage,
					)
					return
				}

				portfolioItemsInfo := make([]string, 0)
				cost := float64(0)
				for _, pi := range portfolioPB.GetItems() {
					c := pi.GetPrice() * pi.GetVolume()
					cost += c
					portfolioItemsInfo = append(
						portfolioItemsInfo,
						fmt.Sprintf(
							"%d. %s",
							pi.GetId(),
							availableMarketItemsMap[pi.GetMarketItemId()].GetCode(),
						),
					)
				}

				h.Send(
					update.Message.Chat.ID,
					fmt.Sprintf(
						CommandDescriptions[command],
						strings.Join(portfolioItemsInfo, "\n"),
					),
				)
				userCache.Command = command

			default:
				h.Send(
					update.Message.Chat.ID,
					UnknownCommandMessage,
				)
				userCache.Command = CommandStart

			}

			userCache, err = h.SetUserCached(userCache)

			return
		}

		userText := strings.TrimSpace(update.Message.Text)

		switch userCache.Command {
		case CommandMarket:

			availableMarketItems, err := h.clnt.GetAvailableMarketItems()
			if err != nil {
				h.Send(update.Message.Chat.ID, ErrorMessage)
				log.Println(err)
				return
			}
			availableMarketItemsMap := make(map[string]*api.MarketItem, 0)
			for _, mi := range availableMarketItems {
				availableMarketItemsMap[mi.GetCode()] = mi
			}

			userCode := strings.ToUpper(userText)
			marketItem, ok := availableMarketItemsMap[userCode]
			if !ok {
				h.Send(
					update.Message.Chat.ID,
					NoSufficientDataMessage,
				)
				return
			}
			portfolioPB, err := h.clnt.GetOrCreatePortfolio(userCache.User.GetId())
			if err != nil {
				h.Send(update.Message.Chat.ID, ErrorMessage)
				log.Println(err)
				return
			}

			var portfolioMarketItem *api.PortfolioItem = nil
			for _, pi := range portfolioPB.GetItems() {
				if pi.GetMarketItemId() == marketItem.GetId() {
					portfolioMarketItem = pi
					break
				}
			}

			interval := int64(60 * 60)
			timeShift := int64(60 * 60 * 24 * 7)
			if marketItem.GetType() != domain.MarketItemCryptoCurrencyType {
				timeShift *= 2
				interval *= 12
			}

			marketItemPrices, err := h.clnt.GetMarketItemPrices(marketItem.GetId(),
				time.Now().Unix()-timeShift,
				time.Now().Unix(),
				interval,
			)
			if err != nil {
				h.Send(update.Message.Chat.ID, NoSufficientDataMessage)
				log.Println(err)
				return
			}

			buffer, err := h.plt.MarketItem(marketItem, marketItemPrices, portfolioMarketItem)
			if err != nil {
				h.Send(update.Message.Chat.ID, ErrorMessage)
				log.Println(err)
				return
			}

			file := tgbotapi.FileBytes{
				Name:  "image.png",
				Bytes: buffer.Bytes(),
			}

			h.SendFile(update.Message.Chat.ID, file)

		case CommandAddOrUpdatePortfolioItem:

			availableMarketItems, err := h.clnt.GetAvailableMarketItems()
			if err != nil {
				h.Send(update.Message.Chat.ID, ErrorMessage)
				log.Println(err)
				return
			}
			availableMarketItemsMap := make(map[string]*api.MarketItem, 0)
			for _, mi := range availableMarketItems {
				availableMarketItemsMap[mi.GetCode()] = mi
			}
			portfolioPB, err := h.clnt.GetOrCreatePortfolio(userCache.User.GetId())
			if err != nil {
				h.Send(update.Message.Chat.ID, ErrorMessage)
				log.Println(err)
				return
			}

			splittedText := strings.Split(userText, " ")
			if len(splittedText) != 3 {
				h.Send(
					update.Message.Chat.ID,
					ParsingErrorMessage,
				)
				return
			}
			userCode := strings.ToUpper(splittedText[0])
			marketItem, ok := availableMarketItemsMap[userCode]
			if !ok {
				h.Send(
					update.Message.Chat.ID,
					NoSufficientDataMessage,
				)
				return
			}
			userPrice, err := strconv.ParseFloat(splittedText[1], 64)
			if err != nil {
				h.Send(
					update.Message.Chat.ID,
					ParsingErrorMessage,
				)
				return
			}
			userVolume, err := strconv.ParseFloat(splittedText[2], 64)
			if err != nil {
				h.Send(
					update.Message.Chat.ID,
					ParsingErrorMessage,
				)
				return
			}

			var existingPortfolioItem *api.PortfolioItem = nil
			for _, pi := range portfolioPB.GetItems() {
				if pi.GetMarketItemId() == marketItem.GetId() {
					existingPortfolioItem = pi
				}
			}

			_, err = h.clnt.CreateOrUpdatePortfolioItem(&domain.PortfolioItemCreate{
				PortfolioId:  portfolioPB.GetId(),
				MarketItemId: marketItem.GetId(),
				Price:        userPrice,
				Volume:       userVolume,
			})
			if err != nil {
				h.Send(
					update.Message.Chat.ID,
					ErrorMessage,
				)
				return
			}

			if existingPortfolioItem != nil {
				h.Send(
					update.Message.Chat.ID,
					fmt.Sprintf(
						PortfolioItemUpdated,
						marketItem.GetCode(),
					),
				)
			} else {
				h.Send(
					update.Message.Chat.ID,
					fmt.Sprintf(
						PortfolioItemCreated,
						marketItem.GetCode(),
					),
				)
			}

		case CommandDeletePortfolioItem:
			portfolioPB, err := h.clnt.GetOrCreatePortfolio(userCache.User.GetId())
			if err != nil {
				h.Send(update.Message.Chat.ID, ErrorMessage)
				log.Println(err)
				return
			}

			userPIId, err := strconv.ParseInt(userText, 10, 64)
			if err != nil {
				h.Send(
					update.Message.Chat.ID,
					ParsingErrorMessage,
				)
				return
			}

			var existingPortfolioItem *api.PortfolioItem = nil
			for _, pi := range portfolioPB.GetItems() {
				if pi.GetId() == userPIId {
					existingPortfolioItem = pi
				}
			}

			if existingPortfolioItem == nil {
				h.Send(
					update.Message.Chat.ID,
					ParsingErrorMessage,
				)
				return
			}

			_, err = h.clnt.DeletePortfolioItem(existingPortfolioItem.GetId())
			if err != nil {
				h.Send(
					update.Message.Chat.ID,
					ErrorMessage,
				)
				return
			}

			h.Send(
				update.Message.Chat.ID,
				fmt.Sprintf(
					PortfolioItemDeleted,
					fmt.Sprintf("%d", existingPortfolioItem.GetId()),
				),
			)
		}
	}
}

func (h *TelegramHandler) GetUserCached(userId domain.UserId) (*cache.UserCache, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	channel := h.userCache.Get(userId)
	select {
	case resp := <-channel:
		if resp.Error != nil {
			return nil, resp.Error
		}
		return &resp.Data, nil

	case <-ctx.Done():
		return nil, TimeoutError
	}
}

func (h *TelegramHandler) SetUserCached(userData *cache.UserCache) (*cache.UserCache, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	channel := h.userCache.Set(userData)
	select {
	case resp := <-channel:
		if resp.Error != nil {
			return nil, resp.Error
		}
		return &resp.Data, nil

	case <-ctx.Done():
		return nil, TimeoutError
	}
}
