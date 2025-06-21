package service

import (
	"github.com/eraxyso/go-template/api"
	"github.com/eraxyso/go-template/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type LotteryServiceImpl struct {
	lotteryRepository repository.LotteryRepository
}

func NewLotteryServiceImpl(lotteryRepository repository.LotteryRepository) *LotteryServiceImpl {
	return &LotteryServiceImpl{
		lotteryRepository: lotteryRepository,
	}
}

func (ls LotteryServiceImpl) CreateLottery(ctx echo.Context, eventID uuid.UUID, lottery api.PostLotteriesJSONRequestBody) (api.Lottery, error) {
	if _, err := ls.lotteryRepository.InsertLottery(ctx, eventID, lottery); err != nil {
		return api.Lottery{}, err
	}
	return api.Lottery{}, nil
}

func (ls *LotteryServiceImpl) GetLotteries(ctx echo.Context, eventID uuid.UUID, ifDeleted bool) ([]api.Lottery, error) {
	lotteries, err := ls.lotteryRepository.GetLotteries(ctx, eventID, ifDeleted)
	if err != nil {
		return nil, err
	}
	return lotteries, nil
}

func (ls *LotteryServiceImpl) DeleteLottery(ctx echo.Context, eventID uuid.UUID, lotteryID uuid.UUID) error {
	err := ls.lotteryRepository.DeleteLottery(ctx, lotteryID)
	if err != nil {
		return err
	}
	return nil
}
