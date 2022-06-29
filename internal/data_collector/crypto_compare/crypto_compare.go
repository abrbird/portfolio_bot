package crypto_compare

import (
	"encoding/json"
	"fmt"
	"github.com/abrbird/portfolio_bot/config"
	"github.com/abrbird/portfolio_bot/internal/domain"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const ServiceName = "CryptoCompare"
const MaxLimit = 2000

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

func (c *Client) GetHistoricalHour(marketItem domain.MarketItem, toCode string, limit int, ts time.Time) (*Historical, error) {
	if limit > MaxLimit {
		limit = MaxLimit
	}
	if limit <= 0 {
		limit = 1
	}

	cl := http.Client{Timeout: 10 * time.Second}

	urlReq := fmt.Sprintf("%s/data/v2/histohour?api_key=%s&limit=%d&fsym=%s&tsym=%s&toTs=%d",
		c.url, c.apiKey, limit, marketItem.Code, toCode, ts.Unix())

	log.Println(urlReq)

	req, err := http.NewRequest(http.MethodPut, urlReq, nil)
	if err != nil {
		return nil, err
	}

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

	historical := Historical{}

	err = json.Unmarshal(b, &historical)
	if err != nil {
		log.Print(string(b))
		return nil, err
	}

	return &historical, nil
}
