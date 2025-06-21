package repository

import (
	"github.com/eraxyso/go-template/api"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type LotteryRepository interface {
	InsertLottery(ctx echo.Context, eventID uuid.UUID, lottery api.PostLotteryJSONRequestBody) (uuid.UUID, error)
	GetLotteries(ctx echo.Context, eventID uuid.UUID, ifDeleted bool) ([]api.Lottery, error)
	DeleteLottery(ctx echo.Context, lotteryID uuid.UUID) error
}
