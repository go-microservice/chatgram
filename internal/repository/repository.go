package repository

import (
	"context"
	"time"

	"github.com/go-eagle/eagle/pkg/client/consulclient"
	"github.com/go-eagle/eagle/pkg/client/etcdclient"
	"github.com/go-eagle/eagle/pkg/client/nacosclient"
	"github.com/go-eagle/eagle/pkg/registry"
	"github.com/go-eagle/eagle/pkg/registry/consul"
	"github.com/go-eagle/eagle/pkg/registry/etcd"
	"github.com/go-eagle/eagle/pkg/registry/nacos"
	"github.com/go-eagle/eagle/pkg/transport/grpc"
	"github.com/google/wire"

	momentv1 "github.com/go-microservice/moment-service/api/moment/v1"
	relationV1 "github.com/go-microservice/relation-service/api/relation/v1"
	userv1 "github.com/go-microservice/user-service/api/user/v1"
)

// ProviderSet is repo providers.
var ProviderSet = wire.NewSet(NewUserClient, NewRelationClient, NewPostClient, NewCommentClient, NewLikeClient)

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
	client, err := consulclient.New()
	if err != nil {
		panic(err)
	}
	return consul.New(client)
}

// NOTE: endpoint is special, with suffix grpc
// example: endpoint := "discovery:///user-svc.grpc"
func getNacosDiscovery() registry.Discovery {
	client, err := nacosclient.New()
	if err != nil {
		panic(err)
	}
	return nacos.New(client)
}

func NewUserClient() userv1.UserServiceClient {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// etcd
	endpoint := "discovery:///user-svc"
	conn, err := grpc.DialInsecure(
		ctx,
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(getConsulDiscovery()),
		grpc.WithLog(),
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

	endpoint := "discovery:///relation-svc"
	conn, err := grpc.DialInsecure(
		ctx,
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(getConsulDiscovery()),
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

	endpoint := "discovery:///moment-svc"
	conn, err := grpc.DialInsecure(
		ctx,
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(getConsulDiscovery()),
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

	endpoint := "discovery:///moment-svc"
	conn, err := grpc.DialInsecure(
		ctx,
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(getConsulDiscovery()),
	)
	if err != nil {
		panic(err)
	}
	c := momentv1.NewCommentServiceClient(conn)
	return c
}

func NewLikeClient() momentv1.LikeServiceClient {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	endpoint := "discovery:///moment-svc"
	conn, err := grpc.DialInsecure(
		ctx,
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(getConsulDiscovery()),
	)
	if err != nil {
		panic(err)
	}
	c := momentv1.NewLikeServiceClient(conn)
	return c
}
