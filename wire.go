// +build wireinject

package main

import (
	eagle "github.com/go-eagle/eagle/pkg/app"
	"github.com/go-microservice/ins-api/internal/server"
	"github.com/go-microservice/ins-api/internal/service"
	"github.com/go-microservice/ins-api/internal/repository"
	"github.com/google/wire"
)

func InitApp(cfg *eagle.Config, config *eagle.ServerConfig) (*eagle.App, error) {
	wire.Build(server.ProviderSet, service.ProviderSet, repository.ProviderSet, newApp)
	return &eagle.App{}, nil
}
