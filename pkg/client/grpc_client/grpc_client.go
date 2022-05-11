package grpc_client

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

type Client struct {
	apiKey string
	target string
}

func New(apiKey string, target string) *Client {
	return &Client{
		apiKey: apiKey,
		target: target,
	}
}

func (c *Client) prepareClientContext() (*grpc.ClientConn, api.UserPortfolioServiceClient, context.Context, error) {
	conn, err := grpc.Dial(c.target, grpc.WithInsecure())
	if err != nil {
		return nil, nil, nil, err
	}

	clnt := api.NewUserPortfolioServiceClient(conn)
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx,
		"authorization", c.apiKey,
		"when", time.Now().Format(time.RFC3339),
	)
	return conn, clnt, ctx, nil
}

func (c *Client) GetOrCreateUser(user *domain.User) (*api.User, error) {
	conn, clnt, ctx, err := c.prepareClientContext()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	userPB, err := clnt.RetrieveOrCreateUser(
		ctx,
		&api.CreateUserRequest{
			Id:        user.Id.ToInt64(),
			UserName:  user.UserName,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return userPB, nil
}

func (c *Client) GetOrCreatePortfolio(userId int64) (*api.Portfolio, error) {
	conn, clnt, ctx, err := c.prepareClientContext()
	if err != nil {
		return nil, err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	portfolioPB, err := clnt.RetrieveOrCreatePortfolio(
		ctx,
		&api.CreatePortfolioRequest{UserId: userId},
	)
	if err != nil {
		return nil, err
	}

	return portfolioPB, nil
}

func (c *Client) GetAvailableMarketItems() ([]*api.MarketItem, error) {
	conn, clnt, ctx, err := c.prepareClientContext()
	if err != nil {
		return nil, err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	marketItemsResponse, err := clnt.AvailableMarketItems(
		ctx,
		&api.Empty{},
	)
	if err != nil {
		return nil, err
	}

	return marketItemsResponse.MarketItems, nil
}

func (c *Client) GetMarketItemPrices(
	marketItemId int64,
	startTimeStamp int64,
	endTimeStamp int64,
	interval int64,
) ([]*api.MarketPrice, error) {

	conn, clnt, ctx, err := c.prepareClientContext()
	if err != nil {
		return nil, err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	marketItemPricesResponse, err := clnt.MarketItemsPrices(
		ctx,
		&api.MarketItemPricesRequest{
			MarketItemId:   marketItemId,
			StartTimestamp: startTimeStamp,
			EndTimestamp:   endTimeStamp,
			Interval:       interval,
		},
	)
	if err != nil {
		return nil, err
	}

	return marketItemPricesResponse.MarketPrices, nil
}
