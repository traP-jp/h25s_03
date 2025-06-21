package handler

import (
	"github.com/eraxyso/go-template/service"
)

type Handler struct {
	ExampleService    service.ExampleService
	MiddlewareService service.MiddlewareService
}

func NewHandler(exampleService service.ExampleService, middlewareService service.MiddlewareService) *Handler {
	return &Handler{
		ExampleService:    exampleService,
		MiddlewareService: middlewareService,
	}
}
