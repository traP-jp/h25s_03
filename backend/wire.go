//go:build wireinject

//go:generate wire

package main

import (
	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/eraxyso/go-template/handler"
	"github.com/eraxyso/go-template/repository"
	"github.com/eraxyso/go-template/service"
)

var (
	exampleServiceBind = wire.Bind(
		new(service.ExampleService),
		new(*service.ExampleServiceImpl),
	)
	middlewareServiceBind = wire.Bind(
		new(service.MiddlewareService),
		new(*service.MiddlewareServiceImpl),
	)

	exampleRepositoryBind = wire.Bind(
		new(repository.ExampleRepository),
		new(*repository.ExampleRepositoryImpl),
	)
)

func InitializeServer(db *gorm.DB) *handler.Handler {
	wire.Build(
		handler.NewHandler,

		exampleServiceBind,
		service.NewExampleServiceImpl,
		middlewareServiceBind,
		service.NewMiddlewareServiceImpl,

		exampleRepositoryBind,
		repository.NewExampleRepositoryImpl,
	)

	return &handler.Handler{}
}
