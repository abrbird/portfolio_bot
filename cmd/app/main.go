package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"

	"gitlab.ozon.dev/zBlur/homework-2/config"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository/sql_repository"
	"gitlab.ozon.dev/zBlur/homework-2/internal/server"
	"gitlab.ozon.dev/zBlur/homework-2/internal/service/service_impl"
	"gitlab.ozon.dev/zBlur/homework-2/pkg/api"
)

func main() {
	configFile, err := config.ParseConfig("config/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(configFile.Database.Uri())

	for i, source := range configFile.DataSources {
		fmt.Printf("dataSource %v %v %v \n", i, source.Name, source.Url)
	}

	runServer(configFile)
}

func runServer(configFile *config.File) {

	db, err := server.NewDB(configFile.Database.Uri())
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
		repository,
		service,
	)
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	interceptors := make([]grpc.UnaryServerInterceptor, 0)
	if configFile.Application.ValidateInternal {
		interceptors = append(interceptors, server.AuthInterceptorBuilder(configFile.InternalAPIKeys))
	}
	interceptors = append(interceptors, server.TimeoutInterceptor)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(interceptors...),
	)

	api.RegisterUserServiceServer(grpcServer, newServer)

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
	err = api.RegisterUserServiceHandler(context.Background(), gwmux, conn)
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
