package service

import (
	"github.com/eraxyso/go-template/api"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type LotteryService interface {
	CreateLottery(ctx echo.Context, eventID uuid.UUID, lottery api.PostLotteryJSONBody) (api.Lottery, error)
	GetLotteries(ctx echo.Context, eventID uuid.UUID, ifDeleted bool) ([]api.Lottery, error)
	DeleteLottery(ctx echo.Context, eventID uuid.UUID, lotteryID uuid.UUID) error
}
