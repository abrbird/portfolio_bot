package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	DefaultTimeout = 10 * time.Second
	Url            = "http://app:8090/v1"
)

type User struct {
	Id        string `json:"Id"`
	UserName  string `json:"UserName"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}

type Client struct{}

func New() *Client {
	return &Client{}
}

func (c *Client) GetOrCreateUser(user *domain.User) (*domain.User, error) {
	cl := http.Client{Timeout: DefaultTimeout}

	urlReq := fmt.Sprintf("%s/users/%d", Url, user.Id)
	postBody, _ := json.Marshal(map[string]string{
		"UserName":  user.UserName,
		"FirstName": user.FirstName,
		"LastName":  user.LastName,
	})

	req, err := http.NewRequest(http.MethodPost, urlReq, bytes.NewBuffer(postBody))
	if err != nil {
		return nil, err
	}

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
		log.Print(string(b))
		return nil, err
	}

	userId, err := strconv.ParseInt(userData.Id, 10, 64)
	if err != nil {
		log.Print(string(b))
		return nil, err
	}

	user = &domain.User{
		Id:        domain.UserId(userId),
		UserName:  userData.UserName,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
	}
	return user, err
}
