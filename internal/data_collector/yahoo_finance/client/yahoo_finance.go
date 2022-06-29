package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/abrbird/portfolio_bot/config"
	"github.com/abrbird/portfolio_bot/internal/data_collector/yahoo_finance"
	"github.com/abrbird/portfolio_bot/internal/domain"
	"github.com/abrbird/portfolio_bot/internal/repository"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const ServiceName = "YahooFinance"

type YahooFinanceClient struct {
	apiKey string
	url    string
}

func New(cfg config.DataSource) *YahooFinanceClient {
	return &YahooFinanceClient{
		apiKey: cfg.ApiKey,
		url:    cfg.Url,
	}
}

func (c *YahooFinanceClient) GetHistoricalMap(
	marketItems []domain.MarketItem,
	interval string,
	range_ string,
) (*map[domain.MarketItem]yahoo_finance.Historical, error) {
	cl := http.Client{Timeout: 10 * time.Second}

	itemCodes := make([]string, len(marketItems))
	for i, item := range marketItems {
		if item.Type == domain.MarketItemCryptoCurrencyType {
			itemCodes[i] = fmt.Sprintf("%s-%s", item.Code, "USD")
		} else {
			itemCodes[i] = item.Code
		}
	}

	urlReq := fmt.Sprintf("%s/v8/finance/spark?symbols=%s&interval=%s&range=%s",
		c.url, strings.Join(itemCodes, ","), interval, range_)

	req, err := http.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-API-KEY", c.apiKey)

	r, err := cl.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	historicalMapStr := make(map[string]yahoo_finance.Historical, 0)

	err = json.Unmarshal(b, &historicalMapStr)
	if err != nil {
		log.Print(string(b))
		return nil, err
	}

	historicalMap := make(map[domain.MarketItem]yahoo_finance.Historical, len(historicalMapStr))
	for _, item := range marketItems {
		var codeKey string
		if item.Type == domain.MarketItemCryptoCurrencyType {
			codeKey = fmt.Sprintf("%s-%s", item.Code, "USD")
		} else {
			codeKey = item.Code
		}

		if historical, ok := historicalMapStr[codeKey]; ok {
			historicalMap[item] = historical
		}
	}

	return &historicalMap, nil
}

func Collect(
	availableMarketItems []domain.MarketItem,
	period string,
	range_ string,
	repo repository.MarketPriceRepository,
	clnt Client,
) (success int64, errors int64, err error) {
	historicalMap, err := clnt.GetHistoricalMap(
		availableMarketItems,
		period,
		range_,
	)
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}

	errorsCount := int64(0)
	successCount := int64(0)
	for marketItem, historical := range *historicalMap {
		marketPrices := historical.ToMarketPriceArray(marketItem)
		inserted, err := repo.BulkCreate(context.Background(), marketPrices)
		if err != nil {
			log.Println(err)
		}
		successCount += inserted
		errorsCount += int64(len(*marketPrices)) - inserted
	}
	return successCount, errorsCount, nil
}
