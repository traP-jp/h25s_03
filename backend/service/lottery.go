package service

import (
	"context"

	"github.com/google/uuid"
)

type LotteryService interface {
	CreateLottery(ctx context.Context, eventID uuid.UUID, lottery LotteryOnCreate) (uuid.UUID, error)
	GetLotteries(ctx context.Context, eventID uuid.UUID, ifDeleted bool) ([]Lottery, error)
	GetLottery(ctx context.Context, eventID uuid.UUID, lotteryID uuid.UUID) (Lottery, error)
	DeleteLottery(ctx context.Context, eventID uuid.UUID, lotteryID uuid.UUID) error
	RollLottery(ctx context.Context, eventID uuid.UUID, lotteryID uuid.UUID, ifDuplicated bool) (string, error)
}
