package handler

import (
	"traquji/api"
	"traquji/service"
)

type Handler struct {
	MiddlewareService service.MiddlewareService
	EventService      service.EventService
	AttendanceService service.AttendanceService
	LotteryService    service.LotteryService
}

var _ api.ServerInterface = &Handler{}

func NewHandler(middlewareService service.MiddlewareService, eventService service.EventService, attendanceService service.AttendanceService, lotteryService service.LotteryService) *Handler {
	return &Handler{
		MiddlewareService: middlewareService,
		EventService:      eventService,
		AttendanceService: attendanceService,
		LotteryService:    lotteryService,
	}
}
