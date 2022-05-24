package server

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-2/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

type Void struct{}

var VOID Void

var (
	ErrorMetadataRetrieve = status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
	ErrorAuthTokenMissed  = status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
	ErrorAuth             = status.Errorf(codes.Unauthenticated, "Authorization if failed")
)

func ErrorLogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	response, err := handler(ctx, req)
	if err != nil {
		log.Println(err)
	}
	return response, err
}

func TimeoutInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	return handler(ctxWithTimeout, req)
}

func AuthInterceptorBuilder(keys config.ClientAPIKeys) func(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	validTokens := map[string]Void{keys.AnyClient: VOID}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := Authorize(ctx, validTokens); err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

func Authorize(ctx context.Context, validTokens map[string]Void) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ErrorMetadataRetrieve
	}
	authHeader, ok := md["authorization"]
	if !ok {
		return ErrorAuthTokenMissed
	}

	if _, ok := validTokens[authHeader[0]]; !ok {
		return ErrorAuth
	}

	return nil
}
