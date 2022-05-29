package repository

import (
	"context"
	"time"

	"github.com/go-eagle/eagle/pkg/client/etcdclient"
	"github.com/go-eagle/eagle/pkg/registry"
	"github.com/go-eagle/eagle/pkg/registry/etcd"
	"github.com/go-eagle/eagle/pkg/transport/grpc"
	"github.com/google/wire"

	momentv1 "github.com/go-microservice/moment-service/api/moment/v1"
	relationV1 "github.com/go-microservice/relation-service/api/relation/v1"
	userv1 "github.com/go-microservice/user-service/api/user/v1"
)

// ProviderSet is repo providers.
var ProviderSet = wire.NewSet(NewUserClient, NewRelationClient)

func getEtcdDiscovery() registry.Discovery {
	// create a etcd register
	client, err := etcdclient.New()
	if err != nil {
		panic(err)
	}
	return etcd.New(client.Client)
}

// TODO
func getConsulDiscovery() registry.Discovery {

	return nil
}

// TODO
func getNacosDiscovery() registry.Discovery {

	return nil
}

func NewUserClient() userv1.UserServiceClient {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	endpoint := "discovery:///user-svc"
	//endpoint := "127.0.0.1:9090"
	conn, err := grpc.DialInsecure(
		ctx,
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(getEtcdDiscovery()),
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

func NewPostClient() momentv1.PostServiceClient {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialInsecure(
		ctx,
		grpc.WithEndpoint("localhost:9092"),
	)
	if err != nil {
		panic(err)
	}
	c := momentv1.NewPostServiceClient(conn)
	return c
}

func NewCommentClient() momentv1.CommentServiceClient {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialInsecure(
		ctx,
		grpc.WithEndpoint("localhost:9093"),
	)
	if err != nil {
		panic(err)
	}
	c := momentv1.NewCommentServiceClient(conn)
	return c
}
