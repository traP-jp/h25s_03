package repository

import (
	"github.com/eraxyso/go-template/api"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type LotteryRepository interface {
	InsertLottery(ctx echo.Context, eventID uuid.UUID, lottery api.PostLotteriesJSONRequestBody) (uuid.UUID, error)
	GetLotteries(ctx echo.Context, eventID uuid.UUID) ([]api.Lottery, error)
	RemoveLottery(ctx echo.Context, eventID uuid.UUID, lotteryID uuid.UUID) error
}
