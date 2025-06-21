package service

import (
	"context"
	"fmt"
	"time"

	"github.com/eraxyso/go-template/repository"
	"github.com/google/uuid"
)

type LotteryServiceImpl struct {
	lotteryRepository repository.LotteryRepository
}

var _ LotteryService = &LotteryServiceImpl{}

func NewLotteryServiceImpl(lotteryRepository repository.LotteryRepository) *LotteryServiceImpl {
	return &LotteryServiceImpl{
		lotteryRepository: lotteryRepository,
	}
}

type Lottery struct {
	LotteryId uuid.UUID `json:"lottery_id"`
	EventId   uuid.UUID `json:"event_id"`
	Title     string    `json:"title"`
	IsDeleted bool      `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Winners   []string  `json:"winners"`
}

type LotteryOnCreate struct {
	Title string
}

func (ls *LotteryServiceImpl) CreateLottery(ctx context.Context, eventID uuid.UUID, lottery LotteryOnCreate) (uuid.UUID, error) {
	lotteryOnCreate := repository.LotteryOnCreate{
		Title: lottery.Title,
	}
	createdID, err := ls.lotteryRepository.InsertLottery(ctx, eventID, lotteryOnCreate)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("insert lottery (service): %w", err)
	}
	return createdID, nil
}

func (ls *LotteryServiceImpl) GetLotteries(ctx context.Context, eventID uuid.UUID, ifDeleted bool) ([]Lottery, error) {
	lotteries, err := ls.lotteryRepository.GetLotteries(ctx, eventID, ifDeleted)
	if err != nil {
		return nil, fmt.Errorf("get lotteries (service): %w", err)
	}
	lotteriesResult := make([]Lottery, len(lotteries))
	for i, lottery := range lotteries {
		lotteriesResult[i] = Lottery{
			LotteryId: lottery.LotteryId,
			EventId:   lottery.EventId,
			Title:     lottery.Title,
			IsDeleted: lottery.IsDeleted,
			CreatedAt: lottery.CreatedAt,
			UpdatedAt: lottery.UpdatedAt,
			Winners:   lottery.Winners,
		}
	}
	return lotteriesResult, nil
}

func (ls *LotteryServiceImpl) DeleteLottery(ctx context.Context, eventID uuid.UUID, lotteryID uuid.UUID) error {
	err := ls.lotteryRepository.DeleteLottery(ctx, lotteryID)
	if err != nil {
		return fmt.Errorf("delete lottery (service): %w", err)
	}
	return nil
}
