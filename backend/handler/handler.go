package handler

import (
	"github.com/eraxyso/go-template/service"
)

type Handler struct {
	EventService service.EventService
	MiddlewareService service.MiddlewareService
}

func NewHandler(eventService service.EventService, middlewareService service.MiddlewareService) *Handler {
	return &Handler{
		EventService:    eventService,
		MiddlewareService: middlewareService,
	}
}
