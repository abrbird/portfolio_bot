package data_collector

import (
	"context"
	"github.com/abrbird/portfolio_bot/config"
	"github.com/abrbird/portfolio_bot/internal/data_collector/yahoo_finance"
	"github.com/abrbird/portfolio_bot/internal/data_collector/yahoo_finance/client"
	"github.com/abrbird/portfolio_bot/internal/domain"
	"github.com/abrbird/portfolio_bot/internal/repository"
	"github.com/abrbird/portfolio_bot/internal/service"
	"log"
	"time"
)

func getMinTimeStamp(tss []int64, defaultMinTimeStamp int64) int64 {
	minTimeStamp := defaultMinTimeStamp
	for _, item := range tss {
		if minTimeStamp > item {
			minTimeStamp = item
		}
	}
	return minTimeStamp
}

func getLastTimeStamps(mis []domain.MarketItem, defaultTimeStamp int64, repo repository.Repository) []int64 {
	lastTimeStamps := make([]int64, len(mis))
	for i, item := range mis {
		lastMarketPriceRetrieved := repo.MarketPrice().RetrieveLast(context.Background(), item.Id)
		if lastMarketPriceRetrieved.Error == nil && lastMarketPriceRetrieved.MarketPrice != nil {
			lastTimeStamps[i] = lastMarketPriceRetrieved.MarketPrice.Timestamp
		} else {
			lastTimeStamps[i] = defaultTimeStamp
		}
	}
	return lastTimeStamps
}

func Collect(config_ *config.Config, serv service.Service, repo repository.Repository) error {
	log.Println("running task...")

	//baseCurrency := repo.Currency().Retrieve(context.Background(), config_.Application.BaseCurrency)
	//log.Println("base currency: ", baseCurrency.Currency)

	configMarketItems := config_.Application.GetDomainMarketItems()
	availableMarketItemsRetrieve := serv.MarketItem().RetrieveMany(context.Background(), configMarketItems, repo.MarketItem())

	if availableMarketItemsRetrieve.Error != nil {
		log.Println(availableMarketItemsRetrieve.Error)
		return availableMarketItemsRetrieve.Error
	}
	availableMarketItems := availableMarketItemsRetrieve.MarketItems

	if len(availableMarketItems) == 0 {
		log.Printf("no available MarketItem")
		return nil
	}

	yfDataSource, ok := config_.DataSourcesMap[client.ServiceName]
	if ok {
		now := time.Now()
		lastTimeStamps := getLastTimeStamps(
			availableMarketItemsRetrieve.MarketItems,
			config_.Application.HistoryStartTimeStamp,
			repo,
		)
		minTimeStamp := getMinTimeStamp(lastTimeStamps, now.Unix())

		rangeYF := yahoo_finance.GetRange(minTimeStamp, now.Unix())
		intervalYF, err := yahoo_finance.GetInterval(config_.Application.HistoryInterval)
		if err != nil {
			log.Println(err)
			return err
		}

		log.Printf("%s started. range: %s, interval: %s\n", yfDataSource.Name, rangeYF, intervalYF)

		marketPriceRepo := repo.MarketPrice()
		cl := client.New(yfDataSource)
		successCount, errorsCount, err := client.Collect(availableMarketItems, intervalYF, rangeYF, marketPriceRepo, cl)

		log.Printf("data collection %s done. Success: %d, Errors: %d\n", yfDataSource.Name, successCount, errorsCount)

		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}
