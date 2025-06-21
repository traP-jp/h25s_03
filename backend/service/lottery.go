package service

import (
	"context"

	"github.com/eraxyso/go-template/api"
	"github.com/google/uuid"
)

type LotteryService interface {
	CreateLottery(ctx context.Context, eventID uuid.UUID, lottery api.PostLotteryJSONBody) (api.Lottery, error)
	GetLotteries(ctx context.Context, eventID uuid.UUID, ifDeleted bool) ([]api.Lottery, error)
	DeleteLottery(ctx context.Context, eventID uuid.UUID, lotteryID uuid.UUID) error
}
