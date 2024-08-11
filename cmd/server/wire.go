//go:build wireinject
// +build wireinject

package main

import (
	"fmt"
	"log"
	"net/http"

	eagle "github.com/go-eagle/eagle/pkg/app"
	logger "github.com/go-eagle/eagle/pkg/log"
	transHttp "github.com/go-eagle/eagle/pkg/transport/http"
	"github.com/go-microservice/chatgram/internal/repository"
	"github.com/go-microservice/chatgram/internal/server"
	"github.com/go-microservice/chatgram/internal/service"
	"github.com/google/wire"
)

func InitApp(cfg *eagle.Config, config *eagle.ServerConfig) (*eagle.App, error) {
	wire.Build(server.ProviderSet, service.ProviderSet, repository.ProviderSet, newApp)
	return &eagle.App{}, nil
}

func newApp(cfg *eagle.Config, httpSrv *transHttp.Server) *eagle.App {
	// init pprof server
	go func() {
		fmt.Printf("Listening and serving PProf HTTP on %s\n", cfg.PprofPort)
		if err := http.ListenAndServe(cfg.PprofPort, http.DefaultServeMux); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen ListenAndServe for PProf, err: %s", err.Error())
		}
	}()

	return eagle.New(
		eagle.WithName(cfg.Name),
		eagle.WithVersion(cfg.Version),
		eagle.WithLogger(logger.GetLogger()),
		eagle.WithServer(httpSrv),
	)
}
