package http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/pkg/api"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Client struct {
	apiKey  string
	timeout time.Duration
	Url     string
}

func New(apiKey string, url string, timeout time.Duration) *Client {
	return &Client{
		apiKey:  apiKey,
		timeout: timeout,
		Url:     url,
	}
}

func (c *Client) GetOrCreateUser(user *domain.User) (*api.User, error) {
	cl := http.Client{Timeout: c.timeout}

	urlReq := fmt.Sprintf("%s/users/%d", c.Url, user.Id)
	postBody, _ := json.Marshal(map[string]string{
		"UserName":  user.UserName,
		"FirstName": user.FirstName,
		"LastName":  user.LastName,
	})

	req, err := http.NewRequest(http.MethodPost, urlReq, bytes.NewBuffer(postBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.apiKey)

	r, err := cl.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	userData := User{}
	err = json.Unmarshal(b, &userData)
	if err != nil {
		return nil, err
	}
	userId, err := strconv.ParseInt(userData.Id, 10, 64)
	if err != nil {
		return nil, err
	}

	userApi := api.User{
		Id:        userId,
		UserName:  userData.UserName,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
	}
	return &userApi, err
}

func (c *Client) GetOrCreatePortfolio(userId int64) (*api.Portfolio, error) {
	//TODO implement me
	return nil, nil
}

func (c *Client) CreateOrUpdatePortfolioItem(portfolioItemData *domain.PortfolioItemCreate) (*api.PortfolioItem, error) {
	//TODO implement me
	return nil, nil
}

func (c *Client) DeletePortfolioItem(portfolioItemId int64) (*api.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetAvailableMarketItems() ([]*api.MarketItem, error) {
	//TODO implement me
	return nil, nil
}

func (c *Client) GetMarketItemPrices(marketItemId int64, startTimeStamp int64, endTimeStamp int64, interval int64) ([]*api.MarketPrice, error) {
	//TODO implement me
	return nil, nil
}

func (c *Client) GetMarketLastPrices(marketItemIds []int64) ([]*api.MarketPrice, error) {
	//TODO implement me
	panic("implement me")
}
