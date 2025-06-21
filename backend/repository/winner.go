package repository

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type WinnerRepository interface {
	InsertWinner(ctx echo.Context, eventID uuid.UUID, lotteryID uuid.UUID, userID string) error
}
