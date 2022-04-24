package repository

import (
	"context"
	"time"

	"github.com/go-eagle/eagle/pkg/transport/grpc"
	"github.com/google/wire"

	relationV1 "github.com/go-microservice/relation-service/api/relation/v1"
	userv1 "github.com/go-microservice/user-service/api/user/v1"
)

// ProviderSet is repo providers.
var ProviderSet = wire.NewSet(NewUserClient, NewRelationClient)

func NewUserClient() userv1.UserServiceClient {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialInsecure(
		ctx,
		grpc.WithEndpoint("localhost:9090"),
	)
	if err != nil {
		panic(err)
	}
	c := userv1.NewUserServiceClient(conn)
	return c
}

func NewRelationClient() relationV1.RelationServiceClient {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialInsecure(
		ctx,
		grpc.WithEndpoint("localhost:9091"),
	)
	if err != nil {
		panic(err)
	}
	c := relationV1.NewRelationServiceClient(conn)
	return c
}
