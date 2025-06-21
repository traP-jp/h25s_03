package service

import (
	"context"

	"github.com/eraxyso/go-template/api"
	"github.com/eraxyso/go-template/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
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

func (ls *LotteryServiceImpl) CreateLottery(ctx context.Context, eventID uuid.UUID, lottery api.PostLotteryJSONRequestBody) (uuid.UUID, error) {
	createdID, err := ls.lotteryRepository.InsertLottery(ctx, eventID, lottery)
	if err != nil {
		return uuid.UUID{}, err
	}
	return createdID, nil
}

func (ls *LotteryServiceImpl) GetLotteries(ctx context.Context, eventID uuid.UUID, ifDeleted bool) ([]api.Lottery, error) {
	lotteries, err := ls.lotteryRepository.GetLotteries(ctx, eventID, ifDeleted)
	if err != nil {
		return nil, err
	}
	return lotteries, nil
}

func (ls *LotteryServiceImpl) DeleteLottery(ctx context.Context, eventID uuid.UUID, lotteryID uuid.UUID) error {
	err := ls.lotteryRepository.DeleteLottery(ctx, lotteryID)
	if err != nil {
		return err
	}
	return nil
}

func (ls *LotteryServiceImpl) CreateLottery(ctx echo.Context, eventID uuid.UUID, requestBody api.PostLotteriesJSONRequestBody) (uuid.UUID, error) {
	u, err := ls.lotteryRepository.InsertLottery(ctx, eventID, requestBody)
	if err != nil {
		return uuid.Nil, err
	}
	return u, nil

}
