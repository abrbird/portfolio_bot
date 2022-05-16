package main

import (
	"context"
	"github.com/jasonlvhit/gocron"
	"gitlab.ozon.dev/zBlur/homework-2/config"
	"gitlab.ozon.dev/zBlur/homework-2/internal/data_collector/yahoo_finance"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository/sql_repository"
	"gitlab.ozon.dev/zBlur/homework-2/internal/server"
	"gitlab.ozon.dev/zBlur/homework-2/internal/service"
	"gitlab.ozon.dev/zBlur/homework-2/internal/service/service_impl"
	"log"
	"time"
)

func GetMinTimeStamp(tss []int64, defaultMinTimeStamp int64) int64 {
	minTimeStamp := defaultMinTimeStamp
	for _, item := range tss {
		if minTimeStamp > item {
			minTimeStamp = item
		}
	}
	return minTimeStamp
}

func GetLastTimeStamps(mis []domain.MarketItem, defaultTimeStamp int64, repo repository.Repository) []int64 {
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

func collect(config_ *config.Config, serv service.Service, repo repository.Repository) error {
	log.Println("running task...")

	baseCurrency := repo.Currency().Retrieve(context.Background(), config_.Application.BaseCurrency)
	log.Println("base currency: ", baseCurrency.Currency)

	configMarketItems := config_.Application.GetDomainMarketItems()
	availableMarketItemsRetrieve := serv.MarketItem().RetrieveMany(context.Background(), configMarketItems, repo.MarketItem())

	if availableMarketItemsRetrieve.Error != nil {
		log.Println(availableMarketItemsRetrieve.Error)
		return availableMarketItemsRetrieve.Error
	}
	availableMarketItems := availableMarketItemsRetrieve.MarketItems

	yfDataSource, ok := config_.DataSourcesMap[yahoo_finance.ServiceName]
	if ok && len(availableMarketItems) > 0 {
		cl := yahoo_finance.New(yfDataSource)

		now := time.Now()
		lastTimeStamps := GetLastTimeStamps(
			availableMarketItemsRetrieve.MarketItems,
			config_.Application.HistoryStartTimeStamp,
			repo,
		)
		minTimeStamp := GetMinTimeStamp(lastTimeStamps, now.Unix())

		rangeYF := yahoo_finance.GetRange(minTimeStamp, now.Unix())
		periodYF, err := yahoo_finance.GetInterval(config_.Application.HistoryInterval)
		if err != nil {
			log.Println(err)
			return err
		}

		if len(availableMarketItems) > 0 {
			log.Printf("%s started. range: %s, period: %s\n", yfDataSource.Name, rangeYF, periodYF)

			historicalMap, err := cl.GetHistoricalMap(
				availableMarketItems,
				periodYF,
				rangeYF,
			)
			if err != nil {
				log.Println(err)
				return err
			}

			errorsCount := int64(0)
			successCount := int64(0)
			for marketItem, historical := range *historicalMap {
				marketPrices := historical.ToMarketPriceArray(marketItem)
				inserted, err := repo.MarketPrice().BulkCreate(context.Background(), marketPrices)
				if err != nil {
					log.Println(err)
				}
				successCount += inserted
				errorsCount += int64(len(*marketPrices)) - inserted
			}
			log.Printf("data collection %s done. Success: %d, Errors: %d\n", yfDataSource.Name, successCount, errorsCount)
		} else {
			log.Printf("no available MarketItem")
		}
	}

	return nil
}

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

	err = gocron.Every(config_.Application.HistoryInterval).Seconds().From(gocron.NextTick()).Do(collect, config_, serv, repo)
	if err != nil {
		log.Fatal(err)
	}

	<-gocron.Start()
}
