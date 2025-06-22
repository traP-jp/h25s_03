package repository

import (
	"context"

	"github.com/google/uuid"
)

type LotteryRepository interface {
	InsertLottery(ctx context.Context, eventID uuid.UUID, lottery LotteryOnCreate) (uuid.UUID, error)
	GetLotteries(ctx context.Context, eventID uuid.UUID, ifDeleted bool) ([]LotteryWithWinners, error)
	DeleteLottery(ctx context.Context, lotteryID uuid.UUID) error
}
