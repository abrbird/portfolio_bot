package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"gitlab.ozon.dev/zBlur/homework-2/config"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository/sql_repository"
	"gitlab.ozon.dev/zBlur/homework-2/internal/server"
	"gitlab.ozon.dev/zBlur/homework-2/internal/service/service_impl"
	"gitlab.ozon.dev/zBlur/homework-2/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

func main() {
	config_, err := config.ParseConfig("config/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(config_.Database.Uri())

	runServer(config_)
}

func runServer(config_ *config.Config) {

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

	repository := sql_repository.New(db)
	service := service_impl.New()

	//startInterval :=
	//endInterval :=
	//
	//intervals := service.MarketPrice().GetIntervals(startInterval, endInterval, )
	//// 9300 - btc, 4094 - aapl
	//marketPricesRetrieved := service.MarketPrice().RetrieveInterval(4094, startInterval, endInterval, repository.MarketPrice())
	//if marketPricesRetrieved.Error != nil {
	//	log.Fatal(err)
	//}
	//marketPrices, blanksCount := service.MarketPrice().FillBlanks(marketPricesRetrieved.MarketPrices, intervals)
	//
	//blanksRatio := float64(blanksCount) / float64(len(intervals))
	//
	//fmt.Println(blanksRatio)
	//fmt.Println(intervals[len(intervals)-5:])
	//fmt.Println(marketPrices[len(marketPrices)-5:])

	//d, err := service.MarketPrice().GetMarketItemPrices(9300, int64(1648771200), time.Now().Unix(), config_.Application.HistoryInterval, repository.MarketPrice())
	//if err != nil {
	//
	//}

	newServer := server.NewServer(
		*config_,
		repository,
		service,
	)
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	interceptors := make([]grpc.UnaryServerInterceptor, 0)
	if config_.Application.ValidateInternal {
		interceptors = append(interceptors, server.AuthInterceptorBuilder(config_.ClientAPIKeys))
	}
	interceptors = append(interceptors, server.TimeoutInterceptor, server.ErrorLogInterceptor)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(interceptors...),
	)

	api.RegisterUserPortfolioServiceServer(grpcServer, newServer)

	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		err = grpcServer.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(server.HeaderMatcher),
	)
	err = api.RegisterUserPortfolioServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	err = gwServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
