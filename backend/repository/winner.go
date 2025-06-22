package repository

import (
	"context"

	"github.com/google/uuid"
)

type WinnerRepository interface {
	InsertWinner(ctx context.Context, winner Winner) error
	GetEventWinners(ctx context.Context, eventID uuid.UUID)([]string,error)
	GetLotteryWinnners(ctx context.Context, lotteryID uuid.UUID)([]string,error)
}
