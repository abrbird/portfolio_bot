package main

import (
	"context"
	"fmt"
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

	newServer := server.NewServer(
		*config_,
		repository,
		service,
	)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config_.Application.GrpcPort))
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

	log.Printf("Serving gRPC on 0.0.0.0:%s", config_.Application.GrpcPort)
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
		fmt.Sprintf("0.0.0.0:%s", config_.Application.GrpcPort),
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
		Addr:    fmt.Sprintf(":%s", config_.Application.GrpcGatewayPort),
		Handler: gwmux,
	}

	log.Printf("Serving gRPC-Gateway on http://0.0.0.0:%s", config_.Application.GrpcGatewayPort)
	err = gwServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
