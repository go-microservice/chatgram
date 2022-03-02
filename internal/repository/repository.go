package repository

import (
	"context"

	"github.com/go-eagle/eagle/pkg/transport/grpc"
	"github.com/google/wire"

	userv1 "github.com/go-microservice/user-service/api/user/v1"
)

// ProviderSet is repo providers.
var ProviderSet = wire.NewSet(NewUserClient)

func NewUserClient() userv1.UserServiceClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("localhost:9090"),
	)
	if err != nil {
		panic(err)
	}
	c := userv1.NewUserServiceClient(conn)
	return c
}
