package main

import (
	"github.com/jasonlvhit/gocron"
	"gitlab.ozon.dev/zBlur/homework-2/config"
	"gitlab.ozon.dev/zBlur/homework-2/internal/data_collector/yahoo_finance"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository/sql_repository"
	"gitlab.ozon.dev/zBlur/homework-2/internal/server"
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
		lastMarketPriceRetrieved := repo.MarketPrice().RetrieveLast(item.Id)
		if lastMarketPriceRetrieved.Error == nil && lastMarketPriceRetrieved.MarketPrice != nil {
			lastTimeStamps[i] = lastMarketPriceRetrieved.MarketPrice.Timestamp
		} else {
			lastTimeStamps[i] = defaultTimeStamp
		}
	}
	return lastTimeStamps
}

func collect(config_ *config.Config, repo repository.Repository) error {
	log.Println("running task...")

	baseCurrency := repo.Currency().Retrieve(config_.Application.BaseCurrency)
	log.Println("base currency: ", baseCurrency.Currency)

	marketItemsTypeCodesMap := make(map[string][]string, 0)
	for _, mi := range config_.Application.AvailableMarketItems {
		if _, ok := marketItemsTypeCodesMap[mi.Type]; !ok {
			marketItemsTypeCodesMap[mi.Type] = make([]string, 0)
		}
		marketItemsTypeCodesMap[mi.Type] = append(marketItemsTypeCodesMap[mi.Type], mi.Code)
	}

	availableMarketItems := make([]domain.MarketItem, 0)
	for type_, codes := range marketItemsTypeCodesMap {
		marketItemsRetrieve := repo.MarketItem().RetrieveByType(codes, type_)
		if marketItemsRetrieve.Error != nil {
			log.Println(marketItemsRetrieve.Error)
			return marketItemsRetrieve.Error
		}
		availableMarketItems = append(availableMarketItems, marketItemsRetrieve.MarketItems...)
	}

	if yfDataSource, ok := config_.DataSourcesMap[yahoo_finance.ServiceName]; ok && len(availableMarketItems) > 0 {
		cl, err := yahoo_finance.New(yfDataSource)
		if err != nil {
			log.Println(err)
			return err
		}

		now := time.Now()
		lastTimeStamps := GetLastTimeStamps(availableMarketItems, config_.Application.HistoryStartTimeStamp, repo)
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

			errorsCount := 0
			successCount := 0
			for marketItem, historical := range *historicalMap {
				marketPrices := historical.ToMarketPriceArray(marketItem)
				for _, mp := range *marketPrices {
					err := repo.MarketPrice().Create(&mp)
					if err != nil {
						errorsCount++
						log.Println(err)
					} else {
						successCount++
					}
				}
			}
			log.Printf("data collection %s done. Success: %d, Errors: %d\n", yfDataSource.Name, successCount, errorsCount)
		} else {
			log.Printf("no available MarketItem")
		}
	}

	//if ccDataSource, ok := config_.DataSourcesMap[crypto_compare.ServiceName]; ok {
	//		cl, err := crypto_compare.New(ccDataSource)
	//		if err != nil {
	//			return err
	//		}
	//
	//		historicalHourBTC, err := cl.GetHistoricalHour(
	//			"BTC",
	//			data_collector.DefaultBaseCurrencyCode,
	//			crypto_compare.MaxLimit,
	//			time.Now(),
	//		)
	//		if err != nil {
	//			return err
	//		}
	//
	//		fmt.Println(*historicalHourBTC)
	//}

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
	repo := sql_repository.New(db)

	err = gocron.Every(config_.Application.HistoryInterval).Seconds().From(gocron.NextTick()).Do(collect, config_, repo)
	if err != nil {
		log.Fatal(err)
	}

	<-gocron.Start()
}
