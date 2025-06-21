package handler

import (
	"github.com/eraxyso/go-template/service"
)

type Handler struct {
	EventService service.EventService
	MiddlewareService service.MiddlewareService
	AttendanceService service.AttendanceService
}

func NewHandler(eventService service.EventService, middlewareService service.MiddlewareService,attendanceService service.AttendanceService) *Handler {
	return &Handler{
		EventService:    eventService,
		MiddlewareService: middlewareService,
		AttendanceService: attendanceService,
	}
}
