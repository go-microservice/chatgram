package server

import (
	"github.com/go-eagle/eagle/pkg/app"
	"github.com/go-eagle/eagle/pkg/transport/http"
	"github.com/google/wire"

	momentv1 "github.com/go-microservice/chatgram/api/micro/moment/v1"
	relationv1 "github.com/go-microservice/chatgram/api/micro/relation/v1"
	userv1 "github.com/go-microservice/chatgram/api/micro/user/v1"
	"github.com/go-microservice/chatgram/internal/routers"
	"github.com/go-microservice/chatgram/internal/service"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHTTPServer)

// NewHTTPServer creates a HTTP server
func NewHTTPServer(c *app.ServerConfig,
	userSvc *service.UserServiceServer,
	relSvc *service.RelationServiceServer,
	postSvc *service.PostServiceServer,
	commentSvc *service.CommentServiceServer,
	likeSvc *service.LikeServiceServer,
) *http.Server {
	router := routers.NewRouter()

	srv := http.NewServer(
		http.WithAddress(c.Addr),
		http.WithReadTimeout(c.ReadTimeout),
		http.WithWriteTimeout(c.WriteTimeout),
	)

	srv.Handler = router

	userv1.RegisterUserServiceHTTPServer(router, userSvc)
	relationv1.RegisterRelationServiceHTTPServer(router, relSvc)
	momentv1.RegisterPostServiceHTTPServer(router, postSvc)
	momentv1.RegisterCommentServiceHTTPServer(router, commentSvc)
	momentv1.RegisterLikeServiceHTTPServer(router, likeSvc)

	return srv
}
