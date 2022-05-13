package yahoo_finance

import (
	"encoding/json"
	"fmt"
	"gitlab.ozon.dev/zBlur/homework-2/config"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strings"
	"time"
)

const ServiceName = "YahooFinance"

const Interval1M = "1m"
const Interval5M = "5m"
const Interval15M = "15m"
const Interval30M = "30m"
const Interval1H = "1h"
const Interval3H = "3h"
const Interval6H = "6h"
const Interval12H = "12h"
const Interval1D = "1d"
const Interval1WK = "1wk"
const Interval1MO = "1mo"

const Range1D = "1d"
const Range5D = "5d"
const Range1MO = "1mo"
const Range3MO = "3mo"
const Range6MO = "6mo"
const Range1Y = "1y"
const Range5Y = "5y"
const RangeMAX = "max"

func GetRange(tsStart int64, tsEnd int64) string {
	tsStart_ := time.Unix(tsStart, 0)
	tsEnd_ := time.Unix(tsEnd, 0)

	differenceInDays := math.Ceil(tsEnd_.Sub(tsStart_).Hours() / 24)

	switch {
	case differenceInDays < 1:
		return Range1D
	case differenceInDays < 5:
		return Range5D
	case differenceInDays < 30:
		return Range1MO
	case differenceInDays < 30*3:
		return Range3MO
	case differenceInDays < 30*6:
		return Range6MO
	case differenceInDays < 365:
		return Range1Y
	case differenceInDays < 365*5:
		return Range5Y
	default:
		return RangeMAX
	}
}

func GetInterval(seconds uint64) (string, error) {
	switch {
	case seconds == 60:
		return Interval1M, nil
	case seconds == 60*5:
		return Interval5M, nil
	case seconds == 60*15:
		return Interval15M, nil
	case seconds == 60*30:
		return Interval30M, nil
	case seconds == 60*60:
		return Interval1H, nil
	case seconds == 60*60*3:
		return Interval3H, nil
	case seconds == 60*60*6:
		return Interval6H, nil
	case seconds == 60*60*12:
		return Interval12H, nil
	case seconds == 60*60*24:
		return Interval1D, nil
	case seconds == 60*60*24*7:
		return Interval1WK, nil
	case seconds == 60*60*24*30:
		return Interval1MO, nil
	default:
		return "", fmt.Errorf("can not find interval for %d seconds", seconds)
	}
}

type Client struct {
	apiKey string
	url    string
}

func New(cfg config.DataSource) (*Client, error) {
	return &Client{
		apiKey: cfg.ApiKey,
		url:    cfg.Url,
	}, nil
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
