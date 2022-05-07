package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"gitlab.ozon.dev/zBlur/homework-2/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	client := api.NewUserServiceClient(conn)
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx,
		"sender", "testClient",
		"when", time.Now().Format(time.RFC3339),
		"sender", "route256",
	)

	resp, err := client.RetrieveOrCreate(ctx, &api.CreateUserRequest{
		Id:        int64(123456789),
		UserName:  "TestUserName",
		FirstName: "FN",
		LastName:  "",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Respond: <%v> <%v> <%v> <%v>\n", resp.Id, resp.UserName, resp.FirstName, resp.LastName)
}
