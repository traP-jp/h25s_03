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
	eventServiceBind = wire.Bind(
		new(service.EventService),
		new(*service.EventServiceImpl),
	)
	attendanceServiceBind = wire.Bind(
		new(service.AttendanceService),
		new(*service.AttendanceServiceImpl),
	)
	lotteryServiceBind = wire.Bind(
		new(service.LotteryService),
		new(*service.LotteryServiceImpl),
	)
	middlewareServiceBind = wire.Bind(
		new(service.MiddlewareService),
		new(*service.MiddlewareServiceImpl),
	)

	adminRepositoryBind = wire.Bind(
		new(repository.AdminRepository),
		new(*repository.AdminRepositoryImpl),
	)
	attendeeRepositoryBind = wire.Bind(
		new(repository.AttendeeRepository),
		new(*repository.AttendeeRepositoryImpl),
	)
	eventRepositoryBind = wire.Bind(
		new(repository.EventRepository),
		new(*repository.EventRepositoryImpl),
	)
	lotteryRepositoryBind = wire.Bind(
		new(repository.LotteryRepository),
		new(*repository.LotteryRepositoryImpl),
	)
	// winnerRepositoryBind = wire.Bind(
	// 	new(repository.WinnerRepository),
	// 	new(*repository.WinnerRepositoryImpl),
	// )
)

func InitializeServer(db *gorm.DB) *handler.Handler {
	wire.Build(
		handler.NewHandler,

		eventServiceBind,
		service.NewEventServiceImpl,
		attendanceServiceBind,
		service.NewAttendanceServiceImpl,
		lotteryServiceBind,
		service.NewLotteryServiceImpl,
		middlewareServiceBind,
		service.NewMiddlewareServiceImpl,

		adminRepositoryBind,
		repository.NewAdminRepositoryImpl,
		attendeeRepositoryBind,
		repository.NewAttendeeRepositoryImpl,
		eventRepositoryBind,
		repository.NewEventRepositoryImpl,
		lotteryRepositoryBind,
		repository.NewLotteryRepositoryImpl,
		// winnerRepositoryBind,
		// repository.NewWinnerRepositoryImpl,
	)

	return &handler.Handler{}
}
