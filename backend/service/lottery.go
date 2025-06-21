package service

import (
	"github.com/eraxyso/go-template/api"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type LotteryService interface {
	GetLotteries(ctx echo.Context, eventID uuid.UUID, ifDeleted bool) ([]api.Lottery, error)
}
