package yahoo_finance

import (
	"encoding/json"
	"fmt"
	"gitlab.ozon.dev/zBlur/homework-2/config"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const ServiceName = "YahooFinance"

type Client struct {
	apiKey string
	url    string
}

func New(cfg config.DataSource) *Client {
	return &Client{
		apiKey: cfg.ApiKey,
		url:    cfg.Url,
	}
}

func (c *Client) GetHistoricalMap(
	marketItems []domain.MarketItem,
	interval string,
	range_ string,
) (*map[domain.MarketItem]Historical, error) {
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

	//fmt.Println(urlReq)

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

	historicalMapStr := make(map[string]Historical, 0)

	err = json.Unmarshal(b, &historicalMapStr)
	if err != nil {
		log.Print(string(b))
		return nil, err
	}

	historicalMap := make(map[domain.MarketItem]Historical, len(historicalMapStr))
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
