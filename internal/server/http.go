package server

import (
	"github.com/go-eagle/eagle/pkg/app"
	"github.com/go-eagle/eagle/pkg/transport/http"
	"github.com/google/wire"

	v1 "github.com/go-microservice/ins-api/api/micro/user/v1"
	"github.com/go-microservice/ins-api/internal/routers"
	"github.com/go-microservice/ins-api/internal/service"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHTTPServer)

// NewHTTPServer creates a HTTP server
func NewHTTPServer(c *app.ServerConfig, svc *service.UserServiceServer) *http.Server {
	router := routers.NewRouter()

	srv := http.NewServer(
		http.WithAddress(c.Addr),
		http.WithReadTimeout(c.ReadTimeout),
		http.WithWriteTimeout(c.WriteTimeout),
	)

	srv.Handler = router

	v1.RegisterUserServiceHTTPServer(router, svc)

	return srv
}
